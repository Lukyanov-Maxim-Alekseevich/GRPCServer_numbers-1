// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.32.0
// 	protoc        v4.25.1
// source: adder.proto

package api

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

type YesRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	FindNumber int64 `protobuf:"varint,1,opt,name=find_number,json=findNumber,proto3" json:"find_number,omitempty"`
}

func (x *YesRequest) Reset() {
	*x = YesRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_adder_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *YesRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*YesRequest) ProtoMessage() {}

func (x *YesRequest) ProtoReflect() protoreflect.Message {
	mi := &file_adder_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use YesRequest.ProtoReflect.Descriptor instead.
func (*YesRequest) Descriptor() ([]byte, []int) {
	return file_adder_proto_rawDescGZIP(), []int{0}
}

func (x *YesRequest) GetFindNumber() int64 {
	if x != nil {
		return x.FindNumber
	}
	return 0
}

type YesResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	AnswerFindNumber int64 `protobuf:"varint,1,opt,name=answer_find_number,json=answerFindNumber,proto3" json:"answer_find_number,omitempty"`
}

func (x *YesResponse) Reset() {
	*x = YesResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_adder_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *YesResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*YesResponse) ProtoMessage() {}

func (x *YesResponse) ProtoReflect() protoreflect.Message {
	mi := &file_adder_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use YesResponse.ProtoReflect.Descriptor instead.
func (*YesResponse) Descriptor() ([]byte, []int) {
	return file_adder_proto_rawDescGZIP(), []int{1}
}

func (x *YesResponse) GetAnswerFindNumber() int64 {
	if x != nil {
		return x.AnswerFindNumber
	}
	return 0
}

type NoRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	NoFindNumber int64 `protobuf:"varint,1,opt,name=no_find_number,json=noFindNumber,proto3" json:"no_find_number,omitempty"`
}

func (x *NoRequest) Reset() {
	*x = NoRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_adder_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *NoRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*NoRequest) ProtoMessage() {}

func (x *NoRequest) ProtoReflect() protoreflect.Message {
	mi := &file_adder_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use NoRequest.ProtoReflect.Descriptor instead.
func (*NoRequest) Descriptor() ([]byte, []int) {
	return file_adder_proto_rawDescGZIP(), []int{2}
}

func (x *NoRequest) GetNoFindNumber() int64 {
	if x != nil {
		return x.NoFindNumber
	}
	return 0
}

type NoResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	AnswerNoFindNumber int64 `protobuf:"varint,1,opt,name=answer_no_find_number,json=answerNoFindNumber,proto3" json:"answer_no_find_number,omitempty"`
}

func (x *NoResponse) Reset() {
	*x = NoResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_adder_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *NoResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*NoResponse) ProtoMessage() {}

func (x *NoResponse) ProtoReflect() protoreflect.Message {
	mi := &file_adder_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use NoResponse.ProtoReflect.Descriptor instead.
func (*NoResponse) Descriptor() ([]byte, []int) {
	return file_adder_proto_rawDescGZIP(), []int{3}
}

func (x *NoResponse) GetAnswerNoFindNumber() int64 {
	if x != nil {
		return x.AnswerNoFindNumber
	}
	return 0
}

var File_adder_proto protoreflect.FileDescriptor

var file_adder_proto_rawDesc = []byte{
	0x0a, 0x0b, 0x61, 0x64, 0x64, 0x65, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x03, 0x61,
	0x70, 0x69, 0x22, 0x2d, 0x0a, 0x0a, 0x59, 0x65, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x12, 0x1f, 0x0a, 0x0b, 0x66, 0x69, 0x6e, 0x64, 0x5f, 0x6e, 0x75, 0x6d, 0x62, 0x65, 0x72, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x0a, 0x66, 0x69, 0x6e, 0x64, 0x4e, 0x75, 0x6d, 0x62, 0x65,
	0x72, 0x22, 0x3b, 0x0a, 0x0b, 0x59, 0x65, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65,
	0x12, 0x2c, 0x0a, 0x12, 0x61, 0x6e, 0x73, 0x77, 0x65, 0x72, 0x5f, 0x66, 0x69, 0x6e, 0x64, 0x5f,
	0x6e, 0x75, 0x6d, 0x62, 0x65, 0x72, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x10, 0x61, 0x6e,
	0x73, 0x77, 0x65, 0x72, 0x46, 0x69, 0x6e, 0x64, 0x4e, 0x75, 0x6d, 0x62, 0x65, 0x72, 0x22, 0x31,
	0x0a, 0x09, 0x4e, 0x6f, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x24, 0x0a, 0x0e, 0x6e,
	0x6f, 0x5f, 0x66, 0x69, 0x6e, 0x64, 0x5f, 0x6e, 0x75, 0x6d, 0x62, 0x65, 0x72, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x03, 0x52, 0x0c, 0x6e, 0x6f, 0x46, 0x69, 0x6e, 0x64, 0x4e, 0x75, 0x6d, 0x62, 0x65,
	0x72, 0x22, 0x3f, 0x0a, 0x0a, 0x4e, 0x6f, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12,
	0x31, 0x0a, 0x15, 0x61, 0x6e, 0x73, 0x77, 0x65, 0x72, 0x5f, 0x6e, 0x6f, 0x5f, 0x66, 0x69, 0x6e,
	0x64, 0x5f, 0x6e, 0x75, 0x6d, 0x62, 0x65, 0x72, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x12,
	0x61, 0x6e, 0x73, 0x77, 0x65, 0x72, 0x4e, 0x6f, 0x46, 0x69, 0x6e, 0x64, 0x4e, 0x75, 0x6d, 0x62,
	0x65, 0x72, 0x32, 0x65, 0x0a, 0x06, 0x4e, 0x75, 0x6d, 0x62, 0x65, 0x72, 0x12, 0x2e, 0x0a, 0x09,
	0x59, 0x65, 0x73, 0x4e, 0x75, 0x6d, 0x62, 0x65, 0x72, 0x12, 0x0f, 0x2e, 0x61, 0x70, 0x69, 0x2e,
	0x59, 0x65, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x10, 0x2e, 0x61, 0x70, 0x69,
	0x2e, 0x59, 0x65, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x2b, 0x0a, 0x08,
	0x4e, 0x6f, 0x4e, 0x75, 0x6d, 0x62, 0x65, 0x72, 0x12, 0x0e, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x4e,
	0x6f, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x0f, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x4e,
	0x6f, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x42, 0x09, 0x5a, 0x07, 0x70, 0x6b, 0x67,
	0x2f, 0x61, 0x70, 0x69, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_adder_proto_rawDescOnce sync.Once
	file_adder_proto_rawDescData = file_adder_proto_rawDesc
)

func file_adder_proto_rawDescGZIP() []byte {
	file_adder_proto_rawDescOnce.Do(func() {
		file_adder_proto_rawDescData = protoimpl.X.CompressGZIP(file_adder_proto_rawDescData)
	})
	return file_adder_proto_rawDescData
}

var file_adder_proto_msgTypes = make([]protoimpl.MessageInfo, 4)
var file_adder_proto_goTypes = []interface{}{
	(*YesRequest)(nil),  // 0: api.YesRequest
	(*YesResponse)(nil), // 1: api.YesResponse
	(*NoRequest)(nil),   // 2: api.NoRequest
	(*NoResponse)(nil),  // 3: api.NoResponse
}
var file_adder_proto_depIdxs = []int32{
	0, // 0: api.Number.YesNumber:input_type -> api.YesRequest
	2, // 1: api.Number.NoNumber:input_type -> api.NoRequest
	1, // 2: api.Number.YesNumber:output_type -> api.YesResponse
	3, // 3: api.Number.NoNumber:output_type -> api.NoResponse
	2, // [2:4] is the sub-list for method output_type
	0, // [0:2] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_adder_proto_init() }
func file_adder_proto_init() {
	if File_adder_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_adder_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*YesRequest); i {
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
		file_adder_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*YesResponse); i {
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
		file_adder_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*NoRequest); i {
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
		file_adder_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*NoResponse); i {
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
			RawDescriptor: file_adder_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   4,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_adder_proto_goTypes,
		DependencyIndexes: file_adder_proto_depIdxs,
		MessageInfos:      file_adder_proto_msgTypes,
	}.Build()
	File_adder_proto = out.File
	file_adder_proto_rawDesc = nil
	file_adder_proto_goTypes = nil
	file_adder_proto_depIdxs = nil
}