/*
 * Copyright 2020 The Multicluster-Scheduler Authors.
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 *     http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 */

// Code generated by client-gen. DO NOT EDIT.

package fake

import (
	v1alpha1 "admiralty.io/multicluster-scheduler/pkg/apis/multicluster/v1alpha1"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	labels "k8s.io/apimachinery/pkg/labels"
	schema "k8s.io/apimachinery/pkg/runtime/schema"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	testing "k8s.io/client-go/testing"
)

// FakePodChaperons implements PodChaperonInterface
type FakePodChaperons struct {
	Fake *FakeMulticlusterV1alpha1
	ns   string
}

var podchaperonsResource = schema.GroupVersionResource{Group: "multicluster.admiralty.io", Version: "v1alpha1", Resource: "podchaperons"}

var podchaperonsKind = schema.GroupVersionKind{Group: "multicluster.admiralty.io", Version: "v1alpha1", Kind: "PodChaperon"}

// Get takes name of the podChaperon, and returns the corresponding podChaperon object, and an error if there is any.
func (c *FakePodChaperons) Get(name string, options v1.GetOptions) (result *v1alpha1.PodChaperon, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewGetAction(podchaperonsResource, c.ns, name), &v1alpha1.PodChaperon{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.PodChaperon), err
}

// List takes label and field selectors, and returns the list of PodChaperons that match those selectors.
func (c *FakePodChaperons) List(opts v1.ListOptions) (result *v1alpha1.PodChaperonList, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewListAction(podchaperonsResource, podchaperonsKind, c.ns, opts), &v1alpha1.PodChaperonList{})

	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &v1alpha1.PodChaperonList{ListMeta: obj.(*v1alpha1.PodChaperonList).ListMeta}
	for _, item := range obj.(*v1alpha1.PodChaperonList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested podChaperons.
func (c *FakePodChaperons) Watch(opts v1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewWatchAction(podchaperonsResource, c.ns, opts))

}

// Create takes the representation of a podChaperon and creates it.  Returns the server's representation of the podChaperon, and an error, if there is any.
func (c *FakePodChaperons) Create(podChaperon *v1alpha1.PodChaperon) (result *v1alpha1.PodChaperon, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewCreateAction(podchaperonsResource, c.ns, podChaperon), &v1alpha1.PodChaperon{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.PodChaperon), err
}

// Update takes the representation of a podChaperon and updates it. Returns the server's representation of the podChaperon, and an error, if there is any.
func (c *FakePodChaperons) Update(podChaperon *v1alpha1.PodChaperon) (result *v1alpha1.PodChaperon, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateAction(podchaperonsResource, c.ns, podChaperon), &v1alpha1.PodChaperon{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.PodChaperon), err
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
func (c *FakePodChaperons) UpdateStatus(podChaperon *v1alpha1.PodChaperon) (*v1alpha1.PodChaperon, error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateSubresourceAction(podchaperonsResource, "status", c.ns, podChaperon), &v1alpha1.PodChaperon{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.PodChaperon), err
}

// Delete takes name of the podChaperon and deletes it. Returns an error if one occurs.
func (c *FakePodChaperons) Delete(name string, options *v1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewDeleteAction(podchaperonsResource, c.ns, name), &v1alpha1.PodChaperon{})

	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakePodChaperons) DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error {
	action := testing.NewDeleteCollectionAction(podchaperonsResource, c.ns, listOptions)

	_, err := c.Fake.Invokes(action, &v1alpha1.PodChaperonList{})
	return err
}

// Patch applies the patch and returns the patched podChaperon.
func (c *FakePodChaperons) Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1alpha1.PodChaperon, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewPatchSubresourceAction(podchaperonsResource, c.ns, name, pt, data, subresources...), &v1alpha1.PodChaperon{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.PodChaperon), err
}
