// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.35.1
// 	protoc        v5.28.3
// source: creation/methods/get.proto

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

// 视频详细页
type GetCreationRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	CreationId  int64               `protobuf:"varint,1,opt,name=creation_id,json=creationId,proto3" json:"creation_id,omitempty"`
	AccessToken *common.AccessToken `protobuf:"bytes,2,opt,name=access_token,json=accessToken,proto3" json:"access_token,omitempty"` // 个性化，登录后返回的
}

func (x *GetCreationRequest) Reset() {
	*x = GetCreationRequest{}
	mi := &file_creation_methods_get_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *GetCreationRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetCreationRequest) ProtoMessage() {}

func (x *GetCreationRequest) ProtoReflect() protoreflect.Message {
	mi := &file_creation_methods_get_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetCreationRequest.ProtoReflect.Descriptor instead.
func (*GetCreationRequest) Descriptor() ([]byte, []int) {
	return file_creation_methods_get_proto_rawDescGZIP(), []int{0}
}

func (x *GetCreationRequest) GetCreationId() int64 {
	if x != nil {
		return x.CreationId
	}
	return 0
}

func (x *GetCreationRequest) GetAccessToken() *common.AccessToken {
	if x != nil {
		return x.AccessToken
	}
	return nil
}

type GetCreationResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	CreationInfo *CreationInfo       `protobuf:"bytes,1,opt,name=creation_info,json=creationInfo,proto3" json:"creation_info,omitempty"`
	Msg          *common.ApiResponse `protobuf:"bytes,2,opt,name=msg,proto3" json:"msg,omitempty"`
}

func (x *GetCreationResponse) Reset() {
	*x = GetCreationResponse{}
	mi := &file_creation_methods_get_proto_msgTypes[1]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *GetCreationResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetCreationResponse) ProtoMessage() {}

func (x *GetCreationResponse) ProtoReflect() protoreflect.Message {
	mi := &file_creation_methods_get_proto_msgTypes[1]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetCreationResponse.ProtoReflect.Descriptor instead.
func (*GetCreationResponse) Descriptor() ([]byte, []int) {
	return file_creation_methods_get_proto_rawDescGZIP(), []int{1}
}

func (x *GetCreationResponse) GetCreationInfo() *CreationInfo {
	if x != nil {
		return x.CreationInfo
	}
	return nil
}

func (x *GetCreationResponse) GetMsg() *common.ApiResponse {
	if x != nil {
		return x.Msg
	}
	return nil
}

// 空间的作品
type GetSpaceCreationListRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	UserId      int64               `protobuf:"varint,1,opt,name=user_id,json=userId,proto3" json:"user_id,omitempty"`
	AccessToken *common.AccessToken `protobuf:"bytes,2,opt,name=access_token,json=accessToken,proto3" json:"access_token,omitempty"` // 个性化，登录后返回的
}

func (x *GetSpaceCreationListRequest) Reset() {
	*x = GetSpaceCreationListRequest{}
	mi := &file_creation_methods_get_proto_msgTypes[2]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *GetSpaceCreationListRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetSpaceCreationListRequest) ProtoMessage() {}

func (x *GetSpaceCreationListRequest) ProtoReflect() protoreflect.Message {
	mi := &file_creation_methods_get_proto_msgTypes[2]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetSpaceCreationListRequest.ProtoReflect.Descriptor instead.
func (*GetSpaceCreationListRequest) Descriptor() ([]byte, []int) {
	return file_creation_methods_get_proto_rawDescGZIP(), []int{2}
}

func (x *GetSpaceCreationListRequest) GetUserId() int64 {
	if x != nil {
		return x.UserId
	}
	return 0
}

func (x *GetSpaceCreationListRequest) GetAccessToken() *common.AccessToken {
	if x != nil {
		return x.AccessToken
	}
	return nil
}

// 主页的推荐视频
type GetHomePageCreationListRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// 请求的视频数组
	CreationId []int64 `protobuf:"varint,1,rep,packed,name=creation_id,json=creationId,proto3" json:"creation_id,omitempty"`
}

func (x *GetHomePageCreationListRequest) Reset() {
	*x = GetHomePageCreationListRequest{}
	mi := &file_creation_methods_get_proto_msgTypes[3]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *GetHomePageCreationListRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetHomePageCreationListRequest) ProtoMessage() {}

func (x *GetHomePageCreationListRequest) ProtoReflect() protoreflect.Message {
	mi := &file_creation_methods_get_proto_msgTypes[3]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetHomePageCreationListRequest.ProtoReflect.Descriptor instead.
func (*GetHomePageCreationListRequest) Descriptor() ([]byte, []int) {
	return file_creation_methods_get_proto_rawDescGZIP(), []int{3}
}

func (x *GetHomePageCreationListRequest) GetCreationId() []int64 {
	if x != nil {
		return x.CreationId
	}
	return nil
}

type GetCreationListResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Msg          *common.ApiResponse `protobuf:"bytes,1,opt,name=msg,proto3" json:"msg,omitempty"`
	CreationInfo *CreationInfo       `protobuf:"bytes,2,opt,name=creation_info,json=creationInfo,proto3" json:"creation_info,omitempty"`
}

func (x *GetCreationListResponse) Reset() {
	*x = GetCreationListResponse{}
	mi := &file_creation_methods_get_proto_msgTypes[4]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *GetCreationListResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetCreationListResponse) ProtoMessage() {}

func (x *GetCreationListResponse) ProtoReflect() protoreflect.Message {
	mi := &file_creation_methods_get_proto_msgTypes[4]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetCreationListResponse.ProtoReflect.Descriptor instead.
func (*GetCreationListResponse) Descriptor() ([]byte, []int) {
	return file_creation_methods_get_proto_rawDescGZIP(), []int{4}
}

func (x *GetCreationListResponse) GetMsg() *common.ApiResponse {
	if x != nil {
		return x.Msg
	}
	return nil
}

func (x *GetCreationListResponse) GetCreationInfo() *CreationInfo {
	if x != nil {
		return x.CreationInfo
	}
	return nil
}

var File_creation_methods_get_proto protoreflect.FileDescriptor

var file_creation_methods_get_proto_rawDesc = []byte{
	0x0a, 0x1a, 0x63, 0x72, 0x65, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2f, 0x6d, 0x65, 0x74, 0x68, 0x6f,
	0x64, 0x73, 0x2f, 0x67, 0x65, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x10, 0x63, 0x72,
	0x65, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x6d, 0x65, 0x74, 0x68, 0x6f, 0x64, 0x73, 0x1a, 0x19,
	0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2f, 0x61, 0x63, 0x63, 0x65, 0x73, 0x73, 0x5f, 0x74, 0x6f,
	0x6b, 0x65, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x19, 0x63, 0x6f, 0x6d, 0x6d, 0x6f,
	0x6e, 0x2f, 0x61, 0x70, 0x69, 0x5f, 0x72, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x20, 0x63, 0x72, 0x65, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2f, 0x6d,
	0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x73, 0x2f, 0x63, 0x72, 0x65, 0x61, 0x74, 0x69, 0x6f, 0x6e,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x6d, 0x0a, 0x12, 0x47, 0x65, 0x74, 0x43, 0x72, 0x65,
	0x61, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x1f, 0x0a, 0x0b,
	0x63, 0x72, 0x65, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x03, 0x52, 0x0a, 0x63, 0x72, 0x65, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x49, 0x64, 0x12, 0x36, 0x0a,
	0x0c, 0x61, 0x63, 0x63, 0x65, 0x73, 0x73, 0x5f, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x0b, 0x32, 0x13, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x41, 0x63, 0x63,
	0x65, 0x73, 0x73, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x52, 0x0b, 0x61, 0x63, 0x63, 0x65, 0x73, 0x73,
	0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x22, 0x82, 0x01, 0x0a, 0x13, 0x47, 0x65, 0x74, 0x43, 0x72, 0x65,
	0x61, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x44, 0x0a,
	0x0d, 0x63, 0x72, 0x65, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x69, 0x6e, 0x66, 0x6f, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x0b, 0x32, 0x1f, 0x2e, 0x63, 0x72, 0x65, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2e,
	0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x73, 0x2e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x69, 0x6f,
	0x6e, 0x49, 0x6e, 0x66, 0x6f, 0x52, 0x0c, 0x63, 0x72, 0x65, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x49,
	0x6e, 0x66, 0x6f, 0x12, 0x25, 0x0a, 0x03, 0x6d, 0x73, 0x67, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b,
	0x32, 0x13, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x41, 0x70, 0x69, 0x52, 0x65, 0x73,
	0x70, 0x6f, 0x6e, 0x73, 0x65, 0x52, 0x03, 0x6d, 0x73, 0x67, 0x22, 0x6e, 0x0a, 0x1b, 0x47, 0x65,
	0x74, 0x53, 0x70, 0x61, 0x63, 0x65, 0x43, 0x72, 0x65, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x4c, 0x69,
	0x73, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x17, 0x0a, 0x07, 0x75, 0x73, 0x65,
	0x72, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x06, 0x75, 0x73, 0x65, 0x72,
	0x49, 0x64, 0x12, 0x36, 0x0a, 0x0c, 0x61, 0x63, 0x63, 0x65, 0x73, 0x73, 0x5f, 0x74, 0x6f, 0x6b,
	0x65, 0x6e, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x13, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f,
	0x6e, 0x2e, 0x41, 0x63, 0x63, 0x65, 0x73, 0x73, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x52, 0x0b, 0x61,
	0x63, 0x63, 0x65, 0x73, 0x73, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x22, 0x41, 0x0a, 0x1e, 0x47, 0x65,
	0x74, 0x48, 0x6f, 0x6d, 0x65, 0x50, 0x61, 0x67, 0x65, 0x43, 0x72, 0x65, 0x61, 0x74, 0x69, 0x6f,
	0x6e, 0x4c, 0x69, 0x73, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x1f, 0x0a, 0x0b,
	0x63, 0x72, 0x65, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x03, 0x28,
	0x03, 0x52, 0x0a, 0x63, 0x72, 0x65, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x49, 0x64, 0x22, 0x86, 0x01,
	0x0a, 0x17, 0x47, 0x65, 0x74, 0x43, 0x72, 0x65, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x4c, 0x69, 0x73,
	0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x25, 0x0a, 0x03, 0x6d, 0x73, 0x67,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x13, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e,
	0x41, 0x70, 0x69, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x52, 0x03, 0x6d, 0x73, 0x67,
	0x12, 0x44, 0x0a, 0x0d, 0x63, 0x72, 0x65, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x69, 0x6e, 0x66,
	0x6f, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1f, 0x2e, 0x63, 0x72, 0x65, 0x61, 0x74, 0x69,
	0x6f, 0x6e, 0x2e, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x73, 0x2e, 0x43, 0x72, 0x65, 0x61,
	0x74, 0x69, 0x6f, 0x6e, 0x49, 0x6e, 0x66, 0x6f, 0x52, 0x0c, 0x63, 0x72, 0x65, 0x61, 0x74, 0x69,
	0x6f, 0x6e, 0x49, 0x6e, 0x66, 0x6f, 0x42, 0x39, 0x5a, 0x37, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62,
	0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x59, 0x75, 0x78, 0x37, 0x37, 0x59, 0x75, 0x78, 0x2f, 0x70, 0x6c,
	0x61, 0x74, 0x66, 0x6f, 0x72, 0x6d, 0x5f, 0x62, 0x61, 0x63, 0x6b, 0x65, 0x6e, 0x64, 0x2f, 0x67,
	0x65, 0x6e, 0x65, 0x72, 0x61, 0x74, 0x65, 0x64, 0x2f, 0x63, 0x72, 0x65, 0x61, 0x74, 0x69, 0x6f,
	0x6e, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_creation_methods_get_proto_rawDescOnce sync.Once
	file_creation_methods_get_proto_rawDescData = file_creation_methods_get_proto_rawDesc
)

func file_creation_methods_get_proto_rawDescGZIP() []byte {
	file_creation_methods_get_proto_rawDescOnce.Do(func() {
		file_creation_methods_get_proto_rawDescData = protoimpl.X.CompressGZIP(file_creation_methods_get_proto_rawDescData)
	})
	return file_creation_methods_get_proto_rawDescData
}

var file_creation_methods_get_proto_msgTypes = make([]protoimpl.MessageInfo, 5)
var file_creation_methods_get_proto_goTypes = []any{
	(*GetCreationRequest)(nil),             // 0: creation.methods.GetCreationRequest
	(*GetCreationResponse)(nil),            // 1: creation.methods.GetCreationResponse
	(*GetSpaceCreationListRequest)(nil),    // 2: creation.methods.GetSpaceCreationListRequest
	(*GetHomePageCreationListRequest)(nil), // 3: creation.methods.GetHomePageCreationListRequest
	(*GetCreationListResponse)(nil),        // 4: creation.methods.GetCreationListResponse
	(*common.AccessToken)(nil),             // 5: common.AccessToken
	(*CreationInfo)(nil),                   // 6: creation.messages.CreationInfo
	(*common.ApiResponse)(nil),             // 7: common.ApiResponse
}
var file_creation_methods_get_proto_depIdxs = []int32{
	5, // 0: creation.methods.GetCreationRequest.access_token:type_name -> common.AccessToken
	6, // 1: creation.methods.GetCreationResponse.creation_info:type_name -> creation.messages.CreationInfo
	7, // 2: creation.methods.GetCreationResponse.msg:type_name -> common.ApiResponse
	5, // 3: creation.methods.GetSpaceCreationListRequest.access_token:type_name -> common.AccessToken
	7, // 4: creation.methods.GetCreationListResponse.msg:type_name -> common.ApiResponse
	6, // 5: creation.methods.GetCreationListResponse.creation_info:type_name -> creation.messages.CreationInfo
	6, // [6:6] is the sub-list for method output_type
	6, // [6:6] is the sub-list for method input_type
	6, // [6:6] is the sub-list for extension type_name
	6, // [6:6] is the sub-list for extension extendee
	0, // [0:6] is the sub-list for field type_name
}

func init() { file_creation_methods_get_proto_init() }
func file_creation_methods_get_proto_init() {
	if File_creation_methods_get_proto != nil {
		return
	}
	file_creation_messages_creation_proto_init()
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_creation_methods_get_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   5,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_creation_methods_get_proto_goTypes,
		DependencyIndexes: file_creation_methods_get_proto_depIdxs,
		MessageInfos:      file_creation_methods_get_proto_msgTypes,
	}.Build()
	File_creation_methods_get_proto = out.File
	file_creation_methods_get_proto_rawDesc = nil
	file_creation_methods_get_proto_goTypes = nil
	file_creation_methods_get_proto_depIdxs = nil
}
