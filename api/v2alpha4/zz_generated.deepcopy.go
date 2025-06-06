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

package v2alpha4

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
	in.AddressSettings.DeepCopyInto(&out.AddressSettings)
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
func (in *AddressSettingType) DeepCopyInto(out *AddressSettingType) {
	*out = *in
	if in.DeadLetterAddress != nil {
		in, out := &in.DeadLetterAddress, &out.DeadLetterAddress
		*out = new(string)
		**out = **in
	}
	if in.AutoCreateDeadLetterResources != nil {
		in, out := &in.AutoCreateDeadLetterResources, &out.AutoCreateDeadLetterResources
		*out = new(bool)
		**out = **in
	}
	if in.DeadLetterQueuePrefix != nil {
		in, out := &in.DeadLetterQueuePrefix, &out.DeadLetterQueuePrefix
		*out = new(string)
		**out = **in
	}
	if in.DeadLetterQueueSuffix != nil {
		in, out := &in.DeadLetterQueueSuffix, &out.DeadLetterQueueSuffix
		*out = new(string)
		**out = **in
	}
	if in.ExpiryAddress != nil {
		in, out := &in.ExpiryAddress, &out.ExpiryAddress
		*out = new(string)
		**out = **in
	}
	if in.AutoCreateExpiryResources != nil {
		in, out := &in.AutoCreateExpiryResources, &out.AutoCreateExpiryResources
		*out = new(bool)
		**out = **in
	}
	if in.ExpiryQueuePrefix != nil {
		in, out := &in.ExpiryQueuePrefix, &out.ExpiryQueuePrefix
		*out = new(string)
		**out = **in
	}
	if in.ExpiryQueueSuffix != nil {
		in, out := &in.ExpiryQueueSuffix, &out.ExpiryQueueSuffix
		*out = new(string)
		**out = **in
	}
	if in.ExpiryDelay != nil {
		in, out := &in.ExpiryDelay, &out.ExpiryDelay
		*out = new(int32)
		**out = **in
	}
	if in.MinExpiryDelay != nil {
		in, out := &in.MinExpiryDelay, &out.MinExpiryDelay
		*out = new(int32)
		**out = **in
	}
	if in.MaxExpiryDelay != nil {
		in, out := &in.MaxExpiryDelay, &out.MaxExpiryDelay
		*out = new(int32)
		**out = **in
	}
	if in.RedeliveryDelay != nil {
		in, out := &in.RedeliveryDelay, &out.RedeliveryDelay
		*out = new(int32)
		**out = **in
	}
	if in.RedeliveryDelayMultiplier != nil {
		in, out := &in.RedeliveryDelayMultiplier, &out.RedeliveryDelayMultiplier
		*out = new(int32)
		**out = **in
	}
	if in.RedeliveryCollisionAvoidanceFactor != nil {
		in, out := &in.RedeliveryCollisionAvoidanceFactor, &out.RedeliveryCollisionAvoidanceFactor
		*out = new(float32)
		**out = **in
	}
	if in.MaxRedeliveryDelay != nil {
		in, out := &in.MaxRedeliveryDelay, &out.MaxRedeliveryDelay
		*out = new(int32)
		**out = **in
	}
	if in.MaxDeliveryAttempts != nil {
		in, out := &in.MaxDeliveryAttempts, &out.MaxDeliveryAttempts
		*out = new(int32)
		**out = **in
	}
	if in.MaxSizeBytes != nil {
		in, out := &in.MaxSizeBytes, &out.MaxSizeBytes
		*out = new(string)
		**out = **in
	}
	if in.MaxSizeBytesRejectThreshold != nil {
		in, out := &in.MaxSizeBytesRejectThreshold, &out.MaxSizeBytesRejectThreshold
		*out = new(int32)
		**out = **in
	}
	if in.PageSizeBytes != nil {
		in, out := &in.PageSizeBytes, &out.PageSizeBytes
		*out = new(string)
		**out = **in
	}
	if in.PageMaxCacheSize != nil {
		in, out := &in.PageMaxCacheSize, &out.PageMaxCacheSize
		*out = new(int32)
		**out = **in
	}
	if in.AddressFullPolicy != nil {
		in, out := &in.AddressFullPolicy, &out.AddressFullPolicy
		*out = new(string)
		**out = **in
	}
	if in.MessageCounterHistoryDayLimit != nil {
		in, out := &in.MessageCounterHistoryDayLimit, &out.MessageCounterHistoryDayLimit
		*out = new(int32)
		**out = **in
	}
	if in.LastValueQueue != nil {
		in, out := &in.LastValueQueue, &out.LastValueQueue
		*out = new(bool)
		**out = **in
	}
	if in.DefaultLastValueQueue != nil {
		in, out := &in.DefaultLastValueQueue, &out.DefaultLastValueQueue
		*out = new(bool)
		**out = **in
	}
	if in.DefaultLastValueKey != nil {
		in, out := &in.DefaultLastValueKey, &out.DefaultLastValueKey
		*out = new(string)
		**out = **in
	}
	if in.DefaultNonDestructive != nil {
		in, out := &in.DefaultNonDestructive, &out.DefaultNonDestructive
		*out = new(bool)
		**out = **in
	}
	if in.DefaultExclusiveQueue != nil {
		in, out := &in.DefaultExclusiveQueue, &out.DefaultExclusiveQueue
		*out = new(bool)
		**out = **in
	}
	if in.DefaultGroupRebalance != nil {
		in, out := &in.DefaultGroupRebalance, &out.DefaultGroupRebalance
		*out = new(bool)
		**out = **in
	}
	if in.DefaultGroupRebalancePauseDispatch != nil {
		in, out := &in.DefaultGroupRebalancePauseDispatch, &out.DefaultGroupRebalancePauseDispatch
		*out = new(bool)
		**out = **in
	}
	if in.DefaultGroupBuckets != nil {
		in, out := &in.DefaultGroupBuckets, &out.DefaultGroupBuckets
		*out = new(int32)
		**out = **in
	}
	if in.DefaultGroupFirstKey != nil {
		in, out := &in.DefaultGroupFirstKey, &out.DefaultGroupFirstKey
		*out = new(string)
		**out = **in
	}
	if in.DefaultConsumersBeforeDispatch != nil {
		in, out := &in.DefaultConsumersBeforeDispatch, &out.DefaultConsumersBeforeDispatch
		*out = new(int32)
		**out = **in
	}
	if in.DefaultDelayBeforeDispatch != nil {
		in, out := &in.DefaultDelayBeforeDispatch, &out.DefaultDelayBeforeDispatch
		*out = new(int32)
		**out = **in
	}
	if in.RedistributionDelay != nil {
		in, out := &in.RedistributionDelay, &out.RedistributionDelay
		*out = new(int32)
		**out = **in
	}
	if in.SendToDlaOnNoRoute != nil {
		in, out := &in.SendToDlaOnNoRoute, &out.SendToDlaOnNoRoute
		*out = new(bool)
		**out = **in
	}
	if in.SlowConsumerThreshold != nil {
		in, out := &in.SlowConsumerThreshold, &out.SlowConsumerThreshold
		*out = new(int32)
		**out = **in
	}
	if in.SlowConsumerPolicy != nil {
		in, out := &in.SlowConsumerPolicy, &out.SlowConsumerPolicy
		*out = new(string)
		**out = **in
	}
	if in.SlowConsumerCheckPeriod != nil {
		in, out := &in.SlowConsumerCheckPeriod, &out.SlowConsumerCheckPeriod
		*out = new(int32)
		**out = **in
	}
	if in.AutoCreateJmsQueues != nil {
		in, out := &in.AutoCreateJmsQueues, &out.AutoCreateJmsQueues
		*out = new(bool)
		**out = **in
	}
	if in.AutoDeleteJmsQueues != nil {
		in, out := &in.AutoDeleteJmsQueues, &out.AutoDeleteJmsQueues
		*out = new(bool)
		**out = **in
	}
	if in.AutoCreateJmsTopics != nil {
		in, out := &in.AutoCreateJmsTopics, &out.AutoCreateJmsTopics
		*out = new(bool)
		**out = **in
	}
	if in.AutoDeleteJmsTopics != nil {
		in, out := &in.AutoDeleteJmsTopics, &out.AutoDeleteJmsTopics
		*out = new(bool)
		**out = **in
	}
	if in.AutoCreateQueues != nil {
		in, out := &in.AutoCreateQueues, &out.AutoCreateQueues
		*out = new(bool)
		**out = **in
	}
	if in.AutoDeleteQueues != nil {
		in, out := &in.AutoDeleteQueues, &out.AutoDeleteQueues
		*out = new(bool)
		**out = **in
	}
	if in.AutoDeleteCreatedQueues != nil {
		in, out := &in.AutoDeleteCreatedQueues, &out.AutoDeleteCreatedQueues
		*out = new(bool)
		**out = **in
	}
	if in.AutoDeleteQueuesDelay != nil {
		in, out := &in.AutoDeleteQueuesDelay, &out.AutoDeleteQueuesDelay
		*out = new(int32)
		**out = **in
	}
	if in.AutoDeleteQueuesMessageCount != nil {
		in, out := &in.AutoDeleteQueuesMessageCount, &out.AutoDeleteQueuesMessageCount
		*out = new(int32)
		**out = **in
	}
	if in.ConfigDeleteQueues != nil {
		in, out := &in.ConfigDeleteQueues, &out.ConfigDeleteQueues
		*out = new(string)
		**out = **in
	}
	if in.AutoCreateAddresses != nil {
		in, out := &in.AutoCreateAddresses, &out.AutoCreateAddresses
		*out = new(bool)
		**out = **in
	}
	if in.AutoDeleteAddresses != nil {
		in, out := &in.AutoDeleteAddresses, &out.AutoDeleteAddresses
		*out = new(bool)
		**out = **in
	}
	if in.AutoDeleteAddressesDelay != nil {
		in, out := &in.AutoDeleteAddressesDelay, &out.AutoDeleteAddressesDelay
		*out = new(int32)
		**out = **in
	}
	if in.ConfigDeleteAddresses != nil {
		in, out := &in.ConfigDeleteAddresses, &out.ConfigDeleteAddresses
		*out = new(string)
		**out = **in
	}
	if in.ManagementBrowsePageSize != nil {
		in, out := &in.ManagementBrowsePageSize, &out.ManagementBrowsePageSize
		*out = new(int32)
		**out = **in
	}
	if in.DefaultPurgeOnNoConsumers != nil {
		in, out := &in.DefaultPurgeOnNoConsumers, &out.DefaultPurgeOnNoConsumers
		*out = new(bool)
		**out = **in
	}
	if in.DefaultMaxConsumers != nil {
		in, out := &in.DefaultMaxConsumers, &out.DefaultMaxConsumers
		*out = new(int32)
		**out = **in
	}
	if in.DefaultQueueRoutingType != nil {
		in, out := &in.DefaultQueueRoutingType, &out.DefaultQueueRoutingType
		*out = new(string)
		**out = **in
	}
	if in.DefaultAddressRoutingType != nil {
		in, out := &in.DefaultAddressRoutingType, &out.DefaultAddressRoutingType
		*out = new(string)
		**out = **in
	}
	if in.DefaultConsumerWindowSize != nil {
		in, out := &in.DefaultConsumerWindowSize, &out.DefaultConsumerWindowSize
		*out = new(int32)
		**out = **in
	}
	if in.DefaultRingSize != nil {
		in, out := &in.DefaultRingSize, &out.DefaultRingSize
		*out = new(int32)
		**out = **in
	}
	if in.RetroactiveMessageCount != nil {
		in, out := &in.RetroactiveMessageCount, &out.RetroactiveMessageCount
		*out = new(int32)
		**out = **in
	}
	if in.EnableMetrics != nil {
		in, out := &in.EnableMetrics, &out.EnableMetrics
		*out = new(bool)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new AddressSettingType.
func (in *AddressSettingType) DeepCopy() *AddressSettingType {
	if in == nil {
		return nil
	}
	out := new(AddressSettingType)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *AddressSettingsType) DeepCopyInto(out *AddressSettingsType) {
	*out = *in
	if in.ApplyRule != nil {
		in, out := &in.ApplyRule, &out.ApplyRule
		*out = new(string)
		**out = **in
	}
	if in.AddressSetting != nil {
		in, out := &in.AddressSetting, &out.AddressSetting
		*out = make([]AddressSettingType, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new AddressSettingsType.
func (in *AddressSettingsType) DeepCopy() *AddressSettingsType {
	if in == nil {
		return nil
	}
	out := new(AddressSettingsType)
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
	in.Resources.DeepCopyInto(&out.Resources)
	out.Storage = in.Storage
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

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *StorageType) DeepCopyInto(out *StorageType) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new StorageType.
func (in *StorageType) DeepCopy() *StorageType {
	if in == nil {
		return nil
	}
	out := new(StorageType)
	in.DeepCopyInto(out)
	return out
}
