// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.35.2
// 	protoc        v5.28.3
// source: api/api.proto

package api

import (
	agent "github.com/erayarslan/multiverse/agent"
	common "github.com/erayarslan/multiverse/common"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	timestamppb "google.golang.org/protobuf/types/known/timestamppb"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type Node struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Name     string                 `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	LastSync *timestamppb.Timestamp `protobuf:"bytes,2,opt,name=last_sync,json=lastSync,proto3" json:"last_sync,omitempty"`
	Ipv4     []string               `protobuf:"bytes,3,rep,name=ipv4,proto3" json:"ipv4,omitempty"`
	Resource *agent.Resource        `protobuf:"bytes,4,opt,name=resource,proto3" json:"resource,omitempty"`
}

func (x *Node) Reset() {
	*x = Node{}
	mi := &file_api_api_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *Node) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Node) ProtoMessage() {}

func (x *Node) ProtoReflect() protoreflect.Message {
	mi := &file_api_api_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Node.ProtoReflect.Descriptor instead.
func (*Node) Descriptor() ([]byte, []int) {
	return file_api_api_proto_rawDescGZIP(), []int{0}
}

func (x *Node) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *Node) GetLastSync() *timestamppb.Timestamp {
	if x != nil {
		return x.LastSync
	}
	return nil
}

func (x *Node) GetIpv4() []string {
	if x != nil {
		return x.Ipv4
	}
	return nil
}

func (x *Node) GetResource() *agent.Resource {
	if x != nil {
		return x.Resource
	}
	return nil
}

type GetNodesRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *GetNodesRequest) Reset() {
	*x = GetNodesRequest{}
	mi := &file_api_api_proto_msgTypes[1]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *GetNodesRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetNodesRequest) ProtoMessage() {}

func (x *GetNodesRequest) ProtoReflect() protoreflect.Message {
	mi := &file_api_api_proto_msgTypes[1]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetNodesRequest.ProtoReflect.Descriptor instead.
func (*GetNodesRequest) Descriptor() ([]byte, []int) {
	return file_api_api_proto_rawDescGZIP(), []int{1}
}

type GetNodesReply struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Nodes []*Node `protobuf:"bytes,1,rep,name=nodes,proto3" json:"nodes,omitempty"`
}

func (x *GetNodesReply) Reset() {
	*x = GetNodesReply{}
	mi := &file_api_api_proto_msgTypes[2]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *GetNodesReply) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetNodesReply) ProtoMessage() {}

func (x *GetNodesReply) ProtoReflect() protoreflect.Message {
	mi := &file_api_api_proto_msgTypes[2]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetNodesReply.ProtoReflect.Descriptor instead.
func (*GetNodesReply) Descriptor() ([]byte, []int) {
	return file_api_api_proto_rawDescGZIP(), []int{2}
}

func (x *GetNodesReply) GetNodes() []*Node {
	if x != nil {
		return x.Nodes
	}
	return nil
}

type Instance struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	NodeName string          `protobuf:"bytes,1,opt,name=node_name,json=nodeName,proto3" json:"node_name,omitempty"`
	Instance *agent.Instance `protobuf:"bytes,2,opt,name=instance,proto3" json:"instance,omitempty"`
}

func (x *Instance) Reset() {
	*x = Instance{}
	mi := &file_api_api_proto_msgTypes[3]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *Instance) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Instance) ProtoMessage() {}

func (x *Instance) ProtoReflect() protoreflect.Message {
	mi := &file_api_api_proto_msgTypes[3]
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
	return file_api_api_proto_rawDescGZIP(), []int{3}
}

func (x *Instance) GetNodeName() string {
	if x != nil {
		return x.NodeName
	}
	return ""
}

func (x *Instance) GetInstance() *agent.Instance {
	if x != nil {
		return x.Instance
	}
	return nil
}

type GetInstancesRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *GetInstancesRequest) Reset() {
	*x = GetInstancesRequest{}
	mi := &file_api_api_proto_msgTypes[4]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *GetInstancesRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetInstancesRequest) ProtoMessage() {}

func (x *GetInstancesRequest) ProtoReflect() protoreflect.Message {
	mi := &file_api_api_proto_msgTypes[4]
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
	return file_api_api_proto_rawDescGZIP(), []int{4}
}

type GetInstancesReply struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Instances []*Instance `protobuf:"bytes,1,rep,name=instances,proto3" json:"instances,omitempty"`
}

func (x *GetInstancesReply) Reset() {
	*x = GetInstancesReply{}
	mi := &file_api_api_proto_msgTypes[5]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *GetInstancesReply) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetInstancesReply) ProtoMessage() {}

func (x *GetInstancesReply) ProtoReflect() protoreflect.Message {
	mi := &file_api_api_proto_msgTypes[5]
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
	return file_api_api_proto_rawDescGZIP(), []int{5}
}

func (x *GetInstancesReply) GetInstances() []*Instance {
	if x != nil {
		return x.Instances
	}
	return nil
}

type GetInfoInstance struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	NodeName string                  `protobuf:"bytes,1,opt,name=node_name,json=nodeName,proto3" json:"node_name,omitempty"`
	Instance *common.GetInfoInstance `protobuf:"bytes,2,opt,name=instance,proto3" json:"instance,omitempty"`
}

func (x *GetInfoInstance) Reset() {
	*x = GetInfoInstance{}
	mi := &file_api_api_proto_msgTypes[6]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *GetInfoInstance) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetInfoInstance) ProtoMessage() {}

func (x *GetInfoInstance) ProtoReflect() protoreflect.Message {
	mi := &file_api_api_proto_msgTypes[6]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetInfoInstance.ProtoReflect.Descriptor instead.
func (*GetInfoInstance) Descriptor() ([]byte, []int) {
	return file_api_api_proto_rawDescGZIP(), []int{6}
}

func (x *GetInfoInstance) GetNodeName() string {
	if x != nil {
		return x.NodeName
	}
	return ""
}

func (x *GetInfoInstance) GetInstance() *common.GetInfoInstance {
	if x != nil {
		return x.Instance
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
	mi := &file_api_api_proto_msgTypes[7]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *GetInfoRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetInfoRequest) ProtoMessage() {}

func (x *GetInfoRequest) ProtoReflect() protoreflect.Message {
	mi := &file_api_api_proto_msgTypes[7]
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
	return file_api_api_proto_rawDescGZIP(), []int{7}
}

type GetInfoReply struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Instances []*GetInfoInstance `protobuf:"bytes,1,rep,name=instances,proto3" json:"instances,omitempty"`
}

func (x *GetInfoReply) Reset() {
	*x = GetInfoReply{}
	mi := &file_api_api_proto_msgTypes[8]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *GetInfoReply) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetInfoReply) ProtoMessage() {}

func (x *GetInfoReply) ProtoReflect() protoreflect.Message {
	mi := &file_api_api_proto_msgTypes[8]
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
	return file_api_api_proto_rawDescGZIP(), []int{8}
}

func (x *GetInfoReply) GetInstances() []*GetInfoInstance {
	if x != nil {
		return x.Instances
	}
	return nil
}

var File_api_api_proto protoreflect.FileDescriptor

var file_api_api_proto_rawDesc = []byte{
	0x0a, 0x0d, 0x61, 0x70, 0x69, 0x2f, 0x61, 0x70, 0x69, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12,
	0x03, 0x61, 0x70, 0x69, 0x1a, 0x1f, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x13, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2f, 0x63, 0x6f,
	0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x11, 0x61, 0x67, 0x65, 0x6e,
	0x74, 0x2f, 0x61, 0x67, 0x65, 0x6e, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x94, 0x01,
	0x0a, 0x04, 0x4e, 0x6f, 0x64, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x37, 0x0a, 0x09, 0x6c, 0x61,
	0x73, 0x74, 0x5f, 0x73, 0x79, 0x6e, 0x63, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e,
	0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e,
	0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x52, 0x08, 0x6c, 0x61, 0x73, 0x74, 0x53,
	0x79, 0x6e, 0x63, 0x12, 0x12, 0x0a, 0x04, 0x69, 0x70, 0x76, 0x34, 0x18, 0x03, 0x20, 0x03, 0x28,
	0x09, 0x52, 0x04, 0x69, 0x70, 0x76, 0x34, 0x12, 0x2b, 0x0a, 0x08, 0x72, 0x65, 0x73, 0x6f, 0x75,
	0x72, 0x63, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0f, 0x2e, 0x61, 0x67, 0x65, 0x6e,
	0x74, 0x2e, 0x52, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x52, 0x08, 0x72, 0x65, 0x73, 0x6f,
	0x75, 0x72, 0x63, 0x65, 0x22, 0x11, 0x0a, 0x0f, 0x47, 0x65, 0x74, 0x4e, 0x6f, 0x64, 0x65, 0x73,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x22, 0x30, 0x0a, 0x0d, 0x47, 0x65, 0x74, 0x4e, 0x6f,
	0x64, 0x65, 0x73, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x12, 0x1f, 0x0a, 0x05, 0x6e, 0x6f, 0x64, 0x65,
	0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x09, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x4e, 0x6f,
	0x64, 0x65, 0x52, 0x05, 0x6e, 0x6f, 0x64, 0x65, 0x73, 0x22, 0x54, 0x0a, 0x08, 0x49, 0x6e, 0x73,
	0x74, 0x61, 0x6e, 0x63, 0x65, 0x12, 0x1b, 0x0a, 0x09, 0x6e, 0x6f, 0x64, 0x65, 0x5f, 0x6e, 0x61,
	0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x6e, 0x6f, 0x64, 0x65, 0x4e, 0x61,
	0x6d, 0x65, 0x12, 0x2b, 0x0a, 0x08, 0x69, 0x6e, 0x73, 0x74, 0x61, 0x6e, 0x63, 0x65, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x0b, 0x32, 0x0f, 0x2e, 0x61, 0x67, 0x65, 0x6e, 0x74, 0x2e, 0x49, 0x6e, 0x73,
	0x74, 0x61, 0x6e, 0x63, 0x65, 0x52, 0x08, 0x69, 0x6e, 0x73, 0x74, 0x61, 0x6e, 0x63, 0x65, 0x22,
	0x15, 0x0a, 0x13, 0x47, 0x65, 0x74, 0x49, 0x6e, 0x73, 0x74, 0x61, 0x6e, 0x63, 0x65, 0x73, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x22, 0x40, 0x0a, 0x11, 0x47, 0x65, 0x74, 0x49, 0x6e, 0x73,
	0x74, 0x61, 0x6e, 0x63, 0x65, 0x73, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x12, 0x2b, 0x0a, 0x09, 0x69,
	0x6e, 0x73, 0x74, 0x61, 0x6e, 0x63, 0x65, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x0d,
	0x2e, 0x61, 0x70, 0x69, 0x2e, 0x49, 0x6e, 0x73, 0x74, 0x61, 0x6e, 0x63, 0x65, 0x52, 0x09, 0x69,
	0x6e, 0x73, 0x74, 0x61, 0x6e, 0x63, 0x65, 0x73, 0x22, 0x63, 0x0a, 0x0f, 0x47, 0x65, 0x74, 0x49,
	0x6e, 0x66, 0x6f, 0x49, 0x6e, 0x73, 0x74, 0x61, 0x6e, 0x63, 0x65, 0x12, 0x1b, 0x0a, 0x09, 0x6e,
	0x6f, 0x64, 0x65, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08,
	0x6e, 0x6f, 0x64, 0x65, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x33, 0x0a, 0x08, 0x69, 0x6e, 0x73, 0x74,
	0x61, 0x6e, 0x63, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x17, 0x2e, 0x63, 0x6f, 0x6d,
	0x6d, 0x6f, 0x6e, 0x2e, 0x47, 0x65, 0x74, 0x49, 0x6e, 0x66, 0x6f, 0x49, 0x6e, 0x73, 0x74, 0x61,
	0x6e, 0x63, 0x65, 0x52, 0x08, 0x69, 0x6e, 0x73, 0x74, 0x61, 0x6e, 0x63, 0x65, 0x22, 0x10, 0x0a,
	0x0e, 0x47, 0x65, 0x74, 0x49, 0x6e, 0x66, 0x6f, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x22,
	0x42, 0x0a, 0x0c, 0x47, 0x65, 0x74, 0x49, 0x6e, 0x66, 0x6f, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x12,
	0x32, 0x0a, 0x09, 0x69, 0x6e, 0x73, 0x74, 0x61, 0x6e, 0x63, 0x65, 0x73, 0x18, 0x01, 0x20, 0x03,
	0x28, 0x0b, 0x32, 0x14, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x47, 0x65, 0x74, 0x49, 0x6e, 0x66, 0x6f,
	0x49, 0x6e, 0x73, 0x74, 0x61, 0x6e, 0x63, 0x65, 0x52, 0x09, 0x69, 0x6e, 0x73, 0x74, 0x61, 0x6e,
	0x63, 0x65, 0x73, 0x32, 0x9e, 0x02, 0x0a, 0x03, 0x52, 0x70, 0x63, 0x12, 0x3f, 0x0a, 0x09, 0x69,
	0x6e, 0x73, 0x74, 0x61, 0x6e, 0x63, 0x65, 0x73, 0x12, 0x18, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x47,
	0x65, 0x74, 0x49, 0x6e, 0x73, 0x74, 0x61, 0x6e, 0x63, 0x65, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x1a, 0x16, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x47, 0x65, 0x74, 0x49, 0x6e, 0x73, 0x74,
	0x61, 0x6e, 0x63, 0x65, 0x73, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x22, 0x00, 0x12, 0x33, 0x0a, 0x05,
	0x6e, 0x6f, 0x64, 0x65, 0x73, 0x12, 0x14, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x47, 0x65, 0x74, 0x4e,
	0x6f, 0x64, 0x65, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x12, 0x2e, 0x61, 0x70,
	0x69, 0x2e, 0x47, 0x65, 0x74, 0x4e, 0x6f, 0x64, 0x65, 0x73, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x22,
	0x00, 0x12, 0x30, 0x0a, 0x04, 0x69, 0x6e, 0x66, 0x6f, 0x12, 0x13, 0x2e, 0x61, 0x70, 0x69, 0x2e,
	0x47, 0x65, 0x74, 0x49, 0x6e, 0x66, 0x6f, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x11,
	0x2e, 0x61, 0x70, 0x69, 0x2e, 0x47, 0x65, 0x74, 0x49, 0x6e, 0x66, 0x6f, 0x52, 0x65, 0x70, 0x6c,
	0x79, 0x22, 0x00, 0x12, 0x37, 0x0a, 0x05, 0x73, 0x68, 0x65, 0x6c, 0x6c, 0x12, 0x14, 0x2e, 0x63,
	0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x53, 0x68, 0x65, 0x6c, 0x6c, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x1a, 0x12, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x53, 0x68, 0x65, 0x6c,
	0x6c, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x22, 0x00, 0x28, 0x01, 0x30, 0x01, 0x12, 0x36, 0x0a, 0x06,
	0x6c, 0x61, 0x75, 0x6e, 0x63, 0x68, 0x12, 0x15, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e,
	0x4c, 0x61, 0x75, 0x6e, 0x63, 0x68, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x13, 0x2e,
	0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x4c, 0x61, 0x75, 0x6e, 0x63, 0x68, 0x52, 0x65, 0x70,
	0x6c, 0x79, 0x22, 0x00, 0x42, 0x26, 0x5a, 0x24, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63,
	0x6f, 0x6d, 0x2f, 0x65, 0x72, 0x61, 0x79, 0x61, 0x72, 0x73, 0x6c, 0x61, 0x6e, 0x2f, 0x6d, 0x75,
	0x6c, 0x74, 0x69, 0x76, 0x65, 0x72, 0x73, 0x65, 0x2f, 0x61, 0x70, 0x69, 0x62, 0x06, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_api_api_proto_rawDescOnce sync.Once
	file_api_api_proto_rawDescData = file_api_api_proto_rawDesc
)

func file_api_api_proto_rawDescGZIP() []byte {
	file_api_api_proto_rawDescOnce.Do(func() {
		file_api_api_proto_rawDescData = protoimpl.X.CompressGZIP(file_api_api_proto_rawDescData)
	})
	return file_api_api_proto_rawDescData
}

var file_api_api_proto_msgTypes = make([]protoimpl.MessageInfo, 9)
var file_api_api_proto_goTypes = []any{
	(*Node)(nil),                   // 0: api.Node
	(*GetNodesRequest)(nil),        // 1: api.GetNodesRequest
	(*GetNodesReply)(nil),          // 2: api.GetNodesReply
	(*Instance)(nil),               // 3: api.Instance
	(*GetInstancesRequest)(nil),    // 4: api.GetInstancesRequest
	(*GetInstancesReply)(nil),      // 5: api.GetInstancesReply
	(*GetInfoInstance)(nil),        // 6: api.GetInfoInstance
	(*GetInfoRequest)(nil),         // 7: api.GetInfoRequest
	(*GetInfoReply)(nil),           // 8: api.GetInfoReply
	(*timestamppb.Timestamp)(nil),  // 9: google.protobuf.Timestamp
	(*agent.Resource)(nil),         // 10: agent.Resource
	(*agent.Instance)(nil),         // 11: agent.Instance
	(*common.GetInfoInstance)(nil), // 12: common.GetInfoInstance
	(*common.ShellRequest)(nil),    // 13: common.ShellRequest
	(*common.LaunchRequest)(nil),   // 14: common.LaunchRequest
	(*common.ShellReply)(nil),      // 15: common.ShellReply
	(*common.LaunchReply)(nil),     // 16: common.LaunchReply
}
var file_api_api_proto_depIdxs = []int32{
	9,  // 0: api.Node.last_sync:type_name -> google.protobuf.Timestamp
	10, // 1: api.Node.resource:type_name -> agent.Resource
	0,  // 2: api.GetNodesReply.nodes:type_name -> api.Node
	11, // 3: api.Instance.instance:type_name -> agent.Instance
	3,  // 4: api.GetInstancesReply.instances:type_name -> api.Instance
	12, // 5: api.GetInfoInstance.instance:type_name -> common.GetInfoInstance
	6,  // 6: api.GetInfoReply.instances:type_name -> api.GetInfoInstance
	4,  // 7: api.Rpc.instances:input_type -> api.GetInstancesRequest
	1,  // 8: api.Rpc.nodes:input_type -> api.GetNodesRequest
	7,  // 9: api.Rpc.info:input_type -> api.GetInfoRequest
	13, // 10: api.Rpc.shell:input_type -> common.ShellRequest
	14, // 11: api.Rpc.launch:input_type -> common.LaunchRequest
	5,  // 12: api.Rpc.instances:output_type -> api.GetInstancesReply
	2,  // 13: api.Rpc.nodes:output_type -> api.GetNodesReply
	8,  // 14: api.Rpc.info:output_type -> api.GetInfoReply
	15, // 15: api.Rpc.shell:output_type -> common.ShellReply
	16, // 16: api.Rpc.launch:output_type -> common.LaunchReply
	12, // [12:17] is the sub-list for method output_type
	7,  // [7:12] is the sub-list for method input_type
	7,  // [7:7] is the sub-list for extension type_name
	7,  // [7:7] is the sub-list for extension extendee
	0,  // [0:7] is the sub-list for field type_name
}

func init() { file_api_api_proto_init() }
func file_api_api_proto_init() {
	if File_api_api_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_api_api_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   9,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_api_api_proto_goTypes,
		DependencyIndexes: file_api_api_proto_depIdxs,
		MessageInfos:      file_api_api_proto_msgTypes,
	}.Build()
	File_api_api_proto = out.File
	file_api_api_proto_rawDesc = nil
	file_api_api_proto_goTypes = nil
	file_api_api_proto_depIdxs = nil
}
