// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.26.0
// 	protoc        v3.15.8
// source: grpc/zookeeper/proto/zookeeper.proto

package zookeeper

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

// RegisterRequest 请求结构
type RegisterRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	NodeName string `protobuf:"bytes,1,opt,name=nodeName,proto3" json:"nodeName,omitempty"`
	HttpAddr string `protobuf:"bytes,2,opt,name=httpAddr,proto3" json:"httpAddr,omitempty"`
	GrpcAddr string `protobuf:"bytes,3,opt,name=grpcAddr,proto3" json:"grpcAddr,omitempty"`
}

func (x *RegisterRequest) Reset() {
	*x = RegisterRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_grpc_zookeeper_proto_zookeeper_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RegisterRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RegisterRequest) ProtoMessage() {}

func (x *RegisterRequest) ProtoReflect() protoreflect.Message {
	mi := &file_grpc_zookeeper_proto_zookeeper_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RegisterRequest.ProtoReflect.Descriptor instead.
func (*RegisterRequest) Descriptor() ([]byte, []int) {
	return file_grpc_zookeeper_proto_zookeeper_proto_rawDescGZIP(), []int{0}
}

func (x *RegisterRequest) GetNodeName() string {
	if x != nil {
		return x.NodeName
	}
	return ""
}

func (x *RegisterRequest) GetHttpAddr() string {
	if x != nil {
		return x.HttpAddr
	}
	return ""
}

func (x *RegisterRequest) GetGrpcAddr() string {
	if x != nil {
		return x.GrpcAddr
	}
	return ""
}

// RegisterResponse 响应结构
type RegisterResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Code string `protobuf:"bytes,1,opt,name=code,proto3" json:"code,omitempty"`
	Msg  string `protobuf:"bytes,2,opt,name=msg,proto3" json:"msg,omitempty"`
}

func (x *RegisterResponse) Reset() {
	*x = RegisterResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_grpc_zookeeper_proto_zookeeper_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RegisterResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RegisterResponse) ProtoMessage() {}

func (x *RegisterResponse) ProtoReflect() protoreflect.Message {
	mi := &file_grpc_zookeeper_proto_zookeeper_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RegisterResponse.ProtoReflect.Descriptor instead.
func (*RegisterResponse) Descriptor() ([]byte, []int) {
	return file_grpc_zookeeper_proto_zookeeper_proto_rawDescGZIP(), []int{1}
}

func (x *RegisterResponse) GetCode() string {
	if x != nil {
		return x.Code
	}
	return ""
}

func (x *RegisterResponse) GetMsg() string {
	if x != nil {
		return x.Msg
	}
	return ""
}

// HeartbeatRequest 请求结构
type HeartbeatRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	NodeName string `protobuf:"bytes,1,opt,name=nodeName,proto3" json:"nodeName,omitempty"`
	GrpcAddr string `protobuf:"bytes,2,opt,name=grpcAddr,proto3" json:"grpcAddr,omitempty"`
	HttpAddr string `protobuf:"bytes,3,opt,name=httpAddr,proto3" json:"httpAddr,omitempty"`
}

func (x *HeartbeatRequest) Reset() {
	*x = HeartbeatRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_grpc_zookeeper_proto_zookeeper_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *HeartbeatRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*HeartbeatRequest) ProtoMessage() {}

func (x *HeartbeatRequest) ProtoReflect() protoreflect.Message {
	mi := &file_grpc_zookeeper_proto_zookeeper_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use HeartbeatRequest.ProtoReflect.Descriptor instead.
func (*HeartbeatRequest) Descriptor() ([]byte, []int) {
	return file_grpc_zookeeper_proto_zookeeper_proto_rawDescGZIP(), []int{2}
}

func (x *HeartbeatRequest) GetNodeName() string {
	if x != nil {
		return x.NodeName
	}
	return ""
}

func (x *HeartbeatRequest) GetGrpcAddr() string {
	if x != nil {
		return x.GrpcAddr
	}
	return ""
}

func (x *HeartbeatRequest) GetHttpAddr() string {
	if x != nil {
		return x.HttpAddr
	}
	return ""
}

// HeartbeatResponse 响应结构
type HeartbeatResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Code string `protobuf:"bytes,1,opt,name=code,proto3" json:"code,omitempty"`
	Msg  string `protobuf:"bytes,2,opt,name=msg,proto3" json:"msg,omitempty"`
}

func (x *HeartbeatResponse) Reset() {
	*x = HeartbeatResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_grpc_zookeeper_proto_zookeeper_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *HeartbeatResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*HeartbeatResponse) ProtoMessage() {}

func (x *HeartbeatResponse) ProtoReflect() protoreflect.Message {
	mi := &file_grpc_zookeeper_proto_zookeeper_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use HeartbeatResponse.ProtoReflect.Descriptor instead.
func (*HeartbeatResponse) Descriptor() ([]byte, []int) {
	return file_grpc_zookeeper_proto_zookeeper_proto_rawDescGZIP(), []int{3}
}

func (x *HeartbeatResponse) GetCode() string {
	if x != nil {
		return x.Code
	}
	return ""
}

func (x *HeartbeatResponse) GetMsg() string {
	if x != nil {
		return x.Msg
	}
	return ""
}

type UserCheckInRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	UserId   string `protobuf:"bytes,1,opt,name=userId,proto3" json:"userId,omitempty"`
	NodeName string `protobuf:"bytes,2,opt,name=nodeName,proto3" json:"nodeName,omitempty"`
}

func (x *UserCheckInRequest) Reset() {
	*x = UserCheckInRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_grpc_zookeeper_proto_zookeeper_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UserCheckInRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UserCheckInRequest) ProtoMessage() {}

func (x *UserCheckInRequest) ProtoReflect() protoreflect.Message {
	mi := &file_grpc_zookeeper_proto_zookeeper_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UserCheckInRequest.ProtoReflect.Descriptor instead.
func (*UserCheckInRequest) Descriptor() ([]byte, []int) {
	return file_grpc_zookeeper_proto_zookeeper_proto_rawDescGZIP(), []int{4}
}

func (x *UserCheckInRequest) GetUserId() string {
	if x != nil {
		return x.UserId
	}
	return ""
}

func (x *UserCheckInRequest) GetNodeName() string {
	if x != nil {
		return x.NodeName
	}
	return ""
}

type UserCheckInResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Code string `protobuf:"bytes,1,opt,name=code,proto3" json:"code,omitempty"`
	Msg  string `protobuf:"bytes,2,opt,name=msg,proto3" json:"msg,omitempty"`
}

func (x *UserCheckInResponse) Reset() {
	*x = UserCheckInResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_grpc_zookeeper_proto_zookeeper_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UserCheckInResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UserCheckInResponse) ProtoMessage() {}

func (x *UserCheckInResponse) ProtoReflect() protoreflect.Message {
	mi := &file_grpc_zookeeper_proto_zookeeper_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UserCheckInResponse.ProtoReflect.Descriptor instead.
func (*UserCheckInResponse) Descriptor() ([]byte, []int) {
	return file_grpc_zookeeper_proto_zookeeper_proto_rawDescGZIP(), []int{5}
}

func (x *UserCheckInResponse) GetCode() string {
	if x != nil {
		return x.Code
	}
	return ""
}

func (x *UserCheckInResponse) GetMsg() string {
	if x != nil {
		return x.Msg
	}
	return ""
}

type UserCheckOutRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	UserId   string `protobuf:"bytes,1,opt,name=userId,proto3" json:"userId,omitempty"`
	NodeName string `protobuf:"bytes,2,opt,name=nodeName,proto3" json:"nodeName,omitempty"`
}

func (x *UserCheckOutRequest) Reset() {
	*x = UserCheckOutRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_grpc_zookeeper_proto_zookeeper_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UserCheckOutRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UserCheckOutRequest) ProtoMessage() {}

func (x *UserCheckOutRequest) ProtoReflect() protoreflect.Message {
	mi := &file_grpc_zookeeper_proto_zookeeper_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UserCheckOutRequest.ProtoReflect.Descriptor instead.
func (*UserCheckOutRequest) Descriptor() ([]byte, []int) {
	return file_grpc_zookeeper_proto_zookeeper_proto_rawDescGZIP(), []int{6}
}

func (x *UserCheckOutRequest) GetUserId() string {
	if x != nil {
		return x.UserId
	}
	return ""
}

func (x *UserCheckOutRequest) GetNodeName() string {
	if x != nil {
		return x.NodeName
	}
	return ""
}

type UserCheckOutResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Code string `protobuf:"bytes,1,opt,name=code,proto3" json:"code,omitempty"`
	Msg  string `protobuf:"bytes,2,opt,name=msg,proto3" json:"msg,omitempty"`
}

func (x *UserCheckOutResponse) Reset() {
	*x = UserCheckOutResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_grpc_zookeeper_proto_zookeeper_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UserCheckOutResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UserCheckOutResponse) ProtoMessage() {}

func (x *UserCheckOutResponse) ProtoReflect() protoreflect.Message {
	mi := &file_grpc_zookeeper_proto_zookeeper_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UserCheckOutResponse.ProtoReflect.Descriptor instead.
func (*UserCheckOutResponse) Descriptor() ([]byte, []int) {
	return file_grpc_zookeeper_proto_zookeeper_proto_rawDescGZIP(), []int{7}
}

func (x *UserCheckOutResponse) GetCode() string {
	if x != nil {
		return x.Code
	}
	return ""
}

func (x *UserCheckOutResponse) GetMsg() string {
	if x != nil {
		return x.Msg
	}
	return ""
}

type FindUserRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	UserId string `protobuf:"bytes,1,opt,name=userId,proto3" json:"userId,omitempty"`
}

func (x *FindUserRequest) Reset() {
	*x = FindUserRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_grpc_zookeeper_proto_zookeeper_proto_msgTypes[8]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *FindUserRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*FindUserRequest) ProtoMessage() {}

func (x *FindUserRequest) ProtoReflect() protoreflect.Message {
	mi := &file_grpc_zookeeper_proto_zookeeper_proto_msgTypes[8]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use FindUserRequest.ProtoReflect.Descriptor instead.
func (*FindUserRequest) Descriptor() ([]byte, []int) {
	return file_grpc_zookeeper_proto_zookeeper_proto_rawDescGZIP(), []int{8}
}

func (x *FindUserRequest) GetUserId() string {
	if x != nil {
		return x.UserId
	}
	return ""
}

type FindUserResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Code     string `protobuf:"bytes,1,opt,name=code,proto3" json:"code,omitempty"`
	Msg      string `protobuf:"bytes,2,opt,name=msg,proto3" json:"msg,omitempty"`
	NodeName string `protobuf:"bytes,3,opt,name=nodeName,proto3" json:"nodeName,omitempty"`
	GrpcAddr string `protobuf:"bytes,4,opt,name=grpcAddr,proto3" json:"grpcAddr,omitempty"`
	HttpAddr string `protobuf:"bytes,5,opt,name=httpAddr,proto3" json:"httpAddr,omitempty"`
}

func (x *FindUserResponse) Reset() {
	*x = FindUserResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_grpc_zookeeper_proto_zookeeper_proto_msgTypes[9]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *FindUserResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*FindUserResponse) ProtoMessage() {}

func (x *FindUserResponse) ProtoReflect() protoreflect.Message {
	mi := &file_grpc_zookeeper_proto_zookeeper_proto_msgTypes[9]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use FindUserResponse.ProtoReflect.Descriptor instead.
func (*FindUserResponse) Descriptor() ([]byte, []int) {
	return file_grpc_zookeeper_proto_zookeeper_proto_rawDescGZIP(), []int{9}
}

func (x *FindUserResponse) GetCode() string {
	if x != nil {
		return x.Code
	}
	return ""
}

func (x *FindUserResponse) GetMsg() string {
	if x != nil {
		return x.Msg
	}
	return ""
}

func (x *FindUserResponse) GetNodeName() string {
	if x != nil {
		return x.NodeName
	}
	return ""
}

func (x *FindUserResponse) GetGrpcAddr() string {
	if x != nil {
		return x.GrpcAddr
	}
	return ""
}

func (x *FindUserResponse) GetHttpAddr() string {
	if x != nil {
		return x.HttpAddr
	}
	return ""
}

var File_grpc_zookeeper_proto_zookeeper_proto protoreflect.FileDescriptor

var file_grpc_zookeeper_proto_zookeeper_proto_rawDesc = []byte{
	0x0a, 0x24, 0x67, 0x72, 0x70, 0x63, 0x2f, 0x7a, 0x6f, 0x6f, 0x6b, 0x65, 0x65, 0x70, 0x65, 0x72,
	0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x7a, 0x6f, 0x6f, 0x6b, 0x65, 0x65, 0x70, 0x65, 0x72,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x09, 0x7a, 0x6f, 0x6f, 0x6b, 0x65, 0x65, 0x70, 0x65,
	0x72, 0x22, 0x65, 0x0a, 0x0f, 0x52, 0x65, 0x67, 0x69, 0x73, 0x74, 0x65, 0x72, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x12, 0x1a, 0x0a, 0x08, 0x6e, 0x6f, 0x64, 0x65, 0x4e, 0x61, 0x6d, 0x65,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x6e, 0x6f, 0x64, 0x65, 0x4e, 0x61, 0x6d, 0x65,
	0x12, 0x1a, 0x0a, 0x08, 0x68, 0x74, 0x74, 0x70, 0x41, 0x64, 0x64, 0x72, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x08, 0x68, 0x74, 0x74, 0x70, 0x41, 0x64, 0x64, 0x72, 0x12, 0x1a, 0x0a, 0x08,
	0x67, 0x72, 0x70, 0x63, 0x41, 0x64, 0x64, 0x72, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08,
	0x67, 0x72, 0x70, 0x63, 0x41, 0x64, 0x64, 0x72, 0x22, 0x38, 0x0a, 0x10, 0x52, 0x65, 0x67, 0x69,
	0x73, 0x74, 0x65, 0x72, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x12, 0x0a, 0x04,
	0x63, 0x6f, 0x64, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x63, 0x6f, 0x64, 0x65,
	0x12, 0x10, 0x0a, 0x03, 0x6d, 0x73, 0x67, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x6d,
	0x73, 0x67, 0x22, 0x66, 0x0a, 0x10, 0x48, 0x65, 0x61, 0x72, 0x74, 0x62, 0x65, 0x61, 0x74, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x1a, 0x0a, 0x08, 0x6e, 0x6f, 0x64, 0x65, 0x4e, 0x61,
	0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x6e, 0x6f, 0x64, 0x65, 0x4e, 0x61,
	0x6d, 0x65, 0x12, 0x1a, 0x0a, 0x08, 0x67, 0x72, 0x70, 0x63, 0x41, 0x64, 0x64, 0x72, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x67, 0x72, 0x70, 0x63, 0x41, 0x64, 0x64, 0x72, 0x12, 0x1a,
	0x0a, 0x08, 0x68, 0x74, 0x74, 0x70, 0x41, 0x64, 0x64, 0x72, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x08, 0x68, 0x74, 0x74, 0x70, 0x41, 0x64, 0x64, 0x72, 0x22, 0x39, 0x0a, 0x11, 0x48, 0x65,
	0x61, 0x72, 0x74, 0x62, 0x65, 0x61, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12,
	0x12, 0x0a, 0x04, 0x63, 0x6f, 0x64, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x63,
	0x6f, 0x64, 0x65, 0x12, 0x10, 0x0a, 0x03, 0x6d, 0x73, 0x67, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x03, 0x6d, 0x73, 0x67, 0x22, 0x48, 0x0a, 0x12, 0x55, 0x73, 0x65, 0x72, 0x43, 0x68, 0x65,
	0x63, 0x6b, 0x49, 0x6e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x16, 0x0a, 0x06, 0x75,
	0x73, 0x65, 0x72, 0x49, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x75, 0x73, 0x65,
	0x72, 0x49, 0x64, 0x12, 0x1a, 0x0a, 0x08, 0x6e, 0x6f, 0x64, 0x65, 0x4e, 0x61, 0x6d, 0x65, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x6e, 0x6f, 0x64, 0x65, 0x4e, 0x61, 0x6d, 0x65, 0x22,
	0x3b, 0x0a, 0x13, 0x55, 0x73, 0x65, 0x72, 0x43, 0x68, 0x65, 0x63, 0x6b, 0x49, 0x6e, 0x52, 0x65,
	0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x63, 0x6f, 0x64, 0x65, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x63, 0x6f, 0x64, 0x65, 0x12, 0x10, 0x0a, 0x03, 0x6d, 0x73,
	0x67, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x6d, 0x73, 0x67, 0x22, 0x49, 0x0a, 0x13,
	0x55, 0x73, 0x65, 0x72, 0x43, 0x68, 0x65, 0x63, 0x6b, 0x4f, 0x75, 0x74, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x12, 0x16, 0x0a, 0x06, 0x75, 0x73, 0x65, 0x72, 0x49, 0x64, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x06, 0x75, 0x73, 0x65, 0x72, 0x49, 0x64, 0x12, 0x1a, 0x0a, 0x08, 0x6e,
	0x6f, 0x64, 0x65, 0x4e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x6e,
	0x6f, 0x64, 0x65, 0x4e, 0x61, 0x6d, 0x65, 0x22, 0x3c, 0x0a, 0x14, 0x55, 0x73, 0x65, 0x72, 0x43,
	0x68, 0x65, 0x63, 0x6b, 0x4f, 0x75, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12,
	0x12, 0x0a, 0x04, 0x63, 0x6f, 0x64, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x63,
	0x6f, 0x64, 0x65, 0x12, 0x10, 0x0a, 0x03, 0x6d, 0x73, 0x67, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x03, 0x6d, 0x73, 0x67, 0x22, 0x29, 0x0a, 0x0f, 0x46, 0x69, 0x6e, 0x64, 0x55, 0x73, 0x65,
	0x72, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x16, 0x0a, 0x06, 0x75, 0x73, 0x65, 0x72,
	0x49, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x75, 0x73, 0x65, 0x72, 0x49, 0x64,
	0x22, 0x8c, 0x01, 0x0a, 0x10, 0x46, 0x69, 0x6e, 0x64, 0x55, 0x73, 0x65, 0x72, 0x52, 0x65, 0x73,
	0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x63, 0x6f, 0x64, 0x65, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x04, 0x63, 0x6f, 0x64, 0x65, 0x12, 0x10, 0x0a, 0x03, 0x6d, 0x73, 0x67,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x6d, 0x73, 0x67, 0x12, 0x1a, 0x0a, 0x08, 0x6e,
	0x6f, 0x64, 0x65, 0x4e, 0x61, 0x6d, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x6e,
	0x6f, 0x64, 0x65, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x1a, 0x0a, 0x08, 0x67, 0x72, 0x70, 0x63, 0x41,
	0x64, 0x64, 0x72, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x67, 0x72, 0x70, 0x63, 0x41,
	0x64, 0x64, 0x72, 0x12, 0x1a, 0x0a, 0x08, 0x68, 0x74, 0x74, 0x70, 0x41, 0x64, 0x64, 0x72, 0x18,
	0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x68, 0x74, 0x74, 0x70, 0x41, 0x64, 0x64, 0x72, 0x32,
	0x86, 0x03, 0x0a, 0x09, 0x5a, 0x6f, 0x6f, 0x4b, 0x65, 0x65, 0x70, 0x65, 0x72, 0x12, 0x45, 0x0a,
	0x08, 0x52, 0x65, 0x67, 0x69, 0x73, 0x74, 0x65, 0x72, 0x12, 0x1a, 0x2e, 0x7a, 0x6f, 0x6f, 0x6b,
	0x65, 0x65, 0x70, 0x65, 0x72, 0x2e, 0x52, 0x65, 0x67, 0x69, 0x73, 0x74, 0x65, 0x72, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1b, 0x2e, 0x7a, 0x6f, 0x6f, 0x6b, 0x65, 0x65, 0x70, 0x65,
	0x72, 0x2e, 0x52, 0x65, 0x67, 0x69, 0x73, 0x74, 0x65, 0x72, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e,
	0x73, 0x65, 0x22, 0x00, 0x12, 0x48, 0x0a, 0x09, 0x48, 0x65, 0x61, 0x72, 0x74, 0x62, 0x65, 0x61,
	0x74, 0x12, 0x1b, 0x2e, 0x7a, 0x6f, 0x6f, 0x6b, 0x65, 0x65, 0x70, 0x65, 0x72, 0x2e, 0x48, 0x65,
	0x61, 0x72, 0x74, 0x62, 0x65, 0x61, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1c,
	0x2e, 0x7a, 0x6f, 0x6f, 0x6b, 0x65, 0x65, 0x70, 0x65, 0x72, 0x2e, 0x48, 0x65, 0x61, 0x72, 0x74,
	0x62, 0x65, 0x61, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x4e,
	0x0a, 0x0b, 0x55, 0x73, 0x65, 0x72, 0x43, 0x68, 0x65, 0x63, 0x6b, 0x49, 0x6e, 0x12, 0x1d, 0x2e,
	0x7a, 0x6f, 0x6f, 0x6b, 0x65, 0x65, 0x70, 0x65, 0x72, 0x2e, 0x55, 0x73, 0x65, 0x72, 0x43, 0x68,
	0x65, 0x63, 0x6b, 0x49, 0x6e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1e, 0x2e, 0x7a,
	0x6f, 0x6f, 0x6b, 0x65, 0x65, 0x70, 0x65, 0x72, 0x2e, 0x55, 0x73, 0x65, 0x72, 0x43, 0x68, 0x65,
	0x63, 0x6b, 0x49, 0x6e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x51,
	0x0a, 0x0c, 0x55, 0x73, 0x65, 0x72, 0x43, 0x68, 0x65, 0x63, 0x6b, 0x4f, 0x75, 0x74, 0x12, 0x1e,
	0x2e, 0x7a, 0x6f, 0x6f, 0x6b, 0x65, 0x65, 0x70, 0x65, 0x72, 0x2e, 0x55, 0x73, 0x65, 0x72, 0x43,
	0x68, 0x65, 0x63, 0x6b, 0x4f, 0x75, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1f,
	0x2e, 0x7a, 0x6f, 0x6f, 0x6b, 0x65, 0x65, 0x70, 0x65, 0x72, 0x2e, 0x55, 0x73, 0x65, 0x72, 0x43,
	0x68, 0x65, 0x63, 0x6b, 0x4f, 0x75, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22,
	0x00, 0x12, 0x45, 0x0a, 0x08, 0x46, 0x69, 0x6e, 0x64, 0x55, 0x73, 0x65, 0x72, 0x12, 0x1a, 0x2e,
	0x7a, 0x6f, 0x6f, 0x6b, 0x65, 0x65, 0x70, 0x65, 0x72, 0x2e, 0x46, 0x69, 0x6e, 0x64, 0x55, 0x73,
	0x65, 0x72, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1b, 0x2e, 0x7a, 0x6f, 0x6f, 0x6b,
	0x65, 0x65, 0x70, 0x65, 0x72, 0x2e, 0x46, 0x69, 0x6e, 0x64, 0x55, 0x73, 0x65, 0x72, 0x52, 0x65,
	0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x42, 0x0c, 0x5a, 0x0a, 0x2f, 0x7a, 0x6f, 0x6f,
	0x6b, 0x65, 0x65, 0x70, 0x65, 0x72, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_grpc_zookeeper_proto_zookeeper_proto_rawDescOnce sync.Once
	file_grpc_zookeeper_proto_zookeeper_proto_rawDescData = file_grpc_zookeeper_proto_zookeeper_proto_rawDesc
)

func file_grpc_zookeeper_proto_zookeeper_proto_rawDescGZIP() []byte {
	file_grpc_zookeeper_proto_zookeeper_proto_rawDescOnce.Do(func() {
		file_grpc_zookeeper_proto_zookeeper_proto_rawDescData = protoimpl.X.CompressGZIP(file_grpc_zookeeper_proto_zookeeper_proto_rawDescData)
	})
	return file_grpc_zookeeper_proto_zookeeper_proto_rawDescData
}

var file_grpc_zookeeper_proto_zookeeper_proto_msgTypes = make([]protoimpl.MessageInfo, 10)
var file_grpc_zookeeper_proto_zookeeper_proto_goTypes = []interface{}{
	(*RegisterRequest)(nil),      // 0: zookeeper.RegisterRequest
	(*RegisterResponse)(nil),     // 1: zookeeper.RegisterResponse
	(*HeartbeatRequest)(nil),     // 2: zookeeper.HeartbeatRequest
	(*HeartbeatResponse)(nil),    // 3: zookeeper.HeartbeatResponse
	(*UserCheckInRequest)(nil),   // 4: zookeeper.UserCheckInRequest
	(*UserCheckInResponse)(nil),  // 5: zookeeper.UserCheckInResponse
	(*UserCheckOutRequest)(nil),  // 6: zookeeper.UserCheckOutRequest
	(*UserCheckOutResponse)(nil), // 7: zookeeper.UserCheckOutResponse
	(*FindUserRequest)(nil),      // 8: zookeeper.FindUserRequest
	(*FindUserResponse)(nil),     // 9: zookeeper.FindUserResponse
}
var file_grpc_zookeeper_proto_zookeeper_proto_depIdxs = []int32{
	0, // 0: zookeeper.ZooKeeper.Register:input_type -> zookeeper.RegisterRequest
	2, // 1: zookeeper.ZooKeeper.Heartbeat:input_type -> zookeeper.HeartbeatRequest
	4, // 2: zookeeper.ZooKeeper.UserCheckIn:input_type -> zookeeper.UserCheckInRequest
	6, // 3: zookeeper.ZooKeeper.UserCheckOut:input_type -> zookeeper.UserCheckOutRequest
	8, // 4: zookeeper.ZooKeeper.FindUser:input_type -> zookeeper.FindUserRequest
	1, // 5: zookeeper.ZooKeeper.Register:output_type -> zookeeper.RegisterResponse
	3, // 6: zookeeper.ZooKeeper.Heartbeat:output_type -> zookeeper.HeartbeatResponse
	5, // 7: zookeeper.ZooKeeper.UserCheckIn:output_type -> zookeeper.UserCheckInResponse
	7, // 8: zookeeper.ZooKeeper.UserCheckOut:output_type -> zookeeper.UserCheckOutResponse
	9, // 9: zookeeper.ZooKeeper.FindUser:output_type -> zookeeper.FindUserResponse
	5, // [5:10] is the sub-list for method output_type
	0, // [0:5] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_grpc_zookeeper_proto_zookeeper_proto_init() }
func file_grpc_zookeeper_proto_zookeeper_proto_init() {
	if File_grpc_zookeeper_proto_zookeeper_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_grpc_zookeeper_proto_zookeeper_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RegisterRequest); i {
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
		file_grpc_zookeeper_proto_zookeeper_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RegisterResponse); i {
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
		file_grpc_zookeeper_proto_zookeeper_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*HeartbeatRequest); i {
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
		file_grpc_zookeeper_proto_zookeeper_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*HeartbeatResponse); i {
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
		file_grpc_zookeeper_proto_zookeeper_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UserCheckInRequest); i {
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
		file_grpc_zookeeper_proto_zookeeper_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UserCheckInResponse); i {
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
		file_grpc_zookeeper_proto_zookeeper_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UserCheckOutRequest); i {
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
		file_grpc_zookeeper_proto_zookeeper_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UserCheckOutResponse); i {
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
		file_grpc_zookeeper_proto_zookeeper_proto_msgTypes[8].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*FindUserRequest); i {
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
		file_grpc_zookeeper_proto_zookeeper_proto_msgTypes[9].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*FindUserResponse); i {
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
			RawDescriptor: file_grpc_zookeeper_proto_zookeeper_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   10,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_grpc_zookeeper_proto_zookeeper_proto_goTypes,
		DependencyIndexes: file_grpc_zookeeper_proto_zookeeper_proto_depIdxs,
		MessageInfos:      file_grpc_zookeeper_proto_zookeeper_proto_msgTypes,
	}.Build()
	File_grpc_zookeeper_proto_zookeeper_proto = out.File
	file_grpc_zookeeper_proto_zookeeper_proto_rawDesc = nil
	file_grpc_zookeeper_proto_zookeeper_proto_goTypes = nil
	file_grpc_zookeeper_proto_zookeeper_proto_depIdxs = nil
}
