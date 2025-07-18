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

package v1alpha4

import (
	apiconversion "k8s.io/apimachinery/pkg/conversion"
	"sigs.k8s.io/controller-runtime/pkg/conversion"

	infraexpv1 "sigs.k8s.io/cluster-api/test/infrastructure/docker/exp/api/v1beta2"
	utilconversion "sigs.k8s.io/cluster-api/util/conversion"
)

func (src *DockerMachinePool) ConvertTo(dstRaw conversion.Hub) error {
	dst := dstRaw.(*infraexpv1.DockerMachinePool)

	if err := Convert_v1alpha4_DockerMachinePool_To_v1beta2_DockerMachinePool(src, dst, nil); err != nil {
		return err
	}

	// Manually restore data.
	restored := &infraexpv1.DockerMachinePool{}
	if ok, err := utilconversion.UnmarshalData(src, restored); err != nil || !ok {
		return err
	}

	dst.Status.InfrastructureMachineKind = restored.Status.InfrastructureMachineKind

	return nil
}

func (dst *DockerMachinePool) ConvertFrom(srcRaw conversion.Hub) error {
	src := srcRaw.(*infraexpv1.DockerMachinePool)

	if err := Convert_v1beta2_DockerMachinePool_To_v1alpha4_DockerMachinePool(src, dst, nil); err != nil {
		return err
	}

	// Preserve Hub data on down-conversion except for metadata
	return utilconversion.MarshalData(src, dst)
}

func Convert_v1beta2_DockerMachinePoolStatus_To_v1alpha4_DockerMachinePoolStatus(in *infraexpv1.DockerMachinePoolStatus, out *DockerMachinePoolStatus, s apiconversion.Scope) error {
	// NOTE: custom conversion func is required because Status.InfrastructureMachineKind has been added in v1beta1.
	return autoConvert_v1beta2_DockerMachinePoolStatus_To_v1alpha4_DockerMachinePoolStatus(in, out, s)
}
