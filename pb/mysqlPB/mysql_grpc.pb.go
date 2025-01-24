// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.5.1
// - protoc             v3.21.12
// source: pb/mysqlPB/mysql.proto

package __

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.64.0 or later.
const _ = grpc.SupportPackageIsVersion9

const (
	MySQLService_Connect_FullMethodName      = "/mysqlPB.MySQLService/Connect"
	MySQLService_Disconnect_FullMethodName   = "/mysqlPB.MySQLService/Disconnect"
	MySQLService_ExecuteQuery_FullMethodName = "/mysqlPB.MySQLService/ExecuteQuery"
)

// MySQLServiceClient is the client API for MySQLService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
//
// MySQL 服务定义
type MySQLServiceClient interface {
	// 连接到数据库
	Connect(ctx context.Context, in *ConnectRequest, opts ...grpc.CallOption) (*ConnectResponse, error)
	// 断开与数据库的连接
	Disconnect(ctx context.Context, in *DisconnectRequest, opts ...grpc.CallOption) (*DisconnectResponse, error)
	// 执行 SQL 查询
	ExecuteQuery(ctx context.Context, in *ExecuteQueryRequest, opts ...grpc.CallOption) (*ExecuteQueryResponse, error)
}

type mySQLServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewMySQLServiceClient(cc grpc.ClientConnInterface) MySQLServiceClient {
	return &mySQLServiceClient{cc}
}

func (c *mySQLServiceClient) Connect(ctx context.Context, in *ConnectRequest, opts ...grpc.CallOption) (*ConnectResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(ConnectResponse)
	err := c.cc.Invoke(ctx, MySQLService_Connect_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *mySQLServiceClient) Disconnect(ctx context.Context, in *DisconnectRequest, opts ...grpc.CallOption) (*DisconnectResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(DisconnectResponse)
	err := c.cc.Invoke(ctx, MySQLService_Disconnect_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *mySQLServiceClient) ExecuteQuery(ctx context.Context, in *ExecuteQueryRequest, opts ...grpc.CallOption) (*ExecuteQueryResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(ExecuteQueryResponse)
	err := c.cc.Invoke(ctx, MySQLService_ExecuteQuery_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// MySQLServiceServer is the server API for MySQLService service.
// All implementations must embed UnimplementedMySQLServiceServer
// for forward compatibility.
//
// MySQL 服务定义
type MySQLServiceServer interface {
	// 连接到数据库
	Connect(context.Context, *ConnectRequest) (*ConnectResponse, error)
	// 断开与数据库的连接
	Disconnect(context.Context, *DisconnectRequest) (*DisconnectResponse, error)
	// 执行 SQL 查询
	ExecuteQuery(context.Context, *ExecuteQueryRequest) (*ExecuteQueryResponse, error)
	mustEmbedUnimplementedMySQLServiceServer()
}

// UnimplementedMySQLServiceServer must be embedded to have
// forward compatible implementations.
//
// NOTE: this should be embedded by value instead of pointer to avoid a nil
// pointer dereference when methods are called.
type UnimplementedMySQLServiceServer struct{}

func (UnimplementedMySQLServiceServer) Connect(context.Context, *ConnectRequest) (*ConnectResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Connect not implemented")
}
func (UnimplementedMySQLServiceServer) Disconnect(context.Context, *DisconnectRequest) (*DisconnectResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Disconnect not implemented")
}
func (UnimplementedMySQLServiceServer) ExecuteQuery(context.Context, *ExecuteQueryRequest) (*ExecuteQueryResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ExecuteQuery not implemented")
}
func (UnimplementedMySQLServiceServer) mustEmbedUnimplementedMySQLServiceServer() {}
func (UnimplementedMySQLServiceServer) testEmbeddedByValue()                      {}

// UnsafeMySQLServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to MySQLServiceServer will
// result in compilation errors.
type UnsafeMySQLServiceServer interface {
	mustEmbedUnimplementedMySQLServiceServer()
}

func RegisterMySQLServiceServer(s grpc.ServiceRegistrar, srv MySQLServiceServer) {
	// If the following call pancis, it indicates UnimplementedMySQLServiceServer was
	// embedded by pointer and is nil.  This will cause panics if an
	// unimplemented method is ever invoked, so we test this at initialization
	// time to prevent it from happening at runtime later due to I/O.
	if t, ok := srv.(interface{ testEmbeddedByValue() }); ok {
		t.testEmbeddedByValue()
	}
	s.RegisterService(&MySQLService_ServiceDesc, srv)
}

func _MySQLService_Connect_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ConnectRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MySQLServiceServer).Connect(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: MySQLService_Connect_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MySQLServiceServer).Connect(ctx, req.(*ConnectRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _MySQLService_Disconnect_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DisconnectRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MySQLServiceServer).Disconnect(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: MySQLService_Disconnect_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MySQLServiceServer).Disconnect(ctx, req.(*DisconnectRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _MySQLService_ExecuteQuery_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ExecuteQueryRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MySQLServiceServer).ExecuteQuery(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: MySQLService_ExecuteQuery_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MySQLServiceServer).ExecuteQuery(ctx, req.(*ExecuteQueryRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// MySQLService_ServiceDesc is the grpc.ServiceDesc for MySQLService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var MySQLService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "mysqlPB.MySQLService",
	HandlerType: (*MySQLServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Connect",
			Handler:    _MySQLService_Connect_Handler,
		},
		{
			MethodName: "Disconnect",
			Handler:    _MySQLService_Disconnect_Handler,
		},
		{
			MethodName: "ExecuteQuery",
			Handler:    _MySQLService_ExecuteQuery_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "pb/mysqlPB/mysql.proto",
}