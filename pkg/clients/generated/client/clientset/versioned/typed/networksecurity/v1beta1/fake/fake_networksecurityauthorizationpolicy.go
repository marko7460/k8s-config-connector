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

	v1beta1 "github.com/GoogleCloudPlatform/k8s-config-connector/pkg/clients/generated/apis/networksecurity/v1beta1"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	labels "k8s.io/apimachinery/pkg/labels"
	schema "k8s.io/apimachinery/pkg/runtime/schema"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	testing "k8s.io/client-go/testing"
)

// FakeNetworkSecurityAuthorizationPolicies implements NetworkSecurityAuthorizationPolicyInterface
type FakeNetworkSecurityAuthorizationPolicies struct {
	Fake *FakeNetworksecurityV1beta1
	ns   string
}

var networksecurityauthorizationpoliciesResource = schema.GroupVersionResource{Group: "networksecurity.cnrm.cloud.google.com", Version: "v1beta1", Resource: "networksecurityauthorizationpolicies"}

var networksecurityauthorizationpoliciesKind = schema.GroupVersionKind{Group: "networksecurity.cnrm.cloud.google.com", Version: "v1beta1", Kind: "NetworkSecurityAuthorizationPolicy"}

// Get takes name of the networkSecurityAuthorizationPolicy, and returns the corresponding networkSecurityAuthorizationPolicy object, and an error if there is any.
func (c *FakeNetworkSecurityAuthorizationPolicies) Get(ctx context.Context, name string, options v1.GetOptions) (result *v1beta1.NetworkSecurityAuthorizationPolicy, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewGetAction(networksecurityauthorizationpoliciesResource, c.ns, name), &v1beta1.NetworkSecurityAuthorizationPolicy{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1beta1.NetworkSecurityAuthorizationPolicy), err
}

// List takes label and field selectors, and returns the list of NetworkSecurityAuthorizationPolicies that match those selectors.
func (c *FakeNetworkSecurityAuthorizationPolicies) List(ctx context.Context, opts v1.ListOptions) (result *v1beta1.NetworkSecurityAuthorizationPolicyList, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewListAction(networksecurityauthorizationpoliciesResource, networksecurityauthorizationpoliciesKind, c.ns, opts), &v1beta1.NetworkSecurityAuthorizationPolicyList{})

	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &v1beta1.NetworkSecurityAuthorizationPolicyList{ListMeta: obj.(*v1beta1.NetworkSecurityAuthorizationPolicyList).ListMeta}
	for _, item := range obj.(*v1beta1.NetworkSecurityAuthorizationPolicyList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested networkSecurityAuthorizationPolicies.
func (c *FakeNetworkSecurityAuthorizationPolicies) Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewWatchAction(networksecurityauthorizationpoliciesResource, c.ns, opts))

}

// Create takes the representation of a networkSecurityAuthorizationPolicy and creates it.  Returns the server's representation of the networkSecurityAuthorizationPolicy, and an error, if there is any.
func (c *FakeNetworkSecurityAuthorizationPolicies) Create(ctx context.Context, networkSecurityAuthorizationPolicy *v1beta1.NetworkSecurityAuthorizationPolicy, opts v1.CreateOptions) (result *v1beta1.NetworkSecurityAuthorizationPolicy, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewCreateAction(networksecurityauthorizationpoliciesResource, c.ns, networkSecurityAuthorizationPolicy), &v1beta1.NetworkSecurityAuthorizationPolicy{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1beta1.NetworkSecurityAuthorizationPolicy), err
}

// Update takes the representation of a networkSecurityAuthorizationPolicy and updates it. Returns the server's representation of the networkSecurityAuthorizationPolicy, and an error, if there is any.
func (c *FakeNetworkSecurityAuthorizationPolicies) Update(ctx context.Context, networkSecurityAuthorizationPolicy *v1beta1.NetworkSecurityAuthorizationPolicy, opts v1.UpdateOptions) (result *v1beta1.NetworkSecurityAuthorizationPolicy, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateAction(networksecurityauthorizationpoliciesResource, c.ns, networkSecurityAuthorizationPolicy), &v1beta1.NetworkSecurityAuthorizationPolicy{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1beta1.NetworkSecurityAuthorizationPolicy), err
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
func (c *FakeNetworkSecurityAuthorizationPolicies) UpdateStatus(ctx context.Context, networkSecurityAuthorizationPolicy *v1beta1.NetworkSecurityAuthorizationPolicy, opts v1.UpdateOptions) (*v1beta1.NetworkSecurityAuthorizationPolicy, error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateSubresourceAction(networksecurityauthorizationpoliciesResource, "status", c.ns, networkSecurityAuthorizationPolicy), &v1beta1.NetworkSecurityAuthorizationPolicy{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1beta1.NetworkSecurityAuthorizationPolicy), err
}

// Delete takes name of the networkSecurityAuthorizationPolicy and deletes it. Returns an error if one occurs.
func (c *FakeNetworkSecurityAuthorizationPolicies) Delete(ctx context.Context, name string, opts v1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewDeleteActionWithOptions(networksecurityauthorizationpoliciesResource, c.ns, name, opts), &v1beta1.NetworkSecurityAuthorizationPolicy{})

	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakeNetworkSecurityAuthorizationPolicies) DeleteCollection(ctx context.Context, opts v1.DeleteOptions, listOpts v1.ListOptions) error {
	action := testing.NewDeleteCollectionAction(networksecurityauthorizationpoliciesResource, c.ns, listOpts)

	_, err := c.Fake.Invokes(action, &v1beta1.NetworkSecurityAuthorizationPolicyList{})
	return err
}

// Patch applies the patch and returns the patched networkSecurityAuthorizationPolicy.
func (c *FakeNetworkSecurityAuthorizationPolicies) Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string) (result *v1beta1.NetworkSecurityAuthorizationPolicy, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewPatchSubresourceAction(networksecurityauthorizationpoliciesResource, c.ns, name, pt, data, subresources...), &v1beta1.NetworkSecurityAuthorizationPolicy{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1beta1.NetworkSecurityAuthorizationPolicy), err
}
