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
	return fileDescriptor_ae30f4e6357a72f8, []int{1}
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
	return fileDescriptor_ae30f4e6357a72f8, []int{2}
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
	return fileDescriptor_ae30f4e6357a72f8, []int{3}
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
	proto.RegisterType((*PartitionInfo)(nil), "sysinfo.PartitionInfo")
	proto.RegisterType((*HostInfo)(nil), "sysinfo.HostInfo")
	proto.RegisterType((*ProcessDetails)(nil), "sysinfo.ProcessDetails")
	proto.RegisterType((*HealthInfo)(nil), "sysinfo.HealthInfo")
}

func init() { proto.RegisterFile("sysinfo/types.proto", fileDescriptor_ae30f4e6357a72f8) }

var fileDescriptor_ae30f4e6357a72f8 = []byte{
	// 348 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x5c, 0x92, 0x41, 0x4b, 0xc3, 0x40,
	0x10, 0x85, 0xd9, 0x36, 0xad, 0x76, 0x4b, 0x15, 0x22, 0x94, 0x20, 0x1e, 0x42, 0x4e, 0x39, 0x55,
	0x54, 0x0a, 0x9e, 0xd5, 0x43, 0x0b, 0x56, 0xca, 0x2a, 0xde, 0xb7, 0xed, 0x84, 0x2e, 0x24, 0x99,
	0xb0, 0xb3, 0x2d, 0xc4, 0xbb, 0x7f, 0xc7, 0x83, 0xbf, 0x50, 0x76, 0x37, 0x36, 0xd5, 0xdb, 0xbc,
	0x6f, 0x86, 0xe1, 0xcd, 0xdb, 0xe5, 0x17, 0x54, 0x93, 0x2a, 0x33, 0xbc, 0x36, 0x75, 0x05, 0x34,
	0xa9, 0x34, 0x1a, 0x0c, 0x4f, 0x1a, 0x98, 0x7c, 0x32, 0x3e, 0x5a, 0x4a, 0x6d, 0x94, 0x51, 0x58,
	0xce, 0xcb, 0x0c, 0xc3, 0x2b, 0x3e, 0x58, 0xe0, 0xae, 0x34, 0x4b, 0x69, 0xb6, 0x11, 0x8b, 0x59,
	0x3a, 0x10, 0x2d, 0x08, 0x43, 0x1e, 0x90, 0xfa, 0x80, 0xa8, 0x13, 0xb3, 0x34, 0x10, 0xae, 0xb6,
	0x2c, 0xd3, 0x00, 0x51, 0xd7, 0x33, 0x5b, 0x87, 0x63, 0xde, 0xcf, 0xe8, 0xad, 0xae, 0x20, 0x0a,
	0xdc, 0x8a, 0x46, 0x59, 0xbe, 0x81, 0xbd, 0x5a, 0x43, 0xd4, 0xf3, 0xdc, 0xab, 0xe4, 0x9b, 0xf1,
	0xd3, 0x19, 0x92, 0x71, 0x16, 0x2e, 0x7d, 0xfd, 0x22, 0x0b, 0x68, 0x1c, 0x1c, 0xb4, 0x5d, 0x80,
	0xe4, 0x3a, 0x1d, 0xbf, 0xc0, 0x2b, 0x6b, 0x1b, 0xe9, 0x1d, 0x34, 0x29, 0x2c, 0x9d, 0x93, 0x81,
	0x68, 0x81, 0xdd, 0x58, 0xee, 0x8a, 0x47, 0xd4, 0x40, 0xce, 0xd0, 0x48, 0x1c, 0x74, 0x78, 0xcf,
	0x87, 0x99, 0xca, 0xe1, 0xb5, 0x26, 0x03, 0x05, 0x45, 0xbd, 0xb8, 0x9b, 0x0e, 0x6f, 0xc7, 0x93,
	0x26, 0xa1, 0xc9, 0x9f, 0x74, 0xc4, 0xf1, 0x68, 0xf2, 0xc0, 0xcf, 0x96, 0x1a, 0xd7, 0x40, 0xf4,
	0x04, 0x46, 0xaa, 0x9c, 0xac, 0x8b, 0xca, 0x93, 0xf9, 0xc6, 0x59, 0x0f, 0x44, 0x0b, 0x6c, 0x50,
	0x95, 0x4d, 0xd5, 0x3b, 0x77, 0x75, 0xf2, 0xc5, 0x38, 0x9f, 0x81, 0xcc, 0xcd, 0xd6, 0x9d, 0x1e,
	0xf3, 0xa1, 0x41, 0x23, 0xf3, 0x05, 0x14, 0xa8, 0xeb, 0x66, 0xc5, 0x31, 0x0a, 0x53, 0x7e, 0x2e,
	0xf7, 0x52, 0xe5, 0x72, 0x95, 0x43, 0x33, 0xe5, 0x1f, 0xe3, 0x3f, 0xb6, 0x47, 0xaf, 0xab, 0xdd,
	0x33, 0xca, 0xcd, 0x8d, 0x4b, 0x84, 0x89, 0x83, 0x3e, 0xea, 0x4d, 0x5d, 0x20, 0x6d, 0x6f, 0x6a,
	0x8f, 0xf8, 0x9d, 0x9b, 0xba, 0x67, 0x62, 0xa2, 0x05, 0xab, 0xbe, 0xfb, 0x41, 0x77, 0x3f, 0x01,
	0x00, 0x00, 0xff, 0xff, 0xbd, 0x8e, 0x65, 0x50, 0x58, 0x02, 0x00, 0x00,
}
