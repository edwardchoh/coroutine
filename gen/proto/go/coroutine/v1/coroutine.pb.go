// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.31.0
// 	protoc        (unknown)
// source: coroutine/v1/coroutine.proto

package coroutinev1

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

// State is durable coroutine state.
type State struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Build is information about the build in which the coruotine state
	// was generated.
	Build *Build `protobuf:"bytes,1,opt,name=build,proto3" json:"build,omitempty"`
	// State is a serialized representation of the objects in the program
	// that are reachable from the coroutine stack.
	State []byte `protobuf:"bytes,2,opt,name=state,proto3" json:"state,omitempty"`
	// Types is the set of types used by the object graph.
	Types []*Type `protobuf:"bytes,3,rep,name=types,proto3" json:"types,omitempty"`
	// Functions is the set of functions, methods and closures referenced
	// by the object graph.
	Functions []*Function `protobuf:"bytes,4,rep,name=functions,proto3" json:"functions,omitempty"`
	// Regions are encoded regions of memory.
	Regions []*Region `protobuf:"bytes,5,rep,name=regions,proto3" json:"regions,omitempty"`
	// Root is the root object.
	Root *Region `protobuf:"bytes,6,opt,name=root,proto3" json:"root,omitempty"`
	// Strings is the string table.
	Strings []string `protobuf:"bytes,7,rep,name=strings,proto3" json:"strings,omitempty"`
}

func (x *State) Reset() {
	*x = State{}
	if protoimpl.UnsafeEnabled {
		mi := &file_coroutine_v1_coroutine_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *State) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*State) ProtoMessage() {}

func (x *State) ProtoReflect() protoreflect.Message {
	mi := &file_coroutine_v1_coroutine_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use State.ProtoReflect.Descriptor instead.
func (*State) Descriptor() ([]byte, []int) {
	return file_coroutine_v1_coroutine_proto_rawDescGZIP(), []int{0}
}

func (x *State) GetBuild() *Build {
	if x != nil {
		return x.Build
	}
	return nil
}

func (x *State) GetState() []byte {
	if x != nil {
		return x.State
	}
	return nil
}

func (x *State) GetTypes() []*Type {
	if x != nil {
		return x.Types
	}
	return nil
}

func (x *State) GetFunctions() []*Function {
	if x != nil {
		return x.Functions
	}
	return nil
}

func (x *State) GetRegions() []*Region {
	if x != nil {
		return x.Regions
	}
	return nil
}

func (x *State) GetRoot() *Region {
	if x != nil {
		return x.Root
	}
	return nil
}

func (x *State) GetStrings() []string {
	if x != nil {
		return x.Strings
	}
	return nil
}

// Build is info about the build in which a durable coroutine
// is/was running.
type Build struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id   string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Os   string `protobuf:"bytes,2,opt,name=os,proto3" json:"os,omitempty"`
	Arch string `protobuf:"bytes,3,opt,name=arch,proto3" json:"arch,omitempty"`
}

func (x *Build) Reset() {
	*x = Build{}
	if protoimpl.UnsafeEnabled {
		mi := &file_coroutine_v1_coroutine_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Build) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Build) ProtoMessage() {}

func (x *Build) ProtoReflect() protoreflect.Message {
	mi := &file_coroutine_v1_coroutine_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Build.ProtoReflect.Descriptor instead.
func (*Build) Descriptor() ([]byte, []int) {
	return file_coroutine_v1_coroutine_proto_rawDescGZIP(), []int{1}
}

func (x *Build) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *Build) GetOs() string {
	if x != nil {
		return x.Os
	}
	return ""
}

func (x *Build) GetArch() string {
	if x != nil {
		return x.Arch
	}
	return ""
}

// Region is an encoded region of memory.
type Region struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Type is the type of the region, shifted left by one.
	//
	// The least significant bit indicates that this region represents
	// an array, and that the type is the array element type rather
	// than the object that's encoded in this region.
	Type uint32 `protobuf:"varint,1,opt,name=type,proto3" json:"type,omitempty"`
	// Array length, for regions that are arrays.
	ArrayLength uint32 `protobuf:"varint,2,opt,name=array_length,json=arrayLength,proto3" json:"array_length,omitempty"`
	// Data is the encoded contents of the memory region.
	Data []byte `protobuf:"bytes,3,opt,name=data,proto3" json:"data,omitempty"`
}

func (x *Region) Reset() {
	*x = Region{}
	if protoimpl.UnsafeEnabled {
		mi := &file_coroutine_v1_coroutine_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Region) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Region) ProtoMessage() {}

func (x *Region) ProtoReflect() protoreflect.Message {
	mi := &file_coroutine_v1_coroutine_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Region.ProtoReflect.Descriptor instead.
func (*Region) Descriptor() ([]byte, []int) {
	return file_coroutine_v1_coroutine_proto_rawDescGZIP(), []int{2}
}

func (x *Region) GetType() uint32 {
	if x != nil {
		return x.Type
	}
	return 0
}

func (x *Region) GetArrayLength() uint32 {
	if x != nil {
		return x.ArrayLength
	}
	return 0
}

func (x *Region) GetData() []byte {
	if x != nil {
		return x.Data
	}
	return nil
}

var File_coroutine_v1_coroutine_proto protoreflect.FileDescriptor

var file_coroutine_v1_coroutine_proto_rawDesc = []byte{
	0x0a, 0x1c, 0x63, 0x6f, 0x72, 0x6f, 0x75, 0x74, 0x69, 0x6e, 0x65, 0x2f, 0x76, 0x31, 0x2f, 0x63,
	0x6f, 0x72, 0x6f, 0x75, 0x74, 0x69, 0x6e, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0c,
	0x63, 0x6f, 0x72, 0x6f, 0x75, 0x74, 0x69, 0x6e, 0x65, 0x2e, 0x76, 0x31, 0x1a, 0x1b, 0x63, 0x6f,
	0x72, 0x6f, 0x75, 0x74, 0x69, 0x6e, 0x65, 0x2f, 0x76, 0x31, 0x2f, 0x66, 0x75, 0x6e, 0x63, 0x74,
	0x69, 0x6f, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x17, 0x63, 0x6f, 0x72, 0x6f, 0x75,
	0x74, 0x69, 0x6e, 0x65, 0x2f, 0x76, 0x31, 0x2f, 0x74, 0x79, 0x70, 0x65, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x22, 0x9c, 0x02, 0x0a, 0x05, 0x53, 0x74, 0x61, 0x74, 0x65, 0x12, 0x29, 0x0a, 0x05,
	0x62, 0x75, 0x69, 0x6c, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x13, 0x2e, 0x63, 0x6f,
	0x72, 0x6f, 0x75, 0x74, 0x69, 0x6e, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x42, 0x75, 0x69, 0x6c, 0x64,
	0x52, 0x05, 0x62, 0x75, 0x69, 0x6c, 0x64, 0x12, 0x14, 0x0a, 0x05, 0x73, 0x74, 0x61, 0x74, 0x65,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x05, 0x73, 0x74, 0x61, 0x74, 0x65, 0x12, 0x28, 0x0a,
	0x05, 0x74, 0x79, 0x70, 0x65, 0x73, 0x18, 0x03, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x12, 0x2e, 0x63,
	0x6f, 0x72, 0x6f, 0x75, 0x74, 0x69, 0x6e, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x54, 0x79, 0x70, 0x65,
	0x52, 0x05, 0x74, 0x79, 0x70, 0x65, 0x73, 0x12, 0x34, 0x0a, 0x09, 0x66, 0x75, 0x6e, 0x63, 0x74,
	0x69, 0x6f, 0x6e, 0x73, 0x18, 0x04, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x16, 0x2e, 0x63, 0x6f, 0x72,
	0x6f, 0x75, 0x74, 0x69, 0x6e, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x46, 0x75, 0x6e, 0x63, 0x74, 0x69,
	0x6f, 0x6e, 0x52, 0x09, 0x66, 0x75, 0x6e, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x12, 0x2e, 0x0a,
	0x07, 0x72, 0x65, 0x67, 0x69, 0x6f, 0x6e, 0x73, 0x18, 0x05, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x14,
	0x2e, 0x63, 0x6f, 0x72, 0x6f, 0x75, 0x74, 0x69, 0x6e, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x52, 0x65,
	0x67, 0x69, 0x6f, 0x6e, 0x52, 0x07, 0x72, 0x65, 0x67, 0x69, 0x6f, 0x6e, 0x73, 0x12, 0x28, 0x0a,
	0x04, 0x72, 0x6f, 0x6f, 0x74, 0x18, 0x06, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x14, 0x2e, 0x63, 0x6f,
	0x72, 0x6f, 0x75, 0x74, 0x69, 0x6e, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x52, 0x65, 0x67, 0x69, 0x6f,
	0x6e, 0x52, 0x04, 0x72, 0x6f, 0x6f, 0x74, 0x12, 0x18, 0x0a, 0x07, 0x73, 0x74, 0x72, 0x69, 0x6e,
	0x67, 0x73, 0x18, 0x07, 0x20, 0x03, 0x28, 0x09, 0x52, 0x07, 0x73, 0x74, 0x72, 0x69, 0x6e, 0x67,
	0x73, 0x22, 0x3b, 0x0a, 0x05, 0x42, 0x75, 0x69, 0x6c, 0x64, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x12, 0x0e, 0x0a, 0x02, 0x6f, 0x73,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x6f, 0x73, 0x12, 0x12, 0x0a, 0x04, 0x61, 0x72,
	0x63, 0x68, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x61, 0x72, 0x63, 0x68, 0x22, 0x53,
	0x0a, 0x06, 0x52, 0x65, 0x67, 0x69, 0x6f, 0x6e, 0x12, 0x12, 0x0a, 0x04, 0x74, 0x79, 0x70, 0x65,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x04, 0x74, 0x79, 0x70, 0x65, 0x12, 0x21, 0x0a, 0x0c,
	0x61, 0x72, 0x72, 0x61, 0x79, 0x5f, 0x6c, 0x65, 0x6e, 0x67, 0x74, 0x68, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x0d, 0x52, 0x0b, 0x61, 0x72, 0x72, 0x61, 0x79, 0x4c, 0x65, 0x6e, 0x67, 0x74, 0x68, 0x12,
	0x12, 0x0a, 0x04, 0x64, 0x61, 0x74, 0x61, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x04, 0x64,
	0x61, 0x74, 0x61, 0x42, 0xbd, 0x01, 0x0a, 0x10, 0x63, 0x6f, 0x6d, 0x2e, 0x63, 0x6f, 0x72, 0x6f,
	0x75, 0x74, 0x69, 0x6e, 0x65, 0x2e, 0x76, 0x31, 0x42, 0x0e, 0x43, 0x6f, 0x72, 0x6f, 0x75, 0x74,
	0x69, 0x6e, 0x65, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x50, 0x01, 0x5a, 0x48, 0x67, 0x69, 0x74, 0x68,
	0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x73, 0x74, 0x65, 0x61, 0x6c, 0x74, 0x68, 0x72, 0x6f,
	0x63, 0x6b, 0x65, 0x74, 0x2f, 0x63, 0x6f, 0x72, 0x6f, 0x75, 0x74, 0x69, 0x6e, 0x65, 0x2f, 0x67,
	0x65, 0x6e, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x67, 0x6f, 0x2f, 0x63, 0x6f, 0x72, 0x6f,
	0x75, 0x74, 0x69, 0x6e, 0x65, 0x2f, 0x76, 0x31, 0x3b, 0x63, 0x6f, 0x72, 0x6f, 0x75, 0x74, 0x69,
	0x6e, 0x65, 0x76, 0x31, 0xa2, 0x02, 0x03, 0x43, 0x58, 0x58, 0xaa, 0x02, 0x0c, 0x43, 0x6f, 0x72,
	0x6f, 0x75, 0x74, 0x69, 0x6e, 0x65, 0x2e, 0x56, 0x31, 0xca, 0x02, 0x0c, 0x43, 0x6f, 0x72, 0x6f,
	0x75, 0x74, 0x69, 0x6e, 0x65, 0x5c, 0x56, 0x31, 0xe2, 0x02, 0x18, 0x43, 0x6f, 0x72, 0x6f, 0x75,
	0x74, 0x69, 0x6e, 0x65, 0x5c, 0x56, 0x31, 0x5c, 0x47, 0x50, 0x42, 0x4d, 0x65, 0x74, 0x61, 0x64,
	0x61, 0x74, 0x61, 0xea, 0x02, 0x0d, 0x43, 0x6f, 0x72, 0x6f, 0x75, 0x74, 0x69, 0x6e, 0x65, 0x3a,
	0x3a, 0x56, 0x31, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_coroutine_v1_coroutine_proto_rawDescOnce sync.Once
	file_coroutine_v1_coroutine_proto_rawDescData = file_coroutine_v1_coroutine_proto_rawDesc
)

func file_coroutine_v1_coroutine_proto_rawDescGZIP() []byte {
	file_coroutine_v1_coroutine_proto_rawDescOnce.Do(func() {
		file_coroutine_v1_coroutine_proto_rawDescData = protoimpl.X.CompressGZIP(file_coroutine_v1_coroutine_proto_rawDescData)
	})
	return file_coroutine_v1_coroutine_proto_rawDescData
}

var file_coroutine_v1_coroutine_proto_msgTypes = make([]protoimpl.MessageInfo, 3)
var file_coroutine_v1_coroutine_proto_goTypes = []interface{}{
	(*State)(nil),    // 0: coroutine.v1.State
	(*Build)(nil),    // 1: coroutine.v1.Build
	(*Region)(nil),   // 2: coroutine.v1.Region
	(*Type)(nil),     // 3: coroutine.v1.Type
	(*Function)(nil), // 4: coroutine.v1.Function
}
var file_coroutine_v1_coroutine_proto_depIdxs = []int32{
	1, // 0: coroutine.v1.State.build:type_name -> coroutine.v1.Build
	3, // 1: coroutine.v1.State.types:type_name -> coroutine.v1.Type
	4, // 2: coroutine.v1.State.functions:type_name -> coroutine.v1.Function
	2, // 3: coroutine.v1.State.regions:type_name -> coroutine.v1.Region
	2, // 4: coroutine.v1.State.root:type_name -> coroutine.v1.Region
	5, // [5:5] is the sub-list for method output_type
	5, // [5:5] is the sub-list for method input_type
	5, // [5:5] is the sub-list for extension type_name
	5, // [5:5] is the sub-list for extension extendee
	0, // [0:5] is the sub-list for field type_name
}

func init() { file_coroutine_v1_coroutine_proto_init() }
func file_coroutine_v1_coroutine_proto_init() {
	if File_coroutine_v1_coroutine_proto != nil {
		return
	}
	file_coroutine_v1_function_proto_init()
	file_coroutine_v1_type_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_coroutine_v1_coroutine_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*State); i {
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
		file_coroutine_v1_coroutine_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Build); i {
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
		file_coroutine_v1_coroutine_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Region); i {
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
			RawDescriptor: file_coroutine_v1_coroutine_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   3,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_coroutine_v1_coroutine_proto_goTypes,
		DependencyIndexes: file_coroutine_v1_coroutine_proto_depIdxs,
		MessageInfos:      file_coroutine_v1_coroutine_proto_msgTypes,
	}.Build()
	File_coroutine_v1_coroutine_proto = out.File
	file_coroutine_v1_coroutine_proto_rawDesc = nil
	file_coroutine_v1_coroutine_proto_goTypes = nil
	file_coroutine_v1_coroutine_proto_depIdxs = nil
}
