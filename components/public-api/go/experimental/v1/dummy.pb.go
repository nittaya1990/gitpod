// Copyright (c) 2023 Gitpod GmbH. All rights reserved.
// Licensed under the GNU Affero General Public License (AGPL).
// See License.AGPL.txt in the project root for license information.

// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        (unknown)
// source: gitpod/experimental/v1/dummy.proto

package v1

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

type SayHelloRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *SayHelloRequest) Reset() {
	*x = SayHelloRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_gitpod_experimental_v1_dummy_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SayHelloRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SayHelloRequest) ProtoMessage() {}

func (x *SayHelloRequest) ProtoReflect() protoreflect.Message {
	mi := &file_gitpod_experimental_v1_dummy_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SayHelloRequest.ProtoReflect.Descriptor instead.
func (*SayHelloRequest) Descriptor() ([]byte, []int) {
	return file_gitpod_experimental_v1_dummy_proto_rawDescGZIP(), []int{0}
}

type SayHelloResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Reply string `protobuf:"bytes,1,opt,name=reply,proto3" json:"reply,omitempty"`
}

func (x *SayHelloResponse) Reset() {
	*x = SayHelloResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_gitpod_experimental_v1_dummy_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SayHelloResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SayHelloResponse) ProtoMessage() {}

func (x *SayHelloResponse) ProtoReflect() protoreflect.Message {
	mi := &file_gitpod_experimental_v1_dummy_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SayHelloResponse.ProtoReflect.Descriptor instead.
func (*SayHelloResponse) Descriptor() ([]byte, []int) {
	return file_gitpod_experimental_v1_dummy_proto_rawDescGZIP(), []int{1}
}

func (x *SayHelloResponse) GetReply() string {
	if x != nil {
		return x.Reply
	}
	return ""
}

type LotsOfRepliesRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	PreviousCount int32 `protobuf:"varint,1,opt,name=previous_count,json=previousCount,proto3" json:"previous_count,omitempty"`
}

func (x *LotsOfRepliesRequest) Reset() {
	*x = LotsOfRepliesRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_gitpod_experimental_v1_dummy_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *LotsOfRepliesRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*LotsOfRepliesRequest) ProtoMessage() {}

func (x *LotsOfRepliesRequest) ProtoReflect() protoreflect.Message {
	mi := &file_gitpod_experimental_v1_dummy_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use LotsOfRepliesRequest.ProtoReflect.Descriptor instead.
func (*LotsOfRepliesRequest) Descriptor() ([]byte, []int) {
	return file_gitpod_experimental_v1_dummy_proto_rawDescGZIP(), []int{2}
}

func (x *LotsOfRepliesRequest) GetPreviousCount() int32 {
	if x != nil {
		return x.PreviousCount
	}
	return 0
}

type LotsOfRepliesResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Reply string `protobuf:"bytes,1,opt,name=reply,proto3" json:"reply,omitempty"`
	Count int32  `protobuf:"varint,2,opt,name=count,proto3" json:"count,omitempty"`
}

func (x *LotsOfRepliesResponse) Reset() {
	*x = LotsOfRepliesResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_gitpod_experimental_v1_dummy_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *LotsOfRepliesResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*LotsOfRepliesResponse) ProtoMessage() {}

func (x *LotsOfRepliesResponse) ProtoReflect() protoreflect.Message {
	mi := &file_gitpod_experimental_v1_dummy_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use LotsOfRepliesResponse.ProtoReflect.Descriptor instead.
func (*LotsOfRepliesResponse) Descriptor() ([]byte, []int) {
	return file_gitpod_experimental_v1_dummy_proto_rawDescGZIP(), []int{3}
}

func (x *LotsOfRepliesResponse) GetReply() string {
	if x != nil {
		return x.Reply
	}
	return ""
}

func (x *LotsOfRepliesResponse) GetCount() int32 {
	if x != nil {
		return x.Count
	}
	return 0
}

var File_gitpod_experimental_v1_dummy_proto protoreflect.FileDescriptor

var file_gitpod_experimental_v1_dummy_proto_rawDesc = []byte{
	0x0a, 0x22, 0x67, 0x69, 0x74, 0x70, 0x6f, 0x64, 0x2f, 0x65, 0x78, 0x70, 0x65, 0x72, 0x69, 0x6d,
	0x65, 0x6e, 0x74, 0x61, 0x6c, 0x2f, 0x76, 0x31, 0x2f, 0x64, 0x75, 0x6d, 0x6d, 0x79, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x12, 0x16, 0x67, 0x69, 0x74, 0x70, 0x6f, 0x64, 0x2e, 0x65, 0x78, 0x70,
	0x65, 0x72, 0x69, 0x6d, 0x65, 0x6e, 0x74, 0x61, 0x6c, 0x2e, 0x76, 0x31, 0x22, 0x11, 0x0a, 0x0f,
	0x53, 0x61, 0x79, 0x48, 0x65, 0x6c, 0x6c, 0x6f, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x22,
	0x28, 0x0a, 0x10, 0x53, 0x61, 0x79, 0x48, 0x65, 0x6c, 0x6c, 0x6f, 0x52, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x73, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x72, 0x65, 0x70, 0x6c, 0x79, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x05, 0x72, 0x65, 0x70, 0x6c, 0x79, 0x22, 0x3d, 0x0a, 0x14, 0x4c, 0x6f, 0x74,
	0x73, 0x4f, 0x66, 0x52, 0x65, 0x70, 0x6c, 0x69, 0x65, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x12, 0x25, 0x0a, 0x0e, 0x70, 0x72, 0x65, 0x76, 0x69, 0x6f, 0x75, 0x73, 0x5f, 0x63, 0x6f,
	0x75, 0x6e, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x0d, 0x70, 0x72, 0x65, 0x76, 0x69,
	0x6f, 0x75, 0x73, 0x43, 0x6f, 0x75, 0x6e, 0x74, 0x22, 0x43, 0x0a, 0x15, 0x4c, 0x6f, 0x74, 0x73,
	0x4f, 0x66, 0x52, 0x65, 0x70, 0x6c, 0x69, 0x65, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73,
	0x65, 0x12, 0x14, 0x0a, 0x05, 0x72, 0x65, 0x70, 0x6c, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x05, 0x72, 0x65, 0x70, 0x6c, 0x79, 0x12, 0x14, 0x0a, 0x05, 0x63, 0x6f, 0x75, 0x6e, 0x74,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x05, 0x52, 0x05, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x32, 0xdd, 0x01,
	0x0a, 0x0c, 0x48, 0x65, 0x6c, 0x6c, 0x6f, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x5d,
	0x0a, 0x08, 0x53, 0x61, 0x79, 0x48, 0x65, 0x6c, 0x6c, 0x6f, 0x12, 0x27, 0x2e, 0x67, 0x69, 0x74,
	0x70, 0x6f, 0x64, 0x2e, 0x65, 0x78, 0x70, 0x65, 0x72, 0x69, 0x6d, 0x65, 0x6e, 0x74, 0x61, 0x6c,
	0x2e, 0x76, 0x31, 0x2e, 0x53, 0x61, 0x79, 0x48, 0x65, 0x6c, 0x6c, 0x6f, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x1a, 0x28, 0x2e, 0x67, 0x69, 0x74, 0x70, 0x6f, 0x64, 0x2e, 0x65, 0x78, 0x70,
	0x65, 0x72, 0x69, 0x6d, 0x65, 0x6e, 0x74, 0x61, 0x6c, 0x2e, 0x76, 0x31, 0x2e, 0x53, 0x61, 0x79,
	0x48, 0x65, 0x6c, 0x6c, 0x6f, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x6e, 0x0a,
	0x0d, 0x4c, 0x6f, 0x74, 0x73, 0x4f, 0x66, 0x52, 0x65, 0x70, 0x6c, 0x69, 0x65, 0x73, 0x12, 0x2c,
	0x2e, 0x67, 0x69, 0x74, 0x70, 0x6f, 0x64, 0x2e, 0x65, 0x78, 0x70, 0x65, 0x72, 0x69, 0x6d, 0x65,
	0x6e, 0x74, 0x61, 0x6c, 0x2e, 0x76, 0x31, 0x2e, 0x4c, 0x6f, 0x74, 0x73, 0x4f, 0x66, 0x52, 0x65,
	0x70, 0x6c, 0x69, 0x65, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x2d, 0x2e, 0x67,
	0x69, 0x74, 0x70, 0x6f, 0x64, 0x2e, 0x65, 0x78, 0x70, 0x65, 0x72, 0x69, 0x6d, 0x65, 0x6e, 0x74,
	0x61, 0x6c, 0x2e, 0x76, 0x31, 0x2e, 0x4c, 0x6f, 0x74, 0x73, 0x4f, 0x66, 0x52, 0x65, 0x70, 0x6c,
	0x69, 0x65, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x30, 0x01, 0x42, 0x46, 0x5a,
	0x44, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x67, 0x69, 0x74, 0x70,
	0x6f, 0x64, 0x2d, 0x69, 0x6f, 0x2f, 0x67, 0x69, 0x74, 0x70, 0x6f, 0x64, 0x2f, 0x63, 0x6f, 0x6d,
	0x70, 0x6f, 0x6e, 0x65, 0x6e, 0x74, 0x73, 0x2f, 0x70, 0x75, 0x62, 0x6c, 0x69, 0x63, 0x2d, 0x61,
	0x70, 0x69, 0x2f, 0x67, 0x6f, 0x2f, 0x65, 0x78, 0x70, 0x65, 0x72, 0x69, 0x6d, 0x65, 0x6e, 0x74,
	0x61, 0x6c, 0x2f, 0x76, 0x31, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_gitpod_experimental_v1_dummy_proto_rawDescOnce sync.Once
	file_gitpod_experimental_v1_dummy_proto_rawDescData = file_gitpod_experimental_v1_dummy_proto_rawDesc
)

func file_gitpod_experimental_v1_dummy_proto_rawDescGZIP() []byte {
	file_gitpod_experimental_v1_dummy_proto_rawDescOnce.Do(func() {
		file_gitpod_experimental_v1_dummy_proto_rawDescData = protoimpl.X.CompressGZIP(file_gitpod_experimental_v1_dummy_proto_rawDescData)
	})
	return file_gitpod_experimental_v1_dummy_proto_rawDescData
}

var file_gitpod_experimental_v1_dummy_proto_msgTypes = make([]protoimpl.MessageInfo, 4)
var file_gitpod_experimental_v1_dummy_proto_goTypes = []interface{}{
	(*SayHelloRequest)(nil),       // 0: gitpod.experimental.v1.SayHelloRequest
	(*SayHelloResponse)(nil),      // 1: gitpod.experimental.v1.SayHelloResponse
	(*LotsOfRepliesRequest)(nil),  // 2: gitpod.experimental.v1.LotsOfRepliesRequest
	(*LotsOfRepliesResponse)(nil), // 3: gitpod.experimental.v1.LotsOfRepliesResponse
}
var file_gitpod_experimental_v1_dummy_proto_depIdxs = []int32{
	0, // 0: gitpod.experimental.v1.HelloService.SayHello:input_type -> gitpod.experimental.v1.SayHelloRequest
	2, // 1: gitpod.experimental.v1.HelloService.LotsOfReplies:input_type -> gitpod.experimental.v1.LotsOfRepliesRequest
	1, // 2: gitpod.experimental.v1.HelloService.SayHello:output_type -> gitpod.experimental.v1.SayHelloResponse
	3, // 3: gitpod.experimental.v1.HelloService.LotsOfReplies:output_type -> gitpod.experimental.v1.LotsOfRepliesResponse
	2, // [2:4] is the sub-list for method output_type
	0, // [0:2] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_gitpod_experimental_v1_dummy_proto_init() }
func file_gitpod_experimental_v1_dummy_proto_init() {
	if File_gitpod_experimental_v1_dummy_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_gitpod_experimental_v1_dummy_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SayHelloRequest); i {
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
		file_gitpod_experimental_v1_dummy_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SayHelloResponse); i {
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
		file_gitpod_experimental_v1_dummy_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*LotsOfRepliesRequest); i {
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
		file_gitpod_experimental_v1_dummy_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*LotsOfRepliesResponse); i {
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
			RawDescriptor: file_gitpod_experimental_v1_dummy_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   4,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_gitpod_experimental_v1_dummy_proto_goTypes,
		DependencyIndexes: file_gitpod_experimental_v1_dummy_proto_depIdxs,
		MessageInfos:      file_gitpod_experimental_v1_dummy_proto_msgTypes,
	}.Build()
	File_gitpod_experimental_v1_dummy_proto = out.File
	file_gitpod_experimental_v1_dummy_proto_rawDesc = nil
	file_gitpod_experimental_v1_dummy_proto_goTypes = nil
	file_gitpod_experimental_v1_dummy_proto_depIdxs = nil
}
