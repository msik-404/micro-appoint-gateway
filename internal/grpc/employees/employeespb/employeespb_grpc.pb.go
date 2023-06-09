// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.21.12
// source: employeespb.proto

package employeespb

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
	emptypb "google.golang.org/protobuf/types/known/emptypb"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.32.0 or later.
const _ = grpc.SupportPackageIsVersion7

// ApiClient is the client API for Api service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type ApiClient interface {
	AddEmployee(ctx context.Context, in *AddEmployeeRequest, opts ...grpc.CallOption) (*emptypb.Empty, error)
	UpdateEmployee(ctx context.Context, in *UpdateEmployeeRequest, opts ...grpc.CallOption) (*emptypb.Empty, error)
	DeleteEmployee(ctx context.Context, in *DeleteEmployeeRequest, opts ...grpc.CallOption) (*emptypb.Empty, error)
	FindOneEmployee(ctx context.Context, in *EmployeeRequest, opts ...grpc.CallOption) (*EmployeeReply, error)
	FindManyEmployees(ctx context.Context, in *EmployeesRequest, opts ...grpc.CallOption) (*EmployeesReply, error)
}

type apiClient struct {
	cc grpc.ClientConnInterface
}

func NewApiClient(cc grpc.ClientConnInterface) ApiClient {
	return &apiClient{cc}
}

func (c *apiClient) AddEmployee(ctx context.Context, in *AddEmployeeRequest, opts ...grpc.CallOption) (*emptypb.Empty, error) {
	out := new(emptypb.Empty)
	err := c.cc.Invoke(ctx, "/employeespb.Api/AddEmployee", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *apiClient) UpdateEmployee(ctx context.Context, in *UpdateEmployeeRequest, opts ...grpc.CallOption) (*emptypb.Empty, error) {
	out := new(emptypb.Empty)
	err := c.cc.Invoke(ctx, "/employeespb.Api/UpdateEmployee", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *apiClient) DeleteEmployee(ctx context.Context, in *DeleteEmployeeRequest, opts ...grpc.CallOption) (*emptypb.Empty, error) {
	out := new(emptypb.Empty)
	err := c.cc.Invoke(ctx, "/employeespb.Api/DeleteEmployee", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *apiClient) FindOneEmployee(ctx context.Context, in *EmployeeRequest, opts ...grpc.CallOption) (*EmployeeReply, error) {
	out := new(EmployeeReply)
	err := c.cc.Invoke(ctx, "/employeespb.Api/FindOneEmployee", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *apiClient) FindManyEmployees(ctx context.Context, in *EmployeesRequest, opts ...grpc.CallOption) (*EmployeesReply, error) {
	out := new(EmployeesReply)
	err := c.cc.Invoke(ctx, "/employeespb.Api/FindManyEmployees", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// ApiServer is the server API for Api service.
// All implementations must embed UnimplementedApiServer
// for forward compatibility
type ApiServer interface {
	AddEmployee(context.Context, *AddEmployeeRequest) (*emptypb.Empty, error)
	UpdateEmployee(context.Context, *UpdateEmployeeRequest) (*emptypb.Empty, error)
	DeleteEmployee(context.Context, *DeleteEmployeeRequest) (*emptypb.Empty, error)
	FindOneEmployee(context.Context, *EmployeeRequest) (*EmployeeReply, error)
	FindManyEmployees(context.Context, *EmployeesRequest) (*EmployeesReply, error)
	mustEmbedUnimplementedApiServer()
}

// UnimplementedApiServer must be embedded to have forward compatible implementations.
type UnimplementedApiServer struct {
}

func (UnimplementedApiServer) AddEmployee(context.Context, *AddEmployeeRequest) (*emptypb.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method AddEmployee not implemented")
}
func (UnimplementedApiServer) UpdateEmployee(context.Context, *UpdateEmployeeRequest) (*emptypb.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateEmployee not implemented")
}
func (UnimplementedApiServer) DeleteEmployee(context.Context, *DeleteEmployeeRequest) (*emptypb.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteEmployee not implemented")
}
func (UnimplementedApiServer) FindOneEmployee(context.Context, *EmployeeRequest) (*EmployeeReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method FindOneEmployee not implemented")
}
func (UnimplementedApiServer) FindManyEmployees(context.Context, *EmployeesRequest) (*EmployeesReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method FindManyEmployees not implemented")
}
func (UnimplementedApiServer) mustEmbedUnimplementedApiServer() {}

// UnsafeApiServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to ApiServer will
// result in compilation errors.
type UnsafeApiServer interface {
	mustEmbedUnimplementedApiServer()
}

func RegisterApiServer(s grpc.ServiceRegistrar, srv ApiServer) {
	s.RegisterService(&Api_ServiceDesc, srv)
}

func _Api_AddEmployee_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AddEmployeeRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ApiServer).AddEmployee(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/employeespb.Api/AddEmployee",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ApiServer).AddEmployee(ctx, req.(*AddEmployeeRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Api_UpdateEmployee_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateEmployeeRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ApiServer).UpdateEmployee(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/employeespb.Api/UpdateEmployee",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ApiServer).UpdateEmployee(ctx, req.(*UpdateEmployeeRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Api_DeleteEmployee_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteEmployeeRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ApiServer).DeleteEmployee(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/employeespb.Api/DeleteEmployee",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ApiServer).DeleteEmployee(ctx, req.(*DeleteEmployeeRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Api_FindOneEmployee_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(EmployeeRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ApiServer).FindOneEmployee(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/employeespb.Api/FindOneEmployee",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ApiServer).FindOneEmployee(ctx, req.(*EmployeeRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Api_FindManyEmployees_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(EmployeesRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ApiServer).FindManyEmployees(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/employeespb.Api/FindManyEmployees",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ApiServer).FindManyEmployees(ctx, req.(*EmployeesRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// Api_ServiceDesc is the grpc.ServiceDesc for Api service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Api_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "employeespb.Api",
	HandlerType: (*ApiServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "AddEmployee",
			Handler:    _Api_AddEmployee_Handler,
		},
		{
			MethodName: "UpdateEmployee",
			Handler:    _Api_UpdateEmployee_Handler,
		},
		{
			MethodName: "DeleteEmployee",
			Handler:    _Api_DeleteEmployee_Handler,
		},
		{
			MethodName: "FindOneEmployee",
			Handler:    _Api_FindOneEmployee_Handler,
		},
		{
			MethodName: "FindManyEmployees",
			Handler:    _Api_FindManyEmployees_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "employeespb.proto",
}