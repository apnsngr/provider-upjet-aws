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

type BackendServerPolicyInitParameters struct {

	// The instance port to apply the policy to.
	InstancePort *float64 `json:"instancePort,omitempty" tf:"instance_port,omitempty"`

	// The load balancer to attach the policy to.
	// +crossplane:generate:reference:type=github.com/upbound/provider-aws/apis/elb/v1beta1.ELB
	LoadBalancerName *string `json:"loadBalancerName,omitempty" tf:"load_balancer_name,omitempty"`

	// Reference to a ELB in elb to populate loadBalancerName.
	// +kubebuilder:validation:Optional
	LoadBalancerNameRef *v1.Reference `json:"loadBalancerNameRef,omitempty" tf:"-"`

	// Selector for a ELB in elb to populate loadBalancerName.
	// +kubebuilder:validation:Optional
	LoadBalancerNameSelector *v1.Selector `json:"loadBalancerNameSelector,omitempty" tf:"-"`

	// List of Policy Names to apply to the backend server.
	// +listType=set
	PolicyNames []*string `json:"policyNames,omitempty" tf:"policy_names,omitempty"`
}

type BackendServerPolicyObservation struct {

	// The ID of the policy.
	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	// The instance port to apply the policy to.
	InstancePort *float64 `json:"instancePort,omitempty" tf:"instance_port,omitempty"`

	// The load balancer to attach the policy to.
	LoadBalancerName *string `json:"loadBalancerName,omitempty" tf:"load_balancer_name,omitempty"`

	// List of Policy Names to apply to the backend server.
	// +listType=set
	PolicyNames []*string `json:"policyNames,omitempty" tf:"policy_names,omitempty"`
}

type BackendServerPolicyParameters struct {

	// The instance port to apply the policy to.
	// +kubebuilder:validation:Optional
	InstancePort *float64 `json:"instancePort,omitempty" tf:"instance_port,omitempty"`

	// The load balancer to attach the policy to.
	// +crossplane:generate:reference:type=github.com/upbound/provider-aws/apis/elb/v1beta1.ELB
	// +kubebuilder:validation:Optional
	LoadBalancerName *string `json:"loadBalancerName,omitempty" tf:"load_balancer_name,omitempty"`

	// Reference to a ELB in elb to populate loadBalancerName.
	// +kubebuilder:validation:Optional
	LoadBalancerNameRef *v1.Reference `json:"loadBalancerNameRef,omitempty" tf:"-"`

	// Selector for a ELB in elb to populate loadBalancerName.
	// +kubebuilder:validation:Optional
	LoadBalancerNameSelector *v1.Selector `json:"loadBalancerNameSelector,omitempty" tf:"-"`

	// List of Policy Names to apply to the backend server.
	// +kubebuilder:validation:Optional
	// +listType=set
	PolicyNames []*string `json:"policyNames,omitempty" tf:"policy_names,omitempty"`

	// Region is the region you'd like your resource to be created in.
	// +upjet:crd:field:TFTag=-
	// +kubebuilder:validation:Required
	Region *string `json:"region" tf:"-"`
}

// BackendServerPolicySpec defines the desired state of BackendServerPolicy
type BackendServerPolicySpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     BackendServerPolicyParameters `json:"forProvider"`
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
	InitProvider BackendServerPolicyInitParameters `json:"initProvider,omitempty"`
}

// BackendServerPolicyStatus defines the observed state of BackendServerPolicy.
type BackendServerPolicyStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        BackendServerPolicyObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true
// +kubebuilder:subresource:status
// +kubebuilder:storageversion

// BackendServerPolicy is the Schema for the BackendServerPolicys API. Attaches a load balancer policy to an ELB backend server.
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,aws}
type BackendServerPolicy struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.instancePort) || (has(self.initProvider) && has(self.initProvider.instancePort))",message="spec.forProvider.instancePort is a required parameter"
	Spec   BackendServerPolicySpec   `json:"spec"`
	Status BackendServerPolicyStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// BackendServerPolicyList contains a list of BackendServerPolicys
type BackendServerPolicyList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []BackendServerPolicy `json:"items"`
}

// Repository type metadata.
var (
	BackendServerPolicy_Kind             = "BackendServerPolicy"
	BackendServerPolicy_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: BackendServerPolicy_Kind}.String()
	BackendServerPolicy_KindAPIVersion   = BackendServerPolicy_Kind + "." + CRDGroupVersion.String()
	BackendServerPolicy_GroupVersionKind = CRDGroupVersion.WithKind(BackendServerPolicy_Kind)
)

func init() {
	SchemeBuilder.Register(&BackendServerPolicy{}, &BackendServerPolicyList{})
}
