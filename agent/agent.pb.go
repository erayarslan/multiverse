// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.35.1
// 	protoc        v5.28.3
// source: agent/agent.proto

package agent

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

type ListRequest struct {
	state         protoimpl.MessageState
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *ListRequest) Reset() {
	*x = ListRequest{}
	mi := &file_agent_agent_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *ListRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListRequest) ProtoMessage() {}

func (x *ListRequest) ProtoReflect() protoreflect.Message {
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

// Deprecated: Use ListRequest.ProtoReflect.Descriptor instead.
func (*ListRequest) Descriptor() ([]byte, []int) {
	return file_agent_agent_proto_rawDescGZIP(), []int{0}
}

type ListReply struct {
	state         protoimpl.MessageState
	unknownFields protoimpl.UnknownFields
	Names         []string `protobuf:"bytes,1,rep,name=names,proto3" json:"names,omitempty"`
	sizeCache     protoimpl.SizeCache
}

func (x *ListReply) Reset() {
	*x = ListReply{}
	mi := &file_agent_agent_proto_msgTypes[1]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *ListReply) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListReply) ProtoMessage() {}

func (x *ListReply) ProtoReflect() protoreflect.Message {
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

// Deprecated: Use ListReply.ProtoReflect.Descriptor instead.
func (*ListReply) Descriptor() ([]byte, []int) {
	return file_agent_agent_proto_rawDescGZIP(), []int{1}
}

func (x *ListReply) GetNames() []string {
	if x != nil {
		return x.Names
	}
	return nil
}

type ShellRequest struct {
	state         protoimpl.MessageState
	InstanceName  string `protobuf:"bytes,1,opt,name=instance_name,json=instanceName,proto3" json:"instance_name,omitempty"`
	unknownFields protoimpl.UnknownFields
	InBuffer      []byte `protobuf:"bytes,4,opt,name=in_buffer,json=inBuffer,proto3" json:"in_buffer,omitempty"`
	Height        int64  `protobuf:"varint,2,opt,name=height,proto3" json:"height,omitempty"`
	Width         int64  `protobuf:"varint,3,opt,name=width,proto3" json:"width,omitempty"`
	sizeCache     protoimpl.SizeCache
}

func (x *ShellRequest) Reset() {
	*x = ShellRequest{}
	mi := &file_agent_agent_proto_msgTypes[2]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *ShellRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ShellRequest) ProtoMessage() {}

func (x *ShellRequest) ProtoReflect() protoreflect.Message {
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

// Deprecated: Use ShellRequest.ProtoReflect.Descriptor instead.
func (*ShellRequest) Descriptor() ([]byte, []int) {
	return file_agent_agent_proto_rawDescGZIP(), []int{2}
}

func (x *ShellRequest) GetInstanceName() string {
	if x != nil {
		return x.InstanceName
	}
	return ""
}

func (x *ShellRequest) GetHeight() int64 {
	if x != nil {
		return x.Height
	}
	return 0
}

func (x *ShellRequest) GetWidth() int64 {
	if x != nil {
		return x.Width
	}
	return 0
}

func (x *ShellRequest) GetInBuffer() []byte {
	if x != nil {
		return x.InBuffer
	}
	return nil
}

type ShellReply struct {
	state         protoimpl.MessageState
	unknownFields protoimpl.UnknownFields
	OutBuffer     []byte `protobuf:"bytes,1,opt,name=out_buffer,json=outBuffer,proto3" json:"out_buffer,omitempty"`
	ErrBuffer     []byte `protobuf:"bytes,2,opt,name=err_buffer,json=errBuffer,proto3" json:"err_buffer,omitempty"`
	sizeCache     protoimpl.SizeCache
}

func (x *ShellReply) Reset() {
	*x = ShellReply{}
	mi := &file_agent_agent_proto_msgTypes[3]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *ShellReply) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ShellReply) ProtoMessage() {}

func (x *ShellReply) ProtoReflect() protoreflect.Message {
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

// Deprecated: Use ShellReply.ProtoReflect.Descriptor instead.
func (*ShellReply) Descriptor() ([]byte, []int) {
	return file_agent_agent_proto_rawDescGZIP(), []int{3}
}

func (x *ShellReply) GetOutBuffer() []byte {
	if x != nil {
		return x.OutBuffer
	}
	return nil
}

func (x *ShellReply) GetErrBuffer() []byte {
	if x != nil {
		return x.ErrBuffer
	}
	return nil
}

var File_agent_agent_proto protoreflect.FileDescriptor

var file_agent_agent_proto_rawDesc = []byte{
	0x0a, 0x11, 0x61, 0x67, 0x65, 0x6e, 0x74, 0x2f, 0x61, 0x67, 0x65, 0x6e, 0x74, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x12, 0x05, 0x61, 0x67, 0x65, 0x6e, 0x74, 0x22, 0x0d, 0x0a, 0x0b, 0x4c, 0x69,
	0x73, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x22, 0x21, 0x0a, 0x09, 0x4c, 0x69, 0x73,
	0x74, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x12, 0x14, 0x0a, 0x05, 0x6e, 0x61, 0x6d, 0x65, 0x73, 0x18,
	0x01, 0x20, 0x03, 0x28, 0x09, 0x52, 0x05, 0x6e, 0x61, 0x6d, 0x65, 0x73, 0x22, 0x7e, 0x0a, 0x0c,
	0x53, 0x68, 0x65, 0x6c, 0x6c, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x23, 0x0a, 0x0d,
	0x69, 0x6e, 0x73, 0x74, 0x61, 0x6e, 0x63, 0x65, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x0c, 0x69, 0x6e, 0x73, 0x74, 0x61, 0x6e, 0x63, 0x65, 0x4e, 0x61, 0x6d,
	0x65, 0x12, 0x16, 0x0a, 0x06, 0x68, 0x65, 0x69, 0x67, 0x68, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x03, 0x52, 0x06, 0x68, 0x65, 0x69, 0x67, 0x68, 0x74, 0x12, 0x14, 0x0a, 0x05, 0x77, 0x69, 0x64,
	0x74, 0x68, 0x18, 0x03, 0x20, 0x01, 0x28, 0x03, 0x52, 0x05, 0x77, 0x69, 0x64, 0x74, 0x68, 0x12,
	0x1b, 0x0a, 0x09, 0x69, 0x6e, 0x5f, 0x62, 0x75, 0x66, 0x66, 0x65, 0x72, 0x18, 0x04, 0x20, 0x01,
	0x28, 0x0c, 0x52, 0x08, 0x69, 0x6e, 0x42, 0x75, 0x66, 0x66, 0x65, 0x72, 0x22, 0x4a, 0x0a, 0x0a,
	0x53, 0x68, 0x65, 0x6c, 0x6c, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x12, 0x1d, 0x0a, 0x0a, 0x6f, 0x75,
	0x74, 0x5f, 0x62, 0x75, 0x66, 0x66, 0x65, 0x72, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x09,
	0x6f, 0x75, 0x74, 0x42, 0x75, 0x66, 0x66, 0x65, 0x72, 0x12, 0x1d, 0x0a, 0x0a, 0x65, 0x72, 0x72,
	0x5f, 0x62, 0x75, 0x66, 0x66, 0x65, 0x72, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x09, 0x65,
	0x72, 0x72, 0x42, 0x75, 0x66, 0x66, 0x65, 0x72, 0x32, 0x6c, 0x0a, 0x03, 0x52, 0x70, 0x63, 0x12,
	0x2e, 0x0a, 0x04, 0x6c, 0x69, 0x73, 0x74, 0x12, 0x12, 0x2e, 0x61, 0x67, 0x65, 0x6e, 0x74, 0x2e,
	0x4c, 0x69, 0x73, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x10, 0x2e, 0x61, 0x67,
	0x65, 0x6e, 0x74, 0x2e, 0x4c, 0x69, 0x73, 0x74, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x22, 0x00, 0x12,
	0x35, 0x0a, 0x05, 0x73, 0x68, 0x65, 0x6c, 0x6c, 0x12, 0x13, 0x2e, 0x61, 0x67, 0x65, 0x6e, 0x74,
	0x2e, 0x53, 0x68, 0x65, 0x6c, 0x6c, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x11, 0x2e,
	0x61, 0x67, 0x65, 0x6e, 0x74, 0x2e, 0x53, 0x68, 0x65, 0x6c, 0x6c, 0x52, 0x65, 0x70, 0x6c, 0x79,
	0x22, 0x00, 0x28, 0x01, 0x30, 0x01, 0x42, 0x09, 0x5a, 0x07, 0x2e, 0x2f, 0x61, 0x67, 0x65, 0x6e,
	0x74, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
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

var file_agent_agent_proto_msgTypes = make([]protoimpl.MessageInfo, 4)
var file_agent_agent_proto_goTypes = []any{
	(*ListRequest)(nil),  // 0: agent.ListRequest
	(*ListReply)(nil),    // 1: agent.ListReply
	(*ShellRequest)(nil), // 2: agent.ShellRequest
	(*ShellReply)(nil),   // 3: agent.ShellReply
}
var file_agent_agent_proto_depIdxs = []int32{
	0, // 0: agent.Rpc.list:input_type -> agent.ListRequest
	2, // 1: agent.Rpc.shell:input_type -> agent.ShellRequest
	1, // 2: agent.Rpc.list:output_type -> agent.ListReply
	3, // 3: agent.Rpc.shell:output_type -> agent.ShellReply
	2, // [2:4] is the sub-list for method output_type
	0, // [0:2] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
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
			NumMessages:   4,
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