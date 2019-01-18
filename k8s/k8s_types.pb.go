// Code generated by protoc-gen-go. DO NOT EDIT.
// source: k8s/k8s_types.proto

package k8s

import (
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	any "github.com/golang/protobuf/ptypes/any"
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

type NodeAddressType int32

const (
	NodeAddressType_Hostname   NodeAddressType = 0
	NodeAddressType_ExternalIP NodeAddressType = 1
	NodeAddressType_InternalIP NodeAddressType = 3
)

var NodeAddressType_name = map[int32]string{
	0: "Hostname",
	1: "ExternalIP",
	3: "InternalIP",
}

var NodeAddressType_value = map[string]int32{
	"Hostname":   0,
	"ExternalIP": 1,
	"InternalIP": 3,
}

func (x NodeAddressType) String() string {
	return proto.EnumName(NodeAddressType_name, int32(x))
}

func (NodeAddressType) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_d854beef6e32a34d, []int{0}
}

type NodeList struct {
	ApiVersion           string    `protobuf:"bytes,1,opt,name=apiVersion,proto3" json:"apiVersion,omitempty"`
	Items                []*Node   `protobuf:"bytes,2,rep,name=items,proto3" json:"items,omitempty"`
	Metadata             *ListMeta `protobuf:"bytes,3,opt,name=metadata,proto3" json:"metadata,omitempty"`
	XXX_NoUnkeyedLiteral struct{}  `json:"-"`
	XXX_unrecognized     []byte    `json:"-"`
	XXX_sizecache        int32     `json:"-"`
}

func (m *NodeList) Reset()         { *m = NodeList{} }
func (m *NodeList) String() string { return proto.CompactTextString(m) }
func (*NodeList) ProtoMessage()    {}
func (*NodeList) Descriptor() ([]byte, []int) {
	return fileDescriptor_d854beef6e32a34d, []int{0}
}

func (m *NodeList) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_NodeList.Unmarshal(m, b)
}
func (m *NodeList) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_NodeList.Marshal(b, m, deterministic)
}
func (m *NodeList) XXX_Merge(src proto.Message) {
	xxx_messageInfo_NodeList.Merge(m, src)
}
func (m *NodeList) XXX_Size() int {
	return xxx_messageInfo_NodeList.Size(m)
}
func (m *NodeList) XXX_DiscardUnknown() {
	xxx_messageInfo_NodeList.DiscardUnknown(m)
}

var xxx_messageInfo_NodeList proto.InternalMessageInfo

func (m *NodeList) GetApiVersion() string {
	if m != nil {
		return m.ApiVersion
	}
	return ""
}

func (m *NodeList) GetItems() []*Node {
	if m != nil {
		return m.Items
	}
	return nil
}

func (m *NodeList) GetMetadata() *ListMeta {
	if m != nil {
		return m.Metadata
	}
	return nil
}

type ListMeta struct {
	Continue             string   `protobuf:"bytes,1,opt,name=continue,proto3" json:"continue,omitempty"`
	ResourceVersion      string   `protobuf:"bytes,2,opt,name=resourceVersion,proto3" json:"resourceVersion,omitempty"`
	SelfLink             string   `protobuf:"bytes,3,opt,name=selfLink,proto3" json:"selfLink,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ListMeta) Reset()         { *m = ListMeta{} }
func (m *ListMeta) String() string { return proto.CompactTextString(m) }
func (*ListMeta) ProtoMessage()    {}
func (*ListMeta) Descriptor() ([]byte, []int) {
	return fileDescriptor_d854beef6e32a34d, []int{1}
}

func (m *ListMeta) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ListMeta.Unmarshal(m, b)
}
func (m *ListMeta) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ListMeta.Marshal(b, m, deterministic)
}
func (m *ListMeta) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ListMeta.Merge(m, src)
}
func (m *ListMeta) XXX_Size() int {
	return xxx_messageInfo_ListMeta.Size(m)
}
func (m *ListMeta) XXX_DiscardUnknown() {
	xxx_messageInfo_ListMeta.DiscardUnknown(m)
}

var xxx_messageInfo_ListMeta proto.InternalMessageInfo

func (m *ListMeta) GetContinue() string {
	if m != nil {
		return m.Continue
	}
	return ""
}

func (m *ListMeta) GetResourceVersion() string {
	if m != nil {
		return m.ResourceVersion
	}
	return ""
}

func (m *ListMeta) GetSelfLink() string {
	if m != nil {
		return m.SelfLink
	}
	return ""
}

type Node struct {
	ApiVersion           string      `protobuf:"bytes,1,opt,name=apiVersion,proto3" json:"apiVersion,omitempty"`
	Kind                 string      `protobuf:"bytes,2,opt,name=kind,proto3" json:"kind,omitempty"`
	Metadata             *ObjectMeta `protobuf:"bytes,3,opt,name=metadata,proto3" json:"metadata,omitempty"`
	Spec                 *NodeSpec   `protobuf:"bytes,4,opt,name=spec,proto3" json:"spec,omitempty"`
	Status               *NodeStatus `protobuf:"bytes,5,opt,name=status,proto3" json:"status,omitempty"`
	XXX_NoUnkeyedLiteral struct{}    `json:"-"`
	XXX_unrecognized     []byte      `json:"-"`
	XXX_sizecache        int32       `json:"-"`
}

func (m *Node) Reset()         { *m = Node{} }
func (m *Node) String() string { return proto.CompactTextString(m) }
func (*Node) ProtoMessage()    {}
func (*Node) Descriptor() ([]byte, []int) {
	return fileDescriptor_d854beef6e32a34d, []int{2}
}

func (m *Node) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Node.Unmarshal(m, b)
}
func (m *Node) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Node.Marshal(b, m, deterministic)
}
func (m *Node) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Node.Merge(m, src)
}
func (m *Node) XXX_Size() int {
	return xxx_messageInfo_Node.Size(m)
}
func (m *Node) XXX_DiscardUnknown() {
	xxx_messageInfo_Node.DiscardUnknown(m)
}

var xxx_messageInfo_Node proto.InternalMessageInfo

func (m *Node) GetApiVersion() string {
	if m != nil {
		return m.ApiVersion
	}
	return ""
}

func (m *Node) GetKind() string {
	if m != nil {
		return m.Kind
	}
	return ""
}

func (m *Node) GetMetadata() *ObjectMeta {
	if m != nil {
		return m.Metadata
	}
	return nil
}

func (m *Node) GetSpec() *NodeSpec {
	if m != nil {
		return m.Spec
	}
	return nil
}

func (m *Node) GetStatus() *NodeStatus {
	if m != nil {
		return m.Status
	}
	return nil
}

type NodeSpec struct {
	ExternalID           string   `protobuf:"bytes,1,opt,name=externalID,proto3" json:"externalID,omitempty"`
	PodCIDR              string   `protobuf:"bytes,2,opt,name=podCIDR,proto3" json:"podCIDR,omitempty"`
	ProviderID           string   `protobuf:"bytes,3,opt,name=providerID,proto3" json:"providerID,omitempty"`
	Unschedulable        bool     `protobuf:"varint,4,opt,name=unschedulable,proto3" json:"unschedulable,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *NodeSpec) Reset()         { *m = NodeSpec{} }
func (m *NodeSpec) String() string { return proto.CompactTextString(m) }
func (*NodeSpec) ProtoMessage()    {}
func (*NodeSpec) Descriptor() ([]byte, []int) {
	return fileDescriptor_d854beef6e32a34d, []int{3}
}

func (m *NodeSpec) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_NodeSpec.Unmarshal(m, b)
}
func (m *NodeSpec) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_NodeSpec.Marshal(b, m, deterministic)
}
func (m *NodeSpec) XXX_Merge(src proto.Message) {
	xxx_messageInfo_NodeSpec.Merge(m, src)
}
func (m *NodeSpec) XXX_Size() int {
	return xxx_messageInfo_NodeSpec.Size(m)
}
func (m *NodeSpec) XXX_DiscardUnknown() {
	xxx_messageInfo_NodeSpec.DiscardUnknown(m)
}

var xxx_messageInfo_NodeSpec proto.InternalMessageInfo

func (m *NodeSpec) GetExternalID() string {
	if m != nil {
		return m.ExternalID
	}
	return ""
}

func (m *NodeSpec) GetPodCIDR() string {
	if m != nil {
		return m.PodCIDR
	}
	return ""
}

func (m *NodeSpec) GetProviderID() string {
	if m != nil {
		return m.ProviderID
	}
	return ""
}

func (m *NodeSpec) GetUnschedulable() bool {
	if m != nil {
		return m.Unschedulable
	}
	return false
}

type NodeStatus struct {
	Addresses   []*NodeAddress   `protobuf:"bytes,1,rep,name=addresses,proto3" json:"addresses,omitempty"`
	Allocatable *any.Any         `protobuf:"bytes,2,opt,name=allocatable,proto3" json:"allocatable,omitempty"`
	Capacity    *any.Any         `protobuf:"bytes,3,opt,name=capacity,proto3" json:"capacity,omitempty"`
	Conditions  []*NodeCondition `protobuf:"bytes,4,rep,name=conditions,proto3" json:"conditions,omitempty"`
	// daemonEndpoints=5
	Images               []*ContainerImage `protobuf:"bytes,6,rep,name=images,proto3" json:"images,omitempty"`
	NodeInfo             *NodeSystemInfo   `protobuf:"bytes,7,opt,name=nodeInfo,proto3" json:"nodeInfo,omitempty"`
	Phase                string            `protobuf:"bytes,8,opt,name=phase,proto3" json:"phase,omitempty"`
	VolumesAttached      []*AttachedVolume `protobuf:"bytes,9,rep,name=volumesAttached,proto3" json:"volumesAttached,omitempty"`
	VolumesInUse         []string          `protobuf:"bytes,10,rep,name=volumesInUse,proto3" json:"volumesInUse,omitempty"`
	XXX_NoUnkeyedLiteral struct{}          `json:"-"`
	XXX_unrecognized     []byte            `json:"-"`
	XXX_sizecache        int32             `json:"-"`
}

func (m *NodeStatus) Reset()         { *m = NodeStatus{} }
func (m *NodeStatus) String() string { return proto.CompactTextString(m) }
func (*NodeStatus) ProtoMessage()    {}
func (*NodeStatus) Descriptor() ([]byte, []int) {
	return fileDescriptor_d854beef6e32a34d, []int{4}
}

func (m *NodeStatus) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_NodeStatus.Unmarshal(m, b)
}
func (m *NodeStatus) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_NodeStatus.Marshal(b, m, deterministic)
}
func (m *NodeStatus) XXX_Merge(src proto.Message) {
	xxx_messageInfo_NodeStatus.Merge(m, src)
}
func (m *NodeStatus) XXX_Size() int {
	return xxx_messageInfo_NodeStatus.Size(m)
}
func (m *NodeStatus) XXX_DiscardUnknown() {
	xxx_messageInfo_NodeStatus.DiscardUnknown(m)
}

var xxx_messageInfo_NodeStatus proto.InternalMessageInfo

func (m *NodeStatus) GetAddresses() []*NodeAddress {
	if m != nil {
		return m.Addresses
	}
	return nil
}

func (m *NodeStatus) GetAllocatable() *any.Any {
	if m != nil {
		return m.Allocatable
	}
	return nil
}

func (m *NodeStatus) GetCapacity() *any.Any {
	if m != nil {
		return m.Capacity
	}
	return nil
}

func (m *NodeStatus) GetConditions() []*NodeCondition {
	if m != nil {
		return m.Conditions
	}
	return nil
}

func (m *NodeStatus) GetImages() []*ContainerImage {
	if m != nil {
		return m.Images
	}
	return nil
}

func (m *NodeStatus) GetNodeInfo() *NodeSystemInfo {
	if m != nil {
		return m.NodeInfo
	}
	return nil
}

func (m *NodeStatus) GetPhase() string {
	if m != nil {
		return m.Phase
	}
	return ""
}

func (m *NodeStatus) GetVolumesAttached() []*AttachedVolume {
	if m != nil {
		return m.VolumesAttached
	}
	return nil
}

func (m *NodeStatus) GetVolumesInUse() []string {
	if m != nil {
		return m.VolumesInUse
	}
	return nil
}

type NodeCondition struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *NodeCondition) Reset()         { *m = NodeCondition{} }
func (m *NodeCondition) String() string { return proto.CompactTextString(m) }
func (*NodeCondition) ProtoMessage()    {}
func (*NodeCondition) Descriptor() ([]byte, []int) {
	return fileDescriptor_d854beef6e32a34d, []int{5}
}

func (m *NodeCondition) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_NodeCondition.Unmarshal(m, b)
}
func (m *NodeCondition) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_NodeCondition.Marshal(b, m, deterministic)
}
func (m *NodeCondition) XXX_Merge(src proto.Message) {
	xxx_messageInfo_NodeCondition.Merge(m, src)
}
func (m *NodeCondition) XXX_Size() int {
	return xxx_messageInfo_NodeCondition.Size(m)
}
func (m *NodeCondition) XXX_DiscardUnknown() {
	xxx_messageInfo_NodeCondition.DiscardUnknown(m)
}

var xxx_messageInfo_NodeCondition proto.InternalMessageInfo

type ContainerImage struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ContainerImage) Reset()         { *m = ContainerImage{} }
func (m *ContainerImage) String() string { return proto.CompactTextString(m) }
func (*ContainerImage) ProtoMessage()    {}
func (*ContainerImage) Descriptor() ([]byte, []int) {
	return fileDescriptor_d854beef6e32a34d, []int{6}
}

func (m *ContainerImage) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ContainerImage.Unmarshal(m, b)
}
func (m *ContainerImage) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ContainerImage.Marshal(b, m, deterministic)
}
func (m *ContainerImage) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ContainerImage.Merge(m, src)
}
func (m *ContainerImage) XXX_Size() int {
	return xxx_messageInfo_ContainerImage.Size(m)
}
func (m *ContainerImage) XXX_DiscardUnknown() {
	xxx_messageInfo_ContainerImage.DiscardUnknown(m)
}

var xxx_messageInfo_ContainerImage proto.InternalMessageInfo

type NodeSystemInfo struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *NodeSystemInfo) Reset()         { *m = NodeSystemInfo{} }
func (m *NodeSystemInfo) String() string { return proto.CompactTextString(m) }
func (*NodeSystemInfo) ProtoMessage()    {}
func (*NodeSystemInfo) Descriptor() ([]byte, []int) {
	return fileDescriptor_d854beef6e32a34d, []int{7}
}

func (m *NodeSystemInfo) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_NodeSystemInfo.Unmarshal(m, b)
}
func (m *NodeSystemInfo) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_NodeSystemInfo.Marshal(b, m, deterministic)
}
func (m *NodeSystemInfo) XXX_Merge(src proto.Message) {
	xxx_messageInfo_NodeSystemInfo.Merge(m, src)
}
func (m *NodeSystemInfo) XXX_Size() int {
	return xxx_messageInfo_NodeSystemInfo.Size(m)
}
func (m *NodeSystemInfo) XXX_DiscardUnknown() {
	xxx_messageInfo_NodeSystemInfo.DiscardUnknown(m)
}

var xxx_messageInfo_NodeSystemInfo proto.InternalMessageInfo

type AttachedVolume struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *AttachedVolume) Reset()         { *m = AttachedVolume{} }
func (m *AttachedVolume) String() string { return proto.CompactTextString(m) }
func (*AttachedVolume) ProtoMessage()    {}
func (*AttachedVolume) Descriptor() ([]byte, []int) {
	return fileDescriptor_d854beef6e32a34d, []int{8}
}

func (m *AttachedVolume) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_AttachedVolume.Unmarshal(m, b)
}
func (m *AttachedVolume) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_AttachedVolume.Marshal(b, m, deterministic)
}
func (m *AttachedVolume) XXX_Merge(src proto.Message) {
	xxx_messageInfo_AttachedVolume.Merge(m, src)
}
func (m *AttachedVolume) XXX_Size() int {
	return xxx_messageInfo_AttachedVolume.Size(m)
}
func (m *AttachedVolume) XXX_DiscardUnknown() {
	xxx_messageInfo_AttachedVolume.DiscardUnknown(m)
}

var xxx_messageInfo_AttachedVolume proto.InternalMessageInfo

type ObjectMeta struct {
	Annotations                *any.Any   `protobuf:"bytes,1,opt,name=annotations,proto3" json:"annotations,omitempty"`
	ClusterName                string     `protobuf:"bytes,2,opt,name=clusterName,proto3" json:"clusterName,omitempty"`
	CreationTimestamp          string     `protobuf:"bytes,3,opt,name=creationTimestamp,proto3" json:"creationTimestamp,omitempty"`
	DeletionGracePeriodSeconds int32      `protobuf:"varint,4,opt,name=deletionGracePeriodSeconds,proto3" json:"deletionGracePeriodSeconds,omitempty"`
	DeletionTimestamp          string     `protobuf:"bytes,5,opt,name=deletionTimestamp,proto3" json:"deletionTimestamp,omitempty"`
	Finalizers                 []string   `protobuf:"bytes,6,rep,name=finalizers,proto3" json:"finalizers,omitempty"`
	GenerateName               string     `protobuf:"bytes,7,opt,name=generateName,proto3" json:"generateName,omitempty"`
	Generation                 int32      `protobuf:"varint,8,opt,name=generation,proto3" json:"generation,omitempty"`
	Initializers               *any.Any   `protobuf:"bytes,9,opt,name=initializers,proto3" json:"initializers,omitempty"`
	Labels                     *any.Any   `protobuf:"bytes,10,opt,name=labels,proto3" json:"labels,omitempty"`
	Name                       string     `protobuf:"bytes,11,opt,name=name,proto3" json:"name,omitempty"`
	Namespace                  string     `protobuf:"bytes,12,opt,name=namespace,proto3" json:"namespace,omitempty"`
	OwnerReferences            []*any.Any `protobuf:"bytes,13,rep,name=ownerReferences,proto3" json:"ownerReferences,omitempty"`
	ResourceVersion            string     `protobuf:"bytes,14,opt,name=resourceVersion,proto3" json:"resourceVersion,omitempty"`
	SelfLink                   string     `protobuf:"bytes,15,opt,name=selfLink,proto3" json:"selfLink,omitempty"`
	Uid                        string     `protobuf:"bytes,16,opt,name=uid,proto3" json:"uid,omitempty"`
	XXX_NoUnkeyedLiteral       struct{}   `json:"-"`
	XXX_unrecognized           []byte     `json:"-"`
	XXX_sizecache              int32      `json:"-"`
}

func (m *ObjectMeta) Reset()         { *m = ObjectMeta{} }
func (m *ObjectMeta) String() string { return proto.CompactTextString(m) }
func (*ObjectMeta) ProtoMessage()    {}
func (*ObjectMeta) Descriptor() ([]byte, []int) {
	return fileDescriptor_d854beef6e32a34d, []int{9}
}

func (m *ObjectMeta) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ObjectMeta.Unmarshal(m, b)
}
func (m *ObjectMeta) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ObjectMeta.Marshal(b, m, deterministic)
}
func (m *ObjectMeta) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ObjectMeta.Merge(m, src)
}
func (m *ObjectMeta) XXX_Size() int {
	return xxx_messageInfo_ObjectMeta.Size(m)
}
func (m *ObjectMeta) XXX_DiscardUnknown() {
	xxx_messageInfo_ObjectMeta.DiscardUnknown(m)
}

var xxx_messageInfo_ObjectMeta proto.InternalMessageInfo

func (m *ObjectMeta) GetAnnotations() *any.Any {
	if m != nil {
		return m.Annotations
	}
	return nil
}

func (m *ObjectMeta) GetClusterName() string {
	if m != nil {
		return m.ClusterName
	}
	return ""
}

func (m *ObjectMeta) GetCreationTimestamp() string {
	if m != nil {
		return m.CreationTimestamp
	}
	return ""
}

func (m *ObjectMeta) GetDeletionGracePeriodSeconds() int32 {
	if m != nil {
		return m.DeletionGracePeriodSeconds
	}
	return 0
}

func (m *ObjectMeta) GetDeletionTimestamp() string {
	if m != nil {
		return m.DeletionTimestamp
	}
	return ""
}

func (m *ObjectMeta) GetFinalizers() []string {
	if m != nil {
		return m.Finalizers
	}
	return nil
}

func (m *ObjectMeta) GetGenerateName() string {
	if m != nil {
		return m.GenerateName
	}
	return ""
}

func (m *ObjectMeta) GetGeneration() int32 {
	if m != nil {
		return m.Generation
	}
	return 0
}

func (m *ObjectMeta) GetInitializers() *any.Any {
	if m != nil {
		return m.Initializers
	}
	return nil
}

func (m *ObjectMeta) GetLabels() *any.Any {
	if m != nil {
		return m.Labels
	}
	return nil
}

func (m *ObjectMeta) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *ObjectMeta) GetNamespace() string {
	if m != nil {
		return m.Namespace
	}
	return ""
}

func (m *ObjectMeta) GetOwnerReferences() []*any.Any {
	if m != nil {
		return m.OwnerReferences
	}
	return nil
}

func (m *ObjectMeta) GetResourceVersion() string {
	if m != nil {
		return m.ResourceVersion
	}
	return ""
}

func (m *ObjectMeta) GetSelfLink() string {
	if m != nil {
		return m.SelfLink
	}
	return ""
}

func (m *ObjectMeta) GetUid() string {
	if m != nil {
		return m.Uid
	}
	return ""
}

type NodeAddress struct {
	Address              string          `protobuf:"bytes,1,opt,name=address,proto3" json:"address,omitempty"`
	Type                 NodeAddressType `protobuf:"varint,2,opt,name=type,proto3,enum=k8s.NodeAddressType" json:"type,omitempty"`
	XXX_NoUnkeyedLiteral struct{}        `json:"-"`
	XXX_unrecognized     []byte          `json:"-"`
	XXX_sizecache        int32           `json:"-"`
}

func (m *NodeAddress) Reset()         { *m = NodeAddress{} }
func (m *NodeAddress) String() string { return proto.CompactTextString(m) }
func (*NodeAddress) ProtoMessage()    {}
func (*NodeAddress) Descriptor() ([]byte, []int) {
	return fileDescriptor_d854beef6e32a34d, []int{10}
}

func (m *NodeAddress) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_NodeAddress.Unmarshal(m, b)
}
func (m *NodeAddress) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_NodeAddress.Marshal(b, m, deterministic)
}
func (m *NodeAddress) XXX_Merge(src proto.Message) {
	xxx_messageInfo_NodeAddress.Merge(m, src)
}
func (m *NodeAddress) XXX_Size() int {
	return xxx_messageInfo_NodeAddress.Size(m)
}
func (m *NodeAddress) XXX_DiscardUnknown() {
	xxx_messageInfo_NodeAddress.DiscardUnknown(m)
}

var xxx_messageInfo_NodeAddress proto.InternalMessageInfo

func (m *NodeAddress) GetAddress() string {
	if m != nil {
		return m.Address
	}
	return ""
}

func (m *NodeAddress) GetType() NodeAddressType {
	if m != nil {
		return m.Type
	}
	return NodeAddressType_Hostname
}

func init() {
	proto.RegisterEnum("k8s.NodeAddressType", NodeAddressType_name, NodeAddressType_value)
	proto.RegisterType((*NodeList)(nil), "k8s.NodeList")
	proto.RegisterType((*ListMeta)(nil), "k8s.ListMeta")
	proto.RegisterType((*Node)(nil), "k8s.Node")
	proto.RegisterType((*NodeSpec)(nil), "k8s.NodeSpec")
	proto.RegisterType((*NodeStatus)(nil), "k8s.NodeStatus")
	proto.RegisterType((*NodeCondition)(nil), "k8s.NodeCondition")
	proto.RegisterType((*ContainerImage)(nil), "k8s.ContainerImage")
	proto.RegisterType((*NodeSystemInfo)(nil), "k8s.NodeSystemInfo")
	proto.RegisterType((*AttachedVolume)(nil), "k8s.AttachedVolume")
	proto.RegisterType((*ObjectMeta)(nil), "k8s.ObjectMeta")
	proto.RegisterType((*NodeAddress)(nil), "k8s.NodeAddress")
}

func init() { proto.RegisterFile("k8s/k8s_types.proto", fileDescriptor_d854beef6e32a34d) }

var fileDescriptor_d854beef6e32a34d = []byte{
	// 852 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x84, 0x55, 0x6d, 0x6f, 0xe3, 0x44,
	0x10, 0x26, 0xcd, 0x4b, 0xed, 0x49, 0xdb, 0x84, 0x6d, 0x3f, 0x2c, 0x15, 0x82, 0x60, 0x21, 0x11,
	0xb8, 0x53, 0x8a, 0x8a, 0x84, 0xf2, 0x85, 0x43, 0x55, 0x8b, 0x20, 0xd2, 0x71, 0x1c, 0x7b, 0xc7,
	0x7d, 0x45, 0x1b, 0x7b, 0xd2, 0x5b, 0xe2, 0xec, 0x5a, 0xbb, 0xeb, 0x42, 0xf8, 0x09, 0xfc, 0x02,
	0xfe, 0x06, 0xe2, 0x0f, 0x9e, 0x76, 0x6d, 0xc7, 0x76, 0x5a, 0xf5, 0x3e, 0xd5, 0xf3, 0x3c, 0xcf,
	0xee, 0x3c, 0x33, 0x3b, 0x9d, 0xc0, 0xe9, 0x7a, 0x6e, 0x2e, 0xd6, 0x73, 0xf3, 0xbb, 0xdd, 0x66,
	0x68, 0x66, 0x99, 0x56, 0x56, 0x91, 0xee, 0x7a, 0x6e, 0xce, 0x3f, 0xba, 0x55, 0xea, 0x36, 0xc5,
	0x0b, 0x0f, 0x2d, 0xf3, 0xd5, 0x05, 0x97, 0xdb, 0x82, 0x8f, 0xee, 0x20, 0x78, 0xa1, 0x12, 0x7c,
	0x2e, 0x8c, 0x25, 0x9f, 0x00, 0xf0, 0x4c, 0xbc, 0x41, 0x6d, 0x84, 0x92, 0xb4, 0x33, 0xe9, 0x4c,
	0x43, 0xd6, 0x40, 0xc8, 0xa7, 0xd0, 0x17, 0x16, 0x37, 0x86, 0x1e, 0x4c, 0xba, 0xd3, 0xe1, 0x65,
	0x38, 0x5b, 0xcf, 0xcd, 0xcc, 0x9d, 0x66, 0x05, 0x4e, 0xbe, 0x84, 0x60, 0x83, 0x96, 0x27, 0xdc,
	0x72, 0xda, 0x9d, 0x74, 0xa6, 0xc3, 0xcb, 0x63, 0xaf, 0x71, 0xb7, 0xff, 0x8c, 0x96, 0xb3, 0x1d,
	0x1d, 0xa5, 0x10, 0x54, 0x28, 0x39, 0x87, 0x20, 0x56, 0xd2, 0x0a, 0x99, 0x63, 0x99, 0x75, 0x17,
	0x93, 0x29, 0x8c, 0x34, 0x1a, 0x95, 0xeb, 0x18, 0x2b, 0x63, 0x07, 0x5e, 0xb2, 0x0f, 0xbb, 0x5b,
	0x0c, 0xa6, 0xab, 0xe7, 0x42, 0xae, 0x7d, 0xf2, 0x90, 0xed, 0xe2, 0xe8, 0xff, 0x0e, 0xf4, 0x9c,
	0xd1, 0xf7, 0x96, 0x48, 0xa0, 0xb7, 0x16, 0x32, 0x29, 0x73, 0xf8, 0x6f, 0xf2, 0xe4, 0x5e, 0x55,
	0x23, 0x5f, 0xd5, 0x2f, 0xcb, 0x3f, 0x30, 0xde, 0xab, 0x8b, 0x7c, 0x06, 0x3d, 0x93, 0x61, 0x4c,
	0x7b, 0x8d, 0xf2, 0x5d, 0xe6, 0x57, 0x19, 0xc6, 0xcc, 0x53, 0xe4, 0x0b, 0x18, 0x18, 0xcb, 0x6d,
	0x6e, 0x68, 0xbf, 0x71, 0x9b, 0x17, 0x79, 0x98, 0x95, 0x74, 0xf4, 0x4f, 0xa7, 0x78, 0x1c, 0x77,
	0xd6, 0x39, 0xc7, 0xbf, 0x2c, 0x6a, 0xc9, 0xd3, 0xc5, 0x4d, 0xe5, 0xbc, 0x46, 0x08, 0x85, 0xc3,
	0x4c, 0x25, 0xd7, 0x8b, 0x1b, 0x56, 0x9a, 0xaf, 0x42, 0x77, 0x32, 0xd3, 0xea, 0x4e, 0x24, 0xa8,
	0x17, 0x37, 0x65, 0x6b, 0x1a, 0x08, 0xf9, 0x1c, 0x8e, 0x73, 0x69, 0xe2, 0xb7, 0x98, 0xe4, 0x29,
	0x5f, 0xa6, 0xe8, 0xbd, 0x07, 0xac, 0x0d, 0x46, 0xff, 0x75, 0x01, 0x6a, 0x8f, 0x64, 0x06, 0x21,
	0x4f, 0x12, 0x8d, 0xc6, 0xa0, 0xa1, 0x1d, 0x3f, 0x0f, 0xe3, 0x5d, 0x1d, 0x57, 0x05, 0xc3, 0x6a,
	0x09, 0xf9, 0x16, 0x86, 0x3c, 0x4d, 0x55, 0xcc, 0xad, 0x4f, 0x71, 0xe0, 0x2b, 0x3f, 0x9b, 0x15,
	0x83, 0x39, 0xab, 0x06, 0x73, 0x76, 0x25, 0xb7, 0xac, 0x29, 0x24, 0x5f, 0x43, 0x10, 0xf3, 0x8c,
	0xc7, 0xc2, 0x6e, 0xcb, 0xe6, 0x3f, 0x7c, 0x68, 0xa7, 0x22, 0x97, 0x00, 0xb1, 0x92, 0x89, 0xb0,
	0x42, 0x49, 0x43, 0x7b, 0xde, 0x1a, 0xd9, 0x59, 0xbb, 0xae, 0x28, 0xd6, 0x50, 0x91, 0x27, 0x30,
	0x10, 0x1b, 0x7e, 0x8b, 0x86, 0x0e, 0xbc, 0xfe, 0xd4, 0xeb, 0xaf, 0x95, 0xb4, 0x5c, 0x48, 0xd4,
	0x0b, 0xc7, 0xb1, 0x52, 0x42, 0x2e, 0x20, 0x90, 0x2a, 0xc1, 0x85, 0x5c, 0x29, 0x7a, 0xe8, 0x2d,
	0x9d, 0xd6, 0x2f, 0xb8, 0x35, 0x16, 0x37, 0x8e, 0x62, 0x3b, 0x11, 0x39, 0x83, 0x7e, 0xf6, 0x96,
	0x1b, 0xa4, 0x81, 0xef, 0x7d, 0x11, 0x90, 0xef, 0x60, 0x74, 0xa7, 0xd2, 0x7c, 0x83, 0xe6, 0xca,
	0x5a, 0xee, 0x3a, 0x4d, 0xc3, 0x46, 0xf2, 0x0a, 0x7c, 0xe3, 0x35, 0x6c, 0x5f, 0x4b, 0x22, 0x38,
	0x2a, 0xa1, 0x85, 0xfc, 0xcd, 0x20, 0x85, 0x49, 0x77, 0x1a, 0xb2, 0x16, 0x16, 0x8d, 0xe0, 0xb8,
	0x55, 0x73, 0x34, 0x86, 0x93, 0x76, 0x51, 0x0e, 0x69, 0xfb, 0x76, 0x48, 0x3b, 0x77, 0xf4, 0x6f,
	0x1f, 0xa0, 0x1e, 0x76, 0xff, 0x94, 0x52, 0x2a, 0xcb, 0x8b, 0x0e, 0x77, 0x1e, 0x7d, 0xca, 0x5a,
	0x48, 0x26, 0x30, 0x8c, 0xd3, 0xdc, 0x58, 0xd4, 0x2f, 0xf8, 0x06, 0xcb, 0x29, 0x6d, 0x42, 0xe4,
	0x29, 0x7c, 0x18, 0x6b, 0xf4, 0xf2, 0xd7, 0x62, 0x83, 0xc6, 0xf2, 0x4d, 0x56, 0x0e, 0xec, 0x7d,
	0x82, 0x3c, 0x83, 0xf3, 0x04, 0x53, 0x74, 0xe0, 0x8f, 0x9a, 0xc7, 0xf8, 0x12, 0xb5, 0x50, 0xc9,
	0x2b, 0x74, 0xef, 0x6a, 0xfc, 0x10, 0xf7, 0xd9, 0x23, 0x0a, 0x97, 0xad, 0x62, 0xeb, 0x6c, 0xfd,
	0x22, 0xdb, 0x3d, 0xc2, 0xfd, 0x17, 0xad, 0x84, 0xe4, 0xa9, 0xf8, 0x1b, 0x75, 0x31, 0x26, 0x21,
	0x6b, 0x20, 0xee, 0x3d, 0x6e, 0x51, 0xa2, 0xe6, 0x16, 0x7d, 0x79, 0x87, 0xfe, 0xa2, 0x16, 0xe6,
	0xee, 0x28, 0x63, 0xb7, 0x7d, 0x02, 0xef, 0xb0, 0x81, 0x90, 0x39, 0x1c, 0x09, 0x29, 0xac, 0xa8,
	0xb2, 0x84, 0x8f, 0xb4, 0xb6, 0xa5, 0x24, 0x4f, 0x61, 0x90, 0xf2, 0x25, 0xa6, 0x86, 0xc2, 0x23,
	0x67, 0x4a, 0x8d, 0xdb, 0x72, 0xd2, 0x79, 0x1c, 0x16, 0x5b, 0xce, 0x7d, 0x93, 0x8f, 0x21, 0x74,
	0x7f, 0x4d, 0xc6, 0x63, 0xa4, 0x47, 0x9e, 0xa8, 0x01, 0xf2, 0x0c, 0x46, 0xea, 0x4f, 0x89, 0x9a,
	0xe1, 0x0a, 0x35, 0xca, 0x18, 0x0d, 0x3d, 0xf6, 0xc3, 0xfa, 0x70, 0xa2, 0x7d, 0xf1, 0x43, 0x6b,
	0xfc, 0xe4, 0xfd, 0x6b, 0x7c, 0xd4, 0x5e, 0xe3, 0x64, 0x0c, 0xdd, 0x5c, 0x24, 0x74, 0xec, 0x61,
	0xf7, 0x19, 0xfd, 0x0a, 0xc3, 0xc6, 0xc2, 0x71, 0x4b, 0xb0, 0x5c, 0x39, 0xe5, 0x86, 0xac, 0x42,
	0x32, 0x85, 0x9e, 0xfb, 0x59, 0xf4, 0x53, 0x77, 0x72, 0x79, 0xb6, 0xbf, 0xaa, 0x5e, 0x6f, 0x33,
	0x64, 0x5e, 0xf1, 0xd5, 0xf7, 0x30, 0xda, 0x23, 0xc8, 0x11, 0x04, 0x3f, 0x29, 0x63, 0x5d, 0x3b,
	0xc6, 0x1f, 0x90, 0x13, 0x80, 0x1f, 0xaa, 0xbd, 0xfb, 0x72, 0xdc, 0x71, 0xf1, 0x42, 0xee, 0xe2,
	0xee, 0x72, 0xe0, 0x5b, 0xf1, 0xcd, 0xbb, 0x00, 0x00, 0x00, 0xff, 0xff, 0xd3, 0x12, 0x5f, 0xcb,
	0x90, 0x07, 0x00, 0x00,
}
