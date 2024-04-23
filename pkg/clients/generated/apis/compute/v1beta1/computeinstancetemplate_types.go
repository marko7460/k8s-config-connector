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

type InstancetemplateAccessConfig struct {
	// +optional
	NatIpRef *v1alpha1.ResourceRef `json:"natIpRef,omitempty"`

	/* Immutable. The networking tier used for configuring this instance template. This field can take the following values: PREMIUM, STANDARD, FIXED_STANDARD. If this field is not specified, it is assumed to be PREMIUM. */
	// +optional
	NetworkTier *string `json:"networkTier,omitempty"`

	/* The DNS domain name for the public PTR record.The DNS domain name for the public PTR record. */
	// +optional
	PublicPtrDomainName *string `json:"publicPtrDomainName,omitempty"`
}

type InstancetemplateAdvancedMachineFeatures struct {
	/* Immutable. Whether to enable nested virtualization or not. */
	// +optional
	EnableNestedVirtualization *bool `json:"enableNestedVirtualization,omitempty"`

	/* Immutable. The number of threads per physical core. To disable simultaneous multithreading (SMT) set this to 1. If unset, the maximum number of threads supported per core by the underlying processor is assumed. */
	// +optional
	ThreadsPerCore *int `json:"threadsPerCore,omitempty"`

	/* Immutable. The number of physical cores to expose to an instance. Multiply by the number of threads per core to compute the total number of virtual CPUs to expose to the instance. If unset, the number of cores is inferred from the instance\'s nominal CPU count and the underlying platform\'s SMT width. */
	// +optional
	VisibleCoreCount *int `json:"visibleCoreCount,omitempty"`
}

type InstancetemplateAliasIpRange struct {
	/* Immutable. The IP CIDR range represented by this alias IP range. This IP CIDR range must belong to the specified subnetwork and cannot contain IP addresses reserved by system or used by other network interfaces. At the time of writing only a netmask (e.g. /24) may be supplied, with a CIDR format resulting in an API error. */
	IpCidrRange string `json:"ipCidrRange"`

	/* Immutable. The subnetwork secondary range name specifying the secondary range from which to allocate the IP CIDR range for this alias IP range. If left unspecified, the primary range of the subnetwork will be used. */
	// +optional
	SubnetworkRangeName *string `json:"subnetworkRangeName,omitempty"`
}

type InstancetemplateConfidentialInstanceConfig struct {
	/* Immutable. Defines whether the instance should have confidential compute enabled. */
	EnableConfidentialCompute bool `json:"enableConfidentialCompute"`
}

type InstancetemplateDisk struct {
	/* Immutable. Whether or not the disk should be auto-deleted. This defaults to true. */
	// +optional
	AutoDelete *bool `json:"autoDelete,omitempty"`

	/* Immutable. Indicates that this is a boot disk. */
	// +optional
	Boot *bool `json:"boot,omitempty"`

	/* Immutable. A unique device name that is reflected into the /dev/ tree of a Linux operating system running within the instance. If not specified, the server chooses a default device name to apply to this disk. */
	// +optional
	DeviceName *string `json:"deviceName,omitempty"`

	/* Immutable. Encrypts or decrypts a disk using a customer-supplied encryption key. */
	// +optional
	DiskEncryptionKey *InstancetemplateDiskEncryptionKey `json:"diskEncryptionKey,omitempty"`

	/* Immutable. Name of the disk. When not provided, this defaults to the name of the instance. */
	// +optional
	DiskName *string `json:"diskName,omitempty"`

	/* Immutable. The size of the image in gigabytes. If not specified, it will inherit the size of its base image. For SCRATCH disks, the size must be one of 375 or 3000 GB, with a default of 375 GB. */
	// +optional
	DiskSizeGb *int `json:"diskSizeGb,omitempty"`

	/* Immutable. The Google Compute Engine disk type. Such as "pd-ssd", "local-ssd", "pd-balanced" or "pd-standard". */
	// +optional
	DiskType *string `json:"diskType,omitempty"`

	/* Immutable. Specifies the disk interface to use for attaching this disk. */
	// +optional
	Interface *string `json:"interface,omitempty"`

	/* Immutable. A set of key/value label pairs to assign to disks,. */
	// +optional
	Labels map[string]string `json:"labels,omitempty"`

	/* Immutable. The mode in which to attach this disk, either READ_WRITE or READ_ONLY. If you are attaching or creating a boot disk, this must read-write mode. */
	// +optional
	Mode *string `json:"mode,omitempty"`

	/* Immutable. Indicates how many IOPS to provision for the disk. This sets the number of I/O operations per second that the disk can handle. Values must be between 10,000 and 120,000. For more details, see the [Extreme persistent disk documentation](https://cloud.google.com/compute/docs/disks/extreme-persistent-disk). */
	// +optional
	ProvisionedIops *int `json:"provisionedIops,omitempty"`

	// +optional
	ResourcePolicies []v1alpha1.ResourceRef `json:"resourcePolicies,omitempty"`

	// +optional
	SourceDiskRef *v1alpha1.ResourceRef `json:"sourceDiskRef,omitempty"`

	/* Immutable. The customer-supplied encryption key of the source
	image. Required if the source image is protected by a
	customer-supplied encryption key.

	Instance templates do not store customer-supplied
	encryption keys, so you cannot create disks for
	instances in a managed instance group if the source
	images are encrypted with your own keys. */
	// +optional
	SourceImageEncryptionKey *InstancetemplateSourceImageEncryptionKey `json:"sourceImageEncryptionKey,omitempty"`

	// +optional
	SourceImageRef *v1alpha1.ResourceRef `json:"sourceImageRef,omitempty"`

	/* Immutable. The customer-supplied encryption key of the source snapshot. */
	// +optional
	SourceSnapshotEncryptionKey *InstancetemplateSourceSnapshotEncryptionKey `json:"sourceSnapshotEncryptionKey,omitempty"`

	/* The source snapshot to create this disk. When creating a new
	instance, one of initializeParams.sourceSnapshot,
	initializeParams.sourceImage, or disks.source is required except for
	local SSD. */
	// +optional
	SourceSnapshotRef *v1alpha1.ResourceRef `json:"sourceSnapshotRef,omitempty"`

	/* Immutable. The type of Google Compute Engine disk, can be either "SCRATCH" or "PERSISTENT". */
	// +optional
	Type *string `json:"type,omitempty"`
}

type InstancetemplateDiskEncryptionKey struct {
	KmsKeyRef v1alpha1.ResourceRef `json:"kmsKeyRef"`
}

type InstancetemplateGuestAccelerator struct {
	/* Immutable. The number of the guest accelerator cards exposed to this instance. */
	Count int `json:"count"`

	/* Immutable. The accelerator type resource to expose to this instance. E.g. nvidia-tesla-k80. */
	Type string `json:"type"`
}

type InstancetemplateIpv6AccessConfig struct {
	/* The first IPv6 address of the external IPv6 range associated with this instance, prefix length is stored in externalIpv6PrefixLength in ipv6AccessConfig. The field is output only, an IPv6 address from a subnetwork associated with the instance will be allocated dynamically. */
	// +optional
	ExternalIpv6 *string `json:"externalIpv6,omitempty"`

	/* The prefix length of the external IPv6 range. */
	// +optional
	ExternalIpv6PrefixLength *string `json:"externalIpv6PrefixLength,omitempty"`

	/* The name of this access configuration. */
	// +optional
	Name *string `json:"name,omitempty"`

	/* The service-level to be provided for IPv6 traffic when the subnet has an external subnet. Only PREMIUM tier is valid for IPv6. */
	NetworkTier string `json:"networkTier"`

	/* The domain name to be used when creating DNSv6 records for the external IPv6 ranges. */
	// +optional
	PublicPtrDomainName *string `json:"publicPtrDomainName,omitempty"`
}

type InstancetemplateLocalSsdRecoveryTimeout struct {
	/* Immutable. Span of time that's a fraction of a second at nanosecond
	resolution. Durations less than one second are represented
	with a 0 seconds field and a positive nanos field. Must
	be from 0 to 999,999,999 inclusive. */
	// +optional
	Nanos *int `json:"nanos,omitempty"`

	/* Immutable. Span of time at a resolution of a second.
	Must be from 0 to 315,576,000,000 inclusive. */
	Seconds int `json:"seconds"`
}

type InstancetemplateMaxRunDuration struct {
	/* Immutable. Span of time that's a fraction of a second at nanosecond
	resolution. Durations less than one second are represented
	with a 0 seconds field and a positive nanos field. Must
	be from 0 to 999,999,999 inclusive. */
	// +optional
	Nanos *int `json:"nanos,omitempty"`

	/* Immutable. Span of time at a resolution of a second.
	Must be from 0 to 315,576,000,000 inclusive. */
	Seconds int `json:"seconds"`
}

type InstancetemplateMetadata struct {
	Key string `json:"key"`

	Value string `json:"value"`
}

type InstancetemplateNetworkInterface struct {
	// +optional
	AccessConfig []InstancetemplateAccessConfig `json:"accessConfig,omitempty"`

	/* Immutable. An array of alias IP ranges for this network interface. Can only be specified for network interfaces on subnet-mode networks. */
	// +optional
	AliasIpRange []InstancetemplateAliasIpRange `json:"aliasIpRange,omitempty"`

	/* The prefix length of the primary internal IPv6 range. */
	// +optional
	InternalIpv6PrefixLength *int `json:"internalIpv6PrefixLength,omitempty"`

	/* An array of IPv6 access configurations for this interface. Currently, only one IPv6 access config, DIRECT_IPV6, is supported. If there is no ipv6AccessConfig specified, then this instance will have no external IPv6 Internet access. */
	// +optional
	Ipv6AccessConfig []InstancetemplateIpv6AccessConfig `json:"ipv6AccessConfig,omitempty"`

	/* One of EXTERNAL, INTERNAL to indicate whether the IP can be accessed from the Internet. This field is always inherited from its subnetwork. */
	// +optional
	Ipv6AccessType *string `json:"ipv6AccessType,omitempty"`

	/* An IPv6 internal network address for this network interface. If not specified, Google Cloud will automatically assign an internal IPv6 address from the instance's subnetwork. */
	// +optional
	Ipv6Address *string `json:"ipv6Address,omitempty"`

	/* The name of the network_interface. */
	// +optional
	Name *string `json:"name,omitempty"`

	/* Immutable. The URL of the network attachment that this interface should connect to in the following format: projects/{projectNumber}/regions/{region_name}/networkAttachments/{network_attachment_name}. */
	// +optional
	NetworkAttachment *string `json:"networkAttachment,omitempty"`

	/* Immutable. The private IP address to assign to the instance. If empty, the address will be automatically assigned. */
	// +optional
	NetworkIp *string `json:"networkIp,omitempty"`

	// +optional
	NetworkRef *v1alpha1.ResourceRef `json:"networkRef,omitempty"`

	/* Immutable. The type of vNIC to be used on this interface. Possible values:GVNIC, VIRTIO_NET. */
	// +optional
	NicType *string `json:"nicType,omitempty"`

	/* Immutable. The networking queue count that's specified by users for the network interface. Both Rx and Tx queues will be set to this number. It will be empty if not specified. */
	// +optional
	QueueCount *int `json:"queueCount,omitempty"`

	/* The stack type for this network interface to identify whether the IPv6 feature is enabled or not. If not specified, IPV4_ONLY will be used. */
	// +optional
	StackType *string `json:"stackType,omitempty"`

	/* Immutable. The ID of the project in which the subnetwork belongs. If it is not provided, the provider project is used. */
	// +optional
	SubnetworkProject *string `json:"subnetworkProject,omitempty"`

	// +optional
	SubnetworkRef *v1alpha1.ResourceRef `json:"subnetworkRef,omitempty"`
}

type InstancetemplateNetworkPerformanceConfig struct {
	/* Immutable. The egress bandwidth tier to enable. Possible values:TIER_1, DEFAULT. */
	TotalEgressBandwidthTier string `json:"totalEgressBandwidthTier"`
}

type InstancetemplateNodeAffinities struct {
	// +optional
	Value *InstancetemplateValue `json:"value,omitempty"`
}

type InstancetemplateReservationAffinity struct {
	/* Immutable. Specifies the label selector for the reservation to use. */
	// +optional
	SpecificReservation *InstancetemplateSpecificReservation `json:"specificReservation,omitempty"`

	/* Immutable. The type of reservation from which this instance can consume resources. */
	Type string `json:"type"`
}

type InstancetemplateScheduling struct {
	/* Immutable. Specifies whether the instance should be automatically restarted if it is terminated by Compute Engine (not terminated by a user). This defaults to true. */
	// +optional
	AutomaticRestart *bool `json:"automaticRestart,omitempty"`

	/* Immutable. Specifies the action GCE should take when SPOT VM is preempted. */
	// +optional
	InstanceTerminationAction *string `json:"instanceTerminationAction,omitempty"`

	/* Specifies the maximum amount of time a Local Ssd Vm should wait while
	recovery of the Local Ssd state is attempted. Its value should be in
	between 0 and 168 hours with hour granularity and the default value being 1
	hour. */
	// +optional
	LocalSsdRecoveryTimeout []InstancetemplateLocalSsdRecoveryTimeout `json:"localSsdRecoveryTimeout,omitempty"`

	/* Specifies the frequency of planned maintenance events. The accepted values are: PERIODIC. */
	// +optional
	MaintenanceInterval *string `json:"maintenanceInterval,omitempty"`

	/* Immutable. The timeout for new network connections to hosts. */
	// +optional
	MaxRunDuration *InstancetemplateMaxRunDuration `json:"maxRunDuration,omitempty"`

	/* Minimum number of cpus for the instance. */
	// +optional
	MinNodeCpus *int `json:"minNodeCpus,omitempty"`

	// +optional
	NodeAffinities []InstancetemplateNodeAffinities `json:"nodeAffinities,omitempty"`

	/* Immutable. Defines the maintenance behavior for this instance. */
	// +optional
	OnHostMaintenance *string `json:"onHostMaintenance,omitempty"`

	/* Immutable. Allows instance to be preempted. This defaults to false. */
	// +optional
	Preemptible *bool `json:"preemptible,omitempty"`

	/* Immutable. Whether the instance is spot. If this is set as SPOT. */
	// +optional
	ProvisioningModel *string `json:"provisioningModel,omitempty"`
}

type InstancetemplateServiceAccount struct {
	/* Immutable. A list of service scopes. Both OAuth2 URLs and gcloud short names are supported. To allow full access to all Cloud APIs, use the cloud-platform scope. */
	Scopes []string `json:"scopes"`

	// +optional
	ServiceAccountRef *v1alpha1.ResourceRef `json:"serviceAccountRef,omitempty"`
}

type InstancetemplateShieldedInstanceConfig struct {
	/* Immutable. Compare the most recent boot measurements to the integrity policy baseline and return a pair of pass/fail results depending on whether they match or not. Defaults to true. */
	// +optional
	EnableIntegrityMonitoring *bool `json:"enableIntegrityMonitoring,omitempty"`

	/* Immutable. Verify the digital signature of all boot components, and halt the boot process if signature verification fails. Defaults to false. */
	// +optional
	EnableSecureBoot *bool `json:"enableSecureBoot,omitempty"`

	/* Immutable. Use a virtualized trusted platform module, which is a specialized computer chip you can use to encrypt objects like keys and certificates. Defaults to true. */
	// +optional
	EnableVtpm *bool `json:"enableVtpm,omitempty"`
}

type InstancetemplateSourceImageEncryptionKey struct {
	/* The self link of the encryption key that is stored in Google Cloud
	KMS. */
	KmsKeySelfLinkRef v1alpha1.ResourceRef `json:"kmsKeySelfLinkRef"`

	/* The service account being used for the encryption request for the
	given KMS key. If absent, the Compute Engine default service account
	is used. */
	// +optional
	KmsKeyServiceAccountRef *v1alpha1.ResourceRef `json:"kmsKeyServiceAccountRef,omitempty"`
}

type InstancetemplateSourceSnapshotEncryptionKey struct {
	/* The self link of the encryption key that is stored in Google Cloud
	KMS. */
	KmsKeySelfLinkRef v1alpha1.ResourceRef `json:"kmsKeySelfLinkRef"`

	/* The service account being used for the encryption request for the
	given KMS key. If absent, the Compute Engine default service account
	is used. */
	// +optional
	KmsKeyServiceAccountRef *v1alpha1.ResourceRef `json:"kmsKeyServiceAccountRef,omitempty"`
}

type InstancetemplateSpecificReservation struct {
	/* Immutable. Corresponds to the label key of a reservation resource. To target a SPECIFIC_RESERVATION by name, specify compute.googleapis.com/reservation-name as the key and specify the name of your reservation as the only value. */
	Key string `json:"key"`

	/* Immutable. Corresponds to the label values of a reservation resource. */
	Values []string `json:"values"`
}

type InstancetemplateValue struct {
}

type ComputeInstanceTemplateSpec struct {
	/* Immutable. Controls for advanced machine-related behavior features. */
	// +optional
	AdvancedMachineFeatures *InstancetemplateAdvancedMachineFeatures `json:"advancedMachineFeatures,omitempty"`

	/* Immutable. Whether to allow sending and receiving of packets with non-matching source or destination IPs. This defaults to false. */
	// +optional
	CanIpForward *bool `json:"canIpForward,omitempty"`

	/* Immutable. The Confidential VM config being used by the instance. on_host_maintenance has to be set to TERMINATE or this will fail to create. */
	// +optional
	ConfidentialInstanceConfig *InstancetemplateConfidentialInstanceConfig `json:"confidentialInstanceConfig,omitempty"`

	/* Immutable. A brief description of this resource. */
	// +optional
	Description *string `json:"description,omitempty"`

	/* Immutable. Disks to attach to instances created from this template. This can be specified multiple times for multiple disks. */
	Disk []InstancetemplateDisk `json:"disk"`

	/* Immutable. Enable Virtual Displays on this instance. Note: allow_stopping_for_update must be set to true in order to update this field. */
	// +optional
	EnableDisplay *bool `json:"enableDisplay,omitempty"`

	/* Immutable. List of the type and count of accelerator cards attached to the instance. */
	// +optional
	GuestAccelerator []InstancetemplateGuestAccelerator `json:"guestAccelerator,omitempty"`

	/* Immutable. A description of the instance. */
	// +optional
	InstanceDescription *string `json:"instanceDescription,omitempty"`

	/* Immutable. The machine type to create. To create a machine with a custom type (such as extended memory), format the value like custom-VCPUS-MEM_IN_MB like custom-6-20480 for 6 vCPU and 20GB of RAM. */
	MachineType string `json:"machineType"`

	// +optional
	Metadata []InstancetemplateMetadata `json:"metadata,omitempty"`

	/* Immutable. An alternative to using the startup-script metadata key, mostly to match the compute_instance resource. This replaces the startup-script metadata key on the created instance and thus the two mechanisms are not allowed to be used simultaneously. */
	// +optional
	MetadataStartupScript *string `json:"metadataStartupScript,omitempty"`

	/* Immutable. Specifies a minimum CPU platform. Applicable values are the friendly names of CPU platforms, such as Intel Haswell or Intel Skylake. */
	// +optional
	MinCpuPlatform *string `json:"minCpuPlatform,omitempty"`

	/* Immutable. Creates a unique name beginning with the specified prefix. Conflicts with name. */
	// +optional
	NamePrefix *string `json:"namePrefix,omitempty"`

	/* Immutable. Networks to attach to instances created from this template. This can be specified multiple times for multiple networks. */
	// +optional
	NetworkInterface []InstancetemplateNetworkInterface `json:"networkInterface,omitempty"`

	/* Immutable. Configures network performance settings for the instance. If not specified, the instance will be created with its default network performance configuration. */
	// +optional
	NetworkPerformanceConfig *InstancetemplateNetworkPerformanceConfig `json:"networkPerformanceConfig,omitempty"`

	/* Immutable. An instance template is a global resource that is not bound to a zone or a region. However, you can still specify some regional resources in an instance template, which restricts the template to the region where that resource resides. For example, a custom subnetwork resource is tied to a specific region. Defaults to the region of the Provider if no value is given. */
	// +optional
	Region *string `json:"region,omitempty"`

	/* Immutable. Specifies the reservations that this instance can consume from. */
	// +optional
	ReservationAffinity *InstancetemplateReservationAffinity `json:"reservationAffinity,omitempty"`

	/* Immutable. Optional. The name of the resource. Used for creation and acquisition. When unset, the value of `metadata.name` is used as the default. */
	// +optional
	ResourceID *string `json:"resourceID,omitempty"`

	// +optional
	ResourcePolicies []v1alpha1.ResourceRef `json:"resourcePolicies,omitempty"`

	/* Immutable. The scheduling strategy to use. */
	// +optional
	Scheduling *InstancetemplateScheduling `json:"scheduling,omitempty"`

	/* Immutable. Service account to attach to the instance. */
	// +optional
	ServiceAccount *InstancetemplateServiceAccount `json:"serviceAccount,omitempty"`

	/* Immutable. Enable Shielded VM on this instance. Shielded VM provides verifiable integrity to prevent against malware and rootkits. Defaults to disabled. Note: shielded_instance_config can only be used with boot images with shielded vm support. */
	// +optional
	ShieldedInstanceConfig *InstancetemplateShieldedInstanceConfig `json:"shieldedInstanceConfig,omitempty"`

	/* Immutable. Tags to attach to the instance. */
	// +optional
	Tags []string `json:"tags,omitempty"`
}

type ComputeInstanceTemplateStatus struct {
	/* Conditions represent the latest available observations of the
	   ComputeInstanceTemplate's current state. */
	Conditions []v1alpha1.Condition `json:"conditions,omitempty"`
	/* The unique fingerprint of the metadata. */
	// +optional
	MetadataFingerprint *string `json:"metadataFingerprint,omitempty"`

	/* ObservedGeneration is the generation of the resource that was most recently observed by the Config Connector controller. If this is equal to metadata.generation, then that means that the current reported status reflects the most recent desired state of the resource. */
	// +optional
	ObservedGeneration *int `json:"observedGeneration,omitempty"`

	/* The URI of the created resource. */
	// +optional
	SelfLink *string `json:"selfLink,omitempty"`

	/* A special URI of the created resource that uniquely identifies this instance template. */
	// +optional
	SelfLinkUnique *string `json:"selfLinkUnique,omitempty"`

	/* The unique fingerprint of the tags. */
	// +optional
	TagsFingerprint *string `json:"tagsFingerprint,omitempty"`
}

// +genclient
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:resource:categories=gcp,shortName=gcpcomputeinstancetemplate;gcpcomputeinstancetemplates
// +kubebuilder:subresource:status
// +kubebuilder:metadata:labels="cnrm.cloud.google.com/managed-by-kcc=true";"cnrm.cloud.google.com/stability-level=stable";"cnrm.cloud.google.com/system=true";"cnrm.cloud.google.com/tf2crd=true"
// +kubebuilder:printcolumn:name="Age",JSONPath=".metadata.creationTimestamp",type="date"
// +kubebuilder:printcolumn:name="Ready",JSONPath=".status.conditions[?(@.type=='Ready')].status",type="string",description="When 'True', the most recent reconcile of the resource succeeded"
// +kubebuilder:printcolumn:name="Status",JSONPath=".status.conditions[?(@.type=='Ready')].reason",type="string",description="The reason for the value in 'Ready'"
// +kubebuilder:printcolumn:name="Status Age",JSONPath=".status.conditions[?(@.type=='Ready')].lastTransitionTime",type="date",description="The last transition time for the value in 'Status'"

// ComputeInstanceTemplate is the Schema for the compute API
// +k8s:openapi-gen=true
type ComputeInstanceTemplate struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   ComputeInstanceTemplateSpec   `json:"spec,omitempty"`
	Status ComputeInstanceTemplateStatus `json:"status,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// ComputeInstanceTemplateList contains a list of ComputeInstanceTemplate
type ComputeInstanceTemplateList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []ComputeInstanceTemplate `json:"items"`
}

func init() {
	SchemeBuilder.Register(&ComputeInstanceTemplate{}, &ComputeInstanceTemplateList{})
}
