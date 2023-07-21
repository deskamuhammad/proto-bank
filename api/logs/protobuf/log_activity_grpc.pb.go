// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.21.12
// source: proto/log_activity.proto

package __

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

// ActivityLogServiceClient is the client API for ActivityLogService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type ActivityLogServiceClient interface {
	ServeInsertActivityLog(ctx context.Context, in *CreateActivityLogRequest, opts ...grpc.CallOption) (*ActivityLogsResponse, error)
	ServeGetListActivityLog(ctx context.Context, in *ActivityLogParameterRequest, opts ...grpc.CallOption) (*ActivityLogsResponse, error)
	ServeFindActivityByID(ctx context.Context, in *FindActivityLogRequest, opts ...grpc.CallOption) (*ActivityLog, error)
}

type activityLogServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewActivityLogServiceClient(cc grpc.ClientConnInterface) ActivityLogServiceClient {
	return &activityLogServiceClient{cc}
}

func (c *activityLogServiceClient) ServeInsertActivityLog(ctx context.Context, in *CreateActivityLogRequest, opts ...grpc.CallOption) (*ActivityLogsResponse, error) {
	out := new(ActivityLogsResponse)
	err := c.cc.Invoke(ctx, "/activitypb.ActivityLogService/ServeInsertActivityLog", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *activityLogServiceClient) ServeGetListActivityLog(ctx context.Context, in *ActivityLogParameterRequest, opts ...grpc.CallOption) (*ActivityLogsResponse, error) {
	out := new(ActivityLogsResponse)
	err := c.cc.Invoke(ctx, "/activitypb.ActivityLogService/ServeGetListActivityLog", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *activityLogServiceClient) ServeFindActivityByID(ctx context.Context, in *FindActivityLogRequest, opts ...grpc.CallOption) (*ActivityLog, error) {
	out := new(ActivityLog)
	err := c.cc.Invoke(ctx, "/activitypb.ActivityLogService/ServeFindActivityByID", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// ActivityLogServiceServer is the server API for ActivityLogService service.
// All implementations must embed UnimplementedActivityLogServiceServer
// for forward compatibility
type ActivityLogServiceServer interface {
	ServeInsertActivityLog(context.Context, *CreateActivityLogRequest) (*ActivityLogsResponse, error)
	ServeGetListActivityLog(context.Context, *ActivityLogParameterRequest) (*ActivityLogsResponse, error)
	ServeFindActivityByID(context.Context, *FindActivityLogRequest) (*ActivityLog, error)
	mustEmbedUnimplementedActivityLogServiceServer()
}

// UnimplementedActivityLogServiceServer must be embedded to have forward compatible implementations.
type UnimplementedActivityLogServiceServer struct {
}

func (UnimplementedActivityLogServiceServer) ServeInsertActivityLog(context.Context, *CreateActivityLogRequest) (*ActivityLogsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ServeInsertActivityLog not implemented")
}
func (UnimplementedActivityLogServiceServer) ServeGetListActivityLog(context.Context, *ActivityLogParameterRequest) (*ActivityLogsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ServeGetListActivityLog not implemented")
}
func (UnimplementedActivityLogServiceServer) ServeFindActivityByID(context.Context, *FindActivityLogRequest) (*ActivityLog, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ServeFindActivityByID not implemented")
}
func (UnimplementedActivityLogServiceServer) mustEmbedUnimplementedActivityLogServiceServer() {}

// UnsafeActivityLogServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to ActivityLogServiceServer will
// result in compilation errors.
type UnsafeActivityLogServiceServer interface {
	mustEmbedUnimplementedActivityLogServiceServer()
}

func RegisterActivityLogServiceServer(s grpc.ServiceRegistrar, srv ActivityLogServiceServer) {
	s.RegisterService(&ActivityLogService_ServiceDesc, srv)
}

func _ActivityLogService_ServeInsertActivityLog_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateActivityLogRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ActivityLogServiceServer).ServeInsertActivityLog(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/activitypb.ActivityLogService/ServeInsertActivityLog",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ActivityLogServiceServer).ServeInsertActivityLog(ctx, req.(*CreateActivityLogRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ActivityLogService_ServeGetListActivityLog_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ActivityLogParameterRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ActivityLogServiceServer).ServeGetListActivityLog(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/activitypb.ActivityLogService/ServeGetListActivityLog",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ActivityLogServiceServer).ServeGetListActivityLog(ctx, req.(*ActivityLogParameterRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ActivityLogService_ServeFindActivityByID_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(FindActivityLogRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ActivityLogServiceServer).ServeFindActivityByID(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/activitypb.ActivityLogService/ServeFindActivityByID",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ActivityLogServiceServer).ServeFindActivityByID(ctx, req.(*FindActivityLogRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// ActivityLogService_ServiceDesc is the grpc.ServiceDesc for ActivityLogService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var ActivityLogService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "activitypb.ActivityLogService",
	HandlerType: (*ActivityLogServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "ServeInsertActivityLog",
			Handler:    _ActivityLogService_ServeInsertActivityLog_Handler,
		},
		{
			MethodName: "ServeGetListActivityLog",
			Handler:    _ActivityLogService_ServeGetListActivityLog_Handler,
		},
		{
			MethodName: "ServeFindActivityByID",
			Handler:    _ActivityLogService_ServeFindActivityByID_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "proto/log_activity.proto",
}
