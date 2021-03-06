// Code generated by protoc-gen-micro. DO NOT EDIT.
// source: names.proto

package names

import (
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	math "math"
)

import (
	context "context"
	client "github.com/micro/go-micro/v2/client"
	server "github.com/micro/go-micro/v2/server"
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

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ client.Option
var _ server.Option

// Client API for Names service

type NamesService interface {
	GetNames(ctx context.Context, in *Request, opts ...client.CallOption) (*Response, error)
}

type namesService struct {
	c    client.Client
	name string
}

func NewNamesService(name string, c client.Client) NamesService {
	return &namesService{
		c:    c,
		name: name,
	}
}

func (c *namesService) GetNames(ctx context.Context, in *Request, opts ...client.CallOption) (*Response, error) {
	req := c.c.NewRequest(c.name, "Names.GetNames", in)
	out := new(Response)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for Names service

type NamesHandler interface {
	GetNames(context.Context, *Request, *Response) error
}

func RegisterNamesHandler(s server.Server, hdlr NamesHandler, opts ...server.HandlerOption) error {
	type names interface {
		GetNames(ctx context.Context, in *Request, out *Response) error
	}
	type Names struct {
		names
	}
	h := &namesHandler{hdlr}
	return s.Handle(s.NewHandler(&Names{h}, opts...))
}

type namesHandler struct {
	NamesHandler
}

func (h *namesHandler) GetNames(ctx context.Context, in *Request, out *Response) error {
	return h.NamesHandler.GetNames(ctx, in, out)
}
