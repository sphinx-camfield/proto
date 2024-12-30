// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.35.2
// 	protoc        v3.21.12
// source: svc/room-manager.proto

package svc

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

type RoomTokenRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Room string `protobuf:"bytes,1,opt,name=room,proto3" json:"room,omitempty"`
}

func (x *RoomTokenRequest) Reset() {
	*x = RoomTokenRequest{}
	mi := &file_svc_room_manager_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *RoomTokenRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RoomTokenRequest) ProtoMessage() {}

func (x *RoomTokenRequest) ProtoReflect() protoreflect.Message {
	mi := &file_svc_room_manager_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RoomTokenRequest.ProtoReflect.Descriptor instead.
func (*RoomTokenRequest) Descriptor() ([]byte, []int) {
	return file_svc_room_manager_proto_rawDescGZIP(), []int{0}
}

func (x *RoomTokenRequest) GetRoom() string {
	if x != nil {
		return x.Room
	}
	return ""
}

type RoomTokenReply struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Token    string `protobuf:"bytes,1,opt,name=token,proto3" json:"token,omitempty"`
	Endpoint string `protobuf:"bytes,2,opt,name=endpoint,proto3" json:"endpoint,omitempty"`
	Expiry   int64  `protobuf:"varint,3,opt,name=expiry,proto3" json:"expiry,omitempty"`
}

func (x *RoomTokenReply) Reset() {
	*x = RoomTokenReply{}
	mi := &file_svc_room_manager_proto_msgTypes[1]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *RoomTokenReply) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RoomTokenReply) ProtoMessage() {}

func (x *RoomTokenReply) ProtoReflect() protoreflect.Message {
	mi := &file_svc_room_manager_proto_msgTypes[1]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RoomTokenReply.ProtoReflect.Descriptor instead.
func (*RoomTokenReply) Descriptor() ([]byte, []int) {
	return file_svc_room_manager_proto_rawDescGZIP(), []int{1}
}

func (x *RoomTokenReply) GetToken() string {
	if x != nil {
		return x.Token
	}
	return ""
}

func (x *RoomTokenReply) GetEndpoint() string {
	if x != nil {
		return x.Endpoint
	}
	return ""
}

func (x *RoomTokenReply) GetExpiry() int64 {
	if x != nil {
		return x.Expiry
	}
	return 0
}

var File_svc_room_manager_proto protoreflect.FileDescriptor

var file_svc_room_manager_proto_rawDesc = []byte{
	0x0a, 0x16, 0x73, 0x76, 0x63, 0x2f, 0x72, 0x6f, 0x6f, 0x6d, 0x2d, 0x6d, 0x61, 0x6e, 0x61, 0x67,
	0x65, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x13, 0x73, 0x70, 0x68, 0x69, 0x6e, 0x78,
	0x2e, 0x63, 0x61, 0x6d, 0x66, 0x69, 0x65, 0x6c, 0x64, 0x2e, 0x73, 0x76, 0x63, 0x22, 0x26, 0x0a,
	0x10, 0x52, 0x6f, 0x6f, 0x6d, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x12, 0x12, 0x0a, 0x04, 0x72, 0x6f, 0x6f, 0x6d, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x04, 0x72, 0x6f, 0x6f, 0x6d, 0x22, 0x5a, 0x0a, 0x0e, 0x52, 0x6f, 0x6f, 0x6d, 0x54, 0x6f, 0x6b,
	0x65, 0x6e, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x12, 0x14, 0x0a, 0x05, 0x74, 0x6f, 0x6b, 0x65, 0x6e,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x12, 0x1a, 0x0a,
	0x08, 0x65, 0x6e, 0x64, 0x70, 0x6f, 0x69, 0x6e, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x08, 0x65, 0x6e, 0x64, 0x70, 0x6f, 0x69, 0x6e, 0x74, 0x12, 0x16, 0x0a, 0x06, 0x65, 0x78, 0x70,
	0x69, 0x72, 0x79, 0x18, 0x03, 0x20, 0x01, 0x28, 0x03, 0x52, 0x06, 0x65, 0x78, 0x70, 0x69, 0x72,
	0x79, 0x32, 0x64, 0x0a, 0x0b, 0x52, 0x6f, 0x6f, 0x6d, 0x4d, 0x61, 0x6e, 0x61, 0x67, 0x65, 0x72,
	0x12, 0x55, 0x0a, 0x07, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x25, 0x2e, 0x73, 0x70,
	0x68, 0x69, 0x6e, 0x78, 0x2e, 0x63, 0x61, 0x6d, 0x66, 0x69, 0x65, 0x6c, 0x64, 0x2e, 0x73, 0x76,
	0x63, 0x2e, 0x52, 0x6f, 0x6f, 0x6d, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x1a, 0x23, 0x2e, 0x73, 0x70, 0x68, 0x69, 0x6e, 0x78, 0x2e, 0x63, 0x61, 0x6d, 0x66,
	0x69, 0x65, 0x6c, 0x64, 0x2e, 0x73, 0x76, 0x63, 0x2e, 0x52, 0x6f, 0x6f, 0x6d, 0x54, 0x6f, 0x6b,
	0x65, 0x6e, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x42, 0x20, 0x5a, 0x1e, 0x67, 0x69, 0x74, 0x68, 0x75,
	0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x73, 0x70, 0x68, 0x69, 0x6e, 0x78, 0x2d, 0x63, 0x61, 0x6d,
	0x66, 0x69, 0x65, 0x6c, 0x64, 0x2f, 0x73, 0x76, 0x63, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x33,
}

var (
	file_svc_room_manager_proto_rawDescOnce sync.Once
	file_svc_room_manager_proto_rawDescData = file_svc_room_manager_proto_rawDesc
)

func file_svc_room_manager_proto_rawDescGZIP() []byte {
	file_svc_room_manager_proto_rawDescOnce.Do(func() {
		file_svc_room_manager_proto_rawDescData = protoimpl.X.CompressGZIP(file_svc_room_manager_proto_rawDescData)
	})
	return file_svc_room_manager_proto_rawDescData
}

var file_svc_room_manager_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_svc_room_manager_proto_goTypes = []any{
	(*RoomTokenRequest)(nil), // 0: sphinx.camfield.svc.RoomTokenRequest
	(*RoomTokenReply)(nil),   // 1: sphinx.camfield.svc.RoomTokenReply
}
var file_svc_room_manager_proto_depIdxs = []int32{
	0, // 0: sphinx.camfield.svc.RoomManager.Request:input_type -> sphinx.camfield.svc.RoomTokenRequest
	1, // 1: sphinx.camfield.svc.RoomManager.Request:output_type -> sphinx.camfield.svc.RoomTokenReply
	1, // [1:2] is the sub-list for method output_type
	0, // [0:1] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_svc_room_manager_proto_init() }
func file_svc_room_manager_proto_init() {
	if File_svc_room_manager_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_svc_room_manager_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_svc_room_manager_proto_goTypes,
		DependencyIndexes: file_svc_room_manager_proto_depIdxs,
		MessageInfos:      file_svc_room_manager_proto_msgTypes,
	}.Build()
	File_svc_room_manager_proto = out.File
	file_svc_room_manager_proto_rawDesc = nil
	file_svc_room_manager_proto_goTypes = nil
	file_svc_room_manager_proto_depIdxs = nil
}
