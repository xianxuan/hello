// Code generated by protoc-gen-go. DO NOT EDIT.
// source: hello/hello_world.proto

/*
Package hello_world is a generated protocol buffer package.

It is generated from these files:
	hello/hello_world.proto

It has these top-level messages:
	HelloRequest
	HelloResponse
*/
package hello_world

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

import (
	client "github.com/micro/go-micro/client"
	server "github.com/micro/go-micro/server"
	context "golang.org/x/net/context"
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

type HelloRequest struct {
	Name string `protobuf:"bytes,1,opt,name=name" json:"name,omitempty"`
}

func (m *HelloRequest) Reset()                    { *m = HelloRequest{} }
func (m *HelloRequest) String() string            { return proto.CompactTextString(m) }
func (*HelloRequest) ProtoMessage()               {}
func (*HelloRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

func (m *HelloRequest) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

type HelloResponse struct {
	Greeting string `protobuf:"bytes,2,opt,name=greeting" json:"greeting,omitempty"`
}

func (m *HelloResponse) Reset()                    { *m = HelloResponse{} }
func (m *HelloResponse) String() string            { return proto.CompactTextString(m) }
func (*HelloResponse) ProtoMessage()               {}
func (*HelloResponse) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

func (m *HelloResponse) GetGreeting() string {
	if m != nil {
		return m.Greeting
	}
	return ""
}

func init() {
	proto.RegisterType((*HelloRequest)(nil), "HelloRequest")
	proto.RegisterType((*HelloResponse)(nil), "HelloResponse")
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ client.Option
var _ server.Option

// Client API for HelloWorld service

type HelloWorldClient interface {
	Hello(ctx context.Context, in *HelloRequest, opts ...client.CallOption) (*HelloResponse, error)
}

type helloWorldClient struct {
	c           client.Client
	serviceName string
}

func NewHelloWorldClient(serviceName string, c client.Client) HelloWorldClient {
	if c == nil {
		c = client.NewClient()
	}
	if len(serviceName) == 0 {
		serviceName = "helloworld"
	}
	return &helloWorldClient{
		c:           c,
		serviceName: serviceName,
	}
}

func (c *helloWorldClient) Hello(ctx context.Context, in *HelloRequest, opts ...client.CallOption) (*HelloResponse, error) {
	req := c.c.NewRequest(c.serviceName, "HelloWorld.Hello", in)
	out := new(HelloResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for HelloWorld service

type HelloWorldHandler interface {
	Hello(context.Context, *HelloRequest, *HelloResponse) error
}

func RegisterHelloWorldHandler(s server.Server, hdlr HelloWorldHandler, opts ...server.HandlerOption) {
	s.Handle(s.NewHandler(&HelloWorld{hdlr}, opts...))
}

type HelloWorld struct {
	HelloWorldHandler
}

func (h *HelloWorld) Hello(ctx context.Context, in *HelloRequest, out *HelloResponse) error {
	return h.HelloWorldHandler.Hello(ctx, in, out)
}

func init() { proto.RegisterFile("hello/hello_world.proto", fileDescriptor0) }

var fileDescriptor0 = []byte{
	// 139 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x12, 0xcf, 0x48, 0xcd, 0xc9,
	0xc9, 0xd7, 0x07, 0x93, 0xf1, 0xe5, 0xf9, 0x45, 0x39, 0x29, 0x7a, 0x05, 0x45, 0xf9, 0x25, 0xf9,
	0x4a, 0x4a, 0x5c, 0x3c, 0x1e, 0x20, 0xc1, 0xa0, 0xd4, 0xc2, 0xd2, 0xd4, 0xe2, 0x12, 0x21, 0x21,
	0x2e, 0x96, 0xbc, 0xc4, 0xdc, 0x54, 0x09, 0x46, 0x05, 0x46, 0x0d, 0xce, 0x20, 0x30, 0x5b, 0x49,
	0x9b, 0x8b, 0x17, 0xaa, 0xa6, 0xb8, 0x20, 0x3f, 0xaf, 0x38, 0x55, 0x48, 0x8a, 0x8b, 0x23, 0xbd,
	0x28, 0x35, 0xb5, 0x24, 0x33, 0x2f, 0x5d, 0x82, 0x09, 0xac, 0x10, 0xce, 0x37, 0x32, 0xe3, 0xe2,
	0x02, 0x2b, 0x0e, 0x07, 0x59, 0x22, 0xa4, 0xc1, 0xc5, 0x0a, 0xe6, 0x09, 0xf1, 0xea, 0x21, 0x5b,
	0x23, 0xc5, 0xa7, 0x87, 0x62, 0xa2, 0x12, 0x43, 0x12, 0x1b, 0xd8, 0x3d, 0xc6, 0x80, 0x00, 0x00,
	0x00, 0xff, 0xff, 0xcf, 0x90, 0x6d, 0x3b, 0xaa, 0x00, 0x00, 0x00,
}
