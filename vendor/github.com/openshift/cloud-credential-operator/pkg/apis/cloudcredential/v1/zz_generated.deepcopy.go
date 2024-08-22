//go:build !ignore_autogenerated
// +build !ignore_autogenerated

// Code generated by deepcopy-gen. DO NOT EDIT.

package v1

import (
	runtime "k8s.io/apimachinery/pkg/runtime"
)

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *AWSProviderSpec) DeepCopyInto(out *AWSProviderSpec) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	if in.StatementEntries != nil {
		in, out := &in.StatementEntries, &out.StatementEntries
		*out = make([]StatementEntry, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new AWSProviderSpec.
func (in *AWSProviderSpec) DeepCopy() *AWSProviderSpec {
	if in == nil {
		return nil
	}
	out := new(AWSProviderSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *AWSProviderSpec) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *AWSProviderStatus) DeepCopyInto(out *AWSProviderStatus) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new AWSProviderStatus.
func (in *AWSProviderStatus) DeepCopy() *AWSProviderStatus {
	if in == nil {
		return nil
	}
	out := new(AWSProviderStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *AWSProviderStatus) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *AccessPolicy) DeepCopyInto(out *AccessPolicy) {
	*out = *in
	if in.Attributes != nil {
		in, out := &in.Attributes, &out.Attributes
		*out = make([]ResourceAttribute, len(*in))
		copy(*out, *in)
	}
	if in.Roles != nil {
		in, out := &in.Roles, &out.Roles
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new AccessPolicy.
func (in *AccessPolicy) DeepCopy() *AccessPolicy {
	if in == nil {
		return nil
	}
	out := new(AccessPolicy)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *AzureProviderSpec) DeepCopyInto(out *AzureProviderSpec) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	if in.RoleBindings != nil {
		in, out := &in.RoleBindings, &out.RoleBindings
		*out = make([]RoleBinding, len(*in))
		copy(*out, *in)
	}
	if in.Permissions != nil {
		in, out := &in.Permissions, &out.Permissions
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	if in.DataPermissions != nil {
		in, out := &in.DataPermissions, &out.DataPermissions
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new AzureProviderSpec.
func (in *AzureProviderSpec) DeepCopy() *AzureProviderSpec {
	if in == nil {
		return nil
	}
	out := new(AzureProviderSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *AzureProviderSpec) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *AzureProviderStatus) DeepCopyInto(out *AzureProviderStatus) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new AzureProviderStatus.
func (in *AzureProviderStatus) DeepCopy() *AzureProviderStatus {
	if in == nil {
		return nil
	}
	out := new(AzureProviderStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *AzureProviderStatus) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *CredentialsRequest) DeepCopyInto(out *CredentialsRequest) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new CredentialsRequest.
func (in *CredentialsRequest) DeepCopy() *CredentialsRequest {
	if in == nil {
		return nil
	}
	out := new(CredentialsRequest)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *CredentialsRequest) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *CredentialsRequestCondition) DeepCopyInto(out *CredentialsRequestCondition) {
	*out = *in
	in.LastProbeTime.DeepCopyInto(&out.LastProbeTime)
	in.LastTransitionTime.DeepCopyInto(&out.LastTransitionTime)
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new CredentialsRequestCondition.
func (in *CredentialsRequestCondition) DeepCopy() *CredentialsRequestCondition {
	if in == nil {
		return nil
	}
	out := new(CredentialsRequestCondition)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *CredentialsRequestList) DeepCopyInto(out *CredentialsRequestList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]CredentialsRequest, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new CredentialsRequestList.
func (in *CredentialsRequestList) DeepCopy() *CredentialsRequestList {
	if in == nil {
		return nil
	}
	out := new(CredentialsRequestList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *CredentialsRequestList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *CredentialsRequestSpec) DeepCopyInto(out *CredentialsRequestSpec) {
	*out = *in
	out.SecretRef = in.SecretRef
	if in.ProviderSpec != nil {
		in, out := &in.ProviderSpec, &out.ProviderSpec
		*out = new(runtime.RawExtension)
		(*in).DeepCopyInto(*out)
	}
	if in.ServiceAccountNames != nil {
		in, out := &in.ServiceAccountNames, &out.ServiceAccountNames
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new CredentialsRequestSpec.
func (in *CredentialsRequestSpec) DeepCopy() *CredentialsRequestSpec {
	if in == nil {
		return nil
	}
	out := new(CredentialsRequestSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *CredentialsRequestStatus) DeepCopyInto(out *CredentialsRequestStatus) {
	*out = *in
	if in.LastSyncTimestamp != nil {
		in, out := &in.LastSyncTimestamp, &out.LastSyncTimestamp
		*out = (*in).DeepCopy()
	}
	if in.ProviderStatus != nil {
		in, out := &in.ProviderStatus, &out.ProviderStatus
		*out = new(runtime.RawExtension)
		(*in).DeepCopyInto(*out)
	}
	if in.Conditions != nil {
		in, out := &in.Conditions, &out.Conditions
		*out = make([]CredentialsRequestCondition, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new CredentialsRequestStatus.
func (in *CredentialsRequestStatus) DeepCopy() *CredentialsRequestStatus {
	if in == nil {
		return nil
	}
	out := new(CredentialsRequestStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *GCPProviderSpec) DeepCopyInto(out *GCPProviderSpec) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	if in.PredefinedRoles != nil {
		in, out := &in.PredefinedRoles, &out.PredefinedRoles
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	if in.Permissions != nil {
		in, out := &in.Permissions, &out.Permissions
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new GCPProviderSpec.
func (in *GCPProviderSpec) DeepCopy() *GCPProviderSpec {
	if in == nil {
		return nil
	}
	out := new(GCPProviderSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *GCPProviderSpec) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *GCPProviderStatus) DeepCopyInto(out *GCPProviderStatus) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new GCPProviderStatus.
func (in *GCPProviderStatus) DeepCopy() *GCPProviderStatus {
	if in == nil {
		return nil
	}
	out := new(GCPProviderStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *GCPProviderStatus) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *IBMCloudPowerVSProviderSpec) DeepCopyInto(out *IBMCloudPowerVSProviderSpec) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	if in.Policies != nil {
		in, out := &in.Policies, &out.Policies
		*out = make([]AccessPolicy, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new IBMCloudPowerVSProviderSpec.
func (in *IBMCloudPowerVSProviderSpec) DeepCopy() *IBMCloudPowerVSProviderSpec {
	if in == nil {
		return nil
	}
	out := new(IBMCloudPowerVSProviderSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *IBMCloudPowerVSProviderSpec) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *IBMCloudPowerVSProviderStatus) DeepCopyInto(out *IBMCloudPowerVSProviderStatus) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new IBMCloudPowerVSProviderStatus.
func (in *IBMCloudPowerVSProviderStatus) DeepCopy() *IBMCloudPowerVSProviderStatus {
	if in == nil {
		return nil
	}
	out := new(IBMCloudPowerVSProviderStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *IBMCloudPowerVSProviderStatus) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *IBMCloudProviderSpec) DeepCopyInto(out *IBMCloudProviderSpec) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	if in.Policies != nil {
		in, out := &in.Policies, &out.Policies
		*out = make([]AccessPolicy, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new IBMCloudProviderSpec.
func (in *IBMCloudProviderSpec) DeepCopy() *IBMCloudProviderSpec {
	if in == nil {
		return nil
	}
	out := new(IBMCloudProviderSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *IBMCloudProviderSpec) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *IBMCloudProviderStatus) DeepCopyInto(out *IBMCloudProviderStatus) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new IBMCloudProviderStatus.
func (in *IBMCloudProviderStatus) DeepCopy() *IBMCloudProviderStatus {
	if in == nil {
		return nil
	}
	out := new(IBMCloudProviderStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *IBMCloudProviderStatus) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *KubevirtProviderSpec) DeepCopyInto(out *KubevirtProviderSpec) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new KubevirtProviderSpec.
func (in *KubevirtProviderSpec) DeepCopy() *KubevirtProviderSpec {
	if in == nil {
		return nil
	}
	out := new(KubevirtProviderSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *KubevirtProviderSpec) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *KubevirtProviderStatus) DeepCopyInto(out *KubevirtProviderStatus) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new KubevirtProviderStatus.
func (in *KubevirtProviderStatus) DeepCopy() *KubevirtProviderStatus {
	if in == nil {
		return nil
	}
	out := new(KubevirtProviderStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *KubevirtProviderStatus) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *NutanixProviderSpec) DeepCopyInto(out *NutanixProviderSpec) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new NutanixProviderSpec.
func (in *NutanixProviderSpec) DeepCopy() *NutanixProviderSpec {
	if in == nil {
		return nil
	}
	out := new(NutanixProviderSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *NutanixProviderSpec) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *NutanixProviderStatus) DeepCopyInto(out *NutanixProviderStatus) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new NutanixProviderStatus.
func (in *NutanixProviderStatus) DeepCopy() *NutanixProviderStatus {
	if in == nil {
		return nil
	}
	out := new(NutanixProviderStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *NutanixProviderStatus) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *OpenStackProviderSpec) DeepCopyInto(out *OpenStackProviderSpec) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new OpenStackProviderSpec.
func (in *OpenStackProviderSpec) DeepCopy() *OpenStackProviderSpec {
	if in == nil {
		return nil
	}
	out := new(OpenStackProviderSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *OpenStackProviderSpec) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *OpenStackProviderStatus) DeepCopyInto(out *OpenStackProviderStatus) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new OpenStackProviderStatus.
func (in *OpenStackProviderStatus) DeepCopy() *OpenStackProviderStatus {
	if in == nil {
		return nil
	}
	out := new(OpenStackProviderStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *OpenStackProviderStatus) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *OvirtProviderSpec) DeepCopyInto(out *OvirtProviderSpec) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new OvirtProviderSpec.
func (in *OvirtProviderSpec) DeepCopy() *OvirtProviderSpec {
	if in == nil {
		return nil
	}
	out := new(OvirtProviderSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *OvirtProviderSpec) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *OvirtProviderStatus) DeepCopyInto(out *OvirtProviderStatus) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new OvirtProviderStatus.
func (in *OvirtProviderStatus) DeepCopy() *OvirtProviderStatus {
	if in == nil {
		return nil
	}
	out := new(OvirtProviderStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *OvirtProviderStatus) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ResourceAttribute) DeepCopyInto(out *ResourceAttribute) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ResourceAttribute.
func (in *ResourceAttribute) DeepCopy() *ResourceAttribute {
	if in == nil {
		return nil
	}
	out := new(ResourceAttribute)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *RoleBinding) DeepCopyInto(out *RoleBinding) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new RoleBinding.
func (in *RoleBinding) DeepCopy() *RoleBinding {
	if in == nil {
		return nil
	}
	out := new(RoleBinding)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *StatementEntry) DeepCopyInto(out *StatementEntry) {
	*out = *in
	if in.Action != nil {
		in, out := &in.Action, &out.Action
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	in.PolicyCondition.DeepCopyInto(&out.PolicyCondition)
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new StatementEntry.
func (in *StatementEntry) DeepCopy() *StatementEntry {
	if in == nil {
		return nil
	}
	out := new(StatementEntry)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *VSpherePermission) DeepCopyInto(out *VSpherePermission) {
	*out = *in
	if in.Privileges != nil {
		in, out := &in.Privileges, &out.Privileges
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new VSpherePermission.
func (in *VSpherePermission) DeepCopy() *VSpherePermission {
	if in == nil {
		return nil
	}
	out := new(VSpherePermission)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *VSphereProviderSpec) DeepCopyInto(out *VSphereProviderSpec) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	if in.Permissions != nil {
		in, out := &in.Permissions, &out.Permissions
		*out = make([]VSpherePermission, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new VSphereProviderSpec.
func (in *VSphereProviderSpec) DeepCopy() *VSphereProviderSpec {
	if in == nil {
		return nil
	}
	out := new(VSphereProviderSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *VSphereProviderSpec) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *VSphereProviderStatus) DeepCopyInto(out *VSphereProviderStatus) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new VSphereProviderStatus.
func (in *VSphereProviderStatus) DeepCopy() *VSphereProviderStatus {
	if in == nil {
		return nil
	}
	out := new(VSphereProviderStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *VSphereProviderStatus) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}
