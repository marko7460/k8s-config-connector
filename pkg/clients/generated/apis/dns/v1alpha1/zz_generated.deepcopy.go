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
func (in *DNSResponsePolicy) DeepCopyInto(out *DNSResponsePolicy) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new DNSResponsePolicy.
func (in *DNSResponsePolicy) DeepCopy() *DNSResponsePolicy {
	if in == nil {
		return nil
	}
	out := new(DNSResponsePolicy)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *DNSResponsePolicy) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *DNSResponsePolicyList) DeepCopyInto(out *DNSResponsePolicyList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]DNSResponsePolicy, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new DNSResponsePolicyList.
func (in *DNSResponsePolicyList) DeepCopy() *DNSResponsePolicyList {
	if in == nil {
		return nil
	}
	out := new(DNSResponsePolicyList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *DNSResponsePolicyList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *DNSResponsePolicyRule) DeepCopyInto(out *DNSResponsePolicyRule) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new DNSResponsePolicyRule.
func (in *DNSResponsePolicyRule) DeepCopy() *DNSResponsePolicyRule {
	if in == nil {
		return nil
	}
	out := new(DNSResponsePolicyRule)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *DNSResponsePolicyRule) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *DNSResponsePolicyRuleList) DeepCopyInto(out *DNSResponsePolicyRuleList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]DNSResponsePolicyRule, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new DNSResponsePolicyRuleList.
func (in *DNSResponsePolicyRuleList) DeepCopy() *DNSResponsePolicyRuleList {
	if in == nil {
		return nil
	}
	out := new(DNSResponsePolicyRuleList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *DNSResponsePolicyRuleList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *DNSResponsePolicyRuleSpec) DeepCopyInto(out *DNSResponsePolicyRuleSpec) {
	*out = *in
	if in.Behavior != nil {
		in, out := &in.Behavior, &out.Behavior
		*out = new(string)
		**out = **in
	}
	if in.LocalData != nil {
		in, out := &in.LocalData, &out.LocalData
		*out = new(ResponsepolicyruleLocalData)
		(*in).DeepCopyInto(*out)
	}
	out.ProjectRef = in.ProjectRef
	if in.ResourceID != nil {
		in, out := &in.ResourceID, &out.ResourceID
		*out = new(string)
		**out = **in
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new DNSResponsePolicyRuleSpec.
func (in *DNSResponsePolicyRuleSpec) DeepCopy() *DNSResponsePolicyRuleSpec {
	if in == nil {
		return nil
	}
	out := new(DNSResponsePolicyRuleSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *DNSResponsePolicyRuleStatus) DeepCopyInto(out *DNSResponsePolicyRuleStatus) {
	*out = *in
	if in.Conditions != nil {
		in, out := &in.Conditions, &out.Conditions
		*out = make([]k8sv1alpha1.Condition, len(*in))
		copy(*out, *in)
	}
	if in.ObservedGeneration != nil {
		in, out := &in.ObservedGeneration, &out.ObservedGeneration
		*out = new(int64)
		**out = **in
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new DNSResponsePolicyRuleStatus.
func (in *DNSResponsePolicyRuleStatus) DeepCopy() *DNSResponsePolicyRuleStatus {
	if in == nil {
		return nil
	}
	out := new(DNSResponsePolicyRuleStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *DNSResponsePolicySpec) DeepCopyInto(out *DNSResponsePolicySpec) {
	*out = *in
	if in.Description != nil {
		in, out := &in.Description, &out.Description
		*out = new(string)
		**out = **in
	}
	if in.GkeClusters != nil {
		in, out := &in.GkeClusters, &out.GkeClusters
		*out = make([]ResponsepolicyGkeClusters, len(*in))
		copy(*out, *in)
	}
	if in.Networks != nil {
		in, out := &in.Networks, &out.Networks
		*out = make([]ResponsepolicyNetworks, len(*in))
		copy(*out, *in)
	}
	out.ProjectRef = in.ProjectRef
	if in.ResourceID != nil {
		in, out := &in.ResourceID, &out.ResourceID
		*out = new(string)
		**out = **in
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new DNSResponsePolicySpec.
func (in *DNSResponsePolicySpec) DeepCopy() *DNSResponsePolicySpec {
	if in == nil {
		return nil
	}
	out := new(DNSResponsePolicySpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *DNSResponsePolicyStatus) DeepCopyInto(out *DNSResponsePolicyStatus) {
	*out = *in
	if in.Conditions != nil {
		in, out := &in.Conditions, &out.Conditions
		*out = make([]k8sv1alpha1.Condition, len(*in))
		copy(*out, *in)
	}
	if in.ObservedGeneration != nil {
		in, out := &in.ObservedGeneration, &out.ObservedGeneration
		*out = new(int64)
		**out = **in
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new DNSResponsePolicyStatus.
func (in *DNSResponsePolicyStatus) DeepCopy() *DNSResponsePolicyStatus {
	if in == nil {
		return nil
	}
	out := new(DNSResponsePolicyStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ResponsepolicyGkeClusters) DeepCopyInto(out *ResponsepolicyGkeClusters) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ResponsepolicyGkeClusters.
func (in *ResponsepolicyGkeClusters) DeepCopy() *ResponsepolicyGkeClusters {
	if in == nil {
		return nil
	}
	out := new(ResponsepolicyGkeClusters)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ResponsepolicyNetworks) DeepCopyInto(out *ResponsepolicyNetworks) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ResponsepolicyNetworks.
func (in *ResponsepolicyNetworks) DeepCopy() *ResponsepolicyNetworks {
	if in == nil {
		return nil
	}
	out := new(ResponsepolicyNetworks)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ResponsepolicyruleLocalData) DeepCopyInto(out *ResponsepolicyruleLocalData) {
	*out = *in
	if in.LocalDatas != nil {
		in, out := &in.LocalDatas, &out.LocalDatas
		*out = make([]ResponsepolicyruleLocalDatas, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ResponsepolicyruleLocalData.
func (in *ResponsepolicyruleLocalData) DeepCopy() *ResponsepolicyruleLocalData {
	if in == nil {
		return nil
	}
	out := new(ResponsepolicyruleLocalData)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ResponsepolicyruleLocalDatas) DeepCopyInto(out *ResponsepolicyruleLocalDatas) {
	*out = *in
	if in.Rrdatas != nil {
		in, out := &in.Rrdatas, &out.Rrdatas
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	if in.Ttl != nil {
		in, out := &in.Ttl, &out.Ttl
		*out = new(int64)
		**out = **in
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ResponsepolicyruleLocalDatas.
func (in *ResponsepolicyruleLocalDatas) DeepCopy() *ResponsepolicyruleLocalDatas {
	if in == nil {
		return nil
	}
	out := new(ResponsepolicyruleLocalDatas)
	in.DeepCopyInto(out)
	return out
}
