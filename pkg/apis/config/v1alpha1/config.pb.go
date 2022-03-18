//
//Copyright © 2022 Ci4Rail GmbH <engineering@ci4rail.com>
//Licensed under the Apache License, Version 2.0 (the "License");
//you may not use this file except in compliance with the License.
//You may obtain a copy of the License at
//http://www.apache.org/licenses/LICENSE-2.0
//Unless required by applicable law or agreed to in writing, software
//distributed under the License is distributed on an "AS IS" BASIS,
//WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
//See the License for the specific language governing permissions and
//limitations under the License.

// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.27.1
// 	protoc        v3.6.1
// source: config.proto

package v1alpha1

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

type SysAccountRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *SysAccountRequest) Reset() {
	*x = SysAccountRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_config_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SysAccountRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SysAccountRequest) ProtoMessage() {}

func (x *SysAccountRequest) ProtoReflect() protoreflect.Message {
	mi := &file_config_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SysAccountRequest.ProtoReflect.Descriptor instead.
func (*SysAccountRequest) Descriptor() ([]byte, []int) {
	return file_config_proto_rawDescGZIP(), []int{0}
}

type SysAccountResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	OperatorJWT  string `protobuf:"bytes,1,opt,name=operatorJWT,proto3" json:"operatorJWT,omitempty"`
	SysCreds     string `protobuf:"bytes,2,opt,name=sysCreds,proto3" json:"sysCreds,omitempty"`
	SysJWT       string `protobuf:"bytes,3,opt,name=sysJWT,proto3" json:"sysJWT,omitempty"`
	SysPublicKey string `protobuf:"bytes,4,opt,name=sysPublicKey,proto3" json:"sysPublicKey,omitempty"`
}

func (x *SysAccountResponse) Reset() {
	*x = SysAccountResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_config_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SysAccountResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SysAccountResponse) ProtoMessage() {}

func (x *SysAccountResponse) ProtoReflect() protoreflect.Message {
	mi := &file_config_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SysAccountResponse.ProtoReflect.Descriptor instead.
func (*SysAccountResponse) Descriptor() ([]byte, []int) {
	return file_config_proto_rawDescGZIP(), []int{1}
}

func (x *SysAccountResponse) GetOperatorJWT() string {
	if x != nil {
		return x.OperatorJWT
	}
	return ""
}

func (x *SysAccountResponse) GetSysCreds() string {
	if x != nil {
		return x.SysCreds
	}
	return ""
}

func (x *SysAccountResponse) GetSysJWT() string {
	if x != nil {
		return x.SysJWT
	}
	return ""
}

func (x *SysAccountResponse) GetSysPublicKey() string {
	if x != nil {
		return x.SysPublicKey
	}
	return ""
}

type DesiredStateRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Network      string   `protobuf:"bytes,1,opt,name=network,proto3" json:"network,omitempty"`
	Participants []string `protobuf:"bytes,2,rep,name=participants,proto3" json:"participants,omitempty"`
}

func (x *DesiredStateRequest) Reset() {
	*x = DesiredStateRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_config_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DesiredStateRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DesiredStateRequest) ProtoMessage() {}

func (x *DesiredStateRequest) ProtoReflect() protoreflect.Message {
	mi := &file_config_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DesiredStateRequest.ProtoReflect.Descriptor instead.
func (*DesiredStateRequest) Descriptor() ([]byte, []int) {
	return file_config_proto_rawDescGZIP(), []int{2}
}

func (x *DesiredStateRequest) GetNetwork() string {
	if x != nil {
		return x.Network
	}
	return ""
}

func (x *DesiredStateRequest) GetParticipants() []string {
	if x != nil {
		return x.Participants
	}
	return nil
}

type Credentials struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	NetworkParticipant string `protobuf:"bytes,1,opt,name=networkParticipant,proto3" json:"networkParticipant,omitempty"`
	Creds              string `protobuf:"bytes,2,opt,name=creds,proto3" json:"creds,omitempty"`
	AccoutPublicKey    string `protobuf:"bytes,3,opt,name=accoutPublicKey,proto3" json:"accoutPublicKey,omitempty"`
}

func (x *Credentials) Reset() {
	*x = Credentials{}
	if protoimpl.UnsafeEnabled {
		mi := &file_config_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Credentials) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Credentials) ProtoMessage() {}

func (x *Credentials) ProtoReflect() protoreflect.Message {
	mi := &file_config_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Credentials.ProtoReflect.Descriptor instead.
func (*Credentials) Descriptor() ([]byte, []int) {
	return file_config_proto_rawDescGZIP(), []int{3}
}

func (x *Credentials) GetNetworkParticipant() string {
	if x != nil {
		return x.NetworkParticipant
	}
	return ""
}

func (x *Credentials) GetCreds() string {
	if x != nil {
		return x.Creds
	}
	return ""
}

func (x *Credentials) GetAccoutPublicKey() string {
	if x != nil {
		return x.AccoutPublicKey
	}
	return ""
}

type DesiredStateResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Creds               []*Credentials `protobuf:"bytes,1,rep,name=creds,proto3" json:"creds,omitempty"`
	DeletedParticipants []string       `protobuf:"bytes,2,rep,name=deletedParticipants,proto3" json:"deletedParticipants,omitempty"`
}

func (x *DesiredStateResponse) Reset() {
	*x = DesiredStateResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_config_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DesiredStateResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DesiredStateResponse) ProtoMessage() {}

func (x *DesiredStateResponse) ProtoReflect() protoreflect.Message {
	mi := &file_config_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DesiredStateResponse.ProtoReflect.Descriptor instead.
func (*DesiredStateResponse) Descriptor() ([]byte, []int) {
	return file_config_proto_rawDescGZIP(), []int{4}
}

func (x *DesiredStateResponse) GetCreds() []*Credentials {
	if x != nil {
		return x.Creds
	}
	return nil
}

func (x *DesiredStateResponse) GetDeletedParticipants() []string {
	if x != nil {
		return x.DeletedParticipants
	}
	return nil
}

type DeleteNetworkRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Network string `protobuf:"bytes,1,opt,name=network,proto3" json:"network,omitempty"`
}

func (x *DeleteNetworkRequest) Reset() {
	*x = DeleteNetworkRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_config_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DeleteNetworkRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeleteNetworkRequest) ProtoMessage() {}

func (x *DeleteNetworkRequest) ProtoReflect() protoreflect.Message {
	mi := &file_config_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeleteNetworkRequest.ProtoReflect.Descriptor instead.
func (*DeleteNetworkRequest) Descriptor() ([]byte, []int) {
	return file_config_proto_rawDescGZIP(), []int{5}
}

func (x *DeleteNetworkRequest) GetNetwork() string {
	if x != nil {
		return x.Network
	}
	return ""
}

type DeleteNetworkResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *DeleteNetworkResponse) Reset() {
	*x = DeleteNetworkResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_config_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DeleteNetworkResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeleteNetworkResponse) ProtoMessage() {}

func (x *DeleteNetworkResponse) ProtoReflect() protoreflect.Message {
	mi := &file_config_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeleteNetworkResponse.ProtoReflect.Descriptor instead.
func (*DeleteNetworkResponse) Descriptor() ([]byte, []int) {
	return file_config_proto_rawDescGZIP(), []int{6}
}

var File_config_proto protoreflect.FileDescriptor

var file_config_proto_rawDesc = []byte{
	0x0a, 0x0c, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x08,
	0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x31, 0x22, 0x13, 0x0a, 0x11, 0x53, 0x79, 0x73, 0x41,
	0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x22, 0x8e, 0x01,
	0x0a, 0x12, 0x53, 0x79, 0x73, 0x41, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x52, 0x65, 0x73, 0x70,
	0x6f, 0x6e, 0x73, 0x65, 0x12, 0x20, 0x0a, 0x0b, 0x6f, 0x70, 0x65, 0x72, 0x61, 0x74, 0x6f, 0x72,
	0x4a, 0x57, 0x54, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x6f, 0x70, 0x65, 0x72, 0x61,
	0x74, 0x6f, 0x72, 0x4a, 0x57, 0x54, 0x12, 0x1a, 0x0a, 0x08, 0x73, 0x79, 0x73, 0x43, 0x72, 0x65,
	0x64, 0x73, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x73, 0x79, 0x73, 0x43, 0x72, 0x65,
	0x64, 0x73, 0x12, 0x16, 0x0a, 0x06, 0x73, 0x79, 0x73, 0x4a, 0x57, 0x54, 0x18, 0x03, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x06, 0x73, 0x79, 0x73, 0x4a, 0x57, 0x54, 0x12, 0x22, 0x0a, 0x0c, 0x73, 0x79,
	0x73, 0x50, 0x75, 0x62, 0x6c, 0x69, 0x63, 0x4b, 0x65, 0x79, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x0c, 0x73, 0x79, 0x73, 0x50, 0x75, 0x62, 0x6c, 0x69, 0x63, 0x4b, 0x65, 0x79, 0x22, 0x53,
	0x0a, 0x13, 0x44, 0x65, 0x73, 0x69, 0x72, 0x65, 0x64, 0x53, 0x74, 0x61, 0x74, 0x65, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x18, 0x0a, 0x07, 0x6e, 0x65, 0x74, 0x77, 0x6f, 0x72, 0x6b,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x6e, 0x65, 0x74, 0x77, 0x6f, 0x72, 0x6b, 0x12,
	0x22, 0x0a, 0x0c, 0x70, 0x61, 0x72, 0x74, 0x69, 0x63, 0x69, 0x70, 0x61, 0x6e, 0x74, 0x73, 0x18,
	0x02, 0x20, 0x03, 0x28, 0x09, 0x52, 0x0c, 0x70, 0x61, 0x72, 0x74, 0x69, 0x63, 0x69, 0x70, 0x61,
	0x6e, 0x74, 0x73, 0x22, 0x7d, 0x0a, 0x0b, 0x43, 0x72, 0x65, 0x64, 0x65, 0x6e, 0x74, 0x69, 0x61,
	0x6c, 0x73, 0x12, 0x2e, 0x0a, 0x12, 0x6e, 0x65, 0x74, 0x77, 0x6f, 0x72, 0x6b, 0x50, 0x61, 0x72,
	0x74, 0x69, 0x63, 0x69, 0x70, 0x61, 0x6e, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x12,
	0x6e, 0x65, 0x74, 0x77, 0x6f, 0x72, 0x6b, 0x50, 0x61, 0x72, 0x74, 0x69, 0x63, 0x69, 0x70, 0x61,
	0x6e, 0x74, 0x12, 0x14, 0x0a, 0x05, 0x63, 0x72, 0x65, 0x64, 0x73, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x05, 0x63, 0x72, 0x65, 0x64, 0x73, 0x12, 0x28, 0x0a, 0x0f, 0x61, 0x63, 0x63, 0x6f,
	0x75, 0x74, 0x50, 0x75, 0x62, 0x6c, 0x69, 0x63, 0x4b, 0x65, 0x79, 0x18, 0x03, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x0f, 0x61, 0x63, 0x63, 0x6f, 0x75, 0x74, 0x50, 0x75, 0x62, 0x6c, 0x69, 0x63, 0x4b,
	0x65, 0x79, 0x22, 0x75, 0x0a, 0x14, 0x44, 0x65, 0x73, 0x69, 0x72, 0x65, 0x64, 0x53, 0x74, 0x61,
	0x74, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x2b, 0x0a, 0x05, 0x63, 0x72,
	0x65, 0x64, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x15, 0x2e, 0x76, 0x31, 0x61, 0x6c,
	0x70, 0x68, 0x61, 0x31, 0x2e, 0x43, 0x72, 0x65, 0x64, 0x65, 0x6e, 0x74, 0x69, 0x61, 0x6c, 0x73,
	0x52, 0x05, 0x63, 0x72, 0x65, 0x64, 0x73, 0x12, 0x30, 0x0a, 0x13, 0x64, 0x65, 0x6c, 0x65, 0x74,
	0x65, 0x64, 0x50, 0x61, 0x72, 0x74, 0x69, 0x63, 0x69, 0x70, 0x61, 0x6e, 0x74, 0x73, 0x18, 0x02,
	0x20, 0x03, 0x28, 0x09, 0x52, 0x13, 0x64, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x64, 0x50, 0x61, 0x72,
	0x74, 0x69, 0x63, 0x69, 0x70, 0x61, 0x6e, 0x74, 0x73, 0x22, 0x30, 0x0a, 0x14, 0x44, 0x65, 0x6c,
	0x65, 0x74, 0x65, 0x4e, 0x65, 0x74, 0x77, 0x6f, 0x72, 0x6b, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x12, 0x18, 0x0a, 0x07, 0x6e, 0x65, 0x74, 0x77, 0x6f, 0x72, 0x6b, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x07, 0x6e, 0x65, 0x74, 0x77, 0x6f, 0x72, 0x6b, 0x22, 0x17, 0x0a, 0x15, 0x44,
	0x65, 0x6c, 0x65, 0x74, 0x65, 0x4e, 0x65, 0x74, 0x77, 0x6f, 0x72, 0x6b, 0x52, 0x65, 0x73, 0x70,
	0x6f, 0x6e, 0x73, 0x65, 0x32, 0xfb, 0x01, 0x0a, 0x0d, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x53,
	0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x49, 0x0a, 0x0a, 0x53, 0x79, 0x73, 0x41, 0x63, 0x63,
	0x6f, 0x75, 0x6e, 0x74, 0x12, 0x1b, 0x2e, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x31, 0x2e,
	0x53, 0x79, 0x73, 0x41, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x1a, 0x1c, 0x2e, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x31, 0x2e, 0x53, 0x79, 0x73,
	0x41, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22,
	0x00, 0x12, 0x4d, 0x0a, 0x0c, 0x44, 0x65, 0x73, 0x69, 0x72, 0x65, 0x64, 0x53, 0x74, 0x61, 0x74,
	0x65, 0x12, 0x1d, 0x2e, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x31, 0x2e, 0x44, 0x65, 0x73,
	0x69, 0x72, 0x65, 0x64, 0x53, 0x74, 0x61, 0x74, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x1a, 0x1e, 0x2e, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x31, 0x2e, 0x44, 0x65, 0x73, 0x69,
	0x72, 0x65, 0x64, 0x53, 0x74, 0x61, 0x74, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65,
	0x12, 0x50, 0x0a, 0x0d, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x4e, 0x65, 0x74, 0x77, 0x6f, 0x72,
	0x6b, 0x12, 0x1e, 0x2e, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x31, 0x2e, 0x44, 0x65, 0x6c,
	0x65, 0x74, 0x65, 0x4e, 0x65, 0x74, 0x77, 0x6f, 0x72, 0x6b, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x1a, 0x1f, 0x2e, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x31, 0x2e, 0x44, 0x65, 0x6c,
	0x65, 0x74, 0x65, 0x4e, 0x65, 0x74, 0x77, 0x6f, 0x72, 0x6b, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e,
	0x73, 0x65, 0x42, 0x0c, 0x5a, 0x0a, 0x2e, 0x3b, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x31,
	0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_config_proto_rawDescOnce sync.Once
	file_config_proto_rawDescData = file_config_proto_rawDesc
)

func file_config_proto_rawDescGZIP() []byte {
	file_config_proto_rawDescOnce.Do(func() {
		file_config_proto_rawDescData = protoimpl.X.CompressGZIP(file_config_proto_rawDescData)
	})
	return file_config_proto_rawDescData
}

var file_config_proto_msgTypes = make([]protoimpl.MessageInfo, 7)
var file_config_proto_goTypes = []interface{}{
	(*SysAccountRequest)(nil),     // 0: v1alpha1.SysAccountRequest
	(*SysAccountResponse)(nil),    // 1: v1alpha1.SysAccountResponse
	(*DesiredStateRequest)(nil),   // 2: v1alpha1.DesiredStateRequest
	(*Credentials)(nil),           // 3: v1alpha1.Credentials
	(*DesiredStateResponse)(nil),  // 4: v1alpha1.DesiredStateResponse
	(*DeleteNetworkRequest)(nil),  // 5: v1alpha1.DeleteNetworkRequest
	(*DeleteNetworkResponse)(nil), // 6: v1alpha1.DeleteNetworkResponse
}
var file_config_proto_depIdxs = []int32{
	3, // 0: v1alpha1.DesiredStateResponse.creds:type_name -> v1alpha1.Credentials
	0, // 1: v1alpha1.ConfigService.SysAccount:input_type -> v1alpha1.SysAccountRequest
	2, // 2: v1alpha1.ConfigService.DesiredState:input_type -> v1alpha1.DesiredStateRequest
	5, // 3: v1alpha1.ConfigService.DeleteNetwork:input_type -> v1alpha1.DeleteNetworkRequest
	1, // 4: v1alpha1.ConfigService.SysAccount:output_type -> v1alpha1.SysAccountResponse
	4, // 5: v1alpha1.ConfigService.DesiredState:output_type -> v1alpha1.DesiredStateResponse
	6, // 6: v1alpha1.ConfigService.DeleteNetwork:output_type -> v1alpha1.DeleteNetworkResponse
	4, // [4:7] is the sub-list for method output_type
	1, // [1:4] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_config_proto_init() }
func file_config_proto_init() {
	if File_config_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_config_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SysAccountRequest); i {
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
		file_config_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SysAccountResponse); i {
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
		file_config_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DesiredStateRequest); i {
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
		file_config_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Credentials); i {
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
		file_config_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DesiredStateResponse); i {
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
		file_config_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DeleteNetworkRequest); i {
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
		file_config_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DeleteNetworkResponse); i {
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
			RawDescriptor: file_config_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   7,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_config_proto_goTypes,
		DependencyIndexes: file_config_proto_depIdxs,
		MessageInfos:      file_config_proto_msgTypes,
	}.Build()
	File_config_proto = out.File
	file_config_proto_rawDesc = nil
	file_config_proto_goTypes = nil
	file_config_proto_depIdxs = nil
}
