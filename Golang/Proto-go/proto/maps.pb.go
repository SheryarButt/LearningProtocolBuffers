// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.26.0
// 	protoc        v3.21.1
// source: maps.proto

package proto

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

type IdWrapper struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id int32 `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
}

func (x *IdWrapper) Reset() {
	*x = IdWrapper{}
	if protoimpl.UnsafeEnabled {
		mi := &file_maps_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *IdWrapper) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*IdWrapper) ProtoMessage() {}

func (x *IdWrapper) ProtoReflect() protoreflect.Message {
	mi := &file_maps_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use IdWrapper.ProtoReflect.Descriptor instead.
func (*IdWrapper) Descriptor() ([]byte, []int) {
	return file_maps_proto_rawDescGZIP(), []int{0}
}

func (x *IdWrapper) GetId() int32 {
	if x != nil {
		return x.Id
	}
	return 0
}

type MapExample struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Ids map[string]*IdWrapper `protobuf:"bytes,1,rep,name=ids,proto3" json:"ids,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
}

func (x *MapExample) Reset() {
	*x = MapExample{}
	if protoimpl.UnsafeEnabled {
		mi := &file_maps_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *MapExample) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*MapExample) ProtoMessage() {}

func (x *MapExample) ProtoReflect() protoreflect.Message {
	mi := &file_maps_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use MapExample.ProtoReflect.Descriptor instead.
func (*MapExample) Descriptor() ([]byte, []int) {
	return file_maps_proto_rawDescGZIP(), []int{1}
}

func (x *MapExample) GetIds() map[string]*IdWrapper {
	if x != nil {
		return x.Ids
	}
	return nil
}

var File_maps_proto protoreflect.FileDescriptor

var file_maps_proto_rawDesc = []byte{
	0x0a, 0x0a, 0x6d, 0x61, 0x70, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0c, 0x65, 0x78,
	0x61, 0x6d, 0x70, 0x6c, 0x65, 0x2e, 0x6d, 0x61, 0x70, 0x73, 0x22, 0x1b, 0x0a, 0x09, 0x49, 0x64,
	0x57, 0x72, 0x61, 0x70, 0x70, 0x65, 0x72, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x05, 0x52, 0x02, 0x69, 0x64, 0x22, 0x92, 0x01, 0x0a, 0x0a, 0x4d, 0x61, 0x70, 0x45,
	0x78, 0x61, 0x6d, 0x70, 0x6c, 0x65, 0x12, 0x33, 0x0a, 0x03, 0x69, 0x64, 0x73, 0x18, 0x01, 0x20,
	0x03, 0x28, 0x0b, 0x32, 0x21, 0x2e, 0x65, 0x78, 0x61, 0x6d, 0x70, 0x6c, 0x65, 0x2e, 0x6d, 0x61,
	0x70, 0x73, 0x2e, 0x4d, 0x61, 0x70, 0x45, 0x78, 0x61, 0x6d, 0x70, 0x6c, 0x65, 0x2e, 0x49, 0x64,
	0x73, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x52, 0x03, 0x69, 0x64, 0x73, 0x1a, 0x4f, 0x0a, 0x08, 0x49,
	0x64, 0x73, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x12, 0x10, 0x0a, 0x03, 0x6b, 0x65, 0x79, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x6b, 0x65, 0x79, 0x12, 0x2d, 0x0a, 0x05, 0x76, 0x61, 0x6c,
	0x75, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x17, 0x2e, 0x65, 0x78, 0x61, 0x6d, 0x70,
	0x6c, 0x65, 0x2e, 0x6d, 0x61, 0x70, 0x73, 0x2e, 0x49, 0x64, 0x57, 0x72, 0x61, 0x70, 0x70, 0x65,
	0x72, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x3a, 0x02, 0x38, 0x01, 0x42, 0x10, 0x5a, 0x0e,
	0x50, 0x72, 0x6f, 0x74, 0x6f, 0x2d, 0x67, 0x6f, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x06,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_maps_proto_rawDescOnce sync.Once
	file_maps_proto_rawDescData = file_maps_proto_rawDesc
)

func file_maps_proto_rawDescGZIP() []byte {
	file_maps_proto_rawDescOnce.Do(func() {
		file_maps_proto_rawDescData = protoimpl.X.CompressGZIP(file_maps_proto_rawDescData)
	})
	return file_maps_proto_rawDescData
}

var file_maps_proto_msgTypes = make([]protoimpl.MessageInfo, 3)
var file_maps_proto_goTypes = []interface{}{
	(*IdWrapper)(nil),  // 0: example.maps.IdWrapper
	(*MapExample)(nil), // 1: example.maps.MapExample
	nil,                // 2: example.maps.MapExample.IdsEntry
}
var file_maps_proto_depIdxs = []int32{
	2, // 0: example.maps.MapExample.ids:type_name -> example.maps.MapExample.IdsEntry
	0, // 1: example.maps.MapExample.IdsEntry.value:type_name -> example.maps.IdWrapper
	2, // [2:2] is the sub-list for method output_type
	2, // [2:2] is the sub-list for method input_type
	2, // [2:2] is the sub-list for extension type_name
	2, // [2:2] is the sub-list for extension extendee
	0, // [0:2] is the sub-list for field type_name
}

func init() { file_maps_proto_init() }
func file_maps_proto_init() {
	if File_maps_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_maps_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*IdWrapper); i {
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
		file_maps_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*MapExample); i {
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
			RawDescriptor: file_maps_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   3,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_maps_proto_goTypes,
		DependencyIndexes: file_maps_proto_depIdxs,
		MessageInfos:      file_maps_proto_msgTypes,
	}.Build()
	File_maps_proto = out.File
	file_maps_proto_rawDesc = nil
	file_maps_proto_goTypes = nil
	file_maps_proto_depIdxs = nil
}
