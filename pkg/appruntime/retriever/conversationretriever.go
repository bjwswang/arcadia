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

package retriever

import (
	"context"
	"errors"
	"fmt"

	langchaingoschema "github.com/tmc/langchaingo/schema"
	"github.com/tmc/langchaingo/vectorstores"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/apimachinery/pkg/runtime/schema"
	"k8s.io/client-go/dynamic"

	apiretriever "github.com/kubeagi/arcadia/api/app-node/retriever/v1alpha1"
	"github.com/kubeagi/arcadia/apiserver/pkg/common"
	"github.com/kubeagi/arcadia/pkg/appruntime/base"
	"github.com/kubeagi/arcadia/pkg/langchainwrap"
	pkgvectorstore "github.com/kubeagi/arcadia/pkg/vectorstore"
)

const (
	DefaultNumberOfDocs = 5
	DefaultThreshold    = 0.3
)

// ConversationRetriever retrive documents from vectorstore with conversatioin id
type ConversationRetriever struct {
	base.BaseNode
	instance *apiretriever.ConversationRetriever

	langchaingoschema.Retriever
	DocNullReturn string
}

func NewConversationRetriever(baseNode base.BaseNode) *ConversationRetriever {
	return &ConversationRetriever{
		BaseNode: baseNode,
	}
}

func (l *ConversationRetriever) Init(ctx context.Context, cli dynamic.Interface, _ map[string]any) error {
	ns := base.GetAppNamespace(ctx)
	instance := &apiretriever.ConversationRetriever{}
	obj, err := cli.Resource(schema.GroupVersionResource{Group: apiretriever.GroupVersion.Group, Version: apiretriever.GroupVersion.Version, Resource: "conversationretrievers"}).
		Namespace(l.Ref.GetNamespace(ns)).Get(ctx, l.Ref.Name, metav1.GetOptions{})
	if err != nil {
		return fmt.Errorf("can't find the conversation retriever in cluster: %w", err)
	}
	err = runtime.DefaultUnstructuredConverter.FromUnstructured(obj.UnstructuredContent(), instance)
	if err != nil {
		return fmt.Errorf("can't convert obj to ConversationRetriever: %w", err)
	}
	l.instance = instance
	return nil
}

func (l *ConversationRetriever) Run(ctx context.Context, cli dynamic.Interface, args map[string]any) (map[string]any, error) {
	conversationID, ok := args["_conversation_id"].(string)
	if !ok {
		return nil, errors.New("no conversation id")
	}
	// conversation retriever utlize our system embedding suite(Embedder,VectorStore)
	embedder, vectorStore, err := common.SystemEmbeddingSuite(ctx, cli)
	if err != nil {
		return nil, err
	}
	langchainEmbedder, err := langchainwrap.GetLangchainEmbedder(ctx, embedder, nil, cli, "")
	if err != nil {
		return nil, err
	}
	var s vectorstores.VectorStore
	s, _, err = pkgvectorstore.NewVectorStore(ctx, vectorStore, langchainEmbedder, conversationID, nil, cli)
	if err != nil {
		return nil, err
	}
	l.Retriever = vectorstores.ToRetriever(s, l.instance.Spec.NumDocuments, vectorstores.WithScoreThreshold(l.instance.Spec.ScoreThreshold))

	// conversation retriever
	args["conversation_retriever"] = l

	return args, nil
}
