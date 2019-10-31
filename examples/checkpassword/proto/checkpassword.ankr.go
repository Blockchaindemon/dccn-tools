// Code generated by protoc-gen-ankr. DO NOT EDIT.
// source: proto/checkpassword.proto

package checkpassword

import (
	context "context"
	json "encoding/json"
	fmt "fmt"
	logger "github.com/Ankr-network/dccn-tools/logger"
	zap "github.com/Ankr-network/dccn-tools/zap"
	proto "github.com/golang/protobuf/proto"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	metadata "google.golang.org/grpc/metadata"
	status "google.golang.org/grpc/status"
	math "math"
	public "public"
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

type Request struct {
	Username             string   `protobuf:"bytes,1,opt,name=username,proto3" json:"username,omitempty"`
	Password             string   `protobuf:"bytes,2,opt,name=password,proto3" json:"password,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Request) Reset()         { *m = Request{} }
func (m *Request) String() string { return proto.CompactTextString(m) }
func (*Request) ProtoMessage()    {}
func (*Request) Descriptor() ([]byte, []int) {
	return fileDescriptor_16b233f109744c3c, []int{0}
}

func (m *Request) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Request.Unmarshal(m, b)
}
func (m *Request) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Request.Marshal(b, m, deterministic)
}
func (m *Request) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Request.Merge(m, src)
}
func (m *Request) XXX_Size() int {
	return xxx_messageInfo_Request.Size(m)
}
func (m *Request) XXX_DiscardUnknown() {
	xxx_messageInfo_Request.DiscardUnknown(m)
}

var xxx_messageInfo_Request proto.InternalMessageInfo

func (m *Request) GetUsername() string {
	if m != nil {
		return m.Username
	}
	return ""
}

func (m *Request) GetPassword() string {
	if m != nil {
		return m.Password
	}
	return ""
}

type Response struct {
	Rsp                  *public.RspCodeMsg `protobuf:"bytes,1,opt,name=rsp,proto3" json:"rsp,omitempty"`
	XXX_NoUnkeyedLiteral struct{}           `json:"-"`
	XXX_unrecognized     []byte             `json:"-"`
	XXX_sizecache        int32              `json:"-"`
}

func (m *Response) Reset()         { *m = Response{} }
func (m *Response) String() string { return proto.CompactTextString(m) }
func (*Response) ProtoMessage()    {}
func (*Response) Descriptor() ([]byte, []int) {
	return fileDescriptor_16b233f109744c3c, []int{1}
}

func (m *Response) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Response.Unmarshal(m, b)
}
func (m *Response) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Response.Marshal(b, m, deterministic)
}
func (m *Response) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Response.Merge(m, src)
}
func (m *Response) XXX_Size() int {
	return xxx_messageInfo_Response.Size(m)
}
func (m *Response) XXX_DiscardUnknown() {
	xxx_messageInfo_Response.DiscardUnknown(m)
}

var xxx_messageInfo_Response proto.InternalMessageInfo

func (m *Response) GetRsp() *public.RspCodeMsg {
	if m != nil {
		return m.Rsp
	}
	return nil
}

func init() {
	proto.RegisterType((*Request)(nil), "checkpassword.request")
	proto.RegisterType((*Response)(nil), "checkpassword.response")
}

func init() { proto.RegisterFile("proto/checkpassword.proto", fileDescriptor_16b233f109744c3c) }

var fileDescriptor_16b233f109744c3c = []byte{
	// 195 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x92, 0x2c, 0x28, 0xca, 0x2f,
	0xc9, 0xd7, 0x4f, 0xce, 0x48, 0x4d, 0xce, 0x2e, 0x48, 0x2c, 0x2e, 0x2e, 0xcf, 0x2f, 0x4a, 0xd1,
	0x03, 0x8b, 0x09, 0xf1, 0xa2, 0x08, 0x4a, 0x89, 0x14, 0x94, 0x26, 0xe5, 0x64, 0x26, 0xeb, 0xe7,
	0xa6, 0x16, 0x17, 0x27, 0xa6, 0xa7, 0x42, 0x14, 0x29, 0x39, 0x72, 0xb1, 0x17, 0xa5, 0x16, 0x96,
	0xa6, 0x16, 0x97, 0x08, 0x49, 0x71, 0x71, 0x94, 0x16, 0xa7, 0x16, 0xe5, 0x25, 0xe6, 0xa6, 0x4a,
	0x30, 0x2a, 0x30, 0x6a, 0x70, 0x06, 0xc1, 0xf9, 0x20, 0x39, 0x98, 0x41, 0x12, 0x4c, 0x10, 0x39,
	0x18, 0x5f, 0xc9, 0x82, 0x8b, 0xa3, 0x28, 0xb5, 0xb8, 0x20, 0x3f, 0xaf, 0x38, 0x55, 0x48, 0x87,
	0x8b, 0xb9, 0xa8, 0xb8, 0x00, 0xac, 0x9d, 0xdb, 0x48, 0x4a, 0xaf, 0xa0, 0x34, 0x27, 0x33, 0x29,
	0x59, 0x0f, 0x66, 0x65, 0x50, 0x71, 0x81, 0x73, 0x7e, 0x4a, 0xaa, 0x6f, 0x71, 0x7a, 0x10, 0x48,
	0x99, 0x91, 0x1b, 0x17, 0x87, 0x33, 0xc8, 0x8d, 0x01, 0xe1, 0x2e, 0x42, 0x56, 0x5c, 0xac, 0x60,
	0xb6, 0x90, 0x98, 0x1e, 0xaa, 0x67, 0xa0, 0xce, 0x93, 0x12, 0xc7, 0x10, 0x87, 0xd8, 0xa9, 0xc4,
	0x90, 0xc4, 0x06, 0xf6, 0x8b, 0x31, 0x20, 0x00, 0x00, 0xff, 0xff, 0x5b, 0x51, 0xab, 0x75, 0x0d,
	0x01, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn
var l = logger.NewLogger()

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion5

// CheckPWDClient is the client API for CheckPWD service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type CheckPWDClient interface {
	Check(ctx context.Context, in *Request, opts ...grpc.CallOption) (*Response, error)
	Close() error
}

type checkPWDClient struct {
	cc *grpc.ClientConn
}

// origin client method
func NewCheckPWDClient(cc *grpc.ClientConn) CheckPWDClient {
	return &checkPWDClient{cc}
}

// new client method
func NewAnkrCheckPWDClient(addr string) (CheckPWDClient, error) {
	c, err := grpc.Dial(addr, grpc.WithInsecure())
	if err != nil {
		return nil, err
	}
	return &checkPWDClient{c}, nil
}

// new client close method
func (c *checkPWDClient) Close() error {
	return c.cc.Close()
}

func (c *checkPWDClient) Check(ctx context.Context, in *Request, opts ...grpc.CallOption) (*Response, error) {
	out := new(Response)
	err := c.cc.Invoke(ctx, "/checkpassword.CheckPWD/Check", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// CheckPWDServer is the server API for CheckPWD service.
type CheckPWDServer interface {
	Check(context.Context, *Request) (*Response, error)
}

// UnimplementedCheckPWDServer can be embedded to have forward compatible implementations.
type UnimplementedCheckPWDServer struct {
}

func (*UnimplementedCheckPWDServer) Check(ctx context.Context, req *Request) (*Response, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Check not implemented")
}

func RegisterCheckPWDServer(s *grpc.Server, srv CheckPWDServer) {
	s.RegisterService(&_CheckPWD_serviceDesc, srv)
}

func _CheckPWD_Check_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Request)
	if err := dec(in); err != nil {
		return nil, err
	}
	//if exist trace id then set new span id, else set the entire id values
	m := make(map[string]string)
	if md, ok := metadata.FromIncomingContext(ctx); ok {
		if vs := md.Get(logger.TraceID); len(vs) > 0 {
			m[logger.TraceID] = vs[len(vs)-1]
		} else {
			m[logger.TraceID] = l.Generate().String()
		}
		if vs := md.Get(logger.SpanID); len(vs) > 0 {
			m[logger.ParentSpanID] = vs[len(vs)-1]
		} else {
			m[logger.ParentSpanID] = l.Generate().String()
		}
	}
	m[logger.SpanID] = l.Generate().String()
	ctx = metadata.NewIncomingContext(ctx, metadata.New(m))
	if interceptor == nil {
		// output request
		reqBody, err := json.Marshal(&in)
		if err != nil {
			l.Error(ctx, err.Error(), zap.String("type", "request"))
		} else {
			l.Info(ctx, string(reqBody), zap.String("type", "request"))
		}
		rsp, err := srv.(CheckPWDServer).Check(ctx, in)
		// output response
		rspBody, err := json.Marshal(&rsp)
		if err != nil {
			l.Error(ctx, err.Error(), zap.String("type", "response"))
		} else {
			l.Info(ctx, string(rspBody), zap.String("type", "response"))
		}
		return rsp, err
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/checkpassword.CheckPWD/Check",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CheckPWDServer).Check(ctx, req.(*Request))
	}
	return interceptor(ctx, in, info, handler)
}

var _CheckPWD_serviceDesc = grpc.ServiceDesc{
	ServiceName: "checkpassword.CheckPWD",
	HandlerType: (*CheckPWDServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Check",
			Handler:    _CheckPWD_Check_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "proto/checkpassword.proto",
}