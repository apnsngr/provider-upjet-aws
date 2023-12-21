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

type ConfigurationTemplateInitParameters struct {

	// –  name of the application to associate with this configuration template
	// +crossplane:generate:reference:type=github.com/upbound/provider-aws/apis/elasticbeanstalk/v1beta1.Application
	Application *string `json:"application,omitempty" tf:"application,omitempty"`

	// Reference to a Application in elasticbeanstalk to populate application.
	// +kubebuilder:validation:Optional
	ApplicationRef *v1.Reference `json:"applicationRef,omitempty" tf:"-"`

	// Selector for a Application in elasticbeanstalk to populate application.
	// +kubebuilder:validation:Optional
	ApplicationSelector *v1.Selector `json:"applicationSelector,omitempty" tf:"-"`

	// Short description of the Template
	Description *string `json:"description,omitempty" tf:"description,omitempty"`

	// –  The ID of the environment used with this configuration template
	EnvironmentID *string `json:"environmentId,omitempty" tf:"environment_id,omitempty"`

	// –  Option settings to configure the new Environment. These
	// override specific values that are set as defaults. The format is detailed
	// below in Option Settings
	Setting []SettingInitParameters `json:"setting,omitempty" tf:"setting,omitempty"`

	// –  A solution stack to base your Template
	// off of. Example stacks can be found in the Amazon API documentation
	SolutionStackName *string `json:"solutionStackName,omitempty" tf:"solution_stack_name,omitempty"`
}

type ConfigurationTemplateObservation struct {

	// –  name of the application to associate with this configuration template
	Application *string `json:"application,omitempty" tf:"application,omitempty"`

	// Short description of the Template
	Description *string `json:"description,omitempty" tf:"description,omitempty"`

	// –  The ID of the environment used with this configuration template
	EnvironmentID *string `json:"environmentId,omitempty" tf:"environment_id,omitempty"`

	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	// –  Option settings to configure the new Environment. These
	// override specific values that are set as defaults. The format is detailed
	// below in Option Settings
	Setting []SettingObservation `json:"setting,omitempty" tf:"setting,omitempty"`

	// –  A solution stack to base your Template
	// off of. Example stacks can be found in the Amazon API documentation
	SolutionStackName *string `json:"solutionStackName,omitempty" tf:"solution_stack_name,omitempty"`
}

type ConfigurationTemplateParameters struct {

	// –  name of the application to associate with this configuration template
	// +crossplane:generate:reference:type=github.com/upbound/provider-aws/apis/elasticbeanstalk/v1beta1.Application
	// +kubebuilder:validation:Optional
	Application *string `json:"application,omitempty" tf:"application,omitempty"`

	// Reference to a Application in elasticbeanstalk to populate application.
	// +kubebuilder:validation:Optional
	ApplicationRef *v1.Reference `json:"applicationRef,omitempty" tf:"-"`

	// Selector for a Application in elasticbeanstalk to populate application.
	// +kubebuilder:validation:Optional
	ApplicationSelector *v1.Selector `json:"applicationSelector,omitempty" tf:"-"`

	// Short description of the Template
	// +kubebuilder:validation:Optional
	Description *string `json:"description,omitempty" tf:"description,omitempty"`

	// –  The ID of the environment used with this configuration template
	// +kubebuilder:validation:Optional
	EnvironmentID *string `json:"environmentId,omitempty" tf:"environment_id,omitempty"`

	// Region is the region you'd like your resource to be created in.
	// +upjet:crd:field:TFTag=-
	// +kubebuilder:validation:Required
	Region *string `json:"region" tf:"-"`

	// –  Option settings to configure the new Environment. These
	// override specific values that are set as defaults. The format is detailed
	// below in Option Settings
	// +kubebuilder:validation:Optional
	Setting []SettingParameters `json:"setting,omitempty" tf:"setting,omitempty"`

	// –  A solution stack to base your Template
	// off of. Example stacks can be found in the Amazon API documentation
	// +kubebuilder:validation:Optional
	SolutionStackName *string `json:"solutionStackName,omitempty" tf:"solution_stack_name,omitempty"`
}

type SettingInitParameters struct {

	// A unique name for this Template.
	Name *string `json:"name,omitempty" tf:"name,omitempty"`

	// unique namespace identifying the option's associated AWS resource
	Namespace *string `json:"namespace,omitempty" tf:"namespace,omitempty"`

	// resource name for scheduled action
	Resource *string `json:"resource,omitempty" tf:"resource,omitempty"`

	// value for the configuration option
	Value *string `json:"value,omitempty" tf:"value,omitempty"`
}

type SettingObservation struct {

	// A unique name for this Template.
	Name *string `json:"name,omitempty" tf:"name,omitempty"`

	// unique namespace identifying the option's associated AWS resource
	Namespace *string `json:"namespace,omitempty" tf:"namespace,omitempty"`

	// resource name for scheduled action
	Resource *string `json:"resource,omitempty" tf:"resource,omitempty"`

	// value for the configuration option
	Value *string `json:"value,omitempty" tf:"value,omitempty"`
}

type SettingParameters struct {

	// A unique name for this Template.
	// +kubebuilder:validation:Optional
	Name *string `json:"name" tf:"name,omitempty"`

	// unique namespace identifying the option's associated AWS resource
	// +kubebuilder:validation:Optional
	Namespace *string `json:"namespace" tf:"namespace,omitempty"`

	// resource name for scheduled action
	// +kubebuilder:validation:Optional
	Resource *string `json:"resource,omitempty" tf:"resource,omitempty"`

	// value for the configuration option
	// +kubebuilder:validation:Optional
	Value *string `json:"value" tf:"value,omitempty"`
}

// ConfigurationTemplateSpec defines the desired state of ConfigurationTemplate
type ConfigurationTemplateSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     ConfigurationTemplateParameters `json:"forProvider"`
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
	InitProvider ConfigurationTemplateInitParameters `json:"initProvider,omitempty"`
}

// ConfigurationTemplateStatus defines the observed state of ConfigurationTemplate.
type ConfigurationTemplateStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        ConfigurationTemplateObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true
// +kubebuilder:subresource:status
// +kubebuilder:storageversion

// ConfigurationTemplate is the Schema for the ConfigurationTemplates API. Provides an Elastic Beanstalk Configuration Template
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,aws}
type ConfigurationTemplate struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              ConfigurationTemplateSpec   `json:"spec"`
	Status            ConfigurationTemplateStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// ConfigurationTemplateList contains a list of ConfigurationTemplates
type ConfigurationTemplateList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []ConfigurationTemplate `json:"items"`
}

// Repository type metadata.
var (
	ConfigurationTemplate_Kind             = "ConfigurationTemplate"
	ConfigurationTemplate_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: ConfigurationTemplate_Kind}.String()
	ConfigurationTemplate_KindAPIVersion   = ConfigurationTemplate_Kind + "." + CRDGroupVersion.String()
	ConfigurationTemplate_GroupVersionKind = CRDGroupVersion.WithKind(ConfigurationTemplate_Kind)
)

func init() {
	SchemeBuilder.Register(&ConfigurationTemplate{}, &ConfigurationTemplateList{})
}
