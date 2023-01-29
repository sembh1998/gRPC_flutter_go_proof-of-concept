// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.6.1
// source: service.proto

package pb

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

// ModuleServiceClient is the client API for ModuleService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type ModuleServiceClient interface {
	GetStores(ctx context.Context, in *StoreRequest, opts ...grpc.CallOption) (*StoresResponse, error)
	GetStoreById(ctx context.Context, in *StoreByIdRequest, opts ...grpc.CallOption) (*StoreResponse, error)
	GetStoreByName(ctx context.Context, in *StoreByNameRequest, opts ...grpc.CallOption) (*StoreResponse, error)
	GetProducts(ctx context.Context, in *ProductRequest, opts ...grpc.CallOption) (*ProductsResponse, error)
	GetProductById(ctx context.Context, in *ProductByIdRequest, opts ...grpc.CallOption) (*ProductResponse, error)
	GetMembers(ctx context.Context, in *MemberRequest, opts ...grpc.CallOption) (*MembersResponse, error)
	GetMemberById(ctx context.Context, in *MemberByIdRequest, opts ...grpc.CallOption) (*MemberResponse, error)
	CreateMember(ctx context.Context, in *Member, opts ...grpc.CallOption) (*MemberResponse, error)
	GetEstablishmentsPerStore(ctx context.Context, in *EstablishmentsByStoreIdRequest, opts ...grpc.CallOption) (*EstablishmentsResponse, error)
	GetEstablishmentById(ctx context.Context, in *EstablishmentByIdRequest, opts ...grpc.CallOption) (*EstablishmentResponse, error)
}

type moduleServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewModuleServiceClient(cc grpc.ClientConnInterface) ModuleServiceClient {
	return &moduleServiceClient{cc}
}

func (c *moduleServiceClient) GetStores(ctx context.Context, in *StoreRequest, opts ...grpc.CallOption) (*StoresResponse, error) {
	out := new(StoresResponse)
	err := c.cc.Invoke(ctx, "/ModuleService/getStores", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *moduleServiceClient) GetStoreById(ctx context.Context, in *StoreByIdRequest, opts ...grpc.CallOption) (*StoreResponse, error) {
	out := new(StoreResponse)
	err := c.cc.Invoke(ctx, "/ModuleService/getStoreById", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *moduleServiceClient) GetStoreByName(ctx context.Context, in *StoreByNameRequest, opts ...grpc.CallOption) (*StoreResponse, error) {
	out := new(StoreResponse)
	err := c.cc.Invoke(ctx, "/ModuleService/getStoreByName", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *moduleServiceClient) GetProducts(ctx context.Context, in *ProductRequest, opts ...grpc.CallOption) (*ProductsResponse, error) {
	out := new(ProductsResponse)
	err := c.cc.Invoke(ctx, "/ModuleService/getProducts", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *moduleServiceClient) GetProductById(ctx context.Context, in *ProductByIdRequest, opts ...grpc.CallOption) (*ProductResponse, error) {
	out := new(ProductResponse)
	err := c.cc.Invoke(ctx, "/ModuleService/getProductById", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *moduleServiceClient) GetMembers(ctx context.Context, in *MemberRequest, opts ...grpc.CallOption) (*MembersResponse, error) {
	out := new(MembersResponse)
	err := c.cc.Invoke(ctx, "/ModuleService/getMembers", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *moduleServiceClient) GetMemberById(ctx context.Context, in *MemberByIdRequest, opts ...grpc.CallOption) (*MemberResponse, error) {
	out := new(MemberResponse)
	err := c.cc.Invoke(ctx, "/ModuleService/getMemberById", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *moduleServiceClient) CreateMember(ctx context.Context, in *Member, opts ...grpc.CallOption) (*MemberResponse, error) {
	out := new(MemberResponse)
	err := c.cc.Invoke(ctx, "/ModuleService/createMember", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *moduleServiceClient) GetEstablishmentsPerStore(ctx context.Context, in *EstablishmentsByStoreIdRequest, opts ...grpc.CallOption) (*EstablishmentsResponse, error) {
	out := new(EstablishmentsResponse)
	err := c.cc.Invoke(ctx, "/ModuleService/getEstablishmentsPerStore", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *moduleServiceClient) GetEstablishmentById(ctx context.Context, in *EstablishmentByIdRequest, opts ...grpc.CallOption) (*EstablishmentResponse, error) {
	out := new(EstablishmentResponse)
	err := c.cc.Invoke(ctx, "/ModuleService/getEstablishmentById", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// ModuleServiceServer is the server API for ModuleService service.
// All implementations must embed UnimplementedModuleServiceServer
// for forward compatibility
type ModuleServiceServer interface {
	GetStores(context.Context, *StoreRequest) (*StoresResponse, error)
	GetStoreById(context.Context, *StoreByIdRequest) (*StoreResponse, error)
	GetStoreByName(context.Context, *StoreByNameRequest) (*StoreResponse, error)
	GetProducts(context.Context, *ProductRequest) (*ProductsResponse, error)
	GetProductById(context.Context, *ProductByIdRequest) (*ProductResponse, error)
	GetMembers(context.Context, *MemberRequest) (*MembersResponse, error)
	GetMemberById(context.Context, *MemberByIdRequest) (*MemberResponse, error)
	CreateMember(context.Context, *Member) (*MemberResponse, error)
	GetEstablishmentsPerStore(context.Context, *EstablishmentsByStoreIdRequest) (*EstablishmentsResponse, error)
	GetEstablishmentById(context.Context, *EstablishmentByIdRequest) (*EstablishmentResponse, error)
	mustEmbedUnimplementedModuleServiceServer()
}

// UnimplementedModuleServiceServer must be embedded to have forward compatible implementations.
type UnimplementedModuleServiceServer struct {
}

func (UnimplementedModuleServiceServer) GetStores(context.Context, *StoreRequest) (*StoresResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetStores not implemented")
}
func (UnimplementedModuleServiceServer) GetStoreById(context.Context, *StoreByIdRequest) (*StoreResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetStoreById not implemented")
}
func (UnimplementedModuleServiceServer) GetStoreByName(context.Context, *StoreByNameRequest) (*StoreResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetStoreByName not implemented")
}
func (UnimplementedModuleServiceServer) GetProducts(context.Context, *ProductRequest) (*ProductsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetProducts not implemented")
}
func (UnimplementedModuleServiceServer) GetProductById(context.Context, *ProductByIdRequest) (*ProductResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetProductById not implemented")
}
func (UnimplementedModuleServiceServer) GetMembers(context.Context, *MemberRequest) (*MembersResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetMembers not implemented")
}
func (UnimplementedModuleServiceServer) GetMemberById(context.Context, *MemberByIdRequest) (*MemberResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetMemberById not implemented")
}
func (UnimplementedModuleServiceServer) CreateMember(context.Context, *Member) (*MemberResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateMember not implemented")
}
func (UnimplementedModuleServiceServer) GetEstablishmentsPerStore(context.Context, *EstablishmentsByStoreIdRequest) (*EstablishmentsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetEstablishmentsPerStore not implemented")
}
func (UnimplementedModuleServiceServer) GetEstablishmentById(context.Context, *EstablishmentByIdRequest) (*EstablishmentResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetEstablishmentById not implemented")
}
func (UnimplementedModuleServiceServer) mustEmbedUnimplementedModuleServiceServer() {}

// UnsafeModuleServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to ModuleServiceServer will
// result in compilation errors.
type UnsafeModuleServiceServer interface {
	mustEmbedUnimplementedModuleServiceServer()
}

func RegisterModuleServiceServer(s grpc.ServiceRegistrar, srv ModuleServiceServer) {
	s.RegisterService(&ModuleService_ServiceDesc, srv)
}

func _ModuleService_GetStores_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(StoreRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ModuleServiceServer).GetStores(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/ModuleService/getStores",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ModuleServiceServer).GetStores(ctx, req.(*StoreRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ModuleService_GetStoreById_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(StoreByIdRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ModuleServiceServer).GetStoreById(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/ModuleService/getStoreById",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ModuleServiceServer).GetStoreById(ctx, req.(*StoreByIdRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ModuleService_GetStoreByName_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(StoreByNameRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ModuleServiceServer).GetStoreByName(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/ModuleService/getStoreByName",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ModuleServiceServer).GetStoreByName(ctx, req.(*StoreByNameRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ModuleService_GetProducts_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ProductRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ModuleServiceServer).GetProducts(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/ModuleService/getProducts",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ModuleServiceServer).GetProducts(ctx, req.(*ProductRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ModuleService_GetProductById_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ProductByIdRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ModuleServiceServer).GetProductById(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/ModuleService/getProductById",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ModuleServiceServer).GetProductById(ctx, req.(*ProductByIdRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ModuleService_GetMembers_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(MemberRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ModuleServiceServer).GetMembers(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/ModuleService/getMembers",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ModuleServiceServer).GetMembers(ctx, req.(*MemberRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ModuleService_GetMemberById_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(MemberByIdRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ModuleServiceServer).GetMemberById(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/ModuleService/getMemberById",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ModuleServiceServer).GetMemberById(ctx, req.(*MemberByIdRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ModuleService_CreateMember_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Member)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ModuleServiceServer).CreateMember(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/ModuleService/createMember",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ModuleServiceServer).CreateMember(ctx, req.(*Member))
	}
	return interceptor(ctx, in, info, handler)
}

func _ModuleService_GetEstablishmentsPerStore_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(EstablishmentsByStoreIdRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ModuleServiceServer).GetEstablishmentsPerStore(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/ModuleService/getEstablishmentsPerStore",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ModuleServiceServer).GetEstablishmentsPerStore(ctx, req.(*EstablishmentsByStoreIdRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ModuleService_GetEstablishmentById_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(EstablishmentByIdRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ModuleServiceServer).GetEstablishmentById(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/ModuleService/getEstablishmentById",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ModuleServiceServer).GetEstablishmentById(ctx, req.(*EstablishmentByIdRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// ModuleService_ServiceDesc is the grpc.ServiceDesc for ModuleService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var ModuleService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "ModuleService",
	HandlerType: (*ModuleServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "getStores",
			Handler:    _ModuleService_GetStores_Handler,
		},
		{
			MethodName: "getStoreById",
			Handler:    _ModuleService_GetStoreById_Handler,
		},
		{
			MethodName: "getStoreByName",
			Handler:    _ModuleService_GetStoreByName_Handler,
		},
		{
			MethodName: "getProducts",
			Handler:    _ModuleService_GetProducts_Handler,
		},
		{
			MethodName: "getProductById",
			Handler:    _ModuleService_GetProductById_Handler,
		},
		{
			MethodName: "getMembers",
			Handler:    _ModuleService_GetMembers_Handler,
		},
		{
			MethodName: "getMemberById",
			Handler:    _ModuleService_GetMemberById_Handler,
		},
		{
			MethodName: "createMember",
			Handler:    _ModuleService_CreateMember_Handler,
		},
		{
			MethodName: "getEstablishmentsPerStore",
			Handler:    _ModuleService_GetEstablishmentsPerStore_Handler,
		},
		{
			MethodName: "getEstablishmentById",
			Handler:    _ModuleService_GetEstablishmentById_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "service.proto",
}
