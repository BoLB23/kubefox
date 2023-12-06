//go:build !ignore_autogenerated

// Code generated by controller-gen. DO NOT EDIT.

package v1alpha1

import (
	"github.com/xigxog/kubefox/api"
	"github.com/xigxog/kubefox/api/kubernetes"
	runtime "k8s.io/apimachinery/pkg/runtime"
)

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *App) DeepCopyInto(out *App) {
	*out = *in
	out.App = in.App
	in.CommitTime.DeepCopyInto(&out.CommitTime)
	if in.Components != nil {
		in, out := &in.Components, &out.Components
		*out = make(map[string]*Component, len(*in))
		for key, val := range *in {
			var outVal *Component
			if val == nil {
				(*out)[key] = nil
			} else {
				inVal := (*in)[key]
				in, out := &inVal, &outVal
				*out = new(Component)
				(*in).DeepCopyInto(*out)
			}
			(*out)[key] = outVal
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new App.
func (in *App) DeepCopy() *App {
	if in == nil {
		return nil
	}
	out := new(App)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *AppDeployment) DeepCopyInto(out *AppDeployment) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	out.Status = in.Status
	in.Details.DeepCopyInto(&out.Details)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new AppDeployment.
func (in *AppDeployment) DeepCopy() *AppDeployment {
	if in == nil {
		return nil
	}
	out := new(AppDeployment)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *AppDeployment) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *AppDeploymentDetails) DeepCopyInto(out *AppDeploymentDetails) {
	*out = *in
	out.App = in.App
	if in.Components != nil {
		in, out := &in.Components, &out.Components
		*out = make(map[string]api.Details, len(*in))
		for key, val := range *in {
			(*out)[key] = val
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new AppDeploymentDetails.
func (in *AppDeploymentDetails) DeepCopy() *AppDeploymentDetails {
	if in == nil {
		return nil
	}
	out := new(AppDeploymentDetails)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *AppDeploymentList) DeepCopyInto(out *AppDeploymentList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]AppDeployment, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new AppDeploymentList.
func (in *AppDeploymentList) DeepCopy() *AppDeploymentList {
	if in == nil {
		return nil
	}
	out := new(AppDeploymentList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *AppDeploymentList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *AppDeploymentSpec) DeepCopyInto(out *AppDeploymentSpec) {
	*out = *in
	in.App.DeepCopyInto(&out.App)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new AppDeploymentSpec.
func (in *AppDeploymentSpec) DeepCopy() *AppDeploymentSpec {
	if in == nil {
		return nil
	}
	out := new(AppDeploymentSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *AppDeploymentStatus) DeepCopyInto(out *AppDeploymentStatus) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new AppDeploymentStatus.
func (in *AppDeploymentStatus) DeepCopy() *AppDeploymentStatus {
	if in == nil {
		return nil
	}
	out := new(AppDeploymentStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *BrokerSpec) DeepCopyInto(out *BrokerSpec) {
	*out = *in
	in.PodSpec.DeepCopyInto(&out.PodSpec)
	in.ContainerSpec.DeepCopyInto(&out.ContainerSpec)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new BrokerSpec.
func (in *BrokerSpec) DeepCopy() *BrokerSpec {
	if in == nil {
		return nil
	}
	out := new(BrokerSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ClusterEnvSpec) DeepCopyInto(out *ClusterEnvSpec) {
	*out = *in
	if in.ReleasePolicy != nil {
		in, out := &in.ReleasePolicy, &out.ReleasePolicy
		*out = new(EnvReleasePolicy)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ClusterEnvSpec.
func (in *ClusterEnvSpec) DeepCopy() *ClusterEnvSpec {
	if in == nil {
		return nil
	}
	out := new(ClusterEnvSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ClusterVirtualEnv) DeepCopyInto(out *ClusterVirtualEnv) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Data.DeepCopyInto(&out.Data)
	in.Details.DeepCopyInto(&out.Details)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ClusterVirtualEnv.
func (in *ClusterVirtualEnv) DeepCopy() *ClusterVirtualEnv {
	if in == nil {
		return nil
	}
	out := new(ClusterVirtualEnv)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *ClusterVirtualEnv) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ClusterVirtualEnvList) DeepCopyInto(out *ClusterVirtualEnvList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]ClusterVirtualEnv, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ClusterVirtualEnvList.
func (in *ClusterVirtualEnvList) DeepCopy() *ClusterVirtualEnvList {
	if in == nil {
		return nil
	}
	out := new(ClusterVirtualEnvList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *ClusterVirtualEnvList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Component) DeepCopyInto(out *Component) {
	*out = *in
	in.ComponentDefinition.DeepCopyInto(&out.ComponentDefinition)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Component.
func (in *Component) DeepCopy() *Component {
	if in == nil {
		return nil
	}
	out := new(Component)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ComponentRef) DeepCopyInto(out *ComponentRef) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ComponentRef.
func (in *ComponentRef) DeepCopy() *ComponentRef {
	if in == nil {
		return nil
	}
	out := new(ComponentRef)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ComponentSpec) DeepCopyInto(out *ComponentSpec) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ComponentSpec.
func (in *ComponentSpec) DeepCopy() *ComponentSpec {
	if in == nil {
		return nil
	}
	out := new(ComponentSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *ComponentSpec) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ComponentSpecList) DeepCopyInto(out *ComponentSpecList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]ComponentSpec, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ComponentSpecList.
func (in *ComponentSpecList) DeepCopy() *ComponentSpecList {
	if in == nil {
		return nil
	}
	out := new(ComponentSpecList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *ComponentSpecList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ComponentSpecSelector) DeepCopyInto(out *ComponentSpecSelector) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ComponentSpecSelector.
func (in *ComponentSpecSelector) DeepCopy() *ComponentSpecSelector {
	if in == nil {
		return nil
	}
	out := new(ComponentSpecSelector)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ComponentSpecSpec) DeepCopyInto(out *ComponentSpecSpec) {
	*out = *in
	out.Selector = in.Selector
	in.PodSpec.DeepCopyInto(&out.PodSpec)
	in.ContainerSpec.DeepCopyInto(&out.ContainerSpec)
	out.Logger = in.Logger
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ComponentSpecSpec.
func (in *ComponentSpecSpec) DeepCopy() *ComponentSpecSpec {
	if in == nil {
		return nil
	}
	out := new(ComponentSpecSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ComponentSpecStatus) DeepCopyInto(out *ComponentSpecStatus) {
	*out = *in
	if in.Matched != nil {
		in, out := &in.Matched, &out.Matched
		*out = make([]ComponentRef, len(*in))
		copy(*out, *in)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ComponentSpecStatus.
func (in *ComponentSpecStatus) DeepCopy() *ComponentSpecStatus {
	if in == nil {
		return nil
	}
	out := new(ComponentSpecStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ComponentStatus) DeepCopyInto(out *ComponentStatus) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ComponentStatus.
func (in *ComponentStatus) DeepCopy() *ComponentStatus {
	if in == nil {
		return nil
	}
	out := new(ComponentStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *EnvAdapter) DeepCopyInto(out *EnvAdapter) {
	*out = *in
	out.URL = in.URL
	if in.Headers != nil {
		in, out := &in.Headers, &out.Headers
		*out = make(map[string]*kubernetes.StringOrSecret, len(*in))
		for key, val := range *in {
			var outVal *kubernetes.StringOrSecret
			if val == nil {
				(*out)[key] = nil
			} else {
				inVal := (*in)[key]
				in, out := &inVal, &outVal
				*out = new(kubernetes.StringOrSecret)
				**out = **in
			}
			(*out)[key] = outVal
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new EnvAdapter.
func (in *EnvAdapter) DeepCopy() *EnvAdapter {
	if in == nil {
		return nil
	}
	out := new(EnvAdapter)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *EnvData) DeepCopyInto(out *EnvData) {
	*out = *in
	if in.Vars != nil {
		in, out := &in.Vars, &out.Vars
		*out = make(map[string]*api.Val, len(*in))
		for key, val := range *in {
			var outVal *api.Val
			if val == nil {
				(*out)[key] = nil
			} else {
				inVal := (*in)[key]
				in, out := &inVal, &outVal
				*out = new(api.Val)
				(*in).DeepCopyInto(*out)
			}
			(*out)[key] = outVal
		}
	}
	if in.Adapters != nil {
		in, out := &in.Adapters, &out.Adapters
		*out = make(map[string]*EnvAdapter, len(*in))
		for key, val := range *in {
			var outVal *EnvAdapter
			if val == nil {
				(*out)[key] = nil
			} else {
				inVal := (*in)[key]
				in, out := &inVal, &outVal
				*out = new(EnvAdapter)
				(*in).DeepCopyInto(*out)
			}
			(*out)[key] = outVal
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new EnvData.
func (in *EnvData) DeepCopy() *EnvData {
	if in == nil {
		return nil
	}
	out := new(EnvData)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *EnvDetails) DeepCopyInto(out *EnvDetails) {
	*out = *in
	out.Details = in.Details
	if in.Vars != nil {
		in, out := &in.Vars, &out.Vars
		*out = make(map[string]api.Details, len(*in))
		for key, val := range *in {
			(*out)[key] = val
		}
	}
	if in.Adapters != nil {
		in, out := &in.Adapters, &out.Adapters
		*out = make(map[string]api.Details, len(*in))
		for key, val := range *in {
			(*out)[key] = val
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new EnvDetails.
func (in *EnvDetails) DeepCopy() *EnvDetails {
	if in == nil {
		return nil
	}
	out := new(EnvDetails)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *EnvReleasePolicy) DeepCopyInto(out *EnvReleasePolicy) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new EnvReleasePolicy.
func (in *EnvReleasePolicy) DeepCopy() *EnvReleasePolicy {
	if in == nil {
		return nil
	}
	out := new(EnvReleasePolicy)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *EnvSnapshotData) DeepCopyInto(out *EnvSnapshotData) {
	*out = *in
	in.EnvData.DeepCopyInto(&out.EnvData)
	out.Source = in.Source
	in.SnapshotTime.DeepCopyInto(&out.SnapshotTime)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new EnvSnapshotData.
func (in *EnvSnapshotData) DeepCopy() *EnvSnapshotData {
	if in == nil {
		return nil
	}
	out := new(EnvSnapshotData)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *EnvSource) DeepCopyInto(out *EnvSource) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new EnvSource.
func (in *EnvSource) DeepCopy() *EnvSource {
	if in == nil {
		return nil
	}
	out := new(EnvSource)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *EnvSpec) DeepCopyInto(out *EnvSpec) {
	*out = *in
	if in.ReleasePolicy != nil {
		in, out := &in.ReleasePolicy, &out.ReleasePolicy
		*out = new(EnvReleasePolicy)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new EnvSpec.
func (in *EnvSpec) DeepCopy() *EnvSpec {
	if in == nil {
		return nil
	}
	out := new(EnvSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *EventsSpec) DeepCopyInto(out *EventsSpec) {
	*out = *in
	out.MaxSize = in.MaxSize.DeepCopy()
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new EventsSpec.
func (in *EventsSpec) DeepCopy() *EventsSpec {
	if in == nil {
		return nil
	}
	out := new(EventsSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *HTTPSrvPorts) DeepCopyInto(out *HTTPSrvPorts) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new HTTPSrvPorts.
func (in *HTTPSrvPorts) DeepCopy() *HTTPSrvPorts {
	if in == nil {
		return nil
	}
	out := new(HTTPSrvPorts)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *HTTPSrvService) DeepCopyInto(out *HTTPSrvService) {
	*out = *in
	out.Ports = in.Ports
	if in.Labels != nil {
		in, out := &in.Labels, &out.Labels
		*out = make(map[string]string, len(*in))
		for key, val := range *in {
			(*out)[key] = val
		}
	}
	if in.Annotations != nil {
		in, out := &in.Annotations, &out.Annotations
		*out = make(map[string]string, len(*in))
		for key, val := range *in {
			(*out)[key] = val
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new HTTPSrvService.
func (in *HTTPSrvService) DeepCopy() *HTTPSrvService {
	if in == nil {
		return nil
	}
	out := new(HTTPSrvService)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *HTTPSrvSpec) DeepCopyInto(out *HTTPSrvSpec) {
	*out = *in
	in.PodSpec.DeepCopyInto(&out.PodSpec)
	in.ContainerSpec.DeepCopyInto(&out.ContainerSpec)
	in.Service.DeepCopyInto(&out.Service)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new HTTPSrvSpec.
func (in *HTTPSrvSpec) DeepCopy() *HTTPSrvSpec {
	if in == nil {
		return nil
	}
	out := new(HTTPSrvSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *NATSSpec) DeepCopyInto(out *NATSSpec) {
	*out = *in
	in.PodSpec.DeepCopyInto(&out.PodSpec)
	in.ContainerSpec.DeepCopyInto(&out.ContainerSpec)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new NATSSpec.
func (in *NATSSpec) DeepCopy() *NATSSpec {
	if in == nil {
		return nil
	}
	out := new(NATSSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Platform) DeepCopyInto(out *Platform) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
	out.Details = in.Details
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Platform.
func (in *Platform) DeepCopy() *Platform {
	if in == nil {
		return nil
	}
	out := new(Platform)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *Platform) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *PlatformDetails) DeepCopyInto(out *PlatformDetails) {
	*out = *in
	out.Details = in.Details
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new PlatformDetails.
func (in *PlatformDetails) DeepCopy() *PlatformDetails {
	if in == nil {
		return nil
	}
	out := new(PlatformDetails)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *PlatformList) DeepCopyInto(out *PlatformList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]Platform, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new PlatformList.
func (in *PlatformList) DeepCopy() *PlatformList {
	if in == nil {
		return nil
	}
	out := new(PlatformList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *PlatformList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *PlatformSpec) DeepCopyInto(out *PlatformSpec) {
	*out = *in
	in.Events.DeepCopyInto(&out.Events)
	in.Broker.DeepCopyInto(&out.Broker)
	in.HTTPSrv.DeepCopyInto(&out.HTTPSrv)
	in.NATS.DeepCopyInto(&out.NATS)
	out.Logger = in.Logger
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new PlatformSpec.
func (in *PlatformSpec) DeepCopy() *PlatformSpec {
	if in == nil {
		return nil
	}
	out := new(PlatformSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *PlatformStatus) DeepCopyInto(out *PlatformStatus) {
	*out = *in
	if in.Components != nil {
		in, out := &in.Components, &out.Components
		*out = make([]ComponentStatus, len(*in))
		copy(*out, *in)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new PlatformStatus.
func (in *PlatformStatus) DeepCopy() *PlatformStatus {
	if in == nil {
		return nil
	}
	out := new(PlatformStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Release) DeepCopyInto(out *Release) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Release.
func (in *Release) DeepCopy() *Release {
	if in == nil {
		return nil
	}
	out := new(Release)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *Release) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ReleaseAppDeployment) DeepCopyInto(out *ReleaseAppDeployment) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ReleaseAppDeployment.
func (in *ReleaseAppDeployment) DeepCopy() *ReleaseAppDeployment {
	if in == nil {
		return nil
	}
	out := new(ReleaseAppDeployment)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ReleaseHistoryLimit) DeepCopyInto(out *ReleaseHistoryLimit) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ReleaseHistoryLimit.
func (in *ReleaseHistoryLimit) DeepCopy() *ReleaseHistoryLimit {
	if in == nil {
		return nil
	}
	out := new(ReleaseHistoryLimit)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ReleaseList) DeepCopyInto(out *ReleaseList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]Release, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ReleaseList.
func (in *ReleaseList) DeepCopy() *ReleaseList {
	if in == nil {
		return nil
	}
	out := new(ReleaseList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *ReleaseList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ReleaseSpec) DeepCopyInto(out *ReleaseSpec) {
	*out = *in
	out.AppDeployment = in.AppDeployment
	if in.HistoryLimit != nil {
		in, out := &in.HistoryLimit, &out.HistoryLimit
		*out = new(ReleaseHistoryLimit)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ReleaseSpec.
func (in *ReleaseSpec) DeepCopy() *ReleaseSpec {
	if in == nil {
		return nil
	}
	out := new(ReleaseSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ReleaseStatus) DeepCopyInto(out *ReleaseStatus) {
	*out = *in
	if in.Current != nil {
		in, out := &in.Current, &out.Current
		*out = new(ReleaseStatusEntry)
		(*in).DeepCopyInto(*out)
	}
	if in.Requested != nil {
		in, out := &in.Requested, &out.Requested
		*out = new(ReleaseStatusEntry)
		(*in).DeepCopyInto(*out)
	}
	if in.History != nil {
		in, out := &in.History, &out.History
		*out = make([]ReleaseStatusEntry, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ReleaseStatus.
func (in *ReleaseStatus) DeepCopy() *ReleaseStatus {
	if in == nil {
		return nil
	}
	out := new(ReleaseStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ReleaseStatusEntry) DeepCopyInto(out *ReleaseStatusEntry) {
	*out = *in
	out.AppDeployment = in.AppDeployment
	in.RequestTime.DeepCopyInto(&out.RequestTime)
	if in.AvailableTime != nil {
		in, out := &in.AvailableTime, &out.AvailableTime
		*out = (*in).DeepCopy()
	}
	if in.ArchiveTime != nil {
		in, out := &in.ArchiveTime, &out.ArchiveTime
		*out = (*in).DeepCopy()
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ReleaseStatusEntry.
func (in *ReleaseStatusEntry) DeepCopy() *ReleaseStatusEntry {
	if in == nil {
		return nil
	}
	out := new(ReleaseStatusEntry)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *VirtualEnv) DeepCopyInto(out *VirtualEnv) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Data.DeepCopyInto(&out.Data)
	in.Details.DeepCopyInto(&out.Details)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new VirtualEnv.
func (in *VirtualEnv) DeepCopy() *VirtualEnv {
	if in == nil {
		return nil
	}
	out := new(VirtualEnv)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *VirtualEnv) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *VirtualEnvList) DeepCopyInto(out *VirtualEnvList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]VirtualEnv, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new VirtualEnvList.
func (in *VirtualEnvList) DeepCopy() *VirtualEnvList {
	if in == nil {
		return nil
	}
	out := new(VirtualEnvList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *VirtualEnvList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *VirtualEnvSnapshot) DeepCopyInto(out *VirtualEnvSnapshot) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Data.DeepCopyInto(&out.Data)
	in.Details.DeepCopyInto(&out.Details)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new VirtualEnvSnapshot.
func (in *VirtualEnvSnapshot) DeepCopy() *VirtualEnvSnapshot {
	if in == nil {
		return nil
	}
	out := new(VirtualEnvSnapshot)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *VirtualEnvSnapshot) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *VirtualEnvSnapshotList) DeepCopyInto(out *VirtualEnvSnapshotList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]VirtualEnvSnapshot, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new VirtualEnvSnapshotList.
func (in *VirtualEnvSnapshotList) DeepCopy() *VirtualEnvSnapshotList {
	if in == nil {
		return nil
	}
	out := new(VirtualEnvSnapshotList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *VirtualEnvSnapshotList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}
