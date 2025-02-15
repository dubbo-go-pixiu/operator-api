//
// Licensed to the Apache Software Foundation (ASF) under one or more
// contributor license agreements.  See the NOTICE file distributed with
// this work for additional information regarding copyright ownership.
// The ASF licenses this file to You under the Apache License, Version 2.0
// (the "License"); you may not use this file except in compliance with
// the License.  You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.27.1
// 	protoc        (unknown)
// source: extensions/v1alpha1/service_name_mapping.proto

// $schema: istio.extensions.v1alpha1.ServiceNameMapping
// $title: Service Name Mapping
// $description: it's the spec of Service Name Mapping, which is used to map the interface name
// to the application name in application discovery.

// `Service Name Mapping` which is used to map the interface name
// to the application name in application discovery.

package v1alpha1

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

//
// <!-- crd generation tags
// +cue-gen:ServiceNameMapping:groupName:extensions.istio.io
// +cue-gen:ServiceNameMapping:version:v1alpha1
// +cue-gen:ServiceNameMapping:storageVersion
// +cue-gen:ServiceNameMapping:annotations:helm.sh/resource-policy=keep
// +cue-gen:ServiceNameMapping:labels:app=istio-pilot,chart=istio,heritage=Tiller,release=istio
// +cue-gen:ServiceNameMapping:subresource:status
// +cue-gen:ServiceNameMapping:scope:Namespaced
// +cue-gen:ServiceNameMapping:resource:categories=istio-io,extensions-istio-io,shortNames=snp,plural=servicenamemappings
// +cue-gen:ServiceNameMapping:printerColumn:name=Age,type=date,JSONPath=.metadata.creationTimestamp,description="CreationTimestamp is a timestamp
// representing the server time when this object was created. It is not guaranteed to be set in happens-before order across separate operations.
// Clients may not set this value. It is represented in RFC3339 form and is in UTC.
// Populated by the system. Read-only. Null for lists. More info: https://git.k8s.io/community/contributors/devel/api-conventions.md#metadata"
// +cue-gen:ServiceNameMapping:preserveUnknownFields:false
// -->
//
// <!-- go code generation tags
// +kubetype-gen
// +kubetype-gen:groupVersion=extensions.istio.io/v1alpha1
// +genclient
// +k8s:deepcopy-gen=true
// -->
type ServiceNameMapping struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// InterfaceName.
	InterfaceName string `protobuf:"bytes,1,opt,name=interfaceName,proto3" json:"interfaceName,omitempty"`
	// ApplicationNames。
	ApplicationNames []string `protobuf:"bytes,2,rep,name=applicationNames,proto3" json:"applicationNames,omitempty"`
}

func (x *ServiceNameMapping) Reset() {
	*x = ServiceNameMapping{}
	if protoimpl.UnsafeEnabled {
		mi := &file_extensions_v1alpha1_service_name_mapping_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ServiceNameMapping) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ServiceNameMapping) ProtoMessage() {}

func (x *ServiceNameMapping) ProtoReflect() protoreflect.Message {
	mi := &file_extensions_v1alpha1_service_name_mapping_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ServiceNameMapping.ProtoReflect.Descriptor instead.
func (*ServiceNameMapping) Descriptor() ([]byte, []int) {
	return file_extensions_v1alpha1_service_name_mapping_proto_rawDescGZIP(), []int{0}
}

func (x *ServiceNameMapping) GetInterfaceName() string {
	if x != nil {
		return x.InterfaceName
	}
	return ""
}

func (x *ServiceNameMapping) GetApplicationNames() []string {
	if x != nil {
		return x.ApplicationNames
	}
	return nil
}

var File_extensions_v1alpha1_service_name_mapping_proto protoreflect.FileDescriptor

var file_extensions_v1alpha1_service_name_mapping_proto_rawDesc = []byte{
	0x0a, 0x2e, 0x65, 0x78, 0x74, 0x65, 0x6e, 0x73, 0x69, 0x6f, 0x6e, 0x73, 0x2f, 0x76, 0x31, 0x61,
	0x6c, 0x70, 0x68, 0x61, 0x31, 0x2f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x5f, 0x6e, 0x61,
	0x6d, 0x65, 0x5f, 0x6d, 0x61, 0x70, 0x70, 0x69, 0x6e, 0x67, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x12, 0x19, 0x69, 0x73, 0x74, 0x69, 0x6f, 0x2e, 0x65, 0x78, 0x74, 0x65, 0x6e, 0x73, 0x69, 0x6f,
	0x6e, 0x73, 0x2e, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x31, 0x22, 0x66, 0x0a, 0x12, 0x53,
	0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x4e, 0x61, 0x6d, 0x65, 0x4d, 0x61, 0x70, 0x70, 0x69, 0x6e,
	0x67, 0x12, 0x24, 0x0a, 0x0d, 0x69, 0x6e, 0x74, 0x65, 0x72, 0x66, 0x61, 0x63, 0x65, 0x4e, 0x61,
	0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0d, 0x69, 0x6e, 0x74, 0x65, 0x72, 0x66,
	0x61, 0x63, 0x65, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x2a, 0x0a, 0x10, 0x61, 0x70, 0x70, 0x6c, 0x69,
	0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x4e, 0x61, 0x6d, 0x65, 0x73, 0x18, 0x02, 0x20, 0x03, 0x28,
	0x09, 0x52, 0x10, 0x61, 0x70, 0x70, 0x6c, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x4e, 0x61,
	0x6d, 0x65, 0x73, 0x42, 0x22, 0x5a, 0x20, 0x69, 0x73, 0x74, 0x69, 0x6f, 0x2e, 0x69, 0x6f, 0x2f,
	0x61, 0x70, 0x69, 0x2f, 0x65, 0x78, 0x74, 0x65, 0x6e, 0x73, 0x69, 0x6f, 0x6e, 0x73, 0x2f, 0x76,
	0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x31, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_extensions_v1alpha1_service_name_mapping_proto_rawDescOnce sync.Once
	file_extensions_v1alpha1_service_name_mapping_proto_rawDescData = file_extensions_v1alpha1_service_name_mapping_proto_rawDesc
)

func file_extensions_v1alpha1_service_name_mapping_proto_rawDescGZIP() []byte {
	file_extensions_v1alpha1_service_name_mapping_proto_rawDescOnce.Do(func() {
		file_extensions_v1alpha1_service_name_mapping_proto_rawDescData = protoimpl.X.CompressGZIP(file_extensions_v1alpha1_service_name_mapping_proto_rawDescData)
	})
	return file_extensions_v1alpha1_service_name_mapping_proto_rawDescData
}

var file_extensions_v1alpha1_service_name_mapping_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_extensions_v1alpha1_service_name_mapping_proto_goTypes = []interface{}{
	(*ServiceNameMapping)(nil), // 0: istio.extensions.v1alpha1.ServiceNameMapping
}
var file_extensions_v1alpha1_service_name_mapping_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_extensions_v1alpha1_service_name_mapping_proto_init() }
func file_extensions_v1alpha1_service_name_mapping_proto_init() {
	if File_extensions_v1alpha1_service_name_mapping_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_extensions_v1alpha1_service_name_mapping_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ServiceNameMapping); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_extensions_v1alpha1_service_name_mapping_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_extensions_v1alpha1_service_name_mapping_proto_goTypes,
		DependencyIndexes: file_extensions_v1alpha1_service_name_mapping_proto_depIdxs,
		MessageInfos:      file_extensions_v1alpha1_service_name_mapping_proto_msgTypes,
	}.Build()
	File_extensions_v1alpha1_service_name_mapping_proto = out.File
	file_extensions_v1alpha1_service_name_mapping_proto_rawDesc = nil
	file_extensions_v1alpha1_service_name_mapping_proto_goTypes = nil
	file_extensions_v1alpha1_service_name_mapping_proto_depIdxs = nil
}
