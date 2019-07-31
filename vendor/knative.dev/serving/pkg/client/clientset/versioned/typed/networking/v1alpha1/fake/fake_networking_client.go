/*
Copyright 2019 The Knative Authors

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
	rest "k8s.io/client-go/rest"
	testing "k8s.io/client-go/testing"
	v1alpha1 "knative.dev/serving/pkg/client/clientset/versioned/typed/networking/v1alpha1"
)

type FakeNetworkingV1alpha1 struct {
	*testing.Fake
}

func (c *FakeNetworkingV1alpha1) Certificates(namespace string) v1alpha1.CertificateInterface {
	return &FakeCertificates{c, namespace}
}

func (c *FakeNetworkingV1alpha1) ClusterIngresses() v1alpha1.ClusterIngressInterface {
	return &FakeClusterIngresses{c}
}

func (c *FakeNetworkingV1alpha1) Ingresses(namespace string) v1alpha1.IngressInterface {
	return &FakeIngresses{c, namespace}
}

func (c *FakeNetworkingV1alpha1) ServerlessServices(namespace string) v1alpha1.ServerlessServiceInterface {
	return &FakeServerlessServices{c, namespace}
}

// RESTClient returns a RESTClient that is used to communicate
// with API server by this client implementation.
func (c *FakeNetworkingV1alpha1) RESTClient() rest.Interface {
	var ret *rest.RESTClient
	return ret
}
