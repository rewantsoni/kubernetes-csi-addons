//go:build !ignore_autogenerated

/*
Copyright 2022 The Kubernetes-CSI-Addons Authors.

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

package v1alpha1

import (
	"k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	runtime "k8s.io/apimachinery/pkg/runtime"
)

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *CSIAddonsNode) DeepCopyInto(out *CSIAddonsNode) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	out.Spec = in.Spec
	out.Status = in.Status
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new CSIAddonsNode.
func (in *CSIAddonsNode) DeepCopy() *CSIAddonsNode {
	if in == nil {
		return nil
	}
	out := new(CSIAddonsNode)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *CSIAddonsNode) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *CSIAddonsNodeDriver) DeepCopyInto(out *CSIAddonsNodeDriver) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new CSIAddonsNodeDriver.
func (in *CSIAddonsNodeDriver) DeepCopy() *CSIAddonsNodeDriver {
	if in == nil {
		return nil
	}
	out := new(CSIAddonsNodeDriver)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *CSIAddonsNodeList) DeepCopyInto(out *CSIAddonsNodeList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]CSIAddonsNode, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new CSIAddonsNodeList.
func (in *CSIAddonsNodeList) DeepCopy() *CSIAddonsNodeList {
	if in == nil {
		return nil
	}
	out := new(CSIAddonsNodeList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *CSIAddonsNodeList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *CSIAddonsNodeSpec) DeepCopyInto(out *CSIAddonsNodeSpec) {
	*out = *in
	out.Driver = in.Driver
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new CSIAddonsNodeSpec.
func (in *CSIAddonsNodeSpec) DeepCopy() *CSIAddonsNodeSpec {
	if in == nil {
		return nil
	}
	out := new(CSIAddonsNodeSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *CSIAddonsNodeStatus) DeepCopyInto(out *CSIAddonsNodeStatus) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new CSIAddonsNodeStatus.
func (in *CSIAddonsNodeStatus) DeepCopy() *CSIAddonsNodeStatus {
	if in == nil {
		return nil
	}
	out := new(CSIAddonsNodeStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *EncryptionKeyRotationCronJob) DeepCopyInto(out *EncryptionKeyRotationCronJob) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new EncryptionKeyRotationCronJob.
func (in *EncryptionKeyRotationCronJob) DeepCopy() *EncryptionKeyRotationCronJob {
	if in == nil {
		return nil
	}
	out := new(EncryptionKeyRotationCronJob)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *EncryptionKeyRotationCronJob) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *EncryptionKeyRotationCronJobList) DeepCopyInto(out *EncryptionKeyRotationCronJobList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]EncryptionKeyRotationCronJob, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new EncryptionKeyRotationCronJobList.
func (in *EncryptionKeyRotationCronJobList) DeepCopy() *EncryptionKeyRotationCronJobList {
	if in == nil {
		return nil
	}
	out := new(EncryptionKeyRotationCronJobList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *EncryptionKeyRotationCronJobList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *EncryptionKeyRotationCronJobSpec) DeepCopyInto(out *EncryptionKeyRotationCronJobSpec) {
	*out = *in
	if in.StartingDeadlineSeconds != nil {
		in, out := &in.StartingDeadlineSeconds, &out.StartingDeadlineSeconds
		*out = new(int64)
		**out = **in
	}
	if in.Suspend != nil {
		in, out := &in.Suspend, &out.Suspend
		*out = new(bool)
		**out = **in
	}
	in.JobSpec.DeepCopyInto(&out.JobSpec)
	if in.SuccessfulJobsHistoryLimit != nil {
		in, out := &in.SuccessfulJobsHistoryLimit, &out.SuccessfulJobsHistoryLimit
		*out = new(int32)
		**out = **in
	}
	if in.FailedJobsHistoryLimit != nil {
		in, out := &in.FailedJobsHistoryLimit, &out.FailedJobsHistoryLimit
		*out = new(int32)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new EncryptionKeyRotationCronJobSpec.
func (in *EncryptionKeyRotationCronJobSpec) DeepCopy() *EncryptionKeyRotationCronJobSpec {
	if in == nil {
		return nil
	}
	out := new(EncryptionKeyRotationCronJobSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *EncryptionKeyRotationCronJobStatus) DeepCopyInto(out *EncryptionKeyRotationCronJobStatus) {
	*out = *in
	if in.Active != nil {
		in, out := &in.Active, &out.Active
		*out = new(v1.ObjectReference)
		**out = **in
	}
	if in.LastScheduleTime != nil {
		in, out := &in.LastScheduleTime, &out.LastScheduleTime
		*out = (*in).DeepCopy()
	}
	if in.LastSuccessfulTime != nil {
		in, out := &in.LastSuccessfulTime, &out.LastSuccessfulTime
		*out = (*in).DeepCopy()
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new EncryptionKeyRotationCronJobStatus.
func (in *EncryptionKeyRotationCronJobStatus) DeepCopy() *EncryptionKeyRotationCronJobStatus {
	if in == nil {
		return nil
	}
	out := new(EncryptionKeyRotationCronJobStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *EncryptionKeyRotationJob) DeepCopyInto(out *EncryptionKeyRotationJob) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new EncryptionKeyRotationJob.
func (in *EncryptionKeyRotationJob) DeepCopy() *EncryptionKeyRotationJob {
	if in == nil {
		return nil
	}
	out := new(EncryptionKeyRotationJob)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *EncryptionKeyRotationJob) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *EncryptionKeyRotationJobList) DeepCopyInto(out *EncryptionKeyRotationJobList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]EncryptionKeyRotationJob, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new EncryptionKeyRotationJobList.
func (in *EncryptionKeyRotationJobList) DeepCopy() *EncryptionKeyRotationJobList {
	if in == nil {
		return nil
	}
	out := new(EncryptionKeyRotationJobList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *EncryptionKeyRotationJobList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *EncryptionKeyRotationJobSpec) DeepCopyInto(out *EncryptionKeyRotationJobSpec) {
	*out = *in
	out.Target = in.Target
	if in.Timeout != nil {
		in, out := &in.Timeout, &out.Timeout
		*out = new(int64)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new EncryptionKeyRotationJobSpec.
func (in *EncryptionKeyRotationJobSpec) DeepCopy() *EncryptionKeyRotationJobSpec {
	if in == nil {
		return nil
	}
	out := new(EncryptionKeyRotationJobSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *EncryptionKeyRotationJobStatus) DeepCopyInto(out *EncryptionKeyRotationJobStatus) {
	*out = *in
	if in.Conditions != nil {
		in, out := &in.Conditions, &out.Conditions
		*out = make([]metav1.Condition, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.StartTime != nil {
		in, out := &in.StartTime, &out.StartTime
		*out = (*in).DeepCopy()
	}
	if in.CompletionTime != nil {
		in, out := &in.CompletionTime, &out.CompletionTime
		*out = (*in).DeepCopy()
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new EncryptionKeyRotationJobStatus.
func (in *EncryptionKeyRotationJobStatus) DeepCopy() *EncryptionKeyRotationJobStatus {
	if in == nil {
		return nil
	}
	out := new(EncryptionKeyRotationJobStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *EncryptionKeyRotationJobTemplateSpec) DeepCopyInto(out *EncryptionKeyRotationJobTemplateSpec) {
	*out = *in
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new EncryptionKeyRotationJobTemplateSpec.
func (in *EncryptionKeyRotationJobTemplateSpec) DeepCopy() *EncryptionKeyRotationJobTemplateSpec {
	if in == nil {
		return nil
	}
	out := new(EncryptionKeyRotationJobTemplateSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *NetworkFence) DeepCopyInto(out *NetworkFence) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new NetworkFence.
func (in *NetworkFence) DeepCopy() *NetworkFence {
	if in == nil {
		return nil
	}
	out := new(NetworkFence)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *NetworkFence) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *NetworkFenceList) DeepCopyInto(out *NetworkFenceList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]NetworkFence, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new NetworkFenceList.
func (in *NetworkFenceList) DeepCopy() *NetworkFenceList {
	if in == nil {
		return nil
	}
	out := new(NetworkFenceList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *NetworkFenceList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *NetworkFenceSpec) DeepCopyInto(out *NetworkFenceSpec) {
	*out = *in
	if in.Cidrs != nil {
		in, out := &in.Cidrs, &out.Cidrs
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	out.Secret = in.Secret
	if in.Parameters != nil {
		in, out := &in.Parameters, &out.Parameters
		*out = make(map[string]string, len(*in))
		for key, val := range *in {
			(*out)[key] = val
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new NetworkFenceSpec.
func (in *NetworkFenceSpec) DeepCopy() *NetworkFenceSpec {
	if in == nil {
		return nil
	}
	out := new(NetworkFenceSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *NetworkFenceStatus) DeepCopyInto(out *NetworkFenceStatus) {
	*out = *in
	if in.Conditions != nil {
		in, out := &in.Conditions, &out.Conditions
		*out = make([]metav1.Condition, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new NetworkFenceStatus.
func (in *NetworkFenceStatus) DeepCopy() *NetworkFenceStatus {
	if in == nil {
		return nil
	}
	out := new(NetworkFenceStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ReclaimSpaceCronJob) DeepCopyInto(out *ReclaimSpaceCronJob) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ReclaimSpaceCronJob.
func (in *ReclaimSpaceCronJob) DeepCopy() *ReclaimSpaceCronJob {
	if in == nil {
		return nil
	}
	out := new(ReclaimSpaceCronJob)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *ReclaimSpaceCronJob) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ReclaimSpaceCronJobList) DeepCopyInto(out *ReclaimSpaceCronJobList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]ReclaimSpaceCronJob, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ReclaimSpaceCronJobList.
func (in *ReclaimSpaceCronJobList) DeepCopy() *ReclaimSpaceCronJobList {
	if in == nil {
		return nil
	}
	out := new(ReclaimSpaceCronJobList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *ReclaimSpaceCronJobList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ReclaimSpaceCronJobSpec) DeepCopyInto(out *ReclaimSpaceCronJobSpec) {
	*out = *in
	if in.StartingDeadlineSeconds != nil {
		in, out := &in.StartingDeadlineSeconds, &out.StartingDeadlineSeconds
		*out = new(int64)
		**out = **in
	}
	if in.Suspend != nil {
		in, out := &in.Suspend, &out.Suspend
		*out = new(bool)
		**out = **in
	}
	in.JobSpec.DeepCopyInto(&out.JobSpec)
	if in.SuccessfulJobsHistoryLimit != nil {
		in, out := &in.SuccessfulJobsHistoryLimit, &out.SuccessfulJobsHistoryLimit
		*out = new(int32)
		**out = **in
	}
	if in.FailedJobsHistoryLimit != nil {
		in, out := &in.FailedJobsHistoryLimit, &out.FailedJobsHistoryLimit
		*out = new(int32)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ReclaimSpaceCronJobSpec.
func (in *ReclaimSpaceCronJobSpec) DeepCopy() *ReclaimSpaceCronJobSpec {
	if in == nil {
		return nil
	}
	out := new(ReclaimSpaceCronJobSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ReclaimSpaceCronJobStatus) DeepCopyInto(out *ReclaimSpaceCronJobStatus) {
	*out = *in
	if in.Active != nil {
		in, out := &in.Active, &out.Active
		*out = new(v1.ObjectReference)
		**out = **in
	}
	if in.LastScheduleTime != nil {
		in, out := &in.LastScheduleTime, &out.LastScheduleTime
		*out = (*in).DeepCopy()
	}
	if in.LastSuccessfulTime != nil {
		in, out := &in.LastSuccessfulTime, &out.LastSuccessfulTime
		*out = (*in).DeepCopy()
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ReclaimSpaceCronJobStatus.
func (in *ReclaimSpaceCronJobStatus) DeepCopy() *ReclaimSpaceCronJobStatus {
	if in == nil {
		return nil
	}
	out := new(ReclaimSpaceCronJobStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ReclaimSpaceJob) DeepCopyInto(out *ReclaimSpaceJob) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ReclaimSpaceJob.
func (in *ReclaimSpaceJob) DeepCopy() *ReclaimSpaceJob {
	if in == nil {
		return nil
	}
	out := new(ReclaimSpaceJob)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *ReclaimSpaceJob) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ReclaimSpaceJobList) DeepCopyInto(out *ReclaimSpaceJobList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]ReclaimSpaceJob, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ReclaimSpaceJobList.
func (in *ReclaimSpaceJobList) DeepCopy() *ReclaimSpaceJobList {
	if in == nil {
		return nil
	}
	out := new(ReclaimSpaceJobList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *ReclaimSpaceJobList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ReclaimSpaceJobSpec) DeepCopyInto(out *ReclaimSpaceJobSpec) {
	*out = *in
	out.Target = in.Target
	if in.Timeout != nil {
		in, out := &in.Timeout, &out.Timeout
		*out = new(int64)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ReclaimSpaceJobSpec.
func (in *ReclaimSpaceJobSpec) DeepCopy() *ReclaimSpaceJobSpec {
	if in == nil {
		return nil
	}
	out := new(ReclaimSpaceJobSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ReclaimSpaceJobStatus) DeepCopyInto(out *ReclaimSpaceJobStatus) {
	*out = *in
	if in.ReclaimedSpace != nil {
		in, out := &in.ReclaimedSpace, &out.ReclaimedSpace
		x := (*in).DeepCopy()
		*out = &x
	}
	if in.Conditions != nil {
		in, out := &in.Conditions, &out.Conditions
		*out = make([]metav1.Condition, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.StartTime != nil {
		in, out := &in.StartTime, &out.StartTime
		*out = (*in).DeepCopy()
	}
	if in.CompletionTime != nil {
		in, out := &in.CompletionTime, &out.CompletionTime
		*out = (*in).DeepCopy()
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ReclaimSpaceJobStatus.
func (in *ReclaimSpaceJobStatus) DeepCopy() *ReclaimSpaceJobStatus {
	if in == nil {
		return nil
	}
	out := new(ReclaimSpaceJobStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ReclaimSpaceJobTemplateSpec) DeepCopyInto(out *ReclaimSpaceJobTemplateSpec) {
	*out = *in
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ReclaimSpaceJobTemplateSpec.
func (in *ReclaimSpaceJobTemplateSpec) DeepCopy() *ReclaimSpaceJobTemplateSpec {
	if in == nil {
		return nil
	}
	out := new(ReclaimSpaceJobTemplateSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *SecretSpec) DeepCopyInto(out *SecretSpec) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new SecretSpec.
func (in *SecretSpec) DeepCopy() *SecretSpec {
	if in == nil {
		return nil
	}
	out := new(SecretSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *TargetSpec) DeepCopyInto(out *TargetSpec) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new TargetSpec.
func (in *TargetSpec) DeepCopy() *TargetSpec {
	if in == nil {
		return nil
	}
	out := new(TargetSpec)
	in.DeepCopyInto(out)
	return out
}
