// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.35.1
// 	protoc        v5.28.3
// source: comment/methods/get.proto

package comment

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

type GetCommentsRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Ids []int32 `protobuf:"varint,1,rep,packed,name=ids,proto3" json:"ids,omitempty"`
}

func (x *GetCommentsRequest) Reset() {
	*x = GetCommentsRequest{}
	mi := &file_comment_methods_get_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *GetCommentsRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetCommentsRequest) ProtoMessage() {}

func (x *GetCommentsRequest) ProtoReflect() protoreflect.Message {
	mi := &file_comment_methods_get_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetCommentsRequest.ProtoReflect.Descriptor instead.
func (*GetCommentsRequest) Descriptor() ([]byte, []int) {
	return file_comment_methods_get_proto_rawDescGZIP(), []int{0}
}

func (x *GetCommentsRequest) GetIds() []int32 {
	if x != nil {
		return x.Ids
	}
	return nil
}

// 初始化
type InitialCommentsRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	CreationId int64 `protobuf:"varint,1,opt,name=creation_id,json=creationId,proto3" json:"creation_id,omitempty"`
}

func (x *InitialCommentsRequest) Reset() {
	*x = InitialCommentsRequest{}
	mi := &file_comment_methods_get_proto_msgTypes[1]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *InitialCommentsRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*InitialCommentsRequest) ProtoMessage() {}

func (x *InitialCommentsRequest) ProtoReflect() protoreflect.Message {
	mi := &file_comment_methods_get_proto_msgTypes[1]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use InitialCommentsRequest.ProtoReflect.Descriptor instead.
func (*InitialCommentsRequest) Descriptor() ([]byte, []int) {
	return file_comment_methods_get_proto_rawDescGZIP(), []int{1}
}

func (x *InitialCommentsRequest) GetCreationId() int64 {
	if x != nil {
		return x.CreationId
	}
	return 0
}

type InitialCommentsResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Comments    []*TopComment       `protobuf:"bytes,1,rep,name=comments,proto3" json:"comments,omitempty"`
	CommentArea *CommentArea        `protobuf:"bytes,2,opt,name=comment_area,json=commentArea,proto3" json:"comment_area,omitempty"`
	PageCount   int32               `protobuf:"varint,3,opt,name=page_count,json=pageCount,proto3" json:"page_count,omitempty"`
	Msg         *common.ApiResponse `protobuf:"bytes,4,opt,name=msg,proto3" json:"msg,omitempty"`
}

func (x *InitialCommentsResponse) Reset() {
	*x = InitialCommentsResponse{}
	mi := &file_comment_methods_get_proto_msgTypes[2]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *InitialCommentsResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*InitialCommentsResponse) ProtoMessage() {}

func (x *InitialCommentsResponse) ProtoReflect() protoreflect.Message {
	mi := &file_comment_methods_get_proto_msgTypes[2]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use InitialCommentsResponse.ProtoReflect.Descriptor instead.
func (*InitialCommentsResponse) Descriptor() ([]byte, []int) {
	return file_comment_methods_get_proto_rawDescGZIP(), []int{2}
}

func (x *InitialCommentsResponse) GetComments() []*TopComment {
	if x != nil {
		return x.Comments
	}
	return nil
}

func (x *InitialCommentsResponse) GetCommentArea() *CommentArea {
	if x != nil {
		return x.CommentArea
	}
	return nil
}

func (x *InitialCommentsResponse) GetPageCount() int32 {
	if x != nil {
		return x.PageCount
	}
	return 0
}

func (x *InitialCommentsResponse) GetMsg() *common.ApiResponse {
	if x != nil {
		return x.Msg
	}
	return nil
}

// 作品的一级评论
type GetTopCommentsRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	CreationId int64 `protobuf:"varint,1,opt,name=creation_id,json=creationId,proto3" json:"creation_id,omitempty"`
	Page       int32 `protobuf:"varint,2,opt,name=page,proto3" json:"page,omitempty"`
}

func (x *GetTopCommentsRequest) Reset() {
	*x = GetTopCommentsRequest{}
	mi := &file_comment_methods_get_proto_msgTypes[3]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *GetTopCommentsRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetTopCommentsRequest) ProtoMessage() {}

func (x *GetTopCommentsRequest) ProtoReflect() protoreflect.Message {
	mi := &file_comment_methods_get_proto_msgTypes[3]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetTopCommentsRequest.ProtoReflect.Descriptor instead.
func (*GetTopCommentsRequest) Descriptor() ([]byte, []int) {
	return file_comment_methods_get_proto_rawDescGZIP(), []int{3}
}

func (x *GetTopCommentsRequest) GetCreationId() int64 {
	if x != nil {
		return x.CreationId
	}
	return 0
}

func (x *GetTopCommentsRequest) GetPage() int32 {
	if x != nil {
		return x.Page
	}
	return 0
}

type GetTopCommentsResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Comments []*TopComment       `protobuf:"bytes,1,rep,name=comments,proto3" json:"comments,omitempty"`
	Msg      *common.ApiResponse `protobuf:"bytes,2,opt,name=msg,proto3" json:"msg,omitempty"`
}

func (x *GetTopCommentsResponse) Reset() {
	*x = GetTopCommentsResponse{}
	mi := &file_comment_methods_get_proto_msgTypes[4]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *GetTopCommentsResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetTopCommentsResponse) ProtoMessage() {}

func (x *GetTopCommentsResponse) ProtoReflect() protoreflect.Message {
	mi := &file_comment_methods_get_proto_msgTypes[4]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetTopCommentsResponse.ProtoReflect.Descriptor instead.
func (*GetTopCommentsResponse) Descriptor() ([]byte, []int) {
	return file_comment_methods_get_proto_rawDescGZIP(), []int{4}
}

func (x *GetTopCommentsResponse) GetComments() []*TopComment {
	if x != nil {
		return x.Comments
	}
	return nil
}

func (x *GetTopCommentsResponse) GetMsg() *common.ApiResponse {
	if x != nil {
		return x.Msg
	}
	return nil
}

// 二级评论
type GetSecondCommentsRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	CreationId int64 `protobuf:"varint,1,opt,name=creation_id,json=creationId,proto3" json:"creation_id,omitempty"`
	Root       int32 `protobuf:"varint,2,opt,name=root,proto3" json:"root,omitempty"` // 一级评论所在
	Page       int32 `protobuf:"varint,3,opt,name=page,proto3" json:"page,omitempty"`
}

func (x *GetSecondCommentsRequest) Reset() {
	*x = GetSecondCommentsRequest{}
	mi := &file_comment_methods_get_proto_msgTypes[5]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *GetSecondCommentsRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetSecondCommentsRequest) ProtoMessage() {}

func (x *GetSecondCommentsRequest) ProtoReflect() protoreflect.Message {
	mi := &file_comment_methods_get_proto_msgTypes[5]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetSecondCommentsRequest.ProtoReflect.Descriptor instead.
func (*GetSecondCommentsRequest) Descriptor() ([]byte, []int) {
	return file_comment_methods_get_proto_rawDescGZIP(), []int{5}
}

func (x *GetSecondCommentsRequest) GetCreationId() int64 {
	if x != nil {
		return x.CreationId
	}
	return 0
}

func (x *GetSecondCommentsRequest) GetRoot() int32 {
	if x != nil {
		return x.Root
	}
	return 0
}

func (x *GetSecondCommentsRequest) GetPage() int32 {
	if x != nil {
		return x.Page
	}
	return 0
}

type GetSecondCommentsResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Comments []*SecondComment    `protobuf:"bytes,1,rep,name=comments,proto3" json:"comments,omitempty"`
	Msg      *common.ApiResponse `protobuf:"bytes,2,opt,name=msg,proto3" json:"msg,omitempty"`
}

func (x *GetSecondCommentsResponse) Reset() {
	*x = GetSecondCommentsResponse{}
	mi := &file_comment_methods_get_proto_msgTypes[6]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *GetSecondCommentsResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetSecondCommentsResponse) ProtoMessage() {}

func (x *GetSecondCommentsResponse) ProtoReflect() protoreflect.Message {
	mi := &file_comment_methods_get_proto_msgTypes[6]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetSecondCommentsResponse.ProtoReflect.Descriptor instead.
func (*GetSecondCommentsResponse) Descriptor() ([]byte, []int) {
	return file_comment_methods_get_proto_rawDescGZIP(), []int{6}
}

func (x *GetSecondCommentsResponse) GetComments() []*SecondComment {
	if x != nil {
		return x.Comments
	}
	return nil
}

func (x *GetSecondCommentsResponse) GetMsg() *common.ApiResponse {
	if x != nil {
		return x.Msg
	}
	return nil
}

// 消息中心，没做
// 回复我的评论，在页面的消息内显示,权限类
type GetReplyCommentsRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	AccessToken *common.AccessToken `protobuf:"bytes,1,opt,name=access_token,json=accessToken,proto3" json:"access_token,omitempty"` // 自己的id，
	Page        int32               `protobuf:"varint,2,opt,name=page,proto3" json:"page,omitempty"`
}

func (x *GetReplyCommentsRequest) Reset() {
	*x = GetReplyCommentsRequest{}
	mi := &file_comment_methods_get_proto_msgTypes[7]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *GetReplyCommentsRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetReplyCommentsRequest) ProtoMessage() {}

func (x *GetReplyCommentsRequest) ProtoReflect() protoreflect.Message {
	mi := &file_comment_methods_get_proto_msgTypes[7]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetReplyCommentsRequest.ProtoReflect.Descriptor instead.
func (*GetReplyCommentsRequest) Descriptor() ([]byte, []int) {
	return file_comment_methods_get_proto_rawDescGZIP(), []int{7}
}

func (x *GetReplyCommentsRequest) GetAccessToken() *common.AccessToken {
	if x != nil {
		return x.AccessToken
	}
	return nil
}

func (x *GetReplyCommentsRequest) GetPage() int32 {
	if x != nil {
		return x.Page
	}
	return 0
}

type GetCommentsResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Comments []*Comment          `protobuf:"bytes,1,rep,name=comments,proto3" json:"comments,omitempty"`
	Msg      *common.ApiResponse `protobuf:"bytes,2,opt,name=msg,proto3" json:"msg,omitempty"`
}

func (x *GetCommentsResponse) Reset() {
	*x = GetCommentsResponse{}
	mi := &file_comment_methods_get_proto_msgTypes[8]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *GetCommentsResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetCommentsResponse) ProtoMessage() {}

func (x *GetCommentsResponse) ProtoReflect() protoreflect.Message {
	mi := &file_comment_methods_get_proto_msgTypes[8]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetCommentsResponse.ProtoReflect.Descriptor instead.
func (*GetCommentsResponse) Descriptor() ([]byte, []int) {
	return file_comment_methods_get_proto_rawDescGZIP(), []int{8}
}

func (x *GetCommentsResponse) GetComments() []*Comment {
	if x != nil {
		return x.Comments
	}
	return nil
}

func (x *GetCommentsResponse) GetMsg() *common.ApiResponse {
	if x != nil {
		return x.Msg
	}
	return nil
}

type GetCommentRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id int32 `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
}

func (x *GetCommentRequest) Reset() {
	*x = GetCommentRequest{}
	mi := &file_comment_methods_get_proto_msgTypes[9]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *GetCommentRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetCommentRequest) ProtoMessage() {}

func (x *GetCommentRequest) ProtoReflect() protoreflect.Message {
	mi := &file_comment_methods_get_proto_msgTypes[9]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetCommentRequest.ProtoReflect.Descriptor instead.
func (*GetCommentRequest) Descriptor() ([]byte, []int) {
	return file_comment_methods_get_proto_rawDescGZIP(), []int{9}
}

func (x *GetCommentRequest) GetId() int32 {
	if x != nil {
		return x.Id
	}
	return 0
}

type GetCommentResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Comment *Comment            `protobuf:"bytes,1,opt,name=comment,proto3" json:"comment,omitempty"`
	Msg     *common.ApiResponse `protobuf:"bytes,2,opt,name=msg,proto3" json:"msg,omitempty"`
}

func (x *GetCommentResponse) Reset() {
	*x = GetCommentResponse{}
	mi := &file_comment_methods_get_proto_msgTypes[10]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *GetCommentResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetCommentResponse) ProtoMessage() {}

func (x *GetCommentResponse) ProtoReflect() protoreflect.Message {
	mi := &file_comment_methods_get_proto_msgTypes[10]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetCommentResponse.ProtoReflect.Descriptor instead.
func (*GetCommentResponse) Descriptor() ([]byte, []int) {
	return file_comment_methods_get_proto_rawDescGZIP(), []int{10}
}

func (x *GetCommentResponse) GetComment() *Comment {
	if x != nil {
		return x.Comment
	}
	return nil
}

func (x *GetCommentResponse) GetMsg() *common.ApiResponse {
	if x != nil {
		return x.Msg
	}
	return nil
}

var File_comment_methods_get_proto protoreflect.FileDescriptor

var file_comment_methods_get_proto_rawDesc = []byte{
	0x0a, 0x19, 0x63, 0x6f, 0x6d, 0x6d, 0x65, 0x6e, 0x74, 0x2f, 0x6d, 0x65, 0x74, 0x68, 0x6f, 0x64,
	0x73, 0x2f, 0x67, 0x65, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0f, 0x63, 0x6f, 0x6d,
	0x6d, 0x65, 0x6e, 0x74, 0x2e, 0x6d, 0x65, 0x74, 0x68, 0x6f, 0x64, 0x73, 0x1a, 0x1e, 0x63, 0x6f,
	0x6d, 0x6d, 0x65, 0x6e, 0x74, 0x2f, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x73, 0x2f, 0x63,
	0x6f, 0x6d, 0x6d, 0x65, 0x6e, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x19, 0x63, 0x6f,
	0x6d, 0x6d, 0x6f, 0x6e, 0x2f, 0x61, 0x63, 0x63, 0x65, 0x73, 0x73, 0x5f, 0x74, 0x6f, 0x6b, 0x65,
	0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x19, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2f,
	0x61, 0x70, 0x69, 0x5f, 0x72, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x22, 0x26, 0x0a, 0x12, 0x47, 0x65, 0x74, 0x43, 0x6f, 0x6d, 0x6d, 0x65, 0x6e, 0x74,
	0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x10, 0x0a, 0x03, 0x69, 0x64, 0x73, 0x18,
	0x01, 0x20, 0x03, 0x28, 0x05, 0x52, 0x03, 0x69, 0x64, 0x73, 0x22, 0x39, 0x0a, 0x16, 0x49, 0x6e,
	0x69, 0x74, 0x69, 0x61, 0x6c, 0x43, 0x6f, 0x6d, 0x6d, 0x65, 0x6e, 0x74, 0x73, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x12, 0x1f, 0x0a, 0x0b, 0x63, 0x72, 0x65, 0x61, 0x74, 0x69, 0x6f, 0x6e,
	0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x0a, 0x63, 0x72, 0x65, 0x61, 0x74,
	0x69, 0x6f, 0x6e, 0x49, 0x64, 0x22, 0xdb, 0x01, 0x0a, 0x17, 0x49, 0x6e, 0x69, 0x74, 0x69, 0x61,
	0x6c, 0x43, 0x6f, 0x6d, 0x6d, 0x65, 0x6e, 0x74, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73,
	0x65, 0x12, 0x38, 0x0a, 0x08, 0x63, 0x6f, 0x6d, 0x6d, 0x65, 0x6e, 0x74, 0x73, 0x18, 0x01, 0x20,
	0x03, 0x28, 0x0b, 0x32, 0x1c, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x65, 0x6e, 0x74, 0x2e, 0x6d, 0x65,
	0x73, 0x73, 0x61, 0x67, 0x65, 0x73, 0x2e, 0x54, 0x6f, 0x70, 0x43, 0x6f, 0x6d, 0x6d, 0x65, 0x6e,
	0x74, 0x52, 0x08, 0x63, 0x6f, 0x6d, 0x6d, 0x65, 0x6e, 0x74, 0x73, 0x12, 0x40, 0x0a, 0x0c, 0x63,
	0x6f, 0x6d, 0x6d, 0x65, 0x6e, 0x74, 0x5f, 0x61, 0x72, 0x65, 0x61, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x0b, 0x32, 0x1d, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x65, 0x6e, 0x74, 0x2e, 0x6d, 0x65, 0x73, 0x73,
	0x61, 0x67, 0x65, 0x73, 0x2e, 0x43, 0x6f, 0x6d, 0x6d, 0x65, 0x6e, 0x74, 0x41, 0x72, 0x65, 0x61,
	0x52, 0x0b, 0x63, 0x6f, 0x6d, 0x6d, 0x65, 0x6e, 0x74, 0x41, 0x72, 0x65, 0x61, 0x12, 0x1d, 0x0a,
	0x0a, 0x70, 0x61, 0x67, 0x65, 0x5f, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x18, 0x03, 0x20, 0x01, 0x28,
	0x05, 0x52, 0x09, 0x70, 0x61, 0x67, 0x65, 0x43, 0x6f, 0x75, 0x6e, 0x74, 0x12, 0x25, 0x0a, 0x03,
	0x6d, 0x73, 0x67, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x13, 0x2e, 0x63, 0x6f, 0x6d, 0x6d,
	0x6f, 0x6e, 0x2e, 0x41, 0x70, 0x69, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x52, 0x03,
	0x6d, 0x73, 0x67, 0x22, 0x4c, 0x0a, 0x15, 0x47, 0x65, 0x74, 0x54, 0x6f, 0x70, 0x43, 0x6f, 0x6d,
	0x6d, 0x65, 0x6e, 0x74, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x1f, 0x0a, 0x0b,
	0x63, 0x72, 0x65, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x03, 0x52, 0x0a, 0x63, 0x72, 0x65, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x49, 0x64, 0x12, 0x12, 0x0a,
	0x04, 0x70, 0x61, 0x67, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x05, 0x52, 0x04, 0x70, 0x61, 0x67,
	0x65, 0x22, 0x79, 0x0a, 0x16, 0x47, 0x65, 0x74, 0x54, 0x6f, 0x70, 0x43, 0x6f, 0x6d, 0x6d, 0x65,
	0x6e, 0x74, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x38, 0x0a, 0x08, 0x63,
	0x6f, 0x6d, 0x6d, 0x65, 0x6e, 0x74, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x1c, 0x2e,
	0x63, 0x6f, 0x6d, 0x6d, 0x65, 0x6e, 0x74, 0x2e, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x73,
	0x2e, 0x54, 0x6f, 0x70, 0x43, 0x6f, 0x6d, 0x6d, 0x65, 0x6e, 0x74, 0x52, 0x08, 0x63, 0x6f, 0x6d,
	0x6d, 0x65, 0x6e, 0x74, 0x73, 0x12, 0x25, 0x0a, 0x03, 0x6d, 0x73, 0x67, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x0b, 0x32, 0x13, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x41, 0x70, 0x69, 0x52,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x52, 0x03, 0x6d, 0x73, 0x67, 0x22, 0x63, 0x0a, 0x18,
	0x47, 0x65, 0x74, 0x53, 0x65, 0x63, 0x6f, 0x6e, 0x64, 0x43, 0x6f, 0x6d, 0x6d, 0x65, 0x6e, 0x74,
	0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x1f, 0x0a, 0x0b, 0x63, 0x72, 0x65, 0x61,
	0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x0a, 0x63,
	0x72, 0x65, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x49, 0x64, 0x12, 0x12, 0x0a, 0x04, 0x72, 0x6f, 0x6f,
	0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x05, 0x52, 0x04, 0x72, 0x6f, 0x6f, 0x74, 0x12, 0x12, 0x0a,
	0x04, 0x70, 0x61, 0x67, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x05, 0x52, 0x04, 0x70, 0x61, 0x67,
	0x65, 0x22, 0x7f, 0x0a, 0x19, 0x47, 0x65, 0x74, 0x53, 0x65, 0x63, 0x6f, 0x6e, 0x64, 0x43, 0x6f,
	0x6d, 0x6d, 0x65, 0x6e, 0x74, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x3b,
	0x0a, 0x08, 0x63, 0x6f, 0x6d, 0x6d, 0x65, 0x6e, 0x74, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b,
	0x32, 0x1f, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x65, 0x6e, 0x74, 0x2e, 0x6d, 0x65, 0x73, 0x73, 0x61,
	0x67, 0x65, 0x73, 0x2e, 0x53, 0x65, 0x63, 0x6f, 0x6e, 0x64, 0x43, 0x6f, 0x6d, 0x6d, 0x65, 0x6e,
	0x74, 0x52, 0x08, 0x63, 0x6f, 0x6d, 0x6d, 0x65, 0x6e, 0x74, 0x73, 0x12, 0x25, 0x0a, 0x03, 0x6d,
	0x73, 0x67, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x13, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f,
	0x6e, 0x2e, 0x41, 0x70, 0x69, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x52, 0x03, 0x6d,
	0x73, 0x67, 0x22, 0x65, 0x0a, 0x17, 0x47, 0x65, 0x74, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x43, 0x6f,
	0x6d, 0x6d, 0x65, 0x6e, 0x74, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x36, 0x0a,
	0x0c, 0x61, 0x63, 0x63, 0x65, 0x73, 0x73, 0x5f, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x0b, 0x32, 0x13, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x41, 0x63, 0x63,
	0x65, 0x73, 0x73, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x52, 0x0b, 0x61, 0x63, 0x63, 0x65, 0x73, 0x73,
	0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x12, 0x12, 0x0a, 0x04, 0x70, 0x61, 0x67, 0x65, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x05, 0x52, 0x04, 0x70, 0x61, 0x67, 0x65, 0x22, 0x73, 0x0a, 0x13, 0x47, 0x65, 0x74,
	0x43, 0x6f, 0x6d, 0x6d, 0x65, 0x6e, 0x74, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65,
	0x12, 0x35, 0x0a, 0x08, 0x63, 0x6f, 0x6d, 0x6d, 0x65, 0x6e, 0x74, 0x73, 0x18, 0x01, 0x20, 0x03,
	0x28, 0x0b, 0x32, 0x19, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x65, 0x6e, 0x74, 0x2e, 0x6d, 0x65, 0x73,
	0x73, 0x61, 0x67, 0x65, 0x73, 0x2e, 0x43, 0x6f, 0x6d, 0x6d, 0x65, 0x6e, 0x74, 0x52, 0x08, 0x63,
	0x6f, 0x6d, 0x6d, 0x65, 0x6e, 0x74, 0x73, 0x12, 0x25, 0x0a, 0x03, 0x6d, 0x73, 0x67, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x0b, 0x32, 0x13, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x41, 0x70,
	0x69, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x52, 0x03, 0x6d, 0x73, 0x67, 0x22, 0x23,
	0x0a, 0x11, 0x47, 0x65, 0x74, 0x43, 0x6f, 0x6d, 0x6d, 0x65, 0x6e, 0x74, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52,
	0x02, 0x69, 0x64, 0x22, 0x70, 0x0a, 0x12, 0x47, 0x65, 0x74, 0x43, 0x6f, 0x6d, 0x6d, 0x65, 0x6e,
	0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x33, 0x0a, 0x07, 0x63, 0x6f, 0x6d,
	0x6d, 0x65, 0x6e, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x19, 0x2e, 0x63, 0x6f, 0x6d,
	0x6d, 0x65, 0x6e, 0x74, 0x2e, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x73, 0x2e, 0x43, 0x6f,
	0x6d, 0x6d, 0x65, 0x6e, 0x74, 0x52, 0x07, 0x63, 0x6f, 0x6d, 0x6d, 0x65, 0x6e, 0x74, 0x12, 0x25,
	0x0a, 0x03, 0x6d, 0x73, 0x67, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x13, 0x2e, 0x63, 0x6f,
	0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x41, 0x70, 0x69, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65,
	0x52, 0x03, 0x6d, 0x73, 0x67, 0x42, 0x38, 0x5a, 0x36, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e,
	0x63, 0x6f, 0x6d, 0x2f, 0x59, 0x75, 0x78, 0x37, 0x37, 0x59, 0x75, 0x78, 0x2f, 0x70, 0x6c, 0x61,
	0x74, 0x66, 0x6f, 0x72, 0x6d, 0x5f, 0x62, 0x61, 0x63, 0x6b, 0x65, 0x6e, 0x64, 0x2f, 0x67, 0x65,
	0x6e, 0x65, 0x72, 0x61, 0x74, 0x65, 0x64, 0x2f, 0x63, 0x6f, 0x6d, 0x6d, 0x65, 0x6e, 0x74, 0x62,
	0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_comment_methods_get_proto_rawDescOnce sync.Once
	file_comment_methods_get_proto_rawDescData = file_comment_methods_get_proto_rawDesc
)

func file_comment_methods_get_proto_rawDescGZIP() []byte {
	file_comment_methods_get_proto_rawDescOnce.Do(func() {
		file_comment_methods_get_proto_rawDescData = protoimpl.X.CompressGZIP(file_comment_methods_get_proto_rawDescData)
	})
	return file_comment_methods_get_proto_rawDescData
}

var file_comment_methods_get_proto_msgTypes = make([]protoimpl.MessageInfo, 11)
var file_comment_methods_get_proto_goTypes = []any{
	(*GetCommentsRequest)(nil),        // 0: comment.methods.GetCommentsRequest
	(*InitialCommentsRequest)(nil),    // 1: comment.methods.InitialCommentsRequest
	(*InitialCommentsResponse)(nil),   // 2: comment.methods.InitialCommentsResponse
	(*GetTopCommentsRequest)(nil),     // 3: comment.methods.GetTopCommentsRequest
	(*GetTopCommentsResponse)(nil),    // 4: comment.methods.GetTopCommentsResponse
	(*GetSecondCommentsRequest)(nil),  // 5: comment.methods.GetSecondCommentsRequest
	(*GetSecondCommentsResponse)(nil), // 6: comment.methods.GetSecondCommentsResponse
	(*GetReplyCommentsRequest)(nil),   // 7: comment.methods.GetReplyCommentsRequest
	(*GetCommentsResponse)(nil),       // 8: comment.methods.GetCommentsResponse
	(*GetCommentRequest)(nil),         // 9: comment.methods.GetCommentRequest
	(*GetCommentResponse)(nil),        // 10: comment.methods.GetCommentResponse
	(*TopComment)(nil),                // 11: comment.messages.TopComment
	(*CommentArea)(nil),               // 12: comment.messages.CommentArea
	(*common.ApiResponse)(nil),        // 13: common.ApiResponse
	(*SecondComment)(nil),             // 14: comment.messages.SecondComment
	(*common.AccessToken)(nil),        // 15: common.AccessToken
	(*Comment)(nil),                   // 16: comment.messages.Comment
}
var file_comment_methods_get_proto_depIdxs = []int32{
	11, // 0: comment.methods.InitialCommentsResponse.comments:type_name -> comment.messages.TopComment
	12, // 1: comment.methods.InitialCommentsResponse.comment_area:type_name -> comment.messages.CommentArea
	13, // 2: comment.methods.InitialCommentsResponse.msg:type_name -> common.ApiResponse
	11, // 3: comment.methods.GetTopCommentsResponse.comments:type_name -> comment.messages.TopComment
	13, // 4: comment.methods.GetTopCommentsResponse.msg:type_name -> common.ApiResponse
	14, // 5: comment.methods.GetSecondCommentsResponse.comments:type_name -> comment.messages.SecondComment
	13, // 6: comment.methods.GetSecondCommentsResponse.msg:type_name -> common.ApiResponse
	15, // 7: comment.methods.GetReplyCommentsRequest.access_token:type_name -> common.AccessToken
	16, // 8: comment.methods.GetCommentsResponse.comments:type_name -> comment.messages.Comment
	13, // 9: comment.methods.GetCommentsResponse.msg:type_name -> common.ApiResponse
	16, // 10: comment.methods.GetCommentResponse.comment:type_name -> comment.messages.Comment
	13, // 11: comment.methods.GetCommentResponse.msg:type_name -> common.ApiResponse
	12, // [12:12] is the sub-list for method output_type
	12, // [12:12] is the sub-list for method input_type
	12, // [12:12] is the sub-list for extension type_name
	12, // [12:12] is the sub-list for extension extendee
	0,  // [0:12] is the sub-list for field type_name
}

func init() { file_comment_methods_get_proto_init() }
func file_comment_methods_get_proto_init() {
	if File_comment_methods_get_proto != nil {
		return
	}
	file_comment_messages_comment_proto_init()
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_comment_methods_get_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   11,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_comment_methods_get_proto_goTypes,
		DependencyIndexes: file_comment_methods_get_proto_depIdxs,
		MessageInfos:      file_comment_methods_get_proto_msgTypes,
	}.Build()
	File_comment_methods_get_proto = out.File
	file_comment_methods_get_proto_rawDesc = nil
	file_comment_methods_get_proto_goTypes = nil
	file_comment_methods_get_proto_depIdxs = nil
}
