// Code generated by protoc-gen-go. DO NOT EDIT.
// source: model.proto

package model // import "github.com/TheTerribleChild/CloudApp/cloud_appplication_portal/cloud_applications/novel_application/internal/model"

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

type Novel struct {
	NovelMetadata        *NovelMetadata         `protobuf:"bytes,1,opt,name=novel_metadata,json=novelMetadata,proto3" json:"novel_metadata,omitempty"`
	SourceMetadata       []*NovelSourceMetadata `protobuf:"bytes,2,rep,name=source_metadata,json=sourceMetadata,proto3" json:"source_metadata,omitempty"`
	Chapter              []*Chapter             `protobuf:"bytes,3,rep,name=chapter,proto3" json:"chapter,omitempty"`
	XXX_NoUnkeyedLiteral struct{}               `json:"-"`
	XXX_unrecognized     []byte                 `json:"-"`
	XXX_sizecache        int32                  `json:"-"`
}

func (m *Novel) Reset()         { *m = Novel{} }
func (m *Novel) String() string { return proto.CompactTextString(m) }
func (*Novel) ProtoMessage()    {}
func (*Novel) Descriptor() ([]byte, []int) {
	return fileDescriptor_model_406c922677d5b648, []int{0}
}
func (m *Novel) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Novel.Unmarshal(m, b)
}
func (m *Novel) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Novel.Marshal(b, m, deterministic)
}
func (dst *Novel) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Novel.Merge(dst, src)
}
func (m *Novel) XXX_Size() int {
	return xxx_messageInfo_Novel.Size(m)
}
func (m *Novel) XXX_DiscardUnknown() {
	xxx_messageInfo_Novel.DiscardUnknown(m)
}

var xxx_messageInfo_Novel proto.InternalMessageInfo

func (m *Novel) GetNovelMetadata() *NovelMetadata {
	if m != nil {
		return m.NovelMetadata
	}
	return nil
}

func (m *Novel) GetSourceMetadata() []*NovelSourceMetadata {
	if m != nil {
		return m.SourceMetadata
	}
	return nil
}

func (m *Novel) GetChapter() []*Chapter {
	if m != nil {
		return m.Chapter
	}
	return nil
}

type NovelMetadata struct {
	Id                   string   `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Title                string   `protobuf:"bytes,2,opt,name=title,proto3" json:"title,omitempty"`
	Author               string   `protobuf:"bytes,3,opt,name=author,proto3" json:"author,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *NovelMetadata) Reset()         { *m = NovelMetadata{} }
func (m *NovelMetadata) String() string { return proto.CompactTextString(m) }
func (*NovelMetadata) ProtoMessage()    {}
func (*NovelMetadata) Descriptor() ([]byte, []int) {
	return fileDescriptor_model_406c922677d5b648, []int{1}
}
func (m *NovelMetadata) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_NovelMetadata.Unmarshal(m, b)
}
func (m *NovelMetadata) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_NovelMetadata.Marshal(b, m, deterministic)
}
func (dst *NovelMetadata) XXX_Merge(src proto.Message) {
	xxx_messageInfo_NovelMetadata.Merge(dst, src)
}
func (m *NovelMetadata) XXX_Size() int {
	return xxx_messageInfo_NovelMetadata.Size(m)
}
func (m *NovelMetadata) XXX_DiscardUnknown() {
	xxx_messageInfo_NovelMetadata.DiscardUnknown(m)
}

var xxx_messageInfo_NovelMetadata proto.InternalMessageInfo

func (m *NovelMetadata) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

func (m *NovelMetadata) GetTitle() string {
	if m != nil {
		return m.Title
	}
	return ""
}

func (m *NovelMetadata) GetAuthor() string {
	if m != nil {
		return m.Author
	}
	return ""
}

type Chapter struct {
	ChapterMetadata      *ChapterMetadata     `protobuf:"bytes,1,opt,name=chapter_metadata,json=chapterMetadata,proto3" json:"chapter_metadata,omitempty"`
	SourceMetadata       []*ChapterSourceData `protobuf:"bytes,2,rep,name=source_metadata,json=sourceMetadata,proto3" json:"source_metadata,omitempty"`
	XXX_NoUnkeyedLiteral struct{}             `json:"-"`
	XXX_unrecognized     []byte               `json:"-"`
	XXX_sizecache        int32                `json:"-"`
}

func (m *Chapter) Reset()         { *m = Chapter{} }
func (m *Chapter) String() string { return proto.CompactTextString(m) }
func (*Chapter) ProtoMessage()    {}
func (*Chapter) Descriptor() ([]byte, []int) {
	return fileDescriptor_model_406c922677d5b648, []int{2}
}
func (m *Chapter) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Chapter.Unmarshal(m, b)
}
func (m *Chapter) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Chapter.Marshal(b, m, deterministic)
}
func (dst *Chapter) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Chapter.Merge(dst, src)
}
func (m *Chapter) XXX_Size() int {
	return xxx_messageInfo_Chapter.Size(m)
}
func (m *Chapter) XXX_DiscardUnknown() {
	xxx_messageInfo_Chapter.DiscardUnknown(m)
}

var xxx_messageInfo_Chapter proto.InternalMessageInfo

func (m *Chapter) GetChapterMetadata() *ChapterMetadata {
	if m != nil {
		return m.ChapterMetadata
	}
	return nil
}

func (m *Chapter) GetSourceMetadata() []*ChapterSourceData {
	if m != nil {
		return m.SourceMetadata
	}
	return nil
}

type ChapterMetadata struct {
	Id                   string   `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Title                string   `protobuf:"bytes,2,opt,name=title,proto3" json:"title,omitempty"`
	Index                int32    `protobuf:"varint,3,opt,name=index,proto3" json:"index,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ChapterMetadata) Reset()         { *m = ChapterMetadata{} }
func (m *ChapterMetadata) String() string { return proto.CompactTextString(m) }
func (*ChapterMetadata) ProtoMessage()    {}
func (*ChapterMetadata) Descriptor() ([]byte, []int) {
	return fileDescriptor_model_406c922677d5b648, []int{3}
}
func (m *ChapterMetadata) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ChapterMetadata.Unmarshal(m, b)
}
func (m *ChapterMetadata) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ChapterMetadata.Marshal(b, m, deterministic)
}
func (dst *ChapterMetadata) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ChapterMetadata.Merge(dst, src)
}
func (m *ChapterMetadata) XXX_Size() int {
	return xxx_messageInfo_ChapterMetadata.Size(m)
}
func (m *ChapterMetadata) XXX_DiscardUnknown() {
	xxx_messageInfo_ChapterMetadata.DiscardUnknown(m)
}

var xxx_messageInfo_ChapterMetadata proto.InternalMessageInfo

func (m *ChapterMetadata) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

func (m *ChapterMetadata) GetTitle() string {
	if m != nil {
		return m.Title
	}
	return ""
}

func (m *ChapterMetadata) GetIndex() int32 {
	if m != nil {
		return m.Index
	}
	return 0
}

type NovelSourceData struct {
	Metadata             *NovelSourceMetadata     `protobuf:"bytes,1,opt,name=metadata,proto3" json:"metadata,omitempty"`
	Chapters             []*ChapterSourceMetadata `protobuf:"bytes,2,rep,name=chapters,proto3" json:"chapters,omitempty"`
	XXX_NoUnkeyedLiteral struct{}                 `json:"-"`
	XXX_unrecognized     []byte                   `json:"-"`
	XXX_sizecache        int32                    `json:"-"`
}

func (m *NovelSourceData) Reset()         { *m = NovelSourceData{} }
func (m *NovelSourceData) String() string { return proto.CompactTextString(m) }
func (*NovelSourceData) ProtoMessage()    {}
func (*NovelSourceData) Descriptor() ([]byte, []int) {
	return fileDescriptor_model_406c922677d5b648, []int{4}
}
func (m *NovelSourceData) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_NovelSourceData.Unmarshal(m, b)
}
func (m *NovelSourceData) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_NovelSourceData.Marshal(b, m, deterministic)
}
func (dst *NovelSourceData) XXX_Merge(src proto.Message) {
	xxx_messageInfo_NovelSourceData.Merge(dst, src)
}
func (m *NovelSourceData) XXX_Size() int {
	return xxx_messageInfo_NovelSourceData.Size(m)
}
func (m *NovelSourceData) XXX_DiscardUnknown() {
	xxx_messageInfo_NovelSourceData.DiscardUnknown(m)
}

var xxx_messageInfo_NovelSourceData proto.InternalMessageInfo

func (m *NovelSourceData) GetMetadata() *NovelSourceMetadata {
	if m != nil {
		return m.Metadata
	}
	return nil
}

func (m *NovelSourceData) GetChapters() []*ChapterSourceMetadata {
	if m != nil {
		return m.Chapters
	}
	return nil
}

type NovelSourceMetadata struct {
	Id                   string   `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	SourceId             string   `protobuf:"bytes,2,opt,name=source_id,json=sourceId,proto3" json:"source_id,omitempty"`
	NovelSourceId        string   `protobuf:"bytes,3,opt,name=novel_source_id,json=novelSourceId,proto3" json:"novel_source_id,omitempty"`
	Preference           int32    `protobuf:"varint,4,opt,name=preference,proto3" json:"preference,omitempty"`
	Vip                  bool     `protobuf:"varint,5,opt,name=vip,proto3" json:"vip,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *NovelSourceMetadata) Reset()         { *m = NovelSourceMetadata{} }
func (m *NovelSourceMetadata) String() string { return proto.CompactTextString(m) }
func (*NovelSourceMetadata) ProtoMessage()    {}
func (*NovelSourceMetadata) Descriptor() ([]byte, []int) {
	return fileDescriptor_model_406c922677d5b648, []int{5}
}
func (m *NovelSourceMetadata) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_NovelSourceMetadata.Unmarshal(m, b)
}
func (m *NovelSourceMetadata) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_NovelSourceMetadata.Marshal(b, m, deterministic)
}
func (dst *NovelSourceMetadata) XXX_Merge(src proto.Message) {
	xxx_messageInfo_NovelSourceMetadata.Merge(dst, src)
}
func (m *NovelSourceMetadata) XXX_Size() int {
	return xxx_messageInfo_NovelSourceMetadata.Size(m)
}
func (m *NovelSourceMetadata) XXX_DiscardUnknown() {
	xxx_messageInfo_NovelSourceMetadata.DiscardUnknown(m)
}

var xxx_messageInfo_NovelSourceMetadata proto.InternalMessageInfo

func (m *NovelSourceMetadata) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

func (m *NovelSourceMetadata) GetSourceId() string {
	if m != nil {
		return m.SourceId
	}
	return ""
}

func (m *NovelSourceMetadata) GetNovelSourceId() string {
	if m != nil {
		return m.NovelSourceId
	}
	return ""
}

func (m *NovelSourceMetadata) GetPreference() int32 {
	if m != nil {
		return m.Preference
	}
	return 0
}

func (m *NovelSourceMetadata) GetVip() bool {
	if m != nil {
		return m.Vip
	}
	return false
}

type ChapterSourceData struct {
	Metadata             *ChapterSourceMetadata `protobuf:"bytes,1,opt,name=metadata,proto3" json:"metadata,omitempty"`
	Content              string                 `protobuf:"bytes,2,opt,name=content,proto3" json:"content,omitempty"`
	XXX_NoUnkeyedLiteral struct{}               `json:"-"`
	XXX_unrecognized     []byte                 `json:"-"`
	XXX_sizecache        int32                  `json:"-"`
}

func (m *ChapterSourceData) Reset()         { *m = ChapterSourceData{} }
func (m *ChapterSourceData) String() string { return proto.CompactTextString(m) }
func (*ChapterSourceData) ProtoMessage()    {}
func (*ChapterSourceData) Descriptor() ([]byte, []int) {
	return fileDescriptor_model_406c922677d5b648, []int{6}
}
func (m *ChapterSourceData) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ChapterSourceData.Unmarshal(m, b)
}
func (m *ChapterSourceData) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ChapterSourceData.Marshal(b, m, deterministic)
}
func (dst *ChapterSourceData) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ChapterSourceData.Merge(dst, src)
}
func (m *ChapterSourceData) XXX_Size() int {
	return xxx_messageInfo_ChapterSourceData.Size(m)
}
func (m *ChapterSourceData) XXX_DiscardUnknown() {
	xxx_messageInfo_ChapterSourceData.DiscardUnknown(m)
}

var xxx_messageInfo_ChapterSourceData proto.InternalMessageInfo

func (m *ChapterSourceData) GetMetadata() *ChapterSourceMetadata {
	if m != nil {
		return m.Metadata
	}
	return nil
}

func (m *ChapterSourceData) GetContent() string {
	if m != nil {
		return m.Content
	}
	return ""
}

type ChapterSourceMetadata struct {
	Id                   string   `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	ChapterSourceId      string   `protobuf:"bytes,2,opt,name=chapter_source_id,json=chapterSourceId,proto3" json:"chapter_source_id,omitempty"`
	Url                  string   `protobuf:"bytes,3,opt,name=url,proto3" json:"url,omitempty"`
	Index                int32    `protobuf:"varint,4,opt,name=index,proto3" json:"index,omitempty"`
	Length               int32    `protobuf:"varint,5,opt,name=length,proto3" json:"length,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ChapterSourceMetadata) Reset()         { *m = ChapterSourceMetadata{} }
func (m *ChapterSourceMetadata) String() string { return proto.CompactTextString(m) }
func (*ChapterSourceMetadata) ProtoMessage()    {}
func (*ChapterSourceMetadata) Descriptor() ([]byte, []int) {
	return fileDescriptor_model_406c922677d5b648, []int{7}
}
func (m *ChapterSourceMetadata) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ChapterSourceMetadata.Unmarshal(m, b)
}
func (m *ChapterSourceMetadata) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ChapterSourceMetadata.Marshal(b, m, deterministic)
}
func (dst *ChapterSourceMetadata) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ChapterSourceMetadata.Merge(dst, src)
}
func (m *ChapterSourceMetadata) XXX_Size() int {
	return xxx_messageInfo_ChapterSourceMetadata.Size(m)
}
func (m *ChapterSourceMetadata) XXX_DiscardUnknown() {
	xxx_messageInfo_ChapterSourceMetadata.DiscardUnknown(m)
}

var xxx_messageInfo_ChapterSourceMetadata proto.InternalMessageInfo

func (m *ChapterSourceMetadata) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

func (m *ChapterSourceMetadata) GetChapterSourceId() string {
	if m != nil {
		return m.ChapterSourceId
	}
	return ""
}

func (m *ChapterSourceMetadata) GetUrl() string {
	if m != nil {
		return m.Url
	}
	return ""
}

func (m *ChapterSourceMetadata) GetIndex() int32 {
	if m != nil {
		return m.Index
	}
	return 0
}

func (m *ChapterSourceMetadata) GetLength() int32 {
	if m != nil {
		return m.Length
	}
	return 0
}

func init() {
	proto.RegisterType((*Novel)(nil), "model.Novel")
	proto.RegisterType((*NovelMetadata)(nil), "model.NovelMetadata")
	proto.RegisterType((*Chapter)(nil), "model.Chapter")
	proto.RegisterType((*ChapterMetadata)(nil), "model.ChapterMetadata")
	proto.RegisterType((*NovelSourceData)(nil), "model.NovelSourceData")
	proto.RegisterType((*NovelSourceMetadata)(nil), "model.NovelSourceMetadata")
	proto.RegisterType((*ChapterSourceData)(nil), "model.ChapterSourceData")
	proto.RegisterType((*ChapterSourceMetadata)(nil), "model.ChapterSourceMetadata")
}

func init() { proto.RegisterFile("model.proto", fileDescriptor_model_406c922677d5b648) }

var fileDescriptor_model_406c922677d5b648 = []byte{
	// 498 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x8c, 0x54, 0x4d, 0x8f, 0xd3, 0x30,
	0x10, 0x55, 0xda, 0xcd, 0xb6, 0x3b, 0xd5, 0xb6, 0x5d, 0x53, 0x56, 0x11, 0x20, 0x54, 0xe5, 0x80,
	0x2a, 0x0e, 0x8d, 0xb4, 0x48, 0x08, 0x89, 0xd3, 0x52, 0x2e, 0x1c, 0xca, 0x21, 0xbb, 0x27, 0x2e,
	0x95, 0x1b, 0x9b, 0xc6, 0xc2, 0xb5, 0x23, 0xd7, 0x59, 0x71, 0xe7, 0x07, 0x70, 0xe0, 0xca, 0xbf,
	0xe0, 0x0f, 0xa2, 0xd8, 0xce, 0x67, 0xcb, 0xc7, 0xcd, 0x6f, 0x3c, 0x6f, 0xc6, 0xef, 0xcd, 0x24,
	0x30, 0xda, 0x4b, 0x42, 0xf9, 0x32, 0x53, 0x52, 0x4b, 0xe4, 0x1b, 0x10, 0xfe, 0xf2, 0xc0, 0xff,
	0x28, 0x1f, 0x28, 0x47, 0x6f, 0x61, 0x2c, 0x8a, 0xc3, 0x66, 0x4f, 0x35, 0x26, 0x58, 0xe3, 0xc0,
	0x9b, 0x7b, 0x8b, 0xd1, 0xcd, 0x6c, 0x69, 0x69, 0x26, 0x6b, 0xed, 0xee, 0xe2, 0x4b, 0xd1, 0x84,
	0x68, 0x05, 0x93, 0x83, 0xcc, 0x55, 0x42, 0x6b, 0x76, 0x6f, 0xde, 0x5f, 0x8c, 0x6e, 0x9e, 0x34,
	0xd9, 0x77, 0x26, 0xa5, 0xaa, 0x31, 0x3e, 0xb4, 0x30, 0x5a, 0xc0, 0x20, 0x49, 0x71, 0xa6, 0xa9,
	0x0a, 0xfa, 0x86, 0x3c, 0x76, 0xe4, 0x95, 0x8d, 0xc6, 0xe5, 0x75, 0xb8, 0x86, 0xcb, 0xd6, 0x73,
	0xd0, 0x18, 0x7a, 0x8c, 0x98, 0x07, 0x5f, 0xc4, 0x3d, 0x46, 0xd0, 0x0c, 0x7c, 0xcd, 0x34, 0xa7,
	0x41, 0xcf, 0x84, 0x2c, 0x40, 0xd7, 0x70, 0x8e, 0x73, 0x9d, 0xca, 0xa2, 0x7e, 0x11, 0x76, 0x28,
	0xfc, 0xee, 0xc1, 0xc0, 0xf5, 0x40, 0xb7, 0x30, 0x75, 0x5d, 0xba, 0x46, 0x5c, 0xb7, 0x5f, 0x53,
	0xc9, 0x98, 0x24, 0xed, 0x00, 0xba, 0xfd, 0x93, 0x19, 0x41, 0xbb, 0x82, 0xb5, 0xe3, 0xfd, 0x09,
	0x2b, 0xc2, 0x35, 0x4c, 0x3a, 0x6d, 0xfe, 0x53, 0xe2, 0x0c, 0x7c, 0x26, 0x08, 0xfd, 0x6a, 0x14,
	0xfa, 0xb1, 0x05, 0xe1, 0x37, 0x0f, 0x26, 0x8d, 0x09, 0x14, 0x2d, 0xd1, 0x6b, 0x18, 0x76, 0x04,
	0xfe, 0x6d, 0x56, 0x55, 0x2e, 0x7a, 0x03, 0x43, 0x27, 0xf8, 0xe0, 0x64, 0x3d, 0x3b, 0x25, 0xab,
	0x66, 0x96, 0xd9, 0xe1, 0x4f, 0x0f, 0x1e, 0x9d, 0xa8, 0x7d, 0xa4, 0xec, 0x29, 0x5c, 0x38, 0xff,
	0x18, 0x71, 0xea, 0x86, 0x36, 0xf0, 0x81, 0xa0, 0x17, 0x30, 0xb1, 0x6b, 0x5a, 0xa7, 0xd8, 0x61,
	0xda, 0x8d, 0xbc, 0x2b, 0xf3, 0x9e, 0x03, 0x64, 0x8a, 0x7e, 0xa6, 0x8a, 0x8a, 0x84, 0x06, 0x67,
	0xc6, 0x8d, 0x46, 0x04, 0x4d, 0xa1, 0xff, 0xc0, 0xb2, 0xc0, 0x9f, 0x7b, 0x8b, 0x61, 0x5c, 0x1c,
	0xc3, 0x1d, 0x5c, 0x1d, 0x0d, 0xa6, 0x50, 0xdb, 0x71, 0xe9, 0x1f, 0x6a, 0x2b, 0x9f, 0x02, 0x18,
	0x24, 0x52, 0x68, 0x2a, 0xb4, 0xd3, 0x50, 0xc2, 0xf0, 0x87, 0x07, 0x8f, 0x4f, 0xb2, 0x8f, 0x9c,
	0x78, 0x09, 0x57, 0xe5, 0x32, 0x76, 0x1d, 0x29, 0xb7, 0xae, 0x12, 0x3c, 0x85, 0x7e, 0xae, 0xb8,
	0x33, 0xa3, 0x38, 0xd6, 0xbb, 0x70, 0xd6, 0xd8, 0x85, 0xe2, 0x23, 0xe0, 0x54, 0xec, 0x74, 0x6a,
	0xb4, 0xfb, 0xb1, 0x43, 0xef, 0xbe, 0x7c, 0x62, 0x3b, 0xa6, 0xd3, 0x7c, 0xbb, 0x4c, 0xe4, 0x3e,
	0xba, 0x4f, 0xe9, 0x3d, 0x55, 0x8a, 0x6d, 0x39, 0x5d, 0xa5, 0x8c, 0x93, 0x28, 0xe1, 0x32, 0x27,
	0x1b, 0x9c, 0x65, 0x19, 0x67, 0x09, 0xd6, 0x4c, 0x8a, 0x4d, 0x26, 0x95, 0xc6, 0xbc, 0xbe, 0x29,
	0x2f, 0x0e, 0x91, 0x9d, 0x50, 0x23, 0x14, 0x31, 0xa1, 0xa9, 0x12, 0x98, 0x47, 0xc6, 0xbc, 0xed,
	0xb9, 0xf9, 0x09, 0xbd, 0xfa, 0x1d, 0x00, 0x00, 0xff, 0xff, 0xf0, 0xde, 0x1b, 0x8d, 0x93, 0x04,
	0x00, 0x00,
}