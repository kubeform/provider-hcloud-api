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

// Code generated by client-gen. DO NOT EDIT.

package fake

import (
	v1alpha1 "kubeform.dev/provider-hcloud-api/client/clientset/versioned/typed/load/v1alpha1"

	rest "k8s.io/client-go/rest"
	testing "k8s.io/client-go/testing"
)

type FakeLoadV1alpha1 struct {
	*testing.Fake
}

func (c *FakeLoadV1alpha1) Balancers(namespace string) v1alpha1.BalancerInterface {
	return &FakeBalancers{c, namespace}
}

func (c *FakeLoadV1alpha1) BalancerNetworks(namespace string) v1alpha1.BalancerNetworkInterface {
	return &FakeBalancerNetworks{c, namespace}
}

func (c *FakeLoadV1alpha1) BalancerServices(namespace string) v1alpha1.BalancerServiceInterface {
	return &FakeBalancerServices{c, namespace}
}

func (c *FakeLoadV1alpha1) BalancerTargets(namespace string) v1alpha1.BalancerTargetInterface {
	return &FakeBalancerTargets{c, namespace}
}

// RESTClient returns a RESTClient that is used to communicate
// with API server by this client implementation.
func (c *FakeLoadV1alpha1) RESTClient() rest.Interface {
	var ret *rest.RESTClient
	return ret
}
