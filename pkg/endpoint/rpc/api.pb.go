// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.31.0
// 	protoc        v4.24.3
// source: resources/proto/api.proto

package rpc

import (
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

type Exception struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Code    uint32   `protobuf:"varint,1,opt,name=code,proto3" json:"code,omitempty"`
	Message string   `protobuf:"bytes,2,opt,name=message,proto3" json:"message,omitempty"`
	Errors  []string `protobuf:"bytes,3,rep,name=errors,proto3" json:"errors,omitempty"`
}

func (x *Exception) Reset() {
	*x = Exception{}
	if protoimpl.UnsafeEnabled {
		mi := &file_resources_proto_api_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Exception) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Exception) ProtoMessage() {}

func (x *Exception) ProtoReflect() protoreflect.Message {
	mi := &file_resources_proto_api_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Exception.ProtoReflect.Descriptor instead.
func (*Exception) Descriptor() ([]byte, []int) {
	return file_resources_proto_api_proto_rawDescGZIP(), []int{0}
}

func (x *Exception) GetCode() uint32 {
	if x != nil {
		return x.Code
	}
	return 0
}

func (x *Exception) GetMessage() string {
	if x != nil {
		return x.Message
	}
	return ""
}

func (x *Exception) GetErrors() []string {
	if x != nil {
		return x.Errors
	}
	return nil
}

type Principal struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Username           string   `protobuf:"bytes,1,opt,name=username,proto3" json:"username,omitempty"`
	Role               string   `protobuf:"bytes,2,opt,name=role,proto3" json:"role,omitempty"`
	Password           string   `protobuf:"bytes,3,opt,name=password,proto3" json:"password,omitempty"`
	Passphrase         string   `protobuf:"bytes,4,opt,name=passphrase,proto3" json:"passphrase,omitempty"`
	Enabled            bool     `protobuf:"varint,5,opt,name=enabled,proto3" json:"enabled,omitempty"`
	NonLocked          bool     `protobuf:"varint,6,opt,name=non_locked,json=nonLocked,proto3" json:"non_locked,omitempty"`
	NonExpired         bool     `protobuf:"varint,7,opt,name=non_expired,json=nonExpired,proto3" json:"non_expired,omitempty"`
	PasswordNonExpired bool     `protobuf:"varint,8,opt,name=password_non_expired,json=passwordNonExpired,proto3" json:"password_non_expired,omitempty"`
	SignupDone         bool     `protobuf:"varint,9,opt,name=signup_done,json=signupDone,proto3" json:"signup_done,omitempty"`
	Resources          []string `protobuf:"bytes,10,rep,name=resources,proto3" json:"resources,omitempty"`
	Token              string   `protobuf:"bytes,11,opt,name=token,proto3" json:"token,omitempty"`
}

func (x *Principal) Reset() {
	*x = Principal{}
	if protoimpl.UnsafeEnabled {
		mi := &file_resources_proto_api_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Principal) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Principal) ProtoMessage() {}

func (x *Principal) ProtoReflect() protoreflect.Message {
	mi := &file_resources_proto_api_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Principal.ProtoReflect.Descriptor instead.
func (*Principal) Descriptor() ([]byte, []int) {
	return file_resources_proto_api_proto_rawDescGZIP(), []int{1}
}

func (x *Principal) GetUsername() string {
	if x != nil {
		return x.Username
	}
	return ""
}

func (x *Principal) GetRole() string {
	if x != nil {
		return x.Role
	}
	return ""
}

func (x *Principal) GetPassword() string {
	if x != nil {
		return x.Password
	}
	return ""
}

func (x *Principal) GetPassphrase() string {
	if x != nil {
		return x.Passphrase
	}
	return ""
}

func (x *Principal) GetEnabled() bool {
	if x != nil {
		return x.Enabled
	}
	return false
}

func (x *Principal) GetNonLocked() bool {
	if x != nil {
		return x.NonLocked
	}
	return false
}

func (x *Principal) GetNonExpired() bool {
	if x != nil {
		return x.NonExpired
	}
	return false
}

func (x *Principal) GetPasswordNonExpired() bool {
	if x != nil {
		return x.PasswordNonExpired
	}
	return false
}

func (x *Principal) GetSignupDone() bool {
	if x != nil {
		return x.SignupDone
	}
	return false
}

func (x *Principal) GetResources() []string {
	if x != nil {
		return x.Resources
	}
	return nil
}

func (x *Principal) GetToken() string {
	if x != nil {
		return x.Token
	}
	return ""
}

type LoginRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Username string `protobuf:"bytes,1,opt,name=username,proto3" json:"username,omitempty"`
	Password string `protobuf:"bytes,3,opt,name=password,proto3" json:"password,omitempty"`
}

func (x *LoginRequest) Reset() {
	*x = LoginRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_resources_proto_api_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *LoginRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*LoginRequest) ProtoMessage() {}

func (x *LoginRequest) ProtoReflect() protoreflect.Message {
	mi := &file_resources_proto_api_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use LoginRequest.ProtoReflect.Descriptor instead.
func (*LoginRequest) Descriptor() ([]byte, []int) {
	return file_resources_proto_api_proto_rawDescGZIP(), []int{2}
}

func (x *LoginRequest) GetUsername() string {
	if x != nil {
		return x.Username
	}
	return ""
}

func (x *LoginRequest) GetPassword() string {
	if x != nil {
		return x.Password
	}
	return ""
}

type LoginResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Username  string   `protobuf:"bytes,1,opt,name=username,proto3" json:"username,omitempty"`
	Role      string   `protobuf:"bytes,2,opt,name=role,proto3" json:"role,omitempty"`
	Resources []string `protobuf:"bytes,10,rep,name=resources,proto3" json:"resources,omitempty"`
	Token     string   `protobuf:"bytes,11,opt,name=token,proto3" json:"token,omitempty"`
}

func (x *LoginResponse) Reset() {
	*x = LoginResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_resources_proto_api_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *LoginResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*LoginResponse) ProtoMessage() {}

func (x *LoginResponse) ProtoReflect() protoreflect.Message {
	mi := &file_resources_proto_api_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use LoginResponse.ProtoReflect.Descriptor instead.
func (*LoginResponse) Descriptor() ([]byte, []int) {
	return file_resources_proto_api_proto_rawDescGZIP(), []int{3}
}

func (x *LoginResponse) GetUsername() string {
	if x != nil {
		return x.Username
	}
	return ""
}

func (x *LoginResponse) GetRole() string {
	if x != nil {
		return x.Role
	}
	return ""
}

func (x *LoginResponse) GetResources() []string {
	if x != nil {
		return x.Resources
	}
	return nil
}

func (x *LoginResponse) GetToken() string {
	if x != nil {
		return x.Token
	}
	return ""
}

var File_resources_proto_api_proto protoreflect.FileDescriptor

var file_resources_proto_api_proto_rawDesc = []byte{
	0x0a, 0x19, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x73, 0x2f, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x2f, 0x61, 0x70, 0x69, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1b, 0x67, 0x6f, 0x6f,
	0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x65, 0x6d, 0x70,
	0x74, 0x79, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x51, 0x0a, 0x09, 0x45, 0x78, 0x63, 0x65,
	0x70, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x12, 0x0a, 0x04, 0x63, 0x6f, 0x64, 0x65, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x0d, 0x52, 0x04, 0x63, 0x6f, 0x64, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x6d, 0x65, 0x73,
	0x73, 0x61, 0x67, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x6d, 0x65, 0x73, 0x73,
	0x61, 0x67, 0x65, 0x12, 0x16, 0x0a, 0x06, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x73, 0x18, 0x03, 0x20,
	0x03, 0x28, 0x09, 0x52, 0x06, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x73, 0x22, 0xd8, 0x02, 0x0a, 0x09,
	0x50, 0x72, 0x69, 0x6e, 0x63, 0x69, 0x70, 0x61, 0x6c, 0x12, 0x1a, 0x0a, 0x08, 0x75, 0x73, 0x65,
	0x72, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x75, 0x73, 0x65,
	0x72, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x72, 0x6f, 0x6c, 0x65, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x04, 0x72, 0x6f, 0x6c, 0x65, 0x12, 0x1a, 0x0a, 0x08, 0x70, 0x61, 0x73,
	0x73, 0x77, 0x6f, 0x72, 0x64, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x70, 0x61, 0x73,
	0x73, 0x77, 0x6f, 0x72, 0x64, 0x12, 0x1e, 0x0a, 0x0a, 0x70, 0x61, 0x73, 0x73, 0x70, 0x68, 0x72,
	0x61, 0x73, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x70, 0x61, 0x73, 0x73, 0x70,
	0x68, 0x72, 0x61, 0x73, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x65, 0x6e, 0x61, 0x62, 0x6c, 0x65, 0x64,
	0x18, 0x05, 0x20, 0x01, 0x28, 0x08, 0x52, 0x07, 0x65, 0x6e, 0x61, 0x62, 0x6c, 0x65, 0x64, 0x12,
	0x1d, 0x0a, 0x0a, 0x6e, 0x6f, 0x6e, 0x5f, 0x6c, 0x6f, 0x63, 0x6b, 0x65, 0x64, 0x18, 0x06, 0x20,
	0x01, 0x28, 0x08, 0x52, 0x09, 0x6e, 0x6f, 0x6e, 0x4c, 0x6f, 0x63, 0x6b, 0x65, 0x64, 0x12, 0x1f,
	0x0a, 0x0b, 0x6e, 0x6f, 0x6e, 0x5f, 0x65, 0x78, 0x70, 0x69, 0x72, 0x65, 0x64, 0x18, 0x07, 0x20,
	0x01, 0x28, 0x08, 0x52, 0x0a, 0x6e, 0x6f, 0x6e, 0x45, 0x78, 0x70, 0x69, 0x72, 0x65, 0x64, 0x12,
	0x30, 0x0a, 0x14, 0x70, 0x61, 0x73, 0x73, 0x77, 0x6f, 0x72, 0x64, 0x5f, 0x6e, 0x6f, 0x6e, 0x5f,
	0x65, 0x78, 0x70, 0x69, 0x72, 0x65, 0x64, 0x18, 0x08, 0x20, 0x01, 0x28, 0x08, 0x52, 0x12, 0x70,
	0x61, 0x73, 0x73, 0x77, 0x6f, 0x72, 0x64, 0x4e, 0x6f, 0x6e, 0x45, 0x78, 0x70, 0x69, 0x72, 0x65,
	0x64, 0x12, 0x1f, 0x0a, 0x0b, 0x73, 0x69, 0x67, 0x6e, 0x75, 0x70, 0x5f, 0x64, 0x6f, 0x6e, 0x65,
	0x18, 0x09, 0x20, 0x01, 0x28, 0x08, 0x52, 0x0a, 0x73, 0x69, 0x67, 0x6e, 0x75, 0x70, 0x44, 0x6f,
	0x6e, 0x65, 0x12, 0x1c, 0x0a, 0x09, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x73, 0x18,
	0x0a, 0x20, 0x03, 0x28, 0x09, 0x52, 0x09, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x73,
	0x12, 0x14, 0x0a, 0x05, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x18, 0x0b, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x05, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x22, 0x46, 0x0a, 0x0c, 0x4c, 0x6f, 0x67, 0x69, 0x6e, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x1a, 0x0a, 0x08, 0x75, 0x73, 0x65, 0x72, 0x6e, 0x61,
	0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x75, 0x73, 0x65, 0x72, 0x6e, 0x61,
	0x6d, 0x65, 0x12, 0x1a, 0x0a, 0x08, 0x70, 0x61, 0x73, 0x73, 0x77, 0x6f, 0x72, 0x64, 0x18, 0x03,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x70, 0x61, 0x73, 0x73, 0x77, 0x6f, 0x72, 0x64, 0x22, 0x73,
	0x0a, 0x0d, 0x4c, 0x6f, 0x67, 0x69, 0x6e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12,
	0x1a, 0x0a, 0x08, 0x75, 0x73, 0x65, 0x72, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x08, 0x75, 0x73, 0x65, 0x72, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x72,
	0x6f, 0x6c, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x72, 0x6f, 0x6c, 0x65, 0x12,
	0x1c, 0x0a, 0x09, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x73, 0x18, 0x0a, 0x20, 0x03,
	0x28, 0x09, 0x52, 0x09, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x73, 0x12, 0x14, 0x0a,
	0x05, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x18, 0x0b, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x74, 0x6f,
	0x6b, 0x65, 0x6e, 0x32, 0x67, 0x0a, 0x09, 0x41, 0x70, 0x69, 0x53, 0x61, 0x6d, 0x70, 0x6c, 0x65,
	0x12, 0x26, 0x0a, 0x05, 0x4c, 0x6f, 0x67, 0x69, 0x6e, 0x12, 0x0d, 0x2e, 0x4c, 0x6f, 0x67, 0x69,
	0x6e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x0e, 0x2e, 0x4c, 0x6f, 0x67, 0x69, 0x6e,
	0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x32, 0x0a, 0x0c, 0x47, 0x65, 0x74, 0x50,
	0x72, 0x69, 0x6e, 0x63, 0x69, 0x70, 0x61, 0x6c, 0x12, 0x16, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c,
	0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79,
	0x1a, 0x0a, 0x2e, 0x50, 0x72, 0x69, 0x6e, 0x63, 0x69, 0x70, 0x61, 0x6c, 0x42, 0x14, 0x5a, 0x12,
	0x2e, 0x2f, 0x70, 0x6b, 0x67, 0x2f, 0x65, 0x6e, 0x64, 0x70, 0x6f, 0x69, 0x6e, 0x74, 0x2f, 0x72,
	0x70, 0x63, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_resources_proto_api_proto_rawDescOnce sync.Once
	file_resources_proto_api_proto_rawDescData = file_resources_proto_api_proto_rawDesc
)

func file_resources_proto_api_proto_rawDescGZIP() []byte {
	file_resources_proto_api_proto_rawDescOnce.Do(func() {
		file_resources_proto_api_proto_rawDescData = protoimpl.X.CompressGZIP(file_resources_proto_api_proto_rawDescData)
	})
	return file_resources_proto_api_proto_rawDescData
}

var file_resources_proto_api_proto_msgTypes = make([]protoimpl.MessageInfo, 4)
var file_resources_proto_api_proto_goTypes = []interface{}{
	(*Exception)(nil),     // 0: Exception
	(*Principal)(nil),     // 1: Principal
	(*LoginRequest)(nil),  // 2: LoginRequest
	(*LoginResponse)(nil), // 3: LoginResponse
	(*emptypb.Empty)(nil), // 4: google.protobuf.Empty
}
var file_resources_proto_api_proto_depIdxs = []int32{
	2, // 0: ApiSample.Login:input_type -> LoginRequest
	4, // 1: ApiSample.GetPrincipal:input_type -> google.protobuf.Empty
	3, // 2: ApiSample.Login:output_type -> LoginResponse
	1, // 3: ApiSample.GetPrincipal:output_type -> Principal
	2, // [2:4] is the sub-list for method output_type
	0, // [0:2] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_resources_proto_api_proto_init() }
func file_resources_proto_api_proto_init() {
	if File_resources_proto_api_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_resources_proto_api_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Exception); i {
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
		file_resources_proto_api_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Principal); i {
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
		file_resources_proto_api_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*LoginRequest); i {
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
		file_resources_proto_api_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*LoginResponse); i {
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
			RawDescriptor: file_resources_proto_api_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   4,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_resources_proto_api_proto_goTypes,
		DependencyIndexes: file_resources_proto_api_proto_depIdxs,
		MessageInfos:      file_resources_proto_api_proto_msgTypes,
	}.Build()
	File_resources_proto_api_proto = out.File
	file_resources_proto_api_proto_rawDesc = nil
	file_resources_proto_api_proto_goTypes = nil
	file_resources_proto_api_proto_depIdxs = nil
}
