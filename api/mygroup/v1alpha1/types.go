package v1alpha1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

type PodCreator struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec PodCreatorSpec `json:"spec"`
}

type PodCreatorSpec struct {
	Name   string            `json:"name"`
	Labels map[string]string `json:"labels"`
}

type PodCreatorList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata"`

	Items []PodCreator `json:"items"`
}
