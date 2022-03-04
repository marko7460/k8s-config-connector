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

	v1beta1 "github.com/GoogleCloudPlatform/k8s-config-connector/pkg/clients/generated/apis/privateca/v1beta1"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	labels "k8s.io/apimachinery/pkg/labels"
	schema "k8s.io/apimachinery/pkg/runtime/schema"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	testing "k8s.io/client-go/testing"
)

// FakePrivateCACertificateTemplates implements PrivateCACertificateTemplateInterface
type FakePrivateCACertificateTemplates struct {
	Fake *FakePrivatecaV1beta1
	ns   string
}

var privatecacertificatetemplatesResource = schema.GroupVersionResource{Group: "privateca.cnrm.cloud.google.com", Version: "v1beta1", Resource: "privatecacertificatetemplates"}

var privatecacertificatetemplatesKind = schema.GroupVersionKind{Group: "privateca.cnrm.cloud.google.com", Version: "v1beta1", Kind: "PrivateCACertificateTemplate"}

// Get takes name of the privateCACertificateTemplate, and returns the corresponding privateCACertificateTemplate object, and an error if there is any.
func (c *FakePrivateCACertificateTemplates) Get(ctx context.Context, name string, options v1.GetOptions) (result *v1beta1.PrivateCACertificateTemplate, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewGetAction(privatecacertificatetemplatesResource, c.ns, name), &v1beta1.PrivateCACertificateTemplate{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1beta1.PrivateCACertificateTemplate), err
}

// List takes label and field selectors, and returns the list of PrivateCACertificateTemplates that match those selectors.
func (c *FakePrivateCACertificateTemplates) List(ctx context.Context, opts v1.ListOptions) (result *v1beta1.PrivateCACertificateTemplateList, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewListAction(privatecacertificatetemplatesResource, privatecacertificatetemplatesKind, c.ns, opts), &v1beta1.PrivateCACertificateTemplateList{})

	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &v1beta1.PrivateCACertificateTemplateList{ListMeta: obj.(*v1beta1.PrivateCACertificateTemplateList).ListMeta}
	for _, item := range obj.(*v1beta1.PrivateCACertificateTemplateList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested privateCACertificateTemplates.
func (c *FakePrivateCACertificateTemplates) Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewWatchAction(privatecacertificatetemplatesResource, c.ns, opts))

}

// Create takes the representation of a privateCACertificateTemplate and creates it.  Returns the server's representation of the privateCACertificateTemplate, and an error, if there is any.
func (c *FakePrivateCACertificateTemplates) Create(ctx context.Context, privateCACertificateTemplate *v1beta1.PrivateCACertificateTemplate, opts v1.CreateOptions) (result *v1beta1.PrivateCACertificateTemplate, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewCreateAction(privatecacertificatetemplatesResource, c.ns, privateCACertificateTemplate), &v1beta1.PrivateCACertificateTemplate{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1beta1.PrivateCACertificateTemplate), err
}

// Update takes the representation of a privateCACertificateTemplate and updates it. Returns the server's representation of the privateCACertificateTemplate, and an error, if there is any.
func (c *FakePrivateCACertificateTemplates) Update(ctx context.Context, privateCACertificateTemplate *v1beta1.PrivateCACertificateTemplate, opts v1.UpdateOptions) (result *v1beta1.PrivateCACertificateTemplate, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateAction(privatecacertificatetemplatesResource, c.ns, privateCACertificateTemplate), &v1beta1.PrivateCACertificateTemplate{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1beta1.PrivateCACertificateTemplate), err
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
func (c *FakePrivateCACertificateTemplates) UpdateStatus(ctx context.Context, privateCACertificateTemplate *v1beta1.PrivateCACertificateTemplate, opts v1.UpdateOptions) (*v1beta1.PrivateCACertificateTemplate, error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateSubresourceAction(privatecacertificatetemplatesResource, "status", c.ns, privateCACertificateTemplate), &v1beta1.PrivateCACertificateTemplate{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1beta1.PrivateCACertificateTemplate), err
}

// Delete takes name of the privateCACertificateTemplate and deletes it. Returns an error if one occurs.
func (c *FakePrivateCACertificateTemplates) Delete(ctx context.Context, name string, opts v1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewDeleteActionWithOptions(privatecacertificatetemplatesResource, c.ns, name, opts), &v1beta1.PrivateCACertificateTemplate{})

	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakePrivateCACertificateTemplates) DeleteCollection(ctx context.Context, opts v1.DeleteOptions, listOpts v1.ListOptions) error {
	action := testing.NewDeleteCollectionAction(privatecacertificatetemplatesResource, c.ns, listOpts)

	_, err := c.Fake.Invokes(action, &v1beta1.PrivateCACertificateTemplateList{})
	return err
}

// Patch applies the patch and returns the patched privateCACertificateTemplate.
func (c *FakePrivateCACertificateTemplates) Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string) (result *v1beta1.PrivateCACertificateTemplate, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewPatchSubresourceAction(privatecacertificatetemplatesResource, c.ns, name, pt, data, subresources...), &v1beta1.PrivateCACertificateTemplate{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1beta1.PrivateCACertificateTemplate), err
}
