// Scry Info.  All rights reserved.
// license that can be found in the license file.

// Code generated by protoc-gen-go. DO NOT EDIT.
// source: interface-service.proto

package scryinfo

import (
    context "context"
    fmt "fmt"
    proto "github.com/golang/protobuf/proto"
    grpc "google.golang.org/grpc"
    math "math"
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

type Status int32

const (
    Status_OK    Status = 0
    Status_ERROR Status = 1
)

var Status_name = map[int32]string{
    0: "OK",
    1: "ERROR",
}

var Status_value = map[string]int32{
    "OK":    0,
    "ERROR": 1,
}

func (x Status) String() string {
    return proto.EnumName(Status_name, int32(x))
}

func (Status) EnumDescriptor() ([]byte, []int) {
    return fileDescriptor_bcbe4554547bad70, []int{0}
}

type ImportParameter struct {
    ContentPassword      string   `protobuf:"bytes,1,opt,name=content_password,json=contentPassword,proto3" json:"content_password,omitempty"`
    ImportPsd            string   `protobuf:"bytes,2,opt,name=import_psd,json=importPsd,proto3" json:"import_psd,omitempty"`
    Content              []byte   `protobuf:"bytes,3,opt,name=content,proto3" json:"content,omitempty"`
    XXX_NoUnkeyedLiteral struct{} `json:"-"`
    XXX_unrecognized     []byte   `json:"-"`
    XXX_sizecache        int32    `json:"-"`
}

func (m *ImportParameter) Reset()         { *m = ImportParameter{} }
func (m *ImportParameter) String() string { return proto.CompactTextString(m) }
func (*ImportParameter) ProtoMessage()    {}
func (*ImportParameter) Descriptor() ([]byte, []int) {
    return fileDescriptor_bcbe4554547bad70, []int{0}
}

func (m *ImportParameter) XXX_Unmarshal(b []byte) error {
    return xxx_messageInfo_ImportParameter.Unmarshal(m, b)
}
func (m *ImportParameter) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
    return xxx_messageInfo_ImportParameter.Marshal(b, m, deterministic)
}
func (m *ImportParameter) XXX_Merge(src proto.Message) {
    xxx_messageInfo_ImportParameter.Merge(m, src)
}
func (m *ImportParameter) XXX_Size() int {
    return xxx_messageInfo_ImportParameter.Size(m)
}
func (m *ImportParameter) XXX_DiscardUnknown() {
    xxx_messageInfo_ImportParameter.DiscardUnknown(m)
}

var xxx_messageInfo_ImportParameter proto.InternalMessageInfo

func (m *ImportParameter) GetContentPassword() string {
    if m != nil {
        return m.ContentPassword
    }
    return ""
}

func (m *ImportParameter) GetImportPsd() string {
    if m != nil {
        return m.ImportPsd
    }
    return ""
}

func (m *ImportParameter) GetContent() []byte {
    if m != nil {
        return m.Content
    }
    return nil
}

type AddressParameter struct {
    Password             string   `protobuf:"bytes,1,opt,name=password,proto3" json:"password,omitempty"`
    Address              string   `protobuf:"bytes,2,opt,name=address,proto3" json:"address,omitempty"`
    XXX_NoUnkeyedLiteral struct{} `json:"-"`
    XXX_unrecognized     []byte   `json:"-"`
    XXX_sizecache        int32    `json:"-"`
}

func (m *AddressParameter) Reset()         { *m = AddressParameter{} }
func (m *AddressParameter) String() string { return proto.CompactTextString(m) }
func (*AddressParameter) ProtoMessage()    {}
func (*AddressParameter) Descriptor() ([]byte, []int) {
    return fileDescriptor_bcbe4554547bad70, []int{1}
}

func (m *AddressParameter) XXX_Unmarshal(b []byte) error {
    return xxx_messageInfo_AddressParameter.Unmarshal(m, b)
}
func (m *AddressParameter) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
    return xxx_messageInfo_AddressParameter.Marshal(b, m, deterministic)
}
func (m *AddressParameter) XXX_Merge(src proto.Message) {
    xxx_messageInfo_AddressParameter.Merge(m, src)
}
func (m *AddressParameter) XXX_Size() int {
    return xxx_messageInfo_AddressParameter.Size(m)
}
func (m *AddressParameter) XXX_DiscardUnknown() {
    xxx_messageInfo_AddressParameter.DiscardUnknown(m)
}

var xxx_messageInfo_AddressParameter proto.InternalMessageInfo

func (m *AddressParameter) GetPassword() string {
    if m != nil {
        return m.Password
    }
    return ""
}

func (m *AddressParameter) GetAddress() string {
    if m != nil {
        return m.Address
    }
    return ""
}

type AddressInfo struct {
    Status               Status   `protobuf:"varint,1,opt,name=status,proto3,enum=scryinfo.Status" json:"status,omitempty"`
    Address              string   `protobuf:"bytes,2,opt,name=address,proto3" json:"address,omitempty"`
    Msg                  string   `protobuf:"bytes,3,opt,name=msg,proto3" json:"msg,omitempty"`
    XXX_NoUnkeyedLiteral struct{} `json:"-"`
    XXX_unrecognized     []byte   `json:"-"`
    XXX_sizecache        int32    `json:"-"`
}

func (m *AddressInfo) Reset()         { *m = AddressInfo{} }
func (m *AddressInfo) String() string { return proto.CompactTextString(m) }
func (*AddressInfo) ProtoMessage()    {}
func (*AddressInfo) Descriptor() ([]byte, []int) {
    return fileDescriptor_bcbe4554547bad70, []int{2}
}

func (m *AddressInfo) XXX_Unmarshal(b []byte) error {
    return xxx_messageInfo_AddressInfo.Unmarshal(m, b)
}
func (m *AddressInfo) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
    return xxx_messageInfo_AddressInfo.Marshal(b, m, deterministic)
}
func (m *AddressInfo) XXX_Merge(src proto.Message) {
    xxx_messageInfo_AddressInfo.Merge(m, src)
}
func (m *AddressInfo) XXX_Size() int {
    return xxx_messageInfo_AddressInfo.Size(m)
}
func (m *AddressInfo) XXX_DiscardUnknown() {
    xxx_messageInfo_AddressInfo.DiscardUnknown(m)
}

var xxx_messageInfo_AddressInfo proto.InternalMessageInfo

func (m *AddressInfo) GetStatus() Status {
    if m != nil {
        return m.Status
    }
    return Status_OK
}

func (m *AddressInfo) GetAddress() string {
    if m != nil {
        return m.Address
    }
    return ""
}

func (m *AddressInfo) GetMsg() string {
    if m != nil {
        return m.Msg
    }
    return ""
}

type CipherParameter struct {
    Password             string   `protobuf:"bytes,1,opt,name=password,proto3" json:"password,omitempty"`
    Address              string   `protobuf:"bytes,2,opt,name=address,proto3" json:"address,omitempty"`
    Message              []byte   `protobuf:"bytes,3,opt,name=message,proto3" json:"message,omitempty"`
    XXX_NoUnkeyedLiteral struct{} `json:"-"`
    XXX_unrecognized     []byte   `json:"-"`
    XXX_sizecache        int32    `json:"-"`
}

func (m *CipherParameter) Reset()         { *m = CipherParameter{} }
func (m *CipherParameter) String() string { return proto.CompactTextString(m) }
func (*CipherParameter) ProtoMessage()    {}
func (*CipherParameter) Descriptor() ([]byte, []int) {
    return fileDescriptor_bcbe4554547bad70, []int{3}
}

func (m *CipherParameter) XXX_Unmarshal(b []byte) error {
    return xxx_messageInfo_CipherParameter.Unmarshal(m, b)
}
func (m *CipherParameter) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
    return xxx_messageInfo_CipherParameter.Marshal(b, m, deterministic)
}
func (m *CipherParameter) XXX_Merge(src proto.Message) {
    xxx_messageInfo_CipherParameter.Merge(m, src)
}
func (m *CipherParameter) XXX_Size() int {
    return xxx_messageInfo_CipherParameter.Size(m)
}
func (m *CipherParameter) XXX_DiscardUnknown() {
    xxx_messageInfo_CipherParameter.DiscardUnknown(m)
}

var xxx_messageInfo_CipherParameter proto.InternalMessageInfo

func (m *CipherParameter) GetPassword() string {
    if m != nil {
        return m.Password
    }
    return ""
}

func (m *CipherParameter) GetAddress() string {
    if m != nil {
        return m.Address
    }
    return ""
}

func (m *CipherParameter) GetMessage() []byte {
    if m != nil {
        return m.Message
    }
    return nil
}

type CipherText struct {
    Status               Status   `protobuf:"varint,1,opt,name=status,proto3,enum=scryinfo.Status" json:"status,omitempty"`
    Data                 []byte   `protobuf:"bytes,2,opt,name=data,proto3" json:"data,omitempty"`
    Msg                  string   `protobuf:"bytes,3,opt,name=msg,proto3" json:"msg,omitempty"`
    XXX_NoUnkeyedLiteral struct{} `json:"-"`
    XXX_unrecognized     []byte   `json:"-"`
    XXX_sizecache        int32    `json:"-"`
}

func (m *CipherText) Reset()         { *m = CipherText{} }
func (m *CipherText) String() string { return proto.CompactTextString(m) }
func (*CipherText) ProtoMessage()    {}
func (*CipherText) Descriptor() ([]byte, []int) {
    return fileDescriptor_bcbe4554547bad70, []int{4}
}

func (m *CipherText) XXX_Unmarshal(b []byte) error {
    return xxx_messageInfo_CipherText.Unmarshal(m, b)
}
func (m *CipherText) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
    return xxx_messageInfo_CipherText.Marshal(b, m, deterministic)
}
func (m *CipherText) XXX_Merge(src proto.Message) {
    xxx_messageInfo_CipherText.Merge(m, src)
}
func (m *CipherText) XXX_Size() int {
    return xxx_messageInfo_CipherText.Size(m)
}
func (m *CipherText) XXX_DiscardUnknown() {
    xxx_messageInfo_CipherText.DiscardUnknown(m)
}

var xxx_messageInfo_CipherText proto.InternalMessageInfo

func (m *CipherText) GetStatus() Status {
    if m != nil {
        return m.Status
    }
    return Status_OK
}

func (m *CipherText) GetData() []byte {
    if m != nil {
        return m.Data
    }
    return nil
}

func (m *CipherText) GetMsg() string {
    if m != nil {
        return m.Msg
    }
    return ""
}

func init() {
    proto.RegisterEnum("scryinfo.Status", Status_name, Status_value)
    proto.RegisterType((*ImportParameter)(nil), "scryinfo.ImportParameter")
    proto.RegisterType((*AddressParameter)(nil), "scryinfo.AddressParameter")
    proto.RegisterType((*AddressInfo)(nil), "scryinfo.AddressInfo")
    proto.RegisterType((*CipherParameter)(nil), "scryinfo.CipherParameter")
    proto.RegisterType((*CipherText)(nil), "scryinfo.CipherText")
}

func init() { proto.RegisterFile("interface-service.proto", fileDescriptor_bcbe4554547bad70) }

var fileDescriptor_bcbe4554547bad70 = []byte{
    // 400 bytes of a gzipped FileDescriptorProto
    0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xa4, 0x53, 0xcf, 0xab, 0xd3, 0x40,
    0x10, 0x7e, 0x79, 0xd5, 0xf8, 0x32, 0x3e, 0x9b, 0xb0, 0x58, 0x88, 0x11, 0xa1, 0xe4, 0x54, 0x05,
    0x7b, 0xa8, 0x77, 0x41, 0xda, 0xaa, 0xa5, 0x87, 0x96, 0xad, 0x78, 0x12, 0xca, 0x9a, 0x4c, 0x6b,
    0x90, 0xec, 0x86, 0xdd, 0x8d, 0x9a, 0xbf, 0xc9, 0x7f, 0x52, 0xb2, 0xd9, 0x58, 0x9a, 0xd2, 0xc3,
    0x6b, 0x6f, 0x3b, 0xbf, 0xbe, 0x6f, 0xe7, 0x9b, 0x19, 0x18, 0xb0, 0x24, 0x11, 0x25, 0xd7, 0x6f,
    0x15, 0xca, 0x5f, 0x59, 0x82, 0xe3, 0x42, 0x0a, 0x2d, 0xc8, 0x9d, 0x4a, 0x64, 0x95, 0xf1, 0x9d,
    0x88, 0x4b, 0xf0, 0x17, 0x79, 0x21, 0xa4, 0x5e, 0x33, 0xc9, 0x72, 0xd4, 0x28, 0xc9, 0x6b, 0x08,
    0x12, 0xc1, 0x35, 0x72, 0xbd, 0x2d, 0x98, 0x52, 0xbf, 0x85, 0x4c, 0x43, 0x67, 0xe8, 0x8c, 0x3c,
    0xea, 0x5b, 0xff, 0xda, 0xba, 0xc9, 0x2b, 0x80, 0xcc, 0x54, 0x6f, 0x0b, 0x95, 0x86, 0xb7, 0x26,
    0xc9, 0x6b, 0x3c, 0x6b, 0x95, 0x92, 0x10, 0x9e, 0xd8, 0x8a, 0xb0, 0x37, 0x74, 0x46, 0xf7, 0xb4,
    0x35, 0xe3, 0xcf, 0x10, 0x7c, 0x48, 0x53, 0x89, 0x4a, 0x1d, 0x78, 0x23, 0xb8, 0xeb, 0xf0, 0xfd,
    0xb7, 0x6b, 0x24, 0xd6, 0xe4, 0x5b, 0x96, 0xd6, 0x8c, 0x13, 0x78, 0x6a, 0x91, 0x16, 0x7c, 0x27,
    0xc8, 0x08, 0x5c, 0xa5, 0x99, 0x2e, 0x95, 0x81, 0xe8, 0x4f, 0x82, 0x71, 0xdb, 0xea, 0x78, 0x63,
    0xfc, 0xd4, 0xc6, 0xcf, 0x43, 0x92, 0x00, 0x7a, 0xb9, 0xda, 0x9b, 0x2f, 0x7b, 0xb4, 0x7e, 0xc6,
    0x0c, 0xfc, 0x69, 0x56, 0xfc, 0x40, 0x79, 0xe5, 0x6f, 0xeb, 0x48, 0x8e, 0x4a, 0xb1, 0x3d, 0xb6,
    0x8a, 0x58, 0x33, 0xfe, 0x06, 0xd0, 0x50, 0x7c, 0xc1, 0x3f, 0xfa, 0x01, 0x6d, 0x10, 0x78, 0x94,
    0x32, 0xcd, 0x0c, 0xd1, 0x3d, 0x35, 0xef, 0xd3, 0x06, 0xde, 0xbc, 0x04, 0xb7, 0xa9, 0x23, 0x2e,
    0xdc, 0xae, 0x96, 0xc1, 0x0d, 0xf1, 0xe0, 0xf1, 0x9c, 0xd2, 0x15, 0x0d, 0x9c, 0xc9, 0xdf, 0x1e,
    0xc0, 0x12, 0xab, 0x4d, 0xb3, 0x22, 0xe4, 0x23, 0xf8, 0x9f, 0x90, 0xa3, 0x64, 0x1a, 0xad, 0xb2,
    0x24, 0x3a, 0xd0, 0x77, 0xc7, 0x16, 0x0d, 0x4e, 0x62, 0xf5, 0x20, 0xe2, 0x1b, 0x32, 0x83, 0x67,
    0x5f, 0x51, 0x66, 0xbb, 0xea, 0x2a, 0x94, 0x29, 0xf4, 0xa7, 0xcd, 0xd2, 0xcc, 0x79, 0x22, 0xab,
    0x42, 0x93, 0x17, 0x87, 0xd4, 0xce, 0x50, 0xa2, 0xe7, 0xdd, 0x50, 0x2d, 0xe6, 0x11, 0xc8, 0x0c,
    0x2f, 0x06, 0x79, 0x0f, 0xde, 0x26, 0xdb, 0x73, 0xa6, 0x4b, 0x89, 0x97, 0xd4, 0xcf, 0xc1, 0xb7,
    0xc7, 0xf2, 0x13, 0x2b, 0xa5, 0xc5, 0x31, 0x4a, 0xe7, 0x0a, 0xcf, 0x0a, 0xf2, 0xdd, 0x35, 0x27,
    0xfc, 0xee, 0x5f, 0x00, 0x00, 0x00, 0xff, 0xff, 0x21, 0x45, 0x03, 0xbc, 0xdb, 0x03, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// KeyServiceClient is the client API for KeyService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type KeyServiceClient interface {
    //Generate address
    GenerateAddress(ctx context.Context, in *AddressParameter, opts ...grpc.CallOption) (*AddressInfo, error)
    //Proof address
    VerifyAddress(ctx context.Context, in *AddressParameter, opts ...grpc.CallOption) (*AddressInfo, error)
    //Content encryption
    ContentEncrypt(ctx context.Context, in *CipherParameter, opts ...grpc.CallOption) (*CipherText, error)
    //Content decryption
    ContentDecrypt(ctx context.Context, in *CipherParameter, opts ...grpc.CallOption) (*CipherText, error)
    //info signature
    Signature(ctx context.Context, in *CipherParameter, opts ...grpc.CallOption) (*CipherText, error)
    //Input keystore file content
    ImportKeystore(ctx context.Context, in *ImportParameter, opts ...grpc.CallOption) (*AddressInfo, error)
}

type keyServiceClient struct {
    cc *grpc.ClientConn
}

func NewKeyServiceClient(cc *grpc.ClientConn) KeyServiceClient {
    return &keyServiceClient{cc}
}

func (c *keyServiceClient) GenerateAddress(ctx context.Context, in *AddressParameter, opts ...grpc.CallOption) (*AddressInfo, error) {
    out := new(AddressInfo)
    err := c.cc.Invoke(ctx, "/scryinfo.KeyService/GenerateAddress", in, out, opts...)
    if err != nil {
        return nil, err
    }
    return out, nil
}

func (c *keyServiceClient) VerifyAddress(ctx context.Context, in *AddressParameter, opts ...grpc.CallOption) (*AddressInfo, error) {
    out := new(AddressInfo)
    err := c.cc.Invoke(ctx, "/scryinfo.KeyService/VerifyAddress", in, out, opts...)
    if err != nil {
        return nil, err
    }
    return out, nil
}

func (c *keyServiceClient) ContentEncrypt(ctx context.Context, in *CipherParameter, opts ...grpc.CallOption) (*CipherText, error) {
    out := new(CipherText)
    err := c.cc.Invoke(ctx, "/scryinfo.KeyService/ContentEncrypt", in, out, opts...)
    if err != nil {
        return nil, err
    }
    return out, nil
}

func (c *keyServiceClient) ContentDecrypt(ctx context.Context, in *CipherParameter, opts ...grpc.CallOption) (*CipherText, error) {
    out := new(CipherText)
    err := c.cc.Invoke(ctx, "/scryinfo.KeyService/ContentDecrypt", in, out, opts...)
    if err != nil {
        return nil, err
    }
    return out, nil
}

func (c *keyServiceClient) Signature(ctx context.Context, in *CipherParameter, opts ...grpc.CallOption) (*CipherText, error) {
    out := new(CipherText)
    err := c.cc.Invoke(ctx, "/scryinfo.KeyService/Signature", in, out, opts...)
    if err != nil {
        return nil, err
    }
    return out, nil
}

func (c *keyServiceClient) ImportKeystore(ctx context.Context, in *ImportParameter, opts ...grpc.CallOption) (*AddressInfo, error) {
    out := new(AddressInfo)
    err := c.cc.Invoke(ctx, "/scryinfo.KeyService/import_keystore", in, out, opts...)
    if err != nil {
        return nil, err
    }
    return out, nil
}

// KeyServiceServer is the server API for KeyService service.
type KeyServiceServer interface {
    //Generate address
    GenerateAddress(context.Context, *AddressParameter) (*AddressInfo, error)
    //Proof address
    VerifyAddress(context.Context, *AddressParameter) (*AddressInfo, error)
    //Content encryption
    ContentEncrypt(context.Context, *CipherParameter) (*CipherText, error)
    //Content decryption
    ContentDecrypt(context.Context, *CipherParameter) (*CipherText, error)
    //Info signature
    Signature(context.Context, *CipherParameter) (*CipherText, error)
    //Input keystore file content
    ImportKeystore(context.Context, *ImportParameter) (*AddressInfo, error)
}

func RegisterKeyServiceServer(s *grpc.Server, srv KeyServiceServer) {
    s.RegisterService(&_KeyService_serviceDesc, srv)
}

func _KeyService_GenerateAddress_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
    in := new(AddressParameter)
    if err := dec(in); err != nil {
        return nil, err
    }
    if interceptor == nil {
        return srv.(KeyServiceServer).GenerateAddress(ctx, in)
    }
    info := &grpc.UnaryServerInfo{
        Server:     srv,
        FullMethod: "/scryinfo.KeyService/GenerateAddress",
    }
    handler := func(ctx context.Context, req interface{}) (interface{}, error) {
        return srv.(KeyServiceServer).GenerateAddress(ctx, req.(*AddressParameter))
    }
    return interceptor(ctx, in, info, handler)
}

func _KeyService_VerifyAddress_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
    in := new(AddressParameter)
    if err := dec(in); err != nil {
        return nil, err
    }
    if interceptor == nil {
        return srv.(KeyServiceServer).VerifyAddress(ctx, in)
    }
    info := &grpc.UnaryServerInfo{
        Server:     srv,
        FullMethod: "/scryinfo.KeyService/VerifyAddress",
    }
    handler := func(ctx context.Context, req interface{}) (interface{}, error) {
        return srv.(KeyServiceServer).VerifyAddress(ctx, req.(*AddressParameter))
    }
    return interceptor(ctx, in, info, handler)
}

func _KeyService_ContentEncrypt_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
    in := new(CipherParameter)
    if err := dec(in); err != nil {
        return nil, err
    }
    if interceptor == nil {
        return srv.(KeyServiceServer).ContentEncrypt(ctx, in)
    }
    info := &grpc.UnaryServerInfo{
        Server:     srv,
        FullMethod: "/scryinfo.KeyService/ContentEncrypt",
    }
    handler := func(ctx context.Context, req interface{}) (interface{}, error) {
        return srv.(KeyServiceServer).ContentEncrypt(ctx, req.(*CipherParameter))
    }
    return interceptor(ctx, in, info, handler)
}

func _KeyService_ContentDecrypt_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
    in := new(CipherParameter)
    if err := dec(in); err != nil {
        return nil, err
    }
    if interceptor == nil {
        return srv.(KeyServiceServer).ContentDecrypt(ctx, in)
    }
    info := &grpc.UnaryServerInfo{
        Server:     srv,
        FullMethod: "/scryinfo.KeyService/ContentDecrypt",
    }
    handler := func(ctx context.Context, req interface{}) (interface{}, error) {
        return srv.(KeyServiceServer).ContentDecrypt(ctx, req.(*CipherParameter))
    }
    return interceptor(ctx, in, info, handler)
}

func _KeyService_Signature_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
    in := new(CipherParameter)
    if err := dec(in); err != nil {
        return nil, err
    }
    if interceptor == nil {
        return srv.(KeyServiceServer).Signature(ctx, in)
    }
    info := &grpc.UnaryServerInfo{
        Server:     srv,
        FullMethod: "/scryinfo.KeyService/Signature",
    }
    handler := func(ctx context.Context, req interface{}) (interface{}, error) {
        return srv.(KeyServiceServer).Signature(ctx, req.(*CipherParameter))
    }
    return interceptor(ctx, in, info, handler)
}

func _KeyService_ImportKeystore_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
    in := new(ImportParameter)
    if err := dec(in); err != nil {
        return nil, err
    }
    if interceptor == nil {
        return srv.(KeyServiceServer).ImportKeystore(ctx, in)
    }
    info := &grpc.UnaryServerInfo{
        Server:     srv,
        FullMethod: "/scryinfo.KeyService/ImportKeystore",
    }
    handler := func(ctx context.Context, req interface{}) (interface{}, error) {
        return srv.(KeyServiceServer).ImportKeystore(ctx, req.(*ImportParameter))
    }
    return interceptor(ctx, in, info, handler)
}

var _KeyService_serviceDesc = grpc.ServiceDesc{
    ServiceName: "scryinfo.KeyService",
    HandlerType: (*KeyServiceServer)(nil),
    Methods: []grpc.MethodDesc{
        {
            MethodName: "GenerateAddress",
            Handler:    _KeyService_GenerateAddress_Handler,
        },
        {
            MethodName: "VerifyAddress",
            Handler:    _KeyService_VerifyAddress_Handler,
        },
        {
            MethodName: "ContentEncrypt",
            Handler:    _KeyService_ContentEncrypt_Handler,
        },
        {
            MethodName: "ContentDecrypt",
            Handler:    _KeyService_ContentDecrypt_Handler,
        },
        {
            MethodName: "Signature",
            Handler:    _KeyService_Signature_Handler,
        },
        {
            MethodName: "import_keystore",
            Handler:    _KeyService_ImportKeystore_Handler,
        },
    },
    Streams:  []grpc.StreamDesc{},
    Metadata: "interface-service.proto",
}
