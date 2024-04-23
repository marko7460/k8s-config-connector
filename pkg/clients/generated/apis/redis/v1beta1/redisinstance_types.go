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

type InstanceMaintenancePolicy struct {
	/* Output only. The time when the policy was created.
	A timestamp in RFC3339 UTC "Zulu" format, with nanosecond
	resolution and up to nine fractional digits. */
	// +optional
	CreateTime *string `json:"createTime,omitempty"`

	/* Optional. Description of what this policy is for.
	Create/Update methods return INVALID_ARGUMENT if the
	length is greater than 512. */
	// +optional
	Description *string `json:"description,omitempty"`

	/* Output only. The time when the policy was last updated.
	A timestamp in RFC3339 UTC "Zulu" format, with nanosecond
	resolution and up to nine fractional digits. */
	// +optional
	UpdateTime *string `json:"updateTime,omitempty"`

	/* Optional. Maintenance window that is applied to resources covered by this policy.
	Minimum 1. For the current version, the maximum number
	of weekly_window is expected to be one. */
	// +optional
	WeeklyMaintenanceWindow []InstanceWeeklyMaintenanceWindow `json:"weeklyMaintenanceWindow,omitempty"`
}

type InstanceMaintenanceSchedule struct {
	/* Output only. The end time of any upcoming scheduled maintenance for this instance.
	A timestamp in RFC3339 UTC "Zulu" format, with nanosecond
	resolution and up to nine fractional digits. */
	// +optional
	EndTime *string `json:"endTime,omitempty"`

	/* Output only. The deadline that the maintenance schedule start time
	can not go beyond, including reschedule.
	A timestamp in RFC3339 UTC "Zulu" format, with nanosecond
	resolution and up to nine fractional digits. */
	// +optional
	ScheduleDeadlineTime *string `json:"scheduleDeadlineTime,omitempty"`

	/* Output only. The start time of any upcoming scheduled maintenance for this instance.
	A timestamp in RFC3339 UTC "Zulu" format, with nanosecond
	resolution and up to nine fractional digits. */
	// +optional
	StartTime *string `json:"startTime,omitempty"`
}

type InstancePersistenceConfig struct {
	/* Optional. Controls whether Persistence features are enabled. If not provided, the existing value will be used.

	- DISABLED: 	Persistence is disabled for the instance, and any existing snapshots are deleted.
	- RDB: RDB based Persistence is enabled. Possible values: ["DISABLED", "RDB"]. */
	// +optional
	PersistenceMode *string `json:"persistenceMode,omitempty"`

	/* Output only. The next time that a snapshot attempt is scheduled to occur.
	A timestamp in RFC3339 UTC "Zulu" format, with nanosecond resolution and up
	to nine fractional digits.
	Examples: "2014-10-02T15:01:23Z" and "2014-10-02T15:01:23.045123456Z". */
	// +optional
	RdbNextSnapshotTime *string `json:"rdbNextSnapshotTime,omitempty"`

	/* Optional. Available snapshot periods for scheduling.

	- ONE_HOUR:	Snapshot every 1 hour.
	- SIX_HOURS:	Snapshot every 6 hours.
	- TWELVE_HOURS:	Snapshot every 12 hours.
	- TWENTY_FOUR_HOURS:	Snapshot every 24 hours. Possible values: ["ONE_HOUR", "SIX_HOURS", "TWELVE_HOURS", "TWENTY_FOUR_HOURS"]. */
	// +optional
	RdbSnapshotPeriod *string `json:"rdbSnapshotPeriod,omitempty"`

	/* Optional. Date and time that the first snapshot was/will be attempted,
	and to which future snapshots will be aligned. If not provided,
	the current time will be used.
	A timestamp in RFC3339 UTC "Zulu" format, with nanosecond resolution
	and up to nine fractional digits.
	Examples: "2014-10-02T15:01:23Z" and "2014-10-02T15:01:23.045123456Z". */
	// +optional
	RdbSnapshotStartTime *string `json:"rdbSnapshotStartTime,omitempty"`
}

type InstanceStartTime struct {
	/* Hours of day in 24 hour format. Should be from 0 to 23.
	An API may choose to allow the value "24:00:00" for scenarios like business closing time. */
	// +optional
	Hours *int `json:"hours,omitempty"`

	/* Minutes of hour of day. Must be from 0 to 59. */
	// +optional
	Minutes *int `json:"minutes,omitempty"`

	/* Fractions of seconds in nanoseconds. Must be from 0 to 999,999,999. */
	// +optional
	Nanos *int `json:"nanos,omitempty"`

	/* Seconds of minutes of the time. Must normally be from 0 to 59.
	An API may allow the value 60 if it allows leap-seconds. */
	// +optional
	Seconds *int `json:"seconds,omitempty"`
}

type InstanceWeeklyMaintenanceWindow struct {
	/* Required. The day of week that maintenance updates occur.

	- DAY_OF_WEEK_UNSPECIFIED: The day of the week is unspecified.
	- MONDAY: Monday
	- TUESDAY: Tuesday
	- WEDNESDAY: Wednesday
	- THURSDAY: Thursday
	- FRIDAY: Friday
	- SATURDAY: Saturday
	- SUNDAY: Sunday Possible values: ["DAY_OF_WEEK_UNSPECIFIED", "MONDAY", "TUESDAY", "WEDNESDAY", "THURSDAY", "FRIDAY", "SATURDAY", "SUNDAY"]. */
	Day string `json:"day"`

	/* Output only. Duration of the maintenance window.
	The current window is fixed at 1 hour.
	A duration in seconds with up to nine fractional digits,
	terminated by 's'. Example: "3.5s". */
	// +optional
	Duration *string `json:"duration,omitempty"`

	/* Required. Start time of the window in UTC time. */
	StartTime InstanceStartTime `json:"startTime"`
}

type RedisInstanceSpec struct {
	/* Immutable. Only applicable to STANDARD_HA tier which protects the instance
	against zonal failures by provisioning it across two zones.
	If provided, it must be a different zone from the one provided in
	[locationId]. */
	// +optional
	AlternativeLocationId *string `json:"alternativeLocationId,omitempty"`

	/* Optional. Indicates whether OSS Redis AUTH is enabled for the
	instance. If set to "true" AUTH is enabled on the instance.
	Default value is "false" meaning AUTH is disabled. */
	// +optional
	AuthEnabled *bool `json:"authEnabled,omitempty"`

	/* Output only. AUTH String set on the instance. This field will only be populated if auth_enabled is true. */
	// +optional
	AuthString *string `json:"authString,omitempty"`

	/* The network to which the instance is connected. If left
	unspecified, the default network will be used. */
	// +optional
	AuthorizedNetworkRef *v1alpha1.ResourceRef `json:"authorizedNetworkRef,omitempty"`

	/* Immutable. The connection mode of the Redis instance. Default value: "DIRECT_PEERING" Possible values: ["DIRECT_PEERING", "PRIVATE_SERVICE_ACCESS"]. */
	// +optional
	ConnectMode *string `json:"connectMode,omitempty"`

	/* Immutable. Optional. The KMS key reference that you want to use to
	encrypt the data at rest for this Redis instance. If this is
	provided, CMEK is enabled. */
	// +optional
	CustomerManagedKeyRef *v1alpha1.ResourceRef `json:"customerManagedKeyRef,omitempty"`

	/* An arbitrary and optional user-provided name for the instance. */
	// +optional
	DisplayName *string `json:"displayName,omitempty"`

	/* Immutable. The zone where the instance will be provisioned. If not provided,
	the service will choose a zone for the instance. For STANDARD_HA tier,
	instances will be created across two zones for protection against
	zonal failures. If [alternativeLocationId] is also provided, it must
	be different from [locationId]. */
	// +optional
	LocationId *string `json:"locationId,omitempty"`

	/* Maintenance policy for an instance. */
	// +optional
	MaintenancePolicy *InstanceMaintenancePolicy `json:"maintenancePolicy,omitempty"`

	/* Upcoming maintenance schedule. */
	// +optional
	MaintenanceSchedule []InstanceMaintenanceSchedule `json:"maintenanceSchedule,omitempty"`

	/* Redis memory size in GiB. */
	MemorySizeGb int `json:"memorySizeGb"`

	/* Persistence configuration for an instance. */
	// +optional
	PersistenceConfig *InstancePersistenceConfig `json:"persistenceConfig,omitempty"`

	/* Optional. Read replica mode. Can only be specified when trying to create the instance.
	If not set, Memorystore Redis backend will default to READ_REPLICAS_DISABLED.
	- READ_REPLICAS_DISABLED: If disabled, read endpoint will not be provided and the
	instance cannot scale up or down the number of replicas.
	- READ_REPLICAS_ENABLED: If enabled, read endpoint will be provided and the instance
	can scale up and down the number of replicas. Possible values: ["READ_REPLICAS_DISABLED", "READ_REPLICAS_ENABLED"]. */
	// +optional
	ReadReplicasMode *string `json:"readReplicasMode,omitempty"`

	/* Redis configuration parameters, according to http://redis.io/topics/config.
	Please check Memorystore documentation for the list of supported parameters:
	https://cloud.google.com/memorystore/docs/redis/reference/rest/v1/projects.locations.instances#Instance.FIELDS.redis_configs. */
	// +optional
	RedisConfigs map[string]string `json:"redisConfigs,omitempty"`

	/* The version of Redis software. If not provided, latest supported
	version will be used. Please check the API documentation linked
	at the top for the latest valid values. */
	// +optional
	RedisVersion *string `json:"redisVersion,omitempty"`

	/* Immutable. The name of the Redis region of the instance. */
	Region string `json:"region"`

	/* Optional. The number of replica nodes. The valid range for the Standard Tier with
	read replicas enabled is [1-5] and defaults to 2. If read replicas are not enabled
	for a Standard Tier instance, the only valid value is 1 and the default is 1.
	The valid value for basic tier is 0 and the default is also 0. */
	// +optional
	ReplicaCount *int `json:"replicaCount,omitempty"`

	/* Immutable. The CIDR range of internal addresses that are reserved for this
	instance. If not provided, the service will choose an unused /29
	block, for example, 10.0.0.0/29 or 192.168.0.0/29. Ranges must be
	unique and non-overlapping with existing subnets in an authorized
	network. */
	// +optional
	ReservedIpRange *string `json:"reservedIpRange,omitempty"`

	/* Immutable. Optional. The name of the resource. Used for creation and acquisition. When unset, the value of `metadata.name` is used as the default. */
	// +optional
	ResourceID *string `json:"resourceID,omitempty"`

	/* Optional. Additional IP range for node placement. Required when enabling read replicas on
	an existing instance. For DIRECT_PEERING mode value must be a CIDR range of size /28, or
	"auto". For PRIVATE_SERVICE_ACCESS mode value must be the name of an allocated address
	range associated with the private service access connection, or "auto". */
	// +optional
	SecondaryIpRange *string `json:"secondaryIpRange,omitempty"`

	/* Immutable. The service tier of the instance. Must be one of these values:

	- BASIC: standalone instance
	- STANDARD_HA: highly available primary/replica instances Default value: "BASIC" Possible values: ["BASIC", "STANDARD_HA"]. */
	// +optional
	Tier *string `json:"tier,omitempty"`

	/* Immutable. The TLS mode of the Redis instance, If not provided, TLS is disabled for the instance.

	- SERVER_AUTHENTICATION: Client to Server traffic encryption enabled with server authentication Default value: "DISABLED" Possible values: ["SERVER_AUTHENTICATION", "DISABLED"]. */
	// +optional
	TransitEncryptionMode *string `json:"transitEncryptionMode,omitempty"`
}

type InstanceMaintenanceScheduleStatus struct {
	/* Output only. The end time of any upcoming scheduled maintenance for this instance.
	A timestamp in RFC3339 UTC "Zulu" format, with nanosecond
	resolution and up to nine fractional digits. */
	// +optional
	EndTime *string `json:"endTime,omitempty"`

	/* Output only. The deadline that the maintenance schedule start time
	can not go beyond, including reschedule.
	A timestamp in RFC3339 UTC "Zulu" format, with nanosecond
	resolution and up to nine fractional digits. */
	// +optional
	ScheduleDeadlineTime *string `json:"scheduleDeadlineTime,omitempty"`

	/* Output only. The start time of any upcoming scheduled maintenance for this instance.
	A timestamp in RFC3339 UTC "Zulu" format, with nanosecond
	resolution and up to nine fractional digits. */
	// +optional
	StartTime *string `json:"startTime,omitempty"`
}

type InstanceNodesStatus struct {
	/* Node identifying string. e.g. 'node-0', 'node-1'. */
	// +optional
	Id *string `json:"id,omitempty"`

	/* Location of the node. */
	// +optional
	Zone *string `json:"zone,omitempty"`
}

type InstanceObservedStateStatus struct {
	/* Output only. AUTH String set on the instance. This field will only be populated if auth_enabled is true. */
	// +optional
	AuthString *string `json:"authString,omitempty"`
}

type InstanceServerCaCertsStatus struct {
	/* The certificate data in PEM format. */
	// +optional
	Cert *string `json:"cert,omitempty"`

	/* The time when the certificate was created. */
	// +optional
	CreateTime *string `json:"createTime,omitempty"`

	/* The time when the certificate expires. */
	// +optional
	ExpireTime *string `json:"expireTime,omitempty"`

	/* Serial number, as extracted from the certificate. */
	// +optional
	SerialNumber *string `json:"serialNumber,omitempty"`

	/* Sha1 Fingerprint of the certificate. */
	// +optional
	Sha1Fingerprint *string `json:"sha1Fingerprint,omitempty"`
}

type RedisInstanceStatus struct {
	/* Conditions represent the latest available observations of the
	   RedisInstance's current state. */
	Conditions []v1alpha1.Condition `json:"conditions,omitempty"`
	/* The time the instance was created in RFC3339 UTC "Zulu" format,
	accurate to nanoseconds. */
	// +optional
	CreateTime *string `json:"createTime,omitempty"`

	/* The current zone where the Redis endpoint is placed.
	For Basic Tier instances, this will always be the same as the
	[locationId] provided by the user at creation time. For Standard Tier
	instances, this can be either [locationId] or [alternativeLocationId]
	and can change after a failover event. */
	// +optional
	CurrentLocationId *string `json:"currentLocationId,omitempty"`

	/* Hostname or IP address of the exposed Redis endpoint used by clients
	to connect to the service. */
	// +optional
	Host *string `json:"host,omitempty"`

	/* Upcoming maintenance schedule. */
	// +optional
	MaintenanceSchedule []InstanceMaintenanceScheduleStatus `json:"maintenanceSchedule,omitempty"`

	/* Output only. Info per node. */
	// +optional
	Nodes []InstanceNodesStatus `json:"nodes,omitempty"`

	/* ObservedGeneration is the generation of the resource that was most recently observed by the Config Connector controller. If this is equal to metadata.generation, then that means that the current reported status reflects the most recent desired state of the resource. */
	// +optional
	ObservedGeneration *int `json:"observedGeneration,omitempty"`

	/* The observed state of the underlying GCP resource. */
	// +optional
	ObservedState *InstanceObservedStateStatus `json:"observedState,omitempty"`

	/* Output only. Cloud IAM identity used by import / export operations
	to transfer data to/from Cloud Storage. Format is "serviceAccount:".
	The value may change over time for a given instance so should be
	checked before each import/export operation. */
	// +optional
	PersistenceIamIdentity *string `json:"persistenceIamIdentity,omitempty"`

	/* The port number of the exposed Redis endpoint. */
	// +optional
	Port *int `json:"port,omitempty"`

	/* Output only. Hostname or IP address of the exposed readonly Redis endpoint. Standard tier only.
	Targets all healthy replica nodes in instance. Replication is asynchronous and replica nodes
	will exhibit some lag behind the primary. Write requests must target 'host'. */
	// +optional
	ReadEndpoint *string `json:"readEndpoint,omitempty"`

	/* Output only. The port number of the exposed readonly redis endpoint. Standard tier only.
	Write requests should target 'port'. */
	// +optional
	ReadEndpointPort *int `json:"readEndpointPort,omitempty"`

	/* List of server CA certificates for the instance. */
	// +optional
	ServerCaCerts []InstanceServerCaCertsStatus `json:"serverCaCerts,omitempty"`
}

// +genclient
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:resource:categories=gcp,shortName=gcpredisinstance;gcpredisinstances
// +kubebuilder:subresource:status
// +kubebuilder:metadata:labels="cnrm.cloud.google.com/managed-by-kcc=true";"cnrm.cloud.google.com/stability-level=stable";"cnrm.cloud.google.com/system=true";"cnrm.cloud.google.com/tf2crd=true"
// +kubebuilder:printcolumn:name="Age",JSONPath=".metadata.creationTimestamp",type="date"
// +kubebuilder:printcolumn:name="Ready",JSONPath=".status.conditions[?(@.type=='Ready')].status",type="string",description="When 'True', the most recent reconcile of the resource succeeded"
// +kubebuilder:printcolumn:name="Status",JSONPath=".status.conditions[?(@.type=='Ready')].reason",type="string",description="The reason for the value in 'Ready'"
// +kubebuilder:printcolumn:name="Status Age",JSONPath=".status.conditions[?(@.type=='Ready')].lastTransitionTime",type="date",description="The last transition time for the value in 'Status'"

// RedisInstance is the Schema for the redis API
// +k8s:openapi-gen=true
type RedisInstance struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   RedisInstanceSpec   `json:"spec,omitempty"`
	Status RedisInstanceStatus `json:"status,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// RedisInstanceList contains a list of RedisInstance
type RedisInstanceList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []RedisInstance `json:"items"`
}

func init() {
	SchemeBuilder.Register(&RedisInstance{}, &RedisInstanceList{})
}
