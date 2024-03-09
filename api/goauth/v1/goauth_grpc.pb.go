// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             v3.21.6
// source: api/goauth/v1/goauth.proto

package v1

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

const (
	GoAuth_Register_FullMethodName = "/api.goauth.v1.GoAuth/Register"
	GoAuth_Login_FullMethodName    = "/api.goauth.v1.GoAuth/Login"
)

// GoAuthClient is the client API for GoAuth service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type GoAuthClient interface {
	// User api
	Register(ctx context.Context, in *RegisterRequest, opts ...grpc.CallOption) (*RegisterReply, error)
	Login(ctx context.Context, in *LoginRequest, opts ...grpc.CallOption) (*LoginReply, error)
}

type goAuthClient struct {
	cc grpc.ClientConnInterface
}

func NewGoAuthClient(cc grpc.ClientConnInterface) GoAuthClient {
	return &goAuthClient{cc}
}

func (c *goAuthClient) Register(ctx context.Context, in *RegisterRequest, opts ...grpc.CallOption) (*RegisterReply, error) {
	out := new(RegisterReply)
	err := c.cc.Invoke(ctx, GoAuth_Register_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *goAuthClient) Login(ctx context.Context, in *LoginRequest, opts ...grpc.CallOption) (*LoginReply, error) {
	out := new(LoginReply)
	err := c.cc.Invoke(ctx, GoAuth_Login_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// GoAuthServer is the server API for GoAuth service.
// All implementations must embed UnimplementedGoAuthServer
// for forward compatibility
type GoAuthServer interface {
	// User api
	Register(context.Context, *RegisterRequest) (*RegisterReply, error)
	Login(context.Context, *LoginRequest) (*LoginReply, error)
	mustEmbedUnimplementedGoAuthServer()
}

// UnimplementedGoAuthServer must be embedded to have forward compatible implementations.
type UnimplementedGoAuthServer struct {
}

func (UnimplementedGoAuthServer) Register(context.Context, *RegisterRequest) (*RegisterReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Register not implemented")
}
func (UnimplementedGoAuthServer) Login(context.Context, *LoginRequest) (*LoginReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Login not implemented")
}
func (UnimplementedGoAuthServer) mustEmbedUnimplementedGoAuthServer() {}

// UnsafeGoAuthServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to GoAuthServer will
// result in compilation errors.
type UnsafeGoAuthServer interface {
	mustEmbedUnimplementedGoAuthServer()
}

func RegisterGoAuthServer(s grpc.ServiceRegistrar, srv GoAuthServer) {
	s.RegisterService(&GoAuth_ServiceDesc, srv)
}

func _GoAuth_Register_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RegisterRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GoAuthServer).Register(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: GoAuth_Register_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GoAuthServer).Register(ctx, req.(*RegisterRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _GoAuth_Login_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(LoginRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GoAuthServer).Login(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: GoAuth_Login_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GoAuthServer).Login(ctx, req.(*LoginRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// GoAuth_ServiceDesc is the grpc.ServiceDesc for GoAuth service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var GoAuth_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "api.goauth.v1.GoAuth",
	HandlerType: (*GoAuthServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Register",
			Handler:    _GoAuth_Register_Handler,
		},
		{
			MethodName: "Login",
			Handler:    _GoAuth_Login_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "api/goauth/v1/goauth.proto",
}