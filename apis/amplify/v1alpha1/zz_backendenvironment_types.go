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

type BackendEnvironmentObservation struct {
	Arn *string `json:"arn,omitempty" tf:"arn,omitempty"`

	ID *string `json:"id,omitempty" tf:"id,omitempty"`
}

type BackendEnvironmentParameters struct {

	// +kubebuilder:validation:Required
	AppID *string `json:"appId" tf:"app_id,omitempty"`

	// +kubebuilder:validation:Optional
	DeploymentArtifacts *string `json:"deploymentArtifacts,omitempty" tf:"deployment_artifacts,omitempty"`

	// +kubebuilder:validation:Required
	EnvironmentName *string `json:"environmentName" tf:"environment_name,omitempty"`

	// Region is the region you'd like your resource to be created in.
	// +terrajet:crd:field:TFTag=-
	// +kubebuilder:validation:Required
	Region *string `json:"region" tf:"-"`

	// +kubebuilder:validation:Optional
	StackName *string `json:"stackName,omitempty" tf:"stack_name,omitempty"`
}

// BackendEnvironmentSpec defines the desired state of BackendEnvironment
type BackendEnvironmentSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     BackendEnvironmentParameters `json:"forProvider"`
}

// BackendEnvironmentStatus defines the observed state of BackendEnvironment.
type BackendEnvironmentStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        BackendEnvironmentObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// BackendEnvironment is the Schema for the BackendEnvironments API
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,awsjet}
type BackendEnvironment struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              BackendEnvironmentSpec   `json:"spec"`
	Status            BackendEnvironmentStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// BackendEnvironmentList contains a list of BackendEnvironments
type BackendEnvironmentList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []BackendEnvironment `json:"items"`
}

// Repository type metadata.
var (
	BackendEnvironment_Kind             = "BackendEnvironment"
	BackendEnvironment_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: BackendEnvironment_Kind}.String()
	BackendEnvironment_KindAPIVersion   = BackendEnvironment_Kind + "." + CRDGroupVersion.String()
	BackendEnvironment_GroupVersionKind = CRDGroupVersion.WithKind(BackendEnvironment_Kind)
)

func init() {
	SchemeBuilder.Register(&BackendEnvironment{}, &BackendEnvironmentList{})
}
