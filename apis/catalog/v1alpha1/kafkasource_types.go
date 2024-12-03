/*
Copyright 2022 The Crossplane Authors.

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
	"reflect"

	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime/schema"

	xpv1 "github.com/crossplane/crossplane-runtime/apis/common/v1"
)

// KafkaSourceParameters are the configurable fields of a KafkaSource.
type KafkaSourceParameters struct {
	ConfigurableField string `json:"configurableField"`
}

// KafkaSourceObservation are the observable fields of a KafkaSource.
type KafkaSourceObservation struct {
	ObservableField string `json:"observableField,omitempty"`
}

// A KafkaSourceSpec defines the desired state of a KafkaSource.
type KafkaSourceSpec struct {
	xpv1.ResourceSpec `json:",inline"`
	ForProvider       KafkaSourceParameters `json:"forProvider"`
}

// A KafkaSourceStatus represents the observed state of a KafkaSource.
type KafkaSourceStatus struct {
	xpv1.ResourceStatus `json:",inline"`
	AtProvider          KafkaSourceObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// A KafkaSource is an example API type.
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,collateprovider}
type KafkaSource struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   KafkaSourceSpec   `json:"spec"`
	Status KafkaSourceStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// KafkaSourceList contains a list of KafkaSource
type KafkaSourceList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []KafkaSource `json:"items"`
}

// KafkaSource type metadata.
var (
	KafkaSourceKind             = reflect.TypeOf(KafkaSource{}).Name()
	KafkaSourceGroupKind        = schema.GroupKind{Group: Group, Kind: KafkaSourceKind}.String()
	KafkaSourceKindAPIVersion   = KafkaSourceKind + "." + SchemeGroupVersion.String()
	KafkaSourceGroupVersionKind = SchemeGroupVersion.WithKind(KafkaSourceKind)
)

func init() {
	SchemeBuilder.Register(&KafkaSource{}, &KafkaSourceList{})
}
