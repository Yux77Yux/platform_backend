// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.35.1
// 	protoc        v5.28.3
// source: comment/methods/post.proto

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

type PublishCommentRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Comment     *Comment            `protobuf:"bytes,1,opt,name=comment,proto3" json:"comment,omitempty"`
	AccessToken *common.AccessToken `protobuf:"bytes,2,opt,name=access_token,json=accessToken,proto3" json:"access_token,omitempty"`
}

func (x *PublishCommentRequest) Reset() {
	*x = PublishCommentRequest{}
	mi := &file_comment_methods_post_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *PublishCommentRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PublishCommentRequest) ProtoMessage() {}

func (x *PublishCommentRequest) ProtoReflect() protoreflect.Message {
	mi := &file_comment_methods_post_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PublishCommentRequest.ProtoReflect.Descriptor instead.
func (*PublishCommentRequest) Descriptor() ([]byte, []int) {
	return file_comment_methods_post_proto_rawDescGZIP(), []int{0}
}

func (x *PublishCommentRequest) GetComment() *Comment {
	if x != nil {
		return x.Comment
	}
	return nil
}

func (x *PublishCommentRequest) GetAccessToken() *common.AccessToken {
	if x != nil {
		return x.AccessToken
	}
	return nil
}

type PublishCommentResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Msg *common.ApiResponse `protobuf:"bytes,1,opt,name=msg,proto3" json:"msg,omitempty"`
}

func (x *PublishCommentResponse) Reset() {
	*x = PublishCommentResponse{}
	mi := &file_comment_methods_post_proto_msgTypes[1]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *PublishCommentResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PublishCommentResponse) ProtoMessage() {}

func (x *PublishCommentResponse) ProtoReflect() protoreflect.Message {
	mi := &file_comment_methods_post_proto_msgTypes[1]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PublishCommentResponse.ProtoReflect.Descriptor instead.
func (*PublishCommentResponse) Descriptor() ([]byte, []int) {
	return file_comment_methods_post_proto_rawDescGZIP(), []int{1}
}

func (x *PublishCommentResponse) GetMsg() *common.ApiResponse {
	if x != nil {
		return x.Msg
	}
	return nil
}

var File_comment_methods_post_proto protoreflect.FileDescriptor

var file_comment_methods_post_proto_rawDesc = []byte{
	0x0a, 0x1a, 0x63, 0x6f, 0x6d, 0x6d, 0x65, 0x6e, 0x74, 0x2f, 0x6d, 0x65, 0x74, 0x68, 0x6f, 0x64,
	0x73, 0x2f, 0x70, 0x6f, 0x73, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0f, 0x63, 0x6f,
	0x6d, 0x6d, 0x65, 0x6e, 0x74, 0x2e, 0x6d, 0x65, 0x74, 0x68, 0x6f, 0x64, 0x73, 0x1a, 0x19, 0x63,
	0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2f, 0x61, 0x70, 0x69, 0x5f, 0x72, 0x65, 0x73, 0x70, 0x6f, 0x6e,
	0x73, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1e, 0x63, 0x6f, 0x6d, 0x6d, 0x65, 0x6e,
	0x74, 0x2f, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x73, 0x2f, 0x63, 0x6f, 0x6d, 0x6d, 0x65,
	0x6e, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x19, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e,
	0x2f, 0x61, 0x63, 0x63, 0x65, 0x73, 0x73, 0x5f, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x22, 0x84, 0x01, 0x0a, 0x15, 0x50, 0x75, 0x62, 0x6c, 0x69, 0x73, 0x68, 0x43,
	0x6f, 0x6d, 0x6d, 0x65, 0x6e, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x33, 0x0a,
	0x07, 0x63, 0x6f, 0x6d, 0x6d, 0x65, 0x6e, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x19,
	0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x65, 0x6e, 0x74, 0x2e, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65,
	0x73, 0x2e, 0x43, 0x6f, 0x6d, 0x6d, 0x65, 0x6e, 0x74, 0x52, 0x07, 0x63, 0x6f, 0x6d, 0x6d, 0x65,
	0x6e, 0x74, 0x12, 0x36, 0x0a, 0x0c, 0x61, 0x63, 0x63, 0x65, 0x73, 0x73, 0x5f, 0x74, 0x6f, 0x6b,
	0x65, 0x6e, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x13, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f,
	0x6e, 0x2e, 0x41, 0x63, 0x63, 0x65, 0x73, 0x73, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x52, 0x0b, 0x61,
	0x63, 0x63, 0x65, 0x73, 0x73, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x22, 0x3f, 0x0a, 0x16, 0x50, 0x75,
	0x62, 0x6c, 0x69, 0x73, 0x68, 0x43, 0x6f, 0x6d, 0x6d, 0x65, 0x6e, 0x74, 0x52, 0x65, 0x73, 0x70,
	0x6f, 0x6e, 0x73, 0x65, 0x12, 0x25, 0x0a, 0x03, 0x6d, 0x73, 0x67, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x0b, 0x32, 0x13, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x41, 0x70, 0x69, 0x52, 0x65,
	0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x52, 0x03, 0x6d, 0x73, 0x67, 0x42, 0x38, 0x5a, 0x36, 0x67,
	0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x59, 0x75, 0x78, 0x37, 0x37, 0x59,
	0x75, 0x78, 0x2f, 0x70, 0x6c, 0x61, 0x74, 0x66, 0x6f, 0x72, 0x6d, 0x5f, 0x62, 0x61, 0x63, 0x6b,
	0x65, 0x6e, 0x64, 0x2f, 0x67, 0x65, 0x6e, 0x65, 0x72, 0x61, 0x74, 0x65, 0x64, 0x2f, 0x63, 0x6f,
	0x6d, 0x6d, 0x65, 0x6e, 0x74, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_comment_methods_post_proto_rawDescOnce sync.Once
	file_comment_methods_post_proto_rawDescData = file_comment_methods_post_proto_rawDesc
)

func file_comment_methods_post_proto_rawDescGZIP() []byte {
	file_comment_methods_post_proto_rawDescOnce.Do(func() {
		file_comment_methods_post_proto_rawDescData = protoimpl.X.CompressGZIP(file_comment_methods_post_proto_rawDescData)
	})
	return file_comment_methods_post_proto_rawDescData
}

var file_comment_methods_post_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_comment_methods_post_proto_goTypes = []any{
	(*PublishCommentRequest)(nil),  // 0: comment.methods.PublishCommentRequest
	(*PublishCommentResponse)(nil), // 1: comment.methods.PublishCommentResponse
	(*Comment)(nil),                // 2: comment.messages.Comment
	(*common.AccessToken)(nil),     // 3: common.AccessToken
	(*common.ApiResponse)(nil),     // 4: common.ApiResponse
}
var file_comment_methods_post_proto_depIdxs = []int32{
	2, // 0: comment.methods.PublishCommentRequest.comment:type_name -> comment.messages.Comment
	3, // 1: comment.methods.PublishCommentRequest.access_token:type_name -> common.AccessToken
	4, // 2: comment.methods.PublishCommentResponse.msg:type_name -> common.ApiResponse
	3, // [3:3] is the sub-list for method output_type
	3, // [3:3] is the sub-list for method input_type
	3, // [3:3] is the sub-list for extension type_name
	3, // [3:3] is the sub-list for extension extendee
	0, // [0:3] is the sub-list for field type_name
}

func init() { file_comment_methods_post_proto_init() }
func file_comment_methods_post_proto_init() {
	if File_comment_methods_post_proto != nil {
		return
	}
	file_comment_messages_comment_proto_init()
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_comment_methods_post_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_comment_methods_post_proto_goTypes,
		DependencyIndexes: file_comment_methods_post_proto_depIdxs,
		MessageInfos:      file_comment_methods_post_proto_msgTypes,
	}.Build()
	File_comment_methods_post_proto = out.File
	file_comment_methods_post_proto_rawDesc = nil
	file_comment_methods_post_proto_goTypes = nil
	file_comment_methods_post_proto_depIdxs = nil
}
