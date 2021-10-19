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

type RouteTableObservation struct {
	Arn *string `json:"arn,omitempty" tf:"arn,omitempty"`

	OwnerID *string `json:"ownerId,omitempty" tf:"owner_id,omitempty"`

	TagsAll map[string]*string `json:"tagsAll,omitempty" tf:"tags_all,omitempty"`
}

type RouteTableParameters struct {

	// +kubebuilder:validation:Optional
	PropagatingVgws []*string `json:"propagatingVgws,omitempty" tf:"propagating_vgws,omitempty"`

	// Region is the region you'd like your resource to be created in.
	// +terrajet:crd:field:TFTag=-
	// +kubebuilder:validation:Required
	Region *string `json:"region" tf:"-"`

	// +kubebuilder:validation:Optional
	Route []RouteTableRouteParameters `json:"route,omitempty" tf:"route,omitempty"`

	// +kubebuilder:validation:Optional
	Tags map[string]*string `json:"tags,omitempty" tf:"tags,omitempty"`

	// +crossplane:generate:reference:type=VPC
	// +kubebuilder:validation:Optional
	VpcID *string `json:"vpcId,omitempty" tf:"vpc_id,omitempty"`

	// +kubebuilder:validation:Optional
	VpcIDRef *v1.Reference `json:"vpcIdRef,omitempty" tf:"-"`

	// +kubebuilder:validation:Optional
	VpcIDSelector *v1.Selector `json:"vpcIdSelector,omitempty" tf:"-"`
}

type RouteTableRouteObservation struct {
}

type RouteTableRouteParameters struct {

	// +kubebuilder:validation:Optional
	CarrierGatewayID *string `json:"carrierGatewayId,omitempty" tf:"carrier_gateway_id,omitempty"`

	// +kubebuilder:validation:Optional
	CidrBlock *string `json:"cidrBlock,omitempty" tf:"cidr_block,omitempty"`

	// +kubebuilder:validation:Optional
	DestinationPrefixListID *string `json:"destinationPrefixListId,omitempty" tf:"destination_prefix_list_id,omitempty"`

	// +kubebuilder:validation:Optional
	EgressOnlyGatewayID *string `json:"egressOnlyGatewayId,omitempty" tf:"egress_only_gateway_id,omitempty"`

	// +kubebuilder:validation:Optional
	GatewayID *string `json:"gatewayId,omitempty" tf:"gateway_id,omitempty"`

	// +kubebuilder:validation:Optional
	IPv6CidrBlock *string `json:"ipv6CidrBlock,omitempty" tf:"ipv6_cidr_block,omitempty"`

	// +crossplane:generate:reference:type=Instance
	// +kubebuilder:validation:Optional
	InstanceID *string `json:"instanceId,omitempty" tf:"instance_id,omitempty"`

	// +kubebuilder:validation:Optional
	InstanceIDRef *v1.Reference `json:"instanceIdRef,omitempty" tf:"-"`

	// +kubebuilder:validation:Optional
	InstanceIDSelector *v1.Selector `json:"instanceIdSelector,omitempty" tf:"-"`

	// +kubebuilder:validation:Optional
	LocalGatewayID *string `json:"localGatewayId,omitempty" tf:"local_gateway_id,omitempty"`

	// +kubebuilder:validation:Optional
	NatGatewayID *string `json:"natGatewayId,omitempty" tf:"nat_gateway_id,omitempty"`

	// +crossplane:generate:reference:type=EC2NetworkInterface
	// +kubebuilder:validation:Optional
	NetworkInterfaceID *string `json:"networkInterfaceId,omitempty" tf:"network_interface_id,omitempty"`

	// +kubebuilder:validation:Optional
	NetworkInterfaceIDRef *v1.Reference `json:"networkInterfaceIdRef,omitempty" tf:"-"`

	// +kubebuilder:validation:Optional
	NetworkInterfaceIDSelector *v1.Selector `json:"networkInterfaceIdSelector,omitempty" tf:"-"`

	// +kubebuilder:validation:Optional
	TransitGatewayID *string `json:"transitGatewayId,omitempty" tf:"transit_gateway_id,omitempty"`

	// +crossplane:generate:reference:type=VpcEndpoint
	// +kubebuilder:validation:Optional
	VpcEndpointID *string `json:"vpcEndpointId,omitempty" tf:"vpc_endpoint_id,omitempty"`

	// +kubebuilder:validation:Optional
	VpcEndpointIDRef *v1.Reference `json:"vpcEndpointIdRef,omitempty" tf:"-"`

	// +kubebuilder:validation:Optional
	VpcEndpointIDSelector *v1.Selector `json:"vpcEndpointIdSelector,omitempty" tf:"-"`

	// +crossplane:generate:reference:type=VpcPeeringConnection
	// +kubebuilder:validation:Optional
	VpcPeeringConnectionID *string `json:"vpcPeeringConnectionId,omitempty" tf:"vpc_peering_connection_id,omitempty"`

	// +kubebuilder:validation:Optional
	VpcPeeringConnectionIDRef *v1.Reference `json:"vpcPeeringConnectionIdRef,omitempty" tf:"-"`

	// +kubebuilder:validation:Optional
	VpcPeeringConnectionIDSelector *v1.Selector `json:"vpcPeeringConnectionIdSelector,omitempty" tf:"-"`
}

// RouteTableSpec defines the desired state of RouteTable
type RouteTableSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     RouteTableParameters `json:"forProvider"`
}

// RouteTableStatus defines the observed state of RouteTable.
type RouteTableStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        RouteTableObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// RouteTable is the Schema for the RouteTables API
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,tfaws}
type RouteTable struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              RouteTableSpec   `json:"spec"`
	Status            RouteTableStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// RouteTableList contains a list of RouteTables
type RouteTableList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []RouteTable `json:"items"`
}

// Repository type metadata.
var (
	RouteTableKind             = "RouteTable"
	RouteTableGroupKind        = schema.GroupKind{Group: Group, Kind: RouteTableKind}.String()
	RouteTableKindAPIVersion   = RouteTableKind + "." + GroupVersion.String()
	RouteTableGroupVersionKind = GroupVersion.WithKind(RouteTableKind)
)

func init() {
	SchemeBuilder.Register(&RouteTable{}, &RouteTableList{})
}
