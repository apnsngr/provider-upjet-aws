// SPDX-FileCopyrightText: 2023 The Crossplane Authors <https://crossplane.io>
//
// SPDX-License-Identifier: Apache-2.0

/*
Copyright 2022 Upbound Inc.
*/

// Code generated by upjet. DO NOT EDIT.

package v1beta1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime/schema"

	v1 "github.com/crossplane/crossplane-runtime/apis/common/v1"
)

type AccepterInitParameters struct {
}

type AccepterObservation struct {

	// Allow a local VPC to resolve public DNS hostnames to
	// private IP addresses when queried from instances in the peer VPC.
	AllowRemoteVPCDNSResolution *bool `json:"allowRemoteVpcDnsResolution,omitempty" tf:"allow_remote_vpc_dns_resolution,omitempty"`
}

type AccepterParameters struct {
}

type RequesterInitParameters struct {
}

type RequesterObservation struct {

	// Allow a local VPC to resolve public DNS hostnames to
	// private IP addresses when queried from instances in the peer VPC.
	AllowRemoteVPCDNSResolution *bool `json:"allowRemoteVpcDnsResolution,omitempty" tf:"allow_remote_vpc_dns_resolution,omitempty"`
}

type RequesterParameters struct {
}

type VPCPeeringConnectionInitParameters_2 struct {

	// Accept the peering (both VPCs need to be in the same AWS account and region).
	AutoAccept *bool `json:"autoAccept,omitempty" tf:"auto_accept,omitempty"`

	// The AWS account ID of the owner of the peer VPC.
	// Defaults to the account ID the AWS provider is currently connected to.
	PeerOwnerID *string `json:"peerOwnerId,omitempty" tf:"peer_owner_id,omitempty"`

	// The region of the accepter VPC of the VPC Peering Connection. auto_accept must be false,
	// and use the aws_vpc_peering_connection_accepter to manage the accepter side.
	PeerRegion *string `json:"peerRegion,omitempty" tf:"peer_region,omitempty"`

	// The ID of the VPC with which you are creating the VPC Peering Connection.
	// +crossplane:generate:reference:type=VPC
	PeerVPCID *string `json:"peerVpcId,omitempty" tf:"peer_vpc_id,omitempty"`

	// Reference to a VPC to populate peerVpcId.
	// +kubebuilder:validation:Optional
	PeerVPCIDRef *v1.Reference `json:"peerVpcIdRef,omitempty" tf:"-"`

	// Selector for a VPC to populate peerVpcId.
	// +kubebuilder:validation:Optional
	PeerVPCIDSelector *v1.Selector `json:"peerVpcIdSelector,omitempty" tf:"-"`

	// Key-value map of resource tags.
	// +mapType=granular
	Tags map[string]*string `json:"tags,omitempty" tf:"tags,omitempty"`

	// The ID of the requester VPC.
	// +crossplane:generate:reference:type=github.com/upbound/provider-aws/apis/ec2/v1beta1.VPC
	VPCID *string `json:"vpcId,omitempty" tf:"vpc_id,omitempty"`

	// Reference to a VPC in ec2 to populate vpcId.
	// +kubebuilder:validation:Optional
	VPCIDRef *v1.Reference `json:"vpcIdRef,omitempty" tf:"-"`

	// Selector for a VPC in ec2 to populate vpcId.
	// +kubebuilder:validation:Optional
	VPCIDSelector *v1.Selector `json:"vpcIdSelector,omitempty" tf:"-"`
}

type VPCPeeringConnectionObservation_2 struct {

	// The status of the VPC Peering Connection request.
	AcceptStatus *string `json:"acceptStatus,omitempty" tf:"accept_status,omitempty"`

	// An optional configuration block that allows for VPC Peering Connection options to be set for the VPC that accepts
	// the peering connection (a maximum of one).
	Accepter []AccepterObservation `json:"accepter,omitempty" tf:"accepter,omitempty"`

	// Accept the peering (both VPCs need to be in the same AWS account and region).
	AutoAccept *bool `json:"autoAccept,omitempty" tf:"auto_accept,omitempty"`

	// The ID of the VPC Peering Connection.
	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	// The AWS account ID of the owner of the peer VPC.
	// Defaults to the account ID the AWS provider is currently connected to.
	PeerOwnerID *string `json:"peerOwnerId,omitempty" tf:"peer_owner_id,omitempty"`

	// The region of the accepter VPC of the VPC Peering Connection. auto_accept must be false,
	// and use the aws_vpc_peering_connection_accepter to manage the accepter side.
	PeerRegion *string `json:"peerRegion,omitempty" tf:"peer_region,omitempty"`

	// The ID of the VPC with which you are creating the VPC Peering Connection.
	PeerVPCID *string `json:"peerVpcId,omitempty" tf:"peer_vpc_id,omitempty"`

	// A optional configuration block that allows for VPC Peering Connection options to be set for the VPC that requests
	// the peering connection (a maximum of one).
	Requester []RequesterObservation `json:"requester,omitempty" tf:"requester,omitempty"`

	// Key-value map of resource tags.
	// +mapType=granular
	Tags map[string]*string `json:"tags,omitempty" tf:"tags,omitempty"`

	// A map of tags assigned to the resource, including those inherited from the provider default_tags configuration block.
	// +mapType=granular
	TagsAll map[string]*string `json:"tagsAll,omitempty" tf:"tags_all,omitempty"`

	// The ID of the requester VPC.
	VPCID *string `json:"vpcId,omitempty" tf:"vpc_id,omitempty"`
}

type VPCPeeringConnectionParameters_2 struct {

	// Accept the peering (both VPCs need to be in the same AWS account and region).
	// +kubebuilder:validation:Optional
	AutoAccept *bool `json:"autoAccept,omitempty" tf:"auto_accept,omitempty"`

	// The AWS account ID of the owner of the peer VPC.
	// Defaults to the account ID the AWS provider is currently connected to.
	// +kubebuilder:validation:Optional
	PeerOwnerID *string `json:"peerOwnerId,omitempty" tf:"peer_owner_id,omitempty"`

	// The region of the accepter VPC of the VPC Peering Connection. auto_accept must be false,
	// and use the aws_vpc_peering_connection_accepter to manage the accepter side.
	// +kubebuilder:validation:Optional
	PeerRegion *string `json:"peerRegion,omitempty" tf:"peer_region,omitempty"`

	// The ID of the VPC with which you are creating the VPC Peering Connection.
	// +crossplane:generate:reference:type=VPC
	// +kubebuilder:validation:Optional
	PeerVPCID *string `json:"peerVpcId,omitempty" tf:"peer_vpc_id,omitempty"`

	// Reference to a VPC to populate peerVpcId.
	// +kubebuilder:validation:Optional
	PeerVPCIDRef *v1.Reference `json:"peerVpcIdRef,omitempty" tf:"-"`

	// Selector for a VPC to populate peerVpcId.
	// +kubebuilder:validation:Optional
	PeerVPCIDSelector *v1.Selector `json:"peerVpcIdSelector,omitempty" tf:"-"`

	// Region is the region you'd like your resource to be created in.
	// +upjet:crd:field:TFTag=-
	// +kubebuilder:validation:Required
	Region *string `json:"region" tf:"-"`

	// Key-value map of resource tags.
	// +kubebuilder:validation:Optional
	// +mapType=granular
	Tags map[string]*string `json:"tags,omitempty" tf:"tags,omitempty"`

	// The ID of the requester VPC.
	// +crossplane:generate:reference:type=github.com/upbound/provider-aws/apis/ec2/v1beta1.VPC
	// +kubebuilder:validation:Optional
	VPCID *string `json:"vpcId,omitempty" tf:"vpc_id,omitempty"`

	// Reference to a VPC in ec2 to populate vpcId.
	// +kubebuilder:validation:Optional
	VPCIDRef *v1.Reference `json:"vpcIdRef,omitempty" tf:"-"`

	// Selector for a VPC in ec2 to populate vpcId.
	// +kubebuilder:validation:Optional
	VPCIDSelector *v1.Selector `json:"vpcIdSelector,omitempty" tf:"-"`
}

// VPCPeeringConnectionSpec defines the desired state of VPCPeeringConnection
type VPCPeeringConnectionSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     VPCPeeringConnectionParameters_2 `json:"forProvider"`
	// THIS IS A BETA FIELD. It will be honored
	// unless the Management Policies feature flag is disabled.
	// InitProvider holds the same fields as ForProvider, with the exception
	// of Identifier and other resource reference fields. The fields that are
	// in InitProvider are merged into ForProvider when the resource is created.
	// The same fields are also added to the terraform ignore_changes hook, to
	// avoid updating them after creation. This is useful for fields that are
	// required on creation, but we do not desire to update them after creation,
	// for example because of an external controller is managing them, like an
	// autoscaler.
	InitProvider VPCPeeringConnectionInitParameters_2 `json:"initProvider,omitempty"`
}

// VPCPeeringConnectionStatus defines the observed state of VPCPeeringConnection.
type VPCPeeringConnectionStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        VPCPeeringConnectionObservation_2 `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true
// +kubebuilder:subresource:status
// +kubebuilder:storageversion

// VPCPeeringConnection is the Schema for the VPCPeeringConnections API. Provides a resource to manage a VPC peering connection.
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,aws}
type VPCPeeringConnection struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              VPCPeeringConnectionSpec   `json:"spec"`
	Status            VPCPeeringConnectionStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// VPCPeeringConnectionList contains a list of VPCPeeringConnections
type VPCPeeringConnectionList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []VPCPeeringConnection `json:"items"`
}

// Repository type metadata.
var (
	VPCPeeringConnection_Kind             = "VPCPeeringConnection"
	VPCPeeringConnection_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: VPCPeeringConnection_Kind}.String()
	VPCPeeringConnection_KindAPIVersion   = VPCPeeringConnection_Kind + "." + CRDGroupVersion.String()
	VPCPeeringConnection_GroupVersionKind = CRDGroupVersion.WithKind(VPCPeeringConnection_Kind)
)

func init() {
	SchemeBuilder.Register(&VPCPeeringConnection{}, &VPCPeeringConnectionList{})
}
