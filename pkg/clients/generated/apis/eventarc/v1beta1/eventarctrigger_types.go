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

package v1beta1

import (
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/clients/generated/apis/k8s/v1alpha1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

type TriggerCloudRunService struct {
	/* Optional. The relative path on the Cloud Run service the events should be sent to. The value must conform to the definition of URI path segment (section 3.3 of RFC2396). Examples: "/route", "route", "route/subroute". */
	// +optional
	Path *string `json:"path,omitempty"`

	/* Required. The region the Cloud Run service is deployed in. */
	Region string `json:"region"`

	ServiceRef v1alpha1.ResourceRef `json:"serviceRef"`
}

type TriggerDestination struct {
	// +optional
	CloudFunctionRef *v1alpha1.ResourceRef `json:"cloudFunctionRef,omitempty"`

	/* Cloud Run fully-managed service that receives the events. The service should be running in the same project of the trigger. */
	// +optional
	CloudRunService *TriggerCloudRunService `json:"cloudRunService,omitempty"`

	/* A GKE service capable of receiving events. The service should be running in the same project as the trigger. */
	// +optional
	Gke *TriggerGke `json:"gke,omitempty"`

	/* An HTTP endpoint destination described by an URI. */
	// +optional
	HttpEndpoint *TriggerHttpEndpoint `json:"httpEndpoint,omitempty"`

	/* Optional. Network config is used to configure how Eventarc resolves and connect to a destination. This should only be used with HttpEndpoint destination type. */
	// +optional
	NetworkConfig *TriggerNetworkConfig `json:"networkConfig,omitempty"`

	// +optional
	WorkflowRef *v1alpha1.ResourceRef `json:"workflowRef,omitempty"`
}

type TriggerGke struct {
	ClusterRef v1alpha1.ResourceRef `json:"clusterRef"`

	/* Required. The name of the Google Compute Engine in which the cluster resides, which can either be compute zone (for example, us-central1-a) for the zonal clusters or region (for example, us-central1) for regional clusters. */
	Location string `json:"location"`

	/* Required. The namespace the GKE service is running in. */
	Namespace string `json:"namespace"`

	/* Optional. The relative path on the GKE service the events should be sent to. The value must conform to the definition of a URI path segment (section 3.3 of RFC2396). Examples: "/route", "route", "route/subroute". */
	// +optional
	Path *string `json:"path,omitempty"`

	/* Required. Name of the GKE service. */
	Service string `json:"service"`
}

type TriggerHttpEndpoint struct {
	/* Required. The URI of the HTTP enpdoint. The value must be a RFC2396 URI string. Examples: `http://10.10.10.8:80/route`, `http://svc.us-central1.p.local:8080/`. Only HTTP and HTTPS protocols are supported. The host can be either a static IP addressable from the VPC specified by the network config, or an internal DNS hostname of the service resolvable via Cloud DNS. */
	Uri string `json:"uri"`
}

type TriggerMatchingCriteria struct {
	/* Required. The name of a CloudEvents attribute. Currently, only a subset of attributes are supported for filtering. All triggers MUST provide a filter for the 'type' attribute. */
	Attribute string `json:"attribute"`

	/* Optional. The operator used for matching the events with the value of the filter. If not specified, only events that have an exact key-value pair specified in the filter are matched. The only allowed value is `match-path-pattern`. */
	// +optional
	Operator *string `json:"operator,omitempty"`

	/* Required. The value for the attribute. See https://cloud.google.com/eventarc/docs/creating-triggers#trigger-gcloud for available values. */
	Value string `json:"value"`
}

type TriggerNetworkConfig struct {
	NetworkAttachmentRef v1alpha1.ResourceRef `json:"networkAttachmentRef"`
}

type TriggerPubsub struct {
	/* Immutable. */
	// +optional
	TopicRef *v1alpha1.ResourceRef `json:"topicRef,omitempty"`
}

type TriggerTransport struct {
	/* Immutable. The Pub/Sub topic and subscription used by Eventarc as delivery intermediary. */
	// +optional
	Pubsub *TriggerPubsub `json:"pubsub,omitempty"`
}

type EventarcTriggerSpec struct {
	/* Immutable. */
	// +optional
	ChannelRef *v1alpha1.ResourceRef `json:"channelRef,omitempty"`

	/* Required. Destination specifies where the events should be sent to. */
	Destination TriggerDestination `json:"destination"`

	/* Optional. EventDataContentType specifies the type of payload in MIME format that is expected from the CloudEvent data field. This is set to `application/json` if the value is not defined. */
	// +optional
	EventDataContentType *string `json:"eventDataContentType,omitempty"`

	/* Immutable. The location for the resource */
	Location string `json:"location"`

	/* Required. null The list of filters that applies to event attributes. Only events that match all the provided filters will be sent to the destination. */
	MatchingCriteria []TriggerMatchingCriteria `json:"matchingCriteria"`

	/* Immutable. The Project that this resource belongs to. */
	ProjectRef v1alpha1.ResourceRef `json:"projectRef"`

	/* Immutable. Optional. The name of the resource. Used for creation and acquisition. When unset, the value of `metadata.name` is used as the default. */
	// +optional
	ResourceID *string `json:"resourceID,omitempty"`

	// +optional
	ServiceAccountRef *v1alpha1.ResourceRef `json:"serviceAccountRef,omitempty"`

	/* Immutable. Optional. In order to deliver messages, Eventarc may use other GCP products as transport intermediary. This field contains a reference to that transport intermediary. This information can be used for debugging purposes. */
	// +optional
	Transport *TriggerTransport `json:"transport,omitempty"`
}

type TriggerPubsubStatus struct {
	/* Output only. The name of the Pub/Sub subscription created and managed by Eventarc system as a transport for the event delivery. Format: `projects/{PROJECT_ID}/subscriptions/{SUBSCRIPTION_NAME}`. */
	// +optional
	Subscription *string `json:"subscription,omitempty"`
}

type TriggerTransportStatus struct {
	// +optional
	Pubsub *TriggerPubsubStatus `json:"pubsub,omitempty"`
}

type EventarcTriggerStatus struct {
	/* Conditions represent the latest available observations of the
	   EventarcTrigger's current state. */
	Conditions []v1alpha1.Condition `json:"conditions,omitempty"`
	/* Output only. The creation time. */
	// +optional
	CreateTime *string `json:"createTime,omitempty"`

	/* Output only. This checksum is computed by the server based on the value of other fields, and may be sent only on create requests to ensure the client has an up-to-date value before proceeding. */
	// +optional
	Etag *string `json:"etag,omitempty"`

	/* ObservedGeneration is the generation of the resource that was most recently observed by the Config Connector controller. If this is equal to metadata.generation, then that means that the current reported status reflects the most recent desired state of the resource. */
	// +optional
	ObservedGeneration *int `json:"observedGeneration,omitempty"`

	/* Output only. The reason(s) why a trigger is in FAILED state. */
	// +optional
	ResourceConditions map[string]string `json:"resourceConditions,omitempty"`

	// +optional
	Transport *TriggerTransportStatus `json:"transport,omitempty"`

	/* Output only. Server assigned unique identifier for the trigger. The value is a UUID4 string and guaranteed to remain unchanged until the resource is deleted. */
	// +optional
	Uid *string `json:"uid,omitempty"`

	/* Output only. The last-modified time. */
	// +optional
	UpdateTime *string `json:"updateTime,omitempty"`
}

// +genclient
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:resource:categories=gcp,shortName=gcpeventarctrigger;gcpeventarctriggers
// +kubebuilder:subresource:status
// +kubebuilder:metadata:labels="cnrm.cloud.google.com/dcl2crd=true";"cnrm.cloud.google.com/managed-by-kcc=true";"cnrm.cloud.google.com/stability-level=stable";"cnrm.cloud.google.com/system=true"
// +kubebuilder:printcolumn:name="Age",JSONPath=".metadata.creationTimestamp",type="date"
// +kubebuilder:printcolumn:name="Ready",JSONPath=".status.conditions[?(@.type=='Ready')].status",type="string",description="When 'True', the most recent reconcile of the resource succeeded"
// +kubebuilder:printcolumn:name="Status",JSONPath=".status.conditions[?(@.type=='Ready')].reason",type="string",description="The reason for the value in 'Ready'"
// +kubebuilder:printcolumn:name="Status Age",JSONPath=".status.conditions[?(@.type=='Ready')].lastTransitionTime",type="date",description="The last transition time for the value in 'Status'"

// EventarcTrigger is the Schema for the eventarc API
// +k8s:openapi-gen=true
type EventarcTrigger struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   EventarcTriggerSpec   `json:"spec,omitempty"`
	Status EventarcTriggerStatus `json:"status,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// EventarcTriggerList contains a list of EventarcTrigger
type EventarcTriggerList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []EventarcTrigger `json:"items"`
}

func init() {
	SchemeBuilder.Register(&EventarcTrigger{}, &EventarcTriggerList{})
}
