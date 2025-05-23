// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.5.1
// - protoc             v5.28.3
// source: creation/creation_service.proto

package creation

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
	emptypb "google.golang.org/protobuf/types/known/emptypb"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.64.0 or later.
const _ = grpc.SupportPackageIsVersion9

const (
	CreationService_UploadCreation_FullMethodName        = "/creation.CreationService/UploadCreation"
	CreationService_GetCreation_FullMethodName           = "/creation.CreationService/GetCreation"
	CreationService_GetCreationPrivate_FullMethodName    = "/creation.CreationService/GetCreationPrivate"
	CreationService_GetSpaceCreations_FullMethodName     = "/creation.CreationService/GetSpaceCreations"
	CreationService_GetUserCreations_FullMethodName      = "/creation.CreationService/GetUserCreations"
	CreationService_SearchCreation_FullMethodName        = "/creation.CreationService/SearchCreation"
	CreationService_GetCreationList_FullMethodName       = "/creation.CreationService/GetCreationList"
	CreationService_GetPublicCreationList_FullMethodName = "/creation.CreationService/GetPublicCreationList"
	CreationService_DeleteCreation_FullMethodName        = "/creation.CreationService/DeleteCreation"
	CreationService_UpdateCreation_FullMethodName        = "/creation.CreationService/UpdateCreation"
	CreationService_PublishDraftCreation_FullMethodName  = "/creation.CreationService/PublishDraftCreation"
)

// CreationServiceClient is the client API for CreationService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type CreationServiceClient interface {
	// POST
	UploadCreation(ctx context.Context, in *UploadCreationRequest, opts ...grpc.CallOption) (*UploadCreationResponse, error)
	// GET
	GetCreation(ctx context.Context, in *GetCreationRequest, opts ...grpc.CallOption) (*GetCreationResponse, error)
	GetCreationPrivate(ctx context.Context, in *GetCreationPrivateRequest, opts ...grpc.CallOption) (*GetCreationResponse, error)
	GetSpaceCreations(ctx context.Context, in *GetSpaceCreationsRequest, opts ...grpc.CallOption) (*GetCreationListResponse, error)
	GetUserCreations(ctx context.Context, in *GetUserCreationsRequest, opts ...grpc.CallOption) (*GetCreationListResponse, error)
	SearchCreation(ctx context.Context, in *SearchCreationRequest, opts ...grpc.CallOption) (*GetCreationListResponse, error)
	GetCreationList(ctx context.Context, in *GetCreationListRequest, opts ...grpc.CallOption) (*GetCreationListResponse, error)
	GetPublicCreationList(ctx context.Context, in *GetCreationListRequest, opts ...grpc.CallOption) (*GetCreationListResponse, error)
	// DELETE
	DeleteCreation(ctx context.Context, in *DeleteCreationRequest, opts ...grpc.CallOption) (*emptypb.Empty, error)
	// UPDATE
	UpdateCreation(ctx context.Context, in *UpdateCreationRequest, opts ...grpc.CallOption) (*UpdateCreationResponse, error)
	PublishDraftCreation(ctx context.Context, in *UpdateCreationStatusRequest, opts ...grpc.CallOption) (*UpdateCreationResponse, error)
}

type creationServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewCreationServiceClient(cc grpc.ClientConnInterface) CreationServiceClient {
	return &creationServiceClient{cc}
}

func (c *creationServiceClient) UploadCreation(ctx context.Context, in *UploadCreationRequest, opts ...grpc.CallOption) (*UploadCreationResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(UploadCreationResponse)
	err := c.cc.Invoke(ctx, CreationService_UploadCreation_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *creationServiceClient) GetCreation(ctx context.Context, in *GetCreationRequest, opts ...grpc.CallOption) (*GetCreationResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(GetCreationResponse)
	err := c.cc.Invoke(ctx, CreationService_GetCreation_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *creationServiceClient) GetCreationPrivate(ctx context.Context, in *GetCreationPrivateRequest, opts ...grpc.CallOption) (*GetCreationResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(GetCreationResponse)
	err := c.cc.Invoke(ctx, CreationService_GetCreationPrivate_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *creationServiceClient) GetSpaceCreations(ctx context.Context, in *GetSpaceCreationsRequest, opts ...grpc.CallOption) (*GetCreationListResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(GetCreationListResponse)
	err := c.cc.Invoke(ctx, CreationService_GetSpaceCreations_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *creationServiceClient) GetUserCreations(ctx context.Context, in *GetUserCreationsRequest, opts ...grpc.CallOption) (*GetCreationListResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(GetCreationListResponse)
	err := c.cc.Invoke(ctx, CreationService_GetUserCreations_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *creationServiceClient) SearchCreation(ctx context.Context, in *SearchCreationRequest, opts ...grpc.CallOption) (*GetCreationListResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(GetCreationListResponse)
	err := c.cc.Invoke(ctx, CreationService_SearchCreation_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *creationServiceClient) GetCreationList(ctx context.Context, in *GetCreationListRequest, opts ...grpc.CallOption) (*GetCreationListResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(GetCreationListResponse)
	err := c.cc.Invoke(ctx, CreationService_GetCreationList_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *creationServiceClient) GetPublicCreationList(ctx context.Context, in *GetCreationListRequest, opts ...grpc.CallOption) (*GetCreationListResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(GetCreationListResponse)
	err := c.cc.Invoke(ctx, CreationService_GetPublicCreationList_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *creationServiceClient) DeleteCreation(ctx context.Context, in *DeleteCreationRequest, opts ...grpc.CallOption) (*emptypb.Empty, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(emptypb.Empty)
	err := c.cc.Invoke(ctx, CreationService_DeleteCreation_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *creationServiceClient) UpdateCreation(ctx context.Context, in *UpdateCreationRequest, opts ...grpc.CallOption) (*UpdateCreationResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(UpdateCreationResponse)
	err := c.cc.Invoke(ctx, CreationService_UpdateCreation_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *creationServiceClient) PublishDraftCreation(ctx context.Context, in *UpdateCreationStatusRequest, opts ...grpc.CallOption) (*UpdateCreationResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(UpdateCreationResponse)
	err := c.cc.Invoke(ctx, CreationService_PublishDraftCreation_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// CreationServiceServer is the server API for CreationService service.
// All implementations must embed UnimplementedCreationServiceServer
// for forward compatibility.
type CreationServiceServer interface {
	// POST
	UploadCreation(context.Context, *UploadCreationRequest) (*UploadCreationResponse, error)
	// GET
	GetCreation(context.Context, *GetCreationRequest) (*GetCreationResponse, error)
	GetCreationPrivate(context.Context, *GetCreationPrivateRequest) (*GetCreationResponse, error)
	GetSpaceCreations(context.Context, *GetSpaceCreationsRequest) (*GetCreationListResponse, error)
	GetUserCreations(context.Context, *GetUserCreationsRequest) (*GetCreationListResponse, error)
	SearchCreation(context.Context, *SearchCreationRequest) (*GetCreationListResponse, error)
	GetCreationList(context.Context, *GetCreationListRequest) (*GetCreationListResponse, error)
	GetPublicCreationList(context.Context, *GetCreationListRequest) (*GetCreationListResponse, error)
	// DELETE
	DeleteCreation(context.Context, *DeleteCreationRequest) (*emptypb.Empty, error)
	// UPDATE
	UpdateCreation(context.Context, *UpdateCreationRequest) (*UpdateCreationResponse, error)
	PublishDraftCreation(context.Context, *UpdateCreationStatusRequest) (*UpdateCreationResponse, error)
	mustEmbedUnimplementedCreationServiceServer()
}

// UnimplementedCreationServiceServer must be embedded to have
// forward compatible implementations.
//
// NOTE: this should be embedded by value instead of pointer to avoid a nil
// pointer dereference when methods are called.
type UnimplementedCreationServiceServer struct{}

func (UnimplementedCreationServiceServer) UploadCreation(context.Context, *UploadCreationRequest) (*UploadCreationResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UploadCreation not implemented")
}
func (UnimplementedCreationServiceServer) GetCreation(context.Context, *GetCreationRequest) (*GetCreationResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetCreation not implemented")
}
func (UnimplementedCreationServiceServer) GetCreationPrivate(context.Context, *GetCreationPrivateRequest) (*GetCreationResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetCreationPrivate not implemented")
}
func (UnimplementedCreationServiceServer) GetSpaceCreations(context.Context, *GetSpaceCreationsRequest) (*GetCreationListResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetSpaceCreations not implemented")
}
func (UnimplementedCreationServiceServer) GetUserCreations(context.Context, *GetUserCreationsRequest) (*GetCreationListResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetUserCreations not implemented")
}
func (UnimplementedCreationServiceServer) SearchCreation(context.Context, *SearchCreationRequest) (*GetCreationListResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SearchCreation not implemented")
}
func (UnimplementedCreationServiceServer) GetCreationList(context.Context, *GetCreationListRequest) (*GetCreationListResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetCreationList not implemented")
}
func (UnimplementedCreationServiceServer) GetPublicCreationList(context.Context, *GetCreationListRequest) (*GetCreationListResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetPublicCreationList not implemented")
}
func (UnimplementedCreationServiceServer) DeleteCreation(context.Context, *DeleteCreationRequest) (*emptypb.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteCreation not implemented")
}
func (UnimplementedCreationServiceServer) UpdateCreation(context.Context, *UpdateCreationRequest) (*UpdateCreationResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateCreation not implemented")
}
func (UnimplementedCreationServiceServer) PublishDraftCreation(context.Context, *UpdateCreationStatusRequest) (*UpdateCreationResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method PublishDraftCreation not implemented")
}
func (UnimplementedCreationServiceServer) mustEmbedUnimplementedCreationServiceServer() {}
func (UnimplementedCreationServiceServer) testEmbeddedByValue()                         {}

// UnsafeCreationServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to CreationServiceServer will
// result in compilation errors.
type UnsafeCreationServiceServer interface {
	mustEmbedUnimplementedCreationServiceServer()
}

func RegisterCreationServiceServer(s grpc.ServiceRegistrar, srv CreationServiceServer) {
	// If the following call pancis, it indicates UnimplementedCreationServiceServer was
	// embedded by pointer and is nil.  This will cause panics if an
	// unimplemented method is ever invoked, so we test this at initialization
	// time to prevent it from happening at runtime later due to I/O.
	if t, ok := srv.(interface{ testEmbeddedByValue() }); ok {
		t.testEmbeddedByValue()
	}
	s.RegisterService(&CreationService_ServiceDesc, srv)
}

func _CreationService_UploadCreation_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UploadCreationRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CreationServiceServer).UploadCreation(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: CreationService_UploadCreation_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CreationServiceServer).UploadCreation(ctx, req.(*UploadCreationRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _CreationService_GetCreation_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetCreationRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CreationServiceServer).GetCreation(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: CreationService_GetCreation_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CreationServiceServer).GetCreation(ctx, req.(*GetCreationRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _CreationService_GetCreationPrivate_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetCreationPrivateRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CreationServiceServer).GetCreationPrivate(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: CreationService_GetCreationPrivate_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CreationServiceServer).GetCreationPrivate(ctx, req.(*GetCreationPrivateRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _CreationService_GetSpaceCreations_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetSpaceCreationsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CreationServiceServer).GetSpaceCreations(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: CreationService_GetSpaceCreations_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CreationServiceServer).GetSpaceCreations(ctx, req.(*GetSpaceCreationsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _CreationService_GetUserCreations_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetUserCreationsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CreationServiceServer).GetUserCreations(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: CreationService_GetUserCreations_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CreationServiceServer).GetUserCreations(ctx, req.(*GetUserCreationsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _CreationService_SearchCreation_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SearchCreationRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CreationServiceServer).SearchCreation(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: CreationService_SearchCreation_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CreationServiceServer).SearchCreation(ctx, req.(*SearchCreationRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _CreationService_GetCreationList_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetCreationListRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CreationServiceServer).GetCreationList(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: CreationService_GetCreationList_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CreationServiceServer).GetCreationList(ctx, req.(*GetCreationListRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _CreationService_GetPublicCreationList_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetCreationListRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CreationServiceServer).GetPublicCreationList(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: CreationService_GetPublicCreationList_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CreationServiceServer).GetPublicCreationList(ctx, req.(*GetCreationListRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _CreationService_DeleteCreation_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteCreationRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CreationServiceServer).DeleteCreation(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: CreationService_DeleteCreation_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CreationServiceServer).DeleteCreation(ctx, req.(*DeleteCreationRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _CreationService_UpdateCreation_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateCreationRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CreationServiceServer).UpdateCreation(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: CreationService_UpdateCreation_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CreationServiceServer).UpdateCreation(ctx, req.(*UpdateCreationRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _CreationService_PublishDraftCreation_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateCreationStatusRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CreationServiceServer).PublishDraftCreation(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: CreationService_PublishDraftCreation_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CreationServiceServer).PublishDraftCreation(ctx, req.(*UpdateCreationStatusRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// CreationService_ServiceDesc is the grpc.ServiceDesc for CreationService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var CreationService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "creation.CreationService",
	HandlerType: (*CreationServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "UploadCreation",
			Handler:    _CreationService_UploadCreation_Handler,
		},
		{
			MethodName: "GetCreation",
			Handler:    _CreationService_GetCreation_Handler,
		},
		{
			MethodName: "GetCreationPrivate",
			Handler:    _CreationService_GetCreationPrivate_Handler,
		},
		{
			MethodName: "GetSpaceCreations",
			Handler:    _CreationService_GetSpaceCreations_Handler,
		},
		{
			MethodName: "GetUserCreations",
			Handler:    _CreationService_GetUserCreations_Handler,
		},
		{
			MethodName: "SearchCreation",
			Handler:    _CreationService_SearchCreation_Handler,
		},
		{
			MethodName: "GetCreationList",
			Handler:    _CreationService_GetCreationList_Handler,
		},
		{
			MethodName: "GetPublicCreationList",
			Handler:    _CreationService_GetPublicCreationList_Handler,
		},
		{
			MethodName: "DeleteCreation",
			Handler:    _CreationService_DeleteCreation_Handler,
		},
		{
			MethodName: "UpdateCreation",
			Handler:    _CreationService_UpdateCreation_Handler,
		},
		{
			MethodName: "PublishDraftCreation",
			Handler:    _CreationService_PublishDraftCreation_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "creation/creation_service.proto",
}
