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

type ConfigRuleInitParameters struct {

	// Description of the rule
	Description *string `json:"description,omitempty" tf:"description,omitempty"`

	// The modes the Config rule can be evaluated in. See Evaluation Mode for more details.
	EvaluationMode []EvaluationModeInitParameters `json:"evaluationMode,omitempty" tf:"evaluation_mode,omitempty"`

	// A string in JSON format that is passed to the AWS Config rule Lambda function.
	InputParameters *string `json:"inputParameters,omitempty" tf:"input_parameters,omitempty"`

	// The maximum frequency with which AWS Config runs evaluations for a rule.
	MaximumExecutionFrequency *string `json:"maximumExecutionFrequency,omitempty" tf:"maximum_execution_frequency,omitempty"`

	// Scope defines which resources can trigger an evaluation for the rule. See Source Below.
	Scope []ScopeInitParameters `json:"scope,omitempty" tf:"scope,omitempty"`

	// Source specifies the rule owner, the rule identifier, and the notifications that cause the function to evaluate your AWS resources. See Scope Below.
	Source []SourceInitParameters `json:"source,omitempty" tf:"source,omitempty"`

	// Key-value map of resource tags.
	// +mapType=granular
	Tags map[string]*string `json:"tags,omitempty" tf:"tags,omitempty"`
}

type ConfigRuleObservation struct {

	// The ARN of the config rule
	Arn *string `json:"arn,omitempty" tf:"arn,omitempty"`

	// Description of the rule
	Description *string `json:"description,omitempty" tf:"description,omitempty"`

	// The modes the Config rule can be evaluated in. See Evaluation Mode for more details.
	EvaluationMode []EvaluationModeObservation `json:"evaluationMode,omitempty" tf:"evaluation_mode,omitempty"`

	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	// A string in JSON format that is passed to the AWS Config rule Lambda function.
	InputParameters *string `json:"inputParameters,omitempty" tf:"input_parameters,omitempty"`

	// The maximum frequency with which AWS Config runs evaluations for a rule.
	MaximumExecutionFrequency *string `json:"maximumExecutionFrequency,omitempty" tf:"maximum_execution_frequency,omitempty"`

	// The ID of the config rule
	RuleID *string `json:"ruleId,omitempty" tf:"rule_id,omitempty"`

	// Scope defines which resources can trigger an evaluation for the rule. See Source Below.
	Scope []ScopeObservation `json:"scope,omitempty" tf:"scope,omitempty"`

	// Source specifies the rule owner, the rule identifier, and the notifications that cause the function to evaluate your AWS resources. See Scope Below.
	Source []SourceObservation `json:"source,omitempty" tf:"source,omitempty"`

	// Key-value map of resource tags.
	// +mapType=granular
	Tags map[string]*string `json:"tags,omitempty" tf:"tags,omitempty"`

	// A map of tags assigned to the resource, including those inherited from the provider default_tags configuration block.
	// +mapType=granular
	TagsAll map[string]*string `json:"tagsAll,omitempty" tf:"tags_all,omitempty"`
}

type ConfigRuleParameters struct {

	// Description of the rule
	// +kubebuilder:validation:Optional
	Description *string `json:"description,omitempty" tf:"description,omitempty"`

	// The modes the Config rule can be evaluated in. See Evaluation Mode for more details.
	// +kubebuilder:validation:Optional
	EvaluationMode []EvaluationModeParameters `json:"evaluationMode,omitempty" tf:"evaluation_mode,omitempty"`

	// A string in JSON format that is passed to the AWS Config rule Lambda function.
	// +kubebuilder:validation:Optional
	InputParameters *string `json:"inputParameters,omitempty" tf:"input_parameters,omitempty"`

	// The maximum frequency with which AWS Config runs evaluations for a rule.
	// +kubebuilder:validation:Optional
	MaximumExecutionFrequency *string `json:"maximumExecutionFrequency,omitempty" tf:"maximum_execution_frequency,omitempty"`

	// Region is the region you'd like your resource to be created in.
	// +upjet:crd:field:TFTag=-
	// +kubebuilder:validation:Required
	Region *string `json:"region" tf:"-"`

	// Scope defines which resources can trigger an evaluation for the rule. See Source Below.
	// +kubebuilder:validation:Optional
	Scope []ScopeParameters `json:"scope,omitempty" tf:"scope,omitempty"`

	// Source specifies the rule owner, the rule identifier, and the notifications that cause the function to evaluate your AWS resources. See Scope Below.
	// +kubebuilder:validation:Optional
	Source []SourceParameters `json:"source,omitempty" tf:"source,omitempty"`

	// Key-value map of resource tags.
	// +kubebuilder:validation:Optional
	// +mapType=granular
	Tags map[string]*string `json:"tags,omitempty" tf:"tags,omitempty"`
}

type CustomPolicyDetailsInitParameters struct {

	// The boolean expression for enabling debug logging for your Config Custom Policy rule. The default value is false.
	EnableDebugLogDelivery *bool `json:"enableDebugLogDelivery,omitempty" tf:"enable_debug_log_delivery,omitempty"`

	// The runtime system for your Config Custom Policy rule. Guard is a policy-as-code language that allows you to write policies that are enforced by Config Custom Policy rules. For more information about Guard, see the Guard GitHub Repository.
	PolicyRuntime *string `json:"policyRuntime,omitempty" tf:"policy_runtime,omitempty"`

	// The policy definition containing the logic for your Config Custom Policy rule.
	PolicyText *string `json:"policyText,omitempty" tf:"policy_text,omitempty"`
}

type CustomPolicyDetailsObservation struct {

	// The boolean expression for enabling debug logging for your Config Custom Policy rule. The default value is false.
	EnableDebugLogDelivery *bool `json:"enableDebugLogDelivery,omitempty" tf:"enable_debug_log_delivery,omitempty"`

	// The runtime system for your Config Custom Policy rule. Guard is a policy-as-code language that allows you to write policies that are enforced by Config Custom Policy rules. For more information about Guard, see the Guard GitHub Repository.
	PolicyRuntime *string `json:"policyRuntime,omitempty" tf:"policy_runtime,omitempty"`

	// The policy definition containing the logic for your Config Custom Policy rule.
	PolicyText *string `json:"policyText,omitempty" tf:"policy_text,omitempty"`
}

type CustomPolicyDetailsParameters struct {

	// The boolean expression for enabling debug logging for your Config Custom Policy rule. The default value is false.
	// +kubebuilder:validation:Optional
	EnableDebugLogDelivery *bool `json:"enableDebugLogDelivery,omitempty" tf:"enable_debug_log_delivery,omitempty"`

	// The runtime system for your Config Custom Policy rule. Guard is a policy-as-code language that allows you to write policies that are enforced by Config Custom Policy rules. For more information about Guard, see the Guard GitHub Repository.
	// +kubebuilder:validation:Optional
	PolicyRuntime *string `json:"policyRuntime" tf:"policy_runtime,omitempty"`

	// The policy definition containing the logic for your Config Custom Policy rule.
	// +kubebuilder:validation:Optional
	PolicyText *string `json:"policyText" tf:"policy_text,omitempty"`
}

type EvaluationModeInitParameters struct {

	// The mode of an evaluation.
	Mode *string `json:"mode,omitempty" tf:"mode,omitempty"`
}

type EvaluationModeObservation struct {

	// The mode of an evaluation.
	Mode *string `json:"mode,omitempty" tf:"mode,omitempty"`
}

type EvaluationModeParameters struct {

	// The mode of an evaluation.
	// +kubebuilder:validation:Optional
	Mode *string `json:"mode,omitempty" tf:"mode,omitempty"`
}

type ScopeInitParameters struct {

	// The IDs of the only AWS resource that you want to trigger an evaluation for the rule. If you specify a resource ID, you must specify one resource type for compliance_resource_types.
	ComplianceResourceID *string `json:"complianceResourceId,omitempty" tf:"compliance_resource_id,omitempty"`

	// A list of resource types of only those AWS resources that you want to trigger an evaluation for the ruleE.g., AWS::EC2::Instance. You can only specify one type if you also specify a resource ID for compliance_resource_id. See relevant part of AWS Docs for available types.
	// +listType=set
	ComplianceResourceTypes []*string `json:"complianceResourceTypes,omitempty" tf:"compliance_resource_types,omitempty"`

	// The tag key that is applied to only those AWS resources that you want you want to trigger an evaluation for the rule.
	TagKey *string `json:"tagKey,omitempty" tf:"tag_key,omitempty"`

	// The tag value applied to only those AWS resources that you want to trigger an evaluation for the rule.
	TagValue *string `json:"tagValue,omitempty" tf:"tag_value,omitempty"`
}

type ScopeObservation struct {

	// The IDs of the only AWS resource that you want to trigger an evaluation for the rule. If you specify a resource ID, you must specify one resource type for compliance_resource_types.
	ComplianceResourceID *string `json:"complianceResourceId,omitempty" tf:"compliance_resource_id,omitempty"`

	// A list of resource types of only those AWS resources that you want to trigger an evaluation for the ruleE.g., AWS::EC2::Instance. You can only specify one type if you also specify a resource ID for compliance_resource_id. See relevant part of AWS Docs for available types.
	// +listType=set
	ComplianceResourceTypes []*string `json:"complianceResourceTypes,omitempty" tf:"compliance_resource_types,omitempty"`

	// The tag key that is applied to only those AWS resources that you want you want to trigger an evaluation for the rule.
	TagKey *string `json:"tagKey,omitempty" tf:"tag_key,omitempty"`

	// The tag value applied to only those AWS resources that you want to trigger an evaluation for the rule.
	TagValue *string `json:"tagValue,omitempty" tf:"tag_value,omitempty"`
}

type ScopeParameters struct {

	// The IDs of the only AWS resource that you want to trigger an evaluation for the rule. If you specify a resource ID, you must specify one resource type for compliance_resource_types.
	// +kubebuilder:validation:Optional
	ComplianceResourceID *string `json:"complianceResourceId,omitempty" tf:"compliance_resource_id,omitempty"`

	// A list of resource types of only those AWS resources that you want to trigger an evaluation for the ruleE.g., AWS::EC2::Instance. You can only specify one type if you also specify a resource ID for compliance_resource_id. See relevant part of AWS Docs for available types.
	// +kubebuilder:validation:Optional
	// +listType=set
	ComplianceResourceTypes []*string `json:"complianceResourceTypes,omitempty" tf:"compliance_resource_types,omitempty"`

	// The tag key that is applied to only those AWS resources that you want you want to trigger an evaluation for the rule.
	// +kubebuilder:validation:Optional
	TagKey *string `json:"tagKey,omitempty" tf:"tag_key,omitempty"`

	// The tag value applied to only those AWS resources that you want to trigger an evaluation for the rule.
	// +kubebuilder:validation:Optional
	TagValue *string `json:"tagValue,omitempty" tf:"tag_value,omitempty"`
}

type SourceDetailInitParameters struct {

	// The source of the event, such as an AWS service, that triggers AWS Config to evaluate your AWSresources. This defaults to aws.config and is the only valid value.
	EventSource *string `json:"eventSource,omitempty" tf:"event_source,omitempty"`

	// The maximum frequency with which AWS Config runs evaluations for a rule.
	MaximumExecutionFrequency *string `json:"maximumExecutionFrequency,omitempty" tf:"maximum_execution_frequency,omitempty"`

	// The type of notification that triggers AWS Config to run an evaluation for a rule. You canspecify the following notification types:
	MessageType *string `json:"messageType,omitempty" tf:"message_type,omitempty"`
}

type SourceDetailObservation struct {

	// The source of the event, such as an AWS service, that triggers AWS Config to evaluate your AWSresources. This defaults to aws.config and is the only valid value.
	EventSource *string `json:"eventSource,omitempty" tf:"event_source,omitempty"`

	// The maximum frequency with which AWS Config runs evaluations for a rule.
	MaximumExecutionFrequency *string `json:"maximumExecutionFrequency,omitempty" tf:"maximum_execution_frequency,omitempty"`

	// The type of notification that triggers AWS Config to run an evaluation for a rule. You canspecify the following notification types:
	MessageType *string `json:"messageType,omitempty" tf:"message_type,omitempty"`
}

type SourceDetailParameters struct {

	// The source of the event, such as an AWS service, that triggers AWS Config to evaluate your AWSresources. This defaults to aws.config and is the only valid value.
	// +kubebuilder:validation:Optional
	EventSource *string `json:"eventSource,omitempty" tf:"event_source,omitempty"`

	// The maximum frequency with which AWS Config runs evaluations for a rule.
	// +kubebuilder:validation:Optional
	MaximumExecutionFrequency *string `json:"maximumExecutionFrequency,omitempty" tf:"maximum_execution_frequency,omitempty"`

	// The type of notification that triggers AWS Config to run an evaluation for a rule. You canspecify the following notification types:
	// +kubebuilder:validation:Optional
	MessageType *string `json:"messageType,omitempty" tf:"message_type,omitempty"`
}

type SourceInitParameters struct {

	// Provides the runtime system, policy definition, and whether debug logging is enabled. Required when owner is set to CUSTOM_POLICY. See Custom Policy Details Below.
	CustomPolicyDetails []CustomPolicyDetailsInitParameters `json:"customPolicyDetails,omitempty" tf:"custom_policy_details,omitempty"`

	// Indicates whether AWS or the customer owns and manages the AWS Config rule. Valid values are AWS, CUSTOM_LAMBDA or CUSTOM_POLICY. For more information about managed rules, see the AWS Config Managed Rules documentation. For more information about custom rules, see the AWS Config Custom Rules documentation. Custom Lambda Functions require permissions to allow the AWS Config service to invoke them, e.g., via the aws_lambda_permission resource.
	Owner *string `json:"owner,omitempty" tf:"owner,omitempty"`

	// Provides the source and type of the event that causes AWS Config to evaluate your AWS resources. Only valid if owner is CUSTOM_LAMBDA or CUSTOM_POLICY. See Source Detail Below.
	SourceDetail []SourceDetailInitParameters `json:"sourceDetail,omitempty" tf:"source_detail,omitempty"`

	// For AWS Config managed rules, a predefined identifier, e.g IAM_PASSWORD_POLICY. For custom Lambda rules, the identifier is the ARN of the Lambda Function, such as arn:aws:lambda:us-east-1:123456789012:function:custom_rule_name or the arn attribute of the aws_lambda_function resource.
	// +crossplane:generate:reference:type=github.com/upbound/provider-aws/apis/lambda/v1beta1.Function
	// +crossplane:generate:reference:extractor=github.com/crossplane/upjet/pkg/resource.ExtractParamPath("arn",true)
	SourceIdentifier *string `json:"sourceIdentifier,omitempty" tf:"source_identifier,omitempty"`

	// Reference to a Function in lambda to populate sourceIdentifier.
	// +kubebuilder:validation:Optional
	SourceIdentifierRef *v1.Reference `json:"sourceIdentifierRef,omitempty" tf:"-"`

	// Selector for a Function in lambda to populate sourceIdentifier.
	// +kubebuilder:validation:Optional
	SourceIdentifierSelector *v1.Selector `json:"sourceIdentifierSelector,omitempty" tf:"-"`
}

type SourceObservation struct {

	// Provides the runtime system, policy definition, and whether debug logging is enabled. Required when owner is set to CUSTOM_POLICY. See Custom Policy Details Below.
	CustomPolicyDetails []CustomPolicyDetailsObservation `json:"customPolicyDetails,omitempty" tf:"custom_policy_details,omitempty"`

	// Indicates whether AWS or the customer owns and manages the AWS Config rule. Valid values are AWS, CUSTOM_LAMBDA or CUSTOM_POLICY. For more information about managed rules, see the AWS Config Managed Rules documentation. For more information about custom rules, see the AWS Config Custom Rules documentation. Custom Lambda Functions require permissions to allow the AWS Config service to invoke them, e.g., via the aws_lambda_permission resource.
	Owner *string `json:"owner,omitempty" tf:"owner,omitempty"`

	// Provides the source and type of the event that causes AWS Config to evaluate your AWS resources. Only valid if owner is CUSTOM_LAMBDA or CUSTOM_POLICY. See Source Detail Below.
	SourceDetail []SourceDetailObservation `json:"sourceDetail,omitempty" tf:"source_detail,omitempty"`

	// For AWS Config managed rules, a predefined identifier, e.g IAM_PASSWORD_POLICY. For custom Lambda rules, the identifier is the ARN of the Lambda Function, such as arn:aws:lambda:us-east-1:123456789012:function:custom_rule_name or the arn attribute of the aws_lambda_function resource.
	SourceIdentifier *string `json:"sourceIdentifier,omitempty" tf:"source_identifier,omitempty"`
}

type SourceParameters struct {

	// Provides the runtime system, policy definition, and whether debug logging is enabled. Required when owner is set to CUSTOM_POLICY. See Custom Policy Details Below.
	// +kubebuilder:validation:Optional
	CustomPolicyDetails []CustomPolicyDetailsParameters `json:"customPolicyDetails,omitempty" tf:"custom_policy_details,omitempty"`

	// Indicates whether AWS or the customer owns and manages the AWS Config rule. Valid values are AWS, CUSTOM_LAMBDA or CUSTOM_POLICY. For more information about managed rules, see the AWS Config Managed Rules documentation. For more information about custom rules, see the AWS Config Custom Rules documentation. Custom Lambda Functions require permissions to allow the AWS Config service to invoke them, e.g., via the aws_lambda_permission resource.
	// +kubebuilder:validation:Optional
	Owner *string `json:"owner" tf:"owner,omitempty"`

	// Provides the source and type of the event that causes AWS Config to evaluate your AWS resources. Only valid if owner is CUSTOM_LAMBDA or CUSTOM_POLICY. See Source Detail Below.
	// +kubebuilder:validation:Optional
	SourceDetail []SourceDetailParameters `json:"sourceDetail,omitempty" tf:"source_detail,omitempty"`

	// For AWS Config managed rules, a predefined identifier, e.g IAM_PASSWORD_POLICY. For custom Lambda rules, the identifier is the ARN of the Lambda Function, such as arn:aws:lambda:us-east-1:123456789012:function:custom_rule_name or the arn attribute of the aws_lambda_function resource.
	// +crossplane:generate:reference:type=github.com/upbound/provider-aws/apis/lambda/v1beta1.Function
	// +crossplane:generate:reference:extractor=github.com/crossplane/upjet/pkg/resource.ExtractParamPath("arn",true)
	// +kubebuilder:validation:Optional
	SourceIdentifier *string `json:"sourceIdentifier,omitempty" tf:"source_identifier,omitempty"`

	// Reference to a Function in lambda to populate sourceIdentifier.
	// +kubebuilder:validation:Optional
	SourceIdentifierRef *v1.Reference `json:"sourceIdentifierRef,omitempty" tf:"-"`

	// Selector for a Function in lambda to populate sourceIdentifier.
	// +kubebuilder:validation:Optional
	SourceIdentifierSelector *v1.Selector `json:"sourceIdentifierSelector,omitempty" tf:"-"`
}

// ConfigRuleSpec defines the desired state of ConfigRule
type ConfigRuleSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     ConfigRuleParameters `json:"forProvider"`
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
	InitProvider ConfigRuleInitParameters `json:"initProvider,omitempty"`
}

// ConfigRuleStatus defines the observed state of ConfigRule.
type ConfigRuleStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        ConfigRuleObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true
// +kubebuilder:subresource:status
// +kubebuilder:storageversion

// ConfigRule is the Schema for the ConfigRules API. Provides an AWS Config Rule.
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,aws}
type ConfigRule struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.source) || (has(self.initProvider) && has(self.initProvider.source))",message="spec.forProvider.source is a required parameter"
	Spec   ConfigRuleSpec   `json:"spec"`
	Status ConfigRuleStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// ConfigRuleList contains a list of ConfigRules
type ConfigRuleList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []ConfigRule `json:"items"`
}

// Repository type metadata.
var (
	ConfigRule_Kind             = "ConfigRule"
	ConfigRule_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: ConfigRule_Kind}.String()
	ConfigRule_KindAPIVersion   = ConfigRule_Kind + "." + CRDGroupVersion.String()
	ConfigRule_GroupVersionKind = CRDGroupVersion.WithKind(ConfigRule_Kind)
)

func init() {
	SchemeBuilder.Register(&ConfigRule{}, &ConfigRuleList{})
}
