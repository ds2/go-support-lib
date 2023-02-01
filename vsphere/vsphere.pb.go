// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v3.21.12
// source: vsphere/vsphere.proto

package vsphere

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

type PowerState int32

const (
	PowerState_UnknownPowerState PowerState = 0
	PowerState_PoweredOn         PowerState = 1
	PowerState_PoweredOff        PowerState = 2
	PowerState_Suspended         PowerState = 3
)

// Enum value maps for PowerState.
var (
	PowerState_name = map[int32]string{
		0: "UnknownPowerState",
		1: "PoweredOn",
		2: "PoweredOff",
		3: "Suspended",
	}
	PowerState_value = map[string]int32{
		"UnknownPowerState": 0,
		"PoweredOn":         1,
		"PoweredOff":        2,
		"Suspended":         3,
	}
)

func (x PowerState) Enum() *PowerState {
	p := new(PowerState)
	*p = x
	return p
}

func (x PowerState) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (PowerState) Descriptor() protoreflect.EnumDescriptor {
	return file_vsphere_vsphere_proto_enumTypes[0].Descriptor()
}

func (PowerState) Type() protoreflect.EnumType {
	return &file_vsphere_vsphere_proto_enumTypes[0]
}

func (x PowerState) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use PowerState.Descriptor instead.
func (PowerState) EnumDescriptor() ([]byte, []int) {
	return file_vsphere_vsphere_proto_rawDescGZIP(), []int{0}
}

type OverallStatus int32

const (
	OverallStatus_UnknownOverallStatus OverallStatus = 0
	OverallStatus_GRAY                 OverallStatus = 1
	OverallStatus_GREEN                OverallStatus = 2
	OverallStatus_YELLOW               OverallStatus = 3
	OverallStatus_RED                  OverallStatus = 4
)

// Enum value maps for OverallStatus.
var (
	OverallStatus_name = map[int32]string{
		0: "UnknownOverallStatus",
		1: "GRAY",
		2: "GREEN",
		3: "YELLOW",
		4: "RED",
	}
	OverallStatus_value = map[string]int32{
		"UnknownOverallStatus": 0,
		"GRAY":                 1,
		"GREEN":                2,
		"YELLOW":               3,
		"RED":                  4,
	}
)

func (x OverallStatus) Enum() *OverallStatus {
	p := new(OverallStatus)
	*p = x
	return p
}

func (x OverallStatus) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (OverallStatus) Descriptor() protoreflect.EnumDescriptor {
	return file_vsphere_vsphere_proto_enumTypes[1].Descriptor()
}

func (OverallStatus) Type() protoreflect.EnumType {
	return &file_vsphere_vsphere_proto_enumTypes[1]
}

func (x OverallStatus) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use OverallStatus.Descriptor instead.
func (OverallStatus) EnumDescriptor() ([]byte, []int) {
	return file_vsphere_vsphere_proto_rawDescGZIP(), []int{1}
}

type DsVirtualMachine struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id            string        `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Name          string        `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	IpV4Addresses []string      `protobuf:"bytes,3,rep,name=ipV4Addresses,proto3" json:"ipV4Addresses,omitempty"`
	IpV6Addresses []string      `protobuf:"bytes,4,rep,name=ipV6Addresses,proto3" json:"ipV6Addresses,omitempty"`
	CpuCount      uint32        `protobuf:"varint,5,opt,name=cpuCount,proto3" json:"cpuCount,omitempty"`
	MemoryMB      uint64        `protobuf:"varint,6,opt,name=memoryMB,proto3" json:"memoryMB,omitempty"`
	PowerState    PowerState    `protobuf:"varint,7,opt,name=powerState,proto3,enum=vsphere.PowerState" json:"powerState,omitempty"`
	GuestId       string        `protobuf:"bytes,8,opt,name=guestId,proto3" json:"guestId,omitempty"`
	Overall       OverallStatus `protobuf:"varint,9,opt,name=overall,proto3,enum=vsphere.OverallStatus" json:"overall,omitempty"`
	Hostname      string        `protobuf:"bytes,10,opt,name=hostname,proto3" json:"hostname,omitempty"`
	UptimeSeconds uint32        `protobuf:"varint,11,opt,name=uptimeSeconds,proto3" json:"uptimeSeconds,omitempty"`
	InstanceUuid  string        `protobuf:"bytes,12,opt,name=instanceUuid,proto3" json:"instanceUuid,omitempty"`
}

func (x *DsVirtualMachine) Reset() {
	*x = DsVirtualMachine{}
	if protoimpl.UnsafeEnabled {
		mi := &file_vsphere_vsphere_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DsVirtualMachine) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DsVirtualMachine) ProtoMessage() {}

func (x *DsVirtualMachine) ProtoReflect() protoreflect.Message {
	mi := &file_vsphere_vsphere_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DsVirtualMachine.ProtoReflect.Descriptor instead.
func (*DsVirtualMachine) Descriptor() ([]byte, []int) {
	return file_vsphere_vsphere_proto_rawDescGZIP(), []int{0}
}

func (x *DsVirtualMachine) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *DsVirtualMachine) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *DsVirtualMachine) GetIpV4Addresses() []string {
	if x != nil {
		return x.IpV4Addresses
	}
	return nil
}

func (x *DsVirtualMachine) GetIpV6Addresses() []string {
	if x != nil {
		return x.IpV6Addresses
	}
	return nil
}

func (x *DsVirtualMachine) GetCpuCount() uint32 {
	if x != nil {
		return x.CpuCount
	}
	return 0
}

func (x *DsVirtualMachine) GetMemoryMB() uint64 {
	if x != nil {
		return x.MemoryMB
	}
	return 0
}

func (x *DsVirtualMachine) GetPowerState() PowerState {
	if x != nil {
		return x.PowerState
	}
	return PowerState_UnknownPowerState
}

func (x *DsVirtualMachine) GetGuestId() string {
	if x != nil {
		return x.GuestId
	}
	return ""
}

func (x *DsVirtualMachine) GetOverall() OverallStatus {
	if x != nil {
		return x.Overall
	}
	return OverallStatus_UnknownOverallStatus
}

func (x *DsVirtualMachine) GetHostname() string {
	if x != nil {
		return x.Hostname
	}
	return ""
}

func (x *DsVirtualMachine) GetUptimeSeconds() uint32 {
	if x != nil {
		return x.UptimeSeconds
	}
	return 0
}

func (x *DsVirtualMachine) GetInstanceUuid() string {
	if x != nil {
		return x.InstanceUuid
	}
	return ""
}

var File_vsphere_vsphere_proto protoreflect.FileDescriptor

var file_vsphere_vsphere_proto_rawDesc = []byte{
	0x0a, 0x15, 0x76, 0x73, 0x70, 0x68, 0x65, 0x72, 0x65, 0x2f, 0x76, 0x73, 0x70, 0x68, 0x65, 0x72,
	0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x07, 0x76, 0x73, 0x70, 0x68, 0x65, 0x72, 0x65,
	0x22, 0xa1, 0x03, 0x0a, 0x10, 0x44, 0x73, 0x56, 0x69, 0x72, 0x74, 0x75, 0x61, 0x6c, 0x4d, 0x61,
	0x63, 0x68, 0x69, 0x6e, 0x65, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x02, 0x69, 0x64, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x24, 0x0a, 0x0d, 0x69, 0x70, 0x56,
	0x34, 0x41, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x65, 0x73, 0x18, 0x03, 0x20, 0x03, 0x28, 0x09,
	0x52, 0x0d, 0x69, 0x70, 0x56, 0x34, 0x41, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x65, 0x73, 0x12,
	0x24, 0x0a, 0x0d, 0x69, 0x70, 0x56, 0x36, 0x41, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x65, 0x73,
	0x18, 0x04, 0x20, 0x03, 0x28, 0x09, 0x52, 0x0d, 0x69, 0x70, 0x56, 0x36, 0x41, 0x64, 0x64, 0x72,
	0x65, 0x73, 0x73, 0x65, 0x73, 0x12, 0x1a, 0x0a, 0x08, 0x63, 0x70, 0x75, 0x43, 0x6f, 0x75, 0x6e,
	0x74, 0x18, 0x05, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x08, 0x63, 0x70, 0x75, 0x43, 0x6f, 0x75, 0x6e,
	0x74, 0x12, 0x1a, 0x0a, 0x08, 0x6d, 0x65, 0x6d, 0x6f, 0x72, 0x79, 0x4d, 0x42, 0x18, 0x06, 0x20,
	0x01, 0x28, 0x04, 0x52, 0x08, 0x6d, 0x65, 0x6d, 0x6f, 0x72, 0x79, 0x4d, 0x42, 0x12, 0x33, 0x0a,
	0x0a, 0x70, 0x6f, 0x77, 0x65, 0x72, 0x53, 0x74, 0x61, 0x74, 0x65, 0x18, 0x07, 0x20, 0x01, 0x28,
	0x0e, 0x32, 0x13, 0x2e, 0x76, 0x73, 0x70, 0x68, 0x65, 0x72, 0x65, 0x2e, 0x50, 0x6f, 0x77, 0x65,
	0x72, 0x53, 0x74, 0x61, 0x74, 0x65, 0x52, 0x0a, 0x70, 0x6f, 0x77, 0x65, 0x72, 0x53, 0x74, 0x61,
	0x74, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x67, 0x75, 0x65, 0x73, 0x74, 0x49, 0x64, 0x18, 0x08, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x07, 0x67, 0x75, 0x65, 0x73, 0x74, 0x49, 0x64, 0x12, 0x30, 0x0a, 0x07,
	0x6f, 0x76, 0x65, 0x72, 0x61, 0x6c, 0x6c, 0x18, 0x09, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x16, 0x2e,
	0x76, 0x73, 0x70, 0x68, 0x65, 0x72, 0x65, 0x2e, 0x4f, 0x76, 0x65, 0x72, 0x61, 0x6c, 0x6c, 0x53,
	0x74, 0x61, 0x74, 0x75, 0x73, 0x52, 0x07, 0x6f, 0x76, 0x65, 0x72, 0x61, 0x6c, 0x6c, 0x12, 0x1a,
	0x0a, 0x08, 0x68, 0x6f, 0x73, 0x74, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x0a, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x08, 0x68, 0x6f, 0x73, 0x74, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x24, 0x0a, 0x0d, 0x75, 0x70,
	0x74, 0x69, 0x6d, 0x65, 0x53, 0x65, 0x63, 0x6f, 0x6e, 0x64, 0x73, 0x18, 0x0b, 0x20, 0x01, 0x28,
	0x0d, 0x52, 0x0d, 0x75, 0x70, 0x74, 0x69, 0x6d, 0x65, 0x53, 0x65, 0x63, 0x6f, 0x6e, 0x64, 0x73,
	0x12, 0x22, 0x0a, 0x0c, 0x69, 0x6e, 0x73, 0x74, 0x61, 0x6e, 0x63, 0x65, 0x55, 0x75, 0x69, 0x64,
	0x18, 0x0c, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0c, 0x69, 0x6e, 0x73, 0x74, 0x61, 0x6e, 0x63, 0x65,
	0x55, 0x75, 0x69, 0x64, 0x2a, 0x51, 0x0a, 0x0a, 0x50, 0x6f, 0x77, 0x65, 0x72, 0x53, 0x74, 0x61,
	0x74, 0x65, 0x12, 0x15, 0x0a, 0x11, 0x55, 0x6e, 0x6b, 0x6e, 0x6f, 0x77, 0x6e, 0x50, 0x6f, 0x77,
	0x65, 0x72, 0x53, 0x74, 0x61, 0x74, 0x65, 0x10, 0x00, 0x12, 0x0d, 0x0a, 0x09, 0x50, 0x6f, 0x77,
	0x65, 0x72, 0x65, 0x64, 0x4f, 0x6e, 0x10, 0x01, 0x12, 0x0e, 0x0a, 0x0a, 0x50, 0x6f, 0x77, 0x65,
	0x72, 0x65, 0x64, 0x4f, 0x66, 0x66, 0x10, 0x02, 0x12, 0x0d, 0x0a, 0x09, 0x53, 0x75, 0x73, 0x70,
	0x65, 0x6e, 0x64, 0x65, 0x64, 0x10, 0x03, 0x2a, 0x53, 0x0a, 0x0d, 0x4f, 0x76, 0x65, 0x72, 0x61,
	0x6c, 0x6c, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0x18, 0x0a, 0x14, 0x55, 0x6e, 0x6b, 0x6e,
	0x6f, 0x77, 0x6e, 0x4f, 0x76, 0x65, 0x72, 0x61, 0x6c, 0x6c, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73,
	0x10, 0x00, 0x12, 0x08, 0x0a, 0x04, 0x47, 0x52, 0x41, 0x59, 0x10, 0x01, 0x12, 0x09, 0x0a, 0x05,
	0x47, 0x52, 0x45, 0x45, 0x4e, 0x10, 0x02, 0x12, 0x0a, 0x0a, 0x06, 0x59, 0x45, 0x4c, 0x4c, 0x4f,
	0x57, 0x10, 0x03, 0x12, 0x07, 0x0a, 0x03, 0x52, 0x45, 0x44, 0x10, 0x04, 0x42, 0x0a, 0x5a, 0x08,
	0x2f, 0x76, 0x73, 0x70, 0x68, 0x65, 0x72, 0x65, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_vsphere_vsphere_proto_rawDescOnce sync.Once
	file_vsphere_vsphere_proto_rawDescData = file_vsphere_vsphere_proto_rawDesc
)

func file_vsphere_vsphere_proto_rawDescGZIP() []byte {
	file_vsphere_vsphere_proto_rawDescOnce.Do(func() {
		file_vsphere_vsphere_proto_rawDescData = protoimpl.X.CompressGZIP(file_vsphere_vsphere_proto_rawDescData)
	})
	return file_vsphere_vsphere_proto_rawDescData
}

var file_vsphere_vsphere_proto_enumTypes = make([]protoimpl.EnumInfo, 2)
var file_vsphere_vsphere_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_vsphere_vsphere_proto_goTypes = []interface{}{
	(PowerState)(0),          // 0: vsphere.PowerState
	(OverallStatus)(0),       // 1: vsphere.OverallStatus
	(*DsVirtualMachine)(nil), // 2: vsphere.DsVirtualMachine
}
var file_vsphere_vsphere_proto_depIdxs = []int32{
	0, // 0: vsphere.DsVirtualMachine.powerState:type_name -> vsphere.PowerState
	1, // 1: vsphere.DsVirtualMachine.overall:type_name -> vsphere.OverallStatus
	2, // [2:2] is the sub-list for method output_type
	2, // [2:2] is the sub-list for method input_type
	2, // [2:2] is the sub-list for extension type_name
	2, // [2:2] is the sub-list for extension extendee
	0, // [0:2] is the sub-list for field type_name
}

func init() { file_vsphere_vsphere_proto_init() }
func file_vsphere_vsphere_proto_init() {
	if File_vsphere_vsphere_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_vsphere_vsphere_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DsVirtualMachine); i {
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
			RawDescriptor: file_vsphere_vsphere_proto_rawDesc,
			NumEnums:      2,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_vsphere_vsphere_proto_goTypes,
		DependencyIndexes: file_vsphere_vsphere_proto_depIdxs,
		EnumInfos:         file_vsphere_vsphere_proto_enumTypes,
		MessageInfos:      file_vsphere_vsphere_proto_msgTypes,
	}.Build()
	File_vsphere_vsphere_proto = out.File
	file_vsphere_vsphere_proto_rawDesc = nil
	file_vsphere_vsphere_proto_goTypes = nil
	file_vsphere_vsphere_proto_depIdxs = nil
}