/*


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

package controller

import (
	"bytes"
	"context"
	"fmt"
	"hash/fnv"
	"strings"

	performancev2 "github.com/openshift/cluster-node-tuning-operator/pkg/apis/performanceprofile/v2"
	tunedv1 "github.com/openshift/cluster-node-tuning-operator/pkg/apis/tuned/v1"
	performanceprofilecomponents "github.com/openshift/cluster-node-tuning-operator/pkg/performanceprofile/controller/performanceprofile/components"
	"github.com/openshift/cluster-node-tuning-operator/pkg/performanceprofile/controller/performanceprofile/components/manifestset"
	"github.com/openshift/cluster-node-tuning-operator/pkg/performanceprofile/controller/performanceprofile/components/runtimeclass"
	mcfgv1 "github.com/openshift/hypershift/thirdparty/machineconfigoperator/pkg/apis/machineconfiguration.openshift.io/v1"
	mcov1 "github.com/openshift/machine-config-operator/pkg/apis/machineconfiguration.openshift.io/v1"

	corev1 "k8s.io/api/core/v1"
	nodev1 "k8s.io/api/node/v1"
	apiequality "k8s.io/apimachinery/pkg/api/equality"
	k8serrors "k8s.io/apimachinery/pkg/api/errors"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/apimachinery/pkg/runtime/schema"
	serializer "k8s.io/apimachinery/pkg/runtime/serializer/json"
	"k8s.io/apimachinery/pkg/types"
	"k8s.io/klog/v2"

	ctrl "sigs.k8s.io/controller-runtime"
	"sigs.k8s.io/controller-runtime/pkg/builder"
	"sigs.k8s.io/controller-runtime/pkg/client"
	"sigs.k8s.io/controller-runtime/pkg/cluster"
	"sigs.k8s.io/controller-runtime/pkg/event"
	"sigs.k8s.io/controller-runtime/pkg/handler"
	"sigs.k8s.io/controller-runtime/pkg/predicate"
	"sigs.k8s.io/controller-runtime/pkg/reconcile"
	"sigs.k8s.io/controller-runtime/pkg/source"
)

const (
	hypershiftPerformanceProfileNameLabel = "hypershift.openshift.io/performanceProfileName"
	hypershiftNodePoolNameLabel           = "hypershift.openshift.io/nodePoolName"
	hypershiftNodePoolLabel               = "hypershift.openshift.io/nodePool"
	controllerGeneratedMachineConfig      = "hypershift.openshift.io/performanceprofile-config"

	tunedConfigMapLabel     = "hypershift.openshift.io/tuned-config"
	tunedConfigMapConfigKey = "tuning"

	mcoConfigMapConfigKey          = "config"
	ntoGeneratedMachineConfigLabel = "hypershift.openshift.io/nto-generated-machine-config"

	hypershiftFinalizer = "hypershift.openshift.io/foreground-deletion"
)

func configureScheme(scheme *runtime.Scheme) error {
	if err := tunedv1.AddToScheme(scheme); err != nil {
		return fmt.Errorf("unable to add tuned/v1 to scheme. %w", err)
	}
	if err := mcfgv1.AddToScheme(scheme); err != nil {
		return fmt.Errorf("unable to add machineconfiguration.openshift.io/v1 to scheme. %w", err)
	}
	if err := performancev2.AddToScheme(scheme); err != nil {
		return fmt.Errorf("unable to add performanceprofile/v2 to scheme. %w", err)
	}
	return nil
}

func (r *PerformanceProfileReconciler) HypershiftSetupWithManager(mgr ctrl.Manager, managementCluster cluster.Cluster) error {
	// In hypershift just have to reconcile ConfigMaps created by Hypershift Operator in the
	// controller namespace with the right label.
	p := predicate.Funcs{
		UpdateFunc: func(ue event.UpdateEvent) bool {
			if !validateUpdateEvent(&ue) {
				klog.Infof("[%s] UpdateEvent NOT VALID", ue.ObjectOld.GetName())
				return false
			}

			_, hasLabel := ue.ObjectNew.GetLabels()[controllerGeneratedMachineConfig]
			if hasLabel {
				klog.Infof("[%s] UpdateEvent has label %s", ue.ObjectOld.GetName(), controllerGeneratedMachineConfig)
			}
			return hasLabel
		},
		CreateFunc: func(ce event.CreateEvent) bool {
			if ce.Object == nil {
				klog.Error("Create event has no runtime object")
				return false
			}

			_, hasLabel := ce.Object.GetLabels()[controllerGeneratedMachineConfig]
			if hasLabel {
				klog.Infof("[%s] CreateEvent has label", ce.Object.GetName(), controllerGeneratedMachineConfig)
			}
			return hasLabel
		},
		DeleteFunc: func(de event.DeleteEvent) bool {
			if de.Object == nil {
				klog.Error("Delete event has no runtime object")
				return false
			}
			_, hasLabel := de.Object.GetLabels()[controllerGeneratedMachineConfig]
			if hasLabel {
				klog.Infof("[%s] DeleteEvent has label", de.Object.GetName(), controllerGeneratedMachineConfig)
			}
			return hasLabel
		},
	}

	if err := configureScheme(r.Scheme); err != nil {
		klog.Errorf("unable to configure scheme %v", err)
		return err
	}

	return ctrl.NewControllerManagedBy(mgr).
		Named("performanceprofile_controller").
		WatchesRawSource(source.Kind(managementCluster.GetCache(), &corev1.ConfigMap{}),
			&handler.EnqueueRequestForObject{},
			builder.WithPredicates(p)).Complete(r)
}

func (r *PerformanceProfileReconciler) hypershiftReconcile(ctx context.Context, req ctrl.Request) (ctrl.Result, error) {
	klog.Infof("[%s] *** Entering ReconcileLoop ***", req.NamespacedName)
	defer klog.Infof("[%s] *** Exiting ReconcileLoop ***", req.NamespacedName)

	instance := &corev1.ConfigMap{}

	err := r.ManagementClient.Get(ctx, req.NamespacedName, instance)
	if err != nil {
		if k8serrors.IsNotFound(err) {
			// Request object not found, could have been deleted after reconcile request.
			// Owned objects are automatically garbage collected. For additional cleanup logic use finalizers.
			// Return and don't requeue
			klog.Infof("[%s] Error: Instance not found", req.NamespacedName)
			return reconcile.Result{}, nil
		}
		// Error reading the object - requeue the request.
		klog.Infof("[%s] Error: Reading failure. err: %v", req.NamespacedName, err)
		return reconcile.Result{}, err
	}

	if instance.DeletionTimestamp != nil {
		klog.Infof("[%s] deletion timestamp NOT NULL ", instance.Name)
		// ConfigMap is marked for deletion and waiting for finalizers to be empty
		// so better to clean-up and delete the objects.
		if err := hypershiftDeleteComponents(r.Client, ctx, instance); err != nil {
			klog.Errorf("[%s] Error: failed to delete components. err: %v", req.NamespacedName, err)
			r.Recorder.Eventf(instance, corev1.EventTypeWarning, "Deletion failed", "[hypershift:%s] Failed to delete components: %v", req.NamespacedName, err)
			return reconcile.Result{}, err
		}
		r.Recorder.Eventf(instance, corev1.EventTypeNormal, "Deletion succeeded", "[hypershift: %s] Succeeded to delete all components", req.NamespacedName)

		// remove finalizer
		if configMapHasFinalizer(instance, hypershiftFinalizer) {
			klog.Infof("[%s] Has finalizer. Removing ... ", instance.Name)
			cm := configMapRemoveFinalizer(instance, hypershiftFinalizer)
			if err := r.ManagementClient.Update(ctx, cm); err != nil {
				klog.Errorf("[%s] Error: while trying to update configmap.err: %v", instance.Name, err)
				return reconcile.Result{}, err
			}
			klog.Infof("[%s] Configmap updated, finalizer deleted", instance.Name)
			return reconcile.Result{}, nil
		}
	}

	//add finalizer
	if !configMapHasFinalizer(instance, hypershiftFinalizer) {
		klog.Infof("[%s] Do NOT has finalizer. Adding ... ", instance.Name)
		instance.Finalizers = append(instance.Finalizers, hypershiftFinalizer)
		if err := r.ManagementClient.Update(ctx, instance); err != nil {
			klog.Errorf("[%s] Error: while trying to update configmap.err: %v", instance.Name, err)
			return reconcile.Result{}, err
		}
		klog.Infof("[%s] Configmap updated, finalizer added", instance.Name)
		return reconcile.Result{}, nil
	}

	performanceProfileString, ok := instance.Data[tunedConfigMapConfigKey]
	if !ok {
		klog.Errorf("[%s] Error: ConfigMap has no PerformanceProfile info inside (data entry for %s)", instance.Name, tunedConfigMapConfigKey)
		return reconcile.Result{}, fmt.Errorf("configmap %q has no PerformanceProfile info inside (no entry %q)", instance.Name, tunedConfigMapConfigKey)
	}

	cmNodePoolNamespacedName, ok := instance.Annotations[hypershiftNodePoolLabel]
	if !ok {
		klog.Errorf("[%s] Error: ConfigMap has no Annotation %s (NodePool namespacedname)", instance.Name, hypershiftNodePoolLabel)
		// Return and don't requeue
		return reconcile.Result{}, nil
	}
	nodePoolName := parseNamespacedName(cmNodePoolNamespacedName)

	performanceProfileFromConfigMap, err := parsePerformanceProfileManifest([]byte(performanceProfileString), nodePoolName, r.Scheme)
	if err != nil {
		klog.Errorf("[%s] Error: failed to parse PerformanceProfile manifest from Configmap data data %v", instance.Name, err)
		// Return and don't requeue
		return reconcile.Result{}, fmt.Errorf("failed to parse PerformanceProfile manifest from Configmap %q data: %w", instance.Name, err)
	}

	updatePerformanceProfileName(performanceProfileFromConfigMap, nodePoolName)
	klog.Infof("[%s] PerformanceProfile name updated to %s", instance.Name, performanceProfileFromConfigMap.Name)

	pinningMode, err := r.getInfraPartitioningMode()
	if err != nil {
		return ctrl.Result{}, err
	}

	//ContainerRuntimeConfiguration is not yet supported in Hypershift
	// see : https://github.com/openshift/hypershift/blob/1586b54b9ea0b60ecdd5e4d4fb57f99da51b581d/hypershift-operator/controllers/nodepool/nodepool_controller.go#L1962
	// so we gonna go with the default Runtime for now
	//TODO - Get ContainerRuntimeConfiguration when supported by Hypershift
	//NOTE - ContainerRuntimeConfiguration is intended to be listed in NodePools `spec.config` field as a ConfigMap reference
	//       so the way to read it could be difficult unless we could use some other way
	var ctrRuntime mcov1.ContainerRuntimeDefaultRuntime = mcov1.ContainerRuntimeDefaultRuntimeDefault
	klog.Infof("[%s]=> hypershift ContainerRuntimeConfig is not supported for Hypershift yet. Using default %q", instance.Name, ctrRuntime)

	componentSet, err := manifestset.GetNewComponents(performanceProfileFromConfigMap,
		&performanceprofilecomponents.Options{
			ProfileMCP: nil,
			MachineConfig: performanceprofilecomponents.MachineConfigOptions{
				PinningMode:    &pinningMode,
				DefaultRuntime: ctrRuntime},
		})

	if err != nil {
		klog.Errorf("[%s] Error: While getting componentset, err:%v", instance.Name, err)
		r.Recorder.Eventf(instance, corev1.EventTypeWarning, "Creation failed", "[hypershift:%s] Failed to create all components: %v", instance.Name, err)
		return reconcile.Result{}, fmt.Errorf("unable to get components from PerformanceProfile in ConfigMap %q. %w", instance.Name, err)
	}

	// Now we have to create a ConfigMap for each of the elements in the componentSet and then handle them to
	// the different agents that would made them effective in the managed cluster.( where workers are)
	tunedEncoded, err := encodeManifest(componentSet.Tuned, r.Scheme)
	if err != nil {
		klog.Errorf("[%s] Error: failed to encode TuneD, err:%v", instance.Name, err)
		return reconcile.Result{}, fmt.Errorf("failed to encode TuneD from %q. %w", instance.Name, err)
	}

	tunedConfigMap := TunedConfigMap(instance, performanceProfileFromConfigMap.Name, cmNodePoolNamespacedName, string(tunedEncoded))
	if err := createOrUpdateTunedConfigMap(tunedConfigMap, ctx, r.ManagementClient); err != nil {
		klog.Errorf("[%s] Error: While creating/updating tuned configmap (%s), err:%v", instance.Name, err, tunedConfigMap.Name)
		return reconcile.Result{}, fmt.Errorf("unable to create/update tuned configmap %q from %q. %w", tunedConfigMap.Name, instance.Name, err)
	}

	machineconfigEncoded, err := encodeManifest(convertMachineConfig(componentSet.MachineConfig), r.Scheme)
	if err != nil {
		klog.Errorf("[%s] Error: While encoding MachineConfig, err:%v", instance.Name, err)
		return reconcile.Result{}, fmt.Errorf("unable to encode MachineConfig from %q. %w", instance.Name, err)
	}

	machineconfigConfigMap := MachineConfigConfigMap(instance, performanceProfileFromConfigMap.Name, cmNodePoolNamespacedName, string(machineconfigEncoded))
	if err := createOrUpdateMachineConfigConfigMap(machineconfigConfigMap, ctx, r.ManagementClient); err != nil {
		klog.Errorf("[%s] Error: While creating/updating machineconfig configmap (%s), err:%v", instance.Name, machineconfigConfigMap.Name, err)
		return reconcile.Result{}, fmt.Errorf("unable to create/update machineconfig configmap %q from %q. %w", machineconfigConfigMap.Name, instance.Name, err)
	}

	kubeletconfigEncoded, err := encodeManifest(convertKubeletConfig(componentSet.KubeletConfig), r.Scheme)
	if err != nil {
		klog.Errorf("[%s] Error: While encoding kubeletconfig, err:%v", instance.Name, err)
		return reconcile.Result{}, fmt.Errorf("unable to encode kubeletconfig from %q. %w", instance.Name, err)
	}

	kubeletconfigConfigMap := KubeletConfigConfigMap(instance, performanceProfileFromConfigMap.Name, cmNodePoolNamespacedName, string(kubeletconfigEncoded))
	if err := createOrUpdateKubeletConfigConfigConfigMap(kubeletconfigConfigMap, ctx, r.ManagementClient); err != nil {
		klog.Errorf("[%s] Error: While creating/updating kubeletconfig configmap (%s), err:%v", instance.Name, kubeletconfigConfigMap.Name, err)
		return reconcile.Result{}, fmt.Errorf("unable to create/update kubeletconfig configmap %q from %q. %w", kubeletconfigConfigMap.Name, instance.Name, err)
	}

	componentSet.RuntimeClass.Name = runtimeclass.BuildRuntimeClassName(instance.Name)
	if err := createOrUpdateRuntimeClass(r.Client, ctx, componentSet.RuntimeClass); err != nil {
		klog.Errorf("[%s] Error: While creating/updating runtimeclass (%s), err:%v", instance.Name, componentSet.RuntimeClass.Name, err)
		return reconcile.Result{}, fmt.Errorf("unable to create/update runtimeclass %q from %q. %w", componentSet.RuntimeClass.Name, instance.Name, err)
	}
	klog.Infof("[%s] Processed ok", instance.Name)
	r.Recorder.Eventf(instance, corev1.EventTypeNormal, "Creation succeeded", "[hypershift:%s] Succeeded to create all components", instance.Name)
	return reconcile.Result{}, nil
}

func readRuntimeClass(cli client.Client, ctx context.Context, name string) (*nodev1.RuntimeClass, error) {
	rtClass := &nodev1.RuntimeClass{}

	key := types.NamespacedName{
		Name: name,
	}

	if err := cli.Get(ctx, key, rtClass); err != nil {
		return nil, err
	}
	return rtClass, nil
}

func encodeManifest(obj runtime.Object, scheme *runtime.Scheme) ([]byte, error) {
	yamlSerializer := serializer.NewSerializerWithOptions(
		serializer.DefaultMetaFactory, scheme, scheme,
		serializer.SerializerOptions{Yaml: true, Pretty: true, Strict: true},
	)

	buff := bytes.Buffer{}
	err := yamlSerializer.Encode(obj, &buff)
	return buff.Bytes(), err
}

func createOrUpdateTunedConfigMap(cm *corev1.ConfigMap, ctx context.Context, cli client.Client) error {
	tunedConfigMapUpdateFunc := func(orig, dst *corev1.ConfigMap) error {
		dst.Data[tunedConfigMapConfigKey] = orig.Data[tunedConfigMapConfigKey]
		return nil
	}
	return createOrUpdateConfigMap(ctx, cli, cm, tunedConfigMapUpdateFunc)
}

func createOrUpdateMachineConfigConfigMap(cm *corev1.ConfigMap, ctx context.Context, cli client.Client) error {
	machineconfigConfigMapUpdateFunc := func(orig, dst *corev1.ConfigMap) error {
		dst.Data[mcoConfigMapConfigKey] = orig.Data[mcoConfigMapConfigKey]
		return nil
	}
	return createOrUpdateConfigMap(ctx, cli, cm, machineconfigConfigMapUpdateFunc)
}

func createOrUpdateKubeletConfigConfigConfigMap(cm *corev1.ConfigMap, ctx context.Context, cli client.Client) error {
	kubeletconfigConfigMapUpdateFunc := func(orig, dst *corev1.ConfigMap) error {
		dst.Data[mcoConfigMapConfigKey] = orig.Data[mcoConfigMapConfigKey]
		return nil
	}
	return createOrUpdateConfigMap(ctx, cli, cm, kubeletconfigConfigMapUpdateFunc)
}

func createOrUpdateConfigMap(ctx context.Context, cli client.Client, cm *corev1.ConfigMap, updateFunc func(origin, destination *corev1.ConfigMap) error) error {
	tcm := &corev1.ConfigMap{}
	err := cli.Get(ctx, client.ObjectKeyFromObject(cm), tcm)
	if err != nil && !k8serrors.IsNotFound(err) {
		return fmt.Errorf("failed to read configmap %q: %w", cm.Name, err)
	} else if k8serrors.IsNotFound(err) {
		//create
		if err := cli.Create(ctx, cm); err != nil {
			return fmt.Errorf("failed to create configmap %q: %w", cm.Name, err)
		}
	} else {
		// update
		if err := updateFunc(cm, tcm); err != nil {
			return fmt.Errorf("failed while updateing configmap content %q: %w", cm.Name, err)
		}
		if err := cli.Update(ctx, tcm); err != nil {
			return fmt.Errorf("failed to update configmap %q: %w", cm.Name, err)
		}
	}
	return nil
}

func newOwnerReference(owner metav1.Object, gvk schema.GroupVersionKind) metav1.OwnerReference {
	blockOwnerDeletion := false
	isController := false
	return metav1.OwnerReference{
		APIVersion:         gvk.GroupVersion().String(),
		Kind:               gvk.Kind,
		Name:               owner.GetName(),
		UID:                owner.GetUID(),
		BlockOwnerDeletion: &blockOwnerDeletion,
		Controller:         &isController,
	}
}

func createOrUpdateRuntimeClass(cli client.Client, ctx context.Context, rtClass *nodev1.RuntimeClass) error {
	existing, err := readRuntimeClass(cli, ctx, rtClass.Name)
	if err != nil {
		if k8serrors.IsNotFound(err) {
			err := cli.Create(ctx, rtClass)
			return err
		}
	}

	mutated := existing.DeepCopy()
	mergeMaps(rtClass.Annotations, mutated.Annotations)
	mergeMaps(rtClass.Labels, mutated.Labels)
	mutated.Handler = rtClass.Handler
	mutated.Scheduling = rtClass.Scheduling

	// we do not need to update if it no change between mutated and existing object
	if apiequality.Semantic.DeepEqual(existing.Handler, mutated.Handler) &&
		apiequality.Semantic.DeepEqual(existing.Scheduling, mutated.Scheduling) &&
		apiequality.Semantic.DeepEqual(existing.Labels, mutated.Labels) &&
		apiequality.Semantic.DeepEqual(existing.Annotations, mutated.Annotations) {
		return nil
	}

	err = cli.Update(ctx, mutated, &client.UpdateOptions{})
	return err
}

func TunedConfigMap(owner *corev1.ConfigMap, performanceProfileName, nodePoolNamespacedName, tunedManifest string) *corev1.ConfigMap {
	return &corev1.ConfigMap{
		ObjectMeta: metav1.ObjectMeta{
			Namespace: owner.GetNamespace(),
			Name:      "tuned-" + owner.Name,
			Labels: map[string]string{
				tunedConfigMapLabel:                   "true",
				hypershiftPerformanceProfileNameLabel: performanceProfileName,
				hypershiftNodePoolLabel:               parseNamespacedName(nodePoolNamespacedName),
			},
			Annotations: map[string]string{
				hypershiftNodePoolLabel: nodePoolNamespacedName,
			},
			OwnerReferences: []metav1.OwnerReference{
				newOwnerReference(owner, owner.GroupVersionKind()),
			},
		},
		Data: map[string]string{
			tunedConfigMapConfigKey: tunedManifest,
		},
	}
}

func MachineConfigConfigMap(owner *corev1.ConfigMap, performanceProfileName string, nodePoolNamespacedName string, machineconfigManifest string) *corev1.ConfigMap {
	return &corev1.ConfigMap{
		ObjectMeta: metav1.ObjectMeta{
			Namespace: owner.GetNamespace(),
			Name:      "mc-" + owner.Name,
			Labels: map[string]string{
				ntoGeneratedMachineConfigLabel:        "true",
				hypershiftPerformanceProfileNameLabel: performanceProfileName,
				hypershiftNodePoolLabel:               parseNamespacedName(nodePoolNamespacedName),
			},
			Annotations: map[string]string{
				hypershiftNodePoolLabel: nodePoolNamespacedName,
			},
			OwnerReferences: []metav1.OwnerReference{
				newOwnerReference(owner, owner.GroupVersionKind()),
			},
		},
		Data: map[string]string{
			mcoConfigMapConfigKey: machineconfigManifest,
		},
	}
}

func KubeletConfigConfigMap(owner *corev1.ConfigMap, performanceProfileName string, nodePoolNamespacedName string, kubeletconfigManifest string) *corev1.ConfigMap {
	return &corev1.ConfigMap{
		ObjectMeta: metav1.ObjectMeta{
			Namespace: owner.GetNamespace(),
			Name:      "kc-" + owner.Name,
			Labels: map[string]string{
				ntoGeneratedMachineConfigLabel:        "true",
				hypershiftPerformanceProfileNameLabel: performanceProfileName,
				hypershiftNodePoolLabel:               parseNamespacedName(nodePoolNamespacedName),
			},
			Annotations: map[string]string{
				hypershiftNodePoolLabel: nodePoolNamespacedName,
			},
			OwnerReferences: []metav1.OwnerReference{
				newOwnerReference(owner, owner.GroupVersionKind()),
			},
		},
		Data: map[string]string{
			mcoConfigMapConfigKey: kubeletconfigManifest,
		},
	}
}

func convertMachineConfig(origMC *mcov1.MachineConfig) *mcfgv1.MachineConfig {
	return &mcfgv1.MachineConfig{
		TypeMeta: metav1.TypeMeta{
			APIVersion: mcfgv1.SchemeGroupVersion.String(),
			Kind:       "MachineConfig",
		},
		ObjectMeta: origMC.ObjectMeta,
		// althoug mco.MC.Spec and hypershift mco.MC.Spec are defined in
		// different files they have the same structure, so the conversion is
		// almost direct ( but for the BaseOSExtensionsContainerImage)
		Spec: mcfgv1.MachineConfigSpec{
			OSImageURL:      origMC.Spec.OSImageURL,
			Config:          origMC.Spec.Config,
			KernelArguments: origMC.Spec.KernelArguments,
			Extensions:      origMC.Spec.Extensions,
			FIPS:            origMC.Spec.FIPS,
			KernelType:      origMC.Spec.KernelType,
		},
	}
}

func convertKubeletConfig(origKC *mcov1.KubeletConfig) *mcfgv1.KubeletConfig {
	return &mcfgv1.KubeletConfig{
		TypeMeta: metav1.TypeMeta{
			APIVersion: mcfgv1.SchemeGroupVersion.String(),
			Kind:       "KubeletConfig",
		},
		ObjectMeta: origKC.ObjectMeta,
		//NOTE - MachineConfigPoolSelector left empty as NodePool is the one
		// that relates nodes with MachineConfigs in hypershift.
		Spec: mcfgv1.KubeletConfigSpec{
			KubeletConfig: origKC.Spec.KubeletConfig,
		},
	}
}

// parseManifests parses a YAML or JSON document that may contain one or more
// kubernetes resources.
func parsePerformanceProfileManifest(data []byte, nodePoolName string, scheme *runtime.Scheme) (*performancev2.PerformanceProfile, error) {
	yamlSerializer := serializer.NewSerializerWithOptions(
		serializer.DefaultMetaFactory, scheme, scheme,
		serializer.SerializerOptions{Yaml: true, Pretty: true, Strict: true},
	)

	cr, _, err := yamlSerializer.Decode(data, nil, nil)
	if err != nil {
		return nil, fmt.Errorf("error decoding PerformanceProfile manifests: %v", err)
	}
	performanceProfile, ok := cr.(*performancev2.PerformanceProfile)
	if !ok {
		return nil, fmt.Errorf("error parsing PerformanceProfile manifests")
	}
	return performanceProfile, nil
}

func updatePerformanceProfileName(performanceProfile *performancev2.PerformanceProfile, nodePoolName string) {
	// Make PerformanceProfile names unique if a PerformanceProfile is duplicated across NodePools
	// for example, if one ConfigMap is referenced in multiple NodePools
	if !strings.HasSuffix(performanceProfile.ObjectMeta.Name, "-"+hashStruct(nodePoolName)) {
		performanceProfile.SetName(performanceProfile.ObjectMeta.Name + "-" + hashStruct(nodePoolName))
	}
}

func hypershiftDeleteComponents(remoteClient client.Client, ctx context.Context, ppConfigMap *corev1.ConfigMap) error {
	// ConfigMap is marked for deletion and waiting for finalizers to be empty
	// so better to clean-up and delete the objects.

	// just delete RtClass, as right now all the other components created from this PP
	// are embedded into ConfigMaps which has an OwnerReference with the PP configmap
	// and will be deleted by k8s machinery trigerring the deletion procedure of the
	// embedded elements.
	rtName := runtimeclass.BuildRuntimeClassName(ppConfigMap.Name)
	rtClass, err := readRuntimeClass(remoteClient, ctx, rtName)
	if err != nil {
		if k8serrors.IsNotFound(err) {
			// rtClass not found so nothing to delete, so delete process has finished
			return nil
		}
		return fmt.Errorf("unable to read RuntimeClass %q, error: %w. Unable to finalize deletion procedure properly", rtName, err)
	}

	err = remoteClient.Delete(ctx, rtClass)
	return err
}

func configMapHasFinalizer(cm *corev1.ConfigMap, finalizer string) bool {
	for _, f := range cm.Finalizers {
		if f == finalizer {
			return true
		}
	}
	return false
}

func configMapRemoveFinalizer(cm *corev1.ConfigMap, finalizer string) *corev1.ConfigMap {
	var finalizers []string
	for _, finalizer := range cm.Finalizers {
		if finalizer != hypershiftFinalizer {
			finalizers = append(finalizers, finalizer)
		}
	}
	cm.Finalizers = finalizers

	return cm
}

func hashStruct(o interface{}) string {
	hash := fnv.New32a()
	hash.Write([]byte(fmt.Sprintf("%v", o)))
	intHash := hash.Sum32()
	return fmt.Sprintf("%08x", intHash)
}

// parseNamespacedName expects a string with the format "namespace/name"
// and returns the name only.
// If given a string in the format "name" returns "name".
func parseNamespacedName(namespacedName string) string {
	parts := strings.SplitN(namespacedName, "/", 2)
	if len(parts) > 1 {
		return parts[1]
	}
	return parts[0]
}
