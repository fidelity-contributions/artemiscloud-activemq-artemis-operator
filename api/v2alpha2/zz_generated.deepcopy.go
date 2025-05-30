//go:build !ignore_autogenerated

/*
Copyright 2021.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

// Code generated by controller-gen. DO NOT EDIT.

package v2alpha2

import (
	"k8s.io/apimachinery/pkg/apis/meta/v1"
	runtime "k8s.io/apimachinery/pkg/runtime"
)

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *AcceptorType) DeepCopyInto(out *AcceptorType) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new AcceptorType.
func (in *AcceptorType) DeepCopy() *AcceptorType {
	if in == nil {
		return nil
	}
	out := new(AcceptorType)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ActiveMQArtemis) DeepCopyInto(out *ActiveMQArtemis) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ActiveMQArtemis.
func (in *ActiveMQArtemis) DeepCopy() *ActiveMQArtemis {
	if in == nil {
		return nil
	}
	out := new(ActiveMQArtemis)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *ActiveMQArtemis) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ActiveMQArtemisAddress) DeepCopyInto(out *ActiveMQArtemisAddress) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	out.Spec = in.Spec
	in.Status.DeepCopyInto(&out.Status)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ActiveMQArtemisAddress.
func (in *ActiveMQArtemisAddress) DeepCopy() *ActiveMQArtemisAddress {
	if in == nil {
		return nil
	}
	out := new(ActiveMQArtemisAddress)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *ActiveMQArtemisAddress) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ActiveMQArtemisAddressList) DeepCopyInto(out *ActiveMQArtemisAddressList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]ActiveMQArtemisAddress, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ActiveMQArtemisAddressList.
func (in *ActiveMQArtemisAddressList) DeepCopy() *ActiveMQArtemisAddressList {
	if in == nil {
		return nil
	}
	out := new(ActiveMQArtemisAddressList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *ActiveMQArtemisAddressList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ActiveMQArtemisAddressSpec) DeepCopyInto(out *ActiveMQArtemisAddressSpec) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ActiveMQArtemisAddressSpec.
func (in *ActiveMQArtemisAddressSpec) DeepCopy() *ActiveMQArtemisAddressSpec {
	if in == nil {
		return nil
	}
	out := new(ActiveMQArtemisAddressSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ActiveMQArtemisAddressStatus) DeepCopyInto(out *ActiveMQArtemisAddressStatus) {
	*out = *in
	if in.Conditions != nil {
		in, out := &in.Conditions, &out.Conditions
		*out = make([]v1.Condition, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ActiveMQArtemisAddressStatus.
func (in *ActiveMQArtemisAddressStatus) DeepCopy() *ActiveMQArtemisAddressStatus {
	if in == nil {
		return nil
	}
	out := new(ActiveMQArtemisAddressStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ActiveMQArtemisList) DeepCopyInto(out *ActiveMQArtemisList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]ActiveMQArtemis, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ActiveMQArtemisList.
func (in *ActiveMQArtemisList) DeepCopy() *ActiveMQArtemisList {
	if in == nil {
		return nil
	}
	out := new(ActiveMQArtemisList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *ActiveMQArtemisList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ActiveMQArtemisSpec) DeepCopyInto(out *ActiveMQArtemisSpec) {
	*out = *in
	in.DeploymentPlan.DeepCopyInto(&out.DeploymentPlan)
	if in.Acceptors != nil {
		in, out := &in.Acceptors, &out.Acceptors
		*out = make([]AcceptorType, len(*in))
		copy(*out, *in)
	}
	if in.Connectors != nil {
		in, out := &in.Connectors, &out.Connectors
		*out = make([]ConnectorType, len(*in))
		copy(*out, *in)
	}
	out.Console = in.Console
	out.Upgrades = in.Upgrades
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ActiveMQArtemisSpec.
func (in *ActiveMQArtemisSpec) DeepCopy() *ActiveMQArtemisSpec {
	if in == nil {
		return nil
	}
	out := new(ActiveMQArtemisSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ActiveMQArtemisStatus) DeepCopyInto(out *ActiveMQArtemisStatus) {
	*out = *in
	in.PodStatus.DeepCopyInto(&out.PodStatus)
	if in.Conditions != nil {
		in, out := &in.Conditions, &out.Conditions
		*out = make([]v1.Condition, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ActiveMQArtemisStatus.
func (in *ActiveMQArtemisStatus) DeepCopy() *ActiveMQArtemisStatus {
	if in == nil {
		return nil
	}
	out := new(ActiveMQArtemisStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ActiveMQArtemisUpgrades) DeepCopyInto(out *ActiveMQArtemisUpgrades) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ActiveMQArtemisUpgrades.
func (in *ActiveMQArtemisUpgrades) DeepCopy() *ActiveMQArtemisUpgrades {
	if in == nil {
		return nil
	}
	out := new(ActiveMQArtemisUpgrades)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ConnectorType) DeepCopyInto(out *ConnectorType) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ConnectorType.
func (in *ConnectorType) DeepCopy() *ConnectorType {
	if in == nil {
		return nil
	}
	out := new(ConnectorType)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ConsoleType) DeepCopyInto(out *ConsoleType) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ConsoleType.
func (in *ConsoleType) DeepCopy() *ConsoleType {
	if in == nil {
		return nil
	}
	out := new(ConsoleType)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *DeploymentPlanType) DeepCopyInto(out *DeploymentPlanType) {
	*out = *in
	if in.MessageMigration != nil {
		in, out := &in.MessageMigration, &out.MessageMigration
		*out = new(bool)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new DeploymentPlanType.
func (in *DeploymentPlanType) DeepCopy() *DeploymentPlanType {
	if in == nil {
		return nil
	}
	out := new(DeploymentPlanType)
	in.DeepCopyInto(out)
	return out
}
