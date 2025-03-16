// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.35.1
// 	protoc        v5.28.3
// source: event/review/review.proto

package review

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

type Exchange int32

const (
	Exchange_EXCHANGE_COMMENT_REVIEW         Exchange = 0
	Exchange_EXCHANGE_USER_REVIEW            Exchange = 1
	Exchange_EXCHANGE_CREATION_REVIEW        Exchange = 2
	Exchange_EXCHANGE_NEW_REVIEW             Exchange = 3
	Exchange_EXCHANGE_UPDATE                 Exchange = 4
	Exchange_EXCHANGE_BATCH_UPDATE           Exchange = 5
	Exchange_EXCHANGE_PEND_CREATION          Exchange = 6
	Exchange_EXCHANGE_UPDATE_USER_STATUS     Exchange = 7
	Exchange_EXCHANGE_UPDATE_CREATION_STATUS Exchange = 8
	Exchange_EXCHANGE_DELETE_CREATION        Exchange = 9
	Exchange_EXCHANGE_DELETE_COMMENT         Exchange = 10
)

// Enum value maps for Exchange.
var (
	Exchange_name = map[int32]string{
		0:  "EXCHANGE_COMMENT_REVIEW",
		1:  "EXCHANGE_USER_REVIEW",
		2:  "EXCHANGE_CREATION_REVIEW",
		3:  "EXCHANGE_NEW_REVIEW",
		4:  "EXCHANGE_UPDATE",
		5:  "EXCHANGE_BATCH_UPDATE",
		6:  "EXCHANGE_PEND_CREATION",
		7:  "EXCHANGE_UPDATE_USER_STATUS",
		8:  "EXCHANGE_UPDATE_CREATION_STATUS",
		9:  "EXCHANGE_DELETE_CREATION",
		10: "EXCHANGE_DELETE_COMMENT",
	}
	Exchange_value = map[string]int32{
		"EXCHANGE_COMMENT_REVIEW":         0,
		"EXCHANGE_USER_REVIEW":            1,
		"EXCHANGE_CREATION_REVIEW":        2,
		"EXCHANGE_NEW_REVIEW":             3,
		"EXCHANGE_UPDATE":                 4,
		"EXCHANGE_BATCH_UPDATE":           5,
		"EXCHANGE_PEND_CREATION":          6,
		"EXCHANGE_UPDATE_USER_STATUS":     7,
		"EXCHANGE_UPDATE_CREATION_STATUS": 8,
		"EXCHANGE_DELETE_CREATION":        9,
		"EXCHANGE_DELETE_COMMENT":         10,
	}
)

func (x Exchange) Enum() *Exchange {
	p := new(Exchange)
	*p = x
	return p
}

func (x Exchange) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (Exchange) Descriptor() protoreflect.EnumDescriptor {
	return file_event_review_review_proto_enumTypes[0].Descriptor()
}

func (Exchange) Type() protoreflect.EnumType {
	return &file_event_review_review_proto_enumTypes[0]
}

func (x Exchange) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use Exchange.Descriptor instead.
func (Exchange) EnumDescriptor() ([]byte, []int) {
	return file_event_review_review_proto_rawDescGZIP(), []int{0}
}

type Queue int32

const (
	Queue_QUEUE_COMMENT_REVIEW  Queue = 0
	Queue_QUEUE_USER_REVIEW     Queue = 1
	Queue_QUEUE_CREATION_REVIEW Queue = 2
	Queue_QUEUE_NEW_REVIEW      Queue = 3
	Queue_QUEUE_UPDATE          Queue = 4
	Queue_QUEUE_BATCH_UPDATE    Queue = 5
	Queue_QUEUE_PEND_CREATION   Queue = 6
)

// Enum value maps for Queue.
var (
	Queue_name = map[int32]string{
		0: "QUEUE_COMMENT_REVIEW",
		1: "QUEUE_USER_REVIEW",
		2: "QUEUE_CREATION_REVIEW",
		3: "QUEUE_NEW_REVIEW",
		4: "QUEUE_UPDATE",
		5: "QUEUE_BATCH_UPDATE",
		6: "QUEUE_PEND_CREATION",
	}
	Queue_value = map[string]int32{
		"QUEUE_COMMENT_REVIEW":  0,
		"QUEUE_USER_REVIEW":     1,
		"QUEUE_CREATION_REVIEW": 2,
		"QUEUE_NEW_REVIEW":      3,
		"QUEUE_UPDATE":          4,
		"QUEUE_BATCH_UPDATE":    5,
		"QUEUE_PEND_CREATION":   6,
	}
)

func (x Queue) Enum() *Queue {
	p := new(Queue)
	*p = x
	return p
}

func (x Queue) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (Queue) Descriptor() protoreflect.EnumDescriptor {
	return file_event_review_review_proto_enumTypes[1].Descriptor()
}

func (Queue) Type() protoreflect.EnumType {
	return &file_event_review_review_proto_enumTypes[1]
}

func (x Queue) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use Queue.Descriptor instead.
func (Queue) EnumDescriptor() ([]byte, []int) {
	return file_event_review_review_proto_rawDescGZIP(), []int{1}
}

type RoutingKey int32

const (
	RoutingKey_KEY_COMMENT_REVIEW         RoutingKey = 0
	RoutingKey_KEY_USER_REVIEW            RoutingKey = 1
	RoutingKey_KEY_CREATION_REVIEW        RoutingKey = 2
	RoutingKey_KEY_NEW_REVIEW             RoutingKey = 3
	RoutingKey_KEY_UPDATE                 RoutingKey = 4
	RoutingKey_KEY_BATCH_UPDATE           RoutingKey = 5
	RoutingKey_KEY_PEND_CREATION          RoutingKey = 6
	RoutingKey_KEY_UPDATE_USER_STATUS     RoutingKey = 7
	RoutingKey_KEY_UPDATE_CREATION_STATUS RoutingKey = 8
	RoutingKey_KEY_DELETE_CREATION        RoutingKey = 9
	RoutingKey_KEY_DELETE_COMMENT         RoutingKey = 10
)

// Enum value maps for RoutingKey.
var (
	RoutingKey_name = map[int32]string{
		0:  "KEY_COMMENT_REVIEW",
		1:  "KEY_USER_REVIEW",
		2:  "KEY_CREATION_REVIEW",
		3:  "KEY_NEW_REVIEW",
		4:  "KEY_UPDATE",
		5:  "KEY_BATCH_UPDATE",
		6:  "KEY_PEND_CREATION",
		7:  "KEY_UPDATE_USER_STATUS",
		8:  "KEY_UPDATE_CREATION_STATUS",
		9:  "KEY_DELETE_CREATION",
		10: "KEY_DELETE_COMMENT",
	}
	RoutingKey_value = map[string]int32{
		"KEY_COMMENT_REVIEW":         0,
		"KEY_USER_REVIEW":            1,
		"KEY_CREATION_REVIEW":        2,
		"KEY_NEW_REVIEW":             3,
		"KEY_UPDATE":                 4,
		"KEY_BATCH_UPDATE":           5,
		"KEY_PEND_CREATION":          6,
		"KEY_UPDATE_USER_STATUS":     7,
		"KEY_UPDATE_CREATION_STATUS": 8,
		"KEY_DELETE_CREATION":        9,
		"KEY_DELETE_COMMENT":         10,
	}
)

func (x RoutingKey) Enum() *RoutingKey {
	p := new(RoutingKey)
	*p = x
	return p
}

func (x RoutingKey) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (RoutingKey) Descriptor() protoreflect.EnumDescriptor {
	return file_event_review_review_proto_enumTypes[2].Descriptor()
}

func (RoutingKey) Type() protoreflect.EnumType {
	return &file_event_review_review_proto_enumTypes[2]
}

func (x RoutingKey) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use RoutingKey.Descriptor instead.
func (RoutingKey) EnumDescriptor() ([]byte, []int) {
	return file_event_review_review_proto_rawDescGZIP(), []int{2}
}

var File_event_review_review_proto protoreflect.FileDescriptor

var file_event_review_review_proto_rawDesc = []byte{
	0x0a, 0x19, 0x65, 0x76, 0x65, 0x6e, 0x74, 0x2f, 0x72, 0x65, 0x76, 0x69, 0x65, 0x77, 0x2f, 0x72,
	0x65, 0x76, 0x69, 0x65, 0x77, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0c, 0x65, 0x76, 0x65,
	0x6e, 0x74, 0x2e, 0x72, 0x65, 0x76, 0x69, 0x65, 0x77, 0x2a, 0xc5, 0x02, 0x0a, 0x08, 0x45, 0x78,
	0x63, 0x68, 0x61, 0x6e, 0x67, 0x65, 0x12, 0x1b, 0x0a, 0x17, 0x45, 0x58, 0x43, 0x48, 0x41, 0x4e,
	0x47, 0x45, 0x5f, 0x43, 0x4f, 0x4d, 0x4d, 0x45, 0x4e, 0x54, 0x5f, 0x52, 0x45, 0x56, 0x49, 0x45,
	0x57, 0x10, 0x00, 0x12, 0x18, 0x0a, 0x14, 0x45, 0x58, 0x43, 0x48, 0x41, 0x4e, 0x47, 0x45, 0x5f,
	0x55, 0x53, 0x45, 0x52, 0x5f, 0x52, 0x45, 0x56, 0x49, 0x45, 0x57, 0x10, 0x01, 0x12, 0x1c, 0x0a,
	0x18, 0x45, 0x58, 0x43, 0x48, 0x41, 0x4e, 0x47, 0x45, 0x5f, 0x43, 0x52, 0x45, 0x41, 0x54, 0x49,
	0x4f, 0x4e, 0x5f, 0x52, 0x45, 0x56, 0x49, 0x45, 0x57, 0x10, 0x02, 0x12, 0x17, 0x0a, 0x13, 0x45,
	0x58, 0x43, 0x48, 0x41, 0x4e, 0x47, 0x45, 0x5f, 0x4e, 0x45, 0x57, 0x5f, 0x52, 0x45, 0x56, 0x49,
	0x45, 0x57, 0x10, 0x03, 0x12, 0x13, 0x0a, 0x0f, 0x45, 0x58, 0x43, 0x48, 0x41, 0x4e, 0x47, 0x45,
	0x5f, 0x55, 0x50, 0x44, 0x41, 0x54, 0x45, 0x10, 0x04, 0x12, 0x19, 0x0a, 0x15, 0x45, 0x58, 0x43,
	0x48, 0x41, 0x4e, 0x47, 0x45, 0x5f, 0x42, 0x41, 0x54, 0x43, 0x48, 0x5f, 0x55, 0x50, 0x44, 0x41,
	0x54, 0x45, 0x10, 0x05, 0x12, 0x1a, 0x0a, 0x16, 0x45, 0x58, 0x43, 0x48, 0x41, 0x4e, 0x47, 0x45,
	0x5f, 0x50, 0x45, 0x4e, 0x44, 0x5f, 0x43, 0x52, 0x45, 0x41, 0x54, 0x49, 0x4f, 0x4e, 0x10, 0x06,
	0x12, 0x1f, 0x0a, 0x1b, 0x45, 0x58, 0x43, 0x48, 0x41, 0x4e, 0x47, 0x45, 0x5f, 0x55, 0x50, 0x44,
	0x41, 0x54, 0x45, 0x5f, 0x55, 0x53, 0x45, 0x52, 0x5f, 0x53, 0x54, 0x41, 0x54, 0x55, 0x53, 0x10,
	0x07, 0x12, 0x23, 0x0a, 0x1f, 0x45, 0x58, 0x43, 0x48, 0x41, 0x4e, 0x47, 0x45, 0x5f, 0x55, 0x50,
	0x44, 0x41, 0x54, 0x45, 0x5f, 0x43, 0x52, 0x45, 0x41, 0x54, 0x49, 0x4f, 0x4e, 0x5f, 0x53, 0x54,
	0x41, 0x54, 0x55, 0x53, 0x10, 0x08, 0x12, 0x1c, 0x0a, 0x18, 0x45, 0x58, 0x43, 0x48, 0x41, 0x4e,
	0x47, 0x45, 0x5f, 0x44, 0x45, 0x4c, 0x45, 0x54, 0x45, 0x5f, 0x43, 0x52, 0x45, 0x41, 0x54, 0x49,
	0x4f, 0x4e, 0x10, 0x09, 0x12, 0x1b, 0x0a, 0x17, 0x45, 0x58, 0x43, 0x48, 0x41, 0x4e, 0x47, 0x45,
	0x5f, 0x44, 0x45, 0x4c, 0x45, 0x54, 0x45, 0x5f, 0x43, 0x4f, 0x4d, 0x4d, 0x45, 0x4e, 0x54, 0x10,
	0x0a, 0x2a, 0xac, 0x01, 0x0a, 0x05, 0x51, 0x75, 0x65, 0x75, 0x65, 0x12, 0x18, 0x0a, 0x14, 0x51,
	0x55, 0x45, 0x55, 0x45, 0x5f, 0x43, 0x4f, 0x4d, 0x4d, 0x45, 0x4e, 0x54, 0x5f, 0x52, 0x45, 0x56,
	0x49, 0x45, 0x57, 0x10, 0x00, 0x12, 0x15, 0x0a, 0x11, 0x51, 0x55, 0x45, 0x55, 0x45, 0x5f, 0x55,
	0x53, 0x45, 0x52, 0x5f, 0x52, 0x45, 0x56, 0x49, 0x45, 0x57, 0x10, 0x01, 0x12, 0x19, 0x0a, 0x15,
	0x51, 0x55, 0x45, 0x55, 0x45, 0x5f, 0x43, 0x52, 0x45, 0x41, 0x54, 0x49, 0x4f, 0x4e, 0x5f, 0x52,
	0x45, 0x56, 0x49, 0x45, 0x57, 0x10, 0x02, 0x12, 0x14, 0x0a, 0x10, 0x51, 0x55, 0x45, 0x55, 0x45,
	0x5f, 0x4e, 0x45, 0x57, 0x5f, 0x52, 0x45, 0x56, 0x49, 0x45, 0x57, 0x10, 0x03, 0x12, 0x10, 0x0a,
	0x0c, 0x51, 0x55, 0x45, 0x55, 0x45, 0x5f, 0x55, 0x50, 0x44, 0x41, 0x54, 0x45, 0x10, 0x04, 0x12,
	0x16, 0x0a, 0x12, 0x51, 0x55, 0x45, 0x55, 0x45, 0x5f, 0x42, 0x41, 0x54, 0x43, 0x48, 0x5f, 0x55,
	0x50, 0x44, 0x41, 0x54, 0x45, 0x10, 0x05, 0x12, 0x17, 0x0a, 0x13, 0x51, 0x55, 0x45, 0x55, 0x45,
	0x5f, 0x50, 0x45, 0x4e, 0x44, 0x5f, 0x43, 0x52, 0x45, 0x41, 0x54, 0x49, 0x4f, 0x4e, 0x10, 0x06,
	0x2a, 0x90, 0x02, 0x0a, 0x0a, 0x52, 0x6f, 0x75, 0x74, 0x69, 0x6e, 0x67, 0x4b, 0x65, 0x79, 0x12,
	0x16, 0x0a, 0x12, 0x4b, 0x45, 0x59, 0x5f, 0x43, 0x4f, 0x4d, 0x4d, 0x45, 0x4e, 0x54, 0x5f, 0x52,
	0x45, 0x56, 0x49, 0x45, 0x57, 0x10, 0x00, 0x12, 0x13, 0x0a, 0x0f, 0x4b, 0x45, 0x59, 0x5f, 0x55,
	0x53, 0x45, 0x52, 0x5f, 0x52, 0x45, 0x56, 0x49, 0x45, 0x57, 0x10, 0x01, 0x12, 0x17, 0x0a, 0x13,
	0x4b, 0x45, 0x59, 0x5f, 0x43, 0x52, 0x45, 0x41, 0x54, 0x49, 0x4f, 0x4e, 0x5f, 0x52, 0x45, 0x56,
	0x49, 0x45, 0x57, 0x10, 0x02, 0x12, 0x12, 0x0a, 0x0e, 0x4b, 0x45, 0x59, 0x5f, 0x4e, 0x45, 0x57,
	0x5f, 0x52, 0x45, 0x56, 0x49, 0x45, 0x57, 0x10, 0x03, 0x12, 0x0e, 0x0a, 0x0a, 0x4b, 0x45, 0x59,
	0x5f, 0x55, 0x50, 0x44, 0x41, 0x54, 0x45, 0x10, 0x04, 0x12, 0x14, 0x0a, 0x10, 0x4b, 0x45, 0x59,
	0x5f, 0x42, 0x41, 0x54, 0x43, 0x48, 0x5f, 0x55, 0x50, 0x44, 0x41, 0x54, 0x45, 0x10, 0x05, 0x12,
	0x15, 0x0a, 0x11, 0x4b, 0x45, 0x59, 0x5f, 0x50, 0x45, 0x4e, 0x44, 0x5f, 0x43, 0x52, 0x45, 0x41,
	0x54, 0x49, 0x4f, 0x4e, 0x10, 0x06, 0x12, 0x1a, 0x0a, 0x16, 0x4b, 0x45, 0x59, 0x5f, 0x55, 0x50,
	0x44, 0x41, 0x54, 0x45, 0x5f, 0x55, 0x53, 0x45, 0x52, 0x5f, 0x53, 0x54, 0x41, 0x54, 0x55, 0x53,
	0x10, 0x07, 0x12, 0x1e, 0x0a, 0x1a, 0x4b, 0x45, 0x59, 0x5f, 0x55, 0x50, 0x44, 0x41, 0x54, 0x45,
	0x5f, 0x43, 0x52, 0x45, 0x41, 0x54, 0x49, 0x4f, 0x4e, 0x5f, 0x53, 0x54, 0x41, 0x54, 0x55, 0x53,
	0x10, 0x08, 0x12, 0x17, 0x0a, 0x13, 0x4b, 0x45, 0x59, 0x5f, 0x44, 0x45, 0x4c, 0x45, 0x54, 0x45,
	0x5f, 0x43, 0x52, 0x45, 0x41, 0x54, 0x49, 0x4f, 0x4e, 0x10, 0x09, 0x12, 0x16, 0x0a, 0x12, 0x4b,
	0x45, 0x59, 0x5f, 0x44, 0x45, 0x4c, 0x45, 0x54, 0x45, 0x5f, 0x43, 0x4f, 0x4d, 0x4d, 0x45, 0x4e,
	0x54, 0x10, 0x0a, 0x42, 0x3e, 0x5a, 0x3c, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f,
	0x6d, 0x2f, 0x59, 0x75, 0x78, 0x37, 0x37, 0x59, 0x75, 0x78, 0x2f, 0x70, 0x6c, 0x61, 0x74, 0x66,
	0x6f, 0x72, 0x6d, 0x5f, 0x62, 0x61, 0x63, 0x6b, 0x65, 0x6e, 0x64, 0x2f, 0x67, 0x65, 0x6e, 0x65,
	0x72, 0x61, 0x74, 0x65, 0x64, 0x2f, 0x65, 0x76, 0x65, 0x6e, 0x74, 0x2f, 0x72, 0x65, 0x76, 0x69,
	0x65, 0x77, 0x3b, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_event_review_review_proto_rawDescOnce sync.Once
	file_event_review_review_proto_rawDescData = file_event_review_review_proto_rawDesc
)

func file_event_review_review_proto_rawDescGZIP() []byte {
	file_event_review_review_proto_rawDescOnce.Do(func() {
		file_event_review_review_proto_rawDescData = protoimpl.X.CompressGZIP(file_event_review_review_proto_rawDescData)
	})
	return file_event_review_review_proto_rawDescData
}

var file_event_review_review_proto_enumTypes = make([]protoimpl.EnumInfo, 3)
var file_event_review_review_proto_goTypes = []any{
	(Exchange)(0),   // 0: event.review.Exchange
	(Queue)(0),      // 1: event.review.Queue
	(RoutingKey)(0), // 2: event.review.RoutingKey
}
var file_event_review_review_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_event_review_review_proto_init() }
func file_event_review_review_proto_init() {
	if File_event_review_review_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_event_review_review_proto_rawDesc,
			NumEnums:      3,
			NumMessages:   0,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_event_review_review_proto_goTypes,
		DependencyIndexes: file_event_review_review_proto_depIdxs,
		EnumInfos:         file_event_review_review_proto_enumTypes,
	}.Build()
	File_event_review_review_proto = out.File
	file_event_review_review_proto_rawDesc = nil
	file_event_review_review_proto_goTypes = nil
	file_event_review_review_proto_depIdxs = nil
}
