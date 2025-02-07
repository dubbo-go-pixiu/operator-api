// Code generated by protoc-gen-jsonshim. DO NOT EDIT.
package v1alpha1

import (
	bytes "bytes"
	jsonpb "github.com/golang/protobuf/jsonpb"
)

// MarshalJSON is a custom marshaler for ServiceMetadata
func (this *ServiceMetadata) MarshalJSON() ([]byte, error) {
	str, err := ServiceMetadataMarshaler.MarshalToString(this)
	return []byte(str), err
}

// UnmarshalJSON is a custom unmarshaler for ServiceMetadata
func (this *ServiceMetadata) UnmarshalJSON(b []byte) error {
	return ServiceMetadataUnmarshaler.Unmarshal(bytes.NewReader(b), this)
}

var (
	ServiceMetadataMarshaler   = &jsonpb.Marshaler{}
	ServiceMetadataUnmarshaler = &jsonpb.Unmarshaler{AllowUnknownFields: true}
)
