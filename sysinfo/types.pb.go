// Code generated by protoc-gen-go. DO NOT EDIT.
// source: sysinfo/types.proto

package sysinfo

import (
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	math "math"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion3 // please upgrade the proto package

type IpAddressType int32

const (
	IpAddressType_UNKNOWN_IPADDR_TYPE IpAddressType = 0
	IpAddressType_IPv4                IpAddressType = 1
	IpAddressType_IPv6                IpAddressType = 2
)

var IpAddressType_name = map[int32]string{
	0: "UNKNOWN_IPADDR_TYPE",
	1: "IPv4",
	2: "IPv6",
}

var IpAddressType_value = map[string]int32{
	"UNKNOWN_IPADDR_TYPE": 0,
	"IPv4":                1,
	"IPv6":                2,
}

func (x IpAddressType) String() string {
	return proto.EnumName(IpAddressType_name, int32(x))
}

func (IpAddressType) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_ae30f4e6357a72f8, []int{0}
}

type NetworkInterfaceType int32

const (
	NetworkInterfaceType_UNKNOWN  NetworkInterfaceType = 0
	NetworkInterfaceType_ETHERNET NetworkInterfaceType = 1
	NetworkInterfaceType_WIRELESS NetworkInterfaceType = 2
)

var NetworkInterfaceType_name = map[int32]string{
	0: "UNKNOWN",
	1: "ETHERNET",
	2: "WIRELESS",
}

var NetworkInterfaceType_value = map[string]int32{
	"UNKNOWN":  0,
	"ETHERNET": 1,
	"WIRELESS": 2,
}

func (x NetworkInterfaceType) String() string {
	return proto.EnumName(NetworkInterfaceType_name, int32(x))
}

func (NetworkInterfaceType) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_ae30f4e6357a72f8, []int{1}
}

type PartitionInfo struct {
	MountPath            string   `protobuf:"bytes,1,opt,name=MountPath,proto3" json:"MountPath,omitempty"`
	Size                 uint64   `protobuf:"varint,2,opt,name=size,proto3" json:"size,omitempty"`
	Free                 uint64   `protobuf:"varint,3,opt,name=free,proto3" json:"free,omitempty"`
	FsType               string   `protobuf:"bytes,4,opt,name=fsType,proto3" json:"fsType,omitempty"`
	Device               string   `protobuf:"bytes,5,opt,name=device,proto3" json:"device,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *PartitionInfo) Reset()         { *m = PartitionInfo{} }
func (m *PartitionInfo) String() string { return proto.CompactTextString(m) }
func (*PartitionInfo) ProtoMessage()    {}
func (*PartitionInfo) Descriptor() ([]byte, []int) {
	return fileDescriptor_ae30f4e6357a72f8, []int{0}
}

func (m *PartitionInfo) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_PartitionInfo.Unmarshal(m, b)
}
func (m *PartitionInfo) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_PartitionInfo.Marshal(b, m, deterministic)
}
func (m *PartitionInfo) XXX_Merge(src proto.Message) {
	xxx_messageInfo_PartitionInfo.Merge(m, src)
}
func (m *PartitionInfo) XXX_Size() int {
	return xxx_messageInfo_PartitionInfo.Size(m)
}
func (m *PartitionInfo) XXX_DiscardUnknown() {
	xxx_messageInfo_PartitionInfo.DiscardUnknown(m)
}

var xxx_messageInfo_PartitionInfo proto.InternalMessageInfo

func (m *PartitionInfo) GetMountPath() string {
	if m != nil {
		return m.MountPath
	}
	return ""
}

func (m *PartitionInfo) GetSize() uint64 {
	if m != nil {
		return m.Size
	}
	return 0
}

func (m *PartitionInfo) GetFree() uint64 {
	if m != nil {
		return m.Free
	}
	return 0
}

func (m *PartitionInfo) GetFsType() string {
	if m != nil {
		return m.FsType
	}
	return ""
}

func (m *PartitionInfo) GetDevice() string {
	if m != nil {
		return m.Device
	}
	return ""
}

type IpAddress struct {
	Address              string        `protobuf:"bytes,1,opt,name=address,proto3" json:"address,omitempty"`
	Type                 IpAddressType `protobuf:"varint,2,opt,name=type,proto3,enum=sysinfo.IpAddressType" json:"type,omitempty"`
	XXX_NoUnkeyedLiteral struct{}      `json:"-"`
	XXX_unrecognized     []byte        `json:"-"`
	XXX_sizecache        int32         `json:"-"`
}

func (m *IpAddress) Reset()         { *m = IpAddress{} }
func (m *IpAddress) String() string { return proto.CompactTextString(m) }
func (*IpAddress) ProtoMessage()    {}
func (*IpAddress) Descriptor() ([]byte, []int) {
	return fileDescriptor_ae30f4e6357a72f8, []int{1}
}

func (m *IpAddress) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_IpAddress.Unmarshal(m, b)
}
func (m *IpAddress) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_IpAddress.Marshal(b, m, deterministic)
}
func (m *IpAddress) XXX_Merge(src proto.Message) {
	xxx_messageInfo_IpAddress.Merge(m, src)
}
func (m *IpAddress) XXX_Size() int {
	return xxx_messageInfo_IpAddress.Size(m)
}
func (m *IpAddress) XXX_DiscardUnknown() {
	xxx_messageInfo_IpAddress.DiscardUnknown(m)
}

var xxx_messageInfo_IpAddress proto.InternalMessageInfo

func (m *IpAddress) GetAddress() string {
	if m != nil {
		return m.Address
	}
	return ""
}

func (m *IpAddress) GetType() IpAddressType {
	if m != nil {
		return m.Type
	}
	return IpAddressType_UNKNOWN_IPADDR_TYPE
}

type NetworkInterfaceData struct {
	Name                 string               `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	DeviceName           string               `protobuf:"bytes,2,opt,name=deviceName,proto3" json:"deviceName,omitempty"`
	Type                 NetworkInterfaceType `protobuf:"varint,3,opt,name=type,proto3,enum=sysinfo.NetworkInterfaceType" json:"type,omitempty"`
	HwAddr               string               `protobuf:"bytes,4,opt,name=hwAddr,proto3" json:"hwAddr,omitempty"`
	Mtu                  uint32               `protobuf:"varint,5,opt,name=mtu,proto3" json:"mtu,omitempty"`
	Address              []*IpAddress         `protobuf:"bytes,6,rep,name=address,proto3" json:"address,omitempty"`
	XXX_NoUnkeyedLiteral struct{}             `json:"-"`
	XXX_unrecognized     []byte               `json:"-"`
	XXX_sizecache        int32                `json:"-"`
}

func (m *NetworkInterfaceData) Reset()         { *m = NetworkInterfaceData{} }
func (m *NetworkInterfaceData) String() string { return proto.CompactTextString(m) }
func (*NetworkInterfaceData) ProtoMessage()    {}
func (*NetworkInterfaceData) Descriptor() ([]byte, []int) {
	return fileDescriptor_ae30f4e6357a72f8, []int{2}
}

func (m *NetworkInterfaceData) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_NetworkInterfaceData.Unmarshal(m, b)
}
func (m *NetworkInterfaceData) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_NetworkInterfaceData.Marshal(b, m, deterministic)
}
func (m *NetworkInterfaceData) XXX_Merge(src proto.Message) {
	xxx_messageInfo_NetworkInterfaceData.Merge(m, src)
}
func (m *NetworkInterfaceData) XXX_Size() int {
	return xxx_messageInfo_NetworkInterfaceData.Size(m)
}
func (m *NetworkInterfaceData) XXX_DiscardUnknown() {
	xxx_messageInfo_NetworkInterfaceData.DiscardUnknown(m)
}

var xxx_messageInfo_NetworkInterfaceData proto.InternalMessageInfo

func (m *NetworkInterfaceData) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *NetworkInterfaceData) GetDeviceName() string {
	if m != nil {
		return m.DeviceName
	}
	return ""
}

func (m *NetworkInterfaceData) GetType() NetworkInterfaceType {
	if m != nil {
		return m.Type
	}
	return NetworkInterfaceType_UNKNOWN
}

func (m *NetworkInterfaceData) GetHwAddr() string {
	if m != nil {
		return m.HwAddr
	}
	return ""
}

func (m *NetworkInterfaceData) GetMtu() uint32 {
	if m != nil {
		return m.Mtu
	}
	return 0
}

func (m *NetworkInterfaceData) GetAddress() []*IpAddress {
	if m != nil {
		return m.Address
	}
	return nil
}

type HostInfo struct {
	HostName             string           `protobuf:"bytes,1,opt,name=HostName,proto3" json:"HostName,omitempty"`
	OsName               string           `protobuf:"bytes,2,opt,name=osName,proto3" json:"osName,omitempty"`
	OsVersion            string           `protobuf:"bytes,3,opt,name=osVersion,proto3" json:"osVersion,omitempty"`
	NumCores             uint32           `protobuf:"varint,4,opt,name=numCores,proto3" json:"numCores,omitempty"`
	FileSystems          []*PartitionInfo `protobuf:"bytes,5,rep,name=fileSystems,proto3" json:"fileSystems,omitempty"`
	XXX_NoUnkeyedLiteral struct{}         `json:"-"`
	XXX_unrecognized     []byte           `json:"-"`
	XXX_sizecache        int32            `json:"-"`
}

func (m *HostInfo) Reset()         { *m = HostInfo{} }
func (m *HostInfo) String() string { return proto.CompactTextString(m) }
func (*HostInfo) ProtoMessage()    {}
func (*HostInfo) Descriptor() ([]byte, []int) {
	return fileDescriptor_ae30f4e6357a72f8, []int{3}
}

func (m *HostInfo) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_HostInfo.Unmarshal(m, b)
}
func (m *HostInfo) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_HostInfo.Marshal(b, m, deterministic)
}
func (m *HostInfo) XXX_Merge(src proto.Message) {
	xxx_messageInfo_HostInfo.Merge(m, src)
}
func (m *HostInfo) XXX_Size() int {
	return xxx_messageInfo_HostInfo.Size(m)
}
func (m *HostInfo) XXX_DiscardUnknown() {
	xxx_messageInfo_HostInfo.DiscardUnknown(m)
}

var xxx_messageInfo_HostInfo proto.InternalMessageInfo

func (m *HostInfo) GetHostName() string {
	if m != nil {
		return m.HostName
	}
	return ""
}

func (m *HostInfo) GetOsName() string {
	if m != nil {
		return m.OsName
	}
	return ""
}

func (m *HostInfo) GetOsVersion() string {
	if m != nil {
		return m.OsVersion
	}
	return ""
}

func (m *HostInfo) GetNumCores() uint32 {
	if m != nil {
		return m.NumCores
	}
	return 0
}

func (m *HostInfo) GetFileSystems() []*PartitionInfo {
	if m != nil {
		return m.FileSystems
	}
	return nil
}

type ProcessDetails struct {
	ProcessId            uint64   `protobuf:"varint,1,opt,name=processId,proto3" json:"processId,omitempty"`
	Path                 string   `protobuf:"bytes,2,opt,name=path,proto3" json:"path,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ProcessDetails) Reset()         { *m = ProcessDetails{} }
func (m *ProcessDetails) String() string { return proto.CompactTextString(m) }
func (*ProcessDetails) ProtoMessage()    {}
func (*ProcessDetails) Descriptor() ([]byte, []int) {
	return fileDescriptor_ae30f4e6357a72f8, []int{4}
}

func (m *ProcessDetails) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ProcessDetails.Unmarshal(m, b)
}
func (m *ProcessDetails) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ProcessDetails.Marshal(b, m, deterministic)
}
func (m *ProcessDetails) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ProcessDetails.Merge(m, src)
}
func (m *ProcessDetails) XXX_Size() int {
	return xxx_messageInfo_ProcessDetails.Size(m)
}
func (m *ProcessDetails) XXX_DiscardUnknown() {
	xxx_messageInfo_ProcessDetails.DiscardUnknown(m)
}

var xxx_messageInfo_ProcessDetails proto.InternalMessageInfo

func (m *ProcessDetails) GetProcessId() uint64 {
	if m != nil {
		return m.ProcessId
	}
	return 0
}

func (m *ProcessDetails) GetPath() string {
	if m != nil {
		return m.Path
	}
	return ""
}

type HealthInfo struct {
	TotalMemory          uint64   `protobuf:"varint,1,opt,name=totalMemory,proto3" json:"totalMemory,omitempty"`
	AvailableMemory      uint64   `protobuf:"varint,2,opt,name=availableMemory,proto3" json:"availableMemory,omitempty"`
	CpuLoad1             float64  `protobuf:"fixed64,3,opt,name=cpuLoad1,proto3" json:"cpuLoad1,omitempty"`
	CpuLoad5             float64  `protobuf:"fixed64,4,opt,name=cpuLoad5,proto3" json:"cpuLoad5,omitempty"`
	CpuLoad15            float64  `protobuf:"fixed64,5,opt,name=cpuLoad15,proto3" json:"cpuLoad15,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *HealthInfo) Reset()         { *m = HealthInfo{} }
func (m *HealthInfo) String() string { return proto.CompactTextString(m) }
func (*HealthInfo) ProtoMessage()    {}
func (*HealthInfo) Descriptor() ([]byte, []int) {
	return fileDescriptor_ae30f4e6357a72f8, []int{5}
}

func (m *HealthInfo) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_HealthInfo.Unmarshal(m, b)
}
func (m *HealthInfo) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_HealthInfo.Marshal(b, m, deterministic)
}
func (m *HealthInfo) XXX_Merge(src proto.Message) {
	xxx_messageInfo_HealthInfo.Merge(m, src)
}
func (m *HealthInfo) XXX_Size() int {
	return xxx_messageInfo_HealthInfo.Size(m)
}
func (m *HealthInfo) XXX_DiscardUnknown() {
	xxx_messageInfo_HealthInfo.DiscardUnknown(m)
}

var xxx_messageInfo_HealthInfo proto.InternalMessageInfo

func (m *HealthInfo) GetTotalMemory() uint64 {
	if m != nil {
		return m.TotalMemory
	}
	return 0
}

func (m *HealthInfo) GetAvailableMemory() uint64 {
	if m != nil {
		return m.AvailableMemory
	}
	return 0
}

func (m *HealthInfo) GetCpuLoad1() float64 {
	if m != nil {
		return m.CpuLoad1
	}
	return 0
}

func (m *HealthInfo) GetCpuLoad5() float64 {
	if m != nil {
		return m.CpuLoad5
	}
	return 0
}

func (m *HealthInfo) GetCpuLoad15() float64 {
	if m != nil {
		return m.CpuLoad15
	}
	return 0
}

func init() {
	proto.RegisterEnum("sysinfo.IpAddressType", IpAddressType_name, IpAddressType_value)
	proto.RegisterEnum("sysinfo.NetworkInterfaceType", NetworkInterfaceType_name, NetworkInterfaceType_value)
	proto.RegisterType((*PartitionInfo)(nil), "sysinfo.PartitionInfo")
	proto.RegisterType((*IpAddress)(nil), "sysinfo.IpAddress")
	proto.RegisterType((*NetworkInterfaceData)(nil), "sysinfo.NetworkInterfaceData")
	proto.RegisterType((*HostInfo)(nil), "sysinfo.HostInfo")
	proto.RegisterType((*ProcessDetails)(nil), "sysinfo.ProcessDetails")
	proto.RegisterType((*HealthInfo)(nil), "sysinfo.HealthInfo")
}

func init() { proto.RegisterFile("sysinfo/types.proto", fileDescriptor_ae30f4e6357a72f8) }

var fileDescriptor_ae30f4e6357a72f8 = []byte{
	// 554 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x64, 0x53, 0x5f, 0x6f, 0xd3, 0x3e,
	0x14, 0x9d, 0xdb, 0xac, 0x7f, 0x6e, 0x7f, 0xdd, 0x2f, 0xf2, 0xd0, 0x88, 0x10, 0xa0, 0x2a, 0x4f,
	0x55, 0x85, 0x8a, 0x36, 0x28, 0xe2, 0x01, 0x09, 0x0d, 0x1a, 0xa9, 0x11, 0x5b, 0x08, 0x6e, 0x61,
	0xe2, 0x69, 0xf2, 0x5a, 0x47, 0x8d, 0x48, 0xe3, 0x28, 0x76, 0x3b, 0x95, 0x77, 0xbe, 0x0e, 0x0f,
	0x7c, 0x17, 0xbe, 0x0f, 0xb2, 0xe3, 0x26, 0xdd, 0xf6, 0x76, 0xce, 0xb1, 0x73, 0xef, 0x39, 0xbe,
	0x37, 0x70, 0x2c, 0xb6, 0x22, 0x4e, 0x23, 0xfe, 0x52, 0x6e, 0x33, 0x26, 0x86, 0x59, 0xce, 0x25,
	0xc7, 0x4d, 0x23, 0xba, 0xbf, 0x10, 0x74, 0x43, 0x9a, 0xcb, 0x58, 0xc6, 0x3c, 0xf5, 0xd3, 0x88,
	0xe3, 0xa7, 0xd0, 0xbe, 0xe4, 0xeb, 0x54, 0x86, 0x54, 0x2e, 0x1d, 0xd4, 0x43, 0xfd, 0x36, 0xa9,
	0x04, 0x8c, 0xc1, 0x12, 0xf1, 0x4f, 0xe6, 0xd4, 0x7a, 0xa8, 0x6f, 0x11, 0x8d, 0x95, 0x16, 0xe5,
	0x8c, 0x39, 0xf5, 0x42, 0x53, 0x18, 0x9f, 0x40, 0x23, 0x12, 0xb3, 0x6d, 0xc6, 0x1c, 0x4b, 0x97,
	0x30, 0x4c, 0xe9, 0x0b, 0xb6, 0x89, 0xe7, 0xcc, 0x39, 0x2c, 0xf4, 0x82, 0xb9, 0x5f, 0xa0, 0xed,
	0x67, 0xe7, 0x8b, 0x45, 0xce, 0x84, 0xc0, 0x0e, 0x34, 0x69, 0x01, 0x8d, 0x81, 0x1d, 0xc5, 0x03,
	0xb0, 0x54, 0x0c, 0xdd, 0xfe, 0xe8, 0xec, 0x64, 0x68, 0x62, 0x0c, 0xcb, 0x6f, 0x55, 0x13, 0xa2,
	0xef, 0xb8, 0x7f, 0x11, 0x3c, 0x0a, 0x98, 0xbc, 0xe5, 0xf9, 0x0f, 0x3f, 0x95, 0x2c, 0x8f, 0xe8,
	0x9c, 0x8d, 0xa9, 0xa4, 0xca, 0x6f, 0x4a, 0x57, 0xcc, 0xd4, 0xd6, 0x18, 0x3f, 0x07, 0x28, 0x9c,
	0x04, 0xea, 0xa4, 0xa6, 0x4f, 0xf6, 0x14, 0x7c, 0x6a, 0x1a, 0xd7, 0x75, 0xe3, 0x67, 0x65, 0xe3,
	0xfb, 0x0d, 0xaa, 0xfe, 0x2a, 0xea, 0xf2, 0x56, 0xd9, 0xda, 0x3d, 0x41, 0xc1, 0xb0, 0x0d, 0xf5,
	0x95, 0x5c, 0xeb, 0xfc, 0x5d, 0xa2, 0x20, 0x7e, 0x51, 0xe5, 0x6d, 0xf4, 0xea, 0xfd, 0xce, 0x19,
	0x7e, 0x18, 0xac, 0x7c, 0x03, 0xf7, 0x0f, 0x82, 0xd6, 0x84, 0x0b, 0xa9, 0xa7, 0xf5, 0xa4, 0xc0,
	0x41, 0x95, 0xa7, 0xe4, 0xca, 0x00, 0x17, 0x7b, 0x79, 0x0c, 0x53, 0x13, 0xe6, 0xe2, 0x1b, 0xcb,
	0x45, 0xcc, 0x53, 0x1d, 0xa8, 0x4d, 0x2a, 0x41, 0x55, 0x4c, 0xd7, 0xab, 0x8f, 0x3c, 0x67, 0x42,
	0x1b, 0xef, 0x92, 0x92, 0xe3, 0xb7, 0xd0, 0x89, 0xe2, 0x84, 0x4d, 0xb7, 0x42, 0xb2, 0x95, 0x70,
	0x0e, 0xb5, 0xd9, 0x6a, 0x0a, 0x77, 0x16, 0x89, 0xec, 0x5f, 0x75, 0x3f, 0xc0, 0x51, 0x98, 0xf3,
	0x39, 0x13, 0x62, 0xcc, 0x24, 0x8d, 0x13, 0xa1, 0x5c, 0x64, 0x85, 0xe2, 0x2f, 0xb4, 0x75, 0x8b,
	0x54, 0x82, 0x9a, 0x51, 0xa6, 0x16, 0xb0, 0x70, 0xae, 0xb1, 0xfb, 0x1b, 0x01, 0x4c, 0x18, 0x4d,
	0xe4, 0x52, 0x47, 0xef, 0x41, 0x47, 0x72, 0x49, 0x93, 0x4b, 0xb6, 0xe2, 0xf9, 0xd6, 0x94, 0xd8,
	0x97, 0x70, 0x1f, 0xfe, 0xa7, 0x1b, 0x1a, 0x27, 0xf4, 0x26, 0x61, 0xe6, 0x56, 0xb1, 0xb7, 0xf7,
	0x65, 0x15, 0x7a, 0x9e, 0xad, 0x2f, 0x38, 0x5d, 0x9c, 0xea, 0x17, 0x41, 0xa4, 0xe4, 0x7b, 0x67,
	0x23, 0xfd, 0x20, 0xd5, 0xd9, 0x48, 0x85, 0xd8, 0xdd, 0x1b, 0xe9, 0x89, 0x22, 0x52, 0x09, 0x83,
	0x77, 0xd0, 0xbd, 0xb3, 0x98, 0xf8, 0x31, 0x1c, 0x7f, 0x0d, 0x3e, 0x05, 0x9f, 0xaf, 0x82, 0x6b,
	0x3f, 0x3c, 0x1f, 0x8f, 0xc9, 0xf5, 0xec, 0x7b, 0xe8, 0xd9, 0x07, 0xb8, 0x05, 0x96, 0x1f, 0x6e,
	0x5e, 0xdb, 0xc8, 0xa0, 0x37, 0x76, 0x6d, 0xf0, 0xfe, 0xe1, 0xfa, 0xea, 0x22, 0x1d, 0x68, 0x9a,
	0x22, 0xf6, 0x01, 0xfe, 0x0f, 0x5a, 0xde, 0x6c, 0xe2, 0x91, 0xc0, 0x9b, 0xd9, 0x48, 0xb1, 0x2b,
	0x9f, 0x78, 0x17, 0xde, 0x74, 0x6a, 0xd7, 0x6e, 0x1a, 0xfa, 0x5f, 0x7f, 0xf5, 0x2f, 0x00, 0x00,
	0xff, 0xff, 0x08, 0x38, 0xb8, 0xe8, 0x02, 0x04, 0x00, 0x00,
}
