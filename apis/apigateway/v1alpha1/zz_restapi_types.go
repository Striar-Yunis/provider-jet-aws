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

type RestAPIEndpointConfigurationObservation struct {
}

type RestAPIEndpointConfigurationParameters struct {

	// +kubebuilder:validation:Required
	Types []*string `json:"types" tf:"types,omitempty"`

	// +kubebuilder:validation:Optional
	VPCEndpointIds []*string `json:"vpcEndpointIds,omitempty" tf:"vpc_endpoint_ids,omitempty"`
}

type RestAPIObservation struct {
	Arn *string `json:"arn,omitempty" tf:"arn,omitempty"`

	CreatedDate *string `json:"createdDate,omitempty" tf:"created_date,omitempty"`

	ExecutionArn *string `json:"executionArn,omitempty" tf:"execution_arn,omitempty"`

	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	RootResourceID *string `json:"rootResourceId,omitempty" tf:"root_resource_id,omitempty"`

	TagsAll map[string]*string `json:"tagsAll,omitempty" tf:"tags_all,omitempty"`
}

type RestAPIParameters struct {

	// +kubebuilder:validation:Optional
	APIKeySource *string `json:"apiKeySource,omitempty" tf:"api_key_source,omitempty"`

	// +kubebuilder:validation:Optional
	BinaryMediaTypes []*string `json:"binaryMediaTypes,omitempty" tf:"binary_media_types,omitempty"`

	// +kubebuilder:validation:Optional
	Body *string `json:"body,omitempty" tf:"body,omitempty"`

	// +kubebuilder:validation:Optional
	Description *string `json:"description,omitempty" tf:"description,omitempty"`

	// +kubebuilder:validation:Optional
	DisableExecuteAPIEndpoint *bool `json:"disableExecuteApiEndpoint,omitempty" tf:"disable_execute_api_endpoint,omitempty"`

	// +kubebuilder:validation:Optional
	EndpointConfiguration []RestAPIEndpointConfigurationParameters `json:"endpointConfiguration,omitempty" tf:"endpoint_configuration,omitempty"`

	// +kubebuilder:validation:Optional
	MinimumCompressionSize *float64 `json:"minimumCompressionSize,omitempty" tf:"minimum_compression_size,omitempty"`

	// +kubebuilder:validation:Required
	Name *string `json:"name" tf:"name,omitempty"`

	// +kubebuilder:validation:Optional
	Parameters map[string]*string `json:"parameters,omitempty" tf:"parameters,omitempty"`

	// +kubebuilder:validation:Optional
	Policy *string `json:"policy,omitempty" tf:"policy,omitempty"`

	// Region is the region you'd like your resource to be created in.
	// +terrajet:crd:field:TFTag=-
	// +kubebuilder:validation:Required
	Region *string `json:"region" tf:"-"`

	// +kubebuilder:validation:Optional
	Tags map[string]*string `json:"tags,omitempty" tf:"tags,omitempty"`
}

// RestAPISpec defines the desired state of RestAPI
type RestAPISpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     RestAPIParameters `json:"forProvider"`
}

// RestAPIStatus defines the observed state of RestAPI.
type RestAPIStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        RestAPIObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// RestAPI is the Schema for the RestAPIs API
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,awsjet}
type RestAPI struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              RestAPISpec   `json:"spec"`
	Status            RestAPIStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// RestAPIList contains a list of RestAPIs
type RestAPIList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []RestAPI `json:"items"`
}

// Repository type metadata.
var (
	RestAPI_Kind             = "RestAPI"
	RestAPI_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: RestAPI_Kind}.String()
	RestAPI_KindAPIVersion   = RestAPI_Kind + "." + CRDGroupVersion.String()
	RestAPI_GroupVersionKind = CRDGroupVersion.WithKind(RestAPI_Kind)
)

func init() {
	SchemeBuilder.Register(&RestAPI{}, &RestAPIList{})
}
