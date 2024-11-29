//go:build !ignore_autogenerated
// +build !ignore_autogenerated

/*
Copyright The Kubernetes Authors.

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

package v1beta1

import (
	specs_go "github.com/opencontainers/runtime-spec/specs-go"
	"k8s.io/apimachinery/pkg/apis/meta/v1"
	runtime "k8s.io/apimachinery/pkg/runtime"
)

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *AppArmor) DeepCopyInto(out *AppArmor) {
	*out = *in
	if in.Profiles != nil {
		in, out := &in.Profiles, &out.Profiles
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	if in.Executions != nil {
		in, out := &in.Executions, &out.Executions
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	if in.Files != nil {
		in, out := &in.Files, &out.Files
		*out = make([]File, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.Capabilities != nil {
		in, out := &in.Capabilities, &out.Capabilities
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	if in.Networks != nil {
		in, out := &in.Networks, &out.Networks
		*out = make([]Network, len(*in))
		copy(*out, *in)
	}
	if in.Ptraces != nil {
		in, out := &in.Ptraces, &out.Ptraces
		*out = make([]Ptrace, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.Signals != nil {
		in, out := &in.Signals, &out.Signals
		*out = make([]Signal, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.Unhandled != nil {
		in, out := &in.Unhandled, &out.Unhandled
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new AppArmor.
func (in *AppArmor) DeepCopy() *AppArmor {
	if in == nil {
		return nil
	}
	out := new(AppArmor)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ArmorProfile) DeepCopyInto(out *ArmorProfile) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ArmorProfile.
func (in *ArmorProfile) DeepCopy() *ArmorProfile {
	if in == nil {
		return nil
	}
	out := new(ArmorProfile)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *ArmorProfile) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ArmorProfileCondition) DeepCopyInto(out *ArmorProfileCondition) {
	*out = *in
	in.LastTransitionTime.DeepCopyInto(&out.LastTransitionTime)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ArmorProfileCondition.
func (in *ArmorProfileCondition) DeepCopy() *ArmorProfileCondition {
	if in == nil {
		return nil
	}
	out := new(ArmorProfileCondition)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ArmorProfileList) DeepCopyInto(out *ArmorProfileList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]ArmorProfile, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ArmorProfileList.
func (in *ArmorProfileList) DeepCopy() *ArmorProfileList {
	if in == nil {
		return nil
	}
	out := new(ArmorProfileList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *ArmorProfileList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ArmorProfileModel) DeepCopyInto(out *ArmorProfileModel) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Data.DeepCopyInto(&out.Data)
	in.Status.DeepCopyInto(&out.Status)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ArmorProfileModel.
func (in *ArmorProfileModel) DeepCopy() *ArmorProfileModel {
	if in == nil {
		return nil
	}
	out := new(ArmorProfileModel)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *ArmorProfileModel) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ArmorProfileModelCondition) DeepCopyInto(out *ArmorProfileModelCondition) {
	*out = *in
	in.LastTransitionTime.DeepCopyInto(&out.LastTransitionTime)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ArmorProfileModelCondition.
func (in *ArmorProfileModelCondition) DeepCopy() *ArmorProfileModelCondition {
	if in == nil {
		return nil
	}
	out := new(ArmorProfileModelCondition)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ArmorProfileModelData) DeepCopyInto(out *ArmorProfileModelData) {
	*out = *in
	in.DynamicResult.DeepCopyInto(&out.DynamicResult)
	out.StaticResult = in.StaticResult
	in.Profile.DeepCopyInto(&out.Profile)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ArmorProfileModelData.
func (in *ArmorProfileModelData) DeepCopy() *ArmorProfileModelData {
	if in == nil {
		return nil
	}
	out := new(ArmorProfileModelData)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ArmorProfileModelList) DeepCopyInto(out *ArmorProfileModelList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]ArmorProfileModel, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ArmorProfileModelList.
func (in *ArmorProfileModelList) DeepCopy() *ArmorProfileModelList {
	if in == nil {
		return nil
	}
	out := new(ArmorProfileModelList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *ArmorProfileModelList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ArmorProfileModelStatus) DeepCopyInto(out *ArmorProfileModelStatus) {
	*out = *in
	if in.Conditions != nil {
		in, out := &in.Conditions, &out.Conditions
		*out = make([]ArmorProfileModelCondition, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ArmorProfileModelStatus.
func (in *ArmorProfileModelStatus) DeepCopy() *ArmorProfileModelStatus {
	if in == nil {
		return nil
	}
	out := new(ArmorProfileModelStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ArmorProfileSpec) DeepCopyInto(out *ArmorProfileSpec) {
	*out = *in
	in.Target.DeepCopyInto(&out.Target)
	in.Profile.DeepCopyInto(&out.Profile)
	out.BehaviorModeling = in.BehaviorModeling
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ArmorProfileSpec.
func (in *ArmorProfileSpec) DeepCopy() *ArmorProfileSpec {
	if in == nil {
		return nil
	}
	out := new(ArmorProfileSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ArmorProfileStatus) DeepCopyInto(out *ArmorProfileStatus) {
	*out = *in
	if in.Conditions != nil {
		in, out := &in.Conditions, &out.Conditions
		*out = make([]ArmorProfileCondition, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ArmorProfileStatus.
func (in *ArmorProfileStatus) DeepCopy() *ArmorProfileStatus {
	if in == nil {
		return nil
	}
	out := new(ArmorProfileStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *AttackProtectionRules) DeepCopyInto(out *AttackProtectionRules) {
	*out = *in
	if in.Rules != nil {
		in, out := &in.Rules, &out.Rules
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	if in.Targets != nil {
		in, out := &in.Targets, &out.Targets
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new AttackProtectionRules.
func (in *AttackProtectionRules) DeepCopy() *AttackProtectionRules {
	if in == nil {
		return nil
	}
	out := new(AttackProtectionRules)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *BehaviorModeling) DeepCopyInto(out *BehaviorModeling) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new BehaviorModeling.
func (in *BehaviorModeling) DeepCopy() *BehaviorModeling {
	if in == nil {
		return nil
	}
	out := new(BehaviorModeling)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *BpfContent) DeepCopyInto(out *BpfContent) {
	*out = *in
	if in.Capabilities != nil {
		in, out := &in.Capabilities, &out.Capabilities
		*out = new(CapabilitiesContent)
		**out = **in
	}
	if in.Files != nil {
		in, out := &in.Files, &out.Files
		*out = make([]FileContent, len(*in))
		copy(*out, *in)
	}
	if in.Processes != nil {
		in, out := &in.Processes, &out.Processes
		*out = make([]FileContent, len(*in))
		copy(*out, *in)
	}
	if in.Networks != nil {
		in, out := &in.Networks, &out.Networks
		*out = make([]NetworkContent, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.Ptrace != nil {
		in, out := &in.Ptrace, &out.Ptrace
		*out = new(PtraceContent)
		**out = **in
	}
	if in.Mounts != nil {
		in, out := &in.Mounts, &out.Mounts
		*out = make([]MountContent, len(*in))
		copy(*out, *in)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new BpfContent.
func (in *BpfContent) DeepCopy() *BpfContent {
	if in == nil {
		return nil
	}
	out := new(BpfContent)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *BpfRawRules) DeepCopyInto(out *BpfRawRules) {
	*out = *in
	if in.Files != nil {
		in, out := &in.Files, &out.Files
		*out = make([]FileRule, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.Processes != nil {
		in, out := &in.Processes, &out.Processes
		*out = make([]FileRule, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.Network != nil {
		in, out := &in.Network, &out.Network
		*out = new(NetworkRule)
		(*in).DeepCopyInto(*out)
	}
	if in.Ptrace != nil {
		in, out := &in.Ptrace, &out.Ptrace
		*out = new(PtraceRule)
		(*in).DeepCopyInto(*out)
	}
	if in.Mounts != nil {
		in, out := &in.Mounts, &out.Mounts
		*out = make([]MountRule, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new BpfRawRules.
func (in *BpfRawRules) DeepCopy() *BpfRawRules {
	if in == nil {
		return nil
	}
	out := new(BpfRawRules)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *CapabilitiesContent) DeepCopyInto(out *CapabilitiesContent) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new CapabilitiesContent.
func (in *CapabilitiesContent) DeepCopy() *CapabilitiesContent {
	if in == nil {
		return nil
	}
	out := new(CapabilitiesContent)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *DynamicResult) DeepCopyInto(out *DynamicResult) {
	*out = *in
	if in.AppArmor != nil {
		in, out := &in.AppArmor, &out.AppArmor
		*out = new(AppArmor)
		(*in).DeepCopyInto(*out)
	}
	if in.Seccomp != nil {
		in, out := &in.Seccomp, &out.Seccomp
		*out = new(Seccomp)
		(*in).DeepCopyInto(*out)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new DynamicResult.
func (in *DynamicResult) DeepCopy() *DynamicResult {
	if in == nil {
		return nil
	}
	out := new(DynamicResult)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *EnhanceProtect) DeepCopyInto(out *EnhanceProtect) {
	*out = *in
	if in.HardeningRules != nil {
		in, out := &in.HardeningRules, &out.HardeningRules
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	if in.AttackProtectionRules != nil {
		in, out := &in.AttackProtectionRules, &out.AttackProtectionRules
		*out = make([]AttackProtectionRules, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.VulMitigationRules != nil {
		in, out := &in.VulMitigationRules, &out.VulMitigationRules
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	if in.AppArmorRawRules != nil {
		in, out := &in.AppArmorRawRules, &out.AppArmorRawRules
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	if in.BpfRawRules != nil {
		in, out := &in.BpfRawRules, &out.BpfRawRules
		*out = new(BpfRawRules)
		(*in).DeepCopyInto(*out)
	}
	if in.SyscallRawRules != nil {
		in, out := &in.SyscallRawRules, &out.SyscallRawRules
		*out = make([]specs_go.LinuxSyscall, len(*in))
		linuxSyscallDeepCopyInto(in, out)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new EnhanceProtect.
func (in *EnhanceProtect) DeepCopy() *EnhanceProtect {
	if in == nil {
		return nil
	}
	out := new(EnhanceProtect)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *File) DeepCopyInto(out *File) {
	*out = *in
	if in.Permissions != nil {
		in, out := &in.Permissions, &out.Permissions
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new File.
func (in *File) DeepCopy() *File {
	if in == nil {
		return nil
	}
	out := new(File)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *FileContent) DeepCopyInto(out *FileContent) {
	*out = *in
	out.Pattern = in.Pattern
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new FileContent.
func (in *FileContent) DeepCopy() *FileContent {
	if in == nil {
		return nil
	}
	out := new(FileContent)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *FileRule) DeepCopyInto(out *FileRule) {
	*out = *in
	if in.Permissions != nil {
		in, out := &in.Permissions, &out.Permissions
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new FileRule.
func (in *FileRule) DeepCopy() *FileRule {
	if in == nil {
		return nil
	}
	out := new(FileRule)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ModelingOptions) DeepCopyInto(out *ModelingOptions) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ModelingOptions.
func (in *ModelingOptions) DeepCopy() *ModelingOptions {
	if in == nil {
		return nil
	}
	out := new(ModelingOptions)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *MountContent) DeepCopyInto(out *MountContent) {
	*out = *in
	out.Pattern = in.Pattern
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new MountContent.
func (in *MountContent) DeepCopy() *MountContent {
	if in == nil {
		return nil
	}
	out := new(MountContent)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *MountRule) DeepCopyInto(out *MountRule) {
	*out = *in
	if in.Flags != nil {
		in, out := &in.Flags, &out.Flags
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new MountRule.
func (in *MountRule) DeepCopy() *MountRule {
	if in == nil {
		return nil
	}
	out := new(MountRule)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Network) DeepCopyInto(out *Network) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Network.
func (in *Network) DeepCopy() *Network {
	if in == nil {
		return nil
	}
	out := new(Network)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *NetworkAddress) DeepCopyInto(out *NetworkAddress) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new NetworkAddress.
func (in *NetworkAddress) DeepCopy() *NetworkAddress {
	if in == nil {
		return nil
	}
	out := new(NetworkAddress)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *NetworkContent) DeepCopyInto(out *NetworkContent) {
	*out = *in
	if in.Socket != nil {
		in, out := &in.Socket, &out.Socket
		*out = new(NetworkSocket)
		**out = **in
	}
	if in.Address != nil {
		in, out := &in.Address, &out.Address
		*out = new(NetworkAddress)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new NetworkContent.
func (in *NetworkContent) DeepCopy() *NetworkContent {
	if in == nil {
		return nil
	}
	out := new(NetworkContent)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *NetworkEgressRule) DeepCopyInto(out *NetworkEgressRule) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new NetworkEgressRule.
func (in *NetworkEgressRule) DeepCopy() *NetworkEgressRule {
	if in == nil {
		return nil
	}
	out := new(NetworkEgressRule)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *NetworkRule) DeepCopyInto(out *NetworkRule) {
	*out = *in
	if in.Sockets != nil {
		in, out := &in.Sockets, &out.Sockets
		*out = make([]NetworkSocketRule, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.Egresses != nil {
		in, out := &in.Egresses, &out.Egresses
		*out = make([]NetworkEgressRule, len(*in))
		copy(*out, *in)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new NetworkRule.
func (in *NetworkRule) DeepCopy() *NetworkRule {
	if in == nil {
		return nil
	}
	out := new(NetworkRule)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *NetworkSocket) DeepCopyInto(out *NetworkSocket) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new NetworkSocket.
func (in *NetworkSocket) DeepCopy() *NetworkSocket {
	if in == nil {
		return nil
	}
	out := new(NetworkSocket)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *NetworkSocketRule) DeepCopyInto(out *NetworkSocketRule) {
	*out = *in
	if in.Domains != nil {
		in, out := &in.Domains, &out.Domains
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	if in.Types != nil {
		in, out := &in.Types, &out.Types
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	if in.Protocols != nil {
		in, out := &in.Protocols, &out.Protocols
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new NetworkSocketRule.
func (in *NetworkSocketRule) DeepCopy() *NetworkSocketRule {
	if in == nil {
		return nil
	}
	out := new(NetworkSocketRule)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *PathPattern) DeepCopyInto(out *PathPattern) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new PathPattern.
func (in *PathPattern) DeepCopy() *PathPattern {
	if in == nil {
		return nil
	}
	out := new(PathPattern)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Policy) DeepCopyInto(out *Policy) {
	*out = *in
	if in.EnhanceProtect != nil {
		in, out := &in.EnhanceProtect, &out.EnhanceProtect
		*out = new(EnhanceProtect)
		(*in).DeepCopyInto(*out)
	}
	if in.ModelingOptions != nil {
		in, out := &in.ModelingOptions, &out.ModelingOptions
		*out = new(ModelingOptions)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Policy.
func (in *Policy) DeepCopy() *Policy {
	if in == nil {
		return nil
	}
	out := new(Policy)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Profile) DeepCopyInto(out *Profile) {
	*out = *in
	if in.BpfContent != nil {
		in, out := &in.BpfContent, &out.BpfContent
		*out = new(BpfContent)
		(*in).DeepCopyInto(*out)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Profile.
func (in *Profile) DeepCopy() *Profile {
	if in == nil {
		return nil
	}
	out := new(Profile)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Ptrace) DeepCopyInto(out *Ptrace) {
	*out = *in
	if in.Permissions != nil {
		in, out := &in.Permissions, &out.Permissions
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Ptrace.
func (in *Ptrace) DeepCopy() *Ptrace {
	if in == nil {
		return nil
	}
	out := new(Ptrace)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *PtraceContent) DeepCopyInto(out *PtraceContent) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new PtraceContent.
func (in *PtraceContent) DeepCopy() *PtraceContent {
	if in == nil {
		return nil
	}
	out := new(PtraceContent)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *PtraceRule) DeepCopyInto(out *PtraceRule) {
	*out = *in
	if in.Permissions != nil {
		in, out := &in.Permissions, &out.Permissions
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new PtraceRule.
func (in *PtraceRule) DeepCopy() *PtraceRule {
	if in == nil {
		return nil
	}
	out := new(PtraceRule)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Seccomp) DeepCopyInto(out *Seccomp) {
	*out = *in
	if in.Syscalls != nil {
		in, out := &in.Syscalls, &out.Syscalls
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Seccomp.
func (in *Seccomp) DeepCopy() *Seccomp {
	if in == nil {
		return nil
	}
	out := new(Seccomp)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Signal) DeepCopyInto(out *Signal) {
	*out = *in
	if in.Permissions != nil {
		in, out := &in.Permissions, &out.Permissions
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	if in.Signals != nil {
		in, out := &in.Signals, &out.Signals
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Signal.
func (in *Signal) DeepCopy() *Signal {
	if in == nil {
		return nil
	}
	out := new(Signal)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *StaticResult) DeepCopyInto(out *StaticResult) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new StaticResult.
func (in *StaticResult) DeepCopy() *StaticResult {
	if in == nil {
		return nil
	}
	out := new(StaticResult)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Target) DeepCopyInto(out *Target) {
	*out = *in
	if in.Containers != nil {
		in, out := &in.Containers, &out.Containers
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	if in.Selector != nil {
		in, out := &in.Selector, &out.Selector
		*out = new(v1.LabelSelector)
		(*in).DeepCopyInto(*out)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Target.
func (in *Target) DeepCopy() *Target {
	if in == nil {
		return nil
	}
	out := new(Target)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *VarmorClusterPolicy) DeepCopyInto(out *VarmorClusterPolicy) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new VarmorClusterPolicy.
func (in *VarmorClusterPolicy) DeepCopy() *VarmorClusterPolicy {
	if in == nil {
		return nil
	}
	out := new(VarmorClusterPolicy)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *VarmorClusterPolicy) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *VarmorClusterPolicyList) DeepCopyInto(out *VarmorClusterPolicyList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]VarmorClusterPolicy, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new VarmorClusterPolicyList.
func (in *VarmorClusterPolicyList) DeepCopy() *VarmorClusterPolicyList {
	if in == nil {
		return nil
	}
	out := new(VarmorClusterPolicyList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *VarmorClusterPolicyList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *VarmorPolicy) DeepCopyInto(out *VarmorPolicy) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new VarmorPolicy.
func (in *VarmorPolicy) DeepCopy() *VarmorPolicy {
	if in == nil {
		return nil
	}
	out := new(VarmorPolicy)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *VarmorPolicy) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *VarmorPolicyCondition) DeepCopyInto(out *VarmorPolicyCondition) {
	*out = *in
	in.LastTransitionTime.DeepCopyInto(&out.LastTransitionTime)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new VarmorPolicyCondition.
func (in *VarmorPolicyCondition) DeepCopy() *VarmorPolicyCondition {
	if in == nil {
		return nil
	}
	out := new(VarmorPolicyCondition)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *VarmorPolicyList) DeepCopyInto(out *VarmorPolicyList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]VarmorPolicy, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new VarmorPolicyList.
func (in *VarmorPolicyList) DeepCopy() *VarmorPolicyList {
	if in == nil {
		return nil
	}
	out := new(VarmorPolicyList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *VarmorPolicyList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *VarmorPolicySpec) DeepCopyInto(out *VarmorPolicySpec) {
	*out = *in
	in.Target.DeepCopyInto(&out.Target)
	in.Policy.DeepCopyInto(&out.Policy)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new VarmorPolicySpec.
func (in *VarmorPolicySpec) DeepCopy() *VarmorPolicySpec {
	if in == nil {
		return nil
	}
	out := new(VarmorPolicySpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *VarmorPolicyStatus) DeepCopyInto(out *VarmorPolicyStatus) {
	*out = *in
	if in.Conditions != nil {
		in, out := &in.Conditions, &out.Conditions
		*out = make([]VarmorPolicyCondition, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new VarmorPolicyStatus.
func (in *VarmorPolicyStatus) DeepCopy() *VarmorPolicyStatus {
	if in == nil {
		return nil
	}
	out := new(VarmorPolicyStatus)
	in.DeepCopyInto(out)
	return out
}
