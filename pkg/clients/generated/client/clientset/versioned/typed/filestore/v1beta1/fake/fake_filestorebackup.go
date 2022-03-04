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

	v1beta1 "github.com/GoogleCloudPlatform/k8s-config-connector/pkg/clients/generated/apis/filestore/v1beta1"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	labels "k8s.io/apimachinery/pkg/labels"
	schema "k8s.io/apimachinery/pkg/runtime/schema"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	testing "k8s.io/client-go/testing"
)

// FakeFilestoreBackups implements FilestoreBackupInterface
type FakeFilestoreBackups struct {
	Fake *FakeFilestoreV1beta1
	ns   string
}

var filestorebackupsResource = schema.GroupVersionResource{Group: "filestore.cnrm.cloud.google.com", Version: "v1beta1", Resource: "filestorebackups"}

var filestorebackupsKind = schema.GroupVersionKind{Group: "filestore.cnrm.cloud.google.com", Version: "v1beta1", Kind: "FilestoreBackup"}

// Get takes name of the filestoreBackup, and returns the corresponding filestoreBackup object, and an error if there is any.
func (c *FakeFilestoreBackups) Get(ctx context.Context, name string, options v1.GetOptions) (result *v1beta1.FilestoreBackup, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewGetAction(filestorebackupsResource, c.ns, name), &v1beta1.FilestoreBackup{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1beta1.FilestoreBackup), err
}

// List takes label and field selectors, and returns the list of FilestoreBackups that match those selectors.
func (c *FakeFilestoreBackups) List(ctx context.Context, opts v1.ListOptions) (result *v1beta1.FilestoreBackupList, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewListAction(filestorebackupsResource, filestorebackupsKind, c.ns, opts), &v1beta1.FilestoreBackupList{})

	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &v1beta1.FilestoreBackupList{ListMeta: obj.(*v1beta1.FilestoreBackupList).ListMeta}
	for _, item := range obj.(*v1beta1.FilestoreBackupList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested filestoreBackups.
func (c *FakeFilestoreBackups) Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewWatchAction(filestorebackupsResource, c.ns, opts))

}

// Create takes the representation of a filestoreBackup and creates it.  Returns the server's representation of the filestoreBackup, and an error, if there is any.
func (c *FakeFilestoreBackups) Create(ctx context.Context, filestoreBackup *v1beta1.FilestoreBackup, opts v1.CreateOptions) (result *v1beta1.FilestoreBackup, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewCreateAction(filestorebackupsResource, c.ns, filestoreBackup), &v1beta1.FilestoreBackup{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1beta1.FilestoreBackup), err
}

// Update takes the representation of a filestoreBackup and updates it. Returns the server's representation of the filestoreBackup, and an error, if there is any.
func (c *FakeFilestoreBackups) Update(ctx context.Context, filestoreBackup *v1beta1.FilestoreBackup, opts v1.UpdateOptions) (result *v1beta1.FilestoreBackup, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateAction(filestorebackupsResource, c.ns, filestoreBackup), &v1beta1.FilestoreBackup{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1beta1.FilestoreBackup), err
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
func (c *FakeFilestoreBackups) UpdateStatus(ctx context.Context, filestoreBackup *v1beta1.FilestoreBackup, opts v1.UpdateOptions) (*v1beta1.FilestoreBackup, error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateSubresourceAction(filestorebackupsResource, "status", c.ns, filestoreBackup), &v1beta1.FilestoreBackup{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1beta1.FilestoreBackup), err
}

// Delete takes name of the filestoreBackup and deletes it. Returns an error if one occurs.
func (c *FakeFilestoreBackups) Delete(ctx context.Context, name string, opts v1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewDeleteActionWithOptions(filestorebackupsResource, c.ns, name, opts), &v1beta1.FilestoreBackup{})

	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakeFilestoreBackups) DeleteCollection(ctx context.Context, opts v1.DeleteOptions, listOpts v1.ListOptions) error {
	action := testing.NewDeleteCollectionAction(filestorebackupsResource, c.ns, listOpts)

	_, err := c.Fake.Invokes(action, &v1beta1.FilestoreBackupList{})
	return err
}

// Patch applies the patch and returns the patched filestoreBackup.
func (c *FakeFilestoreBackups) Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string) (result *v1beta1.FilestoreBackup, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewPatchSubresourceAction(filestorebackupsResource, c.ns, name, pt, data, subresources...), &v1beta1.FilestoreBackup{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1beta1.FilestoreBackup), err
}
