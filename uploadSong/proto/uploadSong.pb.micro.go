// Code generated by protoc-gen-micro. DO NOT EDIT.
// source: proto/uploadSong.proto

package uploadSong

import (
	fmt "fmt"
	proto "google.golang.org/protobuf/proto"
	math "math"
)

import (
	context "context"
	api "go-micro.dev/v4/api"
	client "go-micro.dev/v4/client"
	server "go-micro.dev/v4/server"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// Reference imports to suppress errors if they are not otherwise used.
var _ api.Endpoint
var _ context.Context
var _ client.Option
var _ server.Option

// Api Endpoints for UploadSong service

func NewUploadSongEndpoints() []*api.Endpoint {
	return []*api.Endpoint{}
}

// Client API for UploadSong service

type UploadSongService interface {
	Call(ctx context.Context, in *CallRequest, opts ...client.CallOption) (*CallResponse, error)
	ClientStream(ctx context.Context, opts ...client.CallOption) (UploadSong_ClientStreamService, error)
	ServerStream(ctx context.Context, in *ServerStreamRequest, opts ...client.CallOption) (UploadSong_ServerStreamService, error)
	BidiStream(ctx context.Context, opts ...client.CallOption) (UploadSong_BidiStreamService, error)
}

type uploadSongService struct {
	c    client.Client
	name string
}

func NewUploadSongService(name string, c client.Client) UploadSongService {
	return &uploadSongService{
		c:    c,
		name: name,
	}
}

func (c *uploadSongService) Call(ctx context.Context, in *CallRequest, opts ...client.CallOption) (*CallResponse, error) {
	req := c.c.NewRequest(c.name, "UploadSong.Call", in)
	out := new(CallResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *uploadSongService) ClientStream(ctx context.Context, opts ...client.CallOption) (UploadSong_ClientStreamService, error) {
	req := c.c.NewRequest(c.name, "UploadSong.ClientStream", &ClientStreamRequest{})
	stream, err := c.c.Stream(ctx, req, opts...)
	if err != nil {
		return nil, err
	}
	return &uploadSongServiceClientStream{stream}, nil
}

type UploadSong_ClientStreamService interface {
	Context() context.Context
	SendMsg(interface{}) error
	RecvMsg(interface{}) error
	CloseSend() error
	Close() error
	Send(*ClientStreamRequest) error
}

type uploadSongServiceClientStream struct {
	stream client.Stream
}

func (x *uploadSongServiceClientStream) CloseSend() error {
	return x.stream.CloseSend()
}

func (x *uploadSongServiceClientStream) Close() error {
	return x.stream.Close()
}

func (x *uploadSongServiceClientStream) Context() context.Context {
	return x.stream.Context()
}

func (x *uploadSongServiceClientStream) SendMsg(m interface{}) error {
	return x.stream.Send(m)
}

func (x *uploadSongServiceClientStream) RecvMsg(m interface{}) error {
	return x.stream.Recv(m)
}

func (x *uploadSongServiceClientStream) Send(m *ClientStreamRequest) error {
	return x.stream.Send(m)
}

func (c *uploadSongService) ServerStream(ctx context.Context, in *ServerStreamRequest, opts ...client.CallOption) (UploadSong_ServerStreamService, error) {
	req := c.c.NewRequest(c.name, "UploadSong.ServerStream", &ServerStreamRequest{})
	stream, err := c.c.Stream(ctx, req, opts...)
	if err != nil {
		return nil, err
	}
	if err := stream.Send(in); err != nil {
		return nil, err
	}
	return &uploadSongServiceServerStream{stream}, nil
}

type UploadSong_ServerStreamService interface {
	Context() context.Context
	SendMsg(interface{}) error
	RecvMsg(interface{}) error
	CloseSend() error
	Close() error
	Recv() (*ServerStreamResponse, error)
}

type uploadSongServiceServerStream struct {
	stream client.Stream
}

func (x *uploadSongServiceServerStream) CloseSend() error {
	return x.stream.CloseSend()
}

func (x *uploadSongServiceServerStream) Close() error {
	return x.stream.Close()
}

func (x *uploadSongServiceServerStream) Context() context.Context {
	return x.stream.Context()
}

func (x *uploadSongServiceServerStream) SendMsg(m interface{}) error {
	return x.stream.Send(m)
}

func (x *uploadSongServiceServerStream) RecvMsg(m interface{}) error {
	return x.stream.Recv(m)
}

func (x *uploadSongServiceServerStream) Recv() (*ServerStreamResponse, error) {
	m := new(ServerStreamResponse)
	err := x.stream.Recv(m)
	if err != nil {
		return nil, err
	}
	return m, nil
}

func (c *uploadSongService) BidiStream(ctx context.Context, opts ...client.CallOption) (UploadSong_BidiStreamService, error) {
	req := c.c.NewRequest(c.name, "UploadSong.BidiStream", &BidiStreamRequest{})
	stream, err := c.c.Stream(ctx, req, opts...)
	if err != nil {
		return nil, err
	}
	return &uploadSongServiceBidiStream{stream}, nil
}

type UploadSong_BidiStreamService interface {
	Context() context.Context
	SendMsg(interface{}) error
	RecvMsg(interface{}) error
	CloseSend() error
	Close() error
	Send(*BidiStreamRequest) error
	Recv() (*BidiStreamResponse, error)
}

type uploadSongServiceBidiStream struct {
	stream client.Stream
}

func (x *uploadSongServiceBidiStream) CloseSend() error {
	return x.stream.CloseSend()
}

func (x *uploadSongServiceBidiStream) Close() error {
	return x.stream.Close()
}

func (x *uploadSongServiceBidiStream) Context() context.Context {
	return x.stream.Context()
}

func (x *uploadSongServiceBidiStream) SendMsg(m interface{}) error {
	return x.stream.Send(m)
}

func (x *uploadSongServiceBidiStream) RecvMsg(m interface{}) error {
	return x.stream.Recv(m)
}

func (x *uploadSongServiceBidiStream) Send(m *BidiStreamRequest) error {
	return x.stream.Send(m)
}

func (x *uploadSongServiceBidiStream) Recv() (*BidiStreamResponse, error) {
	m := new(BidiStreamResponse)
	err := x.stream.Recv(m)
	if err != nil {
		return nil, err
	}
	return m, nil
}

// Server API for UploadSong service

type UploadSongHandler interface {
	Call(context.Context, *CallRequest, *CallResponse) error
	ClientStream(context.Context, UploadSong_ClientStreamStream) error
	ServerStream(context.Context, *ServerStreamRequest, UploadSong_ServerStreamStream) error
	BidiStream(context.Context, UploadSong_BidiStreamStream) error
}

func RegisterUploadSongHandler(s server.Server, hdlr UploadSongHandler, opts ...server.HandlerOption) error {
	type uploadSong interface {
		Call(ctx context.Context, in *CallRequest, out *CallResponse) error
		ClientStream(ctx context.Context, stream server.Stream) error
		ServerStream(ctx context.Context, stream server.Stream) error
		BidiStream(ctx context.Context, stream server.Stream) error
	}
	type UploadSong struct {
		uploadSong
	}
	h := &uploadSongHandler{hdlr}
	return s.Handle(s.NewHandler(&UploadSong{h}, opts...))
}

type uploadSongHandler struct {
	UploadSongHandler
}

func (h *uploadSongHandler) Call(ctx context.Context, in *CallRequest, out *CallResponse) error {
	return h.UploadSongHandler.Call(ctx, in, out)
}

func (h *uploadSongHandler) ClientStream(ctx context.Context, stream server.Stream) error {
	return h.UploadSongHandler.ClientStream(ctx, &uploadSongClientStreamStream{stream})
}

type UploadSong_ClientStreamStream interface {
	Context() context.Context
	SendMsg(interface{}) error
	RecvMsg(interface{}) error
	Close() error
	Recv() (*ClientStreamRequest, error)
}

type uploadSongClientStreamStream struct {
	stream server.Stream
}

func (x *uploadSongClientStreamStream) Close() error {
	return x.stream.Close()
}

func (x *uploadSongClientStreamStream) Context() context.Context {
	return x.stream.Context()
}

func (x *uploadSongClientStreamStream) SendMsg(m interface{}) error {
	return x.stream.Send(m)
}

func (x *uploadSongClientStreamStream) RecvMsg(m interface{}) error {
	return x.stream.Recv(m)
}

func (x *uploadSongClientStreamStream) Recv() (*ClientStreamRequest, error) {
	m := new(ClientStreamRequest)
	if err := x.stream.Recv(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (h *uploadSongHandler) ServerStream(ctx context.Context, stream server.Stream) error {
	m := new(ServerStreamRequest)
	if err := stream.Recv(m); err != nil {
		return err
	}
	return h.UploadSongHandler.ServerStream(ctx, m, &uploadSongServerStreamStream{stream})
}

type UploadSong_ServerStreamStream interface {
	Context() context.Context
	SendMsg(interface{}) error
	RecvMsg(interface{}) error
	Close() error
	Send(*ServerStreamResponse) error
}

type uploadSongServerStreamStream struct {
	stream server.Stream
}

func (x *uploadSongServerStreamStream) Close() error {
	return x.stream.Close()
}

func (x *uploadSongServerStreamStream) Context() context.Context {
	return x.stream.Context()
}

func (x *uploadSongServerStreamStream) SendMsg(m interface{}) error {
	return x.stream.Send(m)
}

func (x *uploadSongServerStreamStream) RecvMsg(m interface{}) error {
	return x.stream.Recv(m)
}

func (x *uploadSongServerStreamStream) Send(m *ServerStreamResponse) error {
	return x.stream.Send(m)
}

func (h *uploadSongHandler) BidiStream(ctx context.Context, stream server.Stream) error {
	return h.UploadSongHandler.BidiStream(ctx, &uploadSongBidiStreamStream{stream})
}

type UploadSong_BidiStreamStream interface {
	Context() context.Context
	SendMsg(interface{}) error
	RecvMsg(interface{}) error
	Close() error
	Send(*BidiStreamResponse) error
	Recv() (*BidiStreamRequest, error)
}

type uploadSongBidiStreamStream struct {
	stream server.Stream
}

func (x *uploadSongBidiStreamStream) Close() error {
	return x.stream.Close()
}

func (x *uploadSongBidiStreamStream) Context() context.Context {
	return x.stream.Context()
}

func (x *uploadSongBidiStreamStream) SendMsg(m interface{}) error {
	return x.stream.Send(m)
}

func (x *uploadSongBidiStreamStream) RecvMsg(m interface{}) error {
	return x.stream.Recv(m)
}

func (x *uploadSongBidiStreamStream) Send(m *BidiStreamResponse) error {
	return x.stream.Send(m)
}

func (x *uploadSongBidiStreamStream) Recv() (*BidiStreamRequest, error) {
	m := new(BidiStreamRequest)
	if err := x.stream.Recv(m); err != nil {
		return nil, err
	}
	return m, nil
}
