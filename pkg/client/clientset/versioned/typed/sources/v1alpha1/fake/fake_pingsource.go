/*
Copyright 2020 The Knative Authors

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
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	labels "k8s.io/apimachinery/pkg/labels"
	schema "k8s.io/apimachinery/pkg/runtime/schema"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	testing "k8s.io/client-go/testing"
	v1alpha1 "knative.dev/eventing/pkg/apis/sources/v1alpha1"
)

// FakePingSources implements PingSourceInterface
type FakePingSources struct {
	Fake *FakeSourcesV1alpha1
	ns   string
}

var pingsourcesResource = schema.GroupVersionResource{Group: "sources.knative.dev", Version: "v1alpha1", Resource: "pingsources"}

var pingsourcesKind = schema.GroupVersionKind{Group: "sources.knative.dev", Version: "v1alpha1", Kind: "PingSource"}

// Get takes name of the pingSource, and returns the corresponding pingSource object, and an error if there is any.
func (c *FakePingSources) Get(name string, options v1.GetOptions) (result *v1alpha1.PingSource, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewGetAction(pingsourcesResource, c.ns, name), &v1alpha1.PingSource{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.PingSource), err
}

// List takes label and field selectors, and returns the list of PingSources that match those selectors.
func (c *FakePingSources) List(opts v1.ListOptions) (result *v1alpha1.PingSourceList, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewListAction(pingsourcesResource, pingsourcesKind, c.ns, opts), &v1alpha1.PingSourceList{})

	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &v1alpha1.PingSourceList{ListMeta: obj.(*v1alpha1.PingSourceList).ListMeta}
	for _, item := range obj.(*v1alpha1.PingSourceList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested pingSources.
func (c *FakePingSources) Watch(opts v1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewWatchAction(pingsourcesResource, c.ns, opts))

}

// Create takes the representation of a pingSource and creates it.  Returns the server's representation of the pingSource, and an error, if there is any.
func (c *FakePingSources) Create(pingSource *v1alpha1.PingSource) (result *v1alpha1.PingSource, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewCreateAction(pingsourcesResource, c.ns, pingSource), &v1alpha1.PingSource{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.PingSource), err
}

// Update takes the representation of a pingSource and updates it. Returns the server's representation of the pingSource, and an error, if there is any.
func (c *FakePingSources) Update(pingSource *v1alpha1.PingSource) (result *v1alpha1.PingSource, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateAction(pingsourcesResource, c.ns, pingSource), &v1alpha1.PingSource{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.PingSource), err
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
func (c *FakePingSources) UpdateStatus(pingSource *v1alpha1.PingSource) (*v1alpha1.PingSource, error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateSubresourceAction(pingsourcesResource, "status", c.ns, pingSource), &v1alpha1.PingSource{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.PingSource), err
}

// Delete takes name of the pingSource and deletes it. Returns an error if one occurs.
func (c *FakePingSources) Delete(name string, options *v1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewDeleteAction(pingsourcesResource, c.ns, name), &v1alpha1.PingSource{})

	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakePingSources) DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error {
	action := testing.NewDeleteCollectionAction(pingsourcesResource, c.ns, listOptions)

	_, err := c.Fake.Invokes(action, &v1alpha1.PingSourceList{})
	return err
}

// Patch applies the patch and returns the patched pingSource.
func (c *FakePingSources) Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1alpha1.PingSource, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewPatchSubresourceAction(pingsourcesResource, c.ns, name, pt, data, subresources...), &v1alpha1.PingSource{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.PingSource), err
}
