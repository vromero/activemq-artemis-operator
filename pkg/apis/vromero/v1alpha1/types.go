package v1alpha1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

type ArtemisClusterList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata"`
	Items           []ArtemisCluster `json:"items"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

type ArtemisCluster struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata"`
	Spec              ArtemisClusterSpec   `json:"spec"`
	Status            ArtemisClusterStatus `json:"status,omitempty"`
}

type ArtemisClusterSpec struct {
	Version string
	Variant string
	Size int32
}
type ArtemisClusterStatus struct {
	// Fill me
}
