// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.35.2
// 	protoc        v3.12.4
// source: factorize.proto

package proto

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

type PrimeFactorRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Integer int32 `protobuf:"varint,1,opt,name=integer,proto3" json:"integer,omitempty"`
}

func (x *PrimeFactorRequest) Reset() {
	*x = PrimeFactorRequest{}
	mi := &file_factorize_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *PrimeFactorRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PrimeFactorRequest) ProtoMessage() {}

func (x *PrimeFactorRequest) ProtoReflect() protoreflect.Message {
	mi := &file_factorize_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PrimeFactorRequest.ProtoReflect.Descriptor instead.
func (*PrimeFactorRequest) Descriptor() ([]byte, []int) {
	return file_factorize_proto_rawDescGZIP(), []int{0}
}

func (x *PrimeFactorRequest) GetInteger() int32 {
	if x != nil {
		return x.Integer
	}
	return 0
}

type PrimeFactorResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	PrimeFactor int32 `protobuf:"varint,1,opt,name=prime_factor,json=primeFactor,proto3" json:"prime_factor,omitempty"`
}

func (x *PrimeFactorResponse) Reset() {
	*x = PrimeFactorResponse{}
	mi := &file_factorize_proto_msgTypes[1]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *PrimeFactorResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PrimeFactorResponse) ProtoMessage() {}

func (x *PrimeFactorResponse) ProtoReflect() protoreflect.Message {
	mi := &file_factorize_proto_msgTypes[1]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PrimeFactorResponse.ProtoReflect.Descriptor instead.
func (*PrimeFactorResponse) Descriptor() ([]byte, []int) {
	return file_factorize_proto_rawDescGZIP(), []int{1}
}

func (x *PrimeFactorResponse) GetPrimeFactor() int32 {
	if x != nil {
		return x.PrimeFactor
	}
	return 0
}

var File_factorize_proto protoreflect.FileDescriptor

var file_factorize_proto_rawDesc = []byte{
	0x0a, 0x0f, 0x66, 0x61, 0x63, 0x74, 0x6f, 0x72, 0x69, 0x7a, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x12, 0x0a, 0x63, 0x61, 0x6c, 0x63, 0x75, 0x6c, 0x61, 0x74, 0x6f, 0x72, 0x22, 0x2e, 0x0a,
	0x12, 0x50, 0x72, 0x69, 0x6d, 0x65, 0x46, 0x61, 0x63, 0x74, 0x6f, 0x72, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x12, 0x18, 0x0a, 0x07, 0x69, 0x6e, 0x74, 0x65, 0x67, 0x65, 0x72, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x05, 0x52, 0x07, 0x69, 0x6e, 0x74, 0x65, 0x67, 0x65, 0x72, 0x22, 0x38, 0x0a,
	0x13, 0x50, 0x72, 0x69, 0x6d, 0x65, 0x46, 0x61, 0x63, 0x74, 0x6f, 0x72, 0x52, 0x65, 0x73, 0x70,
	0x6f, 0x6e, 0x73, 0x65, 0x12, 0x21, 0x0a, 0x0c, 0x70, 0x72, 0x69, 0x6d, 0x65, 0x5f, 0x66, 0x61,
	0x63, 0x74, 0x6f, 0x72, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x0b, 0x70, 0x72, 0x69, 0x6d,
	0x65, 0x46, 0x61, 0x63, 0x74, 0x6f, 0x72, 0x42, 0x3a, 0x5a, 0x38, 0x67, 0x69, 0x74, 0x68, 0x75,
	0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x63, 0x61, 0x73, 0x70, 0x65, 0x72, 0x65, 0x69, 0x6a, 0x6b,
	0x65, 0x6e, 0x73, 0x2f, 0x67, 0x72, 0x70, 0x63, 0x2d, 0x67, 0x6f, 0x2d, 0x63, 0x6f, 0x75, 0x72,
	0x73, 0x65, 0x2f, 0x63, 0x61, 0x6c, 0x63, 0x75, 0x6c, 0x61, 0x74, 0x6f, 0x72, 0x2f, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_factorize_proto_rawDescOnce sync.Once
	file_factorize_proto_rawDescData = file_factorize_proto_rawDesc
)

func file_factorize_proto_rawDescGZIP() []byte {
	file_factorize_proto_rawDescOnce.Do(func() {
		file_factorize_proto_rawDescData = protoimpl.X.CompressGZIP(file_factorize_proto_rawDescData)
	})
	return file_factorize_proto_rawDescData
}

var file_factorize_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_factorize_proto_goTypes = []any{
	(*PrimeFactorRequest)(nil),  // 0: calculator.PrimeFactorRequest
	(*PrimeFactorResponse)(nil), // 1: calculator.PrimeFactorResponse
}
var file_factorize_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_factorize_proto_init() }
func file_factorize_proto_init() {
	if File_factorize_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_factorize_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_factorize_proto_goTypes,
		DependencyIndexes: file_factorize_proto_depIdxs,
		MessageInfos:      file_factorize_proto_msgTypes,
	}.Build()
	File_factorize_proto = out.File
	file_factorize_proto_rawDesc = nil
	file_factorize_proto_goTypes = nil
	file_factorize_proto_depIdxs = nil
}
