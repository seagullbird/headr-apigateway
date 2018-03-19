// Code generated by protoc-gen-go. DO NOT EDIT.
// source: sitemgrsvc.proto

/*
Package pb is a generated protocol buffer package.

It is generated from these files:
	sitemgrsvc.proto

It has these top-level messages:
	CreateNewSiteRequest
	CreateNewSiteReply
	ProxyDeleteSiteRequest
	ProxyDeleteSiteReply
	CheckSitenameExistsRequest
	CheckSitenameExistsReply
	GetSiteIDByUserIDRequest
	GetSiteIDByUserIDReply
*/
package pb

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

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

type CreateNewSiteRequest struct {
	Sitename string `protobuf:"bytes,1,opt,name=sitename" json:"sitename,omitempty"`
}

func (m *CreateNewSiteRequest) Reset()                    { *m = CreateNewSiteRequest{} }
func (m *CreateNewSiteRequest) String() string            { return proto.CompactTextString(m) }
func (*CreateNewSiteRequest) ProtoMessage()               {}
func (*CreateNewSiteRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

func (m *CreateNewSiteRequest) GetSitename() string {
	if m != nil {
		return m.Sitename
	}
	return ""
}

type CreateNewSiteReply struct {
	SiteId uint64 `protobuf:"varint,1,opt,name=site_id,json=siteId" json:"site_id,omitempty"`
	Err    string `protobuf:"bytes,2,opt,name=err" json:"err,omitempty"`
}

func (m *CreateNewSiteReply) Reset()                    { *m = CreateNewSiteReply{} }
func (m *CreateNewSiteReply) String() string            { return proto.CompactTextString(m) }
func (*CreateNewSiteReply) ProtoMessage()               {}
func (*CreateNewSiteReply) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

func (m *CreateNewSiteReply) GetSiteId() uint64 {
	if m != nil {
		return m.SiteId
	}
	return 0
}

func (m *CreateNewSiteReply) GetErr() string {
	if m != nil {
		return m.Err
	}
	return ""
}

type ProxyDeleteSiteRequest struct {
	SiteId uint64 `protobuf:"varint,1,opt,name=site_id,json=siteId" json:"site_id,omitempty"`
}

func (m *ProxyDeleteSiteRequest) Reset()                    { *m = ProxyDeleteSiteRequest{} }
func (m *ProxyDeleteSiteRequest) String() string            { return proto.CompactTextString(m) }
func (*ProxyDeleteSiteRequest) ProtoMessage()               {}
func (*ProxyDeleteSiteRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{2} }

func (m *ProxyDeleteSiteRequest) GetSiteId() uint64 {
	if m != nil {
		return m.SiteId
	}
	return 0
}

type ProxyDeleteSiteReply struct {
	Err string `protobuf:"bytes,1,opt,name=err" json:"err,omitempty"`
}

func (m *ProxyDeleteSiteReply) Reset()                    { *m = ProxyDeleteSiteReply{} }
func (m *ProxyDeleteSiteReply) String() string            { return proto.CompactTextString(m) }
func (*ProxyDeleteSiteReply) ProtoMessage()               {}
func (*ProxyDeleteSiteReply) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{3} }

func (m *ProxyDeleteSiteReply) GetErr() string {
	if m != nil {
		return m.Err
	}
	return ""
}

type CheckSitenameExistsRequest struct {
	Sitename string `protobuf:"bytes,1,opt,name=sitename" json:"sitename,omitempty"`
}

func (m *CheckSitenameExistsRequest) Reset()                    { *m = CheckSitenameExistsRequest{} }
func (m *CheckSitenameExistsRequest) String() string            { return proto.CompactTextString(m) }
func (*CheckSitenameExistsRequest) ProtoMessage()               {}
func (*CheckSitenameExistsRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{4} }

func (m *CheckSitenameExistsRequest) GetSitename() string {
	if m != nil {
		return m.Sitename
	}
	return ""
}

type CheckSitenameExistsReply struct {
	Exists bool   `protobuf:"varint,1,opt,name=exists" json:"exists,omitempty"`
	Err    string `protobuf:"bytes,2,opt,name=err" json:"err,omitempty"`
}

func (m *CheckSitenameExistsReply) Reset()                    { *m = CheckSitenameExistsReply{} }
func (m *CheckSitenameExistsReply) String() string            { return proto.CompactTextString(m) }
func (*CheckSitenameExistsReply) ProtoMessage()               {}
func (*CheckSitenameExistsReply) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{5} }

func (m *CheckSitenameExistsReply) GetExists() bool {
	if m != nil {
		return m.Exists
	}
	return false
}

func (m *CheckSitenameExistsReply) GetErr() string {
	if m != nil {
		return m.Err
	}
	return ""
}

type GetSiteIDByUserIDRequest struct {
}

func (m *GetSiteIDByUserIDRequest) Reset()                    { *m = GetSiteIDByUserIDRequest{} }
func (m *GetSiteIDByUserIDRequest) String() string            { return proto.CompactTextString(m) }
func (*GetSiteIDByUserIDRequest) ProtoMessage()               {}
func (*GetSiteIDByUserIDRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{6} }

type GetSiteIDByUserIDReply struct {
	SiteId uint64 `protobuf:"varint,1,opt,name=site_id,json=siteId" json:"site_id,omitempty"`
	Err    string `protobuf:"bytes,2,opt,name=err" json:"err,omitempty"`
}

func (m *GetSiteIDByUserIDReply) Reset()                    { *m = GetSiteIDByUserIDReply{} }
func (m *GetSiteIDByUserIDReply) String() string            { return proto.CompactTextString(m) }
func (*GetSiteIDByUserIDReply) ProtoMessage()               {}
func (*GetSiteIDByUserIDReply) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{7} }

func (m *GetSiteIDByUserIDReply) GetSiteId() uint64 {
	if m != nil {
		return m.SiteId
	}
	return 0
}

func (m *GetSiteIDByUserIDReply) GetErr() string {
	if m != nil {
		return m.Err
	}
	return ""
}

func init() {
	proto.RegisterType((*CreateNewSiteRequest)(nil), "pb.CreateNewSiteRequest")
	proto.RegisterType((*CreateNewSiteReply)(nil), "pb.CreateNewSiteReply")
	proto.RegisterType((*ProxyDeleteSiteRequest)(nil), "pb.ProxyDeleteSiteRequest")
	proto.RegisterType((*ProxyDeleteSiteReply)(nil), "pb.ProxyDeleteSiteReply")
	proto.RegisterType((*CheckSitenameExistsRequest)(nil), "pb.CheckSitenameExistsRequest")
	proto.RegisterType((*CheckSitenameExistsReply)(nil), "pb.CheckSitenameExistsReply")
	proto.RegisterType((*GetSiteIDByUserIDRequest)(nil), "pb.GetSiteIDByUserIDRequest")
	proto.RegisterType((*GetSiteIDByUserIDReply)(nil), "pb.GetSiteIDByUserIDReply")
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// Client API for Sitemgr service

type SitemgrClient interface {
	NewSite(ctx context.Context, in *CreateNewSiteRequest, opts ...grpc.CallOption) (*CreateNewSiteReply, error)
	DeleteSite(ctx context.Context, in *ProxyDeleteSiteRequest, opts ...grpc.CallOption) (*ProxyDeleteSiteReply, error)
	CheckSitenameExists(ctx context.Context, in *CheckSitenameExistsRequest, opts ...grpc.CallOption) (*CheckSitenameExistsReply, error)
	GetSiteIDByUserID(ctx context.Context, in *GetSiteIDByUserIDRequest, opts ...grpc.CallOption) (*GetSiteIDByUserIDReply, error)
}

type sitemgrClient struct {
	cc *grpc.ClientConn
}

func NewSitemgrClient(cc *grpc.ClientConn) SitemgrClient {
	return &sitemgrClient{cc}
}

func (c *sitemgrClient) NewSite(ctx context.Context, in *CreateNewSiteRequest, opts ...grpc.CallOption) (*CreateNewSiteReply, error) {
	out := new(CreateNewSiteReply)
	err := grpc.Invoke(ctx, "/pb.Sitemgr/NewSite", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *sitemgrClient) DeleteSite(ctx context.Context, in *ProxyDeleteSiteRequest, opts ...grpc.CallOption) (*ProxyDeleteSiteReply, error) {
	out := new(ProxyDeleteSiteReply)
	err := grpc.Invoke(ctx, "/pb.Sitemgr/DeleteSite", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *sitemgrClient) CheckSitenameExists(ctx context.Context, in *CheckSitenameExistsRequest, opts ...grpc.CallOption) (*CheckSitenameExistsReply, error) {
	out := new(CheckSitenameExistsReply)
	err := grpc.Invoke(ctx, "/pb.Sitemgr/CheckSitenameExists", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *sitemgrClient) GetSiteIDByUserID(ctx context.Context, in *GetSiteIDByUserIDRequest, opts ...grpc.CallOption) (*GetSiteIDByUserIDReply, error) {
	out := new(GetSiteIDByUserIDReply)
	err := grpc.Invoke(ctx, "/pb.Sitemgr/GetSiteIDByUserID", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for Sitemgr service

type SitemgrServer interface {
	NewSite(context.Context, *CreateNewSiteRequest) (*CreateNewSiteReply, error)
	DeleteSite(context.Context, *ProxyDeleteSiteRequest) (*ProxyDeleteSiteReply, error)
	CheckSitenameExists(context.Context, *CheckSitenameExistsRequest) (*CheckSitenameExistsReply, error)
	GetSiteIDByUserID(context.Context, *GetSiteIDByUserIDRequest) (*GetSiteIDByUserIDReply, error)
}

func RegisterSitemgrServer(s *grpc.Server, srv SitemgrServer) {
	s.RegisterService(&_Sitemgr_serviceDesc, srv)
}

func _Sitemgr_NewSite_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateNewSiteRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SitemgrServer).NewSite(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pb.Sitemgr/NewSite",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SitemgrServer).NewSite(ctx, req.(*CreateNewSiteRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Sitemgr_DeleteSite_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ProxyDeleteSiteRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SitemgrServer).DeleteSite(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pb.Sitemgr/DeleteSite",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SitemgrServer).DeleteSite(ctx, req.(*ProxyDeleteSiteRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Sitemgr_CheckSitenameExists_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CheckSitenameExistsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SitemgrServer).CheckSitenameExists(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pb.Sitemgr/CheckSitenameExists",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SitemgrServer).CheckSitenameExists(ctx, req.(*CheckSitenameExistsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Sitemgr_GetSiteIDByUserID_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetSiteIDByUserIDRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SitemgrServer).GetSiteIDByUserID(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pb.Sitemgr/GetSiteIDByUserID",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SitemgrServer).GetSiteIDByUserID(ctx, req.(*GetSiteIDByUserIDRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _Sitemgr_serviceDesc = grpc.ServiceDesc{
	ServiceName: "pb.Sitemgr",
	HandlerType: (*SitemgrServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "NewSite",
			Handler:    _Sitemgr_NewSite_Handler,
		},
		{
			MethodName: "DeleteSite",
			Handler:    _Sitemgr_DeleteSite_Handler,
		},
		{
			MethodName: "CheckSitenameExists",
			Handler:    _Sitemgr_CheckSitenameExists_Handler,
		},
		{
			MethodName: "GetSiteIDByUserID",
			Handler:    _Sitemgr_GetSiteIDByUserID_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "sitemgrsvc.proto",
}

func init() { proto.RegisterFile("sitemgrsvc.proto", fileDescriptor0) }

var fileDescriptor0 = []byte{
	// 311 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x94, 0x52, 0xc1, 0x4a, 0xc3, 0x40,
	0x14, 0x34, 0x51, 0x92, 0xfa, 0x4e, 0x75, 0x2d, 0x71, 0x59, 0x44, 0x64, 0x4f, 0x3d, 0x05, 0xac,
	0x17, 0x2f, 0x22, 0x98, 0x88, 0xe4, 0xa2, 0x92, 0xd0, 0xb3, 0x34, 0xed, 0x43, 0x83, 0xa9, 0x89,
	0xbb, 0xab, 0x36, 0x3f, 0xe6, 0xf7, 0xc9, 0x6e, 0x82, 0x15, 0xbb, 0x01, 0x7b, 0xcb, 0x63, 0xde,
	0xcc, 0xe4, 0xcd, 0x2c, 0x0c, 0x65, 0xa1, 0x70, 0xf9, 0x24, 0xe4, 0xc7, 0x3c, 0xac, 0x45, 0xa5,
	0x2a, 0xe2, 0xd6, 0x39, 0x9f, 0xc0, 0x28, 0x12, 0x38, 0x53, 0x78, 0x87, 0x9f, 0x59, 0xa1, 0x30,
	0xc5, 0xb7, 0x77, 0x94, 0x8a, 0x30, 0x18, 0xe8, 0xfd, 0xd7, 0xd9, 0x12, 0xa9, 0x73, 0xea, 0x8c,
	0xf7, 0xd3, 0x9f, 0x99, 0x5f, 0x01, 0xf9, 0xc3, 0xa9, 0xcb, 0x86, 0x1c, 0x81, 0xaf, 0x37, 0x1e,
	0x8b, 0x85, 0x21, 0xec, 0xa5, 0x9e, 0x1e, 0x93, 0x05, 0x19, 0xc2, 0x2e, 0x0a, 0x41, 0x5d, 0xa3,
	0xa2, 0x3f, 0xf9, 0x19, 0x04, 0x0f, 0xa2, 0x5a, 0x35, 0x31, 0x96, 0xa8, 0xf0, 0xb7, 0x6d, 0x9f,
	0x08, 0x1f, 0xc3, 0x68, 0x83, 0xa2, 0x5d, 0x3b, 0x71, 0x67, 0x2d, 0x7e, 0x01, 0x2c, 0x7a, 0xc6,
	0xf9, 0x4b, 0xd6, 0xfd, 0xee, 0xcd, 0xaa, 0x90, 0x4a, 0xfe, 0xe7, 0xae, 0x18, 0xa8, 0x95, 0xa9,
	0x7d, 0x02, 0xf0, 0xd0, 0x8c, 0x86, 0x35, 0x48, 0xbb, 0xc9, 0x72, 0x1c, 0x03, 0x7a, 0x8b, 0x4a,
	0x6b, 0x24, 0xf1, 0x75, 0x33, 0x95, 0x28, 0x92, 0xb8, 0x73, 0xe7, 0x11, 0x04, 0x16, 0x6c, 0xbb,
	0xf4, 0x26, 0x5f, 0x2e, 0xf8, 0x59, 0xdb, 0x25, 0xb9, 0x04, 0xbf, 0x2b, 0x81, 0xd0, 0xb0, 0xce,
	0x43, 0x5b, 0x97, 0x2c, 0xb0, 0x20, 0x75, 0xd9, 0xf0, 0x1d, 0x12, 0x03, 0xac, 0x03, 0x25, 0x4c,
	0xef, 0xd9, 0x8b, 0x61, 0xd4, 0x8a, 0xb5, 0x2a, 0x53, 0x38, 0xb4, 0xe4, 0x46, 0x4e, 0x8c, 0x6d,
	0x6f, 0x15, 0xec, 0xb8, 0x17, 0x6f, 0x65, 0xef, 0xe1, 0x60, 0x23, 0x2c, 0x62, 0x48, 0x7d, 0xf9,
	0x32, 0xd6, 0x83, 0x1a, 0xc1, 0xdc, 0x33, 0xcf, 0xfe, 0xfc, 0x3b, 0x00, 0x00, 0xff, 0xff, 0x0a,
	0x00, 0x6f, 0x2d, 0x0a, 0x03, 0x00, 0x00,
}
