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

	v1beta1 "github.com/GoogleCloudPlatform/k8s-config-connector/pkg/clients/generated/apis/monitoring/v1beta1"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	labels "k8s.io/apimachinery/pkg/labels"
	schema "k8s.io/apimachinery/pkg/runtime/schema"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	testing "k8s.io/client-go/testing"
)

// FakeMonitoringDashboards implements MonitoringDashboardInterface
type FakeMonitoringDashboards struct {
	Fake *FakeMonitoringV1beta1
	ns   string
}

var monitoringdashboardsResource = schema.GroupVersionResource{Group: "monitoring.cnrm.cloud.google.com", Version: "v1beta1", Resource: "monitoringdashboards"}

var monitoringdashboardsKind = schema.GroupVersionKind{Group: "monitoring.cnrm.cloud.google.com", Version: "v1beta1", Kind: "MonitoringDashboard"}

// Get takes name of the monitoringDashboard, and returns the corresponding monitoringDashboard object, and an error if there is any.
func (c *FakeMonitoringDashboards) Get(ctx context.Context, name string, options v1.GetOptions) (result *v1beta1.MonitoringDashboard, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewGetAction(monitoringdashboardsResource, c.ns, name), &v1beta1.MonitoringDashboard{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1beta1.MonitoringDashboard), err
}

// List takes label and field selectors, and returns the list of MonitoringDashboards that match those selectors.
func (c *FakeMonitoringDashboards) List(ctx context.Context, opts v1.ListOptions) (result *v1beta1.MonitoringDashboardList, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewListAction(monitoringdashboardsResource, monitoringdashboardsKind, c.ns, opts), &v1beta1.MonitoringDashboardList{})

	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &v1beta1.MonitoringDashboardList{ListMeta: obj.(*v1beta1.MonitoringDashboardList).ListMeta}
	for _, item := range obj.(*v1beta1.MonitoringDashboardList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested monitoringDashboards.
func (c *FakeMonitoringDashboards) Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewWatchAction(monitoringdashboardsResource, c.ns, opts))

}

// Create takes the representation of a monitoringDashboard and creates it.  Returns the server's representation of the monitoringDashboard, and an error, if there is any.
func (c *FakeMonitoringDashboards) Create(ctx context.Context, monitoringDashboard *v1beta1.MonitoringDashboard, opts v1.CreateOptions) (result *v1beta1.MonitoringDashboard, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewCreateAction(monitoringdashboardsResource, c.ns, monitoringDashboard), &v1beta1.MonitoringDashboard{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1beta1.MonitoringDashboard), err
}

// Update takes the representation of a monitoringDashboard and updates it. Returns the server's representation of the monitoringDashboard, and an error, if there is any.
func (c *FakeMonitoringDashboards) Update(ctx context.Context, monitoringDashboard *v1beta1.MonitoringDashboard, opts v1.UpdateOptions) (result *v1beta1.MonitoringDashboard, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateAction(monitoringdashboardsResource, c.ns, monitoringDashboard), &v1beta1.MonitoringDashboard{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1beta1.MonitoringDashboard), err
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
func (c *FakeMonitoringDashboards) UpdateStatus(ctx context.Context, monitoringDashboard *v1beta1.MonitoringDashboard, opts v1.UpdateOptions) (*v1beta1.MonitoringDashboard, error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateSubresourceAction(monitoringdashboardsResource, "status", c.ns, monitoringDashboard), &v1beta1.MonitoringDashboard{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1beta1.MonitoringDashboard), err
}

// Delete takes name of the monitoringDashboard and deletes it. Returns an error if one occurs.
func (c *FakeMonitoringDashboards) Delete(ctx context.Context, name string, opts v1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewDeleteActionWithOptions(monitoringdashboardsResource, c.ns, name, opts), &v1beta1.MonitoringDashboard{})

	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakeMonitoringDashboards) DeleteCollection(ctx context.Context, opts v1.DeleteOptions, listOpts v1.ListOptions) error {
	action := testing.NewDeleteCollectionAction(monitoringdashboardsResource, c.ns, listOpts)

	_, err := c.Fake.Invokes(action, &v1beta1.MonitoringDashboardList{})
	return err
}

// Patch applies the patch and returns the patched monitoringDashboard.
func (c *FakeMonitoringDashboards) Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string) (result *v1beta1.MonitoringDashboard, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewPatchSubresourceAction(monitoringdashboardsResource, c.ns, name, pt, data, subresources...), &v1beta1.MonitoringDashboard{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1beta1.MonitoringDashboard), err
}
