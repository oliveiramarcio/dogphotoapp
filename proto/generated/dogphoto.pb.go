// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.26.0
// 	protoc        v4.23.3
// source: dogphoto.proto

package generated

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

type DogPhotoRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Breed string `protobuf:"bytes,1,opt,name=breed,proto3" json:"breed,omitempty"`
}

func (x *DogPhotoRequest) Reset() {
	*x = DogPhotoRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_dogphoto_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DogPhotoRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DogPhotoRequest) ProtoMessage() {}

func (x *DogPhotoRequest) ProtoReflect() protoreflect.Message {
	mi := &file_dogphoto_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DogPhotoRequest.ProtoReflect.Descriptor instead.
func (*DogPhotoRequest) Descriptor() ([]byte, []int) {
	return file_dogphoto_proto_rawDescGZIP(), []int{0}
}

func (x *DogPhotoRequest) GetBreed() string {
	if x != nil {
		return x.Breed
	}
	return ""
}

type DogPhotoResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Status  string `protobuf:"bytes,1,opt,name=status,proto3" json:"status,omitempty"`
	Message string `protobuf:"bytes,2,opt,name=message,proto3" json:"message,omitempty"`
}

func (x *DogPhotoResponse) Reset() {
	*x = DogPhotoResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_dogphoto_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DogPhotoResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DogPhotoResponse) ProtoMessage() {}

func (x *DogPhotoResponse) ProtoReflect() protoreflect.Message {
	mi := &file_dogphoto_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DogPhotoResponse.ProtoReflect.Descriptor instead.
func (*DogPhotoResponse) Descriptor() ([]byte, []int) {
	return file_dogphoto_proto_rawDescGZIP(), []int{1}
}

func (x *DogPhotoResponse) GetStatus() string {
	if x != nil {
		return x.Status
	}
	return ""
}

func (x *DogPhotoResponse) GetMessage() string {
	if x != nil {
		return x.Message
	}
	return ""
}

var File_dogphoto_proto protoreflect.FileDescriptor

var file_dogphoto_proto_rawDesc = []byte{
	0x0a, 0x0e, 0x64, 0x6f, 0x67, 0x70, 0x68, 0x6f, 0x74, 0x6f, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x22, 0x27, 0x0a, 0x0f, 0x44, 0x6f, 0x67, 0x50, 0x68, 0x6f, 0x74, 0x6f, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x12, 0x14, 0x0a, 0x05, 0x62, 0x72, 0x65, 0x65, 0x64, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x05, 0x62, 0x72, 0x65, 0x65, 0x64, 0x22, 0x44, 0x0a, 0x10, 0x44, 0x6f, 0x67,
	0x50, 0x68, 0x6f, 0x74, 0x6f, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x16, 0x0a,
	0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x73,
	0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0x18, 0x0a, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x32,
	0x47, 0x0a, 0x0f, 0x44, 0x6f, 0x67, 0x50, 0x68, 0x6f, 0x74, 0x6f, 0x53, 0x65, 0x72, 0x76, 0x69,
	0x63, 0x65, 0x12, 0x34, 0x0a, 0x0b, 0x47, 0x65, 0x74, 0x44, 0x6f, 0x67, 0x50, 0x68, 0x6f, 0x74,
	0x6f, 0x12, 0x10, 0x2e, 0x44, 0x6f, 0x67, 0x50, 0x68, 0x6f, 0x74, 0x6f, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x1a, 0x11, 0x2e, 0x44, 0x6f, 0x67, 0x50, 0x68, 0x6f, 0x74, 0x6f, 0x52, 0x65,
	0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x42, 0x11, 0x5a, 0x0f, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x2f, 0x67, 0x65, 0x6e, 0x65, 0x72, 0x61, 0x74, 0x65, 0x64, 0x62, 0x06, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x33,
}

var (
	file_dogphoto_proto_rawDescOnce sync.Once
	file_dogphoto_proto_rawDescData = file_dogphoto_proto_rawDesc
)

func file_dogphoto_proto_rawDescGZIP() []byte {
	file_dogphoto_proto_rawDescOnce.Do(func() {
		file_dogphoto_proto_rawDescData = protoimpl.X.CompressGZIP(file_dogphoto_proto_rawDescData)
	})
	return file_dogphoto_proto_rawDescData
}

var file_dogphoto_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_dogphoto_proto_goTypes = []interface{}{
	(*DogPhotoRequest)(nil),  // 0: DogPhotoRequest
	(*DogPhotoResponse)(nil), // 1: DogPhotoResponse
}
var file_dogphoto_proto_depIdxs = []int32{
	0, // 0: DogPhotoService.GetDogPhoto:input_type -> DogPhotoRequest
	1, // 1: DogPhotoService.GetDogPhoto:output_type -> DogPhotoResponse
	1, // [1:2] is the sub-list for method output_type
	0, // [0:1] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_dogphoto_proto_init() }
func file_dogphoto_proto_init() {
	if File_dogphoto_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_dogphoto_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DogPhotoRequest); i {
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
		file_dogphoto_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DogPhotoResponse); i {
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
			RawDescriptor: file_dogphoto_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_dogphoto_proto_goTypes,
		DependencyIndexes: file_dogphoto_proto_depIdxs,
		MessageInfos:      file_dogphoto_proto_msgTypes,
	}.Build()
	File_dogphoto_proto = out.File
	file_dogphoto_proto_rawDesc = nil
	file_dogphoto_proto_goTypes = nil
	file_dogphoto_proto_depIdxs = nil
}
