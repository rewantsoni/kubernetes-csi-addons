// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.4.0
// - protoc             v3.20.2
// source: encryptionkeyrotation.proto

package proto

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.62.0 or later.
const _ = grpc.SupportPackageIsVersion8

const (
	EncryptionKeyRotation_EncryptionKeyRotate_FullMethodName = "/proto.EncryptionKeyRotation/EncryptionKeyRotate"
)

// EncryptionKeyRotationClient is the client API for EncryptionKeyRotation service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
//
// EncryptionKeyRotation holds the RPC method for running
// key rotation operation on the pv.
type EncryptionKeyRotationClient interface {
	EncryptionKeyRotate(ctx context.Context, in *EncryptionKeyRotateRequest, opts ...grpc.CallOption) (*EncryptionKeyRotateResponse, error)
}

type encryptionKeyRotationClient struct {
	cc grpc.ClientConnInterface
}

func NewEncryptionKeyRotationClient(cc grpc.ClientConnInterface) EncryptionKeyRotationClient {
	return &encryptionKeyRotationClient{cc}
}

func (c *encryptionKeyRotationClient) EncryptionKeyRotate(ctx context.Context, in *EncryptionKeyRotateRequest, opts ...grpc.CallOption) (*EncryptionKeyRotateResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(EncryptionKeyRotateResponse)
	err := c.cc.Invoke(ctx, EncryptionKeyRotation_EncryptionKeyRotate_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// EncryptionKeyRotationServer is the server API for EncryptionKeyRotation service.
// All implementations must embed UnimplementedEncryptionKeyRotationServer
// for forward compatibility
//
// EncryptionKeyRotation holds the RPC method for running
// key rotation operation on the pv.
type EncryptionKeyRotationServer interface {
	EncryptionKeyRotate(context.Context, *EncryptionKeyRotateRequest) (*EncryptionKeyRotateResponse, error)
	mustEmbedUnimplementedEncryptionKeyRotationServer()
}

// UnimplementedEncryptionKeyRotationServer must be embedded to have forward compatible implementations.
type UnimplementedEncryptionKeyRotationServer struct {
}

func (UnimplementedEncryptionKeyRotationServer) EncryptionKeyRotate(context.Context, *EncryptionKeyRotateRequest) (*EncryptionKeyRotateResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method EncryptionKeyRotate not implemented")
}
func (UnimplementedEncryptionKeyRotationServer) mustEmbedUnimplementedEncryptionKeyRotationServer() {}

// UnsafeEncryptionKeyRotationServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to EncryptionKeyRotationServer will
// result in compilation errors.
type UnsafeEncryptionKeyRotationServer interface {
	mustEmbedUnimplementedEncryptionKeyRotationServer()
}

func RegisterEncryptionKeyRotationServer(s grpc.ServiceRegistrar, srv EncryptionKeyRotationServer) {
	s.RegisterService(&EncryptionKeyRotation_ServiceDesc, srv)
}

func _EncryptionKeyRotation_EncryptionKeyRotate_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(EncryptionKeyRotateRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(EncryptionKeyRotationServer).EncryptionKeyRotate(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: EncryptionKeyRotation_EncryptionKeyRotate_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(EncryptionKeyRotationServer).EncryptionKeyRotate(ctx, req.(*EncryptionKeyRotateRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// EncryptionKeyRotation_ServiceDesc is the grpc.ServiceDesc for EncryptionKeyRotation service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var EncryptionKeyRotation_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "proto.EncryptionKeyRotation",
	HandlerType: (*EncryptionKeyRotationServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "EncryptionKeyRotate",
			Handler:    _EncryptionKeyRotation_EncryptionKeyRotate_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "encryptionkeyrotation.proto",
}
