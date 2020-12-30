// Code generated by protoc-gen-go-grpc. DO NOT EDIT.

package client

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion7

// DNSClient is the client API for DNS service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type DNSClient interface {
	Ping(ctx context.Context, in *Message, opts ...grpc.CallOption) (*Message, error)
	Action(ctx context.Context, in *Command, opts ...grpc.CallOption) (*VectorClock, error)
	Spread(ctx context.Context, in *Log, opts ...grpc.CallOption) (*Message, error)
}

type dNSClient struct {
	cc grpc.ClientConnInterface
}

func NewDNSClient(cc grpc.ClientConnInterface) DNSClient {
	return &dNSClient{cc}
}

func (c *dNSClient) Ping(ctx context.Context, in *Message, opts ...grpc.CallOption) (*Message, error) {
	out := new(Message)
	err := c.cc.Invoke(ctx, "/lab3.DNS/Ping", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *dNSClient) Action(ctx context.Context, in *Command, opts ...grpc.CallOption) (*VectorClock, error) {
	out := new(VectorClock)
	err := c.cc.Invoke(ctx, "/lab3.DNS/Action", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *dNSClient) Spread(ctx context.Context, in *Log, opts ...grpc.CallOption) (*Message, error) {
	out := new(Message)
	err := c.cc.Invoke(ctx, "/lab3.DNS/Spread", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// DNSServer is the server API for DNS service.
// All implementations must embed UnimplementedDNSServer
// for forward compatibility
type DNSServer interface {
	Ping(context.Context, *Message) (*Message, error)
	Action(context.Context, *Command) (*VectorClock, error)
	Spread(context.Context, *Log) (*Message, error)
	mustEmbedUnimplementedDNSServer()
}

// UnimplementedDNSServer must be embedded to have forward compatible implementations.
type UnimplementedDNSServer struct {
}

func (UnimplementedDNSServer) Ping(context.Context, *Message) (*Message, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Ping not implemented")
}
func (UnimplementedDNSServer) Action(context.Context, *Command) (*VectorClock, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Action not implemented")
}
func (UnimplementedDNSServer) Spread(context.Context, *Log) (*Message, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Spread not implemented")
}
func (UnimplementedDNSServer) mustEmbedUnimplementedDNSServer() {}

// UnsafeDNSServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to DNSServer will
// result in compilation errors.
type UnsafeDNSServer interface {
	mustEmbedUnimplementedDNSServer()
}

func RegisterDNSServer(s grpc.ServiceRegistrar, srv DNSServer) {
	s.RegisterService(&DNS_ServiceDesc, srv)
}

func _DNS_Ping_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Message)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DNSServer).Ping(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/lab3.DNS/Ping",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DNSServer).Ping(ctx, req.(*Message))
	}
	return interceptor(ctx, in, info, handler)
}

func _DNS_Action_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Command)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DNSServer).Action(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/lab3.DNS/Action",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DNSServer).Action(ctx, req.(*Command))
	}
	return interceptor(ctx, in, info, handler)
}

func _DNS_Spread_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Log)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DNSServer).Spread(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/lab3.DNS/Spread",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DNSServer).Spread(ctx, req.(*Log))
	}
	return interceptor(ctx, in, info, handler)
}

// DNS_ServiceDesc is the grpc.ServiceDesc for DNS service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var DNS_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "lab3.DNS",
	HandlerType: (*DNSServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Ping",
			Handler:    _DNS_Ping_Handler,
		},
		{
			MethodName: "Action",
			Handler:    _DNS_Action_Handler,
		},
		{
			MethodName: "Spread",
			Handler:    _DNS_Spread_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "pkg/proto/l3.proto",
}

// BrokerClient is the client API for Broker service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type BrokerClient interface {
	DNSIsAvailable(ctx context.Context, in *Message, opts ...grpc.CallOption) (*DNSState, error)
	GetIP(ctx context.Context, in *Command, opts ...grpc.CallOption) (*PageInfo, error)
}

type brokerClient struct {
	cc grpc.ClientConnInterface
}

func NewBrokerClient(cc grpc.ClientConnInterface) BrokerClient {
	return &brokerClient{cc}
}

func (c *brokerClient) DNSIsAvailable(ctx context.Context, in *Message, opts ...grpc.CallOption) (*DNSState, error) {
	out := new(DNSState)
	err := c.cc.Invoke(ctx, "/lab3.Broker/DNSIsAvailable", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *brokerClient) GetIP(ctx context.Context, in *Command, opts ...grpc.CallOption) (*PageInfo, error) {
	out := new(PageInfo)
	err := c.cc.Invoke(ctx, "/lab3.Broker/GetIP", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// BrokerServer is the server API for Broker service.
// All implementations must embed UnimplementedBrokerServer
// for forward compatibility
type BrokerServer interface {
	DNSIsAvailable(context.Context, *Message) (*DNSState, error)
	GetIP(context.Context, *Command) (*PageInfo, error)
	mustEmbedUnimplementedBrokerServer()
}

// UnimplementedBrokerServer must be embedded to have forward compatible implementations.
type UnimplementedBrokerServer struct {
}

func (UnimplementedBrokerServer) DNSIsAvailable(context.Context, *Message) (*DNSState, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DNSIsAvailable not implemented")
}
func (UnimplementedBrokerServer) GetIP(context.Context, *Command) (*PageInfo, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetIP not implemented")
}
func (UnimplementedBrokerServer) mustEmbedUnimplementedBrokerServer() {}

// UnsafeBrokerServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to BrokerServer will
// result in compilation errors.
type UnsafeBrokerServer interface {
	mustEmbedUnimplementedBrokerServer()
}

func RegisterBrokerServer(s grpc.ServiceRegistrar, srv BrokerServer) {
	s.RegisterService(&Broker_ServiceDesc, srv)
}

func _Broker_DNSIsAvailable_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Message)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BrokerServer).DNSIsAvailable(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/lab3.Broker/DNSIsAvailable",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BrokerServer).DNSIsAvailable(ctx, req.(*Message))
	}
	return interceptor(ctx, in, info, handler)
}

func _Broker_GetIP_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Command)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BrokerServer).GetIP(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/lab3.Broker/GetIP",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BrokerServer).GetIP(ctx, req.(*Command))
	}
	return interceptor(ctx, in, info, handler)
}

// Broker_ServiceDesc is the grpc.ServiceDesc for Broker service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Broker_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "lab3.Broker",
	HandlerType: (*BrokerServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "DNSIsAvailable",
			Handler:    _Broker_DNSIsAvailable_Handler,
		},
		{
			MethodName: "GetIP",
			Handler:    _Broker_GetIP_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "pkg/proto/l3.proto",
}
