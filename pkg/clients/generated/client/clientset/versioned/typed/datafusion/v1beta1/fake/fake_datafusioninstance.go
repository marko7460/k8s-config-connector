// Copyright 2020 Google LLC
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

// *** DISCLAIMER ***
// Config Connector's go-client for CRDs is currently in ALPHA, which means
// that future versions of the go-client may include breaking changes.
// Please try it out and give us feedback!

// Code generated by main. DO NOT EDIT.

package fake

import (
	"context"

	v1beta1 "github.com/GoogleCloudPlatform/k8s-config-connector/pkg/clients/generated/apis/datafusion/v1beta1"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	labels "k8s.io/apimachinery/pkg/labels"
	schema "k8s.io/apimachinery/pkg/runtime/schema"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	testing "k8s.io/client-go/testing"
)

// FakeDataFusionInstances implements DataFusionInstanceInterface
type FakeDataFusionInstances struct {
	Fake *FakeDatafusionV1beta1
	ns   string
}

var datafusioninstancesResource = schema.GroupVersionResource{Group: "datafusion.cnrm.cloud.google.com", Version: "v1beta1", Resource: "datafusioninstances"}

var datafusioninstancesKind = schema.GroupVersionKind{Group: "datafusion.cnrm.cloud.google.com", Version: "v1beta1", Kind: "DataFusionInstance"}

// Get takes name of the dataFusionInstance, and returns the corresponding dataFusionInstance object, and an error if there is any.
func (c *FakeDataFusionInstances) Get(ctx context.Context, name string, options v1.GetOptions) (result *v1beta1.DataFusionInstance, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewGetAction(datafusioninstancesResource, c.ns, name), &v1beta1.DataFusionInstance{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1beta1.DataFusionInstance), err
}

// List takes label and field selectors, and returns the list of DataFusionInstances that match those selectors.
func (c *FakeDataFusionInstances) List(ctx context.Context, opts v1.ListOptions) (result *v1beta1.DataFusionInstanceList, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewListAction(datafusioninstancesResource, datafusioninstancesKind, c.ns, opts), &v1beta1.DataFusionInstanceList{})

	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &v1beta1.DataFusionInstanceList{ListMeta: obj.(*v1beta1.DataFusionInstanceList).ListMeta}
	for _, item := range obj.(*v1beta1.DataFusionInstanceList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested dataFusionInstances.
func (c *FakeDataFusionInstances) Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewWatchAction(datafusioninstancesResource, c.ns, opts))

}

// Create takes the representation of a dataFusionInstance and creates it.  Returns the server's representation of the dataFusionInstance, and an error, if there is any.
func (c *FakeDataFusionInstances) Create(ctx context.Context, dataFusionInstance *v1beta1.DataFusionInstance, opts v1.CreateOptions) (result *v1beta1.DataFusionInstance, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewCreateAction(datafusioninstancesResource, c.ns, dataFusionInstance), &v1beta1.DataFusionInstance{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1beta1.DataFusionInstance), err
}

// Update takes the representation of a dataFusionInstance and updates it. Returns the server's representation of the dataFusionInstance, and an error, if there is any.
func (c *FakeDataFusionInstances) Update(ctx context.Context, dataFusionInstance *v1beta1.DataFusionInstance, opts v1.UpdateOptions) (result *v1beta1.DataFusionInstance, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateAction(datafusioninstancesResource, c.ns, dataFusionInstance), &v1beta1.DataFusionInstance{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1beta1.DataFusionInstance), err
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
func (c *FakeDataFusionInstances) UpdateStatus(ctx context.Context, dataFusionInstance *v1beta1.DataFusionInstance, opts v1.UpdateOptions) (*v1beta1.DataFusionInstance, error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateSubresourceAction(datafusioninstancesResource, "status", c.ns, dataFusionInstance), &v1beta1.DataFusionInstance{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1beta1.DataFusionInstance), err
}

// Delete takes name of the dataFusionInstance and deletes it. Returns an error if one occurs.
func (c *FakeDataFusionInstances) Delete(ctx context.Context, name string, opts v1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewDeleteActionWithOptions(datafusioninstancesResource, c.ns, name, opts), &v1beta1.DataFusionInstance{})

	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakeDataFusionInstances) DeleteCollection(ctx context.Context, opts v1.DeleteOptions, listOpts v1.ListOptions) error {
	action := testing.NewDeleteCollectionAction(datafusioninstancesResource, c.ns, listOpts)

	_, err := c.Fake.Invokes(action, &v1beta1.DataFusionInstanceList{})
	return err
}

// Patch applies the patch and returns the patched dataFusionInstance.
func (c *FakeDataFusionInstances) Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string) (result *v1beta1.DataFusionInstance, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewPatchSubresourceAction(datafusioninstancesResource, c.ns, name, pt, data, subresources...), &v1beta1.DataFusionInstance{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1beta1.DataFusionInstance), err
}
