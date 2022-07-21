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

type AssociationObservation struct {
	ID *string `json:"id,omitempty" tf:"id,omitempty"`
}

type AssociationParameters struct {

	// +kubebuilder:validation:Required
	LicenseConfigurationArn *string `json:"licenseConfigurationArn" tf:"license_configuration_arn,omitempty"`

	// Region is the region you'd like your resource to be created in.
	// +terrajet:crd:field:TFTag=-
	// +kubebuilder:validation:Required
	Region *string `json:"region" tf:"-"`

	// +kubebuilder:validation:Required
	ResourceArn *string `json:"resourceArn" tf:"resource_arn,omitempty"`
}

// AssociationSpec defines the desired state of Association
type AssociationSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     AssociationParameters `json:"forProvider"`
}

// AssociationStatus defines the observed state of Association.
type AssociationStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        AssociationObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// Association is the Schema for the Associations API
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,awsjet}
type Association struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              AssociationSpec   `json:"spec"`
	Status            AssociationStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// AssociationList contains a list of Associations
type AssociationList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []Association `json:"items"`
}

// Repository type metadata.
var (
	Association_Kind             = "Association"
	Association_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: Association_Kind}.String()
	Association_KindAPIVersion   = Association_Kind + "." + CRDGroupVersion.String()
	Association_GroupVersionKind = CRDGroupVersion.WithKind(Association_Kind)
)

func init() {
	SchemeBuilder.Register(&Association{}, &AssociationList{})
}
