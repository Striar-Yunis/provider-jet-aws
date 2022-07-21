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

type XSSMatchSetObservation struct {
	Arn *string `json:"arn,omitempty" tf:"arn,omitempty"`

	ID *string `json:"id,omitempty" tf:"id,omitempty"`
}

type XSSMatchSetParameters struct {

	// +kubebuilder:validation:Required
	Name *string `json:"name" tf:"name,omitempty"`

	// Region is the region you'd like your resource to be created in.
	// +terrajet:crd:field:TFTag=-
	// +kubebuilder:validation:Required
	Region *string `json:"region" tf:"-"`

	// +kubebuilder:validation:Optional
	XSSMatchTuples []XSSMatchTuplesParameters `json:"xssMatchTuples,omitempty" tf:"xss_match_tuples,omitempty"`
}

type XSSMatchTuplesFieldToMatchObservation struct {
}

type XSSMatchTuplesFieldToMatchParameters struct {

	// +kubebuilder:validation:Optional
	Data *string `json:"data,omitempty" tf:"data,omitempty"`

	// +kubebuilder:validation:Required
	Type *string `json:"type" tf:"type,omitempty"`
}

type XSSMatchTuplesObservation struct {
}

type XSSMatchTuplesParameters struct {

	// +kubebuilder:validation:Required
	FieldToMatch []XSSMatchTuplesFieldToMatchParameters `json:"fieldToMatch" tf:"field_to_match,omitempty"`

	// +kubebuilder:validation:Required
	TextTransformation *string `json:"textTransformation" tf:"text_transformation,omitempty"`
}

// XSSMatchSetSpec defines the desired state of XSSMatchSet
type XSSMatchSetSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     XSSMatchSetParameters `json:"forProvider"`
}

// XSSMatchSetStatus defines the observed state of XSSMatchSet.
type XSSMatchSetStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        XSSMatchSetObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// XSSMatchSet is the Schema for the XSSMatchSets API
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,awsjet}
type XSSMatchSet struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              XSSMatchSetSpec   `json:"spec"`
	Status            XSSMatchSetStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// XSSMatchSetList contains a list of XSSMatchSets
type XSSMatchSetList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []XSSMatchSet `json:"items"`
}

// Repository type metadata.
var (
	XSSMatchSet_Kind             = "XSSMatchSet"
	XSSMatchSet_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: XSSMatchSet_Kind}.String()
	XSSMatchSet_KindAPIVersion   = XSSMatchSet_Kind + "." + CRDGroupVersion.String()
	XSSMatchSet_GroupVersionKind = CRDGroupVersion.WithKind(XSSMatchSet_Kind)
)

func init() {
	SchemeBuilder.Register(&XSSMatchSet{}, &XSSMatchSetList{})
}
