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

type TrafficMirrorTargetObservation struct {
	Arn *string `json:"arn,omitempty" tf:"arn,omitempty"`

	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	OwnerID *string `json:"ownerId,omitempty" tf:"owner_id,omitempty"`

	TagsAll map[string]*string `json:"tagsAll,omitempty" tf:"tags_all,omitempty"`
}

type TrafficMirrorTargetParameters struct {

	// +kubebuilder:validation:Optional
	Description *string `json:"description,omitempty" tf:"description,omitempty"`

	// +kubebuilder:validation:Optional
	NetworkInterfaceID *string `json:"networkInterfaceId,omitempty" tf:"network_interface_id,omitempty"`

	// +kubebuilder:validation:Optional
	NetworkLoadBalancerArn *string `json:"networkLoadBalancerArn,omitempty" tf:"network_load_balancer_arn,omitempty"`

	// Region is the region you'd like your resource to be created in.
	// +terrajet:crd:field:TFTag=-
	// +kubebuilder:validation:Required
	Region *string `json:"region" tf:"-"`

	// +kubebuilder:validation:Optional
	Tags map[string]*string `json:"tags,omitempty" tf:"tags,omitempty"`
}

// TrafficMirrorTargetSpec defines the desired state of TrafficMirrorTarget
type TrafficMirrorTargetSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     TrafficMirrorTargetParameters `json:"forProvider"`
}

// TrafficMirrorTargetStatus defines the observed state of TrafficMirrorTarget.
type TrafficMirrorTargetStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        TrafficMirrorTargetObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// TrafficMirrorTarget is the Schema for the TrafficMirrorTargets API
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,awsjet}
type TrafficMirrorTarget struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              TrafficMirrorTargetSpec   `json:"spec"`
	Status            TrafficMirrorTargetStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// TrafficMirrorTargetList contains a list of TrafficMirrorTargets
type TrafficMirrorTargetList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []TrafficMirrorTarget `json:"items"`
}

// Repository type metadata.
var (
	TrafficMirrorTarget_Kind             = "TrafficMirrorTarget"
	TrafficMirrorTarget_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: TrafficMirrorTarget_Kind}.String()
	TrafficMirrorTarget_KindAPIVersion   = TrafficMirrorTarget_Kind + "." + CRDGroupVersion.String()
	TrafficMirrorTarget_GroupVersionKind = CRDGroupVersion.WithKind(TrafficMirrorTarget_Kind)
)

func init() {
	SchemeBuilder.Register(&TrafficMirrorTarget{}, &TrafficMirrorTargetList{})
}
