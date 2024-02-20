/*
Copyright 2024 KubeAGI.

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
	"bytes"
	"context"
	"errors"
	"fmt"
	"io"
	"mime/multipart"
	"path/filepath"

	rbacv1 "k8s.io/api/rbac/v1"

	"github.com/kubeagi/arcadia/pkg/appruntime/base"
	"github.com/tmc/langchaingo/chains"
	"github.com/tmc/langchaingo/documentloaders"
	"github.com/tmc/langchaingo/llms"
	"github.com/tmc/langchaingo/prompts"
	"github.com/tmc/langchaingo/textsplitter"
	"k8s.io/client-go/dynamic"
	"k8s.io/klog/v2"
	"sigs.k8s.io/controller-runtime/pkg/client"
)

const (
	DefaultPromptTemplateForMap = `
		{{.context}}

		With above content, please summarize it with only half content size of it.
		`
	DefaultPromptTemplatForReduce       = `"{{.context}}"`
	DefaultSummaryMaxNumberOfConcurrent = 2

	DefaultDocumentChunkSize    = 1024
	DefaultDocumentChunkOverlap = 100
)

type MapReduceChain struct {
	base.BaseNode

	rbacv1.Role

	chains.MapReduceDocuments
	chainCallOptions []chains.ChainCallOption
}

func NewMapReduceChain(baseNode base.BaseNode) *MapReduceChain {
	return &MapReduceChain{
		MapReduceDocuments: chains.MapReduceDocuments{},
		BaseNode:           baseNode,
	}
}

func (l *MapReduceChain) Init(ctx context.Context, cli client.Client, args map[string]any) error {
	if args == nil {
		return errors.New("no arguments provided for MapReduceChain")
	}

	var chainCallOptions []chains.ChainCallOption
	switch kind := l.BaseNode.Kind(); kind {
	case "llmchain":
		llmchain := NewLLMChain(l.BaseNode)
		if err := llmchain.Init(ctx, cli, nil); err != nil {
			return err
		}
		chainCallOptions = GetChainOptions(llmchain.Instance.Spec.CommonChainConfig)
	case "retrievalqachain":
		retrivalQAChain := NewRetrievalQAChain(l.BaseNode)
		if err := retrivalQAChain.Init(ctx, cli, nil); err != nil {
			return err
		}
		chainCallOptions = GetChainOptions(retrivalQAChain.Instance.Spec.CommonChainConfig)
	default:
		return fmt.Errorf("invalid basenode kind %s for MapReduceChain", kind)
	}
	l.chainCallOptions = append(l.chainCallOptions, chainCallOptions...)

	v1, ok := args["llm"]
	if !ok {
		return errors.New("no llm")
	}
	llm, ok := v1.(llms.LLM)
	if !ok {
		return errors.New("llm not llms.LanguageModel")
	}

	l.MapReduceDocuments = chains.NewMapReduceDocuments(
		chains.NewLLMChain(llm, prompts.NewPromptTemplate(DefaultPromptTemplateForMap, []string{"context"})),
		chains.NewStuffDocuments(
			chains.NewLLMChain(
				llm,
				prompts.NewPromptTemplate(DefaultPromptTemplatForReduce, []string{"context"}),
			),
		),
	)

	return nil
}

func (l *MapReduceChain) Run(ctx context.Context, cli dynamic.Interface, args map[string]any) (outArgs map[string]any, err error) {
	v1, ok := args["documents"]
	if !ok {
		return args, errors.New("no llm")
	}
	doc, ok := v1.(*multipart.FileHeader)
	if !ok {
		return args, errors.New("document not *multipart.FileHeader")
	}

	var chunkSize, chunkOverlap int
	v2, ok := args["chunk_size"]
	if !ok {
		chunkSize = DefaultDocumentChunkSize
	} else {
		chunkSize, ok = v2.(int)
		if !ok {
			chunkSize = DefaultDocumentChunkSize
		}
	}
	v3, ok := args["chunk_overlap"]
	if !ok {
		chunkOverlap = DefaultDocumentChunkOverlap
	} else {
		chunkOverlap, ok = v3.(int)
		if !ok {
			chunkOverlap = DefaultDocumentChunkOverlap
		}
	}

	src, err := doc.Open()
	if err != nil {
		return nil, err
	}
	defer src.Close()
	data, err := io.ReadAll(src)
	if err != nil {
		return nil, err
	}
	dataReader := bytes.NewReader(data)

	var loader documentloaders.Loader
	switch ext := filepath.Ext(doc.Filename); ext {
	case ".pdf":
		loader = documentloaders.NewPDF(dataReader, doc.Size)
	case ".txt":
		loader = documentloaders.NewText(dataReader)
	case ".html", ".htm":
		loader = documentloaders.NewHTML(dataReader)
	default:
		return nil, fmt.Errorf("file with extension %s not supported yet", ext)
	}

	split := textsplitter.NewRecursiveCharacter(
		textsplitter.WithChunkSize(chunkSize),
		textsplitter.WithChunkOverlap(chunkOverlap),
	)
	documents, err := loader.LoadAndSplit(ctx, split)
	if err != nil {
		return nil, err
	}

	needStream := false
	needStream, ok = args["_need_stream"].(bool)
	if ok && needStream {
		l.chainCallOptions = append(l.chainCallOptions, chains.WithStreamingFunc(stream(args)))
	}

	out, err := chains.Run(ctx, l.MapReduceDocuments, documents, l.chainCallOptions...)
	if err != nil {
		return args, fmt.Errorf("failed to run MapReduceChain due to %s", err.Error())
	}
	klog.FromContext(ctx).V(5).Info("use MapReduceChain, blocking out:" + out)
	args["_answer"] = out

	return args, nil
}
