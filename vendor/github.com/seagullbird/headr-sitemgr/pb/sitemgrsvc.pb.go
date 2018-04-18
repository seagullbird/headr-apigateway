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
	GetConfigRequest
	GetConfigReply
	UpdateConfigRequest
	UpdateConfigReply
	GetThemesRequest
	GetThemesReply
	UpdateSiteThemeRequest
	UpdateSiteThemeReply
	PostAboutRequest
	PostAboutReply
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

type GetConfigRequest struct {
	SiteId uint64 `protobuf:"varint,1,opt,name=site_id,json=siteId" json:"site_id,omitempty"`
}

func (m *GetConfigRequest) Reset()                    { *m = GetConfigRequest{} }
func (m *GetConfigRequest) String() string            { return proto.CompactTextString(m) }
func (*GetConfigRequest) ProtoMessage()               {}
func (*GetConfigRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{8} }

func (m *GetConfigRequest) GetSiteId() uint64 {
	if m != nil {
		return m.SiteId
	}
	return 0
}

type GetConfigReply struct {
	Config string `protobuf:"bytes,1,opt,name=config" json:"config,omitempty"`
	Err    string `protobuf:"bytes,2,opt,name=err" json:"err,omitempty"`
}

func (m *GetConfigReply) Reset()                    { *m = GetConfigReply{} }
func (m *GetConfigReply) String() string            { return proto.CompactTextString(m) }
func (*GetConfigReply) ProtoMessage()               {}
func (*GetConfigReply) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{9} }

func (m *GetConfigReply) GetConfig() string {
	if m != nil {
		return m.Config
	}
	return ""
}

func (m *GetConfigReply) GetErr() string {
	if m != nil {
		return m.Err
	}
	return ""
}

type UpdateConfigRequest struct {
	SiteId uint64 `protobuf:"varint,1,opt,name=site_id,json=siteId" json:"site_id,omitempty"`
	Config string `protobuf:"bytes,2,opt,name=config" json:"config,omitempty"`
}

func (m *UpdateConfigRequest) Reset()                    { *m = UpdateConfigRequest{} }
func (m *UpdateConfigRequest) String() string            { return proto.CompactTextString(m) }
func (*UpdateConfigRequest) ProtoMessage()               {}
func (*UpdateConfigRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{10} }

func (m *UpdateConfigRequest) GetSiteId() uint64 {
	if m != nil {
		return m.SiteId
	}
	return 0
}

func (m *UpdateConfigRequest) GetConfig() string {
	if m != nil {
		return m.Config
	}
	return ""
}

type UpdateConfigReply struct {
	Err string `protobuf:"bytes,1,opt,name=err" json:"err,omitempty"`
}

func (m *UpdateConfigReply) Reset()                    { *m = UpdateConfigReply{} }
func (m *UpdateConfigReply) String() string            { return proto.CompactTextString(m) }
func (*UpdateConfigReply) ProtoMessage()               {}
func (*UpdateConfigReply) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{11} }

func (m *UpdateConfigReply) GetErr() string {
	if m != nil {
		return m.Err
	}
	return ""
}

type GetThemesRequest struct {
	SiteId uint64 `protobuf:"varint,1,opt,name=site_id,json=siteId" json:"site_id,omitempty"`
}

func (m *GetThemesRequest) Reset()                    { *m = GetThemesRequest{} }
func (m *GetThemesRequest) String() string            { return proto.CompactTextString(m) }
func (*GetThemesRequest) ProtoMessage()               {}
func (*GetThemesRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{12} }

func (m *GetThemesRequest) GetSiteId() uint64 {
	if m != nil {
		return m.SiteId
	}
	return 0
}

type GetThemesReply struct {
	Themes string `protobuf:"bytes,1,opt,name=themes" json:"themes,omitempty"`
	Err    string `protobuf:"bytes,2,opt,name=err" json:"err,omitempty"`
}

func (m *GetThemesReply) Reset()                    { *m = GetThemesReply{} }
func (m *GetThemesReply) String() string            { return proto.CompactTextString(m) }
func (*GetThemesReply) ProtoMessage()               {}
func (*GetThemesReply) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{13} }

func (m *GetThemesReply) GetThemes() string {
	if m != nil {
		return m.Themes
	}
	return ""
}

func (m *GetThemesReply) GetErr() string {
	if m != nil {
		return m.Err
	}
	return ""
}

type UpdateSiteThemeRequest struct {
	SiteId uint64 `protobuf:"varint,1,opt,name=site_id,json=siteId" json:"site_id,omitempty"`
	Theme  string `protobuf:"bytes,2,opt,name=theme" json:"theme,omitempty"`
}

func (m *UpdateSiteThemeRequest) Reset()                    { *m = UpdateSiteThemeRequest{} }
func (m *UpdateSiteThemeRequest) String() string            { return proto.CompactTextString(m) }
func (*UpdateSiteThemeRequest) ProtoMessage()               {}
func (*UpdateSiteThemeRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{14} }

func (m *UpdateSiteThemeRequest) GetSiteId() uint64 {
	if m != nil {
		return m.SiteId
	}
	return 0
}

func (m *UpdateSiteThemeRequest) GetTheme() string {
	if m != nil {
		return m.Theme
	}
	return ""
}

type UpdateSiteThemeReply struct {
	Err string `protobuf:"bytes,1,opt,name=err" json:"err,omitempty"`
}

func (m *UpdateSiteThemeReply) Reset()                    { *m = UpdateSiteThemeReply{} }
func (m *UpdateSiteThemeReply) String() string            { return proto.CompactTextString(m) }
func (*UpdateSiteThemeReply) ProtoMessage()               {}
func (*UpdateSiteThemeReply) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{15} }

func (m *UpdateSiteThemeReply) GetErr() string {
	if m != nil {
		return m.Err
	}
	return ""
}

type PostAboutRequest struct {
	SiteId  uint64 `protobuf:"varint,1,opt,name=site_id,json=siteId" json:"site_id,omitempty"`
	Content string `protobuf:"bytes,2,opt,name=content" json:"content,omitempty"`
}

func (m *PostAboutRequest) Reset()                    { *m = PostAboutRequest{} }
func (m *PostAboutRequest) String() string            { return proto.CompactTextString(m) }
func (*PostAboutRequest) ProtoMessage()               {}
func (*PostAboutRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{16} }

func (m *PostAboutRequest) GetSiteId() uint64 {
	if m != nil {
		return m.SiteId
	}
	return 0
}

func (m *PostAboutRequest) GetContent() string {
	if m != nil {
		return m.Content
	}
	return ""
}

type PostAboutReply struct {
	Err string `protobuf:"bytes,1,opt,name=err" json:"err,omitempty"`
}

func (m *PostAboutReply) Reset()                    { *m = PostAboutReply{} }
func (m *PostAboutReply) String() string            { return proto.CompactTextString(m) }
func (*PostAboutReply) ProtoMessage()               {}
func (*PostAboutReply) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{17} }

func (m *PostAboutReply) GetErr() string {
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
	proto.RegisterType((*GetConfigRequest)(nil), "pb.GetConfigRequest")
	proto.RegisterType((*GetConfigReply)(nil), "pb.GetConfigReply")
	proto.RegisterType((*UpdateConfigRequest)(nil), "pb.UpdateConfigRequest")
	proto.RegisterType((*UpdateConfigReply)(nil), "pb.UpdateConfigReply")
	proto.RegisterType((*GetThemesRequest)(nil), "pb.GetThemesRequest")
	proto.RegisterType((*GetThemesReply)(nil), "pb.GetThemesReply")
	proto.RegisterType((*UpdateSiteThemeRequest)(nil), "pb.UpdateSiteThemeRequest")
	proto.RegisterType((*UpdateSiteThemeReply)(nil), "pb.UpdateSiteThemeReply")
	proto.RegisterType((*PostAboutRequest)(nil), "pb.PostAboutRequest")
	proto.RegisterType((*PostAboutReply)(nil), "pb.PostAboutReply")
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
	GetConfig(ctx context.Context, in *GetConfigRequest, opts ...grpc.CallOption) (*GetConfigReply, error)
	UpdateConfig(ctx context.Context, in *UpdateConfigRequest, opts ...grpc.CallOption) (*UpdateConfigReply, error)
	GetThemes(ctx context.Context, in *GetThemesRequest, opts ...grpc.CallOption) (*GetThemesReply, error)
	UpdateSiteTheme(ctx context.Context, in *UpdateSiteThemeRequest, opts ...grpc.CallOption) (*UpdateSiteThemeReply, error)
	PostAbout(ctx context.Context, in *PostAboutRequest, opts ...grpc.CallOption) (*PostAboutReply, error)
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

func (c *sitemgrClient) GetConfig(ctx context.Context, in *GetConfigRequest, opts ...grpc.CallOption) (*GetConfigReply, error) {
	out := new(GetConfigReply)
	err := grpc.Invoke(ctx, "/pb.Sitemgr/GetConfig", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *sitemgrClient) UpdateConfig(ctx context.Context, in *UpdateConfigRequest, opts ...grpc.CallOption) (*UpdateConfigReply, error) {
	out := new(UpdateConfigReply)
	err := grpc.Invoke(ctx, "/pb.Sitemgr/UpdateConfig", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *sitemgrClient) GetThemes(ctx context.Context, in *GetThemesRequest, opts ...grpc.CallOption) (*GetThemesReply, error) {
	out := new(GetThemesReply)
	err := grpc.Invoke(ctx, "/pb.Sitemgr/GetThemes", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *sitemgrClient) UpdateSiteTheme(ctx context.Context, in *UpdateSiteThemeRequest, opts ...grpc.CallOption) (*UpdateSiteThemeReply, error) {
	out := new(UpdateSiteThemeReply)
	err := grpc.Invoke(ctx, "/pb.Sitemgr/UpdateSiteTheme", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *sitemgrClient) PostAbout(ctx context.Context, in *PostAboutRequest, opts ...grpc.CallOption) (*PostAboutReply, error) {
	out := new(PostAboutReply)
	err := grpc.Invoke(ctx, "/pb.Sitemgr/PostAbout", in, out, c.cc, opts...)
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
	GetConfig(context.Context, *GetConfigRequest) (*GetConfigReply, error)
	UpdateConfig(context.Context, *UpdateConfigRequest) (*UpdateConfigReply, error)
	GetThemes(context.Context, *GetThemesRequest) (*GetThemesReply, error)
	UpdateSiteTheme(context.Context, *UpdateSiteThemeRequest) (*UpdateSiteThemeReply, error)
	PostAbout(context.Context, *PostAboutRequest) (*PostAboutReply, error)
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

func _Sitemgr_GetConfig_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetConfigRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SitemgrServer).GetConfig(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pb.Sitemgr/GetConfig",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SitemgrServer).GetConfig(ctx, req.(*GetConfigRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Sitemgr_UpdateConfig_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateConfigRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SitemgrServer).UpdateConfig(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pb.Sitemgr/UpdateConfig",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SitemgrServer).UpdateConfig(ctx, req.(*UpdateConfigRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Sitemgr_GetThemes_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetThemesRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SitemgrServer).GetThemes(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pb.Sitemgr/GetThemes",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SitemgrServer).GetThemes(ctx, req.(*GetThemesRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Sitemgr_UpdateSiteTheme_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateSiteThemeRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SitemgrServer).UpdateSiteTheme(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pb.Sitemgr/UpdateSiteTheme",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SitemgrServer).UpdateSiteTheme(ctx, req.(*UpdateSiteThemeRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Sitemgr_PostAbout_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(PostAboutRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SitemgrServer).PostAbout(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pb.Sitemgr/PostAbout",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SitemgrServer).PostAbout(ctx, req.(*PostAboutRequest))
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
		{
			MethodName: "GetConfig",
			Handler:    _Sitemgr_GetConfig_Handler,
		},
		{
			MethodName: "UpdateConfig",
			Handler:    _Sitemgr_UpdateConfig_Handler,
		},
		{
			MethodName: "GetThemes",
			Handler:    _Sitemgr_GetThemes_Handler,
		},
		{
			MethodName: "UpdateSiteTheme",
			Handler:    _Sitemgr_UpdateSiteTheme_Handler,
		},
		{
			MethodName: "PostAbout",
			Handler:    _Sitemgr_PostAbout_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "sitemgrsvc.proto",
}

func init() { proto.RegisterFile("sitemgrsvc.proto", fileDescriptor0) }

var fileDescriptor0 = []byte{
	// 518 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x94, 0x55, 0x5d, 0x6b, 0x13, 0x41,
	0x14, 0x4d, 0x62, 0x4d, 0xda, 0x8b, 0xd4, 0x74, 0x1a, 0xb7, 0xcb, 0x20, 0x22, 0x03, 0x42, 0x41,
	0x08, 0x58, 0x1f, 0x14, 0x41, 0xfc, 0xc8, 0xd6, 0x90, 0x17, 0x2d, 0x89, 0x79, 0x96, 0x7c, 0x5c,
	0xdb, 0xc5, 0x24, 0xbb, 0xee, 0x4c, 0xb5, 0xf9, 0xd3, 0xfe, 0x06, 0x99, 0x8f, 0x6c, 0x26, 0xd9,
	0xbb, 0x24, 0x7d, 0xeb, 0xed, 0xbd, 0xf7, 0x9c, 0x7b, 0x76, 0xce, 0x21, 0xd0, 0x94, 0xb1, 0xc2,
	0xf9, 0x75, 0x26, 0xff, 0x4c, 0xda, 0x69, 0x96, 0xa8, 0x84, 0xd5, 0xd2, 0xb1, 0xb8, 0x80, 0x56,
	0x27, 0xc3, 0x91, 0xc2, 0xaf, 0xf8, 0x77, 0x10, 0x2b, 0xec, 0xe3, 0xef, 0x5b, 0x94, 0x8a, 0x71,
	0x38, 0xd4, 0xf3, 0x8b, 0xd1, 0x1c, 0xc3, 0xea, 0xf3, 0xea, 0xf9, 0x51, 0x3f, 0xaf, 0xc5, 0x07,
	0x60, 0x5b, 0x3b, 0xe9, 0x6c, 0xc9, 0xce, 0xa0, 0xa1, 0x27, 0x7e, 0xc4, 0x53, 0xb3, 0x70, 0xd0,
	0xaf, 0xeb, 0xb2, 0x37, 0x65, 0x4d, 0x78, 0x80, 0x59, 0x16, 0xd6, 0x0c, 0x8a, 0xfe, 0x53, 0xbc,
	0x82, 0xe0, 0x2a, 0x4b, 0xee, 0x96, 0x11, 0xce, 0x50, 0xa1, 0x4f, 0x5b, 0x06, 0x22, 0xce, 0xa1,
	0x55, 0x58, 0xd1, 0xac, 0x0e, 0xbc, 0xba, 0x06, 0x7f, 0x0b, 0xbc, 0x73, 0x83, 0x93, 0x5f, 0x03,
	0x77, 0xee, 0xe5, 0x5d, 0x2c, 0x95, 0xdc, 0x47, 0x57, 0x04, 0x21, 0xb9, 0xa9, 0x79, 0x02, 0xa8,
	0xa3, 0x29, 0xcd, 0xd6, 0x61, 0xdf, 0x55, 0x84, 0x38, 0x0e, 0x61, 0x17, 0x95, 0xc6, 0xe8, 0x45,
	0x9f, 0x97, 0x43, 0x89, 0x59, 0x2f, 0x72, 0xec, 0xa2, 0x03, 0x01, 0xd1, 0xbb, 0xe7, 0xd7, 0x7b,
	0x09, 0xcd, 0x2e, 0xaa, 0x4e, 0xb2, 0xf8, 0x19, 0x5f, 0xef, 0xfc, 0x6e, 0xef, 0xe0, 0xd8, 0x1b,
	0x76, 0x4a, 0x26, 0xa6, 0x74, 0xfa, 0x5d, 0x45, 0x10, 0x7d, 0x81, 0xd3, 0x61, 0x3a, 0x1d, 0x29,
	0xdc, 0x8f, 0xcb, 0x43, 0xae, 0xf9, 0xc8, 0xe2, 0x05, 0x9c, 0x6c, 0xe2, 0xd0, 0x0f, 0x67, 0x75,
	0x7d, 0xbf, 0xc1, 0x39, 0xca, 0x3d, 0x75, 0xad, 0x86, 0x9d, 0x2e, 0x65, 0xca, 0x95, 0x2e, 0x5b,
	0x11, 0xba, 0xba, 0x10, 0xd8, 0x7b, 0xf4, 0x43, 0x18, 0x88, 0x9d, 0xd2, 0x5a, 0xf0, 0xd0, 0xc0,
	0x39, 0x18, 0x5b, 0x68, 0x53, 0x16, 0x80, 0x68, 0x6d, 0x97, 0xd0, 0xbc, 0x4a, 0xa4, 0xfa, 0x34,
	0x4e, 0x6e, 0xd5, 0x4e, 0xb2, 0x10, 0x1a, 0x93, 0x64, 0xa1, 0x70, 0xa1, 0x1c, 0xdd, 0xaa, 0x14,
	0x02, 0x8e, 0x3d, 0x18, 0x92, 0xea, 0xe2, 0xdf, 0x01, 0x34, 0x06, 0x36, 0xea, 0xec, 0x3d, 0x34,
	0x5c, 0x46, 0x59, 0xd8, 0x4e, 0xc7, 0x6d, 0x2a, 0xea, 0x3c, 0x20, 0x3a, 0xe9, 0x6c, 0x29, 0x2a,
	0x2c, 0x02, 0x58, 0xe7, 0x8d, 0x71, 0x3d, 0x47, 0xe7, 0x96, 0x87, 0x64, 0xcf, 0xa2, 0x0c, 0xe1,
	0x94, 0x88, 0x15, 0x7b, 0x66, 0x68, 0x4b, 0x93, 0xca, 0x9f, 0x96, 0xf6, 0x2d, 0xec, 0x37, 0x38,
	0x29, 0x64, 0x89, 0x99, 0xa5, 0xb2, 0xf8, 0x71, 0x5e, 0xd2, 0xb5, 0x80, 0x6f, 0xe0, 0x28, 0x8f,
	0x0a, 0x6b, 0xb9, 0xd1, 0x0d, 0xeb, 0x73, 0xb6, 0xf5, 0x5f, 0xbb, 0xf8, 0x11, 0x1e, 0xf9, 0xfe,
	0x66, 0x67, 0x7a, 0x8a, 0x48, 0x0e, 0x7f, 0x52, 0x6c, 0xf8, 0xd4, 0xd6, 0xcd, 0x39, 0xf5, 0x46,
	0x12, 0x72, 0x6a, 0xcf, 0xf2, 0xa2, 0xc2, 0x7a, 0xf0, 0x78, 0xcb, 0x81, 0xf6, 0x99, 0x68, 0x7f,
	0xdb, 0x67, 0xa2, 0x2c, 0x6b, 0x6f, 0xc8, 0xbd, 0x65, 0x6f, 0xd8, 0x76, 0xac, 0xbd, 0x61, 0xd3,
	0x80, 0xa2, 0x32, 0xae, 0x9b, 0x5f, 0x93, 0xd7, 0xff, 0x03, 0x00, 0x00, 0xff, 0xff, 0xfd, 0x07,
	0x98, 0xb6, 0x61, 0x06, 0x00, 0x00,
}
