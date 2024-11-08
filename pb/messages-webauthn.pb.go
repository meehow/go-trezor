// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.35.1
// 	protoc        v3.21.12
// source: messages-webauthn.proto

package pb

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

// *
// Request: List resident credentials
// @start
// @next WebAuthnCredentials
// @next Failure
type WebAuthnListResidentCredentials struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *WebAuthnListResidentCredentials) Reset() {
	*x = WebAuthnListResidentCredentials{}
	mi := &file_messages_webauthn_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *WebAuthnListResidentCredentials) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*WebAuthnListResidentCredentials) ProtoMessage() {}

func (x *WebAuthnListResidentCredentials) ProtoReflect() protoreflect.Message {
	mi := &file_messages_webauthn_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use WebAuthnListResidentCredentials.ProtoReflect.Descriptor instead.
func (*WebAuthnListResidentCredentials) Descriptor() ([]byte, []int) {
	return file_messages_webauthn_proto_rawDescGZIP(), []int{0}
}

// *
// Request: Add resident credential
// @start
// @next Success
// @next Failure
type WebAuthnAddResidentCredential struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	CredentialId []byte `protobuf:"bytes,1,opt,name=credential_id,json=credentialId" json:"credential_id,omitempty"`
}

func (x *WebAuthnAddResidentCredential) Reset() {
	*x = WebAuthnAddResidentCredential{}
	mi := &file_messages_webauthn_proto_msgTypes[1]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *WebAuthnAddResidentCredential) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*WebAuthnAddResidentCredential) ProtoMessage() {}

func (x *WebAuthnAddResidentCredential) ProtoReflect() protoreflect.Message {
	mi := &file_messages_webauthn_proto_msgTypes[1]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use WebAuthnAddResidentCredential.ProtoReflect.Descriptor instead.
func (*WebAuthnAddResidentCredential) Descriptor() ([]byte, []int) {
	return file_messages_webauthn_proto_rawDescGZIP(), []int{1}
}

func (x *WebAuthnAddResidentCredential) GetCredentialId() []byte {
	if x != nil {
		return x.CredentialId
	}
	return nil
}

// *
// Request: Remove resident credential
// @start
// @next Success
// @next Failure
type WebAuthnRemoveResidentCredential struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Index *uint32 `protobuf:"varint,1,opt,name=index" json:"index,omitempty"`
}

func (x *WebAuthnRemoveResidentCredential) Reset() {
	*x = WebAuthnRemoveResidentCredential{}
	mi := &file_messages_webauthn_proto_msgTypes[2]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *WebAuthnRemoveResidentCredential) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*WebAuthnRemoveResidentCredential) ProtoMessage() {}

func (x *WebAuthnRemoveResidentCredential) ProtoReflect() protoreflect.Message {
	mi := &file_messages_webauthn_proto_msgTypes[2]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use WebAuthnRemoveResidentCredential.ProtoReflect.Descriptor instead.
func (*WebAuthnRemoveResidentCredential) Descriptor() ([]byte, []int) {
	return file_messages_webauthn_proto_rawDescGZIP(), []int{2}
}

func (x *WebAuthnRemoveResidentCredential) GetIndex() uint32 {
	if x != nil && x.Index != nil {
		return *x.Index
	}
	return 0
}

// *
// Response: Resident credential list
// @start
// @next end
type WebAuthnCredentials struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Credentials []*WebAuthnCredentials_WebAuthnCredential `protobuf:"bytes,1,rep,name=credentials" json:"credentials,omitempty"`
}

func (x *WebAuthnCredentials) Reset() {
	*x = WebAuthnCredentials{}
	mi := &file_messages_webauthn_proto_msgTypes[3]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *WebAuthnCredentials) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*WebAuthnCredentials) ProtoMessage() {}

func (x *WebAuthnCredentials) ProtoReflect() protoreflect.Message {
	mi := &file_messages_webauthn_proto_msgTypes[3]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use WebAuthnCredentials.ProtoReflect.Descriptor instead.
func (*WebAuthnCredentials) Descriptor() ([]byte, []int) {
	return file_messages_webauthn_proto_rawDescGZIP(), []int{3}
}

func (x *WebAuthnCredentials) GetCredentials() []*WebAuthnCredentials_WebAuthnCredential {
	if x != nil {
		return x.Credentials
	}
	return nil
}

type WebAuthnCredentials_WebAuthnCredential struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Index           *uint32 `protobuf:"varint,1,opt,name=index" json:"index,omitempty"`
	Id              []byte  `protobuf:"bytes,2,opt,name=id" json:"id,omitempty"`
	RpId            *string `protobuf:"bytes,3,opt,name=rp_id,json=rpId" json:"rp_id,omitempty"`
	RpName          *string `protobuf:"bytes,4,opt,name=rp_name,json=rpName" json:"rp_name,omitempty"`
	UserId          []byte  `protobuf:"bytes,5,opt,name=user_id,json=userId" json:"user_id,omitempty"`
	UserName        *string `protobuf:"bytes,6,opt,name=user_name,json=userName" json:"user_name,omitempty"`
	UserDisplayName *string `protobuf:"bytes,7,opt,name=user_display_name,json=userDisplayName" json:"user_display_name,omitempty"`
	CreationTime    *uint32 `protobuf:"varint,8,opt,name=creation_time,json=creationTime" json:"creation_time,omitempty"`
	HmacSecret      *bool   `protobuf:"varint,9,opt,name=hmac_secret,json=hmacSecret" json:"hmac_secret,omitempty"`
	UseSignCount    *bool   `protobuf:"varint,10,opt,name=use_sign_count,json=useSignCount" json:"use_sign_count,omitempty"`
	Algorithm       *int32  `protobuf:"zigzag32,11,opt,name=algorithm" json:"algorithm,omitempty"`
	Curve           *int32  `protobuf:"zigzag32,12,opt,name=curve" json:"curve,omitempty"`
}

func (x *WebAuthnCredentials_WebAuthnCredential) Reset() {
	*x = WebAuthnCredentials_WebAuthnCredential{}
	mi := &file_messages_webauthn_proto_msgTypes[4]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *WebAuthnCredentials_WebAuthnCredential) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*WebAuthnCredentials_WebAuthnCredential) ProtoMessage() {}

func (x *WebAuthnCredentials_WebAuthnCredential) ProtoReflect() protoreflect.Message {
	mi := &file_messages_webauthn_proto_msgTypes[4]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use WebAuthnCredentials_WebAuthnCredential.ProtoReflect.Descriptor instead.
func (*WebAuthnCredentials_WebAuthnCredential) Descriptor() ([]byte, []int) {
	return file_messages_webauthn_proto_rawDescGZIP(), []int{3, 0}
}

func (x *WebAuthnCredentials_WebAuthnCredential) GetIndex() uint32 {
	if x != nil && x.Index != nil {
		return *x.Index
	}
	return 0
}

func (x *WebAuthnCredentials_WebAuthnCredential) GetId() []byte {
	if x != nil {
		return x.Id
	}
	return nil
}

func (x *WebAuthnCredentials_WebAuthnCredential) GetRpId() string {
	if x != nil && x.RpId != nil {
		return *x.RpId
	}
	return ""
}

func (x *WebAuthnCredentials_WebAuthnCredential) GetRpName() string {
	if x != nil && x.RpName != nil {
		return *x.RpName
	}
	return ""
}

func (x *WebAuthnCredentials_WebAuthnCredential) GetUserId() []byte {
	if x != nil {
		return x.UserId
	}
	return nil
}

func (x *WebAuthnCredentials_WebAuthnCredential) GetUserName() string {
	if x != nil && x.UserName != nil {
		return *x.UserName
	}
	return ""
}

func (x *WebAuthnCredentials_WebAuthnCredential) GetUserDisplayName() string {
	if x != nil && x.UserDisplayName != nil {
		return *x.UserDisplayName
	}
	return ""
}

func (x *WebAuthnCredentials_WebAuthnCredential) GetCreationTime() uint32 {
	if x != nil && x.CreationTime != nil {
		return *x.CreationTime
	}
	return 0
}

func (x *WebAuthnCredentials_WebAuthnCredential) GetHmacSecret() bool {
	if x != nil && x.HmacSecret != nil {
		return *x.HmacSecret
	}
	return false
}

func (x *WebAuthnCredentials_WebAuthnCredential) GetUseSignCount() bool {
	if x != nil && x.UseSignCount != nil {
		return *x.UseSignCount
	}
	return false
}

func (x *WebAuthnCredentials_WebAuthnCredential) GetAlgorithm() int32 {
	if x != nil && x.Algorithm != nil {
		return *x.Algorithm
	}
	return 0
}

func (x *WebAuthnCredentials_WebAuthnCredential) GetCurve() int32 {
	if x != nil && x.Curve != nil {
		return *x.Curve
	}
	return 0
}

var File_messages_webauthn_proto protoreflect.FileDescriptor

var file_messages_webauthn_proto_rawDesc = []byte{
	0x0a, 0x17, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x73, 0x2d, 0x77, 0x65, 0x62, 0x61, 0x75,
	0x74, 0x68, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x1b, 0x68, 0x77, 0x2e, 0x74, 0x72,
	0x65, 0x7a, 0x6f, 0x72, 0x2e, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x73, 0x2e, 0x77, 0x65,
	0x62, 0x61, 0x75, 0x74, 0x68, 0x6e, 0x22, 0x21, 0x0a, 0x1f, 0x57, 0x65, 0x62, 0x41, 0x75, 0x74,
	0x68, 0x6e, 0x4c, 0x69, 0x73, 0x74, 0x52, 0x65, 0x73, 0x69, 0x64, 0x65, 0x6e, 0x74, 0x43, 0x72,
	0x65, 0x64, 0x65, 0x6e, 0x74, 0x69, 0x61, 0x6c, 0x73, 0x22, 0x44, 0x0a, 0x1d, 0x57, 0x65, 0x62,
	0x41, 0x75, 0x74, 0x68, 0x6e, 0x41, 0x64, 0x64, 0x52, 0x65, 0x73, 0x69, 0x64, 0x65, 0x6e, 0x74,
	0x43, 0x72, 0x65, 0x64, 0x65, 0x6e, 0x74, 0x69, 0x61, 0x6c, 0x12, 0x23, 0x0a, 0x0d, 0x63, 0x72,
	0x65, 0x64, 0x65, 0x6e, 0x74, 0x69, 0x61, 0x6c, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x0c, 0x52, 0x0c, 0x63, 0x72, 0x65, 0x64, 0x65, 0x6e, 0x74, 0x69, 0x61, 0x6c, 0x49, 0x64, 0x22,
	0x38, 0x0a, 0x20, 0x57, 0x65, 0x62, 0x41, 0x75, 0x74, 0x68, 0x6e, 0x52, 0x65, 0x6d, 0x6f, 0x76,
	0x65, 0x52, 0x65, 0x73, 0x69, 0x64, 0x65, 0x6e, 0x74, 0x43, 0x72, 0x65, 0x64, 0x65, 0x6e, 0x74,
	0x69, 0x61, 0x6c, 0x12, 0x14, 0x0a, 0x05, 0x69, 0x6e, 0x64, 0x65, 0x78, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x0d, 0x52, 0x05, 0x69, 0x6e, 0x64, 0x65, 0x78, 0x22, 0xe9, 0x03, 0x0a, 0x13, 0x57, 0x65,
	0x62, 0x41, 0x75, 0x74, 0x68, 0x6e, 0x43, 0x72, 0x65, 0x64, 0x65, 0x6e, 0x74, 0x69, 0x61, 0x6c,
	0x73, 0x12, 0x65, 0x0a, 0x0b, 0x63, 0x72, 0x65, 0x64, 0x65, 0x6e, 0x74, 0x69, 0x61, 0x6c, 0x73,
	0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x43, 0x2e, 0x68, 0x77, 0x2e, 0x74, 0x72, 0x65, 0x7a,
	0x6f, 0x72, 0x2e, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x73, 0x2e, 0x77, 0x65, 0x62, 0x61,
	0x75, 0x74, 0x68, 0x6e, 0x2e, 0x57, 0x65, 0x62, 0x41, 0x75, 0x74, 0x68, 0x6e, 0x43, 0x72, 0x65,
	0x64, 0x65, 0x6e, 0x74, 0x69, 0x61, 0x6c, 0x73, 0x2e, 0x57, 0x65, 0x62, 0x41, 0x75, 0x74, 0x68,
	0x6e, 0x43, 0x72, 0x65, 0x64, 0x65, 0x6e, 0x74, 0x69, 0x61, 0x6c, 0x52, 0x0b, 0x63, 0x72, 0x65,
	0x64, 0x65, 0x6e, 0x74, 0x69, 0x61, 0x6c, 0x73, 0x1a, 0xea, 0x02, 0x0a, 0x12, 0x57, 0x65, 0x62,
	0x41, 0x75, 0x74, 0x68, 0x6e, 0x43, 0x72, 0x65, 0x64, 0x65, 0x6e, 0x74, 0x69, 0x61, 0x6c, 0x12,
	0x14, 0x0a, 0x05, 0x69, 0x6e, 0x64, 0x65, 0x78, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x05,
	0x69, 0x6e, 0x64, 0x65, 0x78, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x0c, 0x52, 0x02, 0x69, 0x64, 0x12, 0x13, 0x0a, 0x05, 0x72, 0x70, 0x5f, 0x69, 0x64, 0x18, 0x03,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x72, 0x70, 0x49, 0x64, 0x12, 0x17, 0x0a, 0x07, 0x72, 0x70,
	0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x72, 0x70, 0x4e,
	0x61, 0x6d, 0x65, 0x12, 0x17, 0x0a, 0x07, 0x75, 0x73, 0x65, 0x72, 0x5f, 0x69, 0x64, 0x18, 0x05,
	0x20, 0x01, 0x28, 0x0c, 0x52, 0x06, 0x75, 0x73, 0x65, 0x72, 0x49, 0x64, 0x12, 0x1b, 0x0a, 0x09,
	0x75, 0x73, 0x65, 0x72, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x06, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x08, 0x75, 0x73, 0x65, 0x72, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x2a, 0x0a, 0x11, 0x75, 0x73, 0x65,
	0x72, 0x5f, 0x64, 0x69, 0x73, 0x70, 0x6c, 0x61, 0x79, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x07,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x0f, 0x75, 0x73, 0x65, 0x72, 0x44, 0x69, 0x73, 0x70, 0x6c, 0x61,
	0x79, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x23, 0x0a, 0x0d, 0x63, 0x72, 0x65, 0x61, 0x74, 0x69, 0x6f,
	0x6e, 0x5f, 0x74, 0x69, 0x6d, 0x65, 0x18, 0x08, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x0c, 0x63, 0x72,
	0x65, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x54, 0x69, 0x6d, 0x65, 0x12, 0x1f, 0x0a, 0x0b, 0x68, 0x6d,
	0x61, 0x63, 0x5f, 0x73, 0x65, 0x63, 0x72, 0x65, 0x74, 0x18, 0x09, 0x20, 0x01, 0x28, 0x08, 0x52,
	0x0a, 0x68, 0x6d, 0x61, 0x63, 0x53, 0x65, 0x63, 0x72, 0x65, 0x74, 0x12, 0x24, 0x0a, 0x0e, 0x75,
	0x73, 0x65, 0x5f, 0x73, 0x69, 0x67, 0x6e, 0x5f, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x18, 0x0a, 0x20,
	0x01, 0x28, 0x08, 0x52, 0x0c, 0x75, 0x73, 0x65, 0x53, 0x69, 0x67, 0x6e, 0x43, 0x6f, 0x75, 0x6e,
	0x74, 0x12, 0x1c, 0x0a, 0x09, 0x61, 0x6c, 0x67, 0x6f, 0x72, 0x69, 0x74, 0x68, 0x6d, 0x18, 0x0b,
	0x20, 0x01, 0x28, 0x11, 0x52, 0x09, 0x61, 0x6c, 0x67, 0x6f, 0x72, 0x69, 0x74, 0x68, 0x6d, 0x12,
	0x14, 0x0a, 0x05, 0x63, 0x75, 0x72, 0x76, 0x65, 0x18, 0x0c, 0x20, 0x01, 0x28, 0x11, 0x52, 0x05,
	0x63, 0x75, 0x72, 0x76, 0x65, 0x42, 0x3c, 0x0a, 0x23, 0x63, 0x6f, 0x6d, 0x2e, 0x73, 0x61, 0x74,
	0x6f, 0x73, 0x68, 0x69, 0x6c, 0x61, 0x62, 0x73, 0x2e, 0x74, 0x72, 0x65, 0x7a, 0x6f, 0x72, 0x2e,
	0x6c, 0x69, 0x62, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x42, 0x15, 0x54, 0x72,
	0x65, 0x7a, 0x6f, 0x72, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x57, 0x65, 0x62, 0x41, 0x75,
	0x74, 0x68, 0x6e,
}

var (
	file_messages_webauthn_proto_rawDescOnce sync.Once
	file_messages_webauthn_proto_rawDescData = file_messages_webauthn_proto_rawDesc
)

func file_messages_webauthn_proto_rawDescGZIP() []byte {
	file_messages_webauthn_proto_rawDescOnce.Do(func() {
		file_messages_webauthn_proto_rawDescData = protoimpl.X.CompressGZIP(file_messages_webauthn_proto_rawDescData)
	})
	return file_messages_webauthn_proto_rawDescData
}

var file_messages_webauthn_proto_msgTypes = make([]protoimpl.MessageInfo, 5)
var file_messages_webauthn_proto_goTypes = []any{
	(*WebAuthnListResidentCredentials)(nil),        // 0: hw.trezor.messages.webauthn.WebAuthnListResidentCredentials
	(*WebAuthnAddResidentCredential)(nil),          // 1: hw.trezor.messages.webauthn.WebAuthnAddResidentCredential
	(*WebAuthnRemoveResidentCredential)(nil),       // 2: hw.trezor.messages.webauthn.WebAuthnRemoveResidentCredential
	(*WebAuthnCredentials)(nil),                    // 3: hw.trezor.messages.webauthn.WebAuthnCredentials
	(*WebAuthnCredentials_WebAuthnCredential)(nil), // 4: hw.trezor.messages.webauthn.WebAuthnCredentials.WebAuthnCredential
}
var file_messages_webauthn_proto_depIdxs = []int32{
	4, // 0: hw.trezor.messages.webauthn.WebAuthnCredentials.credentials:type_name -> hw.trezor.messages.webauthn.WebAuthnCredentials.WebAuthnCredential
	1, // [1:1] is the sub-list for method output_type
	1, // [1:1] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_messages_webauthn_proto_init() }
func file_messages_webauthn_proto_init() {
	if File_messages_webauthn_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_messages_webauthn_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   5,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_messages_webauthn_proto_goTypes,
		DependencyIndexes: file_messages_webauthn_proto_depIdxs,
		MessageInfos:      file_messages_webauthn_proto_msgTypes,
	}.Build()
	File_messages_webauthn_proto = out.File
	file_messages_webauthn_proto_rawDesc = nil
	file_messages_webauthn_proto_goTypes = nil
	file_messages_webauthn_proto_depIdxs = nil
}
