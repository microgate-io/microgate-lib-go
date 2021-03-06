// Code generated by protoc-gen-go-grpc. DO NOT EDIT.

package db

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

// DatabaseServiceClient is the client API for DatabaseService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type DatabaseServiceClient interface {
	Begin(ctx context.Context, in *BeginRequest, opts ...grpc.CallOption) (*BeginResponse, error)
	Commit(ctx context.Context, in *CommitRequest, opts ...grpc.CallOption) (*CommitResponse, error)
	Rollback(ctx context.Context, in *RollbackRequest, opts ...grpc.CallOption) (*RollbackResponse, error)
	Query(ctx context.Context, in *QueryRequest, opts ...grpc.CallOption) (*QueryResponse, error)
	Mutate(ctx context.Context, in *MutationRequest, opts ...grpc.CallOption) (*MutationResponse, error)
}

type databaseServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewDatabaseServiceClient(cc grpc.ClientConnInterface) DatabaseServiceClient {
	return &databaseServiceClient{cc}
}

func (c *databaseServiceClient) Begin(ctx context.Context, in *BeginRequest, opts ...grpc.CallOption) (*BeginResponse, error) {
	out := new(BeginResponse)
	err := c.cc.Invoke(ctx, "/microgate.DatabaseService/Begin", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *databaseServiceClient) Commit(ctx context.Context, in *CommitRequest, opts ...grpc.CallOption) (*CommitResponse, error) {
	out := new(CommitResponse)
	err := c.cc.Invoke(ctx, "/microgate.DatabaseService/Commit", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *databaseServiceClient) Rollback(ctx context.Context, in *RollbackRequest, opts ...grpc.CallOption) (*RollbackResponse, error) {
	out := new(RollbackResponse)
	err := c.cc.Invoke(ctx, "/microgate.DatabaseService/Rollback", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *databaseServiceClient) Query(ctx context.Context, in *QueryRequest, opts ...grpc.CallOption) (*QueryResponse, error) {
	out := new(QueryResponse)
	err := c.cc.Invoke(ctx, "/microgate.DatabaseService/Query", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *databaseServiceClient) Mutate(ctx context.Context, in *MutationRequest, opts ...grpc.CallOption) (*MutationResponse, error) {
	out := new(MutationResponse)
	err := c.cc.Invoke(ctx, "/microgate.DatabaseService/Mutate", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// DatabaseServiceServer is the server API for DatabaseService service.
// All implementations must embed UnimplementedDatabaseServiceServer
// for forward compatibility
type DatabaseServiceServer interface {
	Begin(context.Context, *BeginRequest) (*BeginResponse, error)
	Commit(context.Context, *CommitRequest) (*CommitResponse, error)
	Rollback(context.Context, *RollbackRequest) (*RollbackResponse, error)
	Query(context.Context, *QueryRequest) (*QueryResponse, error)
	Mutate(context.Context, *MutationRequest) (*MutationResponse, error)
	mustEmbedUnimplementedDatabaseServiceServer()
}

// UnimplementedDatabaseServiceServer must be embedded to have forward compatible implementations.
type UnimplementedDatabaseServiceServer struct {
}

func (UnimplementedDatabaseServiceServer) Begin(context.Context, *BeginRequest) (*BeginResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Begin not implemented")
}
func (UnimplementedDatabaseServiceServer) Commit(context.Context, *CommitRequest) (*CommitResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Commit not implemented")
}
func (UnimplementedDatabaseServiceServer) Rollback(context.Context, *RollbackRequest) (*RollbackResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Rollback not implemented")
}
func (UnimplementedDatabaseServiceServer) Query(context.Context, *QueryRequest) (*QueryResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Query not implemented")
}
func (UnimplementedDatabaseServiceServer) Mutate(context.Context, *MutationRequest) (*MutationResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Mutate not implemented")
}
func (UnimplementedDatabaseServiceServer) mustEmbedUnimplementedDatabaseServiceServer() {}

// UnsafeDatabaseServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to DatabaseServiceServer will
// result in compilation errors.
type UnsafeDatabaseServiceServer interface {
	mustEmbedUnimplementedDatabaseServiceServer()
}

func RegisterDatabaseServiceServer(s grpc.ServiceRegistrar, srv DatabaseServiceServer) {
	s.RegisterService(&DatabaseService_ServiceDesc, srv)
}

func _DatabaseService_Begin_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(BeginRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DatabaseServiceServer).Begin(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/microgate.DatabaseService/Begin",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DatabaseServiceServer).Begin(ctx, req.(*BeginRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _DatabaseService_Commit_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CommitRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DatabaseServiceServer).Commit(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/microgate.DatabaseService/Commit",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DatabaseServiceServer).Commit(ctx, req.(*CommitRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _DatabaseService_Rollback_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RollbackRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DatabaseServiceServer).Rollback(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/microgate.DatabaseService/Rollback",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DatabaseServiceServer).Rollback(ctx, req.(*RollbackRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _DatabaseService_Query_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(QueryRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DatabaseServiceServer).Query(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/microgate.DatabaseService/Query",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DatabaseServiceServer).Query(ctx, req.(*QueryRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _DatabaseService_Mutate_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(MutationRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DatabaseServiceServer).Mutate(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/microgate.DatabaseService/Mutate",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DatabaseServiceServer).Mutate(ctx, req.(*MutationRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// DatabaseService_ServiceDesc is the grpc.ServiceDesc for DatabaseService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var DatabaseService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "microgate.DatabaseService",
	HandlerType: (*DatabaseServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Begin",
			Handler:    _DatabaseService_Begin_Handler,
		},
		{
			MethodName: "Commit",
			Handler:    _DatabaseService_Commit_Handler,
		},
		{
			MethodName: "Rollback",
			Handler:    _DatabaseService_Rollback_Handler,
		},
		{
			MethodName: "Query",
			Handler:    _DatabaseService_Query_Handler,
		},
		{
			MethodName: "Mutate",
			Handler:    _DatabaseService_Mutate_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "db.proto",
}
