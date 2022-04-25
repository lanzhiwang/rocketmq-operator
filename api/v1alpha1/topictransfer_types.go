/*
Copyright 2022.

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
)

// EDIT THIS FILE!  THIS IS SCAFFOLDING FOR YOU TO OWN!
// NOTE: json tags are required.  Any new fields you add must have json tags for the fields to be serialized.

// TopicTransferSpec defines the desired state of TopicTransfer
type TopicTransferSpec struct {
	// INSERT ADDITIONAL SPEC FIELDS - desired state of cluster
	// Important: Run "make" to regenerate code after modifying this file

	// Topic name
	Topic string `json:"topic,omitempty"`
	// The cluster where the transferred topic from
	SourceCluster string `json:"sourceCluster,omitempty"`
	// The cluster where the topic will be transferred to
	TargetCluster string `json:"targetCluster,omitempty"`
}

// TopicTransferStatus defines the observed state of TopicTransfer
type TopicTransferStatus struct {
	// INSERT ADDITIONAL STATUS FIELD - define observed state of cluster
	// Important: Run "make" to regenerate code after modifying this file
}

//+kubebuilder:object:root=true
//+kubebuilder:subresource:status

// TopicTransfer is the Schema for the topictransfers API
type TopicTransfer struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   TopicTransferSpec   `json:"spec,omitempty"`
	Status TopicTransferStatus `json:"status,omitempty"`
}

//+kubebuilder:object:root=true

// TopicTransferList contains a list of TopicTransfer
type TopicTransferList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []TopicTransfer `json:"items"`
}

func init() {
	SchemeBuilder.Register(&TopicTransfer{}, &TopicTransferList{})
}
