/*
Copyright 2022 Upbound Inc.
*/
// Code generated by angryjet. DO NOT EDIT.

package v1beta1

import (
	"context"
	reference "github.com/crossplane/crossplane-runtime/pkg/reference"
	resource "github.com/crossplane/upjet/pkg/resource"
	errors "github.com/pkg/errors"
	v1beta1 "github.com/upbound/provider-aws/apis/acm/v1beta1"
	v1beta11 "github.com/upbound/provider-aws/apis/ds/v1beta1"
	v1beta12 "github.com/upbound/provider-aws/apis/ec2/v1beta1"
	v1beta13 "github.com/upbound/provider-aws/apis/iam/v1beta1"
	v1beta14 "github.com/upbound/provider-aws/apis/lambda/v1beta1"
	common "github.com/upbound/provider-aws/config/common"
	client "sigs.k8s.io/controller-runtime/pkg/client"
)

// ResolveReferences of this SSHKey.
func (mg *SSHKey) ResolveReferences(ctx context.Context, c client.Reader) error {
	r := reference.NewAPIResolver(c, mg)

	var rsp reference.ResolutionResponse
	var err error

	rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
		CurrentValue: reference.FromPtrValue(mg.Spec.ForProvider.ServerID),
		Extract:      resource.ExtractResourceID(),
		Reference:    mg.Spec.ForProvider.ServerIDRef,
		Selector:     mg.Spec.ForProvider.ServerIDSelector,
		To: reference.To{
			List:    &ServerList{},
			Managed: &Server{},
		},
	})
	if err != nil {
		return errors.Wrap(err, "mg.Spec.ForProvider.ServerID")
	}
	mg.Spec.ForProvider.ServerID = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.ForProvider.ServerIDRef = rsp.ResolvedReference

	rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
		CurrentValue: reference.FromPtrValue(mg.Spec.ForProvider.UserName),
		Extract:      reference.ExternalName(),
		Reference:    mg.Spec.ForProvider.UserNameRef,
		Selector:     mg.Spec.ForProvider.UserNameSelector,
		To: reference.To{
			List:    &UserList{},
			Managed: &User{},
		},
	})
	if err != nil {
		return errors.Wrap(err, "mg.Spec.ForProvider.UserName")
	}
	mg.Spec.ForProvider.UserName = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.ForProvider.UserNameRef = rsp.ResolvedReference

	rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
		CurrentValue: reference.FromPtrValue(mg.Spec.InitProvider.ServerID),
		Extract:      resource.ExtractResourceID(),
		Reference:    mg.Spec.InitProvider.ServerIDRef,
		Selector:     mg.Spec.InitProvider.ServerIDSelector,
		To: reference.To{
			List:    &ServerList{},
			Managed: &Server{},
		},
	})
	if err != nil {
		return errors.Wrap(err, "mg.Spec.InitProvider.ServerID")
	}
	mg.Spec.InitProvider.ServerID = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.InitProvider.ServerIDRef = rsp.ResolvedReference

	rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
		CurrentValue: reference.FromPtrValue(mg.Spec.InitProvider.UserName),
		Extract:      reference.ExternalName(),
		Reference:    mg.Spec.InitProvider.UserNameRef,
		Selector:     mg.Spec.InitProvider.UserNameSelector,
		To: reference.To{
			List:    &UserList{},
			Managed: &User{},
		},
	})
	if err != nil {
		return errors.Wrap(err, "mg.Spec.InitProvider.UserName")
	}
	mg.Spec.InitProvider.UserName = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.InitProvider.UserNameRef = rsp.ResolvedReference

	return nil
}

// ResolveReferences of this Server.
func (mg *Server) ResolveReferences(ctx context.Context, c client.Reader) error {
	r := reference.NewAPIResolver(c, mg)

	var rsp reference.ResolutionResponse
	var err error

	rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
		CurrentValue: reference.FromPtrValue(mg.Spec.ForProvider.Certificate),
		Extract:      resource.ExtractParamPath("arn", true),
		Reference:    mg.Spec.ForProvider.CertificateRef,
		Selector:     mg.Spec.ForProvider.CertificateSelector,
		To: reference.To{
			List:    &v1beta1.CertificateList{},
			Managed: &v1beta1.Certificate{},
		},
	})
	if err != nil {
		return errors.Wrap(err, "mg.Spec.ForProvider.Certificate")
	}
	mg.Spec.ForProvider.Certificate = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.ForProvider.CertificateRef = rsp.ResolvedReference

	rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
		CurrentValue: reference.FromPtrValue(mg.Spec.ForProvider.DirectoryID),
		Extract:      resource.ExtractResourceID(),
		Reference:    mg.Spec.ForProvider.DirectoryIDRef,
		Selector:     mg.Spec.ForProvider.DirectoryIDSelector,
		To: reference.To{
			List:    &v1beta11.DirectoryList{},
			Managed: &v1beta11.Directory{},
		},
	})
	if err != nil {
		return errors.Wrap(err, "mg.Spec.ForProvider.DirectoryID")
	}
	mg.Spec.ForProvider.DirectoryID = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.ForProvider.DirectoryIDRef = rsp.ResolvedReference

	for i3 := 0; i3 < len(mg.Spec.ForProvider.EndpointDetails); i3++ {
		rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
			CurrentValue: reference.FromPtrValue(mg.Spec.ForProvider.EndpointDetails[i3].VPCID),
			Extract:      resource.ExtractResourceID(),
			Reference:    mg.Spec.ForProvider.EndpointDetails[i3].VPCIDRef,
			Selector:     mg.Spec.ForProvider.EndpointDetails[i3].VPCIDSelector,
			To: reference.To{
				List:    &v1beta12.VPCList{},
				Managed: &v1beta12.VPC{},
			},
		})
		if err != nil {
			return errors.Wrap(err, "mg.Spec.ForProvider.EndpointDetails[i3].VPCID")
		}
		mg.Spec.ForProvider.EndpointDetails[i3].VPCID = reference.ToPtrValue(rsp.ResolvedValue)
		mg.Spec.ForProvider.EndpointDetails[i3].VPCIDRef = rsp.ResolvedReference

	}
	rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
		CurrentValue: reference.FromPtrValue(mg.Spec.ForProvider.LoggingRole),
		Extract:      resource.ExtractParamPath("arn", true),
		Reference:    mg.Spec.ForProvider.LoggingRoleRef,
		Selector:     mg.Spec.ForProvider.LoggingRoleSelector,
		To: reference.To{
			List:    &v1beta13.RoleList{},
			Managed: &v1beta13.Role{},
		},
	})
	if err != nil {
		return errors.Wrap(err, "mg.Spec.ForProvider.LoggingRole")
	}
	mg.Spec.ForProvider.LoggingRole = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.ForProvider.LoggingRoleRef = rsp.ResolvedReference

	rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
		CurrentValue: reference.FromPtrValue(mg.Spec.InitProvider.Certificate),
		Extract:      resource.ExtractParamPath("arn", true),
		Reference:    mg.Spec.InitProvider.CertificateRef,
		Selector:     mg.Spec.InitProvider.CertificateSelector,
		To: reference.To{
			List:    &v1beta1.CertificateList{},
			Managed: &v1beta1.Certificate{},
		},
	})
	if err != nil {
		return errors.Wrap(err, "mg.Spec.InitProvider.Certificate")
	}
	mg.Spec.InitProvider.Certificate = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.InitProvider.CertificateRef = rsp.ResolvedReference

	rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
		CurrentValue: reference.FromPtrValue(mg.Spec.InitProvider.DirectoryID),
		Extract:      resource.ExtractResourceID(),
		Reference:    mg.Spec.InitProvider.DirectoryIDRef,
		Selector:     mg.Spec.InitProvider.DirectoryIDSelector,
		To: reference.To{
			List:    &v1beta11.DirectoryList{},
			Managed: &v1beta11.Directory{},
		},
	})
	if err != nil {
		return errors.Wrap(err, "mg.Spec.InitProvider.DirectoryID")
	}
	mg.Spec.InitProvider.DirectoryID = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.InitProvider.DirectoryIDRef = rsp.ResolvedReference

	for i3 := 0; i3 < len(mg.Spec.InitProvider.EndpointDetails); i3++ {
		rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
			CurrentValue: reference.FromPtrValue(mg.Spec.InitProvider.EndpointDetails[i3].VPCID),
			Extract:      resource.ExtractResourceID(),
			Reference:    mg.Spec.InitProvider.EndpointDetails[i3].VPCIDRef,
			Selector:     mg.Spec.InitProvider.EndpointDetails[i3].VPCIDSelector,
			To: reference.To{
				List:    &v1beta12.VPCList{},
				Managed: &v1beta12.VPC{},
			},
		})
		if err != nil {
			return errors.Wrap(err, "mg.Spec.InitProvider.EndpointDetails[i3].VPCID")
		}
		mg.Spec.InitProvider.EndpointDetails[i3].VPCID = reference.ToPtrValue(rsp.ResolvedValue)
		mg.Spec.InitProvider.EndpointDetails[i3].VPCIDRef = rsp.ResolvedReference

	}
	rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
		CurrentValue: reference.FromPtrValue(mg.Spec.InitProvider.LoggingRole),
		Extract:      resource.ExtractParamPath("arn", true),
		Reference:    mg.Spec.InitProvider.LoggingRoleRef,
		Selector:     mg.Spec.InitProvider.LoggingRoleSelector,
		To: reference.To{
			List:    &v1beta13.RoleList{},
			Managed: &v1beta13.Role{},
		},
	})
	if err != nil {
		return errors.Wrap(err, "mg.Spec.InitProvider.LoggingRole")
	}
	mg.Spec.InitProvider.LoggingRole = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.InitProvider.LoggingRoleRef = rsp.ResolvedReference

	rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
		CurrentValue: reference.FromPtrValue(mg.Spec.InitProvider.Certificate),
		Extract:      resource.ExtractParamPath("arn", true),
		Reference:    mg.Spec.InitProvider.CertificateRef,
		Selector:     mg.Spec.InitProvider.CertificateSelector,
		To: reference.To{
			List:    &v1beta1.CertificateList{},
			Managed: &v1beta1.Certificate{},
		},
	})
	if err != nil {
		return errors.Wrap(err, "mg.Spec.InitProvider.Certificate")
	}
	mg.Spec.InitProvider.Certificate = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.InitProvider.CertificateRef = rsp.ResolvedReference

	rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
		CurrentValue: reference.FromPtrValue(mg.Spec.InitProvider.DirectoryID),
		Extract:      resource.ExtractResourceID(),
		Reference:    mg.Spec.InitProvider.DirectoryIDRef,
		Selector:     mg.Spec.InitProvider.DirectoryIDSelector,
		To: reference.To{
			List:    &v1beta11.DirectoryList{},
			Managed: &v1beta11.Directory{},
		},
	})
	if err != nil {
		return errors.Wrap(err, "mg.Spec.InitProvider.DirectoryID")
	}
	mg.Spec.InitProvider.DirectoryID = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.InitProvider.DirectoryIDRef = rsp.ResolvedReference

	for i3 := 0; i3 < len(mg.Spec.InitProvider.EndpointDetails); i3++ {
		rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
			CurrentValue: reference.FromPtrValue(mg.Spec.InitProvider.EndpointDetails[i3].VPCID),
			Extract:      resource.ExtractResourceID(),
			Reference:    mg.Spec.InitProvider.EndpointDetails[i3].VPCIDRef,
			Selector:     mg.Spec.InitProvider.EndpointDetails[i3].VPCIDSelector,
			To: reference.To{
				List:    &v1beta12.VPCList{},
				Managed: &v1beta12.VPC{},
			},
		})
		if err != nil {
			return errors.Wrap(err, "mg.Spec.InitProvider.EndpointDetails[i3].VPCID")
		}
		mg.Spec.InitProvider.EndpointDetails[i3].VPCID = reference.ToPtrValue(rsp.ResolvedValue)
		mg.Spec.InitProvider.EndpointDetails[i3].VPCIDRef = rsp.ResolvedReference

	}
	rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
		CurrentValue: reference.FromPtrValue(mg.Spec.InitProvider.LoggingRole),
		Extract:      resource.ExtractParamPath("arn", true),
		Reference:    mg.Spec.InitProvider.LoggingRoleRef,
		Selector:     mg.Spec.InitProvider.LoggingRoleSelector,
		To: reference.To{
			List:    &v1beta13.RoleList{},
			Managed: &v1beta13.Role{},
		},
	})
	if err != nil {
		return errors.Wrap(err, "mg.Spec.InitProvider.LoggingRole")
	}
	mg.Spec.InitProvider.LoggingRole = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.InitProvider.LoggingRoleRef = rsp.ResolvedReference

	return nil
}

// ResolveReferences of this Tag.
func (mg *Tag) ResolveReferences(ctx context.Context, c client.Reader) error {
	r := reference.NewAPIResolver(c, mg)

	var rsp reference.ResolutionResponse
	var err error

	rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
		CurrentValue: reference.FromPtrValue(mg.Spec.ForProvider.ResourceArn),
		Extract:      resource.ExtractParamPath("arn", true),
		Reference:    mg.Spec.ForProvider.ResourceArnRef,
		Selector:     mg.Spec.ForProvider.ResourceArnSelector,
		To: reference.To{
			List:    &ServerList{},
			Managed: &Server{},
		},
	})
	if err != nil {
		return errors.Wrap(err, "mg.Spec.ForProvider.ResourceArn")
	}
	mg.Spec.ForProvider.ResourceArn = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.ForProvider.ResourceArnRef = rsp.ResolvedReference

	rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
		CurrentValue: reference.FromPtrValue(mg.Spec.InitProvider.ResourceArn),
		Extract:      resource.ExtractParamPath("arn", true),
		Reference:    mg.Spec.InitProvider.ResourceArnRef,
		Selector:     mg.Spec.InitProvider.ResourceArnSelector,
		To: reference.To{
			List:    &ServerList{},
			Managed: &Server{},
		},
	})
	if err != nil {
		return errors.Wrap(err, "mg.Spec.InitProvider.ResourceArn")
	}
	mg.Spec.InitProvider.ResourceArn = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.InitProvider.ResourceArnRef = rsp.ResolvedReference

	return nil
}

// ResolveReferences of this User.
func (mg *User) ResolveReferences(ctx context.Context, c client.Reader) error {
	r := reference.NewAPIResolver(c, mg)

	var rsp reference.ResolutionResponse
	var err error

	rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
		CurrentValue: reference.FromPtrValue(mg.Spec.ForProvider.Role),
		Extract:      common.ARNExtractor(),
		Reference:    mg.Spec.ForProvider.RoleRef,
		Selector:     mg.Spec.ForProvider.RoleSelector,
		To: reference.To{
			List:    &v1beta13.RoleList{},
			Managed: &v1beta13.Role{},
		},
	})
	if err != nil {
		return errors.Wrap(err, "mg.Spec.ForProvider.Role")
	}
	mg.Spec.ForProvider.Role = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.ForProvider.RoleRef = rsp.ResolvedReference

	rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
		CurrentValue: reference.FromPtrValue(mg.Spec.ForProvider.ServerID),
		Extract:      reference.ExternalName(),
		Reference:    mg.Spec.ForProvider.ServerIDRef,
		Selector:     mg.Spec.ForProvider.ServerIDSelector,
		To: reference.To{
			List:    &ServerList{},
			Managed: &Server{},
		},
	})
	if err != nil {
		return errors.Wrap(err, "mg.Spec.ForProvider.ServerID")
	}
	mg.Spec.ForProvider.ServerID = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.ForProvider.ServerIDRef = rsp.ResolvedReference

	rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
		CurrentValue: reference.FromPtrValue(mg.Spec.InitProvider.Role),
		Extract:      common.ARNExtractor(),
		Reference:    mg.Spec.InitProvider.RoleRef,
		Selector:     mg.Spec.InitProvider.RoleSelector,
		To: reference.To{
			List:    &v1beta13.RoleList{},
			Managed: &v1beta13.Role{},
		},
	})
	if err != nil {
		return errors.Wrap(err, "mg.Spec.InitProvider.Role")
	}
	mg.Spec.InitProvider.Role = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.InitProvider.RoleRef = rsp.ResolvedReference

	rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
		CurrentValue: reference.FromPtrValue(mg.Spec.InitProvider.ServerID),
		Extract:      reference.ExternalName(),
		Reference:    mg.Spec.InitProvider.ServerIDRef,
		Selector:     mg.Spec.InitProvider.ServerIDSelector,
		To: reference.To{
			List:    &ServerList{},
			Managed: &Server{},
		},
	})
	if err != nil {
		return errors.Wrap(err, "mg.Spec.InitProvider.ServerID")
	}
	mg.Spec.InitProvider.ServerID = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.InitProvider.ServerIDRef = rsp.ResolvedReference

	return nil
}

// ResolveReferences of this Workflow.
func (mg *Workflow) ResolveReferences(ctx context.Context, c client.Reader) error {
	r := reference.NewAPIResolver(c, mg)

	var rsp reference.ResolutionResponse
	var err error

	for i3 := 0; i3 < len(mg.Spec.ForProvider.Steps); i3++ {
		for i4 := 0; i4 < len(mg.Spec.ForProvider.Steps[i3].CustomStepDetails); i4++ {
			rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
				CurrentValue: reference.FromPtrValue(mg.Spec.ForProvider.Steps[i3].CustomStepDetails[i4].Target),
				Extract:      resource.ExtractParamPath("arn", true),
				Reference:    mg.Spec.ForProvider.Steps[i3].CustomStepDetails[i4].TargetRef,
				Selector:     mg.Spec.ForProvider.Steps[i3].CustomStepDetails[i4].TargetSelector,
				To: reference.To{
					List:    &v1beta14.FunctionList{},
					Managed: &v1beta14.Function{},
				},
			})
			if err != nil {
				return errors.Wrap(err, "mg.Spec.ForProvider.Steps[i3].CustomStepDetails[i4].Target")
			}
			mg.Spec.ForProvider.Steps[i3].CustomStepDetails[i4].Target = reference.ToPtrValue(rsp.ResolvedValue)
			mg.Spec.ForProvider.Steps[i3].CustomStepDetails[i4].TargetRef = rsp.ResolvedReference

		}
	}
	for i3 := 0; i3 < len(mg.Spec.InitProvider.Steps); i3++ {
		for i4 := 0; i4 < len(mg.Spec.InitProvider.Steps[i3].CustomStepDetails); i4++ {
			rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
				CurrentValue: reference.FromPtrValue(mg.Spec.InitProvider.Steps[i3].CustomStepDetails[i4].Target),
				Extract:      resource.ExtractParamPath("arn", true),
				Reference:    mg.Spec.InitProvider.Steps[i3].CustomStepDetails[i4].TargetRef,
				Selector:     mg.Spec.InitProvider.Steps[i3].CustomStepDetails[i4].TargetSelector,
				To: reference.To{
					List:    &v1beta14.FunctionList{},
					Managed: &v1beta14.Function{},
				},
			})
			if err != nil {
				return errors.Wrap(err, "mg.Spec.InitProvider.Steps[i3].CustomStepDetails[i4].Target")
			}
			mg.Spec.InitProvider.Steps[i3].CustomStepDetails[i4].Target = reference.ToPtrValue(rsp.ResolvedValue)
			mg.Spec.InitProvider.Steps[i3].CustomStepDetails[i4].TargetRef = rsp.ResolvedReference

		}
	}

	return nil
}
