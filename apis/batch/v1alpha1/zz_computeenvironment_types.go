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

type ComputeEnvironmentObservation struct {
	Arn *string `json:"arn,omitempty" tf:"arn,omitempty"`

	EcsClusterArn *string `json:"ecsClusterArn,omitempty" tf:"ecs_cluster_arn,omitempty"`

	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	Status *string `json:"status,omitempty" tf:"status,omitempty"`

	StatusReason *string `json:"statusReason,omitempty" tf:"status_reason,omitempty"`

	TagsAll map[string]*string `json:"tagsAll,omitempty" tf:"tags_all,omitempty"`
}

type ComputeEnvironmentParameters struct {

	// +kubebuilder:validation:Optional
	ComputeEnvironmentName *string `json:"computeEnvironmentName,omitempty" tf:"compute_environment_name,omitempty"`

	// +kubebuilder:validation:Optional
	ComputeEnvironmentNamePrefix *string `json:"computeEnvironmentNamePrefix,omitempty" tf:"compute_environment_name_prefix,omitempty"`

	// +kubebuilder:validation:Optional
	ComputeResources []ComputeResourcesParameters `json:"computeResources,omitempty" tf:"compute_resources,omitempty"`

	// Region is the region you'd like your resource to be created in.
	// +terrajet:crd:field:TFTag=-
	// +kubebuilder:validation:Required
	Region *string `json:"region" tf:"-"`

	// +kubebuilder:validation:Optional
	ServiceRole *string `json:"serviceRole,omitempty" tf:"service_role,omitempty"`

	// +kubebuilder:validation:Optional
	State *string `json:"state,omitempty" tf:"state,omitempty"`

	// +kubebuilder:validation:Optional
	Tags map[string]*string `json:"tags,omitempty" tf:"tags,omitempty"`

	// +kubebuilder:validation:Required
	Type *string `json:"type" tf:"type,omitempty"`
}

type ComputeResourcesObservation struct {
}

type ComputeResourcesParameters struct {

	// +kubebuilder:validation:Optional
	AllocationStrategy *string `json:"allocationStrategy,omitempty" tf:"allocation_strategy,omitempty"`

	// +kubebuilder:validation:Optional
	BidPercentage *float64 `json:"bidPercentage,omitempty" tf:"bid_percentage,omitempty"`

	// +kubebuilder:validation:Optional
	DesiredVcpus *float64 `json:"desiredVcpus,omitempty" tf:"desired_vcpus,omitempty"`

	// +kubebuilder:validation:Optional
	EC2KeyPair *string `json:"ec2KeyPair,omitempty" tf:"ec2_key_pair,omitempty"`

	// +kubebuilder:validation:Optional
	ImageID *string `json:"imageId,omitempty" tf:"image_id,omitempty"`

	// +kubebuilder:validation:Optional
	InstanceRole *string `json:"instanceRole,omitempty" tf:"instance_role,omitempty"`

	// +kubebuilder:validation:Optional
	InstanceType []*string `json:"instanceType,omitempty" tf:"instance_type,omitempty"`

	// +kubebuilder:validation:Optional
	LaunchTemplate []LaunchTemplateParameters `json:"launchTemplate,omitempty" tf:"launch_template,omitempty"`

	// +kubebuilder:validation:Required
	MaxVcpus *float64 `json:"maxVcpus" tf:"max_vcpus,omitempty"`

	// +kubebuilder:validation:Optional
	MinVcpus *float64 `json:"minVcpus,omitempty" tf:"min_vcpus,omitempty"`

	// +kubebuilder:validation:Required
	SecurityGroupIds []*string `json:"securityGroupIds" tf:"security_group_ids,omitempty"`

	// +kubebuilder:validation:Optional
	SpotIAMFleetRole *string `json:"spotIamFleetRole,omitempty" tf:"spot_iam_fleet_role,omitempty"`

	// +kubebuilder:validation:Required
	Subnets []*string `json:"subnets" tf:"subnets,omitempty"`

	// +kubebuilder:validation:Optional
	Tags map[string]*string `json:"tags,omitempty" tf:"tags,omitempty"`

	// +kubebuilder:validation:Required
	Type *string `json:"type" tf:"type,omitempty"`
}

type LaunchTemplateObservation struct {
}

type LaunchTemplateParameters struct {

	// +kubebuilder:validation:Optional
	LaunchTemplateID *string `json:"launchTemplateId,omitempty" tf:"launch_template_id,omitempty"`

	// +kubebuilder:validation:Optional
	LaunchTemplateName *string `json:"launchTemplateName,omitempty" tf:"launch_template_name,omitempty"`

	// +kubebuilder:validation:Optional
	Version *string `json:"version,omitempty" tf:"version,omitempty"`
}

// ComputeEnvironmentSpec defines the desired state of ComputeEnvironment
type ComputeEnvironmentSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     ComputeEnvironmentParameters `json:"forProvider"`
}

// ComputeEnvironmentStatus defines the observed state of ComputeEnvironment.
type ComputeEnvironmentStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        ComputeEnvironmentObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// ComputeEnvironment is the Schema for the ComputeEnvironments API
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,awsjet}
type ComputeEnvironment struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              ComputeEnvironmentSpec   `json:"spec"`
	Status            ComputeEnvironmentStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// ComputeEnvironmentList contains a list of ComputeEnvironments
type ComputeEnvironmentList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []ComputeEnvironment `json:"items"`
}

// Repository type metadata.
var (
	ComputeEnvironment_Kind             = "ComputeEnvironment"
	ComputeEnvironment_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: ComputeEnvironment_Kind}.String()
	ComputeEnvironment_KindAPIVersion   = ComputeEnvironment_Kind + "." + CRDGroupVersion.String()
	ComputeEnvironment_GroupVersionKind = CRDGroupVersion.WithKind(ComputeEnvironment_Kind)
)

func init() {
	SchemeBuilder.Register(&ComputeEnvironment{}, &ComputeEnvironmentList{})
}
