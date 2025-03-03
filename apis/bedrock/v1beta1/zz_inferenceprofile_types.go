// SPDX-FileCopyrightText: 2024 The Crossplane Authors <https://crossplane.io>
//
// SPDX-License-Identifier: Apache-2.0

// Code generated by upjet. DO NOT EDIT.

package v1beta1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime/schema"

	v1 "github.com/crossplane/crossplane-runtime/apis/common/v1"
)

type InferenceProfileInitParameters struct {

	// The description of the inference profile.
	Description *string `json:"description,omitempty" tf:"description,omitempty"`

	// The source of the model this inference profile will track metrics and cost for. See model_source.
	ModelSource []ModelSourceInitParameters `json:"modelSource,omitempty" tf:"model_source,omitempty"`

	// The name of the inference profile.
	Name *string `json:"name,omitempty" tf:"name,omitempty"`

	// Key-value map of resource tags.
	// +mapType=granular
	Tags map[string]*string `json:"tags,omitempty" tf:"tags,omitempty"`
}

type InferenceProfileObservation struct {

	// The Amazon Resource Name (ARN) of the inference profile.
	Arn *string `json:"arn,omitempty" tf:"arn,omitempty"`

	// The time at which the inference profile was created.
	CreatedAt *string `json:"createdAt,omitempty" tf:"created_at,omitempty"`

	// The description of the inference profile.
	Description *string `json:"description,omitempty" tf:"description,omitempty"`

	// The unique identifier of the inference profile.
	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	// The source of the model this inference profile will track metrics and cost for. See model_source.
	ModelSource []ModelSourceObservation `json:"modelSource,omitempty" tf:"model_source,omitempty"`

	// A list of information about each model in the inference profile. See models.
	Models []ModelsObservation `json:"models,omitempty" tf:"models,omitempty"`

	// The name of the inference profile.
	Name *string `json:"name,omitempty" tf:"name,omitempty"`

	// The status of the inference profile. ACTIVE means that the inference profile is available to use.
	Status *string `json:"status,omitempty" tf:"status,omitempty"`

	// Key-value map of resource tags.
	// +mapType=granular
	Tags map[string]*string `json:"tags,omitempty" tf:"tags,omitempty"`

	// +mapType=granular
	TagsAll map[string]*string `json:"tagsAll,omitempty" tf:"tags_all,omitempty"`

	// The type of the inference profile. SYSTEM_DEFINED means that the inference profile is defined by Amazon Bedrock. APPLICATION means that the inference profile is defined by the user.
	Type *string `json:"type,omitempty" tf:"type,omitempty"`

	// The time at which the inference profile was last updated.
	UpdatedAt *string `json:"updatedAt,omitempty" tf:"updated_at,omitempty"`
}

type InferenceProfileParameters struct {

	// The description of the inference profile.
	// +kubebuilder:validation:Optional
	Description *string `json:"description,omitempty" tf:"description,omitempty"`

	// The source of the model this inference profile will track metrics and cost for. See model_source.
	// +kubebuilder:validation:Optional
	ModelSource []ModelSourceParameters `json:"modelSource,omitempty" tf:"model_source,omitempty"`

	// The name of the inference profile.
	// +kubebuilder:validation:Optional
	Name *string `json:"name,omitempty" tf:"name,omitempty"`

	// Region is the region you'd like your resource to be created in.
	// +upjet:crd:field:TFTag=-
	// +kubebuilder:validation:Required
	Region *string `json:"region" tf:"-"`

	// Key-value map of resource tags.
	// +kubebuilder:validation:Optional
	// +mapType=granular
	Tags map[string]*string `json:"tags,omitempty" tf:"tags,omitempty"`
}

type ModelSourceInitParameters struct {

	// The Amazon Resource Name (ARN) of the model.
	CopyFrom *string `json:"copyFrom,omitempty" tf:"copy_from,omitempty"`
}

type ModelSourceObservation struct {

	// The Amazon Resource Name (ARN) of the model.
	CopyFrom *string `json:"copyFrom,omitempty" tf:"copy_from,omitempty"`
}

type ModelSourceParameters struct {

	// The Amazon Resource Name (ARN) of the model.
	// +kubebuilder:validation:Optional
	CopyFrom *string `json:"copyFrom" tf:"copy_from,omitempty"`
}

type ModelsInitParameters struct {
}

type ModelsObservation struct {

	// The Amazon Resource Name (ARN) of the model.
	ModelArn *string `json:"modelArn,omitempty" tf:"model_arn,omitempty"`
}

type ModelsParameters struct {
}

// InferenceProfileSpec defines the desired state of InferenceProfile
type InferenceProfileSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     InferenceProfileParameters `json:"forProvider"`
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
	InitProvider InferenceProfileInitParameters `json:"initProvider,omitempty"`
}

// InferenceProfileStatus defines the observed state of InferenceProfile.
type InferenceProfileStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        InferenceProfileObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true
// +kubebuilder:subresource:status
// +kubebuilder:storageversion

// InferenceProfile is the Schema for the InferenceProfiles API.
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,aws}
type InferenceProfile struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.name) || (has(self.initProvider) && has(self.initProvider.name))",message="spec.forProvider.name is a required parameter"
	Spec   InferenceProfileSpec   `json:"spec"`
	Status InferenceProfileStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// InferenceProfileList contains a list of InferenceProfiles
type InferenceProfileList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []InferenceProfile `json:"items"`
}

// Repository type metadata.
var (
	InferenceProfile_Kind             = "InferenceProfile"
	InferenceProfile_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: InferenceProfile_Kind}.String()
	InferenceProfile_KindAPIVersion   = InferenceProfile_Kind + "." + CRDGroupVersion.String()
	InferenceProfile_GroupVersionKind = CRDGroupVersion.WithKind(InferenceProfile_Kind)
)

func init() {
	SchemeBuilder.Register(&InferenceProfile{}, &InferenceProfileList{})
}
