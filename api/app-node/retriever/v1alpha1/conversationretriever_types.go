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

package v1alpha1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"

	node "github.com/kubeagi/arcadia/api/app-node"
	"github.com/kubeagi/arcadia/api/base/v1alpha1"
)

// EDIT THIS FILE!  THIS IS SCAFFOLDING FOR YOU TO OWN!
// NOTE: json tags are required.  Any new fields you add must have json tags for the fields to be serialized.

// ConversationRetrieverSpec defines the desired state of ConversationRetriever
type ConversationRetrieverSpec struct {
	v1alpha1.CommonSpec   `json:",inline"`
	CommonRetrieverConfig `json:",inline"`
}

// ConversationRetrieverStatus defines the observed state of ConversationRetriever
type ConversationRetrieverStatus struct {
	// ObservedGeneration is the last observed generation.
	// +optional
	ObservedGeneration int64 `json:"observedGeneration,omitempty"`

	// ConditionedStatus is the current status
	v1alpha1.ConditionedStatus `json:",inline"`
}

//+kubebuilder:object:root=true
//+kubebuilder:subresource:status

// ConversationRetriever is the Schema for the conversationretrievers API
type ConversationRetriever struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   ConversationRetrieverSpec   `json:"spec,omitempty"`
	Status ConversationRetrieverStatus `json:"status,omitempty"`
}

//+kubebuilder:object:root=true

// ConversationRetrieverList contains a list of ConversationRetriever
type ConversationRetrieverList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []ConversationRetriever `json:"items"`
}

func init() {
	SchemeBuilder.Register(&ConversationRetriever{}, &ConversationRetrieverList{})
}

var _ node.Node = (*ConversationRetriever)(nil)

// SetRef for conversation retriever
func (c *ConversationRetriever) SetRef() {
	//
	annotations := node.SetRefAnnotations(c.GetAnnotations(), []node.Ref{node.InputRef.Len(1)}, []node.Ref{node.LLMChainRef.Len(1), node.RetrievalQAChainRef.Len(1)})
	if c.GetAnnotations() == nil {
		c.SetAnnotations(annotations)
	}
	for k, v := range annotations {
		c.Annotations[k] = v
	}
}
