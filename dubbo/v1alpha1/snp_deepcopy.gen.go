// Code generated by protoc-gen-deepcopy. DO NOT EDIT.
package v1alpha1

import (
	proto "github.com/golang/protobuf/proto"
)

// DeepCopyInto supports using ServiceMappingRequest within kubernetes types, where deepcopy-gen is used.
func (in *ServiceMappingRequest) DeepCopyInto(out *ServiceMappingRequest) {
	p := proto.Clone(in).(*ServiceMappingRequest)
	*out = *p
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ServiceMappingRequest. Required by controller-gen.
func (in *ServiceMappingRequest) DeepCopy() *ServiceMappingRequest {
	if in == nil {
		return nil
	}
	out := new(ServiceMappingRequest)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInterface is an autogenerated deepcopy function, copying the receiver, creating a new ServiceMappingRequest. Required by controller-gen.
func (in *ServiceMappingRequest) DeepCopyInterface() interface{} {
	return in.DeepCopy()
}

// DeepCopyInto supports using ServiceMappingResponse within kubernetes types, where deepcopy-gen is used.
func (in *ServiceMappingResponse) DeepCopyInto(out *ServiceMappingResponse) {
	p := proto.Clone(in).(*ServiceMappingResponse)
	*out = *p
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ServiceMappingResponse. Required by controller-gen.
func (in *ServiceMappingResponse) DeepCopy() *ServiceMappingResponse {
	if in == nil {
		return nil
	}
	out := new(ServiceMappingResponse)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInterface is an autogenerated deepcopy function, copying the receiver, creating a new ServiceMappingResponse. Required by controller-gen.
func (in *ServiceMappingResponse) DeepCopyInterface() interface{} {
	return in.DeepCopy()
}

// DeepCopyInto supports using ServiceMappingXdsResponse within kubernetes types, where deepcopy-gen is used.
func (in *ServiceMappingXdsResponse) DeepCopyInto(out *ServiceMappingXdsResponse) {
	p := proto.Clone(in).(*ServiceMappingXdsResponse)
	*out = *p
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ServiceMappingXdsResponse. Required by controller-gen.
func (in *ServiceMappingXdsResponse) DeepCopy() *ServiceMappingXdsResponse {
	if in == nil {
		return nil
	}
	out := new(ServiceMappingXdsResponse)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInterface is an autogenerated deepcopy function, copying the receiver, creating a new ServiceMappingXdsResponse. Required by controller-gen.
func (in *ServiceMappingXdsResponse) DeepCopyInterface() interface{} {
	return in.DeepCopy()
}
