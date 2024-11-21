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

type AliasInitParameters struct {

	// Set to true if you want Route 53 to determine whether to respond to DNS queries using this resource record set by checking the health of the resource record set. Some resources have special requirements, see related part of documentation.
	EvaluateTargetHealth *bool `json:"evaluateTargetHealth,omitempty" tf:"evaluate_target_health,omitempty"`

	// The name of the record.
	Name *string `json:"name,omitempty" tf:"name,omitempty"`

	// The ID of the hosted zone to contain this record.
	ZoneID *string `json:"zoneId,omitempty" tf:"zone_id,omitempty"`
}

type AliasObservation struct {

	// Set to true if you want Route 53 to determine whether to respond to DNS queries using this resource record set by checking the health of the resource record set. Some resources have special requirements, see related part of documentation.
	EvaluateTargetHealth *bool `json:"evaluateTargetHealth,omitempty" tf:"evaluate_target_health,omitempty"`

	// The name of the record.
	Name *string `json:"name,omitempty" tf:"name,omitempty"`

	// The ID of the hosted zone to contain this record.
	ZoneID *string `json:"zoneId,omitempty" tf:"zone_id,omitempty"`
}

type AliasParameters struct {

	// Set to true if you want Route 53 to determine whether to respond to DNS queries using this resource record set by checking the health of the resource record set. Some resources have special requirements, see related part of documentation.
	// +kubebuilder:validation:Optional
	EvaluateTargetHealth *bool `json:"evaluateTargetHealth" tf:"evaluate_target_health,omitempty"`

	// The name of the record.
	// +kubebuilder:validation:Optional
	Name *string `json:"name" tf:"name,omitempty"`

	// The ID of the hosted zone to contain this record.
	// +kubebuilder:validation:Optional
	ZoneID *string `json:"zoneId" tf:"zone_id,omitempty"`
}

type CidrRoutingPolicyInitParameters struct {

	// The CIDR collection ID. See the aws_route53_cidr_collection resource for more details.
	CollectionID *string `json:"collectionId,omitempty" tf:"collection_id,omitempty"`

	// The CIDR collection location name. See the aws_route53_cidr_location resource for more details. A location_name with an asterisk "*" can be used to create a default CIDR record. collection_id is still required for default record.
	LocationName *string `json:"locationName,omitempty" tf:"location_name,omitempty"`
}

type CidrRoutingPolicyObservation struct {

	// The CIDR collection ID. See the aws_route53_cidr_collection resource for more details.
	CollectionID *string `json:"collectionId,omitempty" tf:"collection_id,omitempty"`

	// The CIDR collection location name. See the aws_route53_cidr_location resource for more details. A location_name with an asterisk "*" can be used to create a default CIDR record. collection_id is still required for default record.
	LocationName *string `json:"locationName,omitempty" tf:"location_name,omitempty"`
}

type CidrRoutingPolicyParameters struct {

	// The CIDR collection ID. See the aws_route53_cidr_collection resource for more details.
	// +kubebuilder:validation:Optional
	CollectionID *string `json:"collectionId" tf:"collection_id,omitempty"`

	// The CIDR collection location name. See the aws_route53_cidr_location resource for more details. A location_name with an asterisk "*" can be used to create a default CIDR record. collection_id is still required for default record.
	// +kubebuilder:validation:Optional
	LocationName *string `json:"locationName" tf:"location_name,omitempty"`
}

type CoordinatesInitParameters struct {
	Latitude *string `json:"latitude,omitempty" tf:"latitude,omitempty"`

	Longitude *string `json:"longitude,omitempty" tf:"longitude,omitempty"`
}

type CoordinatesObservation struct {
	Latitude *string `json:"latitude,omitempty" tf:"latitude,omitempty"`

	Longitude *string `json:"longitude,omitempty" tf:"longitude,omitempty"`
}

type CoordinatesParameters struct {

	// +kubebuilder:validation:Optional
	Latitude *string `json:"latitude" tf:"latitude,omitempty"`

	// +kubebuilder:validation:Optional
	Longitude *string `json:"longitude" tf:"longitude,omitempty"`
}

type FailoverRoutingPolicyInitParameters struct {

	// The record type. Valid values are A, AAAA, CAA, CNAME, DS, MX, NAPTR, NS, PTR, SOA, SPF, SRV and TXT.
	Type *string `json:"type,omitempty" tf:"type,omitempty"`
}

type FailoverRoutingPolicyObservation struct {

	// The record type. Valid values are A, AAAA, CAA, CNAME, DS, MX, NAPTR, NS, PTR, SOA, SPF, SRV and TXT.
	Type *string `json:"type,omitempty" tf:"type,omitempty"`
}

type FailoverRoutingPolicyParameters struct {

	// The record type. Valid values are A, AAAA, CAA, CNAME, DS, MX, NAPTR, NS, PTR, SOA, SPF, SRV and TXT.
	// +kubebuilder:validation:Optional
	Type *string `json:"type" tf:"type,omitempty"`
}

type GeolocationRoutingPolicyInitParameters struct {

	// A two-letter continent code. See http://docs.aws.amazon.com/Route53/latest/APIReference/API_GetGeoLocation.html for code details. Either continent or country must be specified.
	Continent *string `json:"continent,omitempty" tf:"continent,omitempty"`

	// A two-character country code or * to indicate a default resource record set.
	Country *string `json:"country,omitempty" tf:"country,omitempty"`

	// A subdivision code for a country.
	Subdivision *string `json:"subdivision,omitempty" tf:"subdivision,omitempty"`
}

type GeolocationRoutingPolicyObservation struct {

	// A two-letter continent code. See http://docs.aws.amazon.com/Route53/latest/APIReference/API_GetGeoLocation.html for code details. Either continent or country must be specified.
	Continent *string `json:"continent,omitempty" tf:"continent,omitempty"`

	// A two-character country code or * to indicate a default resource record set.
	Country *string `json:"country,omitempty" tf:"country,omitempty"`

	// A subdivision code for a country.
	Subdivision *string `json:"subdivision,omitempty" tf:"subdivision,omitempty"`
}

type GeolocationRoutingPolicyParameters struct {

	// A two-letter continent code. See http://docs.aws.amazon.com/Route53/latest/APIReference/API_GetGeoLocation.html for code details. Either continent or country must be specified.
	// +kubebuilder:validation:Optional
	Continent *string `json:"continent,omitempty" tf:"continent,omitempty"`

	// A two-character country code or * to indicate a default resource record set.
	// +kubebuilder:validation:Optional
	Country *string `json:"country,omitempty" tf:"country,omitempty"`

	// A subdivision code for a country.
	// +kubebuilder:validation:Optional
	Subdivision *string `json:"subdivision,omitempty" tf:"subdivision,omitempty"`
}

type GeoproximityRoutingPolicyInitParameters struct {

	// A AWS region where the resource is present.
	AwsRegion *string `json:"awsRegion,omitempty" tf:"aws_region,omitempty"`

	// Route more traffic or less traffic to the resource by specifying a value ranges between -90 to 90. See https://docs.aws.amazon.com/Route53/latest/DeveloperGuide/routing-policy-geoproximity.html for bias details.
	Bias *float64 `json:"bias,omitempty" tf:"bias,omitempty"`

	// Specify latitude and longitude for routing traffic to non-AWS resources.
	Coordinates []CoordinatesInitParameters `json:"coordinates,omitempty" tf:"coordinates,omitempty"`

	// A AWS local zone group where the resource is present. See https://docs.aws.amazon.com/local-zones/latest/ug/available-local-zones.html for local zone group list.
	LocalZoneGroup *string `json:"localZoneGroup,omitempty" tf:"local_zone_group,omitempty"`
}

type GeoproximityRoutingPolicyObservation struct {

	// A AWS region where the resource is present.
	AwsRegion *string `json:"awsRegion,omitempty" tf:"aws_region,omitempty"`

	// Route more traffic or less traffic to the resource by specifying a value ranges between -90 to 90. See https://docs.aws.amazon.com/Route53/latest/DeveloperGuide/routing-policy-geoproximity.html for bias details.
	Bias *float64 `json:"bias,omitempty" tf:"bias,omitempty"`

	// Specify latitude and longitude for routing traffic to non-AWS resources.
	Coordinates []CoordinatesObservation `json:"coordinates,omitempty" tf:"coordinates,omitempty"`

	// A AWS local zone group where the resource is present. See https://docs.aws.amazon.com/local-zones/latest/ug/available-local-zones.html for local zone group list.
	LocalZoneGroup *string `json:"localZoneGroup,omitempty" tf:"local_zone_group,omitempty"`
}

type GeoproximityRoutingPolicyParameters struct {

	// A AWS region where the resource is present.
	// +kubebuilder:validation:Optional
	AwsRegion *string `json:"awsRegion,omitempty" tf:"aws_region,omitempty"`

	// Route more traffic or less traffic to the resource by specifying a value ranges between -90 to 90. See https://docs.aws.amazon.com/Route53/latest/DeveloperGuide/routing-policy-geoproximity.html for bias details.
	// +kubebuilder:validation:Optional
	Bias *float64 `json:"bias,omitempty" tf:"bias,omitempty"`

	// Specify latitude and longitude for routing traffic to non-AWS resources.
	// +kubebuilder:validation:Optional
	Coordinates []CoordinatesParameters `json:"coordinates,omitempty" tf:"coordinates,omitempty"`

	// A AWS local zone group where the resource is present. See https://docs.aws.amazon.com/local-zones/latest/ug/available-local-zones.html for local zone group list.
	// +kubebuilder:validation:Optional
	LocalZoneGroup *string `json:"localZoneGroup,omitempty" tf:"local_zone_group,omitempty"`
}

type LatencyRoutingPolicyInitParameters struct {
}

type LatencyRoutingPolicyObservation struct {

	// An AWS region from which to measure latency. See http://docs.aws.amazon.com/Route53/latest/DeveloperGuide/routing-policy.html#routing-policy-latency
	Region *string `json:"region,omitempty" tf:"region,omitempty"`
}

type LatencyRoutingPolicyParameters struct {

	// An AWS region from which to measure latency. See http://docs.aws.amazon.com/Route53/latest/DeveloperGuide/routing-policy.html#routing-policy-latency
	// +kubebuilder:validation:Required
	Region *string `json:"region" tf:"region,omitempty"`
}

type RecordInitParameters struct {

	// An alias block. Conflicts with ttl & records.
	// Documented below.
	Alias []AliasInitParameters `json:"alias,omitempty" tf:"alias,omitempty"`

	// false by default. This configuration is not recommended for most environments.
	AllowOverwrite *bool `json:"allowOverwrite,omitempty" tf:"allow_overwrite,omitempty"`

	// A block indicating a routing policy based on the IP network ranges of requestors. Conflicts with any other routing policy. Documented below.
	CidrRoutingPolicy []CidrRoutingPolicyInitParameters `json:"cidrRoutingPolicy,omitempty" tf:"cidr_routing_policy,omitempty"`

	// A block indicating the routing behavior when associated health check fails. Conflicts with any other routing policy. Documented below.
	FailoverRoutingPolicy []FailoverRoutingPolicyInitParameters `json:"failoverRoutingPolicy,omitempty" tf:"failover_routing_policy,omitempty"`

	// A block indicating a routing policy based on the geolocation of the requestor. Conflicts with any other routing policy. Documented below.
	GeolocationRoutingPolicy []GeolocationRoutingPolicyInitParameters `json:"geolocationRoutingPolicy,omitempty" tf:"geolocation_routing_policy,omitempty"`

	// A block indicating a routing policy based on the geoproximity of the requestor. Conflicts with any other routing policy. Documented below.
	GeoproximityRoutingPolicy []GeoproximityRoutingPolicyInitParameters `json:"geoproximityRoutingPolicy,omitempty" tf:"geoproximity_routing_policy,omitempty"`

	// The health check the record should be associated with.
	// +crossplane:generate:reference:type=github.com/upbound/provider-aws/apis/route53/v1beta1.HealthCheck
	HealthCheckID *string `json:"healthCheckId,omitempty" tf:"health_check_id,omitempty"`

	// Reference to a HealthCheck in route53 to populate healthCheckId.
	// +kubebuilder:validation:Optional
	HealthCheckIDRef *v1.Reference `json:"healthCheckIdRef,omitempty" tf:"-"`

	// Selector for a HealthCheck in route53 to populate healthCheckId.
	// +kubebuilder:validation:Optional
	HealthCheckIDSelector *v1.Selector `json:"healthCheckIdSelector,omitempty" tf:"-"`

	// A block indicating a routing policy based on the latency between the requestor and an AWS region. Conflicts with any other routing policy. Documented below.
	LatencyRoutingPolicy []LatencyRoutingPolicyInitParameters `json:"latencyRoutingPolicy,omitempty" tf:"latency_routing_policy,omitempty"`

	// Set to true to indicate a multivalue answer routing policy. Conflicts with any other routing policy.
	MultivalueAnswerRoutingPolicy *bool `json:"multivalueAnswerRoutingPolicy,omitempty" tf:"multivalue_answer_routing_policy,omitempty"`

	// The name of the record.
	Name *string `json:"name,omitempty" tf:"name,omitempty"`

	// A string list of records.g., "first255characters\"\"morecharacters").
	// +crossplane:generate:reference:type=github.com/upbound/provider-aws/apis/ec2/v1beta1.EIP
	// +crossplane:generate:reference:extractor=github.com/crossplane/upjet/pkg/resource.ExtractParamPath("public_ip",true)
	// +listType=set
	Records []*string `json:"records,omitempty" tf:"records,omitempty"`

	// References to EIP in ec2 to populate records.
	// +kubebuilder:validation:Optional
	RecordsRefs []v1.Reference `json:"recordsRefs,omitempty" tf:"-"`

	// Selector for a list of EIP in ec2 to populate records.
	// +kubebuilder:validation:Optional
	RecordsSelector *v1.Selector `json:"recordsSelector,omitempty" tf:"-"`

	// Unique identifier to differentiate records with routing policies from one another. Required if using cidr_routing_policy, failover_routing_policy, geolocation_routing_policy,geoproximity_routing_policy, latency_routing_policy, multivalue_answer_routing_policy, or weighted_routing_policy.
	SetIdentifier *string `json:"setIdentifier,omitempty" tf:"set_identifier,omitempty"`

	// The TTL of the record.
	TTL *float64 `json:"ttl,omitempty" tf:"ttl,omitempty"`

	// The record type. Valid values are A, AAAA, CAA, CNAME, DS, MX, NAPTR, NS, PTR, SOA, SPF, SRV and TXT.
	Type *string `json:"type,omitempty" tf:"type,omitempty"`

	// A block indicating a weighted routing policy. Conflicts with any other routing policy. Documented below.
	WeightedRoutingPolicy []WeightedRoutingPolicyInitParameters `json:"weightedRoutingPolicy,omitempty" tf:"weighted_routing_policy,omitempty"`

	// The ID of the hosted zone to contain this record.
	// +crossplane:generate:reference:type=github.com/upbound/provider-aws/apis/route53/v1beta1.Zone
	ZoneID *string `json:"zoneId,omitempty" tf:"zone_id,omitempty"`

	// Reference to a Zone in route53 to populate zoneId.
	// +kubebuilder:validation:Optional
	ZoneIDRef *v1.Reference `json:"zoneIdRef,omitempty" tf:"-"`

	// Selector for a Zone in route53 to populate zoneId.
	// +kubebuilder:validation:Optional
	ZoneIDSelector *v1.Selector `json:"zoneIdSelector,omitempty" tf:"-"`
}

type RecordObservation struct {

	// An alias block. Conflicts with ttl & records.
	// Documented below.
	Alias []AliasObservation `json:"alias,omitempty" tf:"alias,omitempty"`

	// false by default. This configuration is not recommended for most environments.
	AllowOverwrite *bool `json:"allowOverwrite,omitempty" tf:"allow_overwrite,omitempty"`

	// A block indicating a routing policy based on the IP network ranges of requestors. Conflicts with any other routing policy. Documented below.
	CidrRoutingPolicy []CidrRoutingPolicyObservation `json:"cidrRoutingPolicy,omitempty" tf:"cidr_routing_policy,omitempty"`

	// A block indicating the routing behavior when associated health check fails. Conflicts with any other routing policy. Documented below.
	FailoverRoutingPolicy []FailoverRoutingPolicyObservation `json:"failoverRoutingPolicy,omitempty" tf:"failover_routing_policy,omitempty"`

	// FQDN built using the zone domain and name.
	Fqdn *string `json:"fqdn,omitempty" tf:"fqdn,omitempty"`

	// A block indicating a routing policy based on the geolocation of the requestor. Conflicts with any other routing policy. Documented below.
	GeolocationRoutingPolicy []GeolocationRoutingPolicyObservation `json:"geolocationRoutingPolicy,omitempty" tf:"geolocation_routing_policy,omitempty"`

	// A block indicating a routing policy based on the geoproximity of the requestor. Conflicts with any other routing policy. Documented below.
	GeoproximityRoutingPolicy []GeoproximityRoutingPolicyObservation `json:"geoproximityRoutingPolicy,omitempty" tf:"geoproximity_routing_policy,omitempty"`

	// The health check the record should be associated with.
	HealthCheckID *string `json:"healthCheckId,omitempty" tf:"health_check_id,omitempty"`

	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	// A block indicating a routing policy based on the latency between the requestor and an AWS region. Conflicts with any other routing policy. Documented below.
	LatencyRoutingPolicy []LatencyRoutingPolicyObservation `json:"latencyRoutingPolicy,omitempty" tf:"latency_routing_policy,omitempty"`

	// Set to true to indicate a multivalue answer routing policy. Conflicts with any other routing policy.
	MultivalueAnswerRoutingPolicy *bool `json:"multivalueAnswerRoutingPolicy,omitempty" tf:"multivalue_answer_routing_policy,omitempty"`

	// The name of the record.
	Name *string `json:"name,omitempty" tf:"name,omitempty"`

	// A string list of records.g., "first255characters\"\"morecharacters").
	// +listType=set
	Records []*string `json:"records,omitempty" tf:"records,omitempty"`

	// Unique identifier to differentiate records with routing policies from one another. Required if using cidr_routing_policy, failover_routing_policy, geolocation_routing_policy,geoproximity_routing_policy, latency_routing_policy, multivalue_answer_routing_policy, or weighted_routing_policy.
	SetIdentifier *string `json:"setIdentifier,omitempty" tf:"set_identifier,omitempty"`

	// The TTL of the record.
	TTL *float64 `json:"ttl,omitempty" tf:"ttl,omitempty"`

	// The record type. Valid values are A, AAAA, CAA, CNAME, DS, MX, NAPTR, NS, PTR, SOA, SPF, SRV and TXT.
	Type *string `json:"type,omitempty" tf:"type,omitempty"`

	// A block indicating a weighted routing policy. Conflicts with any other routing policy. Documented below.
	WeightedRoutingPolicy []WeightedRoutingPolicyObservation `json:"weightedRoutingPolicy,omitempty" tf:"weighted_routing_policy,omitempty"`

	// The ID of the hosted zone to contain this record.
	ZoneID *string `json:"zoneId,omitempty" tf:"zone_id,omitempty"`
}

type RecordParameters struct {

	// An alias block. Conflicts with ttl & records.
	// Documented below.
	// +kubebuilder:validation:Optional
	Alias []AliasParameters `json:"alias,omitempty" tf:"alias,omitempty"`

	// false by default. This configuration is not recommended for most environments.
	// +kubebuilder:validation:Optional
	AllowOverwrite *bool `json:"allowOverwrite,omitempty" tf:"allow_overwrite,omitempty"`

	// A block indicating a routing policy based on the IP network ranges of requestors. Conflicts with any other routing policy. Documented below.
	// +kubebuilder:validation:Optional
	CidrRoutingPolicy []CidrRoutingPolicyParameters `json:"cidrRoutingPolicy,omitempty" tf:"cidr_routing_policy,omitempty"`

	// A block indicating the routing behavior when associated health check fails. Conflicts with any other routing policy. Documented below.
	// +kubebuilder:validation:Optional
	FailoverRoutingPolicy []FailoverRoutingPolicyParameters `json:"failoverRoutingPolicy,omitempty" tf:"failover_routing_policy,omitempty"`

	// A block indicating a routing policy based on the geolocation of the requestor. Conflicts with any other routing policy. Documented below.
	// +kubebuilder:validation:Optional
	GeolocationRoutingPolicy []GeolocationRoutingPolicyParameters `json:"geolocationRoutingPolicy,omitempty" tf:"geolocation_routing_policy,omitempty"`

	// A block indicating a routing policy based on the geoproximity of the requestor. Conflicts with any other routing policy. Documented below.
	// +kubebuilder:validation:Optional
	GeoproximityRoutingPolicy []GeoproximityRoutingPolicyParameters `json:"geoproximityRoutingPolicy,omitempty" tf:"geoproximity_routing_policy,omitempty"`

	// The health check the record should be associated with.
	// +crossplane:generate:reference:type=github.com/upbound/provider-aws/apis/route53/v1beta1.HealthCheck
	// +kubebuilder:validation:Optional
	HealthCheckID *string `json:"healthCheckId,omitempty" tf:"health_check_id,omitempty"`

	// Reference to a HealthCheck in route53 to populate healthCheckId.
	// +kubebuilder:validation:Optional
	HealthCheckIDRef *v1.Reference `json:"healthCheckIdRef,omitempty" tf:"-"`

	// Selector for a HealthCheck in route53 to populate healthCheckId.
	// +kubebuilder:validation:Optional
	HealthCheckIDSelector *v1.Selector `json:"healthCheckIdSelector,omitempty" tf:"-"`

	// A block indicating a routing policy based on the latency between the requestor and an AWS region. Conflicts with any other routing policy. Documented below.
	// +kubebuilder:validation:Optional
	LatencyRoutingPolicy []LatencyRoutingPolicyParameters `json:"latencyRoutingPolicy,omitempty" tf:"latency_routing_policy,omitempty"`

	// Set to true to indicate a multivalue answer routing policy. Conflicts with any other routing policy.
	// +kubebuilder:validation:Optional
	MultivalueAnswerRoutingPolicy *bool `json:"multivalueAnswerRoutingPolicy,omitempty" tf:"multivalue_answer_routing_policy,omitempty"`

	// The name of the record.
	// +kubebuilder:validation:Optional
	Name *string `json:"name,omitempty" tf:"name,omitempty"`

	// A string list of records.g., "first255characters\"\"morecharacters").
	// +crossplane:generate:reference:type=github.com/upbound/provider-aws/apis/ec2/v1beta1.EIP
	// +crossplane:generate:reference:extractor=github.com/crossplane/upjet/pkg/resource.ExtractParamPath("public_ip",true)
	// +kubebuilder:validation:Optional
	// +listType=set
	Records []*string `json:"records,omitempty" tf:"records,omitempty"`

	// References to EIP in ec2 to populate records.
	// +kubebuilder:validation:Optional
	RecordsRefs []v1.Reference `json:"recordsRefs,omitempty" tf:"-"`

	// Selector for a list of EIP in ec2 to populate records.
	// +kubebuilder:validation:Optional
	RecordsSelector *v1.Selector `json:"recordsSelector,omitempty" tf:"-"`

	// An AWS region from which to measure latency. See http://docs.aws.amazon.com/Route53/latest/DeveloperGuide/routing-policy.html#routing-policy-latency
	// Region is the region you'd like your resource to be created in.
	// +upjet:crd:field:TFTag=-
	// +kubebuilder:validation:Required
	Region *string `json:"region" tf:"-"`

	// Unique identifier to differentiate records with routing policies from one another. Required if using cidr_routing_policy, failover_routing_policy, geolocation_routing_policy,geoproximity_routing_policy, latency_routing_policy, multivalue_answer_routing_policy, or weighted_routing_policy.
	// +kubebuilder:validation:Optional
	SetIdentifier *string `json:"setIdentifier,omitempty" tf:"set_identifier,omitempty"`

	// The TTL of the record.
	// +kubebuilder:validation:Optional
	TTL *float64 `json:"ttl,omitempty" tf:"ttl,omitempty"`

	// The record type. Valid values are A, AAAA, CAA, CNAME, DS, MX, NAPTR, NS, PTR, SOA, SPF, SRV and TXT.
	// +kubebuilder:validation:Optional
	Type *string `json:"type,omitempty" tf:"type,omitempty"`

	// A block indicating a weighted routing policy. Conflicts with any other routing policy. Documented below.
	// +kubebuilder:validation:Optional
	WeightedRoutingPolicy []WeightedRoutingPolicyParameters `json:"weightedRoutingPolicy,omitempty" tf:"weighted_routing_policy,omitempty"`

	// The ID of the hosted zone to contain this record.
	// +crossplane:generate:reference:type=github.com/upbound/provider-aws/apis/route53/v1beta1.Zone
	// +kubebuilder:validation:Optional
	ZoneID *string `json:"zoneId,omitempty" tf:"zone_id,omitempty"`

	// Reference to a Zone in route53 to populate zoneId.
	// +kubebuilder:validation:Optional
	ZoneIDRef *v1.Reference `json:"zoneIdRef,omitempty" tf:"-"`

	// Selector for a Zone in route53 to populate zoneId.
	// +kubebuilder:validation:Optional
	ZoneIDSelector *v1.Selector `json:"zoneIdSelector,omitempty" tf:"-"`
}

type WeightedRoutingPolicyInitParameters struct {

	// A numeric value indicating the relative weight of the record. See http://docs.aws.amazon.com/Route53/latest/DeveloperGuide/routing-policy.html#routing-policy-weighted.
	Weight *float64 `json:"weight,omitempty" tf:"weight,omitempty"`
}

type WeightedRoutingPolicyObservation struct {

	// A numeric value indicating the relative weight of the record. See http://docs.aws.amazon.com/Route53/latest/DeveloperGuide/routing-policy.html#routing-policy-weighted.
	Weight *float64 `json:"weight,omitempty" tf:"weight,omitempty"`
}

type WeightedRoutingPolicyParameters struct {

	// A numeric value indicating the relative weight of the record. See http://docs.aws.amazon.com/Route53/latest/DeveloperGuide/routing-policy.html#routing-policy-weighted.
	// +kubebuilder:validation:Optional
	Weight *float64 `json:"weight" tf:"weight,omitempty"`
}

// RecordSpec defines the desired state of Record
type RecordSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     RecordParameters `json:"forProvider"`
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
	InitProvider RecordInitParameters `json:"initProvider,omitempty"`
}

// RecordStatus defines the observed state of Record.
type RecordStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        RecordObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true
// +kubebuilder:subresource:status
// +kubebuilder:storageversion

// Record is the Schema for the Records API. Provides a Route53 record resource.
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,aws}
type Record struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.name) || (has(self.initProvider) && has(self.initProvider.name))",message="spec.forProvider.name is a required parameter"
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.type) || (has(self.initProvider) && has(self.initProvider.type))",message="spec.forProvider.type is a required parameter"
	Spec   RecordSpec   `json:"spec"`
	Status RecordStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// RecordList contains a list of Records
type RecordList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []Record `json:"items"`
}

// Repository type metadata.
var (
	Record_Kind             = "Record"
	Record_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: Record_Kind}.String()
	Record_KindAPIVersion   = Record_Kind + "." + CRDGroupVersion.String()
	Record_GroupVersionKind = CRDGroupVersion.WithKind(Record_Kind)
)

func init() {
	SchemeBuilder.Register(&Record{}, &RecordList{})
}
