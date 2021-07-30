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
	"context"
	time "time"

	serverv1alpha1 "kubeform.dev/provider-hcloud-api/apis/server/v1alpha1"
	versioned "kubeform.dev/provider-hcloud-api/client/clientset/versioned"
	internalinterfaces "kubeform.dev/provider-hcloud-api/client/informers/externalversions/internalinterfaces"
	v1alpha1 "kubeform.dev/provider-hcloud-api/client/listers/server/v1alpha1"

	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	runtime "k8s.io/apimachinery/pkg/runtime"
	watch "k8s.io/apimachinery/pkg/watch"
	cache "k8s.io/client-go/tools/cache"
)

// ServerInformer provides access to a shared informer and lister for
// Servers.
type ServerInformer interface {
	Informer() cache.SharedIndexInformer
	Lister() v1alpha1.ServerLister
}

type serverInformer struct {
	factory          internalinterfaces.SharedInformerFactory
	tweakListOptions internalinterfaces.TweakListOptionsFunc
	namespace        string
}

// NewServerInformer constructs a new informer for Server type.
// Always prefer using an informer factory to get a shared informer instead of getting an independent
// one. This reduces memory footprint and number of connections to the server.
func NewServerInformer(client versioned.Interface, namespace string, resyncPeriod time.Duration, indexers cache.Indexers) cache.SharedIndexInformer {
	return NewFilteredServerInformer(client, namespace, resyncPeriod, indexers, nil)
}

// NewFilteredServerInformer constructs a new informer for Server type.
// Always prefer using an informer factory to get a shared informer instead of getting an independent
// one. This reduces memory footprint and number of connections to the server.
func NewFilteredServerInformer(client versioned.Interface, namespace string, resyncPeriod time.Duration, indexers cache.Indexers, tweakListOptions internalinterfaces.TweakListOptionsFunc) cache.SharedIndexInformer {
	return cache.NewSharedIndexInformer(
		&cache.ListWatch{
			ListFunc: func(options v1.ListOptions) (runtime.Object, error) {
				if tweakListOptions != nil {
					tweakListOptions(&options)
				}
				return client.ServerV1alpha1().Servers(namespace).List(context.TODO(), options)
			},
			WatchFunc: func(options v1.ListOptions) (watch.Interface, error) {
				if tweakListOptions != nil {
					tweakListOptions(&options)
				}
				return client.ServerV1alpha1().Servers(namespace).Watch(context.TODO(), options)
			},
		},
		&serverv1alpha1.Server{},
		resyncPeriod,
		indexers,
	)
}

func (f *serverInformer) defaultInformer(client versioned.Interface, resyncPeriod time.Duration) cache.SharedIndexInformer {
	return NewFilteredServerInformer(client, f.namespace, resyncPeriod, cache.Indexers{cache.NamespaceIndex: cache.MetaNamespaceIndexFunc}, f.tweakListOptions)
}

func (f *serverInformer) Informer() cache.SharedIndexInformer {
	return f.factory.InformerFor(&serverv1alpha1.Server{}, f.defaultInformer)
}

func (f *serverInformer) Lister() v1alpha1.ServerLister {
	return v1alpha1.NewServerLister(f.Informer().GetIndexer())
}
