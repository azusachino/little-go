// Code generated by protoc-gen-go-grpc. DO NOT EDIT.

package proto

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.32.0 or later.
const _ = grpc.SupportPackageIsVersion7

// MailServiceClient is the client API for MailService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type MailServiceClient interface {
	SendMail(ctx context.Context, in *MailRequest, opts ...grpc.CallOption) (*MailResponse, error)
}

type mailServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewMailServiceClient(cc grpc.ClientConnInterface) MailServiceClient {
	return &mailServiceClient{cc}
}

func (c *mailServiceClient) SendMail(ctx context.Context, in *MailRequest, opts ...grpc.CallOption) (*MailResponse, error) {
	out := new(MailResponse)
	err := c.cc.Invoke(ctx, "/mail.MailService/SendMail", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// MailServiceServer is the server API for MailService service.
// All implementations must embed UnimplementedMailServiceServer
// for forward compatibility
type MailServiceServer interface {
	SendMail(context.Context, *MailRequest) (*MailResponse, error)
	mustEmbedUnimplementedMailServiceServer()
}

// UnimplementedMailServiceServer must be embedded to have forward compatible implementations.
type UnimplementedMailServiceServer struct {
}

func (UnimplementedMailServiceServer) SendMail(context.Context, *MailRequest) (*MailResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SendMail not implemented")
}
func (UnimplementedMailServiceServer) mustEmbedUnimplementedMailServiceServer() {}

// UnsafeMailServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to MailServiceServer will
// result in compilation errors.
type UnsafeMailServiceServer interface {
	mustEmbedUnimplementedMailServiceServer()
}

func RegisterMailServiceServer(s grpc.ServiceRegistrar, srv MailServiceServer) {
	s.RegisterService(&MailService_ServiceDesc, srv)
}

func _MailService_SendMail_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(MailRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MailServiceServer).SendMail(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/mail.MailService/SendMail",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MailServiceServer).SendMail(ctx, req.(*MailRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// MailService_ServiceDesc is the grpc.ServiceDesc for MailService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var MailService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "mail.MailService",
	HandlerType: (*MailServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "SendMail",
			Handler:    _MailService_SendMail_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "mail/mail.proto",
}
