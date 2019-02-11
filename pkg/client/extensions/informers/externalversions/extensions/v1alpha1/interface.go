// Code generated by informer-gen. DO NOT EDIT.

package v1alpha1

import (
	internalinterfaces "github.com/gardener/gardener/pkg/client/extensions/informers/externalversions/internalinterfaces"
)

// Interface provides access to all the informers in this group version.
type Interface interface {
	// OperatingSystemConfigs returns a OperatingSystemConfigInformer.
	OperatingSystemConfigs() OperatingSystemConfigInformer
	// WorkerPools returns a WorkerPoolInformer.
	WorkerPools() WorkerPoolInformer
}

type version struct {
	factory          internalinterfaces.SharedInformerFactory
	namespace        string
	tweakListOptions internalinterfaces.TweakListOptionsFunc
}

// New returns a new Interface.
func New(f internalinterfaces.SharedInformerFactory, namespace string, tweakListOptions internalinterfaces.TweakListOptionsFunc) Interface {
	return &version{factory: f, namespace: namespace, tweakListOptions: tweakListOptions}
}

// OperatingSystemConfigs returns a OperatingSystemConfigInformer.
func (v *version) OperatingSystemConfigs() OperatingSystemConfigInformer {
	return &operatingSystemConfigInformer{factory: v.factory, namespace: v.namespace, tweakListOptions: v.tweakListOptions}
}

// WorkerPools returns a WorkerPoolInformer.
func (v *version) WorkerPools() WorkerPoolInformer {
	return &workerPoolInformer{factory: v.factory, namespace: v.namespace, tweakListOptions: v.tweakListOptions}
}
