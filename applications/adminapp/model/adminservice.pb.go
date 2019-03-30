// Code generated by protoc-gen-go. DO NOT EDIT.
// source: proto/adminservice.proto

package model

import (
	context "context"
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	_ "google.golang.org/genproto/googleapis/api/annotations"
	grpc "google.golang.org/grpc"
	math "math"
	common "theterriblechild/CloudApp/common"
	model "theterriblechild/CloudApp/common/model"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion3 // please upgrade the proto package

type CreateAccountMessage struct {
	Name                 string             `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	CreateUser           *CreateUserMessage `protobuf:"bytes,2,opt,name=create_user,json=createUser,proto3" json:"create_user,omitempty"`
	XXX_NoUnkeyedLiteral struct{}           `json:"-"`
	XXX_unrecognized     []byte             `json:"-"`
	XXX_sizecache        int32              `json:"-"`
}

func (m *CreateAccountMessage) Reset()         { *m = CreateAccountMessage{} }
func (m *CreateAccountMessage) String() string { return proto.CompactTextString(m) }
func (*CreateAccountMessage) ProtoMessage()    {}
func (*CreateAccountMessage) Descriptor() ([]byte, []int) {
	return fileDescriptor_96292a55e4fedcbc, []int{0}
}

func (m *CreateAccountMessage) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CreateAccountMessage.Unmarshal(m, b)
}
func (m *CreateAccountMessage) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CreateAccountMessage.Marshal(b, m, deterministic)
}
func (m *CreateAccountMessage) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CreateAccountMessage.Merge(m, src)
}
func (m *CreateAccountMessage) XXX_Size() int {
	return xxx_messageInfo_CreateAccountMessage.Size(m)
}
func (m *CreateAccountMessage) XXX_DiscardUnknown() {
	xxx_messageInfo_CreateAccountMessage.DiscardUnknown(m)
}

var xxx_messageInfo_CreateAccountMessage proto.InternalMessageInfo

func (m *CreateAccountMessage) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *CreateAccountMessage) GetCreateUser() *CreateUserMessage {
	if m != nil {
		return m.CreateUser
	}
	return nil
}

type CreateUserMessage struct {
	UserCreationToken    string   `protobuf:"bytes,1,opt,name=user_creation_token,json=userCreationToken,proto3" json:"user_creation_token,omitempty"`
	Password             string   `protobuf:"bytes,2,opt,name=password,proto3" json:"password,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *CreateUserMessage) Reset()         { *m = CreateUserMessage{} }
func (m *CreateUserMessage) String() string { return proto.CompactTextString(m) }
func (*CreateUserMessage) ProtoMessage()    {}
func (*CreateUserMessage) Descriptor() ([]byte, []int) {
	return fileDescriptor_96292a55e4fedcbc, []int{1}
}

func (m *CreateUserMessage) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CreateUserMessage.Unmarshal(m, b)
}
func (m *CreateUserMessage) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CreateUserMessage.Marshal(b, m, deterministic)
}
func (m *CreateUserMessage) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CreateUserMessage.Merge(m, src)
}
func (m *CreateUserMessage) XXX_Size() int {
	return xxx_messageInfo_CreateUserMessage.Size(m)
}
func (m *CreateUserMessage) XXX_DiscardUnknown() {
	xxx_messageInfo_CreateUserMessage.DiscardUnknown(m)
}

var xxx_messageInfo_CreateUserMessage proto.InternalMessageInfo

func (m *CreateUserMessage) GetUserCreationToken() string {
	if m != nil {
		return m.UserCreationToken
	}
	return ""
}

func (m *CreateUserMessage) GetPassword() string {
	if m != nil {
		return m.Password
	}
	return ""
}

type SetPasswordMessage struct {
	Token                string   `protobuf:"bytes,1,opt,name=token,proto3" json:"token,omitempty"`
	Id                   string   `protobuf:"bytes,2,opt,name=id,proto3" json:"id,omitempty"`
	Password             string   `protobuf:"bytes,3,opt,name=password,proto3" json:"password,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *SetPasswordMessage) Reset()         { *m = SetPasswordMessage{} }
func (m *SetPasswordMessage) String() string { return proto.CompactTextString(m) }
func (*SetPasswordMessage) ProtoMessage()    {}
func (*SetPasswordMessage) Descriptor() ([]byte, []int) {
	return fileDescriptor_96292a55e4fedcbc, []int{2}
}

func (m *SetPasswordMessage) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_SetPasswordMessage.Unmarshal(m, b)
}
func (m *SetPasswordMessage) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_SetPasswordMessage.Marshal(b, m, deterministic)
}
func (m *SetPasswordMessage) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SetPasswordMessage.Merge(m, src)
}
func (m *SetPasswordMessage) XXX_Size() int {
	return xxx_messageInfo_SetPasswordMessage.Size(m)
}
func (m *SetPasswordMessage) XXX_DiscardUnknown() {
	xxx_messageInfo_SetPasswordMessage.DiscardUnknown(m)
}

var xxx_messageInfo_SetPasswordMessage proto.InternalMessageInfo

func (m *SetPasswordMessage) GetToken() string {
	if m != nil {
		return m.Token
	}
	return ""
}

func (m *SetPasswordMessage) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

func (m *SetPasswordMessage) GetPassword() string {
	if m != nil {
		return m.Password
	}
	return ""
}

type CreateAgentMessage struct {
	AccountId            string   `protobuf:"bytes,1,opt,name=account_id,json=accountId,proto3" json:"account_id,omitempty"`
	AgentName            string   `protobuf:"bytes,2,opt,name=agent_name,json=agentName,proto3" json:"agent_name,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *CreateAgentMessage) Reset()         { *m = CreateAgentMessage{} }
func (m *CreateAgentMessage) String() string { return proto.CompactTextString(m) }
func (*CreateAgentMessage) ProtoMessage()    {}
func (*CreateAgentMessage) Descriptor() ([]byte, []int) {
	return fileDescriptor_96292a55e4fedcbc, []int{3}
}

func (m *CreateAgentMessage) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CreateAgentMessage.Unmarshal(m, b)
}
func (m *CreateAgentMessage) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CreateAgentMessage.Marshal(b, m, deterministic)
}
func (m *CreateAgentMessage) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CreateAgentMessage.Merge(m, src)
}
func (m *CreateAgentMessage) XXX_Size() int {
	return xxx_messageInfo_CreateAgentMessage.Size(m)
}
func (m *CreateAgentMessage) XXX_DiscardUnknown() {
	xxx_messageInfo_CreateAgentMessage.DiscardUnknown(m)
}

var xxx_messageInfo_CreateAgentMessage proto.InternalMessageInfo

func (m *CreateAgentMessage) GetAccountId() string {
	if m != nil {
		return m.AccountId
	}
	return ""
}

func (m *CreateAgentMessage) GetAgentName() string {
	if m != nil {
		return m.AgentName
	}
	return ""
}

type RegisterUserRequest struct {
	Email                string   `protobuf:"bytes,1,opt,name=email,proto3" json:"email,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *RegisterUserRequest) Reset()         { *m = RegisterUserRequest{} }
func (m *RegisterUserRequest) String() string { return proto.CompactTextString(m) }
func (*RegisterUserRequest) ProtoMessage()    {}
func (*RegisterUserRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_96292a55e4fedcbc, []int{4}
}

func (m *RegisterUserRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_RegisterUserRequest.Unmarshal(m, b)
}
func (m *RegisterUserRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_RegisterUserRequest.Marshal(b, m, deterministic)
}
func (m *RegisterUserRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_RegisterUserRequest.Merge(m, src)
}
func (m *RegisterUserRequest) XXX_Size() int {
	return xxx_messageInfo_RegisterUserRequest.Size(m)
}
func (m *RegisterUserRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_RegisterUserRequest.DiscardUnknown(m)
}

var xxx_messageInfo_RegisterUserRequest proto.InternalMessageInfo

func (m *RegisterUserRequest) GetEmail() string {
	if m != nil {
		return m.Email
	}
	return ""
}

type RegisterUserResponse struct {
	VerificationToken    string   `protobuf:"bytes,1,opt,name=verification_token,json=verificationToken,proto3" json:"verification_token,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *RegisterUserResponse) Reset()         { *m = RegisterUserResponse{} }
func (m *RegisterUserResponse) String() string { return proto.CompactTextString(m) }
func (*RegisterUserResponse) ProtoMessage()    {}
func (*RegisterUserResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_96292a55e4fedcbc, []int{5}
}

func (m *RegisterUserResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_RegisterUserResponse.Unmarshal(m, b)
}
func (m *RegisterUserResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_RegisterUserResponse.Marshal(b, m, deterministic)
}
func (m *RegisterUserResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_RegisterUserResponse.Merge(m, src)
}
func (m *RegisterUserResponse) XXX_Size() int {
	return xxx_messageInfo_RegisterUserResponse.Size(m)
}
func (m *RegisterUserResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_RegisterUserResponse.DiscardUnknown(m)
}

var xxx_messageInfo_RegisterUserResponse proto.InternalMessageInfo

func (m *RegisterUserResponse) GetVerificationToken() string {
	if m != nil {
		return m.VerificationToken
	}
	return ""
}

type ConfirmCodeRequest struct {
	VerificationToken    string   `protobuf:"bytes,1,opt,name=verification_token,json=verificationToken,proto3" json:"verification_token,omitempty"`
	VerificationCode     string   `protobuf:"bytes,2,opt,name=verification_code,json=verificationCode,proto3" json:"verification_code,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ConfirmCodeRequest) Reset()         { *m = ConfirmCodeRequest{} }
func (m *ConfirmCodeRequest) String() string { return proto.CompactTextString(m) }
func (*ConfirmCodeRequest) ProtoMessage()    {}
func (*ConfirmCodeRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_96292a55e4fedcbc, []int{6}
}

func (m *ConfirmCodeRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ConfirmCodeRequest.Unmarshal(m, b)
}
func (m *ConfirmCodeRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ConfirmCodeRequest.Marshal(b, m, deterministic)
}
func (m *ConfirmCodeRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ConfirmCodeRequest.Merge(m, src)
}
func (m *ConfirmCodeRequest) XXX_Size() int {
	return xxx_messageInfo_ConfirmCodeRequest.Size(m)
}
func (m *ConfirmCodeRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_ConfirmCodeRequest.DiscardUnknown(m)
}

var xxx_messageInfo_ConfirmCodeRequest proto.InternalMessageInfo

func (m *ConfirmCodeRequest) GetVerificationToken() string {
	if m != nil {
		return m.VerificationToken
	}
	return ""
}

func (m *ConfirmCodeRequest) GetVerificationCode() string {
	if m != nil {
		return m.VerificationCode
	}
	return ""
}

type ConfirmCodeResponse struct {
	Success              bool     `protobuf:"varint,1,opt,name=success,proto3" json:"success,omitempty"`
	UserCreationToken    string   `protobuf:"bytes,2,opt,name=user_creation_token,json=userCreationToken,proto3" json:"user_creation_token,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ConfirmCodeResponse) Reset()         { *m = ConfirmCodeResponse{} }
func (m *ConfirmCodeResponse) String() string { return proto.CompactTextString(m) }
func (*ConfirmCodeResponse) ProtoMessage()    {}
func (*ConfirmCodeResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_96292a55e4fedcbc, []int{7}
}

func (m *ConfirmCodeResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ConfirmCodeResponse.Unmarshal(m, b)
}
func (m *ConfirmCodeResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ConfirmCodeResponse.Marshal(b, m, deterministic)
}
func (m *ConfirmCodeResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ConfirmCodeResponse.Merge(m, src)
}
func (m *ConfirmCodeResponse) XXX_Size() int {
	return xxx_messageInfo_ConfirmCodeResponse.Size(m)
}
func (m *ConfirmCodeResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_ConfirmCodeResponse.DiscardUnknown(m)
}

var xxx_messageInfo_ConfirmCodeResponse proto.InternalMessageInfo

func (m *ConfirmCodeResponse) GetSuccess() bool {
	if m != nil {
		return m.Success
	}
	return false
}

func (m *ConfirmCodeResponse) GetUserCreationToken() string {
	if m != nil {
		return m.UserCreationToken
	}
	return ""
}

type LoginRequest struct {
	Email                string   `protobuf:"bytes,1,opt,name=email,proto3" json:"email,omitempty"`
	Password             string   `protobuf:"bytes,2,opt,name=password,proto3" json:"password,omitempty"`
	Token                string   `protobuf:"bytes,3,opt,name=token,proto3" json:"token,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *LoginRequest) Reset()         { *m = LoginRequest{} }
func (m *LoginRequest) String() string { return proto.CompactTextString(m) }
func (*LoginRequest) ProtoMessage()    {}
func (*LoginRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_96292a55e4fedcbc, []int{8}
}

func (m *LoginRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_LoginRequest.Unmarshal(m, b)
}
func (m *LoginRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_LoginRequest.Marshal(b, m, deterministic)
}
func (m *LoginRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_LoginRequest.Merge(m, src)
}
func (m *LoginRequest) XXX_Size() int {
	return xxx_messageInfo_LoginRequest.Size(m)
}
func (m *LoginRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_LoginRequest.DiscardUnknown(m)
}

var xxx_messageInfo_LoginRequest proto.InternalMessageInfo

func (m *LoginRequest) GetEmail() string {
	if m != nil {
		return m.Email
	}
	return ""
}

func (m *LoginRequest) GetPassword() string {
	if m != nil {
		return m.Password
	}
	return ""
}

func (m *LoginRequest) GetToken() string {
	if m != nil {
		return m.Token
	}
	return ""
}

type LoginResponse struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *LoginResponse) Reset()         { *m = LoginResponse{} }
func (m *LoginResponse) String() string { return proto.CompactTextString(m) }
func (*LoginResponse) ProtoMessage()    {}
func (*LoginResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_96292a55e4fedcbc, []int{9}
}

func (m *LoginResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_LoginResponse.Unmarshal(m, b)
}
func (m *LoginResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_LoginResponse.Marshal(b, m, deterministic)
}
func (m *LoginResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_LoginResponse.Merge(m, src)
}
func (m *LoginResponse) XXX_Size() int {
	return xxx_messageInfo_LoginResponse.Size(m)
}
func (m *LoginResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_LoginResponse.DiscardUnknown(m)
}

var xxx_messageInfo_LoginResponse proto.InternalMessageInfo

func init() {
	proto.RegisterType((*CreateAccountMessage)(nil), "adminservice.CreateAccountMessage")
	proto.RegisterType((*CreateUserMessage)(nil), "adminservice.CreateUserMessage")
	proto.RegisterType((*SetPasswordMessage)(nil), "adminservice.SetPasswordMessage")
	proto.RegisterType((*CreateAgentMessage)(nil), "adminservice.CreateAgentMessage")
	proto.RegisterType((*RegisterUserRequest)(nil), "adminservice.RegisterUserRequest")
	proto.RegisterType((*RegisterUserResponse)(nil), "adminservice.RegisterUserResponse")
	proto.RegisterType((*ConfirmCodeRequest)(nil), "adminservice.ConfirmCodeRequest")
	proto.RegisterType((*ConfirmCodeResponse)(nil), "adminservice.ConfirmCodeResponse")
	proto.RegisterType((*LoginRequest)(nil), "adminservice.LoginRequest")
	proto.RegisterType((*LoginResponse)(nil), "adminservice.LoginResponse")
}

func init() { proto.RegisterFile("proto/adminservice.proto", fileDescriptor_96292a55e4fedcbc) }

var fileDescriptor_96292a55e4fedcbc = []byte{
	// 775 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x94, 0x55, 0x5d, 0x4f, 0xe3, 0x46,
	0x14, 0x55, 0x4c, 0x29, 0x70, 0x13, 0x3e, 0x32, 0x49, 0x21, 0xb8, 0xd0, 0xc2, 0xa8, 0xaa, 0x10,
	0xa8, 0xb1, 0xa0, 0xaa, 0x2a, 0xf1, 0x54, 0x1a, 0x21, 0x54, 0x89, 0x56, 0x55, 0x68, 0x79, 0x68,
	0x2b, 0x59, 0xc6, 0xbe, 0x24, 0xa3, 0xda, 0x1e, 0xd7, 0x33, 0x86, 0x46, 0x88, 0x97, 0x4a, 0x7d,
	0xd9, 0xd7, 0xfd, 0x11, 0xfb, 0x83, 0xf6, 0x2f, 0xec, 0x0f, 0x59, 0xcd, 0x78, 0x4c, 0x6c, 0x62,
	0xb2, 0xda, 0x97, 0x28, 0x33, 0xf7, 0xfa, 0x9c, 0xe3, 0x33, 0xf7, 0x8c, 0xa1, 0x97, 0xa4, 0x5c,
	0x72, 0xc7, 0x0b, 0x22, 0x16, 0x0b, 0x4c, 0xef, 0x98, 0x8f, 0x7d, 0xbd, 0x45, 0x5a, 0xe5, 0x3d,
	0x7b, 0x67, 0xc4, 0xf9, 0x28, 0x44, 0xc7, 0x4b, 0x98, 0xe3, 0xc5, 0x31, 0x97, 0x9e, 0x64, 0x3c,
	0x16, 0x79, 0xaf, 0xbd, 0xe1, 0xf3, 0x28, 0xe2, 0xb1, 0x9c, 0x24, 0xe6, 0x69, 0xbb, 0x1d, 0xf1,
	0x00, 0x43, 0x47, 0xff, 0xe6, 0x5b, 0x34, 0x84, 0xee, 0x20, 0x45, 0x4f, 0xe2, 0x99, 0xef, 0xf3,
	0x2c, 0x96, 0x3f, 0xa3, 0x10, 0xde, 0x08, 0x09, 0x81, 0x4f, 0x62, 0x2f, 0xc2, 0x5e, 0x63, 0xaf,
	0x71, 0xb0, 0x32, 0xd4, 0xff, 0xc9, 0x0f, 0xd0, 0xf4, 0x75, 0xaf, 0x9b, 0x09, 0x4c, 0x7b, 0xd6,
	0x5e, 0xe3, 0xa0, 0x79, 0xf2, 0x65, 0xbf, 0x22, 0x33, 0x07, 0xfb, 0x5d, 0x60, 0x6a, 0x90, 0x86,
	0xe0, 0x3f, 0x6d, 0x51, 0x17, 0xda, 0x33, 0x0d, 0xa4, 0x0f, 0x1d, 0x85, 0xe7, 0xea, 0x3e, 0xc6,
	0x63, 0x57, 0xf2, 0xbf, 0x31, 0x36, 0xcc, 0x6d, 0x55, 0x1a, 0x98, 0xca, 0x6f, 0xaa, 0x40, 0x6c,
	0x58, 0x4e, 0x3c, 0x21, 0xee, 0x79, 0x1a, 0x68, 0x0d, 0x2b, 0xc3, 0xa7, 0x35, 0xbd, 0x06, 0x72,
	0x85, 0xf2, 0x57, 0xb3, 0x2c, 0x18, 0xba, 0xb0, 0x58, 0xc6, 0xcc, 0x17, 0x64, 0x0d, 0x2c, 0x56,
	0x20, 0x58, 0x2c, 0xa8, 0xe0, 0x2e, 0x3c, 0xc3, 0x1d, 0x02, 0x31, 0x36, 0x8d, 0x70, 0x6a, 0xd2,
	0x2e, 0x80, 0x97, 0xdb, 0xe6, 0xb2, 0xc0, 0x80, 0xaf, 0x98, 0x9d, 0x9f, 0x02, 0x5d, 0x56, 0xed,
	0xae, 0x76, 0xd2, 0x32, 0x65, 0xb5, 0xf3, 0x8b, 0x17, 0x21, 0x3d, 0x82, 0xce, 0x10, 0x47, 0x4c,
	0x48, 0x4c, 0x95, 0x1d, 0x43, 0xfc, 0x27, 0x43, 0x21, 0x95, 0x58, 0x8c, 0x3c, 0x16, 0x16, 0x62,
	0xf5, 0x82, 0x9e, 0x43, 0xb7, 0xda, 0x2c, 0x12, 0x1e, 0x0b, 0x24, 0xdf, 0x00, 0xb9, 0xc3, 0x94,
	0xdd, 0x32, 0xbf, 0xc6, 0xbb, 0x72, 0x45, 0x7b, 0x47, 0x13, 0x20, 0x03, 0x1e, 0xdf, 0xb2, 0x34,
	0x1a, 0xf0, 0x00, 0x0b, 0xca, 0x8f, 0x03, 0x21, 0x47, 0x50, 0xd9, 0x74, 0x7d, 0x1e, 0x14, 0xaf,
	0xb7, 0x51, 0x2e, 0x28, 0x0a, 0xea, 0x42, 0xa7, 0xc2, 0x68, 0x74, 0xf7, 0x60, 0x49, 0x64, 0xbe,
	0x8f, 0x42, 0x68, 0x9e, 0xe5, 0x61, 0xb1, 0x7c, 0x69, 0x1c, 0xac, 0x17, 0xc6, 0x81, 0x5e, 0x43,
	0xeb, 0x92, 0x8f, 0x58, 0x3c, 0xd7, 0xbf, 0x79, 0x43, 0x33, 0x1d, 0x8f, 0x85, 0xd2, 0x78, 0xd0,
	0x75, 0x58, 0x35, 0xb8, 0xb9, 0xe4, 0x93, 0x7b, 0x58, 0x33, 0x21, 0xb9, 0xca, 0x87, 0x9d, 0x20,
	0xac, 0x56, 0xc2, 0x43, 0x68, 0x5d, 0x18, 0xaa, 0xc9, 0xb2, 0xdb, 0xfd, 0x52, 0x2e, 0xcf, 0xa3,
	0x44, 0x4e, 0xe8, 0xee, 0x7f, 0x6f, 0xdf, 0xbd, 0xb6, 0xb6, 0x28, 0xc9, 0x23, 0xef, 0xdc, 0x1d,
	0x3b, 0x66, 0x8a, 0xc4, 0x69, 0xe3, 0xf0, 0xe4, 0x8d, 0x05, 0x4d, 0x75, 0xe8, 0x05, 0xed, 0x5f,
	0x00, 0xd3, 0x14, 0x91, 0x0f, 0x05, 0xb0, 0x8e, 0xd0, 0xd6, 0x84, 0x5d, 0xba, 0x3e, 0x25, 0x54,
	0xc6, 0x2a, 0x36, 0x72, 0x09, 0x4b, 0x17, 0x28, 0x35, 0xf4, 0x66, 0xf9, 0xc9, 0x0b, 0x7c, 0x7a,
	0x85, 0x66, 0x3f, 0xbf, 0x42, 0x74, 0xa6, 0x77, 0x34, 0xd6, 0x26, 0xe9, 0x3e, 0xc3, 0x72, 0x1e,
	0x58, 0xf0, 0x48, 0x42, 0x68, 0x96, 0x02, 0x49, 0xf6, 0xaa, 0x62, 0x67, 0xb3, 0x5a, 0xa7, 0xf6,
	0x40, 0x33, 0x50, 0xba, 0x5b, 0xc7, 0xe0, 0x14, 0xc7, 0xa8, 0x9c, 0xfa, 0xbf, 0x01, 0x2d, 0x9d,
	0xd0, 0xc2, 0xaa, 0x0c, 0x9a, 0xa5, 0xdc, 0x3e, 0xa7, 0x9f, 0x8d, 0x74, 0x1d, 0xfd, 0xb1, 0xa6,
	0x3f, 0xa2, 0x5f, 0xcf, 0x9e, 0x8e, 0xf3, 0x30, 0xcd, 0xff, 0xa3, 0xa3, 0xa3, 0xad, 0x4f, 0xec,
	0x95, 0x55, 0x64, 0x3b, 0xd5, 0x93, 0x5a, 0xc8, 0x99, 0x40, 0xab, 0x9c, 0x62, 0xb2, 0x5f, 0xd5,
	0x53, 0x73, 0x1d, 0xd8, 0x74, 0x5e, 0x4b, 0x3e, 0x99, 0x94, 0x6a, 0x85, 0x3b, 0x74, 0x6b, 0xaa,
	0x30, 0x35, 0x7d, 0xda, 0x29, 0x75, 0xac, 0xff, 0x42, 0xb3, 0x94, 0xc3, 0x19, 0x27, 0x66, 0x2e,
	0x05, 0x7b, 0x7f, 0x4e, 0x87, 0xe1, 0xfd, 0x4a, 0xf3, 0x7e, 0x41, 0xb7, 0x6b, 0x78, 0xf5, 0x45,
	0x30, 0x51, 0x66, 0x48, 0xf8, 0xec, 0x2c, 0x93, 0x63, 0x8c, 0xa5, 0xb9, 0x17, 0x0a, 0x37, 0xfe,
	0x84, 0x45, 0x9d, 0x30, 0x62, 0x57, 0xa9, 0xca, 0x71, 0xb6, 0x3f, 0xaf, 0xad, 0x19, 0x01, 0xdb,
	0x5a, 0x40, 0x87, 0xae, 0x39, 0x5e, 0x26, 0xc7, 0x8a, 0x3f, 0x54, 0xf5, 0xd3, 0xc6, 0xe1, 0x8f,
	0xdf, 0xff, 0xf1, 0x9d, 0x1c, 0xa3, 0xc4, 0x34, 0x65, 0x37, 0x21, 0xfa, 0x63, 0x16, 0x06, 0xce,
	0x20, 0xe4, 0x59, 0x70, 0x96, 0x24, 0x8e, 0x97, 0x24, 0xa1, 0x11, 0x23, 0x72, 0xed, 0x5e, 0x92,
	0xe4, 0xdf, 0xc5, 0x9b, 0x4f, 0xf5, 0x87, 0xf1, 0xdb, 0xf7, 0x01, 0x00, 0x00, 0xff, 0xff, 0xfd,
	0x93, 0xe7, 0x69, 0x85, 0x07, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// AccountServiceClient is the client API for AccountService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type AccountServiceClient interface {
	CreateAccount(ctx context.Context, in *CreateAccountMessage, opts ...grpc.CallOption) (*common.Empty, error)
}

type accountServiceClient struct {
	cc *grpc.ClientConn
}

func NewAccountServiceClient(cc *grpc.ClientConn) AccountServiceClient {
	return &accountServiceClient{cc}
}

func (c *accountServiceClient) CreateAccount(ctx context.Context, in *CreateAccountMessage, opts ...grpc.CallOption) (*common.Empty, error) {
	out := new(common.Empty)
	err := c.cc.Invoke(ctx, "/adminservice.AccountService/CreateAccount", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// AccountServiceServer is the server API for AccountService service.
type AccountServiceServer interface {
	CreateAccount(context.Context, *CreateAccountMessage) (*common.Empty, error)
}

func RegisterAccountServiceServer(s *grpc.Server, srv AccountServiceServer) {
	s.RegisterService(&_AccountService_serviceDesc, srv)
}

func _AccountService_CreateAccount_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateAccountMessage)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AccountServiceServer).CreateAccount(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/adminservice.AccountService/CreateAccount",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AccountServiceServer).CreateAccount(ctx, req.(*CreateAccountMessage))
	}
	return interceptor(ctx, in, info, handler)
}

var _AccountService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "adminservice.AccountService",
	HandlerType: (*AccountServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "CreateAccount",
			Handler:    _AccountService_CreateAccount_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "proto/adminservice.proto",
}

// UserServiceClient is the client API for UserService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type UserServiceClient interface {
	CreateUser(ctx context.Context, in *CreateUserMessage, opts ...grpc.CallOption) (*common.Empty, error)
	GetUser(ctx context.Context, in *common.GetMessage, opts ...grpc.CallOption) (*model.User, error)
	SetPassword(ctx context.Context, in *SetPasswordMessage, opts ...grpc.CallOption) (*common.Empty, error)
}

type userServiceClient struct {
	cc *grpc.ClientConn
}

func NewUserServiceClient(cc *grpc.ClientConn) UserServiceClient {
	return &userServiceClient{cc}
}

func (c *userServiceClient) CreateUser(ctx context.Context, in *CreateUserMessage, opts ...grpc.CallOption) (*common.Empty, error) {
	out := new(common.Empty)
	err := c.cc.Invoke(ctx, "/adminservice.UserService/CreateUser", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userServiceClient) GetUser(ctx context.Context, in *common.GetMessage, opts ...grpc.CallOption) (*model.User, error) {
	out := new(model.User)
	err := c.cc.Invoke(ctx, "/adminservice.UserService/GetUser", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userServiceClient) SetPassword(ctx context.Context, in *SetPasswordMessage, opts ...grpc.CallOption) (*common.Empty, error) {
	out := new(common.Empty)
	err := c.cc.Invoke(ctx, "/adminservice.UserService/SetPassword", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// UserServiceServer is the server API for UserService service.
type UserServiceServer interface {
	CreateUser(context.Context, *CreateUserMessage) (*common.Empty, error)
	GetUser(context.Context, *common.GetMessage) (*model.User, error)
	SetPassword(context.Context, *SetPasswordMessage) (*common.Empty, error)
}

func RegisterUserServiceServer(s *grpc.Server, srv UserServiceServer) {
	s.RegisterService(&_UserService_serviceDesc, srv)
}

func _UserService_CreateUser_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateUserMessage)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserServiceServer).CreateUser(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/adminservice.UserService/CreateUser",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserServiceServer).CreateUser(ctx, req.(*CreateUserMessage))
	}
	return interceptor(ctx, in, info, handler)
}

func _UserService_GetUser_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(common.GetMessage)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserServiceServer).GetUser(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/adminservice.UserService/GetUser",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserServiceServer).GetUser(ctx, req.(*common.GetMessage))
	}
	return interceptor(ctx, in, info, handler)
}

func _UserService_SetPassword_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SetPasswordMessage)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserServiceServer).SetPassword(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/adminservice.UserService/SetPassword",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserServiceServer).SetPassword(ctx, req.(*SetPasswordMessage))
	}
	return interceptor(ctx, in, info, handler)
}

var _UserService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "adminservice.UserService",
	HandlerType: (*UserServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "CreateUser",
			Handler:    _UserService_CreateUser_Handler,
		},
		{
			MethodName: "GetUser",
			Handler:    _UserService_GetUser_Handler,
		},
		{
			MethodName: "SetPassword",
			Handler:    _UserService_SetPassword_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "proto/adminservice.proto",
}

// AgentServiceClient is the client API for AgentService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type AgentServiceClient interface {
	CreateAgent(ctx context.Context, in *CreateAgentMessage, opts ...grpc.CallOption) (*common.Empty, error)
}

type agentServiceClient struct {
	cc *grpc.ClientConn
}

func NewAgentServiceClient(cc *grpc.ClientConn) AgentServiceClient {
	return &agentServiceClient{cc}
}

func (c *agentServiceClient) CreateAgent(ctx context.Context, in *CreateAgentMessage, opts ...grpc.CallOption) (*common.Empty, error) {
	out := new(common.Empty)
	err := c.cc.Invoke(ctx, "/adminservice.AgentService/CreateAgent", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// AgentServiceServer is the server API for AgentService service.
type AgentServiceServer interface {
	CreateAgent(context.Context, *CreateAgentMessage) (*common.Empty, error)
}

func RegisterAgentServiceServer(s *grpc.Server, srv AgentServiceServer) {
	s.RegisterService(&_AgentService_serviceDesc, srv)
}

func _AgentService_CreateAgent_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateAgentMessage)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AgentServiceServer).CreateAgent(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/adminservice.AgentService/CreateAgent",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AgentServiceServer).CreateAgent(ctx, req.(*CreateAgentMessage))
	}
	return interceptor(ctx, in, info, handler)
}

var _AgentService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "adminservice.AgentService",
	HandlerType: (*AgentServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "CreateAgent",
			Handler:    _AgentService_CreateAgent_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "proto/adminservice.proto",
}

// RegistrationServiceClient is the client API for RegistrationService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type RegistrationServiceClient interface {
	RegisterUser(ctx context.Context, in *RegisterUserRequest, opts ...grpc.CallOption) (*RegisterUserResponse, error)
	ConfirmCode(ctx context.Context, in *ConfirmCodeRequest, opts ...grpc.CallOption) (*ConfirmCodeResponse, error)
}

type registrationServiceClient struct {
	cc *grpc.ClientConn
}

func NewRegistrationServiceClient(cc *grpc.ClientConn) RegistrationServiceClient {
	return &registrationServiceClient{cc}
}

func (c *registrationServiceClient) RegisterUser(ctx context.Context, in *RegisterUserRequest, opts ...grpc.CallOption) (*RegisterUserResponse, error) {
	out := new(RegisterUserResponse)
	err := c.cc.Invoke(ctx, "/adminservice.RegistrationService/RegisterUser", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *registrationServiceClient) ConfirmCode(ctx context.Context, in *ConfirmCodeRequest, opts ...grpc.CallOption) (*ConfirmCodeResponse, error) {
	out := new(ConfirmCodeResponse)
	err := c.cc.Invoke(ctx, "/adminservice.RegistrationService/ConfirmCode", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// RegistrationServiceServer is the server API for RegistrationService service.
type RegistrationServiceServer interface {
	RegisterUser(context.Context, *RegisterUserRequest) (*RegisterUserResponse, error)
	ConfirmCode(context.Context, *ConfirmCodeRequest) (*ConfirmCodeResponse, error)
}

func RegisterRegistrationServiceServer(s *grpc.Server, srv RegistrationServiceServer) {
	s.RegisterService(&_RegistrationService_serviceDesc, srv)
}

func _RegistrationService_RegisterUser_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RegisterUserRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RegistrationServiceServer).RegisterUser(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/adminservice.RegistrationService/RegisterUser",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RegistrationServiceServer).RegisterUser(ctx, req.(*RegisterUserRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _RegistrationService_ConfirmCode_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ConfirmCodeRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RegistrationServiceServer).ConfirmCode(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/adminservice.RegistrationService/ConfirmCode",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RegistrationServiceServer).ConfirmCode(ctx, req.(*ConfirmCodeRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _RegistrationService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "adminservice.RegistrationService",
	HandlerType: (*RegistrationServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "RegisterUser",
			Handler:    _RegistrationService_RegisterUser_Handler,
		},
		{
			MethodName: "ConfirmCode",
			Handler:    _RegistrationService_ConfirmCode_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "proto/adminservice.proto",
}

// AuthenticationServiceClient is the client API for AuthenticationService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type AuthenticationServiceClient interface {
	Login(ctx context.Context, in *LoginRequest, opts ...grpc.CallOption) (*LoginResponse, error)
}

type authenticationServiceClient struct {
	cc *grpc.ClientConn
}

func NewAuthenticationServiceClient(cc *grpc.ClientConn) AuthenticationServiceClient {
	return &authenticationServiceClient{cc}
}

func (c *authenticationServiceClient) Login(ctx context.Context, in *LoginRequest, opts ...grpc.CallOption) (*LoginResponse, error) {
	out := new(LoginResponse)
	err := c.cc.Invoke(ctx, "/adminservice.AuthenticationService/Login", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// AuthenticationServiceServer is the server API for AuthenticationService service.
type AuthenticationServiceServer interface {
	Login(context.Context, *LoginRequest) (*LoginResponse, error)
}

func RegisterAuthenticationServiceServer(s *grpc.Server, srv AuthenticationServiceServer) {
	s.RegisterService(&_AuthenticationService_serviceDesc, srv)
}

func _AuthenticationService_Login_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(LoginRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AuthenticationServiceServer).Login(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/adminservice.AuthenticationService/Login",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AuthenticationServiceServer).Login(ctx, req.(*LoginRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _AuthenticationService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "adminservice.AuthenticationService",
	HandlerType: (*AuthenticationServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Login",
			Handler:    _AuthenticationService_Login_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "proto/adminservice.proto",
}
