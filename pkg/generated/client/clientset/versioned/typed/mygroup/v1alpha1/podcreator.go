/*
Copyright The Kubernetes Authors.

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

package v1alpha1

import (
	"context"
	"time"

	v1alpha1 "github.com/Chandan-DK/kubernetes-custom-controller/api/mygroup/v1alpha1"
	scheme "github.com/Chandan-DK/kubernetes-custom-controller/pkg/generated/client/clientset/versioned/scheme"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	rest "k8s.io/client-go/rest"
)

// PodCreatorsGetter has a method to return a PodCreatorInterface.
// A group's client should implement this interface.
type PodCreatorsGetter interface {
	PodCreators(namespace string) PodCreatorInterface
}

// PodCreatorInterface has methods to work with PodCreator resources.
type PodCreatorInterface interface {
	Create(ctx context.Context, podCreator *v1alpha1.PodCreator, opts v1.CreateOptions) (*v1alpha1.PodCreator, error)
	Update(ctx context.Context, podCreator *v1alpha1.PodCreator, opts v1.UpdateOptions) (*v1alpha1.PodCreator, error)
	Delete(ctx context.Context, name string, opts v1.DeleteOptions) error
	DeleteCollection(ctx context.Context, opts v1.DeleteOptions, listOpts v1.ListOptions) error
	Get(ctx context.Context, name string, opts v1.GetOptions) (*v1alpha1.PodCreator, error)
	List(ctx context.Context, opts v1.ListOptions) (*v1alpha1.PodCreatorList, error)
	Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error)
	Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string) (result *v1alpha1.PodCreator, err error)
	PodCreatorExpansion
}

// podCreators implements PodCreatorInterface
type podCreators struct {
	client rest.Interface
	ns     string
}

// newPodCreators returns a PodCreators
func newPodCreators(c *MygroupV1alpha1Client, namespace string) *podCreators {
	return &podCreators{
		client: c.RESTClient(),
		ns:     namespace,
	}
}

// Get takes name of the podCreator, and returns the corresponding podCreator object, and an error if there is any.
func (c *podCreators) Get(ctx context.Context, name string, options v1.GetOptions) (result *v1alpha1.PodCreator, err error) {
	result = &v1alpha1.PodCreator{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("podcreators").
		Name(name).
		VersionedParams(&options, scheme.ParameterCodec).
		Do(ctx).
		Into(result)
	return
}

// List takes label and field selectors, and returns the list of PodCreators that match those selectors.
func (c *podCreators) List(ctx context.Context, opts v1.ListOptions) (result *v1alpha1.PodCreatorList, err error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	result = &v1alpha1.PodCreatorList{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("podcreators").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Do(ctx).
		Into(result)
	return
}

// Watch returns a watch.Interface that watches the requested podCreators.
func (c *podCreators) Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	opts.Watch = true
	return c.client.Get().
		Namespace(c.ns).
		Resource("podcreators").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Watch(ctx)
}

// Create takes the representation of a podCreator and creates it.  Returns the server's representation of the podCreator, and an error, if there is any.
func (c *podCreators) Create(ctx context.Context, podCreator *v1alpha1.PodCreator, opts v1.CreateOptions) (result *v1alpha1.PodCreator, err error) {
	result = &v1alpha1.PodCreator{}
	err = c.client.Post().
		Namespace(c.ns).
		Resource("podcreators").
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(podCreator).
		Do(ctx).
		Into(result)
	return
}

// Update takes the representation of a podCreator and updates it. Returns the server's representation of the podCreator, and an error, if there is any.
func (c *podCreators) Update(ctx context.Context, podCreator *v1alpha1.PodCreator, opts v1.UpdateOptions) (result *v1alpha1.PodCreator, err error) {
	result = &v1alpha1.PodCreator{}
	err = c.client.Put().
		Namespace(c.ns).
		Resource("podcreators").
		Name(podCreator.Name).
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(podCreator).
		Do(ctx).
		Into(result)
	return
}

// Delete takes name of the podCreator and deletes it. Returns an error if one occurs.
func (c *podCreators) Delete(ctx context.Context, name string, opts v1.DeleteOptions) error {
	return c.client.Delete().
		Namespace(c.ns).
		Resource("podcreators").
		Name(name).
		Body(&opts).
		Do(ctx).
		Error()
}

// DeleteCollection deletes a collection of objects.
func (c *podCreators) DeleteCollection(ctx context.Context, opts v1.DeleteOptions, listOpts v1.ListOptions) error {
	var timeout time.Duration
	if listOpts.TimeoutSeconds != nil {
		timeout = time.Duration(*listOpts.TimeoutSeconds) * time.Second
	}
	return c.client.Delete().
		Namespace(c.ns).
		Resource("podcreators").
		VersionedParams(&listOpts, scheme.ParameterCodec).
		Timeout(timeout).
		Body(&opts).
		Do(ctx).
		Error()
}

// Patch applies the patch and returns the patched podCreator.
func (c *podCreators) Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string) (result *v1alpha1.PodCreator, err error) {
	result = &v1alpha1.PodCreator{}
	err = c.client.Patch(pt).
		Namespace(c.ns).
		Resource("podcreators").
		Name(name).
		SubResource(subresources...).
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(data).
		Do(ctx).
		Into(result)
	return
}
