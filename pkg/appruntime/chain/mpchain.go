/*
Copyright 2023 KubeAGI.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package chain

import (
	"context"
	"errors"
	"fmt"

	"github.com/tmc/langchaingo/chains"
	"github.com/tmc/langchaingo/llms"
	"github.com/tmc/langchaingo/prompts"
	langchainschema "github.com/tmc/langchaingo/schema"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/apimachinery/pkg/runtime/schema"
	"k8s.io/client-go/dynamic"
	"k8s.io/klog/v2"

	"github.com/kubeagi/arcadia/api/app-node/chain/v1alpha1"
	"github.com/kubeagi/arcadia/pkg/appruntime/base"
)

const (
	DefaultPromptTemplateForMap = `
		{{.context}}

		With above content, please summarize it with only half content size of it.
		`
	DefaultPromptTemplatForReduce       = `"{{.context}}"`
	DefaultSummaryMaxNumberOfConcurrent = 2

	DefaultDocumentChunkSize    = 2048
	DefaultDocumentChunkOverlap = 200
)

type MapReduceChain struct {
	base.BaseNode
	mpChain  chains.MapReduceDocuments
	Instance *v1alpha1.LLMChain
}

func NewMapReduceChain(baseNode base.BaseNode) *MapReduceChain {
	return &MapReduceChain{
		mpChain:  chains.MapReduceDocuments{},
		BaseNode: baseNode,
	}
}
func (l *MapReduceChain) Init(ctx context.Context, cli dynamic.Interface, _ map[string]any) error {
	ns := base.GetAppNamespace(ctx)
	instance := &v1alpha1.LLMChain{}
	obj, err := cli.Resource(schema.GroupVersionResource{Group: v1alpha1.GroupVersion.Group, Version: v1alpha1.GroupVersion.Version, Resource: "llmchains"}).
		Namespace(l.Ref.GetNamespace(ns)).Get(ctx, l.Ref.Name, metav1.GetOptions{})
	if err != nil {
		return fmt.Errorf("can't find the chain in cluster: %w", err)
	}
	err = runtime.DefaultUnstructuredConverter.FromUnstructured(obj.UnstructuredContent(), instance)
	if err != nil {
		return fmt.Errorf("can't convert obj to LLMChain: %w", err)
	}
	l.Instance = instance
	return nil
}

func (l *MapReduceChain) Run(ctx context.Context, cli dynamic.Interface, args map[string]any) (outArgs map[string]any, err error) {
	v1, ok := args["llm"]
	if !ok {
		return args, errors.New("no llm")
	}
	llm, ok := v1.(llms.LLM)
	if !ok {
		return args, errors.New("llm not llms.LanguageModel")
	}

	// chain options
	instance := l.Instance
	options := GetChainOptions(instance.Spec.CommonChainConfig)

	// initialize map reduce chain
	mpChain := chains.NewMapReduceDocuments(
		chains.NewLLMChain(llm, prompts.NewPromptTemplate(DefaultPromptTemplateForMap, []string{"context"})),
		chains.NewStuffDocuments(
			chains.NewLLMChain(
				llm,
				prompts.NewPromptTemplate(DefaultPromptTemplatForReduce, []string{"context"}),
			),
		),
	)
	mpChain.MaxNumberOfConcurrent = DefaultSummaryMaxNumberOfConcurrent

	l.mpChain = mpChain
	// parse documents
	documents, ok := args["_documents"].([]langchainschema.Document)
	if !ok {
		return args, errors.New("_documents not schema.Document")
	}

	var out string
	needStream := false
	needStream, ok = args["_need_stream"].(bool)
	if ok && needStream {
		options = append(options, chains.WithStreamingFunc(stream(args)))
		out, err = chains.Run(ctx, mpChain, documents, options...)
	} else {
		if len(options) > 0 {
			out, err = chains.Run(ctx, mpChain, documents, options...)
		} else {
			out, err = chains.Run(ctx, mpChain, documents)
		}
	}

	out, err = handleNoErrNoOut(ctx, needStream, out, err, l.mpChain, args, options)
	klog.FromContext(ctx).V(5).Info("use mpchain, blocking out:" + out)
	if err == nil {
		args["_answer"] = out
		return args, nil
	}
	return args, fmt.Errorf("mpchain run error: %w", err)
}
