// Copyright 2020 Google LLC
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//    http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

// ----------------------------------------------------------------------------
//
//     ***     AUTO GENERATED CODE    ***    AUTO GENERATED CODE     ***
//
// ----------------------------------------------------------------------------
//
//     This file is automatically generated by Config Connector and manual
//     changes will be clobbered when the file is regenerated.
//
// ----------------------------------------------------------------------------

// *** DISCLAIMER ***
// Config Connector's go-client for CRDs is currently in ALPHA, which means
// that future versions of the go-client may include breaking changes.
// Please try it out and give us feedback!

package v1alpha1

import (
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/clients/generated/apis/k8s/v1alpha1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

type DatasetaccessDataset struct {
	/* Immutable. The ID of the dataset containing this table. */
	DatasetId string `json:"datasetId"`

	/* Immutable. The ID of the project containing this table. */
	ProjectId string `json:"projectId"`
}

type DatasetaccessView struct {
	/* Immutable. The ID of the dataset containing this table. */
	DatasetId string `json:"datasetId"`

	/* Immutable. The ID of the project containing this table. */
	ProjectId string `json:"projectId"`

	/* Immutable. The ID of the table. The ID must contain only letters (a-z,
	A-Z), numbers (0-9), or underscores (_). The maximum length
	is 1,024 characters. */
	TableId string `json:"tableId"`
}

type BigQueryDatasetAccessSpec struct {
	/* Immutable. Grants all resources of particular types in a particular dataset read access to the current dataset. */
	// +optional
	Dataset *DatasetaccessDataset `json:"dataset,omitempty"`

	/* Immutable. A unique ID for this dataset, without the project name. The ID
	must contain only letters (a-z, A-Z), numbers (0-9), or
	underscores (_). The maximum length is 1,024 characters. */
	DatasetId string `json:"datasetId"`

	/* Immutable. A domain to grant access to. Any users signed in with the
	domain specified will be granted the specified access. */
	// +optional
	Domain *string `json:"domain,omitempty"`

	/* Immutable. An email address of a Google Group to grant access to. */
	// +optional
	GroupByEmail *string `json:"groupByEmail,omitempty"`

	/* Immutable. Some other type of member that appears in the IAM Policy but isn't a user,
	group, domain, or special group. For example: 'allUsers'. */
	// +optional
	IamMember *string `json:"iamMember,omitempty"`

	/* The project that this resource belongs to. */
	ProjectRef v1alpha1.ResourceRef `json:"projectRef"`

	/* Immutable. Optional. The routine of the resource. Used for creation and acquisition. When unset, the value of `metadata.name` is used as the default. */
	// +optional
	ResourceID *string `json:"resourceID,omitempty"`

	/* Immutable. Describes the rights granted to the user specified by the other
	member of the access object. Basic, predefined, and custom roles are
	supported. Predefined roles that have equivalent basic roles are
	swapped by the API to their basic counterparts, and will show a diff
	post-create. See
	[official docs](https://cloud.google.com/bigquery/docs/access-control). */
	// +optional
	Role *string `json:"role,omitempty"`

	/* Immutable. A special group to grant access to. Possible values include:


	* 'projectOwners': Owners of the enclosing project.


	* 'projectReaders': Readers of the enclosing project.


	* 'projectWriters': Writers of the enclosing project.


	* 'allAuthenticatedUsers': All authenticated BigQuery users. */
	// +optional
	SpecialGroup *string `json:"specialGroup,omitempty"`

	/* Immutable. An email address of a user to grant access to. For example:
	fred@example.com. */
	// +optional
	UserByEmail *string `json:"userByEmail,omitempty"`

	/* Immutable. A view from a different dataset to grant access to. Queries
	executed against that view will have read access to tables in
	this dataset. The role field is not required when this field is
	set. If that view is updated by any user, access to the view
	needs to be granted again via an update operation. */
	// +optional
	View *DatasetaccessView `json:"view,omitempty"`
}

type BigQueryDatasetAccessStatus struct {
	/* Conditions represent the latest available observations of the
	   BigQueryDatasetAccess's current state. */
	Conditions []v1alpha1.Condition `json:"conditions,omitempty"`
	/* If true, represents that that the iam_member in the config was translated to a different member type by the API, and is stored in state as a different member type. */
	// +optional
	ApiUpdatedMember *bool `json:"apiUpdatedMember,omitempty"`

	/* ObservedGeneration is the generation of the resource that was most recently observed by the Config Connector controller. If this is equal to metadata.generation, then that means that the current reported status reflects the most recent desired state of the resource. */
	// +optional
	ObservedGeneration *int `json:"observedGeneration,omitempty"`
}

// +genclient
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:resource:categories=gcp,shortName=gcpbigquerydatasetaccess;gcpbigquerydatasetaccesses
// +kubebuilder:subresource:status
// +kubebuilder:metadata:labels="cnrm.cloud.google.com/managed-by-kcc=true";"cnrm.cloud.google.com/stability-level=alpha";"cnrm.cloud.google.com/system=true";"cnrm.cloud.google.com/tf2crd=true"
// +kubebuilder:printcolumn:name="Age",JSONPath=".metadata.creationTimestamp",type="date"
// +kubebuilder:printcolumn:name="Ready",JSONPath=".status.conditions[?(@.type=='Ready')].status",type="string",description="When 'True', the most recent reconcile of the resource succeeded"
// +kubebuilder:printcolumn:name="Status",JSONPath=".status.conditions[?(@.type=='Ready')].reason",type="string",description="The reason for the value in 'Ready'"
// +kubebuilder:printcolumn:name="Status Age",JSONPath=".status.conditions[?(@.type=='Ready')].lastTransitionTime",type="date",description="The last transition time for the value in 'Status'"

// BigQueryDatasetAccess is the Schema for the bigquery API
// +k8s:openapi-gen=true
type BigQueryDatasetAccess struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   BigQueryDatasetAccessSpec   `json:"spec,omitempty"`
	Status BigQueryDatasetAccessStatus `json:"status,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// BigQueryDatasetAccessList contains a list of BigQueryDatasetAccess
type BigQueryDatasetAccessList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []BigQueryDatasetAccess `json:"items"`
}

func init() {
	SchemeBuilder.Register(&BigQueryDatasetAccess{}, &BigQueryDatasetAccessList{})
}
