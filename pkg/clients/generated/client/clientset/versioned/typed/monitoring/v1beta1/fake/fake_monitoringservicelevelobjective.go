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

// FakeMonitoringServiceLevelObjectives implements MonitoringServiceLevelObjectiveInterface
type FakeMonitoringServiceLevelObjectives struct {
	Fake *FakeMonitoringV1beta1
	ns   string
}

var monitoringservicelevelobjectivesResource = schema.GroupVersionResource{Group: "monitoring.cnrm.cloud.google.com", Version: "v1beta1", Resource: "monitoringservicelevelobjectives"}

var monitoringservicelevelobjectivesKind = schema.GroupVersionKind{Group: "monitoring.cnrm.cloud.google.com", Version: "v1beta1", Kind: "MonitoringServiceLevelObjective"}

// Get takes name of the monitoringServiceLevelObjective, and returns the corresponding monitoringServiceLevelObjective object, and an error if there is any.
func (c *FakeMonitoringServiceLevelObjectives) Get(ctx context.Context, name string, options v1.GetOptions) (result *v1beta1.MonitoringServiceLevelObjective, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewGetAction(monitoringservicelevelobjectivesResource, c.ns, name), &v1beta1.MonitoringServiceLevelObjective{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1beta1.MonitoringServiceLevelObjective), err
}

// List takes label and field selectors, and returns the list of MonitoringServiceLevelObjectives that match those selectors.
func (c *FakeMonitoringServiceLevelObjectives) List(ctx context.Context, opts v1.ListOptions) (result *v1beta1.MonitoringServiceLevelObjectiveList, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewListAction(monitoringservicelevelobjectivesResource, monitoringservicelevelobjectivesKind, c.ns, opts), &v1beta1.MonitoringServiceLevelObjectiveList{})

	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &v1beta1.MonitoringServiceLevelObjectiveList{ListMeta: obj.(*v1beta1.MonitoringServiceLevelObjectiveList).ListMeta}
	for _, item := range obj.(*v1beta1.MonitoringServiceLevelObjectiveList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested monitoringServiceLevelObjectives.
func (c *FakeMonitoringServiceLevelObjectives) Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewWatchAction(monitoringservicelevelobjectivesResource, c.ns, opts))

}

// Create takes the representation of a monitoringServiceLevelObjective and creates it.  Returns the server's representation of the monitoringServiceLevelObjective, and an error, if there is any.
func (c *FakeMonitoringServiceLevelObjectives) Create(ctx context.Context, monitoringServiceLevelObjective *v1beta1.MonitoringServiceLevelObjective, opts v1.CreateOptions) (result *v1beta1.MonitoringServiceLevelObjective, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewCreateAction(monitoringservicelevelobjectivesResource, c.ns, monitoringServiceLevelObjective), &v1beta1.MonitoringServiceLevelObjective{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1beta1.MonitoringServiceLevelObjective), err
}

// Update takes the representation of a monitoringServiceLevelObjective and updates it. Returns the server's representation of the monitoringServiceLevelObjective, and an error, if there is any.
func (c *FakeMonitoringServiceLevelObjectives) Update(ctx context.Context, monitoringServiceLevelObjective *v1beta1.MonitoringServiceLevelObjective, opts v1.UpdateOptions) (result *v1beta1.MonitoringServiceLevelObjective, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateAction(monitoringservicelevelobjectivesResource, c.ns, monitoringServiceLevelObjective), &v1beta1.MonitoringServiceLevelObjective{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1beta1.MonitoringServiceLevelObjective), err
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
func (c *FakeMonitoringServiceLevelObjectives) UpdateStatus(ctx context.Context, monitoringServiceLevelObjective *v1beta1.MonitoringServiceLevelObjective, opts v1.UpdateOptions) (*v1beta1.MonitoringServiceLevelObjective, error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateSubresourceAction(monitoringservicelevelobjectivesResource, "status", c.ns, monitoringServiceLevelObjective), &v1beta1.MonitoringServiceLevelObjective{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1beta1.MonitoringServiceLevelObjective), err
}

// Delete takes name of the monitoringServiceLevelObjective and deletes it. Returns an error if one occurs.
func (c *FakeMonitoringServiceLevelObjectives) Delete(ctx context.Context, name string, opts v1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewDeleteActionWithOptions(monitoringservicelevelobjectivesResource, c.ns, name, opts), &v1beta1.MonitoringServiceLevelObjective{})

	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakeMonitoringServiceLevelObjectives) DeleteCollection(ctx context.Context, opts v1.DeleteOptions, listOpts v1.ListOptions) error {
	action := testing.NewDeleteCollectionAction(monitoringservicelevelobjectivesResource, c.ns, listOpts)

	_, err := c.Fake.Invokes(action, &v1beta1.MonitoringServiceLevelObjectiveList{})
	return err
}

// Patch applies the patch and returns the patched monitoringServiceLevelObjective.
func (c *FakeMonitoringServiceLevelObjectives) Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string) (result *v1beta1.MonitoringServiceLevelObjective, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewPatchSubresourceAction(monitoringservicelevelobjectivesResource, c.ns, name, pt, data, subresources...), &v1beta1.MonitoringServiceLevelObjective{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1beta1.MonitoringServiceLevelObjective), err
}
