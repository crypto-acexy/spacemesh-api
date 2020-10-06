// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.25.0
// 	protoc        v3.12.3
// source: spacemesh/v1/gateway_types.proto

package v1

import (
	proto "github.com/golang/protobuf/proto"
	status "google.golang.org/genproto/googleapis/rpc/status"
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

// This is a compile-time assertion that a sufficiently up-to-date version
// of the legacy proto package is being used.
const _ = proto.ProtoPackageIsVersion4

type BroadcastPoetRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Data []byte `protobuf:"bytes,1,opt,name=data,proto3" json:"data,omitempty"` // signed binary transaction
}

func (x *BroadcastPoetRequest) Reset() {
	*x = BroadcastPoetRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_spacemesh_v1_gateway_types_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *BroadcastPoetRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*BroadcastPoetRequest) ProtoMessage() {}

func (x *BroadcastPoetRequest) ProtoReflect() protoreflect.Message {
	mi := &file_spacemesh_v1_gateway_types_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use BroadcastPoetRequest.ProtoReflect.Descriptor instead.
func (*BroadcastPoetRequest) Descriptor() ([]byte, []int) {
	return file_spacemesh_v1_gateway_types_proto_rawDescGZIP(), []int{0}
}

func (x *BroadcastPoetRequest) GetData() []byte {
	if x != nil {
		return x.Data
	}
	return nil
}

type BroadcastPoetResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Status *status.Status `protobuf:"bytes,1,opt,name=status,proto3" json:"status,omitempty"`
}

func (x *BroadcastPoetResponse) Reset() {
	*x = BroadcastPoetResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_spacemesh_v1_gateway_types_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *BroadcastPoetResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*BroadcastPoetResponse) ProtoMessage() {}

func (x *BroadcastPoetResponse) ProtoReflect() protoreflect.Message {
	mi := &file_spacemesh_v1_gateway_types_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use BroadcastPoetResponse.ProtoReflect.Descriptor instead.
func (*BroadcastPoetResponse) Descriptor() ([]byte, []int) {
	return file_spacemesh_v1_gateway_types_proto_rawDescGZIP(), []int{1}
}

func (x *BroadcastPoetResponse) GetStatus() *status.Status {
	if x != nil {
		return x.Status
	}
	return nil
}

var File_spacemesh_v1_gateway_types_proto protoreflect.FileDescriptor

var file_spacemesh_v1_gateway_types_proto_rawDesc = []byte{
	0x0a, 0x20, 0x73, 0x70, 0x61, 0x63, 0x65, 0x6d, 0x65, 0x73, 0x68, 0x2f, 0x76, 0x31, 0x2f, 0x67,
	0x61, 0x74, 0x65, 0x77, 0x61, 0x79, 0x5f, 0x74, 0x79, 0x70, 0x65, 0x73, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x12, 0x0c, 0x73, 0x70, 0x61, 0x63, 0x65, 0x6d, 0x65, 0x73, 0x68, 0x2e, 0x76, 0x31,
	0x1a, 0x17, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x72, 0x70, 0x63, 0x2f, 0x73, 0x74, 0x61,
	0x74, 0x75, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x2a, 0x0a, 0x14, 0x42, 0x72, 0x6f,
	0x61, 0x64, 0x63, 0x61, 0x73, 0x74, 0x50, 0x6f, 0x65, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x12, 0x12, 0x0a, 0x04, 0x64, 0x61, 0x74, 0x61, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0c, 0x52,
	0x04, 0x64, 0x61, 0x74, 0x61, 0x22, 0x43, 0x0a, 0x15, 0x42, 0x72, 0x6f, 0x61, 0x64, 0x63, 0x61,
	0x73, 0x74, 0x50, 0x6f, 0x65, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x2a,
	0x0a, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x12,
	0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x72, 0x70, 0x63, 0x2e, 0x53, 0x74, 0x61, 0x74,
	0x75, 0x73, 0x52, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x42, 0x34, 0x5a, 0x32, 0x67, 0x69,
	0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x73, 0x70, 0x61, 0x63, 0x65, 0x6d, 0x65,
	0x73, 0x68, 0x6f, 0x73, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x72, 0x65, 0x6c, 0x65, 0x61, 0x73, 0x65,
	0x2f, 0x67, 0x6f, 0x2f, 0x73, 0x70, 0x61, 0x63, 0x65, 0x6d, 0x65, 0x73, 0x68, 0x2f, 0x76, 0x31,
	0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_spacemesh_v1_gateway_types_proto_rawDescOnce sync.Once
	file_spacemesh_v1_gateway_types_proto_rawDescData = file_spacemesh_v1_gateway_types_proto_rawDesc
)

func file_spacemesh_v1_gateway_types_proto_rawDescGZIP() []byte {
	file_spacemesh_v1_gateway_types_proto_rawDescOnce.Do(func() {
		file_spacemesh_v1_gateway_types_proto_rawDescData = protoimpl.X.CompressGZIP(file_spacemesh_v1_gateway_types_proto_rawDescData)
	})
	return file_spacemesh_v1_gateway_types_proto_rawDescData
}

var file_spacemesh_v1_gateway_types_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_spacemesh_v1_gateway_types_proto_goTypes = []interface{}{
	(*BroadcastPoetRequest)(nil),  // 0: spacemesh.v1.BroadcastPoetRequest
	(*BroadcastPoetResponse)(nil), // 1: spacemesh.v1.BroadcastPoetResponse
	(*status.Status)(nil),         // 2: google.rpc.Status
}
var file_spacemesh_v1_gateway_types_proto_depIdxs = []int32{
	2, // 0: spacemesh.v1.BroadcastPoetResponse.status:type_name -> google.rpc.Status
	1, // [1:1] is the sub-list for method output_type
	1, // [1:1] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_spacemesh_v1_gateway_types_proto_init() }
func file_spacemesh_v1_gateway_types_proto_init() {
	if File_spacemesh_v1_gateway_types_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_spacemesh_v1_gateway_types_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*BroadcastPoetRequest); i {
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
		file_spacemesh_v1_gateway_types_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*BroadcastPoetResponse); i {
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
			RawDescriptor: file_spacemesh_v1_gateway_types_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_spacemesh_v1_gateway_types_proto_goTypes,
		DependencyIndexes: file_spacemesh_v1_gateway_types_proto_depIdxs,
		MessageInfos:      file_spacemesh_v1_gateway_types_proto_msgTypes,
	}.Build()
	File_spacemesh_v1_gateway_types_proto = out.File
	file_spacemesh_v1_gateway_types_proto_rawDesc = nil
	file_spacemesh_v1_gateway_types_proto_goTypes = nil
	file_spacemesh_v1_gateway_types_proto_depIdxs = nil
}
