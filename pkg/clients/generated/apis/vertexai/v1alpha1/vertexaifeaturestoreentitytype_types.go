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

type FeaturestoreentitytypeCategoricalThresholdConfig struct {
	/* Specify a threshold value that can trigger the alert. For categorical feature, the distribution distance is calculated by L-inifinity norm. Each feature must have a non-zero threshold if they need to be monitored. Otherwise no alert will be triggered for that feature. The default value is 0.3. */
	Value float64 `json:"value"`
}

type FeaturestoreentitytypeImportFeaturesAnalysis struct {
	/* Defines the baseline to do anomaly detection for feature values imported by each [entityTypes.importFeatureValues][] operation. The value must be one of the values below:
	* LATEST_STATS: Choose the later one statistics generated by either most recent snapshot analysis or previous import features analysis. If non of them exists, skip anomaly detection and only generate a statistics.
	* MOST_RECENT_SNAPSHOT_STATS: Use the statistics generated by the most recent snapshot analysis if exists.
	* PREVIOUS_IMPORT_FEATURES_STATS: Use the statistics generated by the previous import features analysis if exists. */
	// +optional
	AnomalyDetectionBaseline *string `json:"anomalyDetectionBaseline,omitempty"`

	/* Whether to enable / disable / inherite default hebavior for import features analysis. The value must be one of the values below:
	* DEFAULT: The default behavior of whether to enable the monitoring. EntityType-level config: disabled.
	* ENABLED: Explicitly enables import features analysis. EntityType-level config: by default enables import features analysis for all Features under it.
	* DISABLED: Explicitly disables import features analysis. EntityType-level config: by default disables import features analysis for all Features under it. */
	// +optional
	State *string `json:"state,omitempty"`
}

type FeaturestoreentitytypeMonitoringConfig struct {
	/* Threshold for categorical features of anomaly detection. This is shared by all types of Featurestore Monitoring for categorical features (i.e. Features with type (Feature.ValueType) BOOL or STRING). */
	// +optional
	CategoricalThresholdConfig *FeaturestoreentitytypeCategoricalThresholdConfig `json:"categoricalThresholdConfig,omitempty"`

	/* The config for ImportFeatures Analysis Based Feature Monitoring. */
	// +optional
	ImportFeaturesAnalysis *FeaturestoreentitytypeImportFeaturesAnalysis `json:"importFeaturesAnalysis,omitempty"`

	/* Threshold for numerical features of anomaly detection. This is shared by all objectives of Featurestore Monitoring for numerical features (i.e. Features with type (Feature.ValueType) DOUBLE or INT64). */
	// +optional
	NumericalThresholdConfig *FeaturestoreentitytypeNumericalThresholdConfig `json:"numericalThresholdConfig,omitempty"`

	/* The config for Snapshot Analysis Based Feature Monitoring. */
	// +optional
	SnapshotAnalysis *FeaturestoreentitytypeSnapshotAnalysis `json:"snapshotAnalysis,omitempty"`
}

type FeaturestoreentitytypeNumericalThresholdConfig struct {
	/* Specify a threshold value that can trigger the alert. For numerical feature, the distribution distance is calculated by Jensen–Shannon divergence. Each feature must have a non-zero threshold if they need to be monitored. Otherwise no alert will be triggered for that feature. The default value is 0.3. */
	Value float64 `json:"value"`
}

type FeaturestoreentitytypeSnapshotAnalysis struct {
	/* The monitoring schedule for snapshot analysis. For EntityType-level config: unset / disabled = true indicates disabled by default for Features under it; otherwise by default enable snapshot analysis monitoring with monitoringInterval for Features under it. */
	// +optional
	Disabled *bool `json:"disabled,omitempty"`

	/* DEPRECATED. `monitoring_interval` is deprecated and will be removed in a future release. Configuration of the snapshot analysis based monitoring pipeline running interval. The value is rolled up to full day.

	A duration in seconds with up to nine fractional digits, terminated by 's'. Example: "3.5s". */
	// +optional
	MonitoringInterval *string `json:"monitoringInterval,omitempty"`

	/* Configuration of the snapshot analysis based monitoring pipeline running interval. The value indicates number of days. The default value is 1.
	If both FeaturestoreMonitoringConfig.SnapshotAnalysis.monitoring_interval_days and [FeaturestoreMonitoringConfig.SnapshotAnalysis.monitoring_interval][] are set when creating/updating EntityTypes/Features, FeaturestoreMonitoringConfig.SnapshotAnalysis.monitoring_interval_days will be used. */
	// +optional
	MonitoringIntervalDays *int `json:"monitoringIntervalDays,omitempty"`

	/* Customized export features time window for snapshot analysis. Unit is one day. The default value is 21 days. Minimum value is 1 day. Maximum value is 4000 days. */
	// +optional
	StalenessDays *int `json:"stalenessDays,omitempty"`
}

type VertexAIFeaturestoreEntityTypeSpec struct {
	/* Optional. Description of the EntityType. */
	// +optional
	Description *string `json:"description,omitempty"`

	/* Immutable. The name of the Featurestore to use, in the format projects/{project}/locations/{location}/featurestores/{featurestore}. */
	Featurestore string `json:"featurestore"`

	/* The default monitoring configuration for all Features under this EntityType.

	If this is populated with [FeaturestoreMonitoringConfig.monitoring_interval] specified, snapshot analysis monitoring is enabled. Otherwise, snapshot analysis monitoring is disabled. */
	// +optional
	MonitoringConfig *FeaturestoreentitytypeMonitoringConfig `json:"monitoringConfig,omitempty"`

	/* Config for data retention policy in offline storage. TTL in days for feature values that will be stored in offline storage. The Feature Store offline storage periodically removes obsolete feature values older than offlineStorageTtlDays since the feature generation time. If unset (or explicitly set to 0), default to 4000 days TTL. */
	// +optional
	OfflineStorageTtlDays *int `json:"offlineStorageTtlDays,omitempty"`

	/* Immutable. Optional. The name of the resource. Used for creation and acquisition. When unset, the value of `metadata.name` is used as the default. */
	// +optional
	ResourceID *string `json:"resourceID,omitempty"`
}

type VertexAIFeaturestoreEntityTypeStatus struct {
	/* Conditions represent the latest available observations of the
	   VertexAIFeaturestoreEntityType's current state. */
	Conditions []v1alpha1.Condition `json:"conditions,omitempty"`
	/* The timestamp of when the featurestore was created in RFC3339 UTC "Zulu" format, with nanosecond resolution and up to nine fractional digits. */
	// +optional
	CreateTime *string `json:"createTime,omitempty"`

	/* Used to perform consistent read-modify-write updates. */
	// +optional
	Etag *string `json:"etag,omitempty"`

	/* ObservedGeneration is the generation of the resource that was most recently observed by the Config Connector controller. If this is equal to metadata.generation, then that means that the current reported status reflects the most recent desired state of the resource. */
	// +optional
	ObservedGeneration *int `json:"observedGeneration,omitempty"`

	/* The region of the EntityType. */
	// +optional
	Region *string `json:"region,omitempty"`

	/* The timestamp of when the featurestore was last updated in RFC3339 UTC "Zulu" format, with nanosecond resolution and up to nine fractional digits. */
	// +optional
	UpdateTime *string `json:"updateTime,omitempty"`
}

// +genclient
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:resource:categories=gcp,shortName=gcpvertexaifeaturestoreentitytype;gcpvertexaifeaturestoreentitytypes
// +kubebuilder:subresource:status
// +kubebuilder:metadata:labels="cnrm.cloud.google.com/managed-by-kcc=true";"cnrm.cloud.google.com/stability-level=alpha";"cnrm.cloud.google.com/system=true";"cnrm.cloud.google.com/tf2crd=true"
// +kubebuilder:printcolumn:name="Age",JSONPath=".metadata.creationTimestamp",type="date"
// +kubebuilder:printcolumn:name="Ready",JSONPath=".status.conditions[?(@.type=='Ready')].status",type="string",description="When 'True', the most recent reconcile of the resource succeeded"
// +kubebuilder:printcolumn:name="Status",JSONPath=".status.conditions[?(@.type=='Ready')].reason",type="string",description="The reason for the value in 'Ready'"
// +kubebuilder:printcolumn:name="Status Age",JSONPath=".status.conditions[?(@.type=='Ready')].lastTransitionTime",type="date",description="The last transition time for the value in 'Status'"

// VertexAIFeaturestoreEntityType is the Schema for the vertexai API
// +k8s:openapi-gen=true
type VertexAIFeaturestoreEntityType struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   VertexAIFeaturestoreEntityTypeSpec   `json:"spec,omitempty"`
	Status VertexAIFeaturestoreEntityTypeStatus `json:"status,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// VertexAIFeaturestoreEntityTypeList contains a list of VertexAIFeaturestoreEntityType
type VertexAIFeaturestoreEntityTypeList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []VertexAIFeaturestoreEntityType `json:"items"`
}

func init() {
	SchemeBuilder.Register(&VertexAIFeaturestoreEntityType{}, &VertexAIFeaturestoreEntityTypeList{})
}
