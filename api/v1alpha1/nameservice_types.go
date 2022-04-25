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

// NameServiceSpec defines the desired state of NameService
type NameServiceSpec struct {
	// INSERT ADDITIONAL SPEC FIELDS - desired state of cluster
	// Important: Run "make" to regenerate code after modifying this file

	// Size is the number of the name service Pod
	Size int32 `json:"size"`
	//NameServiceImage is the name service image
	NameServiceImage string `json:"nameServiceImage"`
	// ImagePullPolicy defines how the image is pulled.
	ImagePullPolicy corev1.PullPolicy `json:"imagePullPolicy"`
	// HostNetwork can be true or false
	HostNetwork bool `json:"hostNetwork"`
	// dnsPolicy defines how a pod's DNS will be configured
	DNSPolicy corev1.DNSPolicy `json:"dnsPolicy"`
	// Resources describes the compute resource requirements
	Resources corev1.ResourceRequirements `json:"resources"`
	// StorageMode can be EmptyDir, HostPath, StorageClass
	StorageMode string `json:"storageMode"`
	// HostPath is the local path to store data
	HostPath string `json:"hostPath,omitempty"`
	Volume   `json:"volume,omitempty"`
}

// NameServiceStatus defines the observed state of NameService
type NameServiceStatus struct {
	// INSERT ADDITIONAL STATUS FIELD - define observed state of cluster
	// Important: Run "make" to regenerate code after modifying this file
	NameServers []string `json:"nameServers"`
}

//+kubebuilder:object:root=true
//+kubebuilder:subresource:status

// NameService is the Schema for the nameservices API
type NameService struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   NameServiceSpec   `json:"spec,omitempty"`
	Status NameServiceStatus `json:"status,omitempty"`
}

//+kubebuilder:object:root=true

// NameServiceList contains a list of NameService
type NameServiceList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []NameService `json:"items"`
}

func init() {
	SchemeBuilder.Register(&NameService{}, &NameServiceList{})
}
