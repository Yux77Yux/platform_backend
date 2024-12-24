// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.35.1
// 	protoc        v5.28.3
// source: creation/methods/post.proto

package creation

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

type UploadCreationRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	BaseInfo    *CreationUpload     `protobuf:"bytes,1,opt,name=base_info,json=baseInfo,proto3" json:"base_info,omitempty"`
	AccessToken *common.AccessToken `protobuf:"bytes,2,opt,name=access_token,json=accessToken,proto3" json:"access_token,omitempty"`
}

func (x *UploadCreationRequest) Reset() {
	*x = UploadCreationRequest{}
	mi := &file_creation_methods_post_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *UploadCreationRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UploadCreationRequest) ProtoMessage() {}

func (x *UploadCreationRequest) ProtoReflect() protoreflect.Message {
	mi := &file_creation_methods_post_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UploadCreationRequest.ProtoReflect.Descriptor instead.
func (*UploadCreationRequest) Descriptor() ([]byte, []int) {
	return file_creation_methods_post_proto_rawDescGZIP(), []int{0}
}

func (x *UploadCreationRequest) GetBaseInfo() *CreationUpload {
	if x != nil {
		return x.BaseInfo
	}
	return nil
}

func (x *UploadCreationRequest) GetAccessToken() *common.AccessToken {
	if x != nil {
		return x.AccessToken
	}
	return nil
}

type UploadCreationResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Msg *common.ApiResponse `protobuf:"bytes,1,opt,name=msg,proto3" json:"msg,omitempty"`
}

func (x *UploadCreationResponse) Reset() {
	*x = UploadCreationResponse{}
	mi := &file_creation_methods_post_proto_msgTypes[1]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *UploadCreationResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UploadCreationResponse) ProtoMessage() {}

func (x *UploadCreationResponse) ProtoReflect() protoreflect.Message {
	mi := &file_creation_methods_post_proto_msgTypes[1]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UploadCreationResponse.ProtoReflect.Descriptor instead.
func (*UploadCreationResponse) Descriptor() ([]byte, []int) {
	return file_creation_methods_post_proto_rawDescGZIP(), []int{1}
}

func (x *UploadCreationResponse) GetMsg() *common.ApiResponse {
	if x != nil {
		return x.Msg
	}
	return nil
}

var File_creation_methods_post_proto protoreflect.FileDescriptor

var file_creation_methods_post_proto_rawDesc = []byte{
	0x0a, 0x1b, 0x63, 0x72, 0x65, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2f, 0x6d, 0x65, 0x74, 0x68, 0x6f,
	0x64, 0x73, 0x2f, 0x70, 0x6f, 0x73, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x10, 0x63,
	0x72, 0x65, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x6d, 0x65, 0x74, 0x68, 0x6f, 0x64, 0x73, 0x1a,
	0x19, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2f, 0x61, 0x70, 0x69, 0x5f, 0x72, 0x65, 0x73, 0x70,
	0x6f, 0x6e, 0x73, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x27, 0x63, 0x72, 0x65, 0x61,
	0x74, 0x69, 0x6f, 0x6e, 0x2f, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x73, 0x2f, 0x63, 0x72,
	0x65, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x75, 0x70, 0x6c, 0x6f, 0x61, 0x64, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x1a, 0x19, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2f, 0x61, 0x63, 0x63, 0x65,
	0x73, 0x73, 0x5f, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x8f,
	0x01, 0x0a, 0x15, 0x55, 0x70, 0x6c, 0x6f, 0x61, 0x64, 0x43, 0x72, 0x65, 0x61, 0x74, 0x69, 0x6f,
	0x6e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x3e, 0x0a, 0x09, 0x62, 0x61, 0x73, 0x65,
	0x5f, 0x69, 0x6e, 0x66, 0x6f, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x21, 0x2e, 0x63, 0x72,
	0x65, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x73, 0x2e,
	0x43, 0x72, 0x65, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x55, 0x70, 0x6c, 0x6f, 0x61, 0x64, 0x52, 0x08,
	0x62, 0x61, 0x73, 0x65, 0x49, 0x6e, 0x66, 0x6f, 0x12, 0x36, 0x0a, 0x0c, 0x61, 0x63, 0x63, 0x65,
	0x73, 0x73, 0x5f, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x13,
	0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x41, 0x63, 0x63, 0x65, 0x73, 0x73, 0x54, 0x6f,
	0x6b, 0x65, 0x6e, 0x52, 0x0b, 0x61, 0x63, 0x63, 0x65, 0x73, 0x73, 0x54, 0x6f, 0x6b, 0x65, 0x6e,
	0x22, 0x3f, 0x0a, 0x16, 0x55, 0x70, 0x6c, 0x6f, 0x61, 0x64, 0x43, 0x72, 0x65, 0x61, 0x74, 0x69,
	0x6f, 0x6e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x25, 0x0a, 0x03, 0x6d, 0x73,
	0x67, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x13, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e,
	0x2e, 0x41, 0x70, 0x69, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x52, 0x03, 0x6d, 0x73,
	0x67, 0x42, 0x39, 0x5a, 0x37, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f,
	0x59, 0x75, 0x78, 0x37, 0x37, 0x59, 0x75, 0x78, 0x2f, 0x70, 0x6c, 0x61, 0x74, 0x66, 0x6f, 0x72,
	0x6d, 0x5f, 0x62, 0x61, 0x63, 0x6b, 0x65, 0x6e, 0x64, 0x2f, 0x67, 0x65, 0x6e, 0x65, 0x72, 0x61,
	0x74, 0x65, 0x64, 0x2f, 0x63, 0x72, 0x65, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x62, 0x06, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_creation_methods_post_proto_rawDescOnce sync.Once
	file_creation_methods_post_proto_rawDescData = file_creation_methods_post_proto_rawDesc
)

func file_creation_methods_post_proto_rawDescGZIP() []byte {
	file_creation_methods_post_proto_rawDescOnce.Do(func() {
		file_creation_methods_post_proto_rawDescData = protoimpl.X.CompressGZIP(file_creation_methods_post_proto_rawDescData)
	})
	return file_creation_methods_post_proto_rawDescData
}

var file_creation_methods_post_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_creation_methods_post_proto_goTypes = []any{
	(*UploadCreationRequest)(nil),  // 0: creation.methods.UploadCreationRequest
	(*UploadCreationResponse)(nil), // 1: creation.methods.UploadCreationResponse
	(*CreationUpload)(nil),         // 2: creation.messages.CreationUpload
	(*common.AccessToken)(nil),     // 3: common.AccessToken
	(*common.ApiResponse)(nil),     // 4: common.ApiResponse
}
var file_creation_methods_post_proto_depIdxs = []int32{
	2, // 0: creation.methods.UploadCreationRequest.base_info:type_name -> creation.messages.CreationUpload
	3, // 1: creation.methods.UploadCreationRequest.access_token:type_name -> common.AccessToken
	4, // 2: creation.methods.UploadCreationResponse.msg:type_name -> common.ApiResponse
	3, // [3:3] is the sub-list for method output_type
	3, // [3:3] is the sub-list for method input_type
	3, // [3:3] is the sub-list for extension type_name
	3, // [3:3] is the sub-list for extension extendee
	0, // [0:3] is the sub-list for field type_name
}

func init() { file_creation_methods_post_proto_init() }
func file_creation_methods_post_proto_init() {
	if File_creation_methods_post_proto != nil {
		return
	}
	file_creation_messages_creation_upload_proto_init()
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_creation_methods_post_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_creation_methods_post_proto_goTypes,
		DependencyIndexes: file_creation_methods_post_proto_depIdxs,
		MessageInfos:      file_creation_methods_post_proto_msgTypes,
	}.Build()
	File_creation_methods_post_proto = out.File
	file_creation_methods_post_proto_rawDesc = nil
	file_creation_methods_post_proto_goTypes = nil
	file_creation_methods_post_proto_depIdxs = nil
}