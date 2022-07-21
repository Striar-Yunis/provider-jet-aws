/*
Copyright 2021 The Crossplane Authors.

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

// Code generated by terrajet. DO NOT EDIT.

package v1alpha1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime/schema"

	v1 "github.com/crossplane/crossplane-runtime/apis/common/v1"
)

type APIDestinationObservation struct {
	Arn *string `json:"arn,omitempty" tf:"arn,omitempty"`

	ID *string `json:"id,omitempty" tf:"id,omitempty"`
}

type APIDestinationParameters struct {

	// +kubebuilder:validation:Required
	ConnectionArn *string `json:"connectionArn" tf:"connection_arn,omitempty"`

	// +kubebuilder:validation:Optional
	Description *string `json:"description,omitempty" tf:"description,omitempty"`

	// +kubebuilder:validation:Required
	HTTPMethod *string `json:"httpMethod" tf:"http_method,omitempty"`

	// +kubebuilder:validation:Required
	InvocationEndpoint *string `json:"invocationEndpoint" tf:"invocation_endpoint,omitempty"`

	// +kubebuilder:validation:Optional
	InvocationRateLimitPerSecond *float64 `json:"invocationRateLimitPerSecond,omitempty" tf:"invocation_rate_limit_per_second,omitempty"`

	// +kubebuilder:validation:Required
	Name *string `json:"name" tf:"name,omitempty"`

	// Region is the region you'd like your resource to be created in.
	// +terrajet:crd:field:TFTag=-
	// +kubebuilder:validation:Required
	Region *string `json:"region" tf:"-"`
}

// APIDestinationSpec defines the desired state of APIDestination
type APIDestinationSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     APIDestinationParameters `json:"forProvider"`
}

// APIDestinationStatus defines the observed state of APIDestination.
type APIDestinationStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        APIDestinationObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// APIDestination is the Schema for the APIDestinations API
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,awsjet}
type APIDestination struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              APIDestinationSpec   `json:"spec"`
	Status            APIDestinationStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// APIDestinationList contains a list of APIDestinations
type APIDestinationList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []APIDestination `json:"items"`
}

// Repository type metadata.
var (
	APIDestination_Kind             = "APIDestination"
	APIDestination_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: APIDestination_Kind}.String()
	APIDestination_KindAPIVersion   = APIDestination_Kind + "." + CRDGroupVersion.String()
	APIDestination_GroupVersionKind = CRDGroupVersion.WithKind(APIDestination_Kind)
)

func init() {
	SchemeBuilder.Register(&APIDestination{}, &APIDestinationList{})
}
