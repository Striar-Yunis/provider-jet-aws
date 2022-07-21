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

type APNSVoIPSandboxChannelObservation struct {
	ID *string `json:"id,omitempty" tf:"id,omitempty"`
}

type APNSVoIPSandboxChannelParameters struct {

	// +kubebuilder:validation:Required
	ApplicationID *string `json:"applicationId" tf:"application_id,omitempty"`

	// +kubebuilder:validation:Optional
	BundleIDSecretRef *v1.SecretKeySelector `json:"bundleIdSecretRef,omitempty" tf:"-"`

	// +kubebuilder:validation:Optional
	CertificateSecretRef *v1.SecretKeySelector `json:"certificateSecretRef,omitempty" tf:"-"`

	// +kubebuilder:validation:Optional
	DefaultAuthenticationMethod *string `json:"defaultAuthenticationMethod,omitempty" tf:"default_authentication_method,omitempty"`

	// +kubebuilder:validation:Optional
	Enabled *bool `json:"enabled,omitempty" tf:"enabled,omitempty"`

	// +kubebuilder:validation:Optional
	PrivateKeySecretRef *v1.SecretKeySelector `json:"privateKeySecretRef,omitempty" tf:"-"`

	// Region is the region you'd like your resource to be created in.
	// +terrajet:crd:field:TFTag=-
	// +kubebuilder:validation:Required
	Region *string `json:"region" tf:"-"`

	// +kubebuilder:validation:Optional
	TeamIDSecretRef *v1.SecretKeySelector `json:"teamIdSecretRef,omitempty" tf:"-"`

	// +kubebuilder:validation:Optional
	TokenKeyIDSecretRef *v1.SecretKeySelector `json:"tokenKeyIdSecretRef,omitempty" tf:"-"`

	// +kubebuilder:validation:Optional
	TokenKeySecretRef *v1.SecretKeySelector `json:"tokenKeySecretRef,omitempty" tf:"-"`
}

// APNSVoIPSandboxChannelSpec defines the desired state of APNSVoIPSandboxChannel
type APNSVoIPSandboxChannelSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     APNSVoIPSandboxChannelParameters `json:"forProvider"`
}

// APNSVoIPSandboxChannelStatus defines the observed state of APNSVoIPSandboxChannel.
type APNSVoIPSandboxChannelStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        APNSVoIPSandboxChannelObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// APNSVoIPSandboxChannel is the Schema for the APNSVoIPSandboxChannels API
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,awsjet}
type APNSVoIPSandboxChannel struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              APNSVoIPSandboxChannelSpec   `json:"spec"`
	Status            APNSVoIPSandboxChannelStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// APNSVoIPSandboxChannelList contains a list of APNSVoIPSandboxChannels
type APNSVoIPSandboxChannelList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []APNSVoIPSandboxChannel `json:"items"`
}

// Repository type metadata.
var (
	APNSVoIPSandboxChannel_Kind             = "APNSVoIPSandboxChannel"
	APNSVoIPSandboxChannel_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: APNSVoIPSandboxChannel_Kind}.String()
	APNSVoIPSandboxChannel_KindAPIVersion   = APNSVoIPSandboxChannel_Kind + "." + CRDGroupVersion.String()
	APNSVoIPSandboxChannel_GroupVersionKind = CRDGroupVersion.WithKind(APNSVoIPSandboxChannel_Kind)
)

func init() {
	SchemeBuilder.Register(&APNSVoIPSandboxChannel{}, &APNSVoIPSandboxChannelList{})
}
