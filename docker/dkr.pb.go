// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v3.21.9
// source: docker/dkr.proto

package docker

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

type ContainerInfo struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id              string   `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Image           string   `protobuf:"bytes,2,opt,name=image,proto3" json:"image,omitempty"`
	Started         int64    `protobuf:"varint,3,opt,name=started,proto3" json:"started,omitempty"`
	Names           []string `protobuf:"bytes,4,rep,name=names,proto3" json:"names,omitempty"`
	Imageid         string   `protobuf:"bytes,5,opt,name=imageid,proto3" json:"imageid,omitempty"`
	InternalCommand string   `protobuf:"bytes,6,opt,name=internalCommand,proto3" json:"internalCommand,omitempty"`
	State           string   `protobuf:"bytes,7,opt,name=state,proto3" json:"state,omitempty"`
	Status          string   `protobuf:"bytes,8,opt,name=status,proto3" json:"status,omitempty"`
}

func (x *ContainerInfo) Reset() {
	*x = ContainerInfo{}
	if protoimpl.UnsafeEnabled {
		mi := &file_docker_dkr_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ContainerInfo) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ContainerInfo) ProtoMessage() {}

func (x *ContainerInfo) ProtoReflect() protoreflect.Message {
	mi := &file_docker_dkr_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ContainerInfo.ProtoReflect.Descriptor instead.
func (*ContainerInfo) Descriptor() ([]byte, []int) {
	return file_docker_dkr_proto_rawDescGZIP(), []int{0}
}

func (x *ContainerInfo) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *ContainerInfo) GetImage() string {
	if x != nil {
		return x.Image
	}
	return ""
}

func (x *ContainerInfo) GetStarted() int64 {
	if x != nil {
		return x.Started
	}
	return 0
}

func (x *ContainerInfo) GetNames() []string {
	if x != nil {
		return x.Names
	}
	return nil
}

func (x *ContainerInfo) GetImageid() string {
	if x != nil {
		return x.Imageid
	}
	return ""
}

func (x *ContainerInfo) GetInternalCommand() string {
	if x != nil {
		return x.InternalCommand
	}
	return ""
}

func (x *ContainerInfo) GetState() string {
	if x != nil {
		return x.State
	}
	return ""
}

func (x *ContainerInfo) GetStatus() string {
	if x != nil {
		return x.Status
	}
	return ""
}

var File_docker_dkr_proto protoreflect.FileDescriptor

var file_docker_dkr_proto_rawDesc = []byte{
	0x0a, 0x10, 0x64, 0x6f, 0x63, 0x6b, 0x65, 0x72, 0x2f, 0x64, 0x6b, 0x72, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x12, 0x06, 0x64, 0x6f, 0x63, 0x6b, 0x65, 0x72, 0x22, 0xd7, 0x01, 0x0a, 0x0d, 0x43,
	0x6f, 0x6e, 0x74, 0x61, 0x69, 0x6e, 0x65, 0x72, 0x49, 0x6e, 0x66, 0x6f, 0x12, 0x0e, 0x0a, 0x02,
	0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x12, 0x14, 0x0a, 0x05,
	0x69, 0x6d, 0x61, 0x67, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x69, 0x6d, 0x61,
	0x67, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x73, 0x74, 0x61, 0x72, 0x74, 0x65, 0x64, 0x18, 0x03, 0x20,
	0x01, 0x28, 0x03, 0x52, 0x07, 0x73, 0x74, 0x61, 0x72, 0x74, 0x65, 0x64, 0x12, 0x14, 0x0a, 0x05,
	0x6e, 0x61, 0x6d, 0x65, 0x73, 0x18, 0x04, 0x20, 0x03, 0x28, 0x09, 0x52, 0x05, 0x6e, 0x61, 0x6d,
	0x65, 0x73, 0x12, 0x18, 0x0a, 0x07, 0x69, 0x6d, 0x61, 0x67, 0x65, 0x69, 0x64, 0x18, 0x05, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x07, 0x69, 0x6d, 0x61, 0x67, 0x65, 0x69, 0x64, 0x12, 0x28, 0x0a, 0x0f,
	0x69, 0x6e, 0x74, 0x65, 0x72, 0x6e, 0x61, 0x6c, 0x43, 0x6f, 0x6d, 0x6d, 0x61, 0x6e, 0x64, 0x18,
	0x06, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0f, 0x69, 0x6e, 0x74, 0x65, 0x72, 0x6e, 0x61, 0x6c, 0x43,
	0x6f, 0x6d, 0x6d, 0x61, 0x6e, 0x64, 0x12, 0x14, 0x0a, 0x05, 0x73, 0x74, 0x61, 0x74, 0x65, 0x18,
	0x07, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x73, 0x74, 0x61, 0x74, 0x65, 0x12, 0x16, 0x0a, 0x06,
	0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x18, 0x08, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x73, 0x74,
	0x61, 0x74, 0x75, 0x73, 0x42, 0x09, 0x5a, 0x07, 0x2f, 0x64, 0x6f, 0x63, 0x6b, 0x65, 0x72, 0x62,
	0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_docker_dkr_proto_rawDescOnce sync.Once
	file_docker_dkr_proto_rawDescData = file_docker_dkr_proto_rawDesc
)

func file_docker_dkr_proto_rawDescGZIP() []byte {
	file_docker_dkr_proto_rawDescOnce.Do(func() {
		file_docker_dkr_proto_rawDescData = protoimpl.X.CompressGZIP(file_docker_dkr_proto_rawDescData)
	})
	return file_docker_dkr_proto_rawDescData
}

var file_docker_dkr_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_docker_dkr_proto_goTypes = []interface{}{
	(*ContainerInfo)(nil), // 0: docker.ContainerInfo
}
var file_docker_dkr_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_docker_dkr_proto_init() }
func file_docker_dkr_proto_init() {
	if File_docker_dkr_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_docker_dkr_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ContainerInfo); i {
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
			RawDescriptor: file_docker_dkr_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_docker_dkr_proto_goTypes,
		DependencyIndexes: file_docker_dkr_proto_depIdxs,
		MessageInfos:      file_docker_dkr_proto_msgTypes,
	}.Build()
	File_docker_dkr_proto = out.File
	file_docker_dkr_proto_rawDesc = nil
	file_docker_dkr_proto_goTypes = nil
	file_docker_dkr_proto_depIdxs = nil
}
