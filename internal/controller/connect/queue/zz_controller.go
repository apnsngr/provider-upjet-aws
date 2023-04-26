/*
Copyright 2022 Upbound Inc.
*/

// Code generated by upjet. DO NOT EDIT.

package queue

import (
	"time"

	"github.com/crossplane/crossplane-runtime/pkg/connection"
	"github.com/crossplane/crossplane-runtime/pkg/event"
	"github.com/crossplane/crossplane-runtime/pkg/ratelimiter"
	"github.com/crossplane/crossplane-runtime/pkg/reconciler/managed"
	xpresource "github.com/crossplane/crossplane-runtime/pkg/resource"
	tjcontroller "github.com/upbound/upjet/pkg/controller"
	"github.com/upbound/upjet/pkg/terraform"
	ctrl "sigs.k8s.io/controller-runtime"

	v1beta1 "github.com/upbound/provider-aws/apis/connect/v1beta1"
	features "github.com/upbound/provider-aws/internal/features"
)

// Setup adds a controller that reconciles Queue managed resources.
func Setup(mgr ctrl.Manager, o tjcontroller.Options) error {
	name := managed.ControllerName(v1beta1.Queue_GroupVersionKind.String())
	var initializers managed.InitializerChain
	for _, i := range o.Provider.Resources["aws_connect_queue"].InitializerFns {
		initializers = append(initializers, i(mgr.GetClient()))
	}
	cps := []managed.ConnectionPublisher{managed.NewAPISecretPublisher(mgr.GetClient(), mgr.GetScheme())}
	if o.SecretStoreConfigGVK != nil {
		cps = append(cps, connection.NewDetailsManager(mgr.GetClient(), *o.SecretStoreConfigGVK, connection.WithTLSConfig(o.ESSOptions.TLSConfig)))
	}
	opts := []managed.ReconcilerOption{
		managed.WithExternalConnecter(tjcontroller.NewConnector(mgr.GetClient(), o.WorkspaceStore, o.SetupFn, o.Provider.Resources["aws_connect_queue"], tjcontroller.WithLogger(o.Logger),
			tjcontroller.WithCallbackProvider(tjcontroller.NewAPICallbacks(mgr, xpresource.ManagedKind(v1beta1.Queue_GroupVersionKind))),
		)),
		managed.WithLogger(o.Logger.WithValues("controller", name)),
		managed.WithRecorder(event.NewAPIRecorder(mgr.GetEventRecorderFor(name))),
		managed.WithFinalizer(terraform.NewWorkspaceFinalizer(o.WorkspaceStore, xpresource.NewAPIFinalizer(mgr.GetClient(), managed.FinalizerName))),
		managed.WithTimeout(3 * time.Minute),
		managed.WithInitializers(initializers),
		managed.WithConnectionPublishers(cps...),
		managed.WithPollInterval(o.PollInterval),
	}
	if o.Features.Enabled(features.EnableAlphaManagementPolicies) {
		opts = append(opts, managed.WithManagementPolicies())
	}
	r := managed.NewReconciler(mgr, xpresource.ManagedKind(v1beta1.Queue_GroupVersionKind), opts...)

	return ctrl.NewControllerManagedBy(mgr).
		Named(name).
		WithOptions(o.ForControllerRuntime()).
		For(&v1beta1.Queue{}).
		Complete(ratelimiter.NewReconciler(name, r, o.GlobalRateLimiter))
}
