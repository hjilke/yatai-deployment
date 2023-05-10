//go:build !ignore_autogenerated
// +build !ignore_autogenerated

/*
Copyright 2022.

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

package v1alpha2

import (
	"github.com/bentoml/yatai-deployment/apis/serving/common"
	"github.com/bentoml/yatai-schemas/modelschemas"
	runtime "k8s.io/apimachinery/pkg/runtime"
)

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *BentoDeployment) DeepCopyInto(out *BentoDeployment) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new BentoDeployment.
func (in *BentoDeployment) DeepCopy() *BentoDeployment {
	if in == nil {
		return nil
	}
	out := new(BentoDeployment)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *BentoDeployment) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *BentoDeploymentIngressSpec) DeepCopyInto(out *BentoDeploymentIngressSpec) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new BentoDeploymentIngressSpec.
func (in *BentoDeploymentIngressSpec) DeepCopy() *BentoDeploymentIngressSpec {
	if in == nil {
		return nil
	}
	out := new(BentoDeploymentIngressSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *BentoDeploymentList) DeepCopyInto(out *BentoDeploymentList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]BentoDeployment, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new BentoDeploymentList.
func (in *BentoDeploymentList) DeepCopy() *BentoDeploymentList {
	if in == nil {
		return nil
	}
	out := new(BentoDeploymentList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *BentoDeploymentList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *BentoDeploymentRunnerSpec) DeepCopyInto(out *BentoDeploymentRunnerSpec) {
	*out = *in
	if in.Resources != nil {
		in, out := &in.Resources, &out.Resources
		*out = new(common.Resources)
		(*in).DeepCopyInto(*out)
	}
	if in.Autoscaling != nil {
		in, out := &in.Autoscaling, &out.Autoscaling
		*out = (*in).DeepCopy()
	}
	if in.Envs != nil {
		in, out := &in.Envs, &out.Envs
		*out = new([]modelschemas.LabelItemSchema)
		if **in != nil {
			in, out := *in, *out
			*out = make([]modelschemas.LabelItemSchema, len(*in))
			for i := range *in {
				(*in)[i].DeepCopyInto(&(*out)[i])
			}
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new BentoDeploymentRunnerSpec.
func (in *BentoDeploymentRunnerSpec) DeepCopy() *BentoDeploymentRunnerSpec {
	if in == nil {
		return nil
	}
	out := new(BentoDeploymentRunnerSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *BentoDeploymentSpec) DeepCopyInto(out *BentoDeploymentSpec) {
	*out = *in
	if in.Resources != nil {
		in, out := &in.Resources, &out.Resources
		*out = new(common.Resources)
		(*in).DeepCopyInto(*out)
	}
	if in.Autoscaling != nil {
		in, out := &in.Autoscaling, &out.Autoscaling
		*out = (*in).DeepCopy()
	}
	if in.Envs != nil {
		in, out := &in.Envs, &out.Envs
		*out = new([]modelschemas.LabelItemSchema)
		if **in != nil {
			in, out := *in, *out
			*out = make([]modelschemas.LabelItemSchema, len(*in))
			for i := range *in {
				(*in)[i].DeepCopyInto(&(*out)[i])
			}
		}
	}
	if in.Runners != nil {
		in, out := &in.Runners, &out.Runners
		*out = make([]BentoDeploymentRunnerSpec, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	out.Ingress = in.Ingress
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new BentoDeploymentSpec.
func (in *BentoDeploymentSpec) DeepCopy() *BentoDeploymentSpec {
	if in == nil {
		return nil
	}
	out := new(BentoDeploymentSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *BentoDeploymentStatus) DeepCopyInto(out *BentoDeploymentStatus) {
	*out = *in
	if in.PodSelector != nil {
		in, out := &in.PodSelector, &out.PodSelector
		*out = make(map[string]string, len(*in))
		for key, val := range *in {
			(*out)[key] = val
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new BentoDeploymentStatus.
func (in *BentoDeploymentStatus) DeepCopy() *BentoDeploymentStatus {
	if in == nil {
		return nil
	}
	out := new(BentoDeploymentStatus)
	in.DeepCopyInto(out)
	return out
}
