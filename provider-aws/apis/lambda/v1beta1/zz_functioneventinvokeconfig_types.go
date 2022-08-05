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

type DestinationConfigOnFailureObservation struct {
}

type DestinationConfigOnFailureParameters struct {

	// +crossplane:generate:reference:type=github.com/upbound/official-providers/provider-aws/apis/sqs/v1beta1.Queue
	// +crossplane:generate:reference:extractor=github.com/upbound/official-providers/provider-aws/config/common.ARNExtractor()
	// +kubebuilder:validation:Optional
	Destination *string `json:"destination,omitempty" tf:"destination,omitempty"`

	// +kubebuilder:validation:Optional
	DestinationRef *v1.Reference `json:"destinationRef,omitempty" tf:"-"`

	// +kubebuilder:validation:Optional
	DestinationSelector *v1.Selector `json:"destinationSelector,omitempty" tf:"-"`
}

type FunctionEventInvokeConfigDestinationConfigObservation struct {
}

type FunctionEventInvokeConfigDestinationConfigParameters struct {

	// +kubebuilder:validation:Optional
	OnFailure []DestinationConfigOnFailureParameters `json:"onFailure,omitempty" tf:"on_failure,omitempty"`

	// +kubebuilder:validation:Optional
	OnSuccess []OnSuccessParameters `json:"onSuccess,omitempty" tf:"on_success,omitempty"`
}

type FunctionEventInvokeConfigObservation struct {
	ID *string `json:"id,omitempty" tf:"id,omitempty"`
}

type FunctionEventInvokeConfigParameters struct {

	// +kubebuilder:validation:Optional
	DestinationConfig []FunctionEventInvokeConfigDestinationConfigParameters `json:"destinationConfig,omitempty" tf:"destination_config,omitempty"`

	// +kubebuilder:validation:Required
	FunctionName *string `json:"functionName" tf:"function_name,omitempty"`

	// +kubebuilder:validation:Optional
	MaximumEventAgeInSeconds *float64 `json:"maximumEventAgeInSeconds,omitempty" tf:"maximum_event_age_in_seconds,omitempty"`

	// +kubebuilder:validation:Optional
	MaximumRetryAttempts *float64 `json:"maximumRetryAttempts,omitempty" tf:"maximum_retry_attempts,omitempty"`

	// +kubebuilder:validation:Optional
	Qualifier *string `json:"qualifier,omitempty" tf:"qualifier,omitempty"`

	// Region is the region you'd like your resource to be created in.
	// +terrajet:crd:field:TFTag=-
	// +kubebuilder:validation:Required
	Region *string `json:"region" tf:"-"`
}

type OnSuccessObservation struct {
}

type OnSuccessParameters struct {

	// +crossplane:generate:reference:type=github.com/upbound/official-providers/provider-aws/apis/sns/v1beta1.Topic
	// +crossplane:generate:reference:extractor=github.com/upbound/official-providers/provider-aws/config/common.ARNExtractor()
	// +kubebuilder:validation:Optional
	Destination *string `json:"destination,omitempty" tf:"destination,omitempty"`

	// +kubebuilder:validation:Optional
	DestinationRef *v1.Reference `json:"destinationRef,omitempty" tf:"-"`

	// +kubebuilder:validation:Optional
	DestinationSelector *v1.Selector `json:"destinationSelector,omitempty" tf:"-"`
}

// FunctionEventInvokeConfigSpec defines the desired state of FunctionEventInvokeConfig
type FunctionEventInvokeConfigSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     FunctionEventInvokeConfigParameters `json:"forProvider"`
}

// FunctionEventInvokeConfigStatus defines the observed state of FunctionEventInvokeConfig.
type FunctionEventInvokeConfigStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        FunctionEventInvokeConfigObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// FunctionEventInvokeConfig is the Schema for the FunctionEventInvokeConfigs API
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,aws}
type FunctionEventInvokeConfig struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              FunctionEventInvokeConfigSpec   `json:"spec"`
	Status            FunctionEventInvokeConfigStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// FunctionEventInvokeConfigList contains a list of FunctionEventInvokeConfigs
type FunctionEventInvokeConfigList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []FunctionEventInvokeConfig `json:"items"`
}

// Repository type metadata.
var (
	FunctionEventInvokeConfig_Kind             = "FunctionEventInvokeConfig"
	FunctionEventInvokeConfig_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: FunctionEventInvokeConfig_Kind}.String()
	FunctionEventInvokeConfig_KindAPIVersion   = FunctionEventInvokeConfig_Kind + "." + CRDGroupVersion.String()
	FunctionEventInvokeConfig_GroupVersionKind = CRDGroupVersion.WithKind(FunctionEventInvokeConfig_Kind)
)

func init() {
	SchemeBuilder.Register(&FunctionEventInvokeConfig{}, &FunctionEventInvokeConfigList{})
}
