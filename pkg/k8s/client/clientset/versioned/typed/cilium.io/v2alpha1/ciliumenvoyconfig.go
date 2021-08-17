// SPDX-License-Identifier: Apache-2.0
// Copyright 2017-2021 Authors of Cilium

// Code generated by client-gen. DO NOT EDIT.

package v2alpha1

import (
	"context"
	"time"

	v2alpha1 "github.com/cilium/cilium/pkg/k8s/apis/cilium.io/v2alpha1"
	scheme "github.com/cilium/cilium/pkg/k8s/client/clientset/versioned/scheme"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	rest "k8s.io/client-go/rest"
)

// CiliumEnvoyConfigsGetter has a method to return a CiliumEnvoyConfigInterface.
// A group's client should implement this interface.
type CiliumEnvoyConfigsGetter interface {
	CiliumEnvoyConfigs() CiliumEnvoyConfigInterface
}

// CiliumEnvoyConfigInterface has methods to work with CiliumEnvoyConfig resources.
type CiliumEnvoyConfigInterface interface {
	Create(ctx context.Context, ciliumEnvoyConfig *v2alpha1.CiliumEnvoyConfig, opts v1.CreateOptions) (*v2alpha1.CiliumEnvoyConfig, error)
	Update(ctx context.Context, ciliumEnvoyConfig *v2alpha1.CiliumEnvoyConfig, opts v1.UpdateOptions) (*v2alpha1.CiliumEnvoyConfig, error)
	Delete(ctx context.Context, name string, opts v1.DeleteOptions) error
	DeleteCollection(ctx context.Context, opts v1.DeleteOptions, listOpts v1.ListOptions) error
	Get(ctx context.Context, name string, opts v1.GetOptions) (*v2alpha1.CiliumEnvoyConfig, error)
	List(ctx context.Context, opts v1.ListOptions) (*v2alpha1.CiliumEnvoyConfigList, error)
	Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error)
	Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string) (result *v2alpha1.CiliumEnvoyConfig, err error)
	CiliumEnvoyConfigExpansion
}

// ciliumEnvoyConfigs implements CiliumEnvoyConfigInterface
type ciliumEnvoyConfigs struct {
	client rest.Interface
}

// newCiliumEnvoyConfigs returns a CiliumEnvoyConfigs
func newCiliumEnvoyConfigs(c *CiliumV2alpha1Client) *ciliumEnvoyConfigs {
	return &ciliumEnvoyConfigs{
		client: c.RESTClient(),
	}
}

// Get takes name of the ciliumEnvoyConfig, and returns the corresponding ciliumEnvoyConfig object, and an error if there is any.
func (c *ciliumEnvoyConfigs) Get(ctx context.Context, name string, options v1.GetOptions) (result *v2alpha1.CiliumEnvoyConfig, err error) {
	result = &v2alpha1.CiliumEnvoyConfig{}
	err = c.client.Get().
		Resource("ciliumenvoyconfigs").
		Name(name).
		VersionedParams(&options, scheme.ParameterCodec).
		Do(ctx).
		Into(result)
	return
}

// List takes label and field selectors, and returns the list of CiliumEnvoyConfigs that match those selectors.
func (c *ciliumEnvoyConfigs) List(ctx context.Context, opts v1.ListOptions) (result *v2alpha1.CiliumEnvoyConfigList, err error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	result = &v2alpha1.CiliumEnvoyConfigList{}
	err = c.client.Get().
		Resource("ciliumenvoyconfigs").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Do(ctx).
		Into(result)
	return
}

// Watch returns a watch.Interface that watches the requested ciliumEnvoyConfigs.
func (c *ciliumEnvoyConfigs) Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	opts.Watch = true
	return c.client.Get().
		Resource("ciliumenvoyconfigs").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Watch(ctx)
}

// Create takes the representation of a ciliumEnvoyConfig and creates it.  Returns the server's representation of the ciliumEnvoyConfig, and an error, if there is any.
func (c *ciliumEnvoyConfigs) Create(ctx context.Context, ciliumEnvoyConfig *v2alpha1.CiliumEnvoyConfig, opts v1.CreateOptions) (result *v2alpha1.CiliumEnvoyConfig, err error) {
	result = &v2alpha1.CiliumEnvoyConfig{}
	err = c.client.Post().
		Resource("ciliumenvoyconfigs").
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(ciliumEnvoyConfig).
		Do(ctx).
		Into(result)
	return
}

// Update takes the representation of a ciliumEnvoyConfig and updates it. Returns the server's representation of the ciliumEnvoyConfig, and an error, if there is any.
func (c *ciliumEnvoyConfigs) Update(ctx context.Context, ciliumEnvoyConfig *v2alpha1.CiliumEnvoyConfig, opts v1.UpdateOptions) (result *v2alpha1.CiliumEnvoyConfig, err error) {
	result = &v2alpha1.CiliumEnvoyConfig{}
	err = c.client.Put().
		Resource("ciliumenvoyconfigs").
		Name(ciliumEnvoyConfig.Name).
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(ciliumEnvoyConfig).
		Do(ctx).
		Into(result)
	return
}

// Delete takes name of the ciliumEnvoyConfig and deletes it. Returns an error if one occurs.
func (c *ciliumEnvoyConfigs) Delete(ctx context.Context, name string, opts v1.DeleteOptions) error {
	return c.client.Delete().
		Resource("ciliumenvoyconfigs").
		Name(name).
		Body(&opts).
		Do(ctx).
		Error()
}

// DeleteCollection deletes a collection of objects.
func (c *ciliumEnvoyConfigs) DeleteCollection(ctx context.Context, opts v1.DeleteOptions, listOpts v1.ListOptions) error {
	var timeout time.Duration
	if listOpts.TimeoutSeconds != nil {
		timeout = time.Duration(*listOpts.TimeoutSeconds) * time.Second
	}
	return c.client.Delete().
		Resource("ciliumenvoyconfigs").
		VersionedParams(&listOpts, scheme.ParameterCodec).
		Timeout(timeout).
		Body(&opts).
		Do(ctx).
		Error()
}

// Patch applies the patch and returns the patched ciliumEnvoyConfig.
func (c *ciliumEnvoyConfigs) Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string) (result *v2alpha1.CiliumEnvoyConfig, err error) {
	result = &v2alpha1.CiliumEnvoyConfig{}
	err = c.client.Patch(pt).
		Resource("ciliumenvoyconfigs").
		Name(name).
		SubResource(subresources...).
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(data).
		Do(ctx).
		Into(result)
	return
}
