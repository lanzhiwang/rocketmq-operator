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
	corev1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// EDIT THIS FILE!  THIS IS SCAFFOLDING FOR YOU TO OWN!
// NOTE: json tags are required.  Any new fields you add must have json tags for the fields to be serialized.

// BrokerSpec defines the desired state of Broker
type BrokerSpec struct {
	// INSERT ADDITIONAL SPEC FIELDS - desired state of cluster
	// Important: Run "make" to regenerate code after modifying this file

	Size int `json:"size"`
	// NameServers defines the name service list e.g. 192.168.1.1:9876;192.168.1.2:9876
	NameServers string `json:"nameServers"`
	// ReplicaPerGroup each broker cluster's replica number
	ReplicaPerGroup int `json:"replicaPerGroup"`
	// BaseImage is the broker image to use for the Pods
	BrokerImage string `json:"brokerImage"`
	// ImagePullPolicy defines how the image is pulled
	ImagePullPolicy corev1.PullPolicy `json:"imagePullPolicy"`
	// AllowRestart defines whether allow pod restart
	AllowRestart bool `json:"allowRestart"`
	// Resources describes the compute resource requirements
	Resources corev1.ResourceRequirements `json:"resources"`
	// StorageMode can be EmptyDir, HostPath, StorageClass
	StorageMode string `json:"storageMode"`
	// HostPath is the local path to store data
	HostPath string `json:"hostPath,omitempty"`
	// Env defines custom env, e.g. BROKER_MEM
	Env []corev1.EnvVar `json:"env"`
	// Volumes define the broker.conf
	Volumes []corev1.Volume `json:"volumes"`
	// VolumeClaimTemplates defines the StorageClass
	Volume `json:"volume,omitempty"`
	// The name of pod where the metadata from
	ScalePodName string `json:"scalePodName"`
}

// BrokerStatus defines the observed state of Broker
type BrokerStatus struct {
	// INSERT ADDITIONAL STATUS FIELD - define observed state of cluster
	// Important: Run "make" to regenerate code after modifying this file
	Nodes []string `json:"nodes"`
	Size  int      `json:"size"`
}

//+kubebuilder:object:root=true
//+kubebuilder:subresource:status

// Broker is the Schema for the brokers API
type Broker struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   BrokerSpec   `json:"spec,omitempty"`
	Status BrokerStatus `json:"status,omitempty"`
}

//+kubebuilder:object:root=true

// BrokerList contains a list of Broker
type BrokerList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []Broker `json:"items"`
}

func init() {
	SchemeBuilder.Register(&Broker{}, &BrokerList{})
}
