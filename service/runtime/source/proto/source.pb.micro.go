// Code generated by protoc-gen-micro. DO NOT EDIT.
// source: github.com/micro/micro/service/runtime/source/proto/source.proto

package go_micro_server

import (
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	math "math"
)

import (
	context "context"
	api "github.com/micro/go-micro/v3/api"
	client "github.com/micro/go-micro/v3/client"
	server "github.com/micro/go-micro/v3/server"
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
var _ api.Endpoint
var _ context.Context
var _ client.Option
var _ server.Option

// Api Endpoints for Source service

func NewSourceEndpoints() []*api.Endpoint {
	return []*api.Endpoint{}
}

// Client API for Source service

type SourceService interface {
	Upload(ctx context.Context, in *UploadRequest, opts ...client.CallOption) (Source_UploadService, error)
	Download(ctx context.Context, in *DownloadRequest, opts ...client.CallOption) (Source_DownloadService, error)
}

type sourceService struct {
	c    client.Client
	name string
}

func NewSourceService(name string, c client.Client) SourceService {
	return &sourceService{
		c:    c,
		name: name,
	}
}

func (c *sourceService) Upload(ctx context.Context, in *UploadRequest, opts ...client.CallOption) (Source_UploadService, error) {
	req := c.c.NewRequest(c.name, "Source.Upload", &UploadRequest{})
	stream, err := c.c.Stream(ctx, req, opts...)
	if err != nil {
		return nil, err
	}
	if err := stream.Send(in); err != nil {
		return nil, err
	}
	return &sourceServiceUpload{stream}, nil
}

type Source_UploadService interface {
	Context() context.Context
	SendMsg(interface{}) error
	RecvMsg(interface{}) error
	Close() error
	Recv() (*Data, error)
}

type sourceServiceUpload struct {
	stream client.Stream
}

func (x *sourceServiceUpload) Close() error {
	return x.stream.Close()
}

func (x *sourceServiceUpload) Context() context.Context {
	return x.stream.Context()
}

func (x *sourceServiceUpload) SendMsg(m interface{}) error {
	return x.stream.Send(m)
}

func (x *sourceServiceUpload) RecvMsg(m interface{}) error {
	return x.stream.Recv(m)
}

func (x *sourceServiceUpload) Recv() (*Data, error) {
	m := new(Data)
	err := x.stream.Recv(m)
	if err != nil {
		return nil, err
	}
	return m, nil
}

func (c *sourceService) Download(ctx context.Context, in *DownloadRequest, opts ...client.CallOption) (Source_DownloadService, error) {
	req := c.c.NewRequest(c.name, "Source.Download", &DownloadRequest{})
	stream, err := c.c.Stream(ctx, req, opts...)
	if err != nil {
		return nil, err
	}
	if err := stream.Send(in); err != nil {
		return nil, err
	}
	return &sourceServiceDownload{stream}, nil
}

type Source_DownloadService interface {
	Context() context.Context
	SendMsg(interface{}) error
	RecvMsg(interface{}) error
	Close() error
	Recv() (*Data, error)
}

type sourceServiceDownload struct {
	stream client.Stream
}

func (x *sourceServiceDownload) Close() error {
	return x.stream.Close()
}

func (x *sourceServiceDownload) Context() context.Context {
	return x.stream.Context()
}

func (x *sourceServiceDownload) SendMsg(m interface{}) error {
	return x.stream.Send(m)
}

func (x *sourceServiceDownload) RecvMsg(m interface{}) error {
	return x.stream.Recv(m)
}

func (x *sourceServiceDownload) Recv() (*Data, error) {
	m := new(Data)
	err := x.stream.Recv(m)
	if err != nil {
		return nil, err
	}
	return m, nil
}

// Server API for Source service

type SourceHandler interface {
	Upload(context.Context, *UploadRequest, Source_UploadStream) error
	Download(context.Context, *DownloadRequest, Source_DownloadStream) error
}

func RegisterSourceHandler(s server.Server, hdlr SourceHandler, opts ...server.HandlerOption) error {
	type source interface {
		Upload(ctx context.Context, stream server.Stream) error
		Download(ctx context.Context, stream server.Stream) error
	}
	type Source struct {
		source
	}
	h := &sourceHandler{hdlr}
	return s.Handle(s.NewHandler(&Source{h}, opts...))
}

type sourceHandler struct {
	SourceHandler
}

func (h *sourceHandler) Upload(ctx context.Context, stream server.Stream) error {
	m := new(UploadRequest)
	if err := stream.Recv(m); err != nil {
		return err
	}
	return h.SourceHandler.Upload(ctx, m, &sourceUploadStream{stream})
}

type Source_UploadStream interface {
	Context() context.Context
	SendMsg(interface{}) error
	RecvMsg(interface{}) error
	Close() error
	Send(*Data) error
}

type sourceUploadStream struct {
	stream server.Stream
}

func (x *sourceUploadStream) Close() error {
	return x.stream.Close()
}

func (x *sourceUploadStream) Context() context.Context {
	return x.stream.Context()
}

func (x *sourceUploadStream) SendMsg(m interface{}) error {
	return x.stream.Send(m)
}

func (x *sourceUploadStream) RecvMsg(m interface{}) error {
	return x.stream.Recv(m)
}

func (x *sourceUploadStream) Send(m *Data) error {
	return x.stream.Send(m)
}

func (h *sourceHandler) Download(ctx context.Context, stream server.Stream) error {
	m := new(DownloadRequest)
	if err := stream.Recv(m); err != nil {
		return err
	}
	return h.SourceHandler.Download(ctx, m, &sourceDownloadStream{stream})
}

type Source_DownloadStream interface {
	Context() context.Context
	SendMsg(interface{}) error
	RecvMsg(interface{}) error
	Close() error
	Send(*Data) error
}

type sourceDownloadStream struct {
	stream server.Stream
}

func (x *sourceDownloadStream) Close() error {
	return x.stream.Close()
}

func (x *sourceDownloadStream) Context() context.Context {
	return x.stream.Context()
}

func (x *sourceDownloadStream) SendMsg(m interface{}) error {
	return x.stream.Send(m)
}

func (x *sourceDownloadStream) RecvMsg(m interface{}) error {
	return x.stream.Recv(m)
}

func (x *sourceDownloadStream) Send(m *Data) error {
	return x.stream.Send(m)
}