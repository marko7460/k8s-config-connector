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

	v1beta1 "github.com/GoogleCloudPlatform/k8s-config-connector/pkg/clients/generated/apis/storage/v1beta1"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	labels "k8s.io/apimachinery/pkg/labels"
	schema "k8s.io/apimachinery/pkg/runtime/schema"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	testing "k8s.io/client-go/testing"
)

// FakeStorageBuckets implements StorageBucketInterface
type FakeStorageBuckets struct {
	Fake *FakeStorageV1beta1
	ns   string
}

var storagebucketsResource = schema.GroupVersionResource{Group: "storage.cnrm.cloud.google.com", Version: "v1beta1", Resource: "storagebuckets"}

var storagebucketsKind = schema.GroupVersionKind{Group: "storage.cnrm.cloud.google.com", Version: "v1beta1", Kind: "StorageBucket"}

// Get takes name of the storageBucket, and returns the corresponding storageBucket object, and an error if there is any.
func (c *FakeStorageBuckets) Get(ctx context.Context, name string, options v1.GetOptions) (result *v1beta1.StorageBucket, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewGetAction(storagebucketsResource, c.ns, name), &v1beta1.StorageBucket{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1beta1.StorageBucket), err
}

// List takes label and field selectors, and returns the list of StorageBuckets that match those selectors.
func (c *FakeStorageBuckets) List(ctx context.Context, opts v1.ListOptions) (result *v1beta1.StorageBucketList, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewListAction(storagebucketsResource, storagebucketsKind, c.ns, opts), &v1beta1.StorageBucketList{})

	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &v1beta1.StorageBucketList{ListMeta: obj.(*v1beta1.StorageBucketList).ListMeta}
	for _, item := range obj.(*v1beta1.StorageBucketList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested storageBuckets.
func (c *FakeStorageBuckets) Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewWatchAction(storagebucketsResource, c.ns, opts))

}

// Create takes the representation of a storageBucket and creates it.  Returns the server's representation of the storageBucket, and an error, if there is any.
func (c *FakeStorageBuckets) Create(ctx context.Context, storageBucket *v1beta1.StorageBucket, opts v1.CreateOptions) (result *v1beta1.StorageBucket, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewCreateAction(storagebucketsResource, c.ns, storageBucket), &v1beta1.StorageBucket{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1beta1.StorageBucket), err
}

// Update takes the representation of a storageBucket and updates it. Returns the server's representation of the storageBucket, and an error, if there is any.
func (c *FakeStorageBuckets) Update(ctx context.Context, storageBucket *v1beta1.StorageBucket, opts v1.UpdateOptions) (result *v1beta1.StorageBucket, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateAction(storagebucketsResource, c.ns, storageBucket), &v1beta1.StorageBucket{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1beta1.StorageBucket), err
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
func (c *FakeStorageBuckets) UpdateStatus(ctx context.Context, storageBucket *v1beta1.StorageBucket, opts v1.UpdateOptions) (*v1beta1.StorageBucket, error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateSubresourceAction(storagebucketsResource, "status", c.ns, storageBucket), &v1beta1.StorageBucket{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1beta1.StorageBucket), err
}

// Delete takes name of the storageBucket and deletes it. Returns an error if one occurs.
func (c *FakeStorageBuckets) Delete(ctx context.Context, name string, opts v1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewDeleteActionWithOptions(storagebucketsResource, c.ns, name, opts), &v1beta1.StorageBucket{})

	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakeStorageBuckets) DeleteCollection(ctx context.Context, opts v1.DeleteOptions, listOpts v1.ListOptions) error {
	action := testing.NewDeleteCollectionAction(storagebucketsResource, c.ns, listOpts)

	_, err := c.Fake.Invokes(action, &v1beta1.StorageBucketList{})
	return err
}

// Patch applies the patch and returns the patched storageBucket.
func (c *FakeStorageBuckets) Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string) (result *v1beta1.StorageBucket, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewPatchSubresourceAction(storagebucketsResource, c.ns, name, pt, data, subresources...), &v1beta1.StorageBucket{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1beta1.StorageBucket), err
}
