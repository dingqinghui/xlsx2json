// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.31.0
// 	protoc        v4.25.1
// source: staticstruct.proto

package staticdata

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

// TestTable
type TestTable struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ID      int32       `protobuf:"varint,1,opt,name=ID,proto3" json:"ID,omitempty"`                //ID
	Enum1   int32       `protobuf:"varint,2,opt,name=Enum1,proto3" json:"Enum1,omitempty"`          //枚举字段
	Struct1 *TestStruct `protobuf:"bytes,3,opt,name=Struct1,proto3" json:"Struct1,omitempty"`       //结构体字段
	Array1  []int32     `protobuf:"varint,4,rep,packed,name=Array1,proto3" json:"Array1,omitempty"` //数组字段
}

func (x *TestTable) Reset() {
	*x = TestTable{}
	if protoimpl.UnsafeEnabled {
		mi := &file_staticstruct_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *TestTable) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*TestTable) ProtoMessage() {}

func (x *TestTable) ProtoReflect() protoreflect.Message {
	mi := &file_staticstruct_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use TestTable.ProtoReflect.Descriptor instead.
func (*TestTable) Descriptor() ([]byte, []int) {
	return file_staticstruct_proto_rawDescGZIP(), []int{0}
}

func (x *TestTable) GetID() int32 {
	if x != nil {
		return x.ID
	}
	return 0
}

func (x *TestTable) GetEnum1() int32 {
	if x != nil {
		return x.Enum1
	}
	return 0
}

func (x *TestTable) GetStruct1() *TestStruct {
	if x != nil {
		return x.Struct1
	}
	return nil
}

func (x *TestTable) GetArray1() []int32 {
	if x != nil {
		return x.Array1
	}
	return nil
}

// TestStruct
type TestStruct struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Field1 int32 `protobuf:"varint,1,opt,name=Field1,proto3" json:"Field1,omitempty"` //Field1
	Field2 int32 `protobuf:"varint,2,opt,name=Field2,proto3" json:"Field2,omitempty"` //Field2
}

func (x *TestStruct) Reset() {
	*x = TestStruct{}
	if protoimpl.UnsafeEnabled {
		mi := &file_staticstruct_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *TestStruct) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*TestStruct) ProtoMessage() {}

func (x *TestStruct) ProtoReflect() protoreflect.Message {
	mi := &file_staticstruct_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use TestStruct.ProtoReflect.Descriptor instead.
func (*TestStruct) Descriptor() ([]byte, []int) {
	return file_staticstruct_proto_rawDescGZIP(), []int{1}
}

func (x *TestStruct) GetField1() int32 {
	if x != nil {
		return x.Field1
	}
	return 0
}

func (x *TestStruct) GetField2() int32 {
	if x != nil {
		return x.Field2
	}
	return 0
}

var File_staticstruct_proto protoreflect.FileDescriptor

var file_staticstruct_proto_rawDesc = []byte{
	0x0a, 0x12, 0x73, 0x74, 0x61, 0x74, 0x69, 0x63, 0x73, 0x74, 0x72, 0x75, 0x63, 0x74, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x12, 0x04, 0x47, 0x61, 0x6d, 0x65, 0x22, 0x75, 0x0a, 0x09, 0x54, 0x65,
	0x73, 0x74, 0x54, 0x61, 0x62, 0x6c, 0x65, 0x12, 0x0e, 0x0a, 0x02, 0x49, 0x44, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x05, 0x52, 0x02, 0x49, 0x44, 0x12, 0x14, 0x0a, 0x05, 0x45, 0x6e, 0x75, 0x6d, 0x31,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x05, 0x52, 0x05, 0x45, 0x6e, 0x75, 0x6d, 0x31, 0x12, 0x2a, 0x0a,
	0x07, 0x53, 0x74, 0x72, 0x75, 0x63, 0x74, 0x31, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x10,
	0x2e, 0x47, 0x61, 0x6d, 0x65, 0x2e, 0x54, 0x65, 0x73, 0x74, 0x53, 0x74, 0x72, 0x75, 0x63, 0x74,
	0x52, 0x07, 0x53, 0x74, 0x72, 0x75, 0x63, 0x74, 0x31, 0x12, 0x16, 0x0a, 0x06, 0x41, 0x72, 0x72,
	0x61, 0x79, 0x31, 0x18, 0x04, 0x20, 0x03, 0x28, 0x05, 0x52, 0x06, 0x41, 0x72, 0x72, 0x61, 0x79,
	0x31, 0x22, 0x3c, 0x0a, 0x0a, 0x54, 0x65, 0x73, 0x74, 0x53, 0x74, 0x72, 0x75, 0x63, 0x74, 0x12,
	0x16, 0x0a, 0x06, 0x46, 0x69, 0x65, 0x6c, 0x64, 0x31, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52,
	0x06, 0x46, 0x69, 0x65, 0x6c, 0x64, 0x31, 0x12, 0x16, 0x0a, 0x06, 0x46, 0x69, 0x65, 0x6c, 0x64,
	0x32, 0x18, 0x02, 0x20, 0x01, 0x28, 0x05, 0x52, 0x06, 0x46, 0x69, 0x65, 0x6c, 0x64, 0x32, 0x42,
	0x0e, 0x5a, 0x0c, 0x2e, 0x3b, 0x73, 0x74, 0x61, 0x74, 0x69, 0x63, 0x64, 0x61, 0x74, 0x61, 0x62,
	0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_staticstruct_proto_rawDescOnce sync.Once
	file_staticstruct_proto_rawDescData = file_staticstruct_proto_rawDesc
)

func file_staticstruct_proto_rawDescGZIP() []byte {
	file_staticstruct_proto_rawDescOnce.Do(func() {
		file_staticstruct_proto_rawDescData = protoimpl.X.CompressGZIP(file_staticstruct_proto_rawDescData)
	})
	return file_staticstruct_proto_rawDescData
}

var file_staticstruct_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_staticstruct_proto_goTypes = []interface{}{
	(*TestTable)(nil),  // 0: Game.TestTable
	(*TestStruct)(nil), // 1: Game.TestStruct
}
var file_staticstruct_proto_depIdxs = []int32{
	1, // 0: Game.TestTable.Struct1:type_name -> Game.TestStruct
	1, // [1:1] is the sub-list for method output_type
	1, // [1:1] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_staticstruct_proto_init() }
func file_staticstruct_proto_init() {
	if File_staticstruct_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_staticstruct_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*TestTable); i {
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
		file_staticstruct_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*TestStruct); i {
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
			RawDescriptor: file_staticstruct_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_staticstruct_proto_goTypes,
		DependencyIndexes: file_staticstruct_proto_depIdxs,
		MessageInfos:      file_staticstruct_proto_msgTypes,
	}.Build()
	File_staticstruct_proto = out.File
	file_staticstruct_proto_rawDesc = nil
	file_staticstruct_proto_goTypes = nil
	file_staticstruct_proto_depIdxs = nil
}
