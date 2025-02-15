// Copyright (c) Mainflux
// SPDX-License-Identifier: Apache-2.0

// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v3.21.12
// source: users/policies/auth.proto

package policies

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

type AuthorizeReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Sub        string `protobuf:"bytes,1,opt,name=sub,proto3" json:"sub,omitempty"`
	Obj        string `protobuf:"bytes,2,opt,name=obj,proto3" json:"obj,omitempty"`
	Act        string `protobuf:"bytes,3,opt,name=act,proto3" json:"act,omitempty"`
	EntityType string `protobuf:"bytes,4,opt,name=entityType,proto3" json:"entityType,omitempty"`
}

func (x *AuthorizeReq) Reset() {
	*x = AuthorizeReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_users_policies_auth_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AuthorizeReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AuthorizeReq) ProtoMessage() {}

func (x *AuthorizeReq) ProtoReflect() protoreflect.Message {
	mi := &file_users_policies_auth_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AuthorizeReq.ProtoReflect.Descriptor instead.
func (*AuthorizeReq) Descriptor() ([]byte, []int) {
	return file_users_policies_auth_proto_rawDescGZIP(), []int{0}
}

func (x *AuthorizeReq) GetSub() string {
	if x != nil {
		return x.Sub
	}
	return ""
}

func (x *AuthorizeReq) GetObj() string {
	if x != nil {
		return x.Obj
	}
	return ""
}

func (x *AuthorizeReq) GetAct() string {
	if x != nil {
		return x.Act
	}
	return ""
}

func (x *AuthorizeReq) GetEntityType() string {
	if x != nil {
		return x.EntityType
	}
	return ""
}

type AuthorizeRes struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Authorized bool `protobuf:"varint,1,opt,name=authorized,proto3" json:"authorized,omitempty"`
}

func (x *AuthorizeRes) Reset() {
	*x = AuthorizeRes{}
	if protoimpl.UnsafeEnabled {
		mi := &file_users_policies_auth_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AuthorizeRes) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AuthorizeRes) ProtoMessage() {}

func (x *AuthorizeRes) ProtoReflect() protoreflect.Message {
	mi := &file_users_policies_auth_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AuthorizeRes.ProtoReflect.Descriptor instead.
func (*AuthorizeRes) Descriptor() ([]byte, []int) {
	return file_users_policies_auth_proto_rawDescGZIP(), []int{1}
}

func (x *AuthorizeRes) GetAuthorized() bool {
	if x != nil {
		return x.Authorized
	}
	return false
}

type IssueReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Email    string `protobuf:"bytes,1,opt,name=email,proto3" json:"email,omitempty"`
	Password string `protobuf:"bytes,2,opt,name=password,proto3" json:"password,omitempty"`
	Type     uint32 `protobuf:"varint,3,opt,name=type,proto3" json:"type,omitempty"`
}

func (x *IssueReq) Reset() {
	*x = IssueReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_users_policies_auth_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *IssueReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*IssueReq) ProtoMessage() {}

func (x *IssueReq) ProtoReflect() protoreflect.Message {
	mi := &file_users_policies_auth_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use IssueReq.ProtoReflect.Descriptor instead.
func (*IssueReq) Descriptor() ([]byte, []int) {
	return file_users_policies_auth_proto_rawDescGZIP(), []int{2}
}

func (x *IssueReq) GetEmail() string {
	if x != nil {
		return x.Email
	}
	return ""
}

func (x *IssueReq) GetPassword() string {
	if x != nil {
		return x.Password
	}
	return ""
}

func (x *IssueReq) GetType() uint32 {
	if x != nil {
		return x.Type
	}
	return 0
}

type Token struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Value string `protobuf:"bytes,1,opt,name=value,proto3" json:"value,omitempty"`
}

func (x *Token) Reset() {
	*x = Token{}
	if protoimpl.UnsafeEnabled {
		mi := &file_users_policies_auth_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Token) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Token) ProtoMessage() {}

func (x *Token) ProtoReflect() protoreflect.Message {
	mi := &file_users_policies_auth_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Token.ProtoReflect.Descriptor instead.
func (*Token) Descriptor() ([]byte, []int) {
	return file_users_policies_auth_proto_rawDescGZIP(), []int{3}
}

func (x *Token) GetValue() string {
	if x != nil {
		return x.Value
	}
	return ""
}

type UserIdentity struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
}

func (x *UserIdentity) Reset() {
	*x = UserIdentity{}
	if protoimpl.UnsafeEnabled {
		mi := &file_users_policies_auth_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UserIdentity) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UserIdentity) ProtoMessage() {}

func (x *UserIdentity) ProtoReflect() protoreflect.Message {
	mi := &file_users_policies_auth_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UserIdentity.ProtoReflect.Descriptor instead.
func (*UserIdentity) Descriptor() ([]byte, []int) {
	return file_users_policies_auth_proto_rawDescGZIP(), []int{4}
}

func (x *UserIdentity) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

type AddPolicyReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Token string   `protobuf:"bytes,1,opt,name=token,proto3" json:"token,omitempty"`
	Sub   string   `protobuf:"bytes,2,opt,name=sub,proto3" json:"sub,omitempty"`
	Obj   string   `protobuf:"bytes,3,opt,name=obj,proto3" json:"obj,omitempty"`
	Act   []string `protobuf:"bytes,4,rep,name=act,proto3" json:"act,omitempty"`
}

func (x *AddPolicyReq) Reset() {
	*x = AddPolicyReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_users_policies_auth_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AddPolicyReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AddPolicyReq) ProtoMessage() {}

func (x *AddPolicyReq) ProtoReflect() protoreflect.Message {
	mi := &file_users_policies_auth_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AddPolicyReq.ProtoReflect.Descriptor instead.
func (*AddPolicyReq) Descriptor() ([]byte, []int) {
	return file_users_policies_auth_proto_rawDescGZIP(), []int{5}
}

func (x *AddPolicyReq) GetToken() string {
	if x != nil {
		return x.Token
	}
	return ""
}

func (x *AddPolicyReq) GetSub() string {
	if x != nil {
		return x.Sub
	}
	return ""
}

func (x *AddPolicyReq) GetObj() string {
	if x != nil {
		return x.Obj
	}
	return ""
}

func (x *AddPolicyReq) GetAct() []string {
	if x != nil {
		return x.Act
	}
	return nil
}

type AddPolicyRes struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Authorized bool `protobuf:"varint,1,opt,name=authorized,proto3" json:"authorized,omitempty"`
}

func (x *AddPolicyRes) Reset() {
	*x = AddPolicyRes{}
	if protoimpl.UnsafeEnabled {
		mi := &file_users_policies_auth_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AddPolicyRes) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AddPolicyRes) ProtoMessage() {}

func (x *AddPolicyRes) ProtoReflect() protoreflect.Message {
	mi := &file_users_policies_auth_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AddPolicyRes.ProtoReflect.Descriptor instead.
func (*AddPolicyRes) Descriptor() ([]byte, []int) {
	return file_users_policies_auth_proto_rawDescGZIP(), []int{6}
}

func (x *AddPolicyRes) GetAuthorized() bool {
	if x != nil {
		return x.Authorized
	}
	return false
}

type DeletePolicyReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Token string `protobuf:"bytes,1,opt,name=token,proto3" json:"token,omitempty"`
	Sub   string `protobuf:"bytes,2,opt,name=sub,proto3" json:"sub,omitempty"`
	Obj   string `protobuf:"bytes,3,opt,name=obj,proto3" json:"obj,omitempty"`
	Act   string `protobuf:"bytes,4,opt,name=act,proto3" json:"act,omitempty"`
}

func (x *DeletePolicyReq) Reset() {
	*x = DeletePolicyReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_users_policies_auth_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DeletePolicyReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeletePolicyReq) ProtoMessage() {}

func (x *DeletePolicyReq) ProtoReflect() protoreflect.Message {
	mi := &file_users_policies_auth_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeletePolicyReq.ProtoReflect.Descriptor instead.
func (*DeletePolicyReq) Descriptor() ([]byte, []int) {
	return file_users_policies_auth_proto_rawDescGZIP(), []int{7}
}

func (x *DeletePolicyReq) GetToken() string {
	if x != nil {
		return x.Token
	}
	return ""
}

func (x *DeletePolicyReq) GetSub() string {
	if x != nil {
		return x.Sub
	}
	return ""
}

func (x *DeletePolicyReq) GetObj() string {
	if x != nil {
		return x.Obj
	}
	return ""
}

func (x *DeletePolicyReq) GetAct() string {
	if x != nil {
		return x.Act
	}
	return ""
}

type DeletePolicyRes struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Deleted bool `protobuf:"varint,1,opt,name=deleted,proto3" json:"deleted,omitempty"`
}

func (x *DeletePolicyRes) Reset() {
	*x = DeletePolicyRes{}
	if protoimpl.UnsafeEnabled {
		mi := &file_users_policies_auth_proto_msgTypes[8]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DeletePolicyRes) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeletePolicyRes) ProtoMessage() {}

func (x *DeletePolicyRes) ProtoReflect() protoreflect.Message {
	mi := &file_users_policies_auth_proto_msgTypes[8]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeletePolicyRes.ProtoReflect.Descriptor instead.
func (*DeletePolicyRes) Descriptor() ([]byte, []int) {
	return file_users_policies_auth_proto_rawDescGZIP(), []int{8}
}

func (x *DeletePolicyRes) GetDeleted() bool {
	if x != nil {
		return x.Deleted
	}
	return false
}

type ListPoliciesReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Token string `protobuf:"bytes,1,opt,name=token,proto3" json:"token,omitempty"`
	Sub   string `protobuf:"bytes,2,opt,name=sub,proto3" json:"sub,omitempty"`
	Obj   string `protobuf:"bytes,3,opt,name=obj,proto3" json:"obj,omitempty"`
	Act   string `protobuf:"bytes,4,opt,name=act,proto3" json:"act,omitempty"`
}

func (x *ListPoliciesReq) Reset() {
	*x = ListPoliciesReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_users_policies_auth_proto_msgTypes[9]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ListPoliciesReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListPoliciesReq) ProtoMessage() {}

func (x *ListPoliciesReq) ProtoReflect() protoreflect.Message {
	mi := &file_users_policies_auth_proto_msgTypes[9]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListPoliciesReq.ProtoReflect.Descriptor instead.
func (*ListPoliciesReq) Descriptor() ([]byte, []int) {
	return file_users_policies_auth_proto_rawDescGZIP(), []int{9}
}

func (x *ListPoliciesReq) GetToken() string {
	if x != nil {
		return x.Token
	}
	return ""
}

func (x *ListPoliciesReq) GetSub() string {
	if x != nil {
		return x.Sub
	}
	return ""
}

func (x *ListPoliciesReq) GetObj() string {
	if x != nil {
		return x.Obj
	}
	return ""
}

func (x *ListPoliciesReq) GetAct() string {
	if x != nil {
		return x.Act
	}
	return ""
}

type ListPoliciesRes struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Objects []string `protobuf:"bytes,1,rep,name=objects,proto3" json:"objects,omitempty"`
}

func (x *ListPoliciesRes) Reset() {
	*x = ListPoliciesRes{}
	if protoimpl.UnsafeEnabled {
		mi := &file_users_policies_auth_proto_msgTypes[10]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ListPoliciesRes) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListPoliciesRes) ProtoMessage() {}

func (x *ListPoliciesRes) ProtoReflect() protoreflect.Message {
	mi := &file_users_policies_auth_proto_msgTypes[10]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListPoliciesRes.ProtoReflect.Descriptor instead.
func (*ListPoliciesRes) Descriptor() ([]byte, []int) {
	return file_users_policies_auth_proto_rawDescGZIP(), []int{10}
}

func (x *ListPoliciesRes) GetObjects() []string {
	if x != nil {
		return x.Objects
	}
	return nil
}

var File_users_policies_auth_proto protoreflect.FileDescriptor

var file_users_policies_auth_proto_rawDesc = []byte{
	0x0a, 0x19, 0x75, 0x73, 0x65, 0x72, 0x73, 0x2f, 0x70, 0x6f, 0x6c, 0x69, 0x63, 0x69, 0x65, 0x73,
	0x2f, 0x61, 0x75, 0x74, 0x68, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x17, 0x6d, 0x61, 0x69,
	0x6e, 0x66, 0x6c, 0x75, 0x78, 0x2e, 0x75, 0x73, 0x65, 0x72, 0x73, 0x2e, 0x70, 0x6f, 0x6c, 0x69,
	0x63, 0x69, 0x65, 0x73, 0x22, 0x64, 0x0a, 0x0c, 0x41, 0x75, 0x74, 0x68, 0x6f, 0x72, 0x69, 0x7a,
	0x65, 0x52, 0x65, 0x71, 0x12, 0x10, 0x0a, 0x03, 0x73, 0x75, 0x62, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x03, 0x73, 0x75, 0x62, 0x12, 0x10, 0x0a, 0x03, 0x6f, 0x62, 0x6a, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x03, 0x6f, 0x62, 0x6a, 0x12, 0x10, 0x0a, 0x03, 0x61, 0x63, 0x74, 0x18,
	0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x61, 0x63, 0x74, 0x12, 0x1e, 0x0a, 0x0a, 0x65, 0x6e,
	0x74, 0x69, 0x74, 0x79, 0x54, 0x79, 0x70, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a,
	0x65, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x54, 0x79, 0x70, 0x65, 0x22, 0x2e, 0x0a, 0x0c, 0x41, 0x75,
	0x74, 0x68, 0x6f, 0x72, 0x69, 0x7a, 0x65, 0x52, 0x65, 0x73, 0x12, 0x1e, 0x0a, 0x0a, 0x61, 0x75,
	0x74, 0x68, 0x6f, 0x72, 0x69, 0x7a, 0x65, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x08, 0x52, 0x0a,
	0x61, 0x75, 0x74, 0x68, 0x6f, 0x72, 0x69, 0x7a, 0x65, 0x64, 0x22, 0x50, 0x0a, 0x08, 0x49, 0x73,
	0x73, 0x75, 0x65, 0x52, 0x65, 0x71, 0x12, 0x14, 0x0a, 0x05, 0x65, 0x6d, 0x61, 0x69, 0x6c, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x65, 0x6d, 0x61, 0x69, 0x6c, 0x12, 0x1a, 0x0a, 0x08,
	0x70, 0x61, 0x73, 0x73, 0x77, 0x6f, 0x72, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08,
	0x70, 0x61, 0x73, 0x73, 0x77, 0x6f, 0x72, 0x64, 0x12, 0x12, 0x0a, 0x04, 0x74, 0x79, 0x70, 0x65,
	0x18, 0x03, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x04, 0x74, 0x79, 0x70, 0x65, 0x22, 0x1d, 0x0a, 0x05,
	0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x12, 0x14, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x22, 0x1e, 0x0a, 0x0c, 0x55,
	0x73, 0x65, 0x72, 0x49, 0x64, 0x65, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x12, 0x0e, 0x0a, 0x02, 0x69,
	0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x22, 0x5a, 0x0a, 0x0c, 0x41,
	0x64, 0x64, 0x50, 0x6f, 0x6c, 0x69, 0x63, 0x79, 0x52, 0x65, 0x71, 0x12, 0x14, 0x0a, 0x05, 0x74,
	0x6f, 0x6b, 0x65, 0x6e, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x74, 0x6f, 0x6b, 0x65,
	0x6e, 0x12, 0x10, 0x0a, 0x03, 0x73, 0x75, 0x62, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03,
	0x73, 0x75, 0x62, 0x12, 0x10, 0x0a, 0x03, 0x6f, 0x62, 0x6a, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x03, 0x6f, 0x62, 0x6a, 0x12, 0x10, 0x0a, 0x03, 0x61, 0x63, 0x74, 0x18, 0x04, 0x20, 0x03,
	0x28, 0x09, 0x52, 0x03, 0x61, 0x63, 0x74, 0x22, 0x2e, 0x0a, 0x0c, 0x41, 0x64, 0x64, 0x50, 0x6f,
	0x6c, 0x69, 0x63, 0x79, 0x52, 0x65, 0x73, 0x12, 0x1e, 0x0a, 0x0a, 0x61, 0x75, 0x74, 0x68, 0x6f,
	0x72, 0x69, 0x7a, 0x65, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x08, 0x52, 0x0a, 0x61, 0x75, 0x74,
	0x68, 0x6f, 0x72, 0x69, 0x7a, 0x65, 0x64, 0x22, 0x5d, 0x0a, 0x0f, 0x44, 0x65, 0x6c, 0x65, 0x74,
	0x65, 0x50, 0x6f, 0x6c, 0x69, 0x63, 0x79, 0x52, 0x65, 0x71, 0x12, 0x14, 0x0a, 0x05, 0x74, 0x6f,
	0x6b, 0x65, 0x6e, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x74, 0x6f, 0x6b, 0x65, 0x6e,
	0x12, 0x10, 0x0a, 0x03, 0x73, 0x75, 0x62, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x73,
	0x75, 0x62, 0x12, 0x10, 0x0a, 0x03, 0x6f, 0x62, 0x6a, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x03, 0x6f, 0x62, 0x6a, 0x12, 0x10, 0x0a, 0x03, 0x61, 0x63, 0x74, 0x18, 0x04, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x03, 0x61, 0x63, 0x74, 0x22, 0x2b, 0x0a, 0x0f, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65,
	0x50, 0x6f, 0x6c, 0x69, 0x63, 0x79, 0x52, 0x65, 0x73, 0x12, 0x18, 0x0a, 0x07, 0x64, 0x65, 0x6c,
	0x65, 0x74, 0x65, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x08, 0x52, 0x07, 0x64, 0x65, 0x6c, 0x65,
	0x74, 0x65, 0x64, 0x22, 0x5d, 0x0a, 0x0f, 0x4c, 0x69, 0x73, 0x74, 0x50, 0x6f, 0x6c, 0x69, 0x63,
	0x69, 0x65, 0x73, 0x52, 0x65, 0x71, 0x12, 0x14, 0x0a, 0x05, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x12, 0x10, 0x0a, 0x03,
	0x73, 0x75, 0x62, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x73, 0x75, 0x62, 0x12, 0x10,
	0x0a, 0x03, 0x6f, 0x62, 0x6a, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x6f, 0x62, 0x6a,
	0x12, 0x10, 0x0a, 0x03, 0x61, 0x63, 0x74, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x61,
	0x63, 0x74, 0x22, 0x2b, 0x0a, 0x0f, 0x4c, 0x69, 0x73, 0x74, 0x50, 0x6f, 0x6c, 0x69, 0x63, 0x69,
	0x65, 0x73, 0x52, 0x65, 0x73, 0x12, 0x18, 0x0a, 0x07, 0x6f, 0x62, 0x6a, 0x65, 0x63, 0x74, 0x73,
	0x18, 0x01, 0x20, 0x03, 0x28, 0x09, 0x52, 0x07, 0x6f, 0x62, 0x6a, 0x65, 0x63, 0x74, 0x73, 0x32,
	0xb6, 0x04, 0x0a, 0x0b, 0x41, 0x75, 0x74, 0x68, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12,
	0x5b, 0x0a, 0x09, 0x41, 0x75, 0x74, 0x68, 0x6f, 0x72, 0x69, 0x7a, 0x65, 0x12, 0x25, 0x2e, 0x6d,
	0x61, 0x69, 0x6e, 0x66, 0x6c, 0x75, 0x78, 0x2e, 0x75, 0x73, 0x65, 0x72, 0x73, 0x2e, 0x70, 0x6f,
	0x6c, 0x69, 0x63, 0x69, 0x65, 0x73, 0x2e, 0x41, 0x75, 0x74, 0x68, 0x6f, 0x72, 0x69, 0x7a, 0x65,
	0x52, 0x65, 0x71, 0x1a, 0x25, 0x2e, 0x6d, 0x61, 0x69, 0x6e, 0x66, 0x6c, 0x75, 0x78, 0x2e, 0x75,
	0x73, 0x65, 0x72, 0x73, 0x2e, 0x70, 0x6f, 0x6c, 0x69, 0x63, 0x69, 0x65, 0x73, 0x2e, 0x41, 0x75,
	0x74, 0x68, 0x6f, 0x72, 0x69, 0x7a, 0x65, 0x52, 0x65, 0x73, 0x22, 0x00, 0x12, 0x4c, 0x0a, 0x05,
	0x49, 0x73, 0x73, 0x75, 0x65, 0x12, 0x21, 0x2e, 0x6d, 0x61, 0x69, 0x6e, 0x66, 0x6c, 0x75, 0x78,
	0x2e, 0x75, 0x73, 0x65, 0x72, 0x73, 0x2e, 0x70, 0x6f, 0x6c, 0x69, 0x63, 0x69, 0x65, 0x73, 0x2e,
	0x49, 0x73, 0x73, 0x75, 0x65, 0x52, 0x65, 0x71, 0x1a, 0x1e, 0x2e, 0x6d, 0x61, 0x69, 0x6e, 0x66,
	0x6c, 0x75, 0x78, 0x2e, 0x75, 0x73, 0x65, 0x72, 0x73, 0x2e, 0x70, 0x6f, 0x6c, 0x69, 0x63, 0x69,
	0x65, 0x73, 0x2e, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x22, 0x00, 0x12, 0x53, 0x0a, 0x08, 0x49, 0x64,
	0x65, 0x6e, 0x74, 0x69, 0x66, 0x79, 0x12, 0x1e, 0x2e, 0x6d, 0x61, 0x69, 0x6e, 0x66, 0x6c, 0x75,
	0x78, 0x2e, 0x75, 0x73, 0x65, 0x72, 0x73, 0x2e, 0x70, 0x6f, 0x6c, 0x69, 0x63, 0x69, 0x65, 0x73,
	0x2e, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x1a, 0x25, 0x2e, 0x6d, 0x61, 0x69, 0x6e, 0x66, 0x6c, 0x75,
	0x78, 0x2e, 0x75, 0x73, 0x65, 0x72, 0x73, 0x2e, 0x70, 0x6f, 0x6c, 0x69, 0x63, 0x69, 0x65, 0x73,
	0x2e, 0x55, 0x73, 0x65, 0x72, 0x49, 0x64, 0x65, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x22, 0x00, 0x12,
	0x5b, 0x0a, 0x09, 0x41, 0x64, 0x64, 0x50, 0x6f, 0x6c, 0x69, 0x63, 0x79, 0x12, 0x25, 0x2e, 0x6d,
	0x61, 0x69, 0x6e, 0x66, 0x6c, 0x75, 0x78, 0x2e, 0x75, 0x73, 0x65, 0x72, 0x73, 0x2e, 0x70, 0x6f,
	0x6c, 0x69, 0x63, 0x69, 0x65, 0x73, 0x2e, 0x41, 0x64, 0x64, 0x50, 0x6f, 0x6c, 0x69, 0x63, 0x79,
	0x52, 0x65, 0x71, 0x1a, 0x25, 0x2e, 0x6d, 0x61, 0x69, 0x6e, 0x66, 0x6c, 0x75, 0x78, 0x2e, 0x75,
	0x73, 0x65, 0x72, 0x73, 0x2e, 0x70, 0x6f, 0x6c, 0x69, 0x63, 0x69, 0x65, 0x73, 0x2e, 0x41, 0x64,
	0x64, 0x50, 0x6f, 0x6c, 0x69, 0x63, 0x79, 0x52, 0x65, 0x73, 0x22, 0x00, 0x12, 0x64, 0x0a, 0x0c,
	0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x50, 0x6f, 0x6c, 0x69, 0x63, 0x79, 0x12, 0x28, 0x2e, 0x6d,
	0x61, 0x69, 0x6e, 0x66, 0x6c, 0x75, 0x78, 0x2e, 0x75, 0x73, 0x65, 0x72, 0x73, 0x2e, 0x70, 0x6f,
	0x6c, 0x69, 0x63, 0x69, 0x65, 0x73, 0x2e, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x50, 0x6f, 0x6c,
	0x69, 0x63, 0x79, 0x52, 0x65, 0x71, 0x1a, 0x28, 0x2e, 0x6d, 0x61, 0x69, 0x6e, 0x66, 0x6c, 0x75,
	0x78, 0x2e, 0x75, 0x73, 0x65, 0x72, 0x73, 0x2e, 0x70, 0x6f, 0x6c, 0x69, 0x63, 0x69, 0x65, 0x73,
	0x2e, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x50, 0x6f, 0x6c, 0x69, 0x63, 0x79, 0x52, 0x65, 0x73,
	0x22, 0x00, 0x12, 0x64, 0x0a, 0x0c, 0x4c, 0x69, 0x73, 0x74, 0x50, 0x6f, 0x6c, 0x69, 0x63, 0x69,
	0x65, 0x73, 0x12, 0x28, 0x2e, 0x6d, 0x61, 0x69, 0x6e, 0x66, 0x6c, 0x75, 0x78, 0x2e, 0x75, 0x73,
	0x65, 0x72, 0x73, 0x2e, 0x70, 0x6f, 0x6c, 0x69, 0x63, 0x69, 0x65, 0x73, 0x2e, 0x4c, 0x69, 0x73,
	0x74, 0x50, 0x6f, 0x6c, 0x69, 0x63, 0x69, 0x65, 0x73, 0x52, 0x65, 0x71, 0x1a, 0x28, 0x2e, 0x6d,
	0x61, 0x69, 0x6e, 0x66, 0x6c, 0x75, 0x78, 0x2e, 0x75, 0x73, 0x65, 0x72, 0x73, 0x2e, 0x70, 0x6f,
	0x6c, 0x69, 0x63, 0x69, 0x65, 0x73, 0x2e, 0x4c, 0x69, 0x73, 0x74, 0x50, 0x6f, 0x6c, 0x69, 0x63,
	0x69, 0x65, 0x73, 0x52, 0x65, 0x73, 0x22, 0x00, 0x42, 0x0c, 0x5a, 0x0a, 0x2e, 0x2f, 0x70, 0x6f,
	0x6c, 0x69, 0x63, 0x69, 0x65, 0x73, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_users_policies_auth_proto_rawDescOnce sync.Once
	file_users_policies_auth_proto_rawDescData = file_users_policies_auth_proto_rawDesc
)

func file_users_policies_auth_proto_rawDescGZIP() []byte {
	file_users_policies_auth_proto_rawDescOnce.Do(func() {
		file_users_policies_auth_proto_rawDescData = protoimpl.X.CompressGZIP(file_users_policies_auth_proto_rawDescData)
	})
	return file_users_policies_auth_proto_rawDescData
}

var file_users_policies_auth_proto_msgTypes = make([]protoimpl.MessageInfo, 11)
var file_users_policies_auth_proto_goTypes = []interface{}{
	(*AuthorizeReq)(nil),    // 0: mainflux.users.policies.AuthorizeReq
	(*AuthorizeRes)(nil),    // 1: mainflux.users.policies.AuthorizeRes
	(*IssueReq)(nil),        // 2: mainflux.users.policies.IssueReq
	(*Token)(nil),           // 3: mainflux.users.policies.Token
	(*UserIdentity)(nil),    // 4: mainflux.users.policies.UserIdentity
	(*AddPolicyReq)(nil),    // 5: mainflux.users.policies.AddPolicyReq
	(*AddPolicyRes)(nil),    // 6: mainflux.users.policies.AddPolicyRes
	(*DeletePolicyReq)(nil), // 7: mainflux.users.policies.DeletePolicyReq
	(*DeletePolicyRes)(nil), // 8: mainflux.users.policies.DeletePolicyRes
	(*ListPoliciesReq)(nil), // 9: mainflux.users.policies.ListPoliciesReq
	(*ListPoliciesRes)(nil), // 10: mainflux.users.policies.ListPoliciesRes
}
var file_users_policies_auth_proto_depIdxs = []int32{
	0,  // 0: mainflux.users.policies.AuthService.Authorize:input_type -> mainflux.users.policies.AuthorizeReq
	2,  // 1: mainflux.users.policies.AuthService.Issue:input_type -> mainflux.users.policies.IssueReq
	3,  // 2: mainflux.users.policies.AuthService.Identify:input_type -> mainflux.users.policies.Token
	5,  // 3: mainflux.users.policies.AuthService.AddPolicy:input_type -> mainflux.users.policies.AddPolicyReq
	7,  // 4: mainflux.users.policies.AuthService.DeletePolicy:input_type -> mainflux.users.policies.DeletePolicyReq
	9,  // 5: mainflux.users.policies.AuthService.ListPolicies:input_type -> mainflux.users.policies.ListPoliciesReq
	1,  // 6: mainflux.users.policies.AuthService.Authorize:output_type -> mainflux.users.policies.AuthorizeRes
	3,  // 7: mainflux.users.policies.AuthService.Issue:output_type -> mainflux.users.policies.Token
	4,  // 8: mainflux.users.policies.AuthService.Identify:output_type -> mainflux.users.policies.UserIdentity
	6,  // 9: mainflux.users.policies.AuthService.AddPolicy:output_type -> mainflux.users.policies.AddPolicyRes
	8,  // 10: mainflux.users.policies.AuthService.DeletePolicy:output_type -> mainflux.users.policies.DeletePolicyRes
	10, // 11: mainflux.users.policies.AuthService.ListPolicies:output_type -> mainflux.users.policies.ListPoliciesRes
	6,  // [6:12] is the sub-list for method output_type
	0,  // [0:6] is the sub-list for method input_type
	0,  // [0:0] is the sub-list for extension type_name
	0,  // [0:0] is the sub-list for extension extendee
	0,  // [0:0] is the sub-list for field type_name
}

func init() { file_users_policies_auth_proto_init() }
func file_users_policies_auth_proto_init() {
	if File_users_policies_auth_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_users_policies_auth_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AuthorizeReq); i {
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
		file_users_policies_auth_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AuthorizeRes); i {
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
		file_users_policies_auth_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*IssueReq); i {
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
		file_users_policies_auth_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Token); i {
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
		file_users_policies_auth_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UserIdentity); i {
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
		file_users_policies_auth_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AddPolicyReq); i {
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
		file_users_policies_auth_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AddPolicyRes); i {
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
		file_users_policies_auth_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DeletePolicyReq); i {
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
		file_users_policies_auth_proto_msgTypes[8].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DeletePolicyRes); i {
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
		file_users_policies_auth_proto_msgTypes[9].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ListPoliciesReq); i {
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
		file_users_policies_auth_proto_msgTypes[10].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ListPoliciesRes); i {
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
			RawDescriptor: file_users_policies_auth_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   11,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_users_policies_auth_proto_goTypes,
		DependencyIndexes: file_users_policies_auth_proto_depIdxs,
		MessageInfos:      file_users_policies_auth_proto_msgTypes,
	}.Build()
	File_users_policies_auth_proto = out.File
	file_users_policies_auth_proto_rawDesc = nil
	file_users_policies_auth_proto_goTypes = nil
	file_users_policies_auth_proto_depIdxs = nil
}
