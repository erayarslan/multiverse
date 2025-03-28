// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.35.2
// 	protoc        v5.28.3
// source: agent/agent.proto

package agent

import (
	common "github.com/erayarslan/multiverse/common"
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

type CPU struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Total     int32 `protobuf:"varint,1,opt,name=total,proto3" json:"total,omitempty"`
	Available int32 `protobuf:"varint,2,opt,name=available,proto3" json:"available,omitempty"`
}

func (x *CPU) Reset() {
	*x = CPU{}
	mi := &file_agent_agent_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *CPU) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CPU) ProtoMessage() {}

func (x *CPU) ProtoReflect() protoreflect.Message {
	mi := &file_agent_agent_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CPU.ProtoReflect.Descriptor instead.
func (*CPU) Descriptor() ([]byte, []int) {
	return file_agent_agent_proto_rawDescGZIP(), []int{0}
}

func (x *CPU) GetTotal() int32 {
	if x != nil {
		return x.Total
	}
	return 0
}

func (x *CPU) GetAvailable() int32 {
	if x != nil {
		return x.Available
	}
	return 0
}

type Memory struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Total     uint64 `protobuf:"varint,1,opt,name=total,proto3" json:"total,omitempty"`
	Available uint64 `protobuf:"varint,2,opt,name=available,proto3" json:"available,omitempty"`
}

func (x *Memory) Reset() {
	*x = Memory{}
	mi := &file_agent_agent_proto_msgTypes[1]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *Memory) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Memory) ProtoMessage() {}

func (x *Memory) ProtoReflect() protoreflect.Message {
	mi := &file_agent_agent_proto_msgTypes[1]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Memory.ProtoReflect.Descriptor instead.
func (*Memory) Descriptor() ([]byte, []int) {
	return file_agent_agent_proto_rawDescGZIP(), []int{1}
}

func (x *Memory) GetTotal() uint64 {
	if x != nil {
		return x.Total
	}
	return 0
}

func (x *Memory) GetAvailable() uint64 {
	if x != nil {
		return x.Available
	}
	return 0
}

type Disk struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Total     uint64 `protobuf:"varint,1,opt,name=total,proto3" json:"total,omitempty"`
	Available uint64 `protobuf:"varint,2,opt,name=available,proto3" json:"available,omitempty"`
}

func (x *Disk) Reset() {
	*x = Disk{}
	mi := &file_agent_agent_proto_msgTypes[2]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *Disk) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Disk) ProtoMessage() {}

func (x *Disk) ProtoReflect() protoreflect.Message {
	mi := &file_agent_agent_proto_msgTypes[2]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Disk.ProtoReflect.Descriptor instead.
func (*Disk) Descriptor() ([]byte, []int) {
	return file_agent_agent_proto_rawDescGZIP(), []int{2}
}

func (x *Disk) GetTotal() uint64 {
	if x != nil {
		return x.Total
	}
	return 0
}

func (x *Disk) GetAvailable() uint64 {
	if x != nil {
		return x.Available
	}
	return 0
}

type Resource struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Cpu    *CPU    `protobuf:"bytes,1,opt,name=cpu,proto3" json:"cpu,omitempty"`
	Memory *Memory `protobuf:"bytes,2,opt,name=memory,proto3" json:"memory,omitempty"`
	Disk   *Disk   `protobuf:"bytes,3,opt,name=disk,proto3" json:"disk,omitempty"`
}

func (x *Resource) Reset() {
	*x = Resource{}
	mi := &file_agent_agent_proto_msgTypes[3]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *Resource) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Resource) ProtoMessage() {}

func (x *Resource) ProtoReflect() protoreflect.Message {
	mi := &file_agent_agent_proto_msgTypes[3]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Resource.ProtoReflect.Descriptor instead.
func (*Resource) Descriptor() ([]byte, []int) {
	return file_agent_agent_proto_rawDescGZIP(), []int{3}
}

func (x *Resource) GetCpu() *CPU {
	if x != nil {
		return x.Cpu
	}
	return nil
}

func (x *Resource) GetMemory() *Memory {
	if x != nil {
		return x.Memory
	}
	return nil
}

func (x *Resource) GetDisk() *Disk {
	if x != nil {
		return x.Disk
	}
	return nil
}

type GetInfoRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *GetInfoRequest) Reset() {
	*x = GetInfoRequest{}
	mi := &file_agent_agent_proto_msgTypes[4]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *GetInfoRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetInfoRequest) ProtoMessage() {}

func (x *GetInfoRequest) ProtoReflect() protoreflect.Message {
	mi := &file_agent_agent_proto_msgTypes[4]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetInfoRequest.ProtoReflect.Descriptor instead.
func (*GetInfoRequest) Descriptor() ([]byte, []int) {
	return file_agent_agent_proto_rawDescGZIP(), []int{4}
}

type GetInfoReply struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Resource *Resource `protobuf:"bytes,1,opt,name=resource,proto3" json:"resource,omitempty"`
}

func (x *GetInfoReply) Reset() {
	*x = GetInfoReply{}
	mi := &file_agent_agent_proto_msgTypes[5]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *GetInfoReply) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetInfoReply) ProtoMessage() {}

func (x *GetInfoReply) ProtoReflect() protoreflect.Message {
	mi := &file_agent_agent_proto_msgTypes[5]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetInfoReply.ProtoReflect.Descriptor instead.
func (*GetInfoReply) Descriptor() ([]byte, []int) {
	return file_agent_agent_proto_rawDescGZIP(), []int{5}
}

func (x *GetInfoReply) GetResource() *Resource {
	if x != nil {
		return x.Resource
	}
	return nil
}

type Instance struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Name  string   `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	State string   `protobuf:"bytes,2,opt,name=state,proto3" json:"state,omitempty"`
	Ipv4  []string `protobuf:"bytes,3,rep,name=ipv4,proto3" json:"ipv4,omitempty"`
	Image string   `protobuf:"bytes,4,opt,name=image,proto3" json:"image,omitempty"`
}

func (x *Instance) Reset() {
	*x = Instance{}
	mi := &file_agent_agent_proto_msgTypes[6]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *Instance) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Instance) ProtoMessage() {}

func (x *Instance) ProtoReflect() protoreflect.Message {
	mi := &file_agent_agent_proto_msgTypes[6]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Instance.ProtoReflect.Descriptor instead.
func (*Instance) Descriptor() ([]byte, []int) {
	return file_agent_agent_proto_rawDescGZIP(), []int{6}
}

func (x *Instance) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *Instance) GetState() string {
	if x != nil {
		return x.State
	}
	return ""
}

func (x *Instance) GetIpv4() []string {
	if x != nil {
		return x.Ipv4
	}
	return nil
}

func (x *Instance) GetImage() string {
	if x != nil {
		return x.Image
	}
	return ""
}

type GetInstancesRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *GetInstancesRequest) Reset() {
	*x = GetInstancesRequest{}
	mi := &file_agent_agent_proto_msgTypes[7]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *GetInstancesRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetInstancesRequest) ProtoMessage() {}

func (x *GetInstancesRequest) ProtoReflect() protoreflect.Message {
	mi := &file_agent_agent_proto_msgTypes[7]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetInstancesRequest.ProtoReflect.Descriptor instead.
func (*GetInstancesRequest) Descriptor() ([]byte, []int) {
	return file_agent_agent_proto_rawDescGZIP(), []int{7}
}

type GetInstancesReply struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Instances []*Instance `protobuf:"bytes,1,rep,name=instances,proto3" json:"instances,omitempty"`
}

func (x *GetInstancesReply) Reset() {
	*x = GetInstancesReply{}
	mi := &file_agent_agent_proto_msgTypes[8]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *GetInstancesReply) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetInstancesReply) ProtoMessage() {}

func (x *GetInstancesReply) ProtoReflect() protoreflect.Message {
	mi := &file_agent_agent_proto_msgTypes[8]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetInstancesReply.ProtoReflect.Descriptor instead.
func (*GetInstancesReply) Descriptor() ([]byte, []int) {
	return file_agent_agent_proto_rawDescGZIP(), []int{8}
}

func (x *GetInstancesReply) GetInstances() []*Instance {
	if x != nil {
		return x.Instances
	}
	return nil
}

var File_agent_agent_proto protoreflect.FileDescriptor

var file_agent_agent_proto_rawDesc = []byte{
	0x0a, 0x11, 0x61, 0x67, 0x65, 0x6e, 0x74, 0x2f, 0x61, 0x67, 0x65, 0x6e, 0x74, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x12, 0x05, 0x61, 0x67, 0x65, 0x6e, 0x74, 0x1a, 0x13, 0x63, 0x6f, 0x6d, 0x6d,
	0x6f, 0x6e, 0x2f, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22,
	0x39, 0x0a, 0x03, 0x43, 0x50, 0x55, 0x12, 0x14, 0x0a, 0x05, 0x74, 0x6f, 0x74, 0x61, 0x6c, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x05, 0x74, 0x6f, 0x74, 0x61, 0x6c, 0x12, 0x1c, 0x0a, 0x09,
	0x61, 0x76, 0x61, 0x69, 0x6c, 0x61, 0x62, 0x6c, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x05, 0x52,
	0x09, 0x61, 0x76, 0x61, 0x69, 0x6c, 0x61, 0x62, 0x6c, 0x65, 0x22, 0x3c, 0x0a, 0x06, 0x4d, 0x65,
	0x6d, 0x6f, 0x72, 0x79, 0x12, 0x14, 0x0a, 0x05, 0x74, 0x6f, 0x74, 0x61, 0x6c, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x04, 0x52, 0x05, 0x74, 0x6f, 0x74, 0x61, 0x6c, 0x12, 0x1c, 0x0a, 0x09, 0x61, 0x76,
	0x61, 0x69, 0x6c, 0x61, 0x62, 0x6c, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x04, 0x52, 0x09, 0x61,
	0x76, 0x61, 0x69, 0x6c, 0x61, 0x62, 0x6c, 0x65, 0x22, 0x3a, 0x0a, 0x04, 0x44, 0x69, 0x73, 0x6b,
	0x12, 0x14, 0x0a, 0x05, 0x74, 0x6f, 0x74, 0x61, 0x6c, 0x18, 0x01, 0x20, 0x01, 0x28, 0x04, 0x52,
	0x05, 0x74, 0x6f, 0x74, 0x61, 0x6c, 0x12, 0x1c, 0x0a, 0x09, 0x61, 0x76, 0x61, 0x69, 0x6c, 0x61,
	0x62, 0x6c, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x04, 0x52, 0x09, 0x61, 0x76, 0x61, 0x69, 0x6c,
	0x61, 0x62, 0x6c, 0x65, 0x22, 0x70, 0x0a, 0x08, 0x52, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65,
	0x12, 0x1c, 0x0a, 0x03, 0x63, 0x70, 0x75, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0a, 0x2e,
	0x61, 0x67, 0x65, 0x6e, 0x74, 0x2e, 0x43, 0x50, 0x55, 0x52, 0x03, 0x63, 0x70, 0x75, 0x12, 0x25,
	0x0a, 0x06, 0x6d, 0x65, 0x6d, 0x6f, 0x72, 0x79, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0d,
	0x2e, 0x61, 0x67, 0x65, 0x6e, 0x74, 0x2e, 0x4d, 0x65, 0x6d, 0x6f, 0x72, 0x79, 0x52, 0x06, 0x6d,
	0x65, 0x6d, 0x6f, 0x72, 0x79, 0x12, 0x1f, 0x0a, 0x04, 0x64, 0x69, 0x73, 0x6b, 0x18, 0x03, 0x20,
	0x01, 0x28, 0x0b, 0x32, 0x0b, 0x2e, 0x61, 0x67, 0x65, 0x6e, 0x74, 0x2e, 0x44, 0x69, 0x73, 0x6b,
	0x52, 0x04, 0x64, 0x69, 0x73, 0x6b, 0x22, 0x10, 0x0a, 0x0e, 0x47, 0x65, 0x74, 0x49, 0x6e, 0x66,
	0x6f, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x22, 0x3b, 0x0a, 0x0c, 0x47, 0x65, 0x74, 0x49,
	0x6e, 0x66, 0x6f, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x12, 0x2b, 0x0a, 0x08, 0x72, 0x65, 0x73, 0x6f,
	0x75, 0x72, 0x63, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0f, 0x2e, 0x61, 0x67, 0x65,
	0x6e, 0x74, 0x2e, 0x52, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x52, 0x08, 0x72, 0x65, 0x73,
	0x6f, 0x75, 0x72, 0x63, 0x65, 0x22, 0x5e, 0x0a, 0x08, 0x49, 0x6e, 0x73, 0x74, 0x61, 0x6e, 0x63,
	0x65, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x73, 0x74, 0x61, 0x74, 0x65, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x73, 0x74, 0x61, 0x74, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x69,
	0x70, 0x76, 0x34, 0x18, 0x03, 0x20, 0x03, 0x28, 0x09, 0x52, 0x04, 0x69, 0x70, 0x76, 0x34, 0x12,
	0x14, 0x0a, 0x05, 0x69, 0x6d, 0x61, 0x67, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05,
	0x69, 0x6d, 0x61, 0x67, 0x65, 0x22, 0x15, 0x0a, 0x13, 0x47, 0x65, 0x74, 0x49, 0x6e, 0x73, 0x74,
	0x61, 0x6e, 0x63, 0x65, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x22, 0x42, 0x0a, 0x11,
	0x47, 0x65, 0x74, 0x49, 0x6e, 0x73, 0x74, 0x61, 0x6e, 0x63, 0x65, 0x73, 0x52, 0x65, 0x70, 0x6c,
	0x79, 0x12, 0x2d, 0x0a, 0x09, 0x69, 0x6e, 0x73, 0x74, 0x61, 0x6e, 0x63, 0x65, 0x73, 0x18, 0x01,
	0x20, 0x03, 0x28, 0x0b, 0x32, 0x0f, 0x2e, 0x61, 0x67, 0x65, 0x6e, 0x74, 0x2e, 0x49, 0x6e, 0x73,
	0x74, 0x61, 0x6e, 0x63, 0x65, 0x52, 0x09, 0x69, 0x6e, 0x73, 0x74, 0x61, 0x6e, 0x63, 0x65, 0x73,
	0x32, 0xf3, 0x01, 0x0a, 0x03, 0x52, 0x70, 0x63, 0x12, 0x43, 0x0a, 0x09, 0x69, 0x6e, 0x73, 0x74,
	0x61, 0x6e, 0x63, 0x65, 0x73, 0x12, 0x1a, 0x2e, 0x61, 0x67, 0x65, 0x6e, 0x74, 0x2e, 0x47, 0x65,
	0x74, 0x49, 0x6e, 0x73, 0x74, 0x61, 0x6e, 0x63, 0x65, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x1a, 0x18, 0x2e, 0x61, 0x67, 0x65, 0x6e, 0x74, 0x2e, 0x47, 0x65, 0x74, 0x49, 0x6e, 0x73,
	0x74, 0x61, 0x6e, 0x63, 0x65, 0x73, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x22, 0x00, 0x12, 0x36, 0x0a,
	0x04, 0x69, 0x6e, 0x66, 0x6f, 0x12, 0x16, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x47,
	0x65, 0x74, 0x49, 0x6e, 0x66, 0x6f, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x14, 0x2e,
	0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x47, 0x65, 0x74, 0x49, 0x6e, 0x66, 0x6f, 0x52, 0x65,
	0x70, 0x6c, 0x79, 0x22, 0x00, 0x12, 0x37, 0x0a, 0x05, 0x73, 0x68, 0x65, 0x6c, 0x6c, 0x12, 0x14,
	0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x53, 0x68, 0x65, 0x6c, 0x6c, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x1a, 0x12, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x53, 0x68,
	0x65, 0x6c, 0x6c, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x22, 0x00, 0x28, 0x01, 0x30, 0x01, 0x12, 0x36,
	0x0a, 0x06, 0x6c, 0x61, 0x75, 0x6e, 0x63, 0x68, 0x12, 0x15, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f,
	0x6e, 0x2e, 0x4c, 0x61, 0x75, 0x6e, 0x63, 0x68, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a,
	0x13, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x4c, 0x61, 0x75, 0x6e, 0x63, 0x68, 0x52,
	0x65, 0x70, 0x6c, 0x79, 0x22, 0x00, 0x42, 0x28, 0x5a, 0x26, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62,
	0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x65, 0x72, 0x61, 0x79, 0x61, 0x72, 0x73, 0x6c, 0x61, 0x6e, 0x2f,
	0x6d, 0x75, 0x6c, 0x74, 0x69, 0x76, 0x65, 0x72, 0x73, 0x65, 0x2f, 0x61, 0x67, 0x65, 0x6e, 0x74,
	0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_agent_agent_proto_rawDescOnce sync.Once
	file_agent_agent_proto_rawDescData = file_agent_agent_proto_rawDesc
)

func file_agent_agent_proto_rawDescGZIP() []byte {
	file_agent_agent_proto_rawDescOnce.Do(func() {
		file_agent_agent_proto_rawDescData = protoimpl.X.CompressGZIP(file_agent_agent_proto_rawDescData)
	})
	return file_agent_agent_proto_rawDescData
}

var file_agent_agent_proto_msgTypes = make([]protoimpl.MessageInfo, 9)
var file_agent_agent_proto_goTypes = []any{
	(*CPU)(nil),                   // 0: agent.CPU
	(*Memory)(nil),                // 1: agent.Memory
	(*Disk)(nil),                  // 2: agent.Disk
	(*Resource)(nil),              // 3: agent.Resource
	(*GetInfoRequest)(nil),        // 4: agent.GetInfoRequest
	(*GetInfoReply)(nil),          // 5: agent.GetInfoReply
	(*Instance)(nil),              // 6: agent.Instance
	(*GetInstancesRequest)(nil),   // 7: agent.GetInstancesRequest
	(*GetInstancesReply)(nil),     // 8: agent.GetInstancesReply
	(*common.GetInfoRequest)(nil), // 9: common.GetInfoRequest
	(*common.ShellRequest)(nil),   // 10: common.ShellRequest
	(*common.LaunchRequest)(nil),  // 11: common.LaunchRequest
	(*common.GetInfoReply)(nil),   // 12: common.GetInfoReply
	(*common.ShellReply)(nil),     // 13: common.ShellReply
	(*common.LaunchReply)(nil),    // 14: common.LaunchReply
}
var file_agent_agent_proto_depIdxs = []int32{
	0,  // 0: agent.Resource.cpu:type_name -> agent.CPU
	1,  // 1: agent.Resource.memory:type_name -> agent.Memory
	2,  // 2: agent.Resource.disk:type_name -> agent.Disk
	3,  // 3: agent.GetInfoReply.resource:type_name -> agent.Resource
	6,  // 4: agent.GetInstancesReply.instances:type_name -> agent.Instance
	7,  // 5: agent.Rpc.instances:input_type -> agent.GetInstancesRequest
	9,  // 6: agent.Rpc.info:input_type -> common.GetInfoRequest
	10, // 7: agent.Rpc.shell:input_type -> common.ShellRequest
	11, // 8: agent.Rpc.launch:input_type -> common.LaunchRequest
	8,  // 9: agent.Rpc.instances:output_type -> agent.GetInstancesReply
	12, // 10: agent.Rpc.info:output_type -> common.GetInfoReply
	13, // 11: agent.Rpc.shell:output_type -> common.ShellReply
	14, // 12: agent.Rpc.launch:output_type -> common.LaunchReply
	9,  // [9:13] is the sub-list for method output_type
	5,  // [5:9] is the sub-list for method input_type
	5,  // [5:5] is the sub-list for extension type_name
	5,  // [5:5] is the sub-list for extension extendee
	0,  // [0:5] is the sub-list for field type_name
}

func init() { file_agent_agent_proto_init() }
func file_agent_agent_proto_init() {
	if File_agent_agent_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_agent_agent_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   9,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_agent_agent_proto_goTypes,
		DependencyIndexes: file_agent_agent_proto_depIdxs,
		MessageInfos:      file_agent_agent_proto_msgTypes,
	}.Build()
	File_agent_agent_proto = out.File
	file_agent_agent_proto_rawDesc = nil
	file_agent_agent_proto_goTypes = nil
	file_agent_agent_proto_depIdxs = nil
}
