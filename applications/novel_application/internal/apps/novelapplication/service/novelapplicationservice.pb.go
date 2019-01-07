// Code generated by protoc-gen-go. DO NOT EDIT.
// source: novelapplicationservice.proto

package service // import "theterriblechild/CloudApp/cloud_appplication_portal/cloud_applications/novel_application/internal/app/novelapplication/service"

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"
import model "theterriblechild/CloudApp/cloud_appplication_portal/cloud_applications/novel_application/internal/model"

import (
	context "golang.org/x/net/context"
	grpc "google.golang.org/grpc"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

type AddNovelRequest struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *AddNovelRequest) Reset()         { *m = AddNovelRequest{} }
func (m *AddNovelRequest) String() string { return proto.CompactTextString(m) }
func (*AddNovelRequest) ProtoMessage()    {}
func (*AddNovelRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_novelapplicationservice_356a0704b7986de5, []int{0}
}
func (m *AddNovelRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_AddNovelRequest.Unmarshal(m, b)
}
func (m *AddNovelRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_AddNovelRequest.Marshal(b, m, deterministic)
}
func (dst *AddNovelRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_AddNovelRequest.Merge(dst, src)
}
func (m *AddNovelRequest) XXX_Size() int {
	return xxx_messageInfo_AddNovelRequest.Size(m)
}
func (m *AddNovelRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_AddNovelRequest.DiscardUnknown(m)
}

var xxx_messageInfo_AddNovelRequest proto.InternalMessageInfo

type GetNovelRequest struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GetNovelRequest) Reset()         { *m = GetNovelRequest{} }
func (m *GetNovelRequest) String() string { return proto.CompactTextString(m) }
func (*GetNovelRequest) ProtoMessage()    {}
func (*GetNovelRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_novelapplicationservice_356a0704b7986de5, []int{1}
}
func (m *GetNovelRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetNovelRequest.Unmarshal(m, b)
}
func (m *GetNovelRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetNovelRequest.Marshal(b, m, deterministic)
}
func (dst *GetNovelRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetNovelRequest.Merge(dst, src)
}
func (m *GetNovelRequest) XXX_Size() int {
	return xxx_messageInfo_GetNovelRequest.Size(m)
}
func (m *GetNovelRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_GetNovelRequest.DiscardUnknown(m)
}

var xxx_messageInfo_GetNovelRequest proto.InternalMessageInfo

type AddChapterRequest struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *AddChapterRequest) Reset()         { *m = AddChapterRequest{} }
func (m *AddChapterRequest) String() string { return proto.CompactTextString(m) }
func (*AddChapterRequest) ProtoMessage()    {}
func (*AddChapterRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_novelapplicationservice_356a0704b7986de5, []int{2}
}
func (m *AddChapterRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_AddChapterRequest.Unmarshal(m, b)
}
func (m *AddChapterRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_AddChapterRequest.Marshal(b, m, deterministic)
}
func (dst *AddChapterRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_AddChapterRequest.Merge(dst, src)
}
func (m *AddChapterRequest) XXX_Size() int {
	return xxx_messageInfo_AddChapterRequest.Size(m)
}
func (m *AddChapterRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_AddChapterRequest.DiscardUnknown(m)
}

var xxx_messageInfo_AddChapterRequest proto.InternalMessageInfo

type GetChapterRequest struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GetChapterRequest) Reset()         { *m = GetChapterRequest{} }
func (m *GetChapterRequest) String() string { return proto.CompactTextString(m) }
func (*GetChapterRequest) ProtoMessage()    {}
func (*GetChapterRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_novelapplicationservice_356a0704b7986de5, []int{3}
}
func (m *GetChapterRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetChapterRequest.Unmarshal(m, b)
}
func (m *GetChapterRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetChapterRequest.Marshal(b, m, deterministic)
}
func (dst *GetChapterRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetChapterRequest.Merge(dst, src)
}
func (m *GetChapterRequest) XXX_Size() int {
	return xxx_messageInfo_GetChapterRequest.Size(m)
}
func (m *GetChapterRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_GetChapterRequest.DiscardUnknown(m)
}

var xxx_messageInfo_GetChapterRequest proto.InternalMessageInfo

func init() {
	proto.RegisterType((*AddNovelRequest)(nil), "novelapplicationservice.AddNovelRequest")
	proto.RegisterType((*GetNovelRequest)(nil), "novelapplicationservice.GetNovelRequest")
	proto.RegisterType((*AddChapterRequest)(nil), "novelapplicationservice.AddChapterRequest")
	proto.RegisterType((*GetChapterRequest)(nil), "novelapplicationservice.GetChapterRequest")
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// NovelServiceClient is the client API for NovelService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type NovelServiceClient interface {
	AddNovel(ctx context.Context, in *AddNovelRequest, opts ...grpc.CallOption) (*model.Novel, error)
	GetNovel(ctx context.Context, in *GetNovelRequest, opts ...grpc.CallOption) (*model.Novel, error)
}

type novelServiceClient struct {
	cc *grpc.ClientConn
}

func NewNovelServiceClient(cc *grpc.ClientConn) NovelServiceClient {
	return &novelServiceClient{cc}
}

func (c *novelServiceClient) AddNovel(ctx context.Context, in *AddNovelRequest, opts ...grpc.CallOption) (*model.Novel, error) {
	out := new(model.Novel)
	err := c.cc.Invoke(ctx, "/novelapplicationservice.NovelService/AddNovel", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *novelServiceClient) GetNovel(ctx context.Context, in *GetNovelRequest, opts ...grpc.CallOption) (*model.Novel, error) {
	out := new(model.Novel)
	err := c.cc.Invoke(ctx, "/novelapplicationservice.NovelService/GetNovel", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// NovelServiceServer is the server API for NovelService service.
type NovelServiceServer interface {
	AddNovel(context.Context, *AddNovelRequest) (*model.Novel, error)
	GetNovel(context.Context, *GetNovelRequest) (*model.Novel, error)
}

func RegisterNovelServiceServer(s *grpc.Server, srv NovelServiceServer) {
	s.RegisterService(&_NovelService_serviceDesc, srv)
}

func _NovelService_AddNovel_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AddNovelRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(NovelServiceServer).AddNovel(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/novelapplicationservice.NovelService/AddNovel",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(NovelServiceServer).AddNovel(ctx, req.(*AddNovelRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _NovelService_GetNovel_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetNovelRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(NovelServiceServer).GetNovel(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/novelapplicationservice.NovelService/GetNovel",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(NovelServiceServer).GetNovel(ctx, req.(*GetNovelRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _NovelService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "novelapplicationservice.NovelService",
	HandlerType: (*NovelServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "AddNovel",
			Handler:    _NovelService_AddNovel_Handler,
		},
		{
			MethodName: "GetNovel",
			Handler:    _NovelService_GetNovel_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "novelapplicationservice.proto",
}

// ChapterServiceClient is the client API for ChapterService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type ChapterServiceClient interface {
	AddChapter(ctx context.Context, in *AddChapterRequest, opts ...grpc.CallOption) (*model.Chapter, error)
	GetChapter(ctx context.Context, in *GetChapterRequest, opts ...grpc.CallOption) (*model.Chapter, error)
}

type chapterServiceClient struct {
	cc *grpc.ClientConn
}

func NewChapterServiceClient(cc *grpc.ClientConn) ChapterServiceClient {
	return &chapterServiceClient{cc}
}

func (c *chapterServiceClient) AddChapter(ctx context.Context, in *AddChapterRequest, opts ...grpc.CallOption) (*model.Chapter, error) {
	out := new(model.Chapter)
	err := c.cc.Invoke(ctx, "/novelapplicationservice.ChapterService/AddChapter", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *chapterServiceClient) GetChapter(ctx context.Context, in *GetChapterRequest, opts ...grpc.CallOption) (*model.Chapter, error) {
	out := new(model.Chapter)
	err := c.cc.Invoke(ctx, "/novelapplicationservice.ChapterService/GetChapter", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// ChapterServiceServer is the server API for ChapterService service.
type ChapterServiceServer interface {
	AddChapter(context.Context, *AddChapterRequest) (*model.Chapter, error)
	GetChapter(context.Context, *GetChapterRequest) (*model.Chapter, error)
}

func RegisterChapterServiceServer(s *grpc.Server, srv ChapterServiceServer) {
	s.RegisterService(&_ChapterService_serviceDesc, srv)
}

func _ChapterService_AddChapter_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AddChapterRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ChapterServiceServer).AddChapter(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/novelapplicationservice.ChapterService/AddChapter",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ChapterServiceServer).AddChapter(ctx, req.(*AddChapterRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ChapterService_GetChapter_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetChapterRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ChapterServiceServer).GetChapter(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/novelapplicationservice.ChapterService/GetChapter",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ChapterServiceServer).GetChapter(ctx, req.(*GetChapterRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _ChapterService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "novelapplicationservice.ChapterService",
	HandlerType: (*ChapterServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "AddChapter",
			Handler:    _ChapterService_AddChapter_Handler,
		},
		{
			MethodName: "GetChapter",
			Handler:    _ChapterService_GetChapter_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "novelapplicationservice.proto",
}

func init() {
	proto.RegisterFile("novelapplicationservice.proto", fileDescriptor_novelapplicationservice_356a0704b7986de5)
}

var fileDescriptor_novelapplicationservice_356a0704b7986de5 = []byte{
	// 279 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x7c, 0x92, 0xc1, 0x4a, 0xc3, 0x40,
	0x10, 0x86, 0x9b, 0x8b, 0xc8, 0x5a, 0x2a, 0x8d, 0x07, 0x21, 0xe0, 0x25, 0x27, 0xf1, 0x90, 0x85,
	0xfa, 0x04, 0xb5, 0x42, 0xc1, 0x83, 0x07, 0xed, 0xc9, 0x4b, 0xd9, 0x64, 0x07, 0xb3, 0xb0, 0xcd,
	0x8e, 0xdb, 0x49, 0xef, 0x82, 0x4f, 0xe1, 0x13, 0xf8, 0x98, 0x92, 0x64, 0xd3, 0x4d, 0x52, 0x9a,
	0xeb, 0xff, 0x6f, 0xbe, 0xf9, 0x66, 0x08, 0xbb, 0x2b, 0xcc, 0x01, 0xb4, 0x40, 0xd4, 0x2a, 0x13,
	0xa4, 0x4c, 0xb1, 0x07, 0x7b, 0x50, 0x19, 0x24, 0x68, 0x0d, 0x99, 0xf0, 0xf6, 0x4c, 0x1d, 0x5d,
	0xed, 0x8c, 0x04, 0xdd, 0xbc, 0x8a, 0xe7, 0xec, 0x7a, 0x29, 0xe5, 0x6b, 0xf5, 0xf4, 0x0d, 0xbe,
	0x4a, 0xd8, 0x53, 0x15, 0xad, 0x81, 0x7a, 0xd1, 0x0d, 0x9b, 0x2f, 0xa5, 0x5c, 0xe5, 0x02, 0x09,
	0x6c, 0x27, 0x5c, 0x03, 0xf5, 0xc3, 0xc5, 0x6f, 0xc0, 0xa6, 0xf5, 0xa7, 0xef, 0xcd, 0xb4, 0xf0,
	0x99, 0x5d, 0xb6, 0x03, 0xc2, 0xfb, 0xe4, 0x9c, 0xf2, 0xc0, 0x21, 0x9a, 0x26, 0x8d, 0x64, 0x1d,
	0xc6, 0x93, 0x8a, 0xd2, 0x3a, 0x8d, 0x50, 0x06, 0xda, 0x43, 0xca, 0xe2, 0x2f, 0x60, 0x33, 0xe7,
	0xdb, 0xea, 0xbd, 0x30, 0xe6, 0x37, 0x0b, 0x1f, 0xc6, 0x04, 0xfb, 0x9b, 0x46, 0x33, 0x07, 0x77,
	0x71, 0x3c, 0xa9, 0x58, 0xfe, 0x20, 0x23, 0xac, 0x93, 0xab, 0x9d, 0xb2, 0x9e, 0x7e, 0x82, 0x8f,
	0xef, 0xe0, 0x53, 0x51, 0x5e, 0xa6, 0x49, 0x66, 0x76, 0x7c, 0x93, 0xc3, 0x06, 0xac, 0x55, 0xa9,
	0x86, 0x55, 0xae, 0xb4, 0xe4, 0x99, 0x36, 0xa5, 0xdc, 0x0a, 0xf4, 0xf8, 0x2d, 0x1a, 0x4b, 0x42,
	0xfb, 0xe6, 0x38, 0x97, 0xd7, 0x26, 0xdd, 0x88, 0xab, 0x82, 0xc0, 0x16, 0x42, 0x73, 0x81, 0xe8,
	0xfa, 0x6e, 0xed, 0x54, 0xd3, 0x8b, 0xfa, 0x2f, 0x79, 0xfc, 0x0f, 0x00, 0x00, 0xff, 0xff, 0x45,
	0x02, 0xc0, 0x03, 0x6c, 0x02, 0x00, 0x00,
}
