// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.35.1
// 	protoc        v5.28.3
// source: user/methods/exit.proto

package user

import (
	common "github.com/Yux77Yux/platform_backend/generated/common"
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

type ExitRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	RefreshToken *common.RefreshToken `protobuf:"bytes,1,opt,name=refresh_token,json=refreshToken,proto3" json:"refresh_token,omitempty"`
}

func (x *ExitRequest) Reset() {
	*x = ExitRequest{}
	mi := &file_user_methods_exit_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *ExitRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ExitRequest) ProtoMessage() {}

func (x *ExitRequest) ProtoReflect() protoreflect.Message {
	mi := &file_user_methods_exit_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ExitRequest.ProtoReflect.Descriptor instead.
func (*ExitRequest) Descriptor() ([]byte, []int) {
	return file_user_methods_exit_proto_rawDescGZIP(), []int{0}
}

func (x *ExitRequest) GetRefreshToken() *common.RefreshToken {
	if x != nil {
		return x.RefreshToken
	}
	return nil
}

type ExitResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Status bool    `protobuf:"varint,1,opt,name=status,proto3" json:"status,omitempty"`
	Error  *string `protobuf:"bytes,2,opt,name=error,proto3,oneof" json:"error,omitempty"`
}

func (x *ExitResponse) Reset() {
	*x = ExitResponse{}
	mi := &file_user_methods_exit_proto_msgTypes[1]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *ExitResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ExitResponse) ProtoMessage() {}

func (x *ExitResponse) ProtoReflect() protoreflect.Message {
	mi := &file_user_methods_exit_proto_msgTypes[1]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ExitResponse.ProtoReflect.Descriptor instead.
func (*ExitResponse) Descriptor() ([]byte, []int) {
	return file_user_methods_exit_proto_rawDescGZIP(), []int{1}
}

func (x *ExitResponse) GetStatus() bool {
	if x != nil {
		return x.Status
	}
	return false
}

func (x *ExitResponse) GetError() string {
	if x != nil && x.Error != nil {
		return *x.Error
	}
	return ""
}

var File_user_methods_exit_proto protoreflect.FileDescriptor

var file_user_methods_exit_proto_rawDesc = []byte{
	0x0a, 0x17, 0x75, 0x73, 0x65, 0x72, 0x2f, 0x6d, 0x65, 0x74, 0x68, 0x6f, 0x64, 0x73, 0x2f, 0x65,
	0x78, 0x69, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0c, 0x75, 0x73, 0x65, 0x72, 0x2e,
	0x6d, 0x65, 0x74, 0x68, 0x6f, 0x64, 0x73, 0x1a, 0x1a, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2f,
	0x72, 0x65, 0x66, 0x72, 0x65, 0x73, 0x68, 0x5f, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x22, 0x48, 0x0a, 0x0b, 0x45, 0x78, 0x69, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x12, 0x39, 0x0a, 0x0d, 0x72, 0x65, 0x66, 0x72, 0x65, 0x73, 0x68, 0x5f, 0x74, 0x6f,
	0x6b, 0x65, 0x6e, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x14, 0x2e, 0x63, 0x6f, 0x6d, 0x6d,
	0x6f, 0x6e, 0x2e, 0x52, 0x65, 0x66, 0x72, 0x65, 0x73, 0x68, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x52,
	0x0c, 0x72, 0x65, 0x66, 0x72, 0x65, 0x73, 0x68, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x22, 0x4b, 0x0a,
	0x0c, 0x45, 0x78, 0x69, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x16, 0x0a,
	0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x18, 0x01, 0x20, 0x01, 0x28, 0x08, 0x52, 0x06, 0x73,
	0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0x19, 0x0a, 0x05, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x09, 0x48, 0x00, 0x52, 0x05, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x88, 0x01, 0x01,
	0x42, 0x08, 0x0a, 0x06, 0x5f, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x42, 0x36, 0x5a, 0x34, 0x67, 0x69,
	0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x59, 0x75, 0x78, 0x37, 0x37, 0x59, 0x75,
	0x78, 0x2f, 0x70, 0x6c, 0x61, 0x74, 0x66, 0x6f, 0x72, 0x6d, 0x5f, 0x62, 0x61, 0x63, 0x6b, 0x65,
	0x6e, 0x64, 0x2f, 0x67, 0x65, 0x6e, 0x65, 0x72, 0x61, 0x74, 0x65, 0x64, 0x2f, 0x75, 0x73, 0x65,
	0x72, 0x3b, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_user_methods_exit_proto_rawDescOnce sync.Once
	file_user_methods_exit_proto_rawDescData = file_user_methods_exit_proto_rawDesc
)

func file_user_methods_exit_proto_rawDescGZIP() []byte {
	file_user_methods_exit_proto_rawDescOnce.Do(func() {
		file_user_methods_exit_proto_rawDescData = protoimpl.X.CompressGZIP(file_user_methods_exit_proto_rawDescData)
	})
	return file_user_methods_exit_proto_rawDescData
}

var file_user_methods_exit_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_user_methods_exit_proto_goTypes = []any{
	(*ExitRequest)(nil),         // 0: user.methods.ExitRequest
	(*ExitResponse)(nil),        // 1: user.methods.ExitResponse
	(*common.RefreshToken)(nil), // 2: common.RefreshToken
}
var file_user_methods_exit_proto_depIdxs = []int32{
	2, // 0: user.methods.ExitRequest.refresh_token:type_name -> common.RefreshToken
	1, // [1:1] is the sub-list for method output_type
	1, // [1:1] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_user_methods_exit_proto_init() }
func file_user_methods_exit_proto_init() {
	if File_user_methods_exit_proto != nil {
		return
	}
	file_user_methods_exit_proto_msgTypes[1].OneofWrappers = []any{}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_user_methods_exit_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_user_methods_exit_proto_goTypes,
		DependencyIndexes: file_user_methods_exit_proto_depIdxs,
		MessageInfos:      file_user_methods_exit_proto_msgTypes,
	}.Build()
	File_user_methods_exit_proto = out.File
	file_user_methods_exit_proto_rawDesc = nil
	file_user_methods_exit_proto_goTypes = nil
	file_user_methods_exit_proto_depIdxs = nil
}
