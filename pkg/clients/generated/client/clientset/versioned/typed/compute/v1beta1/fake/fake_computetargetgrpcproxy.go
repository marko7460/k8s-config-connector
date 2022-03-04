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

	v1beta1 "github.com/GoogleCloudPlatform/k8s-config-connector/pkg/clients/generated/apis/compute/v1beta1"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	labels "k8s.io/apimachinery/pkg/labels"
	schema "k8s.io/apimachinery/pkg/runtime/schema"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	testing "k8s.io/client-go/testing"
)

// FakeComputeTargetGRPCProxies implements ComputeTargetGRPCProxyInterface
type FakeComputeTargetGRPCProxies struct {
	Fake *FakeComputeV1beta1
	ns   string
}

var computetargetgrpcproxiesResource = schema.GroupVersionResource{Group: "compute.cnrm.cloud.google.com", Version: "v1beta1", Resource: "computetargetgrpcproxies"}

var computetargetgrpcproxiesKind = schema.GroupVersionKind{Group: "compute.cnrm.cloud.google.com", Version: "v1beta1", Kind: "ComputeTargetGRPCProxy"}

// Get takes name of the computeTargetGRPCProxy, and returns the corresponding computeTargetGRPCProxy object, and an error if there is any.
func (c *FakeComputeTargetGRPCProxies) Get(ctx context.Context, name string, options v1.GetOptions) (result *v1beta1.ComputeTargetGRPCProxy, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewGetAction(computetargetgrpcproxiesResource, c.ns, name), &v1beta1.ComputeTargetGRPCProxy{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1beta1.ComputeTargetGRPCProxy), err
}

// List takes label and field selectors, and returns the list of ComputeTargetGRPCProxies that match those selectors.
func (c *FakeComputeTargetGRPCProxies) List(ctx context.Context, opts v1.ListOptions) (result *v1beta1.ComputeTargetGRPCProxyList, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewListAction(computetargetgrpcproxiesResource, computetargetgrpcproxiesKind, c.ns, opts), &v1beta1.ComputeTargetGRPCProxyList{})

	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &v1beta1.ComputeTargetGRPCProxyList{ListMeta: obj.(*v1beta1.ComputeTargetGRPCProxyList).ListMeta}
	for _, item := range obj.(*v1beta1.ComputeTargetGRPCProxyList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested computeTargetGRPCProxies.
func (c *FakeComputeTargetGRPCProxies) Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewWatchAction(computetargetgrpcproxiesResource, c.ns, opts))

}

// Create takes the representation of a computeTargetGRPCProxy and creates it.  Returns the server's representation of the computeTargetGRPCProxy, and an error, if there is any.
func (c *FakeComputeTargetGRPCProxies) Create(ctx context.Context, computeTargetGRPCProxy *v1beta1.ComputeTargetGRPCProxy, opts v1.CreateOptions) (result *v1beta1.ComputeTargetGRPCProxy, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewCreateAction(computetargetgrpcproxiesResource, c.ns, computeTargetGRPCProxy), &v1beta1.ComputeTargetGRPCProxy{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1beta1.ComputeTargetGRPCProxy), err
}

// Update takes the representation of a computeTargetGRPCProxy and updates it. Returns the server's representation of the computeTargetGRPCProxy, and an error, if there is any.
func (c *FakeComputeTargetGRPCProxies) Update(ctx context.Context, computeTargetGRPCProxy *v1beta1.ComputeTargetGRPCProxy, opts v1.UpdateOptions) (result *v1beta1.ComputeTargetGRPCProxy, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateAction(computetargetgrpcproxiesResource, c.ns, computeTargetGRPCProxy), &v1beta1.ComputeTargetGRPCProxy{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1beta1.ComputeTargetGRPCProxy), err
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
func (c *FakeComputeTargetGRPCProxies) UpdateStatus(ctx context.Context, computeTargetGRPCProxy *v1beta1.ComputeTargetGRPCProxy, opts v1.UpdateOptions) (*v1beta1.ComputeTargetGRPCProxy, error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateSubresourceAction(computetargetgrpcproxiesResource, "status", c.ns, computeTargetGRPCProxy), &v1beta1.ComputeTargetGRPCProxy{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1beta1.ComputeTargetGRPCProxy), err
}

// Delete takes name of the computeTargetGRPCProxy and deletes it. Returns an error if one occurs.
func (c *FakeComputeTargetGRPCProxies) Delete(ctx context.Context, name string, opts v1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewDeleteActionWithOptions(computetargetgrpcproxiesResource, c.ns, name, opts), &v1beta1.ComputeTargetGRPCProxy{})

	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakeComputeTargetGRPCProxies) DeleteCollection(ctx context.Context, opts v1.DeleteOptions, listOpts v1.ListOptions) error {
	action := testing.NewDeleteCollectionAction(computetargetgrpcproxiesResource, c.ns, listOpts)

	_, err := c.Fake.Invokes(action, &v1beta1.ComputeTargetGRPCProxyList{})
	return err
}

// Patch applies the patch and returns the patched computeTargetGRPCProxy.
func (c *FakeComputeTargetGRPCProxies) Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string) (result *v1beta1.ComputeTargetGRPCProxy, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewPatchSubresourceAction(computetargetgrpcproxiesResource, c.ns, name, pt, data, subresources...), &v1beta1.ComputeTargetGRPCProxy{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1beta1.ComputeTargetGRPCProxy), err
}
