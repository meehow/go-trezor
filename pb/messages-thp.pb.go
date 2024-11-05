// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.35.1
// 	protoc        v3.21.12
// source: messages-thp.proto

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
// Only for internal use.
// @embed
type ThpCredentialMetadata struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	HostName *string `protobuf:"bytes,1,opt,name=host_name,json=hostName" json:"host_name,omitempty"` // Human-readable host name
}

func (x *ThpCredentialMetadata) Reset() {
	*x = ThpCredentialMetadata{}
	mi := &file_messages_thp_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *ThpCredentialMetadata) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ThpCredentialMetadata) ProtoMessage() {}

func (x *ThpCredentialMetadata) ProtoReflect() protoreflect.Message {
	mi := &file_messages_thp_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ThpCredentialMetadata.ProtoReflect.Descriptor instead.
func (*ThpCredentialMetadata) Descriptor() ([]byte, []int) {
	return file_messages_thp_proto_rawDescGZIP(), []int{0}
}

func (x *ThpCredentialMetadata) GetHostName() string {
	if x != nil && x.HostName != nil {
		return *x.HostName
	}
	return ""
}

// *
// Only for internal use.
// @embed
type ThpPairingCredential struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	CredMetadata *ThpCredentialMetadata `protobuf:"bytes,1,opt,name=cred_metadata,json=credMetadata" json:"cred_metadata,omitempty"` // Credential metadata
	Mac          []byte                 `protobuf:"bytes,2,opt,name=mac" json:"mac,omitempty"`                                       // Message authentication code generated by the Trezor
}

func (x *ThpPairingCredential) Reset() {
	*x = ThpPairingCredential{}
	mi := &file_messages_thp_proto_msgTypes[1]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *ThpPairingCredential) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ThpPairingCredential) ProtoMessage() {}

func (x *ThpPairingCredential) ProtoReflect() protoreflect.Message {
	mi := &file_messages_thp_proto_msgTypes[1]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ThpPairingCredential.ProtoReflect.Descriptor instead.
func (*ThpPairingCredential) Descriptor() ([]byte, []int) {
	return file_messages_thp_proto_rawDescGZIP(), []int{1}
}

func (x *ThpPairingCredential) GetCredMetadata() *ThpCredentialMetadata {
	if x != nil {
		return x.CredMetadata
	}
	return nil
}

func (x *ThpPairingCredential) GetMac() []byte {
	if x != nil {
		return x.Mac
	}
	return nil
}

// *
// Only for internal use.
// @embed
type ThpAuthenticatedCredentialData struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	HostStaticPubkey []byte                 `protobuf:"bytes,1,opt,name=host_static_pubkey,json=hostStaticPubkey" json:"host_static_pubkey,omitempty"` // Host's static public key used in the handshake
	CredMetadata     *ThpCredentialMetadata `protobuf:"bytes,2,opt,name=cred_metadata,json=credMetadata" json:"cred_metadata,omitempty"`               // Credential metadata
}

func (x *ThpAuthenticatedCredentialData) Reset() {
	*x = ThpAuthenticatedCredentialData{}
	mi := &file_messages_thp_proto_msgTypes[2]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *ThpAuthenticatedCredentialData) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ThpAuthenticatedCredentialData) ProtoMessage() {}

func (x *ThpAuthenticatedCredentialData) ProtoReflect() protoreflect.Message {
	mi := &file_messages_thp_proto_msgTypes[2]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ThpAuthenticatedCredentialData.ProtoReflect.Descriptor instead.
func (*ThpAuthenticatedCredentialData) Descriptor() ([]byte, []int) {
	return file_messages_thp_proto_rawDescGZIP(), []int{2}
}

func (x *ThpAuthenticatedCredentialData) GetHostStaticPubkey() []byte {
	if x != nil {
		return x.HostStaticPubkey
	}
	return nil
}

func (x *ThpAuthenticatedCredentialData) GetCredMetadata() *ThpCredentialMetadata {
	if x != nil {
		return x.CredMetadata
	}
	return nil
}

var File_messages_thp_proto protoreflect.FileDescriptor

var file_messages_thp_proto_rawDesc = []byte{
	0x0a, 0x12, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x73, 0x2d, 0x74, 0x68, 0x70, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x12, 0x16, 0x68, 0x77, 0x2e, 0x74, 0x72, 0x65, 0x7a, 0x6f, 0x72, 0x2e,
	0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x73, 0x2e, 0x74, 0x68, 0x70, 0x1a, 0x0d, 0x6f, 0x70,
	0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x3a, 0x0a, 0x15, 0x54,
	0x68, 0x70, 0x43, 0x72, 0x65, 0x64, 0x65, 0x6e, 0x74, 0x69, 0x61, 0x6c, 0x4d, 0x65, 0x74, 0x61,
	0x64, 0x61, 0x74, 0x61, 0x12, 0x1b, 0x0a, 0x09, 0x68, 0x6f, 0x73, 0x74, 0x5f, 0x6e, 0x61, 0x6d,
	0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x68, 0x6f, 0x73, 0x74, 0x4e, 0x61, 0x6d,
	0x65, 0x3a, 0x04, 0x98, 0xb2, 0x19, 0x01, 0x22, 0x82, 0x01, 0x0a, 0x14, 0x54, 0x68, 0x70, 0x50,
	0x61, 0x69, 0x72, 0x69, 0x6e, 0x67, 0x43, 0x72, 0x65, 0x64, 0x65, 0x6e, 0x74, 0x69, 0x61, 0x6c,
	0x12, 0x52, 0x0a, 0x0d, 0x63, 0x72, 0x65, 0x64, 0x5f, 0x6d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74,
	0x61, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x2d, 0x2e, 0x68, 0x77, 0x2e, 0x74, 0x72, 0x65,
	0x7a, 0x6f, 0x72, 0x2e, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x73, 0x2e, 0x74, 0x68, 0x70,
	0x2e, 0x54, 0x68, 0x70, 0x43, 0x72, 0x65, 0x64, 0x65, 0x6e, 0x74, 0x69, 0x61, 0x6c, 0x4d, 0x65,
	0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0x52, 0x0c, 0x63, 0x72, 0x65, 0x64, 0x4d, 0x65, 0x74, 0x61,
	0x64, 0x61, 0x74, 0x61, 0x12, 0x10, 0x0a, 0x03, 0x6d, 0x61, 0x63, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x0c, 0x52, 0x03, 0x6d, 0x61, 0x63, 0x3a, 0x04, 0x98, 0xb2, 0x19, 0x01, 0x22, 0xa8, 0x01, 0x0a,
	0x1e, 0x54, 0x68, 0x70, 0x41, 0x75, 0x74, 0x68, 0x65, 0x6e, 0x74, 0x69, 0x63, 0x61, 0x74, 0x65,
	0x64, 0x43, 0x72, 0x65, 0x64, 0x65, 0x6e, 0x74, 0x69, 0x61, 0x6c, 0x44, 0x61, 0x74, 0x61, 0x12,
	0x2c, 0x0a, 0x12, 0x68, 0x6f, 0x73, 0x74, 0x5f, 0x73, 0x74, 0x61, 0x74, 0x69, 0x63, 0x5f, 0x70,
	0x75, 0x62, 0x6b, 0x65, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x10, 0x68, 0x6f, 0x73,
	0x74, 0x53, 0x74, 0x61, 0x74, 0x69, 0x63, 0x50, 0x75, 0x62, 0x6b, 0x65, 0x79, 0x12, 0x52, 0x0a,
	0x0d, 0x63, 0x72, 0x65, 0x64, 0x5f, 0x6d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x0b, 0x32, 0x2d, 0x2e, 0x68, 0x77, 0x2e, 0x74, 0x72, 0x65, 0x7a, 0x6f, 0x72,
	0x2e, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x73, 0x2e, 0x74, 0x68, 0x70, 0x2e, 0x54, 0x68,
	0x70, 0x43, 0x72, 0x65, 0x64, 0x65, 0x6e, 0x74, 0x69, 0x61, 0x6c, 0x4d, 0x65, 0x74, 0x61, 0x64,
	0x61, 0x74, 0x61, 0x52, 0x0c, 0x63, 0x72, 0x65, 0x64, 0x4d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74,
	0x61, 0x3a, 0x04, 0x98, 0xb2, 0x19, 0x01, 0x42, 0x3b, 0x80, 0xa6, 0x1d, 0x01, 0x0a, 0x23, 0x63,
	0x6f, 0x6d, 0x2e, 0x73, 0x61, 0x74, 0x6f, 0x73, 0x68, 0x69, 0x6c, 0x61, 0x62, 0x73, 0x2e, 0x74,
	0x72, 0x65, 0x7a, 0x6f, 0x72, 0x2e, 0x6c, 0x69, 0x62, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62,
	0x75, 0x66, 0x42, 0x10, 0x54, 0x72, 0x65, 0x7a, 0x6f, 0x72, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67,
	0x65, 0x54, 0x68, 0x70,
}

var (
	file_messages_thp_proto_rawDescOnce sync.Once
	file_messages_thp_proto_rawDescData = file_messages_thp_proto_rawDesc
)

func file_messages_thp_proto_rawDescGZIP() []byte {
	file_messages_thp_proto_rawDescOnce.Do(func() {
		file_messages_thp_proto_rawDescData = protoimpl.X.CompressGZIP(file_messages_thp_proto_rawDescData)
	})
	return file_messages_thp_proto_rawDescData
}

var file_messages_thp_proto_msgTypes = make([]protoimpl.MessageInfo, 3)
var file_messages_thp_proto_goTypes = []any{
	(*ThpCredentialMetadata)(nil),          // 0: hw.trezor.messages.thp.ThpCredentialMetadata
	(*ThpPairingCredential)(nil),           // 1: hw.trezor.messages.thp.ThpPairingCredential
	(*ThpAuthenticatedCredentialData)(nil), // 2: hw.trezor.messages.thp.ThpAuthenticatedCredentialData
}
var file_messages_thp_proto_depIdxs = []int32{
	0, // 0: hw.trezor.messages.thp.ThpPairingCredential.cred_metadata:type_name -> hw.trezor.messages.thp.ThpCredentialMetadata
	0, // 1: hw.trezor.messages.thp.ThpAuthenticatedCredentialData.cred_metadata:type_name -> hw.trezor.messages.thp.ThpCredentialMetadata
	2, // [2:2] is the sub-list for method output_type
	2, // [2:2] is the sub-list for method input_type
	2, // [2:2] is the sub-list for extension type_name
	2, // [2:2] is the sub-list for extension extendee
	0, // [0:2] is the sub-list for field type_name
}

func init() { file_messages_thp_proto_init() }
func file_messages_thp_proto_init() {
	if File_messages_thp_proto != nil {
		return
	}
	file_options_proto_init()
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_messages_thp_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   3,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_messages_thp_proto_goTypes,
		DependencyIndexes: file_messages_thp_proto_depIdxs,
		MessageInfos:      file_messages_thp_proto_msgTypes,
	}.Build()
	File_messages_thp_proto = out.File
	file_messages_thp_proto_rawDesc = nil
	file_messages_thp_proto_goTypes = nil
	file_messages_thp_proto_depIdxs = nil
}