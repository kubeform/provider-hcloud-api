/*
Copyright AppsCode Inc. and Contributors

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

// Code generated by informer-gen. DO NOT EDIT.

package v1alpha1

import (
	internalinterfaces "kubeform.dev/provider-hcloud-api/client/informers/externalversions/internalinterfaces"
)

// Interface provides access to all the informers in this group version.
type Interface interface {
	// Balancers returns a BalancerInformer.
	Balancers() BalancerInformer
	// BalancerNetworks returns a BalancerNetworkInformer.
	BalancerNetworks() BalancerNetworkInformer
	// BalancerServices returns a BalancerServiceInformer.
	BalancerServices() BalancerServiceInformer
	// BalancerTargets returns a BalancerTargetInformer.
	BalancerTargets() BalancerTargetInformer
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

// Balancers returns a BalancerInformer.
func (v *version) Balancers() BalancerInformer {
	return &balancerInformer{factory: v.factory, namespace: v.namespace, tweakListOptions: v.tweakListOptions}
}

// BalancerNetworks returns a BalancerNetworkInformer.
func (v *version) BalancerNetworks() BalancerNetworkInformer {
	return &balancerNetworkInformer{factory: v.factory, namespace: v.namespace, tweakListOptions: v.tweakListOptions}
}

// BalancerServices returns a BalancerServiceInformer.
func (v *version) BalancerServices() BalancerServiceInformer {
	return &balancerServiceInformer{factory: v.factory, namespace: v.namespace, tweakListOptions: v.tweakListOptions}
}

// BalancerTargets returns a BalancerTargetInformer.
func (v *version) BalancerTargets() BalancerTargetInformer {
	return &balancerTargetInformer{factory: v.factory, namespace: v.namespace, tweakListOptions: v.tweakListOptions}
}
