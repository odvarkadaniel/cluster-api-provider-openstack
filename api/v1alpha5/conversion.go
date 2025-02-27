/*
Copyright 2021 The Kubernetes Authors.

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

package v1alpha5

import (
	conversion "k8s.io/apimachinery/pkg/conversion"
	utilconversion "sigs.k8s.io/cluster-api/util/conversion"
	ctrlconversion "sigs.k8s.io/controller-runtime/pkg/conversion"

	infrav1 "sigs.k8s.io/cluster-api-provider-openstack/api/v1alpha7"
)

var _ ctrlconversion.Convertible = &OpenStackCluster{}

func (r *OpenStackCluster) ConvertTo(dstRaw ctrlconversion.Hub) error {
	dst := dstRaw.(*infrav1.OpenStackCluster)

	if err := Convert_v1alpha5_OpenStackCluster_To_v1alpha7_OpenStackCluster(r, dst, nil); err != nil {
		return err
	}

	// Manually restore data.
	restored := &infrav1.OpenStackCluster{}
	if ok, err := utilconversion.UnmarshalData(r, restored); err != nil || !ok {
		return err
	}

	return nil
}

func (r *OpenStackCluster) ConvertFrom(srcRaw ctrlconversion.Hub) error {
	src := srcRaw.(*infrav1.OpenStackCluster)

	if err := Convert_v1alpha7_OpenStackCluster_To_v1alpha5_OpenStackCluster(src, r, nil); err != nil {
		return err
	}

	// Preserve Hub data on down-conversion except for metadata
	return utilconversion.MarshalData(src, r)
}

var _ ctrlconversion.Convertible = &OpenStackClusterList{}

func (r *OpenStackClusterList) ConvertTo(dstRaw ctrlconversion.Hub) error {
	dst := dstRaw.(*infrav1.OpenStackClusterList)

	return Convert_v1alpha5_OpenStackClusterList_To_v1alpha7_OpenStackClusterList(r, dst, nil)
}

func (r *OpenStackClusterList) ConvertFrom(srcRaw ctrlconversion.Hub) error {
	src := srcRaw.(*infrav1.OpenStackClusterList)

	return Convert_v1alpha7_OpenStackClusterList_To_v1alpha5_OpenStackClusterList(src, r, nil)
}

var _ ctrlconversion.Convertible = &OpenStackClusterTemplate{}

func (r *OpenStackClusterTemplate) ConvertTo(dstRaw ctrlconversion.Hub) error {
	dst := dstRaw.(*infrav1.OpenStackClusterTemplate)

	if err := Convert_v1alpha5_OpenStackClusterTemplate_To_v1alpha7_OpenStackClusterTemplate(r, dst, nil); err != nil {
		return err
	}

	// Manually restore data.
	restored := &infrav1.OpenStackClusterTemplate{}
	if ok, err := utilconversion.UnmarshalData(r, restored); err != nil || !ok {
		return err
	}

	return nil
}

func (r *OpenStackClusterTemplate) ConvertFrom(srcRaw ctrlconversion.Hub) error {
	src := srcRaw.(*infrav1.OpenStackClusterTemplate)

	if err := Convert_v1alpha7_OpenStackClusterTemplate_To_v1alpha5_OpenStackClusterTemplate(src, r, nil); err != nil {
		return err
	}

	// Preserve Hub data on down-conversion except for metadata
	return utilconversion.MarshalData(src, r)
}

var _ ctrlconversion.Convertible = &OpenStackMachine{}

func (r *OpenStackMachine) ConvertTo(dstRaw ctrlconversion.Hub) error {
	dst := dstRaw.(*infrav1.OpenStackMachine)

	if err := Convert_v1alpha5_OpenStackMachine_To_v1alpha7_OpenStackMachine(r, dst, nil); err != nil {
		return err
	}

	// Manually restore data.
	restored := &infrav1.OpenStackMachine{}
	if ok, err := utilconversion.UnmarshalData(r, restored); err != nil || !ok {
		return err
	}

	return nil
}

func (r *OpenStackMachine) ConvertFrom(srcRaw ctrlconversion.Hub) error {
	src := srcRaw.(*infrav1.OpenStackMachine)

	if err := Convert_v1alpha7_OpenStackMachine_To_v1alpha5_OpenStackMachine(src, r, nil); err != nil {
		return err
	}

	// Preserve Hub data on down-conversion except for metadata
	return utilconversion.MarshalData(src, r)
}

var _ ctrlconversion.Convertible = &OpenStackMachineList{}

func (r *OpenStackMachineList) ConvertTo(dstRaw ctrlconversion.Hub) error {
	dst := dstRaw.(*infrav1.OpenStackMachineList)

	return Convert_v1alpha5_OpenStackMachineList_To_v1alpha7_OpenStackMachineList(r, dst, nil)
}

func (r *OpenStackMachineList) ConvertFrom(srcRaw ctrlconversion.Hub) error {
	src := srcRaw.(*infrav1.OpenStackMachineList)

	return Convert_v1alpha7_OpenStackMachineList_To_v1alpha5_OpenStackMachineList(src, r, nil)
}

var _ ctrlconversion.Convertible = &OpenStackMachineTemplate{}

func (r *OpenStackMachineTemplate) ConvertTo(dstRaw ctrlconversion.Hub) error {
	dst := dstRaw.(*infrav1.OpenStackMachineTemplate)

	if err := Convert_v1alpha5_OpenStackMachineTemplate_To_v1alpha7_OpenStackMachineTemplate(r, dst, nil); err != nil {
		return err
	}

	// Manually restore data.
	restored := &infrav1.OpenStackMachineTemplate{}
	if ok, err := utilconversion.UnmarshalData(r, restored); err != nil || !ok {
		return err
	}

	return nil
}

func (r *OpenStackMachineTemplate) ConvertFrom(srcRaw ctrlconversion.Hub) error {
	src := srcRaw.(*infrav1.OpenStackMachineTemplate)

	if err := Convert_v1alpha7_OpenStackMachineTemplate_To_v1alpha5_OpenStackMachineTemplate(src, r, nil); err != nil {
		return err
	}

	// Preserve Hub data on down-conversion except for metadata
	return utilconversion.MarshalData(src, r)
}

var _ ctrlconversion.Convertible = &OpenStackMachineTemplateList{}

func (r *OpenStackMachineTemplateList) ConvertTo(dstRaw ctrlconversion.Hub) error {
	dst := dstRaw.(*infrav1.OpenStackMachineTemplateList)

	return Convert_v1alpha5_OpenStackMachineTemplateList_To_v1alpha7_OpenStackMachineTemplateList(r, dst, nil)
}

func (r *OpenStackMachineTemplateList) ConvertFrom(srcRaw ctrlconversion.Hub) error {
	src := srcRaw.(*infrav1.OpenStackMachineTemplateList)

	return Convert_v1alpha7_OpenStackMachineTemplateList_To_v1alpha5_OpenStackMachineTemplateList(src, r, nil)
}

func Convert_v1alpha7_OpenStackClusterSpec_To_v1alpha5_OpenStackClusterSpec(in *infrav1.OpenStackClusterSpec, out *OpenStackClusterSpec, s conversion.Scope) error {
	// Our new flag has no equivalent in v1alpha5
	return autoConvert_v1alpha7_OpenStackClusterSpec_To_v1alpha5_OpenStackClusterSpec(in, out, s)
}

func Convert_v1alpha7_LoadBalancer_To_v1alpha5_LoadBalancer(in *infrav1.LoadBalancer, out *LoadBalancer, s conversion.Scope) error {
	return autoConvert_v1alpha7_LoadBalancer_To_v1alpha5_LoadBalancer(in, out, s)
}

func Convert_v1alpha7_PortOpts_To_v1alpha5_PortOpts(in *infrav1.PortOpts, out *PortOpts, s conversion.Scope) error {
	// value specs and propagate uplink status have been added in v1alpha7 but have no equivalent in v1alpha5
	return autoConvert_v1alpha7_PortOpts_To_v1alpha5_PortOpts(in, out, s)
}

func Convert_Slice_v1alpha5_Network_To_Slice_v1alpha7_Network(in *[]Network, out *[]infrav1.Network, s conversion.Scope) error {
	*out = make([]infrav1.Network, len(*in))
	for i := range *in {
		if err := Convert_v1alpha5_Network_To_v1alpha7_Network(&(*in)[i], &(*out)[i], s); err != nil {
			return err
		}
	}
	return nil
}

func Convert_Slice_v1alpha7_Network_To_Slice_v1alpha5_Network(in *[]infrav1.Network, out *[]Network, s conversion.Scope) error {
	*out = make([]Network, len(*in))
	for i := range *in {
		if err := Convert_v1alpha7_Network_To_v1alpha5_Network(&(*in)[i], &(*out)[i], s); err != nil {
			return err
		}
	}
	return nil
}

func Convert_v1alpha5_OpenStackMachineSpec_To_v1alpha7_OpenStackMachineSpec(in *OpenStackMachineSpec, out *infrav1.OpenStackMachineSpec, s conversion.Scope) error {
	return autoConvert_v1alpha5_OpenStackMachineSpec_To_v1alpha7_OpenStackMachineSpec(in, out, s)
}

func Convert_v1alpha7_APIServerLoadBalancer_To_v1alpha5_APIServerLoadBalancer(in *infrav1.APIServerLoadBalancer, out *APIServerLoadBalancer, s conversion.Scope) error {
	// Provider was originally added in v1alpha7, but was backported to v1alpha6, but has no equivalent in v1alpha5
	return autoConvert_v1alpha7_APIServerLoadBalancer_To_v1alpha5_APIServerLoadBalancer(in, out, s)
}

func Convert_v1alpha5_PortOpts_To_v1alpha7_PortOpts(in *PortOpts, out *infrav1.PortOpts, s conversion.Scope) error {
	// SecurityGroups have been removed in v1alpha7.
	return autoConvert_v1alpha5_PortOpts_To_v1alpha7_PortOpts(in, out, s)
}

func Convert_v1alpha5_Instance_To_v1alpha7_BastionStatus(in *Instance, out *infrav1.BastionStatus, _ conversion.Scope) error {
	// BastionStatus is the same as Instance with unused fields removed
	out.ID = in.ID
	out.Name = in.Name
	out.SSHKeyName = in.SSHKeyName
	out.State = infrav1.InstanceState(in.State)
	out.IP = in.IP
	out.FloatingIP = in.FloatingIP
	return nil
}

func Convert_v1alpha7_BastionStatus_To_v1alpha5_Instance(in *infrav1.BastionStatus, out *Instance, _ conversion.Scope) error {
	// BastionStatus is the same as Instance with unused fields removed
	out.ID = in.ID
	out.Name = in.Name
	out.SSHKeyName = in.SSHKeyName
	out.State = InstanceState(in.State)
	out.IP = in.IP
	out.FloatingIP = in.FloatingIP
	return nil
}

func Convert_v1alpha5_Network_To_v1alpha7_Network(in *Network, out *infrav1.Network, s conversion.Scope) error {
	// PortOpts has been removed in v1alpha7
	return autoConvert_v1alpha5_Network_To_v1alpha7_Network(in, out, s)
}

func Convert_v1alpha5_Network_To_v1alpha7_NetworkStatus(in *Network, out *infrav1.NetworkStatus, _ conversion.Scope) error {
	out.ID = in.ID
	out.Name = in.Name
	out.Tags = in.Tags

	return nil
}

func Convert_v1alpha7_NetworkStatus_To_v1alpha5_Network(in *infrav1.NetworkStatus, out *Network, _ conversion.Scope) error {
	out.ID = in.ID
	out.Name = in.Name
	out.Tags = in.Tags

	return nil
}

func Convert_v1alpha5_SecurityGroupParam_To_v1alpha7_SecurityGroupFilter(in *SecurityGroupParam, out *infrav1.SecurityGroupFilter, _ conversion.Scope) error {
	// SecurityGroupParam is replaced by its contained SecurityGroupFilter in v1alpha7
	*out = infrav1.SecurityGroupFilter(in.Filter)
	if in.UUID != "" {
		out.ID = in.UUID
	}
	if in.Name != "" {
		out.Name = in.Name
	}
	return nil
}

func Convert_v1alpha7_SecurityGroupFilter_To_v1alpha5_SecurityGroupParam(in *infrav1.SecurityGroupFilter, out *SecurityGroupParam, _ conversion.Scope) error {
	// SecurityGroupParam is replaced by its contained SecurityGroupFilter in v1alpha7
	out.Filter = SecurityGroupFilter(*in)
	if in.ID != "" {
		out.UUID = in.ID
	}
	if in.Name != "" {
		out.Name = in.Name
	}
	return nil
}

func Convert_v1alpha5_SubnetParam_To_v1alpha7_SubnetFilter(in *SubnetParam, out *infrav1.SubnetFilter, _ conversion.Scope) error {
	*out = infrav1.SubnetFilter(in.Filter)
	if in.UUID != "" {
		out.ID = in.UUID
	}
	return nil
}

func Convert_v1alpha7_SubnetFilter_To_v1alpha5_SubnetParam(in *infrav1.SubnetFilter, out *SubnetParam, _ conversion.Scope) error {
	out.Filter = SubnetFilter(*in)
	out.UUID = in.ID

	return nil
}
