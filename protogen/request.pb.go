// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.34.2
// 	protoc        v5.27.2
// source: proto/type/request.proto

package protogen

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

type Request struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Data string `protobuf:"bytes,1,opt,name=data,proto3" json:"data,omitempty"`
}

func (x *Request) Reset() {
	*x = Request{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_type_request_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Request) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Request) ProtoMessage() {}

func (x *Request) ProtoReflect() protoreflect.Message {
	mi := &file_proto_type_request_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Request.ProtoReflect.Descriptor instead.
func (*Request) Descriptor() ([]byte, []int) {
	return file_proto_type_request_proto_rawDescGZIP(), []int{0}
}

func (x *Request) GetData() string {
	if x != nil {
		return x.Data
	}
	return ""
}

type SumReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Data int32 `protobuf:"varint,1,opt,name=data,proto3" json:"data,omitempty"`
}

func (x *SumReq) Reset() {
	*x = SumReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_type_request_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SumReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SumReq) ProtoMessage() {}

func (x *SumReq) ProtoReflect() protoreflect.Message {
	mi := &file_proto_type_request_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SumReq.ProtoReflect.Descriptor instead.
func (*SumReq) Descriptor() ([]byte, []int) {
	return file_proto_type_request_proto_rawDescGZIP(), []int{1}
}

func (x *SumReq) GetData() int32 {
	if x != nil {
		return x.Data
	}
	return 0
}

var File_proto_type_request_proto protoreflect.FileDescriptor

var file_proto_type_request_proto_rawDesc = []byte{
	0x0a, 0x18, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x74, 0x79, 0x70, 0x65, 0x2f, 0x72, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0a, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x2e, 0x74, 0x79, 0x70, 0x65, 0x22, 0x1d, 0x0a, 0x07, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x12, 0x12, 0x0a, 0x04, 0x64, 0x61, 0x74, 0x61, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x04, 0x64, 0x61, 0x74, 0x61, 0x22, 0x1c, 0x0a, 0x06, 0x53, 0x75, 0x6d, 0x52, 0x65, 0x71, 0x12,
	0x12, 0x0a, 0x04, 0x64, 0x61, 0x74, 0x61, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x04, 0x64,
	0x61, 0x74, 0x61, 0x42, 0x31, 0x5a, 0x2f, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f,
	0x6d, 0x2f, 0x64, 0x77, 0x70, 0x72, 0x7a, 0x2f, 0x62, 0x65, 0x6c, 0x61, 0x6a, 0x61, 0x72, 0x2d,
	0x67, 0x72, 0x70, 0x63, 0x2d, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x67, 0x65, 0x6e, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_proto_type_request_proto_rawDescOnce sync.Once
	file_proto_type_request_proto_rawDescData = file_proto_type_request_proto_rawDesc
)

func file_proto_type_request_proto_rawDescGZIP() []byte {
	file_proto_type_request_proto_rawDescOnce.Do(func() {
		file_proto_type_request_proto_rawDescData = protoimpl.X.CompressGZIP(file_proto_type_request_proto_rawDescData)
	})
	return file_proto_type_request_proto_rawDescData
}

var file_proto_type_request_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_proto_type_request_proto_goTypes = []any{
	(*Request)(nil), // 0: proto.type.Request
	(*SumReq)(nil),  // 1: proto.type.SumReq
}
var file_proto_type_request_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_proto_type_request_proto_init() }
func file_proto_type_request_proto_init() {
	if File_proto_type_request_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_proto_type_request_proto_msgTypes[0].Exporter = func(v any, i int) any {
			switch v := v.(*Request); i {
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
		file_proto_type_request_proto_msgTypes[1].Exporter = func(v any, i int) any {
			switch v := v.(*SumReq); i {
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
			RawDescriptor: file_proto_type_request_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_proto_type_request_proto_goTypes,
		DependencyIndexes: file_proto_type_request_proto_depIdxs,
		MessageInfos:      file_proto_type_request_proto_msgTypes,
	}.Build()
	File_proto_type_request_proto = out.File
	file_proto_type_request_proto_rawDesc = nil
	file_proto_type_request_proto_goTypes = nil
	file_proto_type_request_proto_depIdxs = nil
}
