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

	v1beta1 "github.com/GoogleCloudPlatform/k8s-config-connector/pkg/clients/generated/apis/artifactregistry/v1beta1"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	labels "k8s.io/apimachinery/pkg/labels"
	schema "k8s.io/apimachinery/pkg/runtime/schema"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	testing "k8s.io/client-go/testing"
)

// FakeArtifactRegistryRepositories implements ArtifactRegistryRepositoryInterface
type FakeArtifactRegistryRepositories struct {
	Fake *FakeArtifactregistryV1beta1
	ns   string
}

var artifactregistryrepositoriesResource = schema.GroupVersionResource{Group: "artifactregistry.cnrm.cloud.google.com", Version: "v1beta1", Resource: "artifactregistryrepositories"}

var artifactregistryrepositoriesKind = schema.GroupVersionKind{Group: "artifactregistry.cnrm.cloud.google.com", Version: "v1beta1", Kind: "ArtifactRegistryRepository"}

// Get takes name of the artifactRegistryRepository, and returns the corresponding artifactRegistryRepository object, and an error if there is any.
func (c *FakeArtifactRegistryRepositories) Get(ctx context.Context, name string, options v1.GetOptions) (result *v1beta1.ArtifactRegistryRepository, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewGetAction(artifactregistryrepositoriesResource, c.ns, name), &v1beta1.ArtifactRegistryRepository{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1beta1.ArtifactRegistryRepository), err
}

// List takes label and field selectors, and returns the list of ArtifactRegistryRepositories that match those selectors.
func (c *FakeArtifactRegistryRepositories) List(ctx context.Context, opts v1.ListOptions) (result *v1beta1.ArtifactRegistryRepositoryList, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewListAction(artifactregistryrepositoriesResource, artifactregistryrepositoriesKind, c.ns, opts), &v1beta1.ArtifactRegistryRepositoryList{})

	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &v1beta1.ArtifactRegistryRepositoryList{ListMeta: obj.(*v1beta1.ArtifactRegistryRepositoryList).ListMeta}
	for _, item := range obj.(*v1beta1.ArtifactRegistryRepositoryList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested artifactRegistryRepositories.
func (c *FakeArtifactRegistryRepositories) Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewWatchAction(artifactregistryrepositoriesResource, c.ns, opts))

}

// Create takes the representation of a artifactRegistryRepository and creates it.  Returns the server's representation of the artifactRegistryRepository, and an error, if there is any.
func (c *FakeArtifactRegistryRepositories) Create(ctx context.Context, artifactRegistryRepository *v1beta1.ArtifactRegistryRepository, opts v1.CreateOptions) (result *v1beta1.ArtifactRegistryRepository, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewCreateAction(artifactregistryrepositoriesResource, c.ns, artifactRegistryRepository), &v1beta1.ArtifactRegistryRepository{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1beta1.ArtifactRegistryRepository), err
}

// Update takes the representation of a artifactRegistryRepository and updates it. Returns the server's representation of the artifactRegistryRepository, and an error, if there is any.
func (c *FakeArtifactRegistryRepositories) Update(ctx context.Context, artifactRegistryRepository *v1beta1.ArtifactRegistryRepository, opts v1.UpdateOptions) (result *v1beta1.ArtifactRegistryRepository, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateAction(artifactregistryrepositoriesResource, c.ns, artifactRegistryRepository), &v1beta1.ArtifactRegistryRepository{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1beta1.ArtifactRegistryRepository), err
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
func (c *FakeArtifactRegistryRepositories) UpdateStatus(ctx context.Context, artifactRegistryRepository *v1beta1.ArtifactRegistryRepository, opts v1.UpdateOptions) (*v1beta1.ArtifactRegistryRepository, error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateSubresourceAction(artifactregistryrepositoriesResource, "status", c.ns, artifactRegistryRepository), &v1beta1.ArtifactRegistryRepository{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1beta1.ArtifactRegistryRepository), err
}

// Delete takes name of the artifactRegistryRepository and deletes it. Returns an error if one occurs.
func (c *FakeArtifactRegistryRepositories) Delete(ctx context.Context, name string, opts v1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewDeleteActionWithOptions(artifactregistryrepositoriesResource, c.ns, name, opts), &v1beta1.ArtifactRegistryRepository{})

	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakeArtifactRegistryRepositories) DeleteCollection(ctx context.Context, opts v1.DeleteOptions, listOpts v1.ListOptions) error {
	action := testing.NewDeleteCollectionAction(artifactregistryrepositoriesResource, c.ns, listOpts)

	_, err := c.Fake.Invokes(action, &v1beta1.ArtifactRegistryRepositoryList{})
	return err
}

// Patch applies the patch and returns the patched artifactRegistryRepository.
func (c *FakeArtifactRegistryRepositories) Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string) (result *v1beta1.ArtifactRegistryRepository, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewPatchSubresourceAction(artifactregistryrepositoriesResource, c.ns, name, pt, data, subresources...), &v1beta1.ArtifactRegistryRepository{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1beta1.ArtifactRegistryRepository), err
}
