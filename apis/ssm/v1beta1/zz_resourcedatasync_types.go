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

type ResourceDataSyncInitParameters struct {

	// Amazon S3 configuration details for the sync.
	S3Destination []S3DestinationInitParameters `json:"s3Destination,omitempty" tf:"s3_destination,omitempty"`
}

type ResourceDataSyncObservation struct {
	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	// Amazon S3 configuration details for the sync.
	S3Destination []S3DestinationObservation `json:"s3Destination,omitempty" tf:"s3_destination,omitempty"`
}

type ResourceDataSyncParameters struct {

	// Region with the bucket targeted by the Resource Data Sync.
	// Region is the region you'd like your resource to be created in.
	// +upjet:crd:field:TFTag=-
	// +kubebuilder:validation:Required
	Region *string `json:"region" tf:"-"`

	// Amazon S3 configuration details for the sync.
	// +kubebuilder:validation:Optional
	S3Destination []S3DestinationParameters `json:"s3Destination,omitempty" tf:"s3_destination,omitempty"`
}

type S3DestinationInitParameters struct {

	// Name of S3 bucket where the aggregated data is stored.
	// +crossplane:generate:reference:type=github.com/upbound/provider-aws/apis/s3/v1beta1.Bucket
	BucketName *string `json:"bucketName,omitempty" tf:"bucket_name,omitempty"`

	// Reference to a Bucket in s3 to populate bucketName.
	// +kubebuilder:validation:Optional
	BucketNameRef *v1.Reference `json:"bucketNameRef,omitempty" tf:"-"`

	// Selector for a Bucket in s3 to populate bucketName.
	// +kubebuilder:validation:Optional
	BucketNameSelector *v1.Selector `json:"bucketNameSelector,omitempty" tf:"-"`

	// ARN of an encryption key for a destination in Amazon S3.
	KMSKeyArn *string `json:"kmsKeyArn,omitempty" tf:"kms_key_arn,omitempty"`

	// Prefix for the bucket.
	Prefix *string `json:"prefix,omitempty" tf:"prefix,omitempty"`

	// A supported sync format. Only JsonSerDe is currently supported. Defaults to JsonSerDe.
	SyncFormat *string `json:"syncFormat,omitempty" tf:"sync_format,omitempty"`
}

type S3DestinationObservation struct {

	// Name of S3 bucket where the aggregated data is stored.
	BucketName *string `json:"bucketName,omitempty" tf:"bucket_name,omitempty"`

	// ARN of an encryption key for a destination in Amazon S3.
	KMSKeyArn *string `json:"kmsKeyArn,omitempty" tf:"kms_key_arn,omitempty"`

	// Prefix for the bucket.
	Prefix *string `json:"prefix,omitempty" tf:"prefix,omitempty"`

	// Region with the bucket targeted by the Resource Data Sync.
	Region *string `json:"region,omitempty" tf:"region,omitempty"`

	// A supported sync format. Only JsonSerDe is currently supported. Defaults to JsonSerDe.
	SyncFormat *string `json:"syncFormat,omitempty" tf:"sync_format,omitempty"`
}

type S3DestinationParameters struct {

	// Name of S3 bucket where the aggregated data is stored.
	// +crossplane:generate:reference:type=github.com/upbound/provider-aws/apis/s3/v1beta1.Bucket
	// +kubebuilder:validation:Optional
	BucketName *string `json:"bucketName,omitempty" tf:"bucket_name,omitempty"`

	// Reference to a Bucket in s3 to populate bucketName.
	// +kubebuilder:validation:Optional
	BucketNameRef *v1.Reference `json:"bucketNameRef,omitempty" tf:"-"`

	// Selector for a Bucket in s3 to populate bucketName.
	// +kubebuilder:validation:Optional
	BucketNameSelector *v1.Selector `json:"bucketNameSelector,omitempty" tf:"-"`

	// ARN of an encryption key for a destination in Amazon S3.
	// +kubebuilder:validation:Optional
	KMSKeyArn *string `json:"kmsKeyArn,omitempty" tf:"kms_key_arn,omitempty"`

	// Prefix for the bucket.
	// +kubebuilder:validation:Optional
	Prefix *string `json:"prefix,omitempty" tf:"prefix,omitempty"`

	// Region with the bucket targeted by the Resource Data Sync.
	// +crossplane:generate:reference:type=github.com/upbound/provider-aws/apis/s3/v1beta1.Bucket
	// +crossplane:generate:reference:extractor=github.com/crossplane/upjet/pkg/resource.ExtractParamPath("region",false)
	// +kubebuilder:validation:Optional
	Region *string `json:"region,omitempty" tf:"region,omitempty"`

	// Reference to a Bucket in s3 to populate region.
	// +kubebuilder:validation:Optional
	RegionRef *v1.Reference `json:"regionRef,omitempty" tf:"-"`

	// Selector for a Bucket in s3 to populate region.
	// +kubebuilder:validation:Optional
	RegionSelector *v1.Selector `json:"regionSelector,omitempty" tf:"-"`

	// A supported sync format. Only JsonSerDe is currently supported. Defaults to JsonSerDe.
	// +kubebuilder:validation:Optional
	SyncFormat *string `json:"syncFormat,omitempty" tf:"sync_format,omitempty"`
}

// ResourceDataSyncSpec defines the desired state of ResourceDataSync
type ResourceDataSyncSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     ResourceDataSyncParameters `json:"forProvider"`
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
	InitProvider ResourceDataSyncInitParameters `json:"initProvider,omitempty"`
}

// ResourceDataSyncStatus defines the observed state of ResourceDataSync.
type ResourceDataSyncStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        ResourceDataSyncObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true
// +kubebuilder:subresource:status
// +kubebuilder:storageversion

// ResourceDataSync is the Schema for the ResourceDataSyncs API. Provides a SSM resource data sync.
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,aws}
type ResourceDataSync struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.s3Destination) || (has(self.initProvider) && has(self.initProvider.s3Destination))",message="spec.forProvider.s3Destination is a required parameter"
	Spec   ResourceDataSyncSpec   `json:"spec"`
	Status ResourceDataSyncStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// ResourceDataSyncList contains a list of ResourceDataSyncs
type ResourceDataSyncList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []ResourceDataSync `json:"items"`
}

// Repository type metadata.
var (
	ResourceDataSync_Kind             = "ResourceDataSync"
	ResourceDataSync_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: ResourceDataSync_Kind}.String()
	ResourceDataSync_KindAPIVersion   = ResourceDataSync_Kind + "." + CRDGroupVersion.String()
	ResourceDataSync_GroupVersionKind = CRDGroupVersion.WithKind(ResourceDataSync_Kind)
)

func init() {
	SchemeBuilder.Register(&ResourceDataSync{}, &ResourceDataSyncList{})
}
