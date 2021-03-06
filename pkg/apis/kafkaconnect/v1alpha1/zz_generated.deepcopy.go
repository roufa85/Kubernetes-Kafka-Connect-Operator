// +build !ignore_autogenerated

// Copyright 2020 The Amadeus Authors.
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

// Code generated by deepcopy-gen. DO NOT EDIT.

package v1alpha1

import (
	v2beta2 "k8s.io/api/autoscaling/v2beta2"
	runtime "k8s.io/apimachinery/pkg/runtime"
)

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *IngressSpec) DeepCopyInto(out *IngressSpec) {
	*out = *in
	if in.Style != nil {
		in, out := &in.Style, &out.Style
		*out = new(IngressStyle)
		**out = **in
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new IngressSpec.
func (in *IngressSpec) DeepCopy() *IngressSpec {
	if in == nil {
		return nil
	}
	out := new(IngressSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *KafkaConnect) DeepCopyInto(out *KafkaConnect) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	if in.Status != nil {
		in, out := &in.Status, &out.Status
		*out = new(KafkaConnectStatus)
		(*in).DeepCopyInto(*out)
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new KafkaConnect.
func (in *KafkaConnect) DeepCopy() *KafkaConnect {
	if in == nil {
		return nil
	}
	out := new(KafkaConnect)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *KafkaConnect) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *KafkaConnectAutoScaler) DeepCopyInto(out *KafkaConnectAutoScaler) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new KafkaConnectAutoScaler.
func (in *KafkaConnectAutoScaler) DeepCopy() *KafkaConnectAutoScaler {
	if in == nil {
		return nil
	}
	out := new(KafkaConnectAutoScaler)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *KafkaConnectAutoScaler) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *KafkaConnectAutoScalerList) DeepCopyInto(out *KafkaConnectAutoScalerList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]KafkaConnectAutoScaler, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new KafkaConnectAutoScalerList.
func (in *KafkaConnectAutoScalerList) DeepCopy() *KafkaConnectAutoScalerList {
	if in == nil {
		return nil
	}
	out := new(KafkaConnectAutoScalerList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *KafkaConnectAutoScalerList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *KafkaConnectAutoScalerSpec) DeepCopyInto(out *KafkaConnectAutoScalerSpec) {
	*out = *in
	out.KafkaConnectorScaleTargetRef = in.KafkaConnectorScaleTargetRef
	if in.MinTasks != nil {
		in, out := &in.MinTasks, &out.MinTasks
		*out = new(int32)
		**out = **in
	}
	if in.Metrics != nil {
		in, out := &in.Metrics, &out.Metrics
		*out = make([]v2beta2.MetricSpec, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new KafkaConnectAutoScalerSpec.
func (in *KafkaConnectAutoScalerSpec) DeepCopy() *KafkaConnectAutoScalerSpec {
	if in == nil {
		return nil
	}
	out := new(KafkaConnectAutoScalerSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *KafkaConnectAutoScalerStatus) DeepCopyInto(out *KafkaConnectAutoScalerStatus) {
	*out = *in
	in.HorizontalPodAutoscalerStatus.DeepCopyInto(&out.HorizontalPodAutoscalerStatus)
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new KafkaConnectAutoScalerStatus.
func (in *KafkaConnectAutoScalerStatus) DeepCopy() *KafkaConnectAutoScalerStatus {
	if in == nil {
		return nil
	}
	out := new(KafkaConnectAutoScalerStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *KafkaConnectList) DeepCopyInto(out *KafkaConnectList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]KafkaConnect, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new KafkaConnectList.
func (in *KafkaConnectList) DeepCopy() *KafkaConnectList {
	if in == nil {
		return nil
	}
	out := new(KafkaConnectList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *KafkaConnectList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *KafkaConnectSpec) DeepCopyInto(out *KafkaConnectSpec) {
	*out = *in
	in.PodSpec.DeepCopyInto(&out.PodSpec)
	if in.IngressSpec != nil {
		in, out := &in.IngressSpec, &out.IngressSpec
		*out = new(IngressSpec)
		(*in).DeepCopyInto(*out)
	}
	if in.KafkaConnectorsSpec != nil {
		in, out := &in.KafkaConnectorsSpec, &out.KafkaConnectorsSpec
		*out = new(KafkaConnectorsSpec)
		(*in).DeepCopyInto(*out)
	}
	if in.ScaleStabilizationSec != nil {
		in, out := &in.ScaleStabilizationSec, &out.ScaleStabilizationSec
		*out = new(int64)
		**out = **in
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new KafkaConnectSpec.
func (in *KafkaConnectSpec) DeepCopy() *KafkaConnectSpec {
	if in == nil {
		return nil
	}
	out := new(KafkaConnectSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *KafkaConnectStatus) DeepCopyInto(out *KafkaConnectStatus) {
	*out = *in
	if in.KafkaConnectorStatus != nil {
		in, out := &in.KafkaConnectorStatus, &out.KafkaConnectorStatus
		*out = make([]KafkaConnectorStatus, len(*in))
		copy(*out, *in)
	}
	in.LastScaleTime.DeepCopyInto(&out.LastScaleTime)
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new KafkaConnectStatus.
func (in *KafkaConnectStatus) DeepCopy() *KafkaConnectStatus {
	if in == nil {
		return nil
	}
	out := new(KafkaConnectStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *KafkaConnectorConfig) DeepCopyInto(out *KafkaConnectorConfig) {
	*out = *in
	if in.TasksMax != nil {
		in, out := &in.TasksMax, &out.TasksMax
		*out = new(int32)
		**out = **in
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new KafkaConnectorConfig.
func (in *KafkaConnectorConfig) DeepCopy() *KafkaConnectorConfig {
	if in == nil {
		return nil
	}
	out := new(KafkaConnectorConfig)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *KafkaConnectorReference) DeepCopyInto(out *KafkaConnectorReference) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new KafkaConnectorReference.
func (in *KafkaConnectorReference) DeepCopy() *KafkaConnectorReference {
	if in == nil {
		return nil
	}
	out := new(KafkaConnectorReference)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *KafkaConnectorStatus) DeepCopyInto(out *KafkaConnectorStatus) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new KafkaConnectorStatus.
func (in *KafkaConnectorStatus) DeepCopy() *KafkaConnectorStatus {
	if in == nil {
		return nil
	}
	out := new(KafkaConnectorStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *KafkaConnectorsSpec) DeepCopyInto(out *KafkaConnectorsSpec) {
	*out = *in
	if in.Configs != nil {
		in, out := &in.Configs, &out.Configs
		*out = make([]KafkaConnectorConfig, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.TaskPerPod != nil {
		in, out := &in.TaskPerPod, &out.TaskPerPod
		*out = new(int32)
		**out = **in
	}
	if in.InitDeploymentReplicas != nil {
		in, out := &in.InitDeploymentReplicas, &out.InitDeploymentReplicas
		*out = new(int32)
		**out = **in
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new KafkaConnectorsSpec.
func (in *KafkaConnectorsSpec) DeepCopy() *KafkaConnectorsSpec {
	if in == nil {
		return nil
	}
	out := new(KafkaConnectorsSpec)
	in.DeepCopyInto(out)
	return out
}
