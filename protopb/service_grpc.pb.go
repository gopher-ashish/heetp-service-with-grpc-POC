// Code generated by protoc-gen-go-grpc. DO NOT EDIT.

package protopb

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

// SumClient is the client API for Sum service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type SumClient interface {
	// unary rpc service
	WelcomeEmail(ctx context.Context, in *EmailRequest, opts ...grpc.CallOption) (*EmailResponse, error)
}

type sumClient struct {
	cc grpc.ClientConnInterface
}

func NewSumClient(cc grpc.ClientConnInterface) SumClient {
	return &sumClient{cc}
}

func (c *sumClient) WelcomeEmail(ctx context.Context, in *EmailRequest, opts ...grpc.CallOption) (*EmailResponse, error) {
	out := new(EmailResponse)
	err := c.cc.Invoke(ctx, "/testing.Sum/WelcomeEmail", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// SumServer is the server API for Sum service.
// All implementations must embed UnimplementedSumServer
// for forward compatibility
type SumServer interface {
	// unary rpc service
	WelcomeEmail(context.Context, *EmailRequest) (*EmailResponse, error)
	mustEmbedUnimplementedSumServer()
}

// UnimplementedSumServer must be embedded to have forward compatible implementations.
type UnimplementedSumServer struct {
}

func (UnimplementedSumServer) WelcomeEmail(context.Context, *EmailRequest) (*EmailResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method WelcomeEmail not implemented")
}
func (UnimplementedSumServer) mustEmbedUnimplementedSumServer() {}

// UnsafeSumServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to SumServer will
// result in compilation errors.
type UnsafeSumServer interface {
	mustEmbedUnimplementedSumServer()
}

func RegisterSumServer(s grpc.ServiceRegistrar, srv SumServer) {
	s.RegisterService(&Sum_ServiceDesc, srv)
}

func _Sum_WelcomeEmail_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(EmailRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SumServer).WelcomeEmail(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/testing.Sum/WelcomeEmail",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SumServer).WelcomeEmail(ctx, req.(*EmailRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// Sum_ServiceDesc is the grpc.ServiceDesc for Sum service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Sum_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "testing.Sum",
	HandlerType: (*SumServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "WelcomeEmail",
			Handler:    _Sum_WelcomeEmail_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "service.proto",
}
