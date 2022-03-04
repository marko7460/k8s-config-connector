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

// FakePrivateCACertificateAuthorities implements PrivateCACertificateAuthorityInterface
type FakePrivateCACertificateAuthorities struct {
	Fake *FakePrivatecaV1beta1
	ns   string
}

var privatecacertificateauthoritiesResource = schema.GroupVersionResource{Group: "privateca.cnrm.cloud.google.com", Version: "v1beta1", Resource: "privatecacertificateauthorities"}

var privatecacertificateauthoritiesKind = schema.GroupVersionKind{Group: "privateca.cnrm.cloud.google.com", Version: "v1beta1", Kind: "PrivateCACertificateAuthority"}

// Get takes name of the privateCACertificateAuthority, and returns the corresponding privateCACertificateAuthority object, and an error if there is any.
func (c *FakePrivateCACertificateAuthorities) Get(ctx context.Context, name string, options v1.GetOptions) (result *v1beta1.PrivateCACertificateAuthority, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewGetAction(privatecacertificateauthoritiesResource, c.ns, name), &v1beta1.PrivateCACertificateAuthority{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1beta1.PrivateCACertificateAuthority), err
}

// List takes label and field selectors, and returns the list of PrivateCACertificateAuthorities that match those selectors.
func (c *FakePrivateCACertificateAuthorities) List(ctx context.Context, opts v1.ListOptions) (result *v1beta1.PrivateCACertificateAuthorityList, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewListAction(privatecacertificateauthoritiesResource, privatecacertificateauthoritiesKind, c.ns, opts), &v1beta1.PrivateCACertificateAuthorityList{})

	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &v1beta1.PrivateCACertificateAuthorityList{ListMeta: obj.(*v1beta1.PrivateCACertificateAuthorityList).ListMeta}
	for _, item := range obj.(*v1beta1.PrivateCACertificateAuthorityList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested privateCACertificateAuthorities.
func (c *FakePrivateCACertificateAuthorities) Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewWatchAction(privatecacertificateauthoritiesResource, c.ns, opts))

}

// Create takes the representation of a privateCACertificateAuthority and creates it.  Returns the server's representation of the privateCACertificateAuthority, and an error, if there is any.
func (c *FakePrivateCACertificateAuthorities) Create(ctx context.Context, privateCACertificateAuthority *v1beta1.PrivateCACertificateAuthority, opts v1.CreateOptions) (result *v1beta1.PrivateCACertificateAuthority, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewCreateAction(privatecacertificateauthoritiesResource, c.ns, privateCACertificateAuthority), &v1beta1.PrivateCACertificateAuthority{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1beta1.PrivateCACertificateAuthority), err
}

// Update takes the representation of a privateCACertificateAuthority and updates it. Returns the server's representation of the privateCACertificateAuthority, and an error, if there is any.
func (c *FakePrivateCACertificateAuthorities) Update(ctx context.Context, privateCACertificateAuthority *v1beta1.PrivateCACertificateAuthority, opts v1.UpdateOptions) (result *v1beta1.PrivateCACertificateAuthority, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateAction(privatecacertificateauthoritiesResource, c.ns, privateCACertificateAuthority), &v1beta1.PrivateCACertificateAuthority{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1beta1.PrivateCACertificateAuthority), err
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
func (c *FakePrivateCACertificateAuthorities) UpdateStatus(ctx context.Context, privateCACertificateAuthority *v1beta1.PrivateCACertificateAuthority, opts v1.UpdateOptions) (*v1beta1.PrivateCACertificateAuthority, error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateSubresourceAction(privatecacertificateauthoritiesResource, "status", c.ns, privateCACertificateAuthority), &v1beta1.PrivateCACertificateAuthority{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1beta1.PrivateCACertificateAuthority), err
}

// Delete takes name of the privateCACertificateAuthority and deletes it. Returns an error if one occurs.
func (c *FakePrivateCACertificateAuthorities) Delete(ctx context.Context, name string, opts v1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewDeleteActionWithOptions(privatecacertificateauthoritiesResource, c.ns, name, opts), &v1beta1.PrivateCACertificateAuthority{})

	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakePrivateCACertificateAuthorities) DeleteCollection(ctx context.Context, opts v1.DeleteOptions, listOpts v1.ListOptions) error {
	action := testing.NewDeleteCollectionAction(privatecacertificateauthoritiesResource, c.ns, listOpts)

	_, err := c.Fake.Invokes(action, &v1beta1.PrivateCACertificateAuthorityList{})
	return err
}

// Patch applies the patch and returns the patched privateCACertificateAuthority.
func (c *FakePrivateCACertificateAuthorities) Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string) (result *v1beta1.PrivateCACertificateAuthority, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewPatchSubresourceAction(privatecacertificateauthoritiesResource, c.ns, name, pt, data, subresources...), &v1beta1.PrivateCACertificateAuthority{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1beta1.PrivateCACertificateAuthority), err
}
