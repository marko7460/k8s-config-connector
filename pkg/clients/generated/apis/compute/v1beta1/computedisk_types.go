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

type DiskAsyncPrimaryDisk struct {
	/* Immutable. Primary disk for asynchronous disk replication. */
	DiskRef v1alpha1.ResourceRef `json:"diskRef"`
}

type DiskDiskEncryptionKey struct {
	/* The encryption key used to encrypt the disk. Your project's Compute
	Engine System service account
	('service-{{PROJECT_NUMBER}}@compute-system.iam.gserviceaccount.com')
	must have 'roles/cloudkms.cryptoKeyEncrypterDecrypter' to use this
	feature. See
	https://cloud.google.com/compute/docs/disks/customer-managed-encryption#encrypt_a_new_persistent_disk_with_your_own_keys */
	// +optional
	KmsKeyRef *v1alpha1.ResourceRef `json:"kmsKeyRef,omitempty"`

	/* The service account used for the encryption request for the given KMS key.
	If absent, the Compute Engine Service Agent service account is used. */
	// +optional
	KmsKeyServiceAccountRef *v1alpha1.ResourceRef `json:"kmsKeyServiceAccountRef,omitempty"`

	/* Immutable. Specifies a 256-bit customer-supplied encryption key, encoded in
	RFC 4648 base64 to either encrypt or decrypt this resource. */
	// +optional
	RawKey *DiskRawKey `json:"rawKey,omitempty"`

	/* Immutable. Specifies an RFC 4648 base64 encoded, RSA-wrapped 2048-bit
	customer-supplied encryption key to either encrypt or decrypt
	this resource. You can provide either the rawKey or the rsaEncryptedKey. */
	// +optional
	RsaEncryptedKey *DiskRsaEncryptedKey `json:"rsaEncryptedKey,omitempty"`

	/* The RFC 4648 base64 encoded SHA-256 hash of the customer-supplied
	encryption key that protects this resource. */
	// +optional
	Sha256 *string `json:"sha256,omitempty"`
}

type DiskGuestOsFeatures struct {
	/* Immutable. The type of supported feature. Read [Enabling guest operating system features](https://cloud.google.com/compute/docs/images/create-delete-deprecate-private-images#guest-os-features) to see a list of available options. Possible values: ["MULTI_IP_SUBNET", "SECURE_BOOT", "SEV_CAPABLE", "UEFI_COMPATIBLE", "VIRTIO_SCSI_MULTIQUEUE", "WINDOWS", "GVNIC", "SEV_LIVE_MIGRATABLE", "SEV_SNP_CAPABLE", "SUSPEND_RESUME_COMPATIBLE", "TDX_CAPABLE"]. */
	Type string `json:"type"`
}

type DiskRawKey struct {
	/* Value of the field. Cannot be used if 'valueFrom' is specified. */
	// +optional
	Value *string `json:"value,omitempty"`

	/* Source for the field's value. Cannot be used if 'value' is specified. */
	// +optional
	ValueFrom *DiskValueFrom `json:"valueFrom,omitempty"`
}

type DiskRsaEncryptedKey struct {
	/* Value of the field. Cannot be used if 'valueFrom' is specified. */
	// +optional
	Value *string `json:"value,omitempty"`

	/* Source for the field's value. Cannot be used if 'value' is specified. */
	// +optional
	ValueFrom *DiskValueFrom `json:"valueFrom,omitempty"`
}

type DiskSourceImageEncryptionKey struct {
	/* The encryption key used to encrypt the disk. Your project's Compute
	Engine System service account
	('service-{{PROJECT_NUMBER}}@compute-system.iam.gserviceaccount.com')
	must have 'roles/cloudkms.cryptoKeyEncrypterDecrypter' to use this
	feature. See
	https://cloud.google.com/compute/docs/disks/customer-managed-encryption#encrypt_a_new_persistent_disk_with_your_own_keys */
	// +optional
	KmsKeyRef *v1alpha1.ResourceRef `json:"kmsKeyRef,omitempty"`

	/* The service account used for the encryption request for the given KMS key.
	If absent, the Compute Engine Service Agent service account is used. */
	// +optional
	KmsKeyServiceAccountRef *v1alpha1.ResourceRef `json:"kmsKeyServiceAccountRef,omitempty"`

	/* Immutable. Specifies a 256-bit customer-supplied encryption key, encoded in
	RFC 4648 base64 to either encrypt or decrypt this resource. */
	// +optional
	RawKey *string `json:"rawKey,omitempty"`

	/* The RFC 4648 base64 encoded SHA-256 hash of the customer-supplied
	encryption key that protects this resource. */
	// +optional
	Sha256 *string `json:"sha256,omitempty"`
}

type DiskSourceSnapshotEncryptionKey struct {
	/* The encryption key used to encrypt the disk. Your project's Compute
	Engine System service account
	('service-{{PROJECT_NUMBER}}@compute-system.iam.gserviceaccount.com')
	must have 'roles/cloudkms.cryptoKeyEncrypterDecrypter' to use this
	feature. See
	https://cloud.google.com/compute/docs/disks/customer-managed-encryption#encrypt_a_new_persistent_disk_with_your_own_keys */
	// +optional
	KmsKeyRef *v1alpha1.ResourceRef `json:"kmsKeyRef,omitempty"`

	/* The service account used for the encryption request for the given KMS key.
	If absent, the Compute Engine Service Agent service account is used. */
	// +optional
	KmsKeyServiceAccountRef *v1alpha1.ResourceRef `json:"kmsKeyServiceAccountRef,omitempty"`

	/* Immutable. Specifies a 256-bit customer-supplied encryption key, encoded in
	RFC 4648 base64 to either encrypt or decrypt this resource. */
	// +optional
	RawKey *string `json:"rawKey,omitempty"`

	/* The RFC 4648 base64 encoded SHA-256 hash of the customer-supplied
	encryption key that protects this resource. */
	// +optional
	Sha256 *string `json:"sha256,omitempty"`
}

type DiskValueFrom struct {
	/* Reference to a value with the given key in the given Secret in the resource's namespace. */
	// +optional
	SecretKeyRef *v1alpha1.SecretKeyRef `json:"secretKeyRef,omitempty"`
}

type ComputeDiskSpec struct {
	/* Immutable. A nested object resource. */
	// +optional
	AsyncPrimaryDisk *DiskAsyncPrimaryDisk `json:"asyncPrimaryDisk,omitempty"`

	/* Immutable. An optional description of this resource. Provide this property when
	you create the resource. */
	// +optional
	Description *string `json:"description,omitempty"`

	/* Immutable. Encrypts the disk using a customer-supplied encryption key.

	After you encrypt a disk with a customer-supplied key, you must
	provide the same key if you use the disk later (e.g. to create a disk
	snapshot or an image, or to attach the disk to a virtual machine).

	Customer-supplied encryption keys do not protect access to metadata of
	the disk.

	If you do not provide an encryption key when creating the disk, then
	the disk will be encrypted using an automatically generated key and
	you do not need to provide a key to use the disk later. */
	// +optional
	DiskEncryptionKey *DiskDiskEncryptionKey `json:"diskEncryptionKey,omitempty"`

	/* Immutable. Whether this disk is using confidential compute mode.
	Note: Only supported on hyperdisk skus, disk_encryption_key is required when setting to true. */
	// +optional
	EnableConfidentialCompute *bool `json:"enableConfidentialCompute,omitempty"`

	/* Immutable. A list of features to enable on the guest operating system.
	Applicable only for bootable disks. */
	// +optional
	GuestOsFeatures []DiskGuestOsFeatures `json:"guestOsFeatures,omitempty"`

	/* The image from which to initialize this disk. */
	// +optional
	ImageRef *v1alpha1.ResourceRef `json:"imageRef,omitempty"`

	/* DEPRECATED. `interface` is deprecated. This field is no longer used and can be safely removed from your configurations; disk interfaces are automatically determined on attachment. Immutable. Specifies the disk interface to use for attaching this disk, which is either SCSI or NVME. The default is SCSI. */
	// +optional
	Interface *string `json:"interface,omitempty"`

	/* Immutable. Any applicable license URI. */
	// +optional
	Licenses []string `json:"licenses,omitempty"`

	/* Location represents the geographical location of the ComputeDisk. Specify a region name or a zone name. Reference: GCP definition of regions/zones (https://cloud.google.com/compute/docs/regions-zones/) */
	Location string `json:"location"`

	/* Immutable. Indicates whether or not the disk can be read/write attached to more than one instance. */
	// +optional
	MultiWriter *bool `json:"multiWriter,omitempty"`

	/* Immutable. Physical block size of the persistent disk, in bytes. If not present
	in a request, a default value is used. Currently supported sizes
	are 4096 and 16384, other sizes may be added in the future.
	If an unsupported value is requested, the error message will list
	the supported values for the caller's project. */
	// +optional
	PhysicalBlockSizeBytes *int `json:"physicalBlockSizeBytes,omitempty"`

	/* The project that this resource belongs to. */
	// +optional
	ProjectRef *v1alpha1.ResourceRef `json:"projectRef,omitempty"`

	/* Indicates how many IOPS must be provisioned for the disk.
	Note: Updating currently is only supported by hyperdisk skus without the need to delete and recreate the disk, hyperdisk
	allows for an update of IOPS every 4 hours. To update your hyperdisk more frequently, you'll need to manually delete and recreate it. */
	// +optional
	ProvisionedIops *int `json:"provisionedIops,omitempty"`

	/* Indicates how much Throughput must be provisioned for the disk.
	Note: Updating currently is only supported by hyperdisk skus without the need to delete and recreate the disk, hyperdisk
	allows for an update of Throughput every 4 hours. To update your hyperdisk more frequently, you'll need to manually delete and recreate it. */
	// +optional
	ProvisionedThroughput *int `json:"provisionedThroughput,omitempty"`

	/* Immutable. URLs of the zones where the disk should be replicated to. */
	// +optional
	ReplicaZones []string `json:"replicaZones,omitempty"`

	/* Immutable. Optional. The name of the resource. Used for creation and acquisition. When unset, the value of `metadata.name` is used as the default. */
	// +optional
	ResourceID *string `json:"resourceID,omitempty"`

	// +optional
	ResourcePolicies []v1alpha1.ResourceRef `json:"resourcePolicies,omitempty"`

	/* Size of the persistent disk, specified in GB. You can specify this
	field when creating a persistent disk using the 'image' or
	'snapshot' parameter, or specify it alone to create an empty
	persistent disk.

	If you specify this field along with 'image' or 'snapshot',
	the value must not be less than the size of the image
	or the size of the snapshot.

	Upsizing the disk is mutable, but downsizing the disk
	requires re-creating the resource. */
	// +optional
	Size *int `json:"size,omitempty"`

	/* The source snapshot used to create this disk. */
	// +optional
	SnapshotRef *v1alpha1.ResourceRef `json:"snapshotRef,omitempty"`

	/* The source disk used to create this disk. */
	// +optional
	SourceDiskRef *v1alpha1.ResourceRef `json:"sourceDiskRef,omitempty"`

	/* Immutable. The customer-supplied encryption key of the source image. Required if
	the source image is protected by a customer-supplied encryption key. */
	// +optional
	SourceImageEncryptionKey *DiskSourceImageEncryptionKey `json:"sourceImageEncryptionKey,omitempty"`

	/* Immutable. The customer-supplied encryption key of the source snapshot. Required
	if the source snapshot is protected by a customer-supplied encryption
	key. */
	// +optional
	SourceSnapshotEncryptionKey *DiskSourceSnapshotEncryptionKey `json:"sourceSnapshotEncryptionKey,omitempty"`

	/* Immutable. URL of the disk type resource describing which disk type to use to
	create the disk. Provide this when creating the disk. */
	// +optional
	Type *string `json:"type,omitempty"`
}

type ComputeDiskStatus struct {
	/* Conditions represent the latest available observations of the
	   ComputeDisk's current state. */
	Conditions []v1alpha1.Condition `json:"conditions,omitempty"`
	/* Creation timestamp in RFC3339 text format. */
	// +optional
	CreationTimestamp *string `json:"creationTimestamp,omitempty"`

	/* The fingerprint used for optimistic locking of this resource.  Used
	internally during updates. */
	// +optional
	LabelFingerprint *string `json:"labelFingerprint,omitempty"`

	/* Last attach timestamp in RFC3339 text format. */
	// +optional
	LastAttachTimestamp *string `json:"lastAttachTimestamp,omitempty"`

	/* Last detach timestamp in RFC3339 text format. */
	// +optional
	LastDetachTimestamp *string `json:"lastDetachTimestamp,omitempty"`

	/* ObservedGeneration is the generation of the resource that was most recently observed by the Config Connector controller. If this is equal to metadata.generation, then that means that the current reported status reflects the most recent desired state of the resource. */
	// +optional
	ObservedGeneration *int `json:"observedGeneration,omitempty"`

	// +optional
	SelfLink *string `json:"selfLink,omitempty"`

	/* The ID value of the disk used to create this image. This value may
	be used to determine whether the image was taken from the current
	or a previous instance of a given disk name. */
	// +optional
	SourceDiskId *string `json:"sourceDiskId,omitempty"`

	/* The ID value of the image used to create this disk. This value
	identifies the exact image that was used to create this persistent
	disk. For example, if you created the persistent disk from an image
	that was later deleted and recreated under the same name, the source
	image ID would identify the exact version of the image that was used. */
	// +optional
	SourceImageId *string `json:"sourceImageId,omitempty"`

	/* The unique ID of the snapshot used to create this disk. This value
	identifies the exact snapshot that was used to create this persistent
	disk. For example, if you created the persistent disk from a snapshot
	that was later deleted and recreated under the same name, the source
	snapshot ID would identify the exact version of the snapshot that was
	used. */
	// +optional
	SourceSnapshotId *string `json:"sourceSnapshotId,omitempty"`

	/* Links to the users of the disk (attached instances) in form:
	project/zones/zone/instances/instance. */
	// +optional
	Users []string `json:"users,omitempty"`
}

// +genclient
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:resource:categories=gcp,shortName=gcpcomputedisk;gcpcomputedisks
// +kubebuilder:subresource:status
// +kubebuilder:metadata:labels="cnrm.cloud.google.com/managed-by-kcc=true";"cnrm.cloud.google.com/stability-level=stable";"cnrm.cloud.google.com/system=true";"cnrm.cloud.google.com/tf2crd=true"
// +kubebuilder:printcolumn:name="Age",JSONPath=".metadata.creationTimestamp",type="date"
// +kubebuilder:printcolumn:name="Ready",JSONPath=".status.conditions[?(@.type=='Ready')].status",type="string",description="When 'True', the most recent reconcile of the resource succeeded"
// +kubebuilder:printcolumn:name="Status",JSONPath=".status.conditions[?(@.type=='Ready')].reason",type="string",description="The reason for the value in 'Ready'"
// +kubebuilder:printcolumn:name="Status Age",JSONPath=".status.conditions[?(@.type=='Ready')].lastTransitionTime",type="date",description="The last transition time for the value in 'Status'"

// ComputeDisk is the Schema for the compute API
// +k8s:openapi-gen=true
type ComputeDisk struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   ComputeDiskSpec   `json:"spec,omitempty"`
	Status ComputeDiskStatus `json:"status,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// ComputeDiskList contains a list of ComputeDisk
type ComputeDiskList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []ComputeDisk `json:"items"`
}

func init() {
	SchemeBuilder.Register(&ComputeDisk{}, &ComputeDiskList{})
}
