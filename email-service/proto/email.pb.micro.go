// Code generated by protoc-gen-micro. DO NOT EDIT.
// source: email.proto

package email

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

// Client API for Email service

type EmailService interface {
	GetEmail(ctx context.Context, in *Request, opts ...client.CallOption) (*Response, error)
}

type emailService struct {
	c    client.Client
	name string
}

func NewEmailService(name string, c client.Client) EmailService {
	return &emailService{
		c:    c,
		name: name,
	}
}

func (c *emailService) GetEmail(ctx context.Context, in *Request, opts ...client.CallOption) (*Response, error) {
	req := c.c.NewRequest(c.name, "Email.GetEmail", in)
	out := new(Response)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for Email service

type EmailHandler interface {
	GetEmail(context.Context, *Request, *Response) error
}

func RegisterEmailHandler(s server.Server, hdlr EmailHandler, opts ...server.HandlerOption) error {
	type email interface {
		GetEmail(ctx context.Context, in *Request, out *Response) error
	}
	type Email struct {
		email
	}
	h := &emailHandler{hdlr}
	return s.Handle(s.NewHandler(&Email{h}, opts...))
}

type emailHandler struct {
	EmailHandler
}

func (h *emailHandler) GetEmail(ctx context.Context, in *Request, out *Response) error {
	return h.EmailHandler.GetEmail(ctx, in, out)
}
