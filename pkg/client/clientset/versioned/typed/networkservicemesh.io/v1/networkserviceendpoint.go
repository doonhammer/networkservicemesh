// Copyright (c) 2018 Cisco and/or its affiliates.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at:
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

// Code generated by client-gen. DO NOT EDIT.

package v1

import (
	v1 "github.com/ligato/networkservicemesh/pkg/apis/networkservicemesh.io/v1"
	scheme "github.com/ligato/networkservicemesh/pkg/client/clientset/versioned/scheme"
	meta_v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	rest "k8s.io/client-go/rest"
)

// NetworkServiceEndpointsGetter has a method to return a NetworkServiceEndpointInterface.
// A group's client should implement this interface.
type NetworkServiceEndpointsGetter interface {
	NetworkServiceEndpoints(namespace string) NetworkServiceEndpointInterface
}

// NetworkServiceEndpointInterface has methods to work with NetworkServiceEndpoint resources.
type NetworkServiceEndpointInterface interface {
	Create(*v1.NetworkServiceEndpoint) (*v1.NetworkServiceEndpoint, error)
	Update(*v1.NetworkServiceEndpoint) (*v1.NetworkServiceEndpoint, error)
	UpdateStatus(*v1.NetworkServiceEndpoint) (*v1.NetworkServiceEndpoint, error)
	Delete(name string, options *meta_v1.DeleteOptions) error
	DeleteCollection(options *meta_v1.DeleteOptions, listOptions meta_v1.ListOptions) error
	Get(name string, options meta_v1.GetOptions) (*v1.NetworkServiceEndpoint, error)
	List(opts meta_v1.ListOptions) (*v1.NetworkServiceEndpointList, error)
	Watch(opts meta_v1.ListOptions) (watch.Interface, error)
	Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1.NetworkServiceEndpoint, err error)
	NetworkServiceEndpointExpansion
}

// networkServiceEndpoints implements NetworkServiceEndpointInterface
type networkServiceEndpoints struct {
	client rest.Interface
	ns     string
}

// newNetworkServiceEndpoints returns a NetworkServiceEndpoints
func newNetworkServiceEndpoints(c *NetworkserviceV1Client, namespace string) *networkServiceEndpoints {
	return &networkServiceEndpoints{
		client: c.RESTClient(),
		ns:     namespace,
	}
}

// Get takes name of the networkServiceEndpoint, and returns the corresponding networkServiceEndpoint object, and an error if there is any.
func (c *networkServiceEndpoints) Get(name string, options meta_v1.GetOptions) (result *v1.NetworkServiceEndpoint, err error) {
	result = &v1.NetworkServiceEndpoint{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("networkserviceendpoints").
		Name(name).
		VersionedParams(&options, scheme.ParameterCodec).
		Do().
		Into(result)
	return
}

// List takes label and field selectors, and returns the list of NetworkServiceEndpoints that match those selectors.
func (c *networkServiceEndpoints) List(opts meta_v1.ListOptions) (result *v1.NetworkServiceEndpointList, err error) {
	result = &v1.NetworkServiceEndpointList{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("networkserviceendpoints").
		VersionedParams(&opts, scheme.ParameterCodec).
		Do().
		Into(result)
	return
}

// Watch returns a watch.Interface that watches the requested networkServiceEndpoints.
func (c *networkServiceEndpoints) Watch(opts meta_v1.ListOptions) (watch.Interface, error) {
	opts.Watch = true
	return c.client.Get().
		Namespace(c.ns).
		Resource("networkserviceendpoints").
		VersionedParams(&opts, scheme.ParameterCodec).
		Watch()
}

// Create takes the representation of a networkServiceEndpoint and creates it.  Returns the server's representation of the networkServiceEndpoint, and an error, if there is any.
func (c *networkServiceEndpoints) Create(networkServiceEndpoint *v1.NetworkServiceEndpoint) (result *v1.NetworkServiceEndpoint, err error) {
	result = &v1.NetworkServiceEndpoint{}
	err = c.client.Post().
		Namespace(c.ns).
		Resource("networkserviceendpoints").
		Body(networkServiceEndpoint).
		Do().
		Into(result)
	return
}

// Update takes the representation of a networkServiceEndpoint and updates it. Returns the server's representation of the networkServiceEndpoint, and an error, if there is any.
func (c *networkServiceEndpoints) Update(networkServiceEndpoint *v1.NetworkServiceEndpoint) (result *v1.NetworkServiceEndpoint, err error) {
	result = &v1.NetworkServiceEndpoint{}
	err = c.client.Put().
		Namespace(c.ns).
		Resource("networkserviceendpoints").
		Name(networkServiceEndpoint.Name).
		Body(networkServiceEndpoint).
		Do().
		Into(result)
	return
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().

func (c *networkServiceEndpoints) UpdateStatus(networkServiceEndpoint *v1.NetworkServiceEndpoint) (result *v1.NetworkServiceEndpoint, err error) {
	result = &v1.NetworkServiceEndpoint{}
	err = c.client.Put().
		Namespace(c.ns).
		Resource("networkserviceendpoints").
		Name(networkServiceEndpoint.Name).
		SubResource("status").
		Body(networkServiceEndpoint).
		Do().
		Into(result)
	return
}

// Delete takes name of the networkServiceEndpoint and deletes it. Returns an error if one occurs.
func (c *networkServiceEndpoints) Delete(name string, options *meta_v1.DeleteOptions) error {
	return c.client.Delete().
		Namespace(c.ns).
		Resource("networkserviceendpoints").
		Name(name).
		Body(options).
		Do().
		Error()
}

// DeleteCollection deletes a collection of objects.
func (c *networkServiceEndpoints) DeleteCollection(options *meta_v1.DeleteOptions, listOptions meta_v1.ListOptions) error {
	return c.client.Delete().
		Namespace(c.ns).
		Resource("networkserviceendpoints").
		VersionedParams(&listOptions, scheme.ParameterCodec).
		Body(options).
		Do().
		Error()
}

// Patch applies the patch and returns the patched networkServiceEndpoint.
func (c *networkServiceEndpoints) Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1.NetworkServiceEndpoint, err error) {
	result = &v1.NetworkServiceEndpoint{}
	err = c.client.Patch(pt).
		Namespace(c.ns).
		Resource("networkserviceendpoints").
		SubResource(subresources...).
		Name(name).
		Body(data).
		Do().
		Into(result)
	return
}
