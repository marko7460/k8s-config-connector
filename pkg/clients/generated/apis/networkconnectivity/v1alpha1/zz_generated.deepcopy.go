//go:build !ignore_autogenerated
// +build !ignore_autogenerated

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

// Code generated by deepcopy-gen. DO NOT EDIT.

package v1alpha1

import (
	k8sv1alpha1 "github.com/GoogleCloudPlatform/k8s-config-connector/pkg/clients/generated/apis/k8s/v1alpha1"
	runtime "k8s.io/apimachinery/pkg/runtime"
)

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *NetworkConnectivityServiceConnectionPolicy) DeepCopyInto(out *NetworkConnectivityServiceConnectionPolicy) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new NetworkConnectivityServiceConnectionPolicy.
func (in *NetworkConnectivityServiceConnectionPolicy) DeepCopy() *NetworkConnectivityServiceConnectionPolicy {
	if in == nil {
		return nil
	}
	out := new(NetworkConnectivityServiceConnectionPolicy)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *NetworkConnectivityServiceConnectionPolicy) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *NetworkConnectivityServiceConnectionPolicyList) DeepCopyInto(out *NetworkConnectivityServiceConnectionPolicyList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]NetworkConnectivityServiceConnectionPolicy, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new NetworkConnectivityServiceConnectionPolicyList.
func (in *NetworkConnectivityServiceConnectionPolicyList) DeepCopy() *NetworkConnectivityServiceConnectionPolicyList {
	if in == nil {
		return nil
	}
	out := new(NetworkConnectivityServiceConnectionPolicyList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *NetworkConnectivityServiceConnectionPolicyList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *NetworkConnectivityServiceConnectionPolicySpec) DeepCopyInto(out *NetworkConnectivityServiceConnectionPolicySpec) {
	*out = *in
	if in.Description != nil {
		in, out := &in.Description, &out.Description
		*out = new(string)
		**out = **in
	}
	if in.NetworkRef != nil {
		in, out := &in.NetworkRef, &out.NetworkRef
		*out = new(k8sv1alpha1.ResourceRef)
		**out = **in
	}
	out.ProjectRef = in.ProjectRef
	if in.PscConfig != nil {
		in, out := &in.PscConfig, &out.PscConfig
		*out = new(ServiceconnectionpolicyPscConfig)
		(*in).DeepCopyInto(*out)
	}
	if in.ResourceID != nil {
		in, out := &in.ResourceID, &out.ResourceID
		*out = new(string)
		**out = **in
	}
	if in.ServiceClass != nil {
		in, out := &in.ServiceClass, &out.ServiceClass
		*out = new(string)
		**out = **in
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new NetworkConnectivityServiceConnectionPolicySpec.
func (in *NetworkConnectivityServiceConnectionPolicySpec) DeepCopy() *NetworkConnectivityServiceConnectionPolicySpec {
	if in == nil {
		return nil
	}
	out := new(NetworkConnectivityServiceConnectionPolicySpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *NetworkConnectivityServiceConnectionPolicyStatus) DeepCopyInto(out *NetworkConnectivityServiceConnectionPolicyStatus) {
	*out = *in
	if in.Conditions != nil {
		in, out := &in.Conditions, &out.Conditions
		*out = make([]k8sv1alpha1.Condition, len(*in))
		copy(*out, *in)
	}
	if in.ExternalRef != nil {
		in, out := &in.ExternalRef, &out.ExternalRef
		*out = new(string)
		**out = **in
	}
	if in.ObservedGeneration != nil {
		in, out := &in.ObservedGeneration, &out.ObservedGeneration
		*out = new(int64)
		**out = **in
	}
	if in.ObservedState != nil {
		in, out := &in.ObservedState, &out.ObservedState
		*out = new(ServiceconnectionpolicyObservedStateStatus)
		(*in).DeepCopyInto(*out)
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new NetworkConnectivityServiceConnectionPolicyStatus.
func (in *NetworkConnectivityServiceConnectionPolicyStatus) DeepCopy() *NetworkConnectivityServiceConnectionPolicyStatus {
	if in == nil {
		return nil
	}
	out := new(NetworkConnectivityServiceConnectionPolicyStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ServiceconnectionpolicyErrorInfoStatus) DeepCopyInto(out *ServiceconnectionpolicyErrorInfoStatus) {
	*out = *in
	if in.Domain != nil {
		in, out := &in.Domain, &out.Domain
		*out = new(string)
		**out = **in
	}
	if in.Metadata != nil {
		in, out := &in.Metadata, &out.Metadata
		*out = make(map[string]string, len(*in))
		for key, val := range *in {
			(*out)[key] = val
		}
	}
	if in.Reason != nil {
		in, out := &in.Reason, &out.Reason
		*out = new(string)
		**out = **in
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ServiceconnectionpolicyErrorInfoStatus.
func (in *ServiceconnectionpolicyErrorInfoStatus) DeepCopy() *ServiceconnectionpolicyErrorInfoStatus {
	if in == nil {
		return nil
	}
	out := new(ServiceconnectionpolicyErrorInfoStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ServiceconnectionpolicyErrorStatus) DeepCopyInto(out *ServiceconnectionpolicyErrorStatus) {
	*out = *in
	if in.Code != nil {
		in, out := &in.Code, &out.Code
		*out = new(int32)
		**out = **in
	}
	if in.Message != nil {
		in, out := &in.Message, &out.Message
		*out = new(string)
		**out = **in
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ServiceconnectionpolicyErrorStatus.
func (in *ServiceconnectionpolicyErrorStatus) DeepCopy() *ServiceconnectionpolicyErrorStatus {
	if in == nil {
		return nil
	}
	out := new(ServiceconnectionpolicyErrorStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ServiceconnectionpolicyObservedStateStatus) DeepCopyInto(out *ServiceconnectionpolicyObservedStateStatus) {
	*out = *in
	if in.CreateTime != nil {
		in, out := &in.CreateTime, &out.CreateTime
		*out = new(string)
		**out = **in
	}
	if in.Etag != nil {
		in, out := &in.Etag, &out.Etag
		*out = new(string)
		**out = **in
	}
	if in.Infrastructure != nil {
		in, out := &in.Infrastructure, &out.Infrastructure
		*out = new(string)
		**out = **in
	}
	if in.PscConnections != nil {
		in, out := &in.PscConnections, &out.PscConnections
		*out = make([]ServiceconnectionpolicyPscConnectionsStatus, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.UpdateTime != nil {
		in, out := &in.UpdateTime, &out.UpdateTime
		*out = new(string)
		**out = **in
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ServiceconnectionpolicyObservedStateStatus.
func (in *ServiceconnectionpolicyObservedStateStatus) DeepCopy() *ServiceconnectionpolicyObservedStateStatus {
	if in == nil {
		return nil
	}
	out := new(ServiceconnectionpolicyObservedStateStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ServiceconnectionpolicyPscConfig) DeepCopyInto(out *ServiceconnectionpolicyPscConfig) {
	*out = *in
	if in.Limit != nil {
		in, out := &in.Limit, &out.Limit
		*out = new(int64)
		**out = **in
	}
	if in.ProducerInstanceLocation != nil {
		in, out := &in.ProducerInstanceLocation, &out.ProducerInstanceLocation
		*out = new(string)
		**out = **in
	}
	if in.SubnetworkRefs != nil {
		in, out := &in.SubnetworkRefs, &out.SubnetworkRefs
		*out = make([]k8sv1alpha1.ResourceRef, len(*in))
		copy(*out, *in)
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ServiceconnectionpolicyPscConfig.
func (in *ServiceconnectionpolicyPscConfig) DeepCopy() *ServiceconnectionpolicyPscConfig {
	if in == nil {
		return nil
	}
	out := new(ServiceconnectionpolicyPscConfig)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ServiceconnectionpolicyPscConnectionsStatus) DeepCopyInto(out *ServiceconnectionpolicyPscConnectionsStatus) {
	*out = *in
	if in.ConsumerAddress != nil {
		in, out := &in.ConsumerAddress, &out.ConsumerAddress
		*out = new(string)
		**out = **in
	}
	if in.ConsumerForwardingRule != nil {
		in, out := &in.ConsumerForwardingRule, &out.ConsumerForwardingRule
		*out = new(string)
		**out = **in
	}
	if in.ConsumerTargetProject != nil {
		in, out := &in.ConsumerTargetProject, &out.ConsumerTargetProject
		*out = new(string)
		**out = **in
	}
	if in.Error != nil {
		in, out := &in.Error, &out.Error
		*out = new(ServiceconnectionpolicyErrorStatus)
		(*in).DeepCopyInto(*out)
	}
	if in.ErrorInfo != nil {
		in, out := &in.ErrorInfo, &out.ErrorInfo
		*out = new(ServiceconnectionpolicyErrorInfoStatus)
		(*in).DeepCopyInto(*out)
	}
	if in.ErrorType != nil {
		in, out := &in.ErrorType, &out.ErrorType
		*out = new(string)
		**out = **in
	}
	if in.GceOperation != nil {
		in, out := &in.GceOperation, &out.GceOperation
		*out = new(string)
		**out = **in
	}
	if in.ProducerInstanceID != nil {
		in, out := &in.ProducerInstanceID, &out.ProducerInstanceID
		*out = new(string)
		**out = **in
	}
	if in.PscConnectionID != nil {
		in, out := &in.PscConnectionID, &out.PscConnectionID
		*out = new(string)
		**out = **in
	}
	if in.SelectedSubnetwork != nil {
		in, out := &in.SelectedSubnetwork, &out.SelectedSubnetwork
		*out = new(string)
		**out = **in
	}
	if in.State != nil {
		in, out := &in.State, &out.State
		*out = new(string)
		**out = **in
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ServiceconnectionpolicyPscConnectionsStatus.
func (in *ServiceconnectionpolicyPscConnectionsStatus) DeepCopy() *ServiceconnectionpolicyPscConnectionsStatus {
	if in == nil {
		return nil
	}
	out := new(ServiceconnectionpolicyPscConnectionsStatus)
	in.DeepCopyInto(out)
	return out
}
