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

type ActionDefinitionObservation struct {
}

type ActionDefinitionParameters struct {

	// +kubebuilder:validation:Required
	PublishMetricAction []PublishMetricActionParameters `json:"publishMetricAction" tf:"publish_metric_action,omitempty"`
}

type DimensionObservation struct {
}

type DimensionParameters struct {

	// +kubebuilder:validation:Required
	Value *string `json:"value" tf:"value,omitempty"`
}

type FirewallPolicyFirewallPolicyObservation struct {
}

type FirewallPolicyFirewallPolicyParameters struct {

	// +kubebuilder:validation:Optional
	StatefulRuleGroupReference []StatefulRuleGroupReferenceParameters `json:"statefulRuleGroupReference,omitempty" tf:"stateful_rule_group_reference,omitempty"`

	// +kubebuilder:validation:Optional
	StatelessCustomAction []StatelessCustomActionParameters `json:"statelessCustomAction,omitempty" tf:"stateless_custom_action,omitempty"`

	// +kubebuilder:validation:Required
	StatelessDefaultActions []*string `json:"statelessDefaultActions" tf:"stateless_default_actions,omitempty"`

	// +kubebuilder:validation:Required
	StatelessFragmentDefaultActions []*string `json:"statelessFragmentDefaultActions" tf:"stateless_fragment_default_actions,omitempty"`

	// +kubebuilder:validation:Optional
	StatelessRuleGroupReference []StatelessRuleGroupReferenceParameters `json:"statelessRuleGroupReference,omitempty" tf:"stateless_rule_group_reference,omitempty"`
}

type FirewallPolicyObservation struct {
	Arn *string `json:"arn,omitempty" tf:"arn,omitempty"`

	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	TagsAll map[string]*string `json:"tagsAll,omitempty" tf:"tags_all,omitempty"`

	UpdateToken *string `json:"updateToken,omitempty" tf:"update_token,omitempty"`
}

type FirewallPolicyParameters struct {

	// +kubebuilder:validation:Optional
	Description *string `json:"description,omitempty" tf:"description,omitempty"`

	// +kubebuilder:validation:Required
	FirewallPolicy []FirewallPolicyFirewallPolicyParameters `json:"firewallPolicy" tf:"firewall_policy,omitempty"`

	// +kubebuilder:validation:Required
	Name *string `json:"name" tf:"name,omitempty"`

	// Region is the region you'd like your resource to be created in.
	// +terrajet:crd:field:TFTag=-
	// +kubebuilder:validation:Required
	Region *string `json:"region" tf:"-"`

	// +kubebuilder:validation:Optional
	Tags map[string]*string `json:"tags,omitempty" tf:"tags,omitempty"`
}

type PublishMetricActionObservation struct {
}

type PublishMetricActionParameters struct {

	// +kubebuilder:validation:Required
	Dimension []DimensionParameters `json:"dimension" tf:"dimension,omitempty"`
}

type StatefulRuleGroupReferenceObservation struct {
}

type StatefulRuleGroupReferenceParameters struct {

	// +kubebuilder:validation:Required
	ResourceArn *string `json:"resourceArn" tf:"resource_arn,omitempty"`
}

type StatelessCustomActionObservation struct {
}

type StatelessCustomActionParameters struct {

	// +kubebuilder:validation:Required
	ActionDefinition []ActionDefinitionParameters `json:"actionDefinition" tf:"action_definition,omitempty"`

	// +kubebuilder:validation:Required
	ActionName *string `json:"actionName" tf:"action_name,omitempty"`
}

type StatelessRuleGroupReferenceObservation struct {
}

type StatelessRuleGroupReferenceParameters struct {

	// +kubebuilder:validation:Required
	Priority *float64 `json:"priority" tf:"priority,omitempty"`

	// +kubebuilder:validation:Required
	ResourceArn *string `json:"resourceArn" tf:"resource_arn,omitempty"`
}

// FirewallPolicySpec defines the desired state of FirewallPolicy
type FirewallPolicySpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     FirewallPolicyParameters `json:"forProvider"`
}

// FirewallPolicyStatus defines the observed state of FirewallPolicy.
type FirewallPolicyStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        FirewallPolicyObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// FirewallPolicy is the Schema for the FirewallPolicys API
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,awsjet}
type FirewallPolicy struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              FirewallPolicySpec   `json:"spec"`
	Status            FirewallPolicyStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// FirewallPolicyList contains a list of FirewallPolicys
type FirewallPolicyList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []FirewallPolicy `json:"items"`
}

// Repository type metadata.
var (
	FirewallPolicy_Kind             = "FirewallPolicy"
	FirewallPolicy_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: FirewallPolicy_Kind}.String()
	FirewallPolicy_KindAPIVersion   = FirewallPolicy_Kind + "." + CRDGroupVersion.String()
	FirewallPolicy_GroupVersionKind = CRDGroupVersion.WithKind(FirewallPolicy_Kind)
)

func init() {
	SchemeBuilder.Register(&FirewallPolicy{}, &FirewallPolicyList{})
}
