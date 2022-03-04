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

	v1beta1 "github.com/GoogleCloudPlatform/k8s-config-connector/pkg/clients/generated/apis/logging/v1beta1"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	labels "k8s.io/apimachinery/pkg/labels"
	schema "k8s.io/apimachinery/pkg/runtime/schema"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	testing "k8s.io/client-go/testing"
)

// FakeLoggingLogExclusions implements LoggingLogExclusionInterface
type FakeLoggingLogExclusions struct {
	Fake *FakeLoggingV1beta1
	ns   string
}

var logginglogexclusionsResource = schema.GroupVersionResource{Group: "logging.cnrm.cloud.google.com", Version: "v1beta1", Resource: "logginglogexclusions"}

var logginglogexclusionsKind = schema.GroupVersionKind{Group: "logging.cnrm.cloud.google.com", Version: "v1beta1", Kind: "LoggingLogExclusion"}

// Get takes name of the loggingLogExclusion, and returns the corresponding loggingLogExclusion object, and an error if there is any.
func (c *FakeLoggingLogExclusions) Get(ctx context.Context, name string, options v1.GetOptions) (result *v1beta1.LoggingLogExclusion, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewGetAction(logginglogexclusionsResource, c.ns, name), &v1beta1.LoggingLogExclusion{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1beta1.LoggingLogExclusion), err
}

// List takes label and field selectors, and returns the list of LoggingLogExclusions that match those selectors.
func (c *FakeLoggingLogExclusions) List(ctx context.Context, opts v1.ListOptions) (result *v1beta1.LoggingLogExclusionList, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewListAction(logginglogexclusionsResource, logginglogexclusionsKind, c.ns, opts), &v1beta1.LoggingLogExclusionList{})

	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &v1beta1.LoggingLogExclusionList{ListMeta: obj.(*v1beta1.LoggingLogExclusionList).ListMeta}
	for _, item := range obj.(*v1beta1.LoggingLogExclusionList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested loggingLogExclusions.
func (c *FakeLoggingLogExclusions) Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewWatchAction(logginglogexclusionsResource, c.ns, opts))

}

// Create takes the representation of a loggingLogExclusion and creates it.  Returns the server's representation of the loggingLogExclusion, and an error, if there is any.
func (c *FakeLoggingLogExclusions) Create(ctx context.Context, loggingLogExclusion *v1beta1.LoggingLogExclusion, opts v1.CreateOptions) (result *v1beta1.LoggingLogExclusion, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewCreateAction(logginglogexclusionsResource, c.ns, loggingLogExclusion), &v1beta1.LoggingLogExclusion{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1beta1.LoggingLogExclusion), err
}

// Update takes the representation of a loggingLogExclusion and updates it. Returns the server's representation of the loggingLogExclusion, and an error, if there is any.
func (c *FakeLoggingLogExclusions) Update(ctx context.Context, loggingLogExclusion *v1beta1.LoggingLogExclusion, opts v1.UpdateOptions) (result *v1beta1.LoggingLogExclusion, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateAction(logginglogexclusionsResource, c.ns, loggingLogExclusion), &v1beta1.LoggingLogExclusion{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1beta1.LoggingLogExclusion), err
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
func (c *FakeLoggingLogExclusions) UpdateStatus(ctx context.Context, loggingLogExclusion *v1beta1.LoggingLogExclusion, opts v1.UpdateOptions) (*v1beta1.LoggingLogExclusion, error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateSubresourceAction(logginglogexclusionsResource, "status", c.ns, loggingLogExclusion), &v1beta1.LoggingLogExclusion{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1beta1.LoggingLogExclusion), err
}

// Delete takes name of the loggingLogExclusion and deletes it. Returns an error if one occurs.
func (c *FakeLoggingLogExclusions) Delete(ctx context.Context, name string, opts v1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewDeleteActionWithOptions(logginglogexclusionsResource, c.ns, name, opts), &v1beta1.LoggingLogExclusion{})

	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakeLoggingLogExclusions) DeleteCollection(ctx context.Context, opts v1.DeleteOptions, listOpts v1.ListOptions) error {
	action := testing.NewDeleteCollectionAction(logginglogexclusionsResource, c.ns, listOpts)

	_, err := c.Fake.Invokes(action, &v1beta1.LoggingLogExclusionList{})
	return err
}

// Patch applies the patch and returns the patched loggingLogExclusion.
func (c *FakeLoggingLogExclusions) Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string) (result *v1beta1.LoggingLogExclusion, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewPatchSubresourceAction(logginglogexclusionsResource, c.ns, name, pt, data, subresources...), &v1beta1.LoggingLogExclusion{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1beta1.LoggingLogExclusion), err
}
