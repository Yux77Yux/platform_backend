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

type GetSpaceCreationsRequest_ByCount int32

const (
	GetSpaceCreationsRequest_PUBLISHED_TIME GetSpaceCreationsRequest_ByCount = 0
	GetSpaceCreationsRequest_VIEWS          GetSpaceCreationsRequest_ByCount = 1
	GetSpaceCreationsRequest_LIKES          GetSpaceCreationsRequest_ByCount = 2
	GetSpaceCreationsRequest_COLLECTIONS    GetSpaceCreationsRequest_ByCount = 3
)

// Enum value maps for GetSpaceCreationsRequest_ByCount.
var (
	GetSpaceCreationsRequest_ByCount_name = map[int32]string{
		0: "PUBLISHED_TIME",
		1: "VIEWS",
		2: "LIKES",
		3: "COLLECTIONS",
	}
	GetSpaceCreationsRequest_ByCount_value = map[string]int32{
		"PUBLISHED_TIME": 0,
		"VIEWS":          1,
		"LIKES":          2,
		"COLLECTIONS":    3,
	}
)

func (x GetSpaceCreationsRequest_ByCount) Enum() *GetSpaceCreationsRequest_ByCount {
	p := new(GetSpaceCreationsRequest_ByCount)
	*p = x
	return p
}

func (x GetSpaceCreationsRequest_ByCount) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (GetSpaceCreationsRequest_ByCount) Descriptor() protoreflect.EnumDescriptor {
	return file_creation_methods_get_proto_enumTypes[0].Descriptor()
}

func (GetSpaceCreationsRequest_ByCount) Type() protoreflect.EnumType {
	return &file_creation_methods_get_proto_enumTypes[0]
}

func (x GetSpaceCreationsRequest_ByCount) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use GetSpaceCreationsRequest_ByCount.Descriptor instead.
func (GetSpaceCreationsRequest_ByCount) EnumDescriptor() ([]byte, []int) {
	return file_creation_methods_get_proto_rawDescGZIP(), []int{2, 0}
}

// 视频详细页
type GetCreationRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	CreationId int64 `protobuf:"varint,1,opt,name=creation_id,json=creationId,proto3" json:"creation_id,omitempty"`
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

// 空间
type GetSpaceCreationsRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	UserId int64                            `protobuf:"varint,1,opt,name=user_id,json=userId,proto3" json:"user_id,omitempty"`
	Page   int32                            `protobuf:"varint,2,opt,name=page,proto3" json:"page,omitempty"`
	ByWhat GetSpaceCreationsRequest_ByCount `protobuf:"varint,3,opt,name=by_what,json=byWhat,proto3,enum=creation.methods.GetSpaceCreationsRequest_ByCount" json:"by_what,omitempty"`
}

func (x *GetSpaceCreationsRequest) Reset() {
	*x = GetSpaceCreationsRequest{}
	mi := &file_creation_methods_get_proto_msgTypes[2]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *GetSpaceCreationsRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetSpaceCreationsRequest) ProtoMessage() {}

func (x *GetSpaceCreationsRequest) ProtoReflect() protoreflect.Message {
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

// Deprecated: Use GetSpaceCreationsRequest.ProtoReflect.Descriptor instead.
func (*GetSpaceCreationsRequest) Descriptor() ([]byte, []int) {
	return file_creation_methods_get_proto_rawDescGZIP(), []int{2}
}

func (x *GetSpaceCreationsRequest) GetUserId() int64 {
	if x != nil {
		return x.UserId
	}
	return 0
}

func (x *GetSpaceCreationsRequest) GetPage() int32 {
	if x != nil {
		return x.Page
	}
	return 0
}

func (x *GetSpaceCreationsRequest) GetByWhat() GetSpaceCreationsRequest_ByCount {
	if x != nil {
		return x.ByWhat
	}
	return GetSpaceCreationsRequest_PUBLISHED_TIME
}

// 发布状态的Creation列表
type GetCreationListRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Ids []int64 `protobuf:"varint,1,rep,packed,name=ids,proto3" json:"ids,omitempty"`
}

func (x *GetCreationListRequest) Reset() {
	*x = GetCreationListRequest{}
	mi := &file_creation_methods_get_proto_msgTypes[3]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *GetCreationListRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetCreationListRequest) ProtoMessage() {}

func (x *GetCreationListRequest) ProtoReflect() protoreflect.Message {
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

// Deprecated: Use GetCreationListRequest.ProtoReflect.Descriptor instead.
func (*GetCreationListRequest) Descriptor() ([]byte, []int) {
	return file_creation_methods_get_proto_rawDescGZIP(), []int{3}
}

func (x *GetCreationListRequest) GetIds() []int64 {
	if x != nil {
		return x.Ids
	}
	return nil
}

type GetCreationListResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Msg               *common.ApiResponse `protobuf:"bytes,1,opt,name=msg,proto3" json:"msg,omitempty"`
	CreationInfoGroup []*CreationInfo     `protobuf:"bytes,2,rep,name=creation_info_group,json=creationInfoGroup,proto3" json:"creation_info_group,omitempty"`
	Count             int32               `protobuf:"varint,3,opt,name=count,proto3" json:"count,omitempty"`
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

func (x *GetCreationListResponse) GetCreationInfoGroup() []*CreationInfo {
	if x != nil {
		return x.CreationInfoGroup
	}
	return nil
}

func (x *GetCreationListResponse) GetCount() int32 {
	if x != nil {
		return x.Count
	}
	return 0
}

var File_creation_methods_get_proto protoreflect.FileDescriptor

var file_creation_methods_get_proto_rawDesc = []byte{
	0x0a, 0x1a, 0x63, 0x72, 0x65, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2f, 0x6d, 0x65, 0x74, 0x68, 0x6f,
	0x64, 0x73, 0x2f, 0x67, 0x65, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x10, 0x63, 0x72,
	0x65, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x6d, 0x65, 0x74, 0x68, 0x6f, 0x64, 0x73, 0x1a, 0x19,
	0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2f, 0x61, 0x70, 0x69, 0x5f, 0x72, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x73, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x20, 0x63, 0x72, 0x65, 0x61, 0x74,
	0x69, 0x6f, 0x6e, 0x2f, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x73, 0x2f, 0x63, 0x72, 0x65,
	0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x35, 0x0a, 0x12, 0x47,
	0x65, 0x74, 0x43, 0x72, 0x65, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x12, 0x1f, 0x0a, 0x0b, 0x63, 0x72, 0x65, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x69, 0x64,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x0a, 0x63, 0x72, 0x65, 0x61, 0x74, 0x69, 0x6f, 0x6e,
	0x49, 0x64, 0x22, 0x82, 0x01, 0x0a, 0x13, 0x47, 0x65, 0x74, 0x43, 0x72, 0x65, 0x61, 0x74, 0x69,
	0x6f, 0x6e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x44, 0x0a, 0x0d, 0x63, 0x72,
	0x65, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x69, 0x6e, 0x66, 0x6f, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x0b, 0x32, 0x1f, 0x2e, 0x63, 0x72, 0x65, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x6d, 0x65, 0x73,
	0x73, 0x61, 0x67, 0x65, 0x73, 0x2e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x49, 0x6e,
	0x66, 0x6f, 0x52, 0x0c, 0x63, 0x72, 0x65, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x49, 0x6e, 0x66, 0x6f,
	0x12, 0x25, 0x0a, 0x03, 0x6d, 0x73, 0x67, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x13, 0x2e,
	0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x41, 0x70, 0x69, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e,
	0x73, 0x65, 0x52, 0x03, 0x6d, 0x73, 0x67, 0x22, 0xda, 0x01, 0x0a, 0x18, 0x47, 0x65, 0x74, 0x53,
	0x70, 0x61, 0x63, 0x65, 0x43, 0x72, 0x65, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x12, 0x17, 0x0a, 0x07, 0x75, 0x73, 0x65, 0x72, 0x5f, 0x69, 0x64, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x06, 0x75, 0x73, 0x65, 0x72, 0x49, 0x64, 0x12, 0x12, 0x0a,
	0x04, 0x70, 0x61, 0x67, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x05, 0x52, 0x04, 0x70, 0x61, 0x67,
	0x65, 0x12, 0x4b, 0x0a, 0x07, 0x62, 0x79, 0x5f, 0x77, 0x68, 0x61, 0x74, 0x18, 0x03, 0x20, 0x01,
	0x28, 0x0e, 0x32, 0x32, 0x2e, 0x63, 0x72, 0x65, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x6d, 0x65,
	0x74, 0x68, 0x6f, 0x64, 0x73, 0x2e, 0x47, 0x65, 0x74, 0x53, 0x70, 0x61, 0x63, 0x65, 0x43, 0x72,
	0x65, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x2e, 0x42,
	0x79, 0x43, 0x6f, 0x75, 0x6e, 0x74, 0x52, 0x06, 0x62, 0x79, 0x57, 0x68, 0x61, 0x74, 0x22, 0x44,
	0x0a, 0x07, 0x42, 0x79, 0x43, 0x6f, 0x75, 0x6e, 0x74, 0x12, 0x12, 0x0a, 0x0e, 0x50, 0x55, 0x42,
	0x4c, 0x49, 0x53, 0x48, 0x45, 0x44, 0x5f, 0x54, 0x49, 0x4d, 0x45, 0x10, 0x00, 0x12, 0x09, 0x0a,
	0x05, 0x56, 0x49, 0x45, 0x57, 0x53, 0x10, 0x01, 0x12, 0x09, 0x0a, 0x05, 0x4c, 0x49, 0x4b, 0x45,
	0x53, 0x10, 0x02, 0x12, 0x0f, 0x0a, 0x0b, 0x43, 0x4f, 0x4c, 0x4c, 0x45, 0x43, 0x54, 0x49, 0x4f,
	0x4e, 0x53, 0x10, 0x03, 0x22, 0x2a, 0x0a, 0x16, 0x47, 0x65, 0x74, 0x43, 0x72, 0x65, 0x61, 0x74,
	0x69, 0x6f, 0x6e, 0x4c, 0x69, 0x73, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x10,
	0x0a, 0x03, 0x69, 0x64, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x03, 0x52, 0x03, 0x69, 0x64, 0x73,
	0x22, 0xa7, 0x01, 0x0a, 0x17, 0x47, 0x65, 0x74, 0x43, 0x72, 0x65, 0x61, 0x74, 0x69, 0x6f, 0x6e,
	0x4c, 0x69, 0x73, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x25, 0x0a, 0x03,
	0x6d, 0x73, 0x67, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x13, 0x2e, 0x63, 0x6f, 0x6d, 0x6d,
	0x6f, 0x6e, 0x2e, 0x41, 0x70, 0x69, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x52, 0x03,
	0x6d, 0x73, 0x67, 0x12, 0x4f, 0x0a, 0x13, 0x63, 0x72, 0x65, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x5f,
	0x69, 0x6e, 0x66, 0x6f, 0x5f, 0x67, 0x72, 0x6f, 0x75, 0x70, 0x18, 0x02, 0x20, 0x03, 0x28, 0x0b,
	0x32, 0x1f, 0x2e, 0x63, 0x72, 0x65, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x6d, 0x65, 0x73, 0x73,
	0x61, 0x67, 0x65, 0x73, 0x2e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x49, 0x6e, 0x66,
	0x6f, 0x52, 0x11, 0x63, 0x72, 0x65, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x49, 0x6e, 0x66, 0x6f, 0x47,
	0x72, 0x6f, 0x75, 0x70, 0x12, 0x14, 0x0a, 0x05, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x18, 0x03, 0x20,
	0x01, 0x28, 0x05, 0x52, 0x05, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x42, 0x39, 0x5a, 0x37, 0x67, 0x69,
	0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x59, 0x75, 0x78, 0x37, 0x37, 0x59, 0x75,
	0x78, 0x2f, 0x70, 0x6c, 0x61, 0x74, 0x66, 0x6f, 0x72, 0x6d, 0x5f, 0x62, 0x61, 0x63, 0x6b, 0x65,
	0x6e, 0x64, 0x2f, 0x67, 0x65, 0x6e, 0x65, 0x72, 0x61, 0x74, 0x65, 0x64, 0x2f, 0x63, 0x72, 0x65,
	0x61, 0x74, 0x69, 0x6f, 0x6e, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
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

var file_creation_methods_get_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_creation_methods_get_proto_msgTypes = make([]protoimpl.MessageInfo, 5)
var file_creation_methods_get_proto_goTypes = []any{
	(GetSpaceCreationsRequest_ByCount)(0), // 0: creation.methods.GetSpaceCreationsRequest.ByCount
	(*GetCreationRequest)(nil),            // 1: creation.methods.GetCreationRequest
	(*GetCreationResponse)(nil),           // 2: creation.methods.GetCreationResponse
	(*GetSpaceCreationsRequest)(nil),      // 3: creation.methods.GetSpaceCreationsRequest
	(*GetCreationListRequest)(nil),        // 4: creation.methods.GetCreationListRequest
	(*GetCreationListResponse)(nil),       // 5: creation.methods.GetCreationListResponse
	(*CreationInfo)(nil),                  // 6: creation.messages.CreationInfo
	(*common.ApiResponse)(nil),            // 7: common.ApiResponse
}
var file_creation_methods_get_proto_depIdxs = []int32{
	6, // 0: creation.methods.GetCreationResponse.creation_info:type_name -> creation.messages.CreationInfo
	7, // 1: creation.methods.GetCreationResponse.msg:type_name -> common.ApiResponse
	0, // 2: creation.methods.GetSpaceCreationsRequest.by_what:type_name -> creation.methods.GetSpaceCreationsRequest.ByCount
	7, // 3: creation.methods.GetCreationListResponse.msg:type_name -> common.ApiResponse
	6, // 4: creation.methods.GetCreationListResponse.creation_info_group:type_name -> creation.messages.CreationInfo
	5, // [5:5] is the sub-list for method output_type
	5, // [5:5] is the sub-list for method input_type
	5, // [5:5] is the sub-list for extension type_name
	5, // [5:5] is the sub-list for extension extendee
	0, // [0:5] is the sub-list for field type_name
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
			NumEnums:      1,
			NumMessages:   5,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_creation_methods_get_proto_goTypes,
		DependencyIndexes: file_creation_methods_get_proto_depIdxs,
		EnumInfos:         file_creation_methods_get_proto_enumTypes,
		MessageInfos:      file_creation_methods_get_proto_msgTypes,
	}.Build()
	File_creation_methods_get_proto = out.File
	file_creation_methods_get_proto_rawDesc = nil
	file_creation_methods_get_proto_goTypes = nil
	file_creation_methods_get_proto_depIdxs = nil
}
