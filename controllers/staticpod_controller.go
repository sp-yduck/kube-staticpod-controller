/*
Copyright 2022 Teppei Sudo.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package controllers

import (
	"context"

	"k8s.io/apimachinery/pkg/runtime"
	ctrl "sigs.k8s.io/controller-runtime"
	"sigs.k8s.io/controller-runtime/pkg/client"
	"sigs.k8s.io/controller-runtime/pkg/log"

	staticv1alpha1 "github.com/sp-yduck/kube-staticpod-controller/api/v1alpha1"
)

// StaticPodReconciler reconciles a StaticPod object
type StaticPodReconciler struct {
	client.Client
	Scheme *runtime.Scheme
}

//+kubebuilder:rbac:groups=sp-yduck.com,resources=staticpods,verbs=get;list;watch;create;update;patch;delete
//+kubebuilder:rbac:groups=sp-yduck.com,resources=staticpods/status,verbs=get;update;patch
//+kubebuilder:rbac:groups=sp-yduck.com,resources=staticpods/finalizers,verbs=update

// Reconcile is part of the main kubernetes reconciliation loop which aims to
// move the current state of the cluster closer to the desired state.
// TODO(user): Modify the Reconcile function to compare the state specified by
// the StaticPod object against the actual cluster state, and then
// perform operations to make the cluster state reflect the state specified by
// the user.
//
// For more details, check Reconcile and its Result here:
// - https://pkg.go.dev/sigs.k8s.io/controller-runtime@v0.12.1/pkg/reconcile
func (r *StaticPodReconciler) Reconcile(ctx context.Context, req ctrl.Request) (ctrl.Result, error) {
	logger := log.FromContext(ctx)

	var staticPod staticv1alpha1.StaticPod
	if err := r.Get(ctx, req.NamespacedName, &staticPod); err != nil {
		logger.Error(err, "unable ro fetch StaticPod")
		return ctrl.Result{}, client.IgnoreNotFound(err)
	}

	var childStaticPods staticv1alpha1.StaticPodList
	if err := r.List(ctx, &childStaticPods, client.InNamespace(req.Namespace), &client.ListOptions{}); err != nil {
		logger.Error(err, "unable to list StaticPods")
		return ctrl.Result{}, err
	}

	return ctrl.Result{}, nil
}

// SetupWithManager sets up the controller with the Manager.
func (r *StaticPodReconciler) SetupWithManager(mgr ctrl.Manager) error {
	return ctrl.NewControllerManagedBy(mgr).
		For(&staticv1alpha1.StaticPod{}).
		Complete(r)
}
