/*
Copyright 2023 KubeAGI.

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

package retriever

import (
	"context"
	"reflect"

	"github.com/go-logr/logr"
	"k8s.io/apimachinery/pkg/runtime"
	ctrl "sigs.k8s.io/controller-runtime"
	"sigs.k8s.io/controller-runtime/pkg/client"
	"sigs.k8s.io/controller-runtime/pkg/controller/controllerutil"

	api "github.com/kubeagi/arcadia/api/app-node/retriever/v1alpha1"
	arcadiav1alpha1 "github.com/kubeagi/arcadia/api/base/v1alpha1"
	appnode "github.com/kubeagi/arcadia/controllers/app-node"
)

// ConversationRetrieverReconciler reconciles a ConversationRetriever object
type ConversationRetrieverReconciler struct {
	client.Client
	Scheme *runtime.Scheme
}

//+kubebuilder:rbac:groups=retriever.arcadia.kubeagi.k8s.com.cn,resources=conversationretrievers,verbs=get;list;watch;create;update;patch;delete
//+kubebuilder:rbac:groups=retriever.arcadia.kubeagi.k8s.com.cn,resources=conversationretrievers/status,verbs=get;update;patch
//+kubebuilder:rbac:groups=retriever.arcadia.kubeagi.k8s.com.cn,resources=conversationretrievers/finalizers,verbs=update

// Reconcile is part of the main kubernetes reconciliation loop which aims to
// move the current state of the cluster closer to the desired state.
// For more details, check Reconcile and its Result here:
// - https://pkg.go.dev/sigs.k8s.io/controller-runtime@v0.12.2/pkg/reconcile
func (r *ConversationRetrieverReconciler) Reconcile(ctx context.Context, req ctrl.Request) (ctrl.Result, error) {
	log := ctrl.LoggerFrom(ctx)
	log.V(5).Info("Start ConversationRetriever Reconcile")
	instance := &api.ConversationRetriever{}
	if err := r.Get(ctx, req.NamespacedName, instance); err != nil {
		// There's no need to requeue if the resource no longer exists.
		// Otherwise, we'll be requeued implicitly because we return an error.
		log.V(1).Info("Failed to get ConversationRetriever")
		return ctrl.Result{}, client.IgnoreNotFound(err)
	}
	log = log.WithValues("Generation", instance.GetGeneration(), "ObservedGeneration", instance.Status.ObservedGeneration, "creator", instance.Spec.Creator)
	log.V(5).Info("Get ConversationRetriever instance")

	// Add a finalizer.Then, we can define some operations which should
	// occur before the ConversationRetriever to be deleted.
	// More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/finalizers
	if newAdded := controllerutil.AddFinalizer(instance, arcadiav1alpha1.Finalizer); newAdded {
		log.Info("Try to add Finalizer for ConversationRetriever")
		if err := r.Update(ctx, instance); err != nil {
			log.Error(err, "Failed to update ConversationRetriever to add finalizer, will try again later")
			return ctrl.Result{}, err
		}
		log.Info("Adding Finalizer for ConversationRetriever done")
		return ctrl.Result{}, nil
	}

	// Check if the ConversationRetriever instance is marked to be deleted, which is
	// indicated by the deletion timestamp being set.
	if instance.GetDeletionTimestamp() != nil && controllerutil.ContainsFinalizer(instance, arcadiav1alpha1.Finalizer) {
		log.Info("Performing Finalizer Operations for ConversationRetriever before delete CR")
		// TODO perform the finalizer operations here, for example: remove vectorstore data?
		log.Info("Removing Finalizer for ConversationRetriever after successfully performing the operations")
		controllerutil.RemoveFinalizer(instance, arcadiav1alpha1.Finalizer)
		if err := r.Update(ctx, instance); err != nil {
			log.Error(err, "Failed to remove the finalizer for ConversationRetriever")
			return ctrl.Result{}, err
		}
		log.Info("Remove ConversationRetriever done")
		return ctrl.Result{}, nil
	}

	instance, result, err := r.reconcile(ctx, log, instance)

	// Update status after reconciliation.
	if updateStatusErr := r.patchStatus(ctx, instance); updateStatusErr != nil {
		log.Error(updateStatusErr, "unable to update status after reconciliation")
		return ctrl.Result{Requeue: true}, updateStatusErr
	}

	return result, err
}

func (r *ConversationRetrieverReconciler) reconcile(ctx context.Context, log logr.Logger, instance *api.ConversationRetriever) (*api.ConversationRetriever, ctrl.Result, error) {
	// Observe generation change
	if instance.Status.ObservedGeneration != instance.Generation {
		instance.Status.ObservedGeneration = instance.Generation
		r.setCondition(instance, instance.Status.WaitingCompleteCondition()...)
		if updateStatusErr := r.patchStatus(ctx, instance); updateStatusErr != nil {
			log.Error(updateStatusErr, "unable to update status after generation update")
			return instance, ctrl.Result{Requeue: true}, updateStatusErr
		}
	}

	if instance.Status.IsReady() {
		return instance, ctrl.Result{}, nil
	}

	if err := appnode.CheckAndUpdateAnnotation(ctx, log, r.Client, instance); err != nil {
		instance.Status.SetConditions(instance.Status.ErrorCondition(err.Error())...)
	} else {
		instance.Status.SetConditions(instance.Status.ReadyCondition()...)
	}

	// TODO:Check LLM,Embedder,VectorStore

	return instance, ctrl.Result{}, nil
}

func (r *ConversationRetrieverReconciler) patchStatus(ctx context.Context, instance *api.ConversationRetriever) error {
	latest := &api.ConversationRetriever{}
	if err := r.Client.Get(ctx, client.ObjectKeyFromObject(instance), latest); err != nil {
		return err
	}
	if reflect.DeepEqual(instance.Status, latest.Status) {
		return nil
	}
	patch := client.MergeFrom(latest.DeepCopy())
	latest.Status = instance.Status
	return r.Client.Status().Patch(ctx, latest, patch, client.FieldOwner("ConversationRetriever-controller"))
}

// SetupWithManager sets up the controller with the Manager.
func (r *ConversationRetrieverReconciler) SetupWithManager(mgr ctrl.Manager) error {
	return ctrl.NewControllerManagedBy(mgr).
		For(&api.ConversationRetriever{}).
		Complete(r)
}

func (r *ConversationRetrieverReconciler) setCondition(instance *api.ConversationRetriever, condition ...arcadiav1alpha1.Condition) *api.ConversationRetriever {
	instance.Status.SetConditions(condition...)
	return instance
}
