// Copyright (c) 2022-2023, Zander Schwid & Co. LLC.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.27.1
// 	protoc        v3.13.0
// source: raft_service.proto

package raftpb

import (
	_ "github.com/grpc-ecosystem/grpc-gateway/v2/protoc-gen-openapiv2/options"
	_ "google.golang.org/genproto/googleapis/api/annotations"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	emptypb "google.golang.org/protobuf/types/known/emptypb"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type RaftNode struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	NodeId   string `protobuf:"bytes,1,opt,name=node_id,json=nodeId,proto3" json:"node_id,omitempty"`
	NodeAddr string `protobuf:"bytes,2,opt,name=node_addr,json=nodeAddr,proto3" json:"node_addr,omitempty"`
}

func (x *RaftNode) Reset() {
	*x = RaftNode{}
	if protoimpl.UnsafeEnabled {
		mi := &file_raft_service_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RaftNode) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RaftNode) ProtoMessage() {}

func (x *RaftNode) ProtoReflect() protoreflect.Message {
	mi := &file_raft_service_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RaftNode.ProtoReflect.Descriptor instead.
func (*RaftNode) Descriptor() ([]byte, []int) {
	return file_raft_service_proto_rawDescGZIP(), []int{0}
}

func (x *RaftNode) GetNodeId() string {
	if x != nil {
		return x.NodeId
	}
	return ""
}

func (x *RaftNode) GetNodeAddr() string {
	if x != nil {
		return x.NodeAddr
	}
	return ""
}

type Content struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Content []byte `protobuf:"bytes,1,opt,name=content,proto3" json:"content,omitempty"`
}

func (x *Content) Reset() {
	*x = Content{}
	if protoimpl.UnsafeEnabled {
		mi := &file_raft_service_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Content) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Content) ProtoMessage() {}

func (x *Content) ProtoReflect() protoreflect.Message {
	mi := &file_raft_service_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Content.ProtoReflect.Descriptor instead.
func (*Content) Descriptor() ([]byte, []int) {
	return file_raft_service_proto_rawDescGZIP(), []int{1}
}

func (x *Content) GetContent() []byte {
	if x != nil {
		return x.Content
	}
	return nil
}

type RaftServer struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	NodeId   string `protobuf:"bytes,1,opt,name=node_id,json=nodeId,proto3" json:"node_id,omitempty"`
	RaftAddr string `protobuf:"bytes,2,opt,name=raft_addr,json=raftAddr,proto3" json:"raft_addr,omitempty"`
	Suffrage string `protobuf:"bytes,3,opt,name=suffrage,proto3" json:"suffrage,omitempty"`
	ApiAddr  string `protobuf:"bytes,4,opt,name=api_addr,json=apiAddr,proto3" json:"api_addr,omitempty"`
}

func (x *RaftServer) Reset() {
	*x = RaftServer{}
	if protoimpl.UnsafeEnabled {
		mi := &file_raft_service_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RaftServer) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RaftServer) ProtoMessage() {}

func (x *RaftServer) ProtoReflect() protoreflect.Message {
	mi := &file_raft_service_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RaftServer.ProtoReflect.Descriptor instead.
func (*RaftServer) Descriptor() ([]byte, []int) {
	return file_raft_service_proto_rawDescGZIP(), []int{2}
}

func (x *RaftServer) GetNodeId() string {
	if x != nil {
		return x.NodeId
	}
	return ""
}

func (x *RaftServer) GetRaftAddr() string {
	if x != nil {
		return x.RaftAddr
	}
	return ""
}

func (x *RaftServer) GetSuffrage() string {
	if x != nil {
		return x.Suffrage
	}
	return ""
}

func (x *RaftServer) GetApiAddr() string {
	if x != nil {
		return x.ApiAddr
	}
	return ""
}

type RaftConfiguration struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	State      string        `protobuf:"bytes,1,opt,name=state,proto3" json:"state,omitempty"`
	LastIndex  uint64        `protobuf:"varint,2,opt,name=last_index,json=lastIndex,proto3" json:"last_index,omitempty"`
	ServerList []*RaftServer `protobuf:"bytes,3,rep,name=server_list,json=serverList,proto3" json:"server_list,omitempty"`
}

func (x *RaftConfiguration) Reset() {
	*x = RaftConfiguration{}
	if protoimpl.UnsafeEnabled {
		mi := &file_raft_service_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RaftConfiguration) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RaftConfiguration) ProtoMessage() {}

func (x *RaftConfiguration) ProtoReflect() protoreflect.Message {
	mi := &file_raft_service_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RaftConfiguration.ProtoReflect.Descriptor instead.
func (*RaftConfiguration) Descriptor() ([]byte, []int) {
	return file_raft_service_proto_rawDescGZIP(), []int{3}
}

func (x *RaftConfiguration) GetState() string {
	if x != nil {
		return x.State
	}
	return ""
}

func (x *RaftConfiguration) GetLastIndex() uint64 {
	if x != nil {
		return x.LastIndex
	}
	return 0
}

func (x *RaftConfiguration) GetServerList() []*RaftServer {
	if x != nil {
		return x.ServerList
	}
	return nil
}

type Status struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Updated bool    `protobuf:"varint,1,opt,name=updated,proto3" json:"updated,omitempty"`
	Elapsed float64 `protobuf:"fixed64,2,opt,name=elapsed,proto3" json:"elapsed,omitempty"` // operation cost in seconds
}

func (x *Status) Reset() {
	*x = Status{}
	if protoimpl.UnsafeEnabled {
		mi := &file_raft_service_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Status) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Status) ProtoMessage() {}

func (x *Status) ProtoReflect() protoreflect.Message {
	mi := &file_raft_service_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Status.ProtoReflect.Descriptor instead.
func (*Status) Descriptor() ([]byte, []int) {
	return file_raft_service_proto_rawDescGZIP(), []int{4}
}

func (x *Status) GetUpdated() bool {
	if x != nil {
		return x.Updated
	}
	return false
}

func (x *Status) GetElapsed() float64 {
	if x != nil {
		return x.Elapsed
	}
	return 0
}

type Command struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Payload []byte `protobuf:"bytes,1,opt,name=payload,proto3" json:"payload,omitempty"`
}

func (x *Command) Reset() {
	*x = Command{}
	if protoimpl.UnsafeEnabled {
		mi := &file_raft_service_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Command) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Command) ProtoMessage() {}

func (x *Command) ProtoReflect() protoreflect.Message {
	mi := &file_raft_service_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Command.ProtoReflect.Descriptor instead.
func (*Command) Descriptor() ([]byte, []int) {
	return file_raft_service_proto_rawDescGZIP(), []int{5}
}

func (x *Command) GetPayload() []byte {
	if x != nil {
		return x.Payload
	}
	return nil
}

var File_raft_service_proto protoreflect.FileDescriptor

var file_raft_service_proto_rawDesc = []byte{
	0x0a, 0x12, 0x72, 0x61, 0x66, 0x74, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x12, 0x04, 0x72, 0x61, 0x66, 0x74, 0x1a, 0x1c, 0x67, 0x6f, 0x6f, 0x67,
	0x6c, 0x65, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f,
	0x6e, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x63,
	0x2d, 0x67, 0x65, 0x6e, 0x2d, 0x6f, 0x70, 0x65, 0x6e, 0x61, 0x70, 0x69, 0x76, 0x32, 0x2f, 0x6f,
	0x70, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2f, 0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f,
	0x6e, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1b, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65,
	0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x65, 0x6d, 0x70, 0x74, 0x79, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x40, 0x0a, 0x08, 0x52, 0x61, 0x66, 0x74, 0x4e, 0x6f, 0x64,
	0x65, 0x12, 0x17, 0x0a, 0x07, 0x6e, 0x6f, 0x64, 0x65, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x06, 0x6e, 0x6f, 0x64, 0x65, 0x49, 0x64, 0x12, 0x1b, 0x0a, 0x09, 0x6e, 0x6f,
	0x64, 0x65, 0x5f, 0x61, 0x64, 0x64, 0x72, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x6e,
	0x6f, 0x64, 0x65, 0x41, 0x64, 0x64, 0x72, 0x22, 0x23, 0x0a, 0x07, 0x43, 0x6f, 0x6e, 0x74, 0x65,
	0x6e, 0x74, 0x12, 0x18, 0x0a, 0x07, 0x63, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x0c, 0x52, 0x07, 0x63, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x22, 0x79, 0x0a, 0x0a,
	0x52, 0x61, 0x66, 0x74, 0x53, 0x65, 0x72, 0x76, 0x65, 0x72, 0x12, 0x17, 0x0a, 0x07, 0x6e, 0x6f,
	0x64, 0x65, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x6e, 0x6f, 0x64,
	0x65, 0x49, 0x64, 0x12, 0x1b, 0x0a, 0x09, 0x72, 0x61, 0x66, 0x74, 0x5f, 0x61, 0x64, 0x64, 0x72,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x72, 0x61, 0x66, 0x74, 0x41, 0x64, 0x64, 0x72,
	0x12, 0x1a, 0x0a, 0x08, 0x73, 0x75, 0x66, 0x66, 0x72, 0x61, 0x67, 0x65, 0x18, 0x03, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x08, 0x73, 0x75, 0x66, 0x66, 0x72, 0x61, 0x67, 0x65, 0x12, 0x19, 0x0a, 0x08,
	0x61, 0x70, 0x69, 0x5f, 0x61, 0x64, 0x64, 0x72, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07,
	0x61, 0x70, 0x69, 0x41, 0x64, 0x64, 0x72, 0x22, 0x7b, 0x0a, 0x11, 0x52, 0x61, 0x66, 0x74, 0x43,
	0x6f, 0x6e, 0x66, 0x69, 0x67, 0x75, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x14, 0x0a, 0x05,
	0x73, 0x74, 0x61, 0x74, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x73, 0x74, 0x61,
	0x74, 0x65, 0x12, 0x1d, 0x0a, 0x0a, 0x6c, 0x61, 0x73, 0x74, 0x5f, 0x69, 0x6e, 0x64, 0x65, 0x78,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x04, 0x52, 0x09, 0x6c, 0x61, 0x73, 0x74, 0x49, 0x6e, 0x64, 0x65,
	0x78, 0x12, 0x31, 0x0a, 0x0b, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x5f, 0x6c, 0x69, 0x73, 0x74,
	0x18, 0x03, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x10, 0x2e, 0x72, 0x61, 0x66, 0x74, 0x2e, 0x52, 0x61,
	0x66, 0x74, 0x53, 0x65, 0x72, 0x76, 0x65, 0x72, 0x52, 0x0a, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72,
	0x4c, 0x69, 0x73, 0x74, 0x22, 0x3c, 0x0a, 0x06, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0x18,
	0x0a, 0x07, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x08, 0x52,
	0x07, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x64, 0x12, 0x18, 0x0a, 0x07, 0x65, 0x6c, 0x61, 0x70,
	0x73, 0x65, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x01, 0x52, 0x07, 0x65, 0x6c, 0x61, 0x70, 0x73,
	0x65, 0x64, 0x22, 0x23, 0x0a, 0x07, 0x43, 0x6f, 0x6d, 0x6d, 0x61, 0x6e, 0x64, 0x12, 0x18, 0x0a,
	0x07, 0x70, 0x61, 0x79, 0x6c, 0x6f, 0x61, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x07,
	0x70, 0x61, 0x79, 0x6c, 0x6f, 0x61, 0x64, 0x32, 0xb1, 0x03, 0x0a, 0x0b, 0x52, 0x61, 0x66, 0x74,
	0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x5b, 0x0a, 0x09, 0x42, 0x6f, 0x6f, 0x74, 0x73,
	0x74, 0x72, 0x61, 0x70, 0x12, 0x16, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x1a, 0x16, 0x2e, 0x67,
	0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x45,
	0x6d, 0x70, 0x74, 0x79, 0x22, 0x1e, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x18, 0x1a, 0x16, 0x2f, 0x61,
	0x70, 0x69, 0x2f, 0x76, 0x31, 0x2f, 0x72, 0x61, 0x66, 0x74, 0x2f, 0x62, 0x6f, 0x6f, 0x74, 0x73,
	0x74, 0x72, 0x61, 0x70, 0x12, 0x49, 0x0a, 0x04, 0x4a, 0x6f, 0x69, 0x6e, 0x12, 0x0e, 0x2e, 0x72,
	0x61, 0x66, 0x74, 0x2e, 0x52, 0x61, 0x66, 0x74, 0x4e, 0x6f, 0x64, 0x65, 0x1a, 0x16, 0x2e, 0x67,
	0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x45,
	0x6d, 0x70, 0x74, 0x79, 0x22, 0x19, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x13, 0x1a, 0x11, 0x2f, 0x61,
	0x70, 0x69, 0x2f, 0x76, 0x31, 0x2f, 0x72, 0x61, 0x66, 0x74, 0x2f, 0x6a, 0x6f, 0x69, 0x6e, 0x12,
	0x60, 0x0a, 0x10, 0x47, 0x65, 0x74, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x75, 0x72, 0x61, 0x74,
	0x69, 0x6f, 0x6e, 0x12, 0x16, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x1a, 0x17, 0x2e, 0x72, 0x61,
	0x66, 0x74, 0x2e, 0x52, 0x61, 0x66, 0x74, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x75, 0x72, 0x61,
	0x74, 0x69, 0x6f, 0x6e, 0x22, 0x1b, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x15, 0x12, 0x13, 0x2f, 0x61,
	0x70, 0x69, 0x2f, 0x76, 0x31, 0x2f, 0x72, 0x61, 0x66, 0x74, 0x2f, 0x63, 0x6f, 0x6e, 0x66, 0x69,
	0x67, 0x12, 0x47, 0x0a, 0x0c, 0x41, 0x70, 0x70, 0x6c, 0x79, 0x43, 0x6f, 0x6d, 0x6d, 0x61, 0x6e,
	0x64, 0x12, 0x0d, 0x2e, 0x72, 0x61, 0x66, 0x74, 0x2e, 0x43, 0x6f, 0x6d, 0x6d, 0x61, 0x6e, 0x64,
	0x1a, 0x0c, 0x2e, 0x72, 0x61, 0x66, 0x74, 0x2e, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x22, 0x1a,
	0x82, 0xd3, 0xe4, 0x93, 0x02, 0x14, 0x12, 0x12, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x76, 0x31, 0x2f,
	0x72, 0x61, 0x66, 0x74, 0x2f, 0x61, 0x70, 0x70, 0x6c, 0x79, 0x12, 0x4f, 0x0a, 0x07, 0x52, 0x65,
	0x63, 0x6f, 0x76, 0x65, 0x72, 0x12, 0x0d, 0x2e, 0x72, 0x61, 0x66, 0x74, 0x2e, 0x43, 0x6f, 0x6e,
	0x74, 0x65, 0x6e, 0x74, 0x1a, 0x16, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x22, 0x1b, 0x82, 0xd3,
	0xe4, 0x93, 0x02, 0x15, 0x1a, 0x13, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x76, 0x31, 0x2f, 0x66, 0x73,
	0x6d, 0x2f, 0x72, 0x65, 0x63, 0x6f, 0x76, 0x65, 0x72, 0x28, 0x01, 0x42, 0xea, 0x01, 0x0a, 0x0f,
	0x63, 0x6f, 0x6d, 0x2e, 0x63, 0x6f, 0x64, 0x65, 0x61, 0x6c, 0x6c, 0x65, 0x72, 0x67, 0x79, 0x42,
	0x0a, 0x52, 0x61, 0x66, 0x74, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x73, 0x50, 0x01, 0x5a, 0x09, 0x2e,
	0x2e, 0x2f, 0x72, 0x61, 0x66, 0x74, 0x70, 0x62, 0xa2, 0x02, 0x02, 0x52, 0x50, 0x92, 0x41, 0xb7,
	0x01, 0x12, 0x59, 0x0a, 0x0b, 0x52, 0x61, 0x66, 0x74, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65,
	0x22, 0x45, 0x0a, 0x0b, 0x52, 0x61, 0x66, 0x74, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12,
	0x23, 0x68, 0x74, 0x74, 0x70, 0x73, 0x3a, 0x2f, 0x2f, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e,
	0x63, 0x6f, 0x6d, 0x2f, 0x63, 0x6f, 0x64, 0x65, 0x61, 0x6c, 0x6c, 0x65, 0x72, 0x67, 0x79, 0x2f,
	0x72, 0x61, 0x66, 0x74, 0x1a, 0x11, 0x7a, 0x61, 0x6e, 0x64, 0x65, 0x72, 0x40, 0x73, 0x63, 0x68,
	0x77, 0x69, 0x64, 0x2e, 0x63, 0x6f, 0x6d, 0x32, 0x03, 0x31, 0x2e, 0x30, 0x2a, 0x02, 0x01, 0x02,
	0x32, 0x10, 0x61, 0x70, 0x70, 0x6c, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2f, 0x6a, 0x73,
	0x6f, 0x6e, 0x32, 0x18, 0x61, 0x70, 0x70, 0x6c, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2f,
	0x6f, 0x63, 0x74, 0x65, 0x74, 0x2d, 0x73, 0x74, 0x72, 0x65, 0x61, 0x6d, 0x3a, 0x10, 0x61, 0x70,
	0x70, 0x6c, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2f, 0x6a, 0x73, 0x6f, 0x6e, 0x3a, 0x18,
	0x61, 0x70, 0x70, 0x6c, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2f, 0x6f, 0x63, 0x74, 0x65,
	0x74, 0x2d, 0x73, 0x74, 0x72, 0x65, 0x61, 0x6d, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_raft_service_proto_rawDescOnce sync.Once
	file_raft_service_proto_rawDescData = file_raft_service_proto_rawDesc
)

func file_raft_service_proto_rawDescGZIP() []byte {
	file_raft_service_proto_rawDescOnce.Do(func() {
		file_raft_service_proto_rawDescData = protoimpl.X.CompressGZIP(file_raft_service_proto_rawDescData)
	})
	return file_raft_service_proto_rawDescData
}

var file_raft_service_proto_msgTypes = make([]protoimpl.MessageInfo, 6)
var file_raft_service_proto_goTypes = []interface{}{
	(*RaftNode)(nil),          // 0: raft.RaftNode
	(*Content)(nil),           // 1: raft.Content
	(*RaftServer)(nil),        // 2: raft.RaftServer
	(*RaftConfiguration)(nil), // 3: raft.RaftConfiguration
	(*Status)(nil),            // 4: raft.Status
	(*Command)(nil),           // 5: raft.Command
	(*emptypb.Empty)(nil),     // 6: google.protobuf.Empty
}
var file_raft_service_proto_depIdxs = []int32{
	2, // 0: raft.RaftConfiguration.server_list:type_name -> raft.RaftServer
	6, // 1: raft.RaftService.Bootstrap:input_type -> google.protobuf.Empty
	0, // 2: raft.RaftService.Join:input_type -> raft.RaftNode
	6, // 3: raft.RaftService.GetConfiguration:input_type -> google.protobuf.Empty
	5, // 4: raft.RaftService.ApplyCommand:input_type -> raft.Command
	1, // 5: raft.RaftService.Recover:input_type -> raft.Content
	6, // 6: raft.RaftService.Bootstrap:output_type -> google.protobuf.Empty
	6, // 7: raft.RaftService.Join:output_type -> google.protobuf.Empty
	3, // 8: raft.RaftService.GetConfiguration:output_type -> raft.RaftConfiguration
	4, // 9: raft.RaftService.ApplyCommand:output_type -> raft.Status
	6, // 10: raft.RaftService.Recover:output_type -> google.protobuf.Empty
	6, // [6:11] is the sub-list for method output_type
	1, // [1:6] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_raft_service_proto_init() }
func file_raft_service_proto_init() {
	if File_raft_service_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_raft_service_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RaftNode); i {
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
		file_raft_service_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Content); i {
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
		file_raft_service_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RaftServer); i {
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
		file_raft_service_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RaftConfiguration); i {
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
		file_raft_service_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Status); i {
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
		file_raft_service_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Command); i {
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
			RawDescriptor: file_raft_service_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   6,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_raft_service_proto_goTypes,
		DependencyIndexes: file_raft_service_proto_depIdxs,
		MessageInfos:      file_raft_service_proto_msgTypes,
	}.Build()
	File_raft_service_proto = out.File
	file_raft_service_proto_rawDesc = nil
	file_raft_service_proto_goTypes = nil
	file_raft_service_proto_depIdxs = nil
}
