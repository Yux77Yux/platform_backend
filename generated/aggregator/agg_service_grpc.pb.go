// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.5.1
// - protoc             v5.28.3
// source: aggregator/agg_service.proto

package aggregator

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
	AggregatorService_Search_FullMethodName                = "/aggregator.AggregatorService/Search"
	AggregatorService_Login_FullMethodName                 = "/aggregator.AggregatorService/Login"
	AggregatorService_WatchCreation_FullMethodName         = "/aggregator.AggregatorService/WatchCreation"
	AggregatorService_SimilarCreations_FullMethodName      = "/aggregator.AggregatorService/SimilarCreations"
	AggregatorService_InitialComments_FullMethodName       = "/aggregator.AggregatorService/InitialComments"
	AggregatorService_GetTopComments_FullMethodName        = "/aggregator.AggregatorService/GetTopComments"
	AggregatorService_GetSecondComments_FullMethodName     = "/aggregator.AggregatorService/GetSecondComments"
	AggregatorService_GetUserReviews_FullMethodName        = "/aggregator.AggregatorService/GetUserReviews"
	AggregatorService_GetCreationReviews_FullMethodName    = "/aggregator.AggregatorService/GetCreationReviews"
	AggregatorService_GetCommentReviews_FullMethodName     = "/aggregator.AggregatorService/GetCommentReviews"
	AggregatorService_GetNewUserReviews_FullMethodName     = "/aggregator.AggregatorService/GetNewUserReviews"
	AggregatorService_GetNewCreationReviews_FullMethodName = "/aggregator.AggregatorService/GetNewCreationReviews"
	AggregatorService_GetNewCommentReviews_FullMethodName  = "/aggregator.AggregatorService/GetNewCommentReviews"
	AggregatorService_HomePage_FullMethodName              = "/aggregator.AggregatorService/HomePage"
	AggregatorService_Collections_FullMethodName           = "/aggregator.AggregatorService/Collections"
	AggregatorService_History_FullMethodName               = "/aggregator.AggregatorService/History"
)

// AggregatorServiceClient is the client API for AggregatorService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type AggregatorServiceClient interface {
	Search(ctx context.Context, in *SearchCreationsRequest, opts ...grpc.CallOption) (*SearchCreationsResponse, error)
	// User OK
	Login(ctx context.Context, in *LoginRequest, opts ...grpc.CallOption) (*LoginResponse, error)
	// WatchCreation OK
	WatchCreation(ctx context.Context, in *WatchCreationRequest, opts ...grpc.CallOption) (*WatchCreationResponse, error)
	// 相似视频 OK
	SimilarCreations(ctx context.Context, in *SimilarCreationsRequest, opts ...grpc.CallOption) (*GetCardsResponse, error)
	// Comment OK
	InitialComments(ctx context.Context, in *InitialCommentsRequest, opts ...grpc.CallOption) (*InitialCommentsResponse, error)
	GetTopComments(ctx context.Context, in *GetTopCommentsRequest, opts ...grpc.CallOption) (*GetTopCommentsResponse, error)
	GetSecondComments(ctx context.Context, in *GetSecondCommentsRequest, opts ...grpc.CallOption) (*GetSecondCommentsResponse, error)
	// Review
	GetUserReviews(ctx context.Context, in *GetReviewsRequest, opts ...grpc.CallOption) (*GetUserReviewsResponse, error)
	GetCreationReviews(ctx context.Context, in *GetReviewsRequest, opts ...grpc.CallOption) (*GetCreationReviewsResponse, error)
	GetCommentReviews(ctx context.Context, in *GetReviewsRequest, opts ...grpc.CallOption) (*GetCommentReviewsResponse, error)
	GetNewUserReviews(ctx context.Context, in *GetNewReviewsRequest, opts ...grpc.CallOption) (*GetUserReviewsResponse, error)
	GetNewCreationReviews(ctx context.Context, in *GetNewReviewsRequest, opts ...grpc.CallOption) (*GetCreationReviewsResponse, error)
	GetNewCommentReviews(ctx context.Context, in *GetNewReviewsRequest, opts ...grpc.CallOption) (*GetCommentReviewsResponse, error)
	// 主页
	HomePage(ctx context.Context, in *HomeRequest, opts ...grpc.CallOption) (*GetCardsResponse, error)
	// 收藏夹
	Collections(ctx context.Context, in *CollectionsRequest, opts ...grpc.CallOption) (*GetCardsResponse, error)
	// 历史
	History(ctx context.Context, in *HistoryRequest, opts ...grpc.CallOption) (*GetCardsResponse, error)
}

type aggregatorServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewAggregatorServiceClient(cc grpc.ClientConnInterface) AggregatorServiceClient {
	return &aggregatorServiceClient{cc}
}

func (c *aggregatorServiceClient) Search(ctx context.Context, in *SearchCreationsRequest, opts ...grpc.CallOption) (*SearchCreationsResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(SearchCreationsResponse)
	err := c.cc.Invoke(ctx, AggregatorService_Search_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *aggregatorServiceClient) Login(ctx context.Context, in *LoginRequest, opts ...grpc.CallOption) (*LoginResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(LoginResponse)
	err := c.cc.Invoke(ctx, AggregatorService_Login_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *aggregatorServiceClient) WatchCreation(ctx context.Context, in *WatchCreationRequest, opts ...grpc.CallOption) (*WatchCreationResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(WatchCreationResponse)
	err := c.cc.Invoke(ctx, AggregatorService_WatchCreation_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *aggregatorServiceClient) SimilarCreations(ctx context.Context, in *SimilarCreationsRequest, opts ...grpc.CallOption) (*GetCardsResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(GetCardsResponse)
	err := c.cc.Invoke(ctx, AggregatorService_SimilarCreations_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *aggregatorServiceClient) InitialComments(ctx context.Context, in *InitialCommentsRequest, opts ...grpc.CallOption) (*InitialCommentsResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(InitialCommentsResponse)
	err := c.cc.Invoke(ctx, AggregatorService_InitialComments_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *aggregatorServiceClient) GetTopComments(ctx context.Context, in *GetTopCommentsRequest, opts ...grpc.CallOption) (*GetTopCommentsResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(GetTopCommentsResponse)
	err := c.cc.Invoke(ctx, AggregatorService_GetTopComments_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *aggregatorServiceClient) GetSecondComments(ctx context.Context, in *GetSecondCommentsRequest, opts ...grpc.CallOption) (*GetSecondCommentsResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(GetSecondCommentsResponse)
	err := c.cc.Invoke(ctx, AggregatorService_GetSecondComments_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *aggregatorServiceClient) GetUserReviews(ctx context.Context, in *GetReviewsRequest, opts ...grpc.CallOption) (*GetUserReviewsResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(GetUserReviewsResponse)
	err := c.cc.Invoke(ctx, AggregatorService_GetUserReviews_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *aggregatorServiceClient) GetCreationReviews(ctx context.Context, in *GetReviewsRequest, opts ...grpc.CallOption) (*GetCreationReviewsResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(GetCreationReviewsResponse)
	err := c.cc.Invoke(ctx, AggregatorService_GetCreationReviews_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *aggregatorServiceClient) GetCommentReviews(ctx context.Context, in *GetReviewsRequest, opts ...grpc.CallOption) (*GetCommentReviewsResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(GetCommentReviewsResponse)
	err := c.cc.Invoke(ctx, AggregatorService_GetCommentReviews_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *aggregatorServiceClient) GetNewUserReviews(ctx context.Context, in *GetNewReviewsRequest, opts ...grpc.CallOption) (*GetUserReviewsResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(GetUserReviewsResponse)
	err := c.cc.Invoke(ctx, AggregatorService_GetNewUserReviews_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *aggregatorServiceClient) GetNewCreationReviews(ctx context.Context, in *GetNewReviewsRequest, opts ...grpc.CallOption) (*GetCreationReviewsResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(GetCreationReviewsResponse)
	err := c.cc.Invoke(ctx, AggregatorService_GetNewCreationReviews_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *aggregatorServiceClient) GetNewCommentReviews(ctx context.Context, in *GetNewReviewsRequest, opts ...grpc.CallOption) (*GetCommentReviewsResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(GetCommentReviewsResponse)
	err := c.cc.Invoke(ctx, AggregatorService_GetNewCommentReviews_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *aggregatorServiceClient) HomePage(ctx context.Context, in *HomeRequest, opts ...grpc.CallOption) (*GetCardsResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(GetCardsResponse)
	err := c.cc.Invoke(ctx, AggregatorService_HomePage_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *aggregatorServiceClient) Collections(ctx context.Context, in *CollectionsRequest, opts ...grpc.CallOption) (*GetCardsResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(GetCardsResponse)
	err := c.cc.Invoke(ctx, AggregatorService_Collections_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *aggregatorServiceClient) History(ctx context.Context, in *HistoryRequest, opts ...grpc.CallOption) (*GetCardsResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(GetCardsResponse)
	err := c.cc.Invoke(ctx, AggregatorService_History_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// AggregatorServiceServer is the server API for AggregatorService service.
// All implementations must embed UnimplementedAggregatorServiceServer
// for forward compatibility.
type AggregatorServiceServer interface {
	Search(context.Context, *SearchCreationsRequest) (*SearchCreationsResponse, error)
	// User OK
	Login(context.Context, *LoginRequest) (*LoginResponse, error)
	// WatchCreation OK
	WatchCreation(context.Context, *WatchCreationRequest) (*WatchCreationResponse, error)
	// 相似视频 OK
	SimilarCreations(context.Context, *SimilarCreationsRequest) (*GetCardsResponse, error)
	// Comment OK
	InitialComments(context.Context, *InitialCommentsRequest) (*InitialCommentsResponse, error)
	GetTopComments(context.Context, *GetTopCommentsRequest) (*GetTopCommentsResponse, error)
	GetSecondComments(context.Context, *GetSecondCommentsRequest) (*GetSecondCommentsResponse, error)
	// Review
	GetUserReviews(context.Context, *GetReviewsRequest) (*GetUserReviewsResponse, error)
	GetCreationReviews(context.Context, *GetReviewsRequest) (*GetCreationReviewsResponse, error)
	GetCommentReviews(context.Context, *GetReviewsRequest) (*GetCommentReviewsResponse, error)
	GetNewUserReviews(context.Context, *GetNewReviewsRequest) (*GetUserReviewsResponse, error)
	GetNewCreationReviews(context.Context, *GetNewReviewsRequest) (*GetCreationReviewsResponse, error)
	GetNewCommentReviews(context.Context, *GetNewReviewsRequest) (*GetCommentReviewsResponse, error)
	// 主页
	HomePage(context.Context, *HomeRequest) (*GetCardsResponse, error)
	// 收藏夹
	Collections(context.Context, *CollectionsRequest) (*GetCardsResponse, error)
	// 历史
	History(context.Context, *HistoryRequest) (*GetCardsResponse, error)
	mustEmbedUnimplementedAggregatorServiceServer()
}

// UnimplementedAggregatorServiceServer must be embedded to have
// forward compatible implementations.
//
// NOTE: this should be embedded by value instead of pointer to avoid a nil
// pointer dereference when methods are called.
type UnimplementedAggregatorServiceServer struct{}

func (UnimplementedAggregatorServiceServer) Search(context.Context, *SearchCreationsRequest) (*SearchCreationsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Search not implemented")
}
func (UnimplementedAggregatorServiceServer) Login(context.Context, *LoginRequest) (*LoginResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Login not implemented")
}
func (UnimplementedAggregatorServiceServer) WatchCreation(context.Context, *WatchCreationRequest) (*WatchCreationResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method WatchCreation not implemented")
}
func (UnimplementedAggregatorServiceServer) SimilarCreations(context.Context, *SimilarCreationsRequest) (*GetCardsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SimilarCreations not implemented")
}
func (UnimplementedAggregatorServiceServer) InitialComments(context.Context, *InitialCommentsRequest) (*InitialCommentsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method InitialComments not implemented")
}
func (UnimplementedAggregatorServiceServer) GetTopComments(context.Context, *GetTopCommentsRequest) (*GetTopCommentsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetTopComments not implemented")
}
func (UnimplementedAggregatorServiceServer) GetSecondComments(context.Context, *GetSecondCommentsRequest) (*GetSecondCommentsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetSecondComments not implemented")
}
func (UnimplementedAggregatorServiceServer) GetUserReviews(context.Context, *GetReviewsRequest) (*GetUserReviewsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetUserReviews not implemented")
}
func (UnimplementedAggregatorServiceServer) GetCreationReviews(context.Context, *GetReviewsRequest) (*GetCreationReviewsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetCreationReviews not implemented")
}
func (UnimplementedAggregatorServiceServer) GetCommentReviews(context.Context, *GetReviewsRequest) (*GetCommentReviewsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetCommentReviews not implemented")
}
func (UnimplementedAggregatorServiceServer) GetNewUserReviews(context.Context, *GetNewReviewsRequest) (*GetUserReviewsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetNewUserReviews not implemented")
}
func (UnimplementedAggregatorServiceServer) GetNewCreationReviews(context.Context, *GetNewReviewsRequest) (*GetCreationReviewsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetNewCreationReviews not implemented")
}
func (UnimplementedAggregatorServiceServer) GetNewCommentReviews(context.Context, *GetNewReviewsRequest) (*GetCommentReviewsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetNewCommentReviews not implemented")
}
func (UnimplementedAggregatorServiceServer) HomePage(context.Context, *HomeRequest) (*GetCardsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method HomePage not implemented")
}
func (UnimplementedAggregatorServiceServer) Collections(context.Context, *CollectionsRequest) (*GetCardsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Collections not implemented")
}
func (UnimplementedAggregatorServiceServer) History(context.Context, *HistoryRequest) (*GetCardsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method History not implemented")
}
func (UnimplementedAggregatorServiceServer) mustEmbedUnimplementedAggregatorServiceServer() {}
func (UnimplementedAggregatorServiceServer) testEmbeddedByValue()                           {}

// UnsafeAggregatorServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to AggregatorServiceServer will
// result in compilation errors.
type UnsafeAggregatorServiceServer interface {
	mustEmbedUnimplementedAggregatorServiceServer()
}

func RegisterAggregatorServiceServer(s grpc.ServiceRegistrar, srv AggregatorServiceServer) {
	// If the following call pancis, it indicates UnimplementedAggregatorServiceServer was
	// embedded by pointer and is nil.  This will cause panics if an
	// unimplemented method is ever invoked, so we test this at initialization
	// time to prevent it from happening at runtime later due to I/O.
	if t, ok := srv.(interface{ testEmbeddedByValue() }); ok {
		t.testEmbeddedByValue()
	}
	s.RegisterService(&AggregatorService_ServiceDesc, srv)
}

func _AggregatorService_Search_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SearchCreationsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AggregatorServiceServer).Search(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: AggregatorService_Search_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AggregatorServiceServer).Search(ctx, req.(*SearchCreationsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _AggregatorService_Login_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(LoginRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AggregatorServiceServer).Login(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: AggregatorService_Login_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AggregatorServiceServer).Login(ctx, req.(*LoginRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _AggregatorService_WatchCreation_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(WatchCreationRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AggregatorServiceServer).WatchCreation(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: AggregatorService_WatchCreation_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AggregatorServiceServer).WatchCreation(ctx, req.(*WatchCreationRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _AggregatorService_SimilarCreations_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SimilarCreationsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AggregatorServiceServer).SimilarCreations(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: AggregatorService_SimilarCreations_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AggregatorServiceServer).SimilarCreations(ctx, req.(*SimilarCreationsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _AggregatorService_InitialComments_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(InitialCommentsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AggregatorServiceServer).InitialComments(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: AggregatorService_InitialComments_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AggregatorServiceServer).InitialComments(ctx, req.(*InitialCommentsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _AggregatorService_GetTopComments_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetTopCommentsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AggregatorServiceServer).GetTopComments(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: AggregatorService_GetTopComments_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AggregatorServiceServer).GetTopComments(ctx, req.(*GetTopCommentsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _AggregatorService_GetSecondComments_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetSecondCommentsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AggregatorServiceServer).GetSecondComments(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: AggregatorService_GetSecondComments_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AggregatorServiceServer).GetSecondComments(ctx, req.(*GetSecondCommentsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _AggregatorService_GetUserReviews_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetReviewsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AggregatorServiceServer).GetUserReviews(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: AggregatorService_GetUserReviews_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AggregatorServiceServer).GetUserReviews(ctx, req.(*GetReviewsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _AggregatorService_GetCreationReviews_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetReviewsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AggregatorServiceServer).GetCreationReviews(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: AggregatorService_GetCreationReviews_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AggregatorServiceServer).GetCreationReviews(ctx, req.(*GetReviewsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _AggregatorService_GetCommentReviews_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetReviewsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AggregatorServiceServer).GetCommentReviews(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: AggregatorService_GetCommentReviews_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AggregatorServiceServer).GetCommentReviews(ctx, req.(*GetReviewsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _AggregatorService_GetNewUserReviews_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetNewReviewsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AggregatorServiceServer).GetNewUserReviews(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: AggregatorService_GetNewUserReviews_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AggregatorServiceServer).GetNewUserReviews(ctx, req.(*GetNewReviewsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _AggregatorService_GetNewCreationReviews_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetNewReviewsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AggregatorServiceServer).GetNewCreationReviews(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: AggregatorService_GetNewCreationReviews_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AggregatorServiceServer).GetNewCreationReviews(ctx, req.(*GetNewReviewsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _AggregatorService_GetNewCommentReviews_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetNewReviewsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AggregatorServiceServer).GetNewCommentReviews(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: AggregatorService_GetNewCommentReviews_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AggregatorServiceServer).GetNewCommentReviews(ctx, req.(*GetNewReviewsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _AggregatorService_HomePage_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(HomeRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AggregatorServiceServer).HomePage(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: AggregatorService_HomePage_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AggregatorServiceServer).HomePage(ctx, req.(*HomeRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _AggregatorService_Collections_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CollectionsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AggregatorServiceServer).Collections(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: AggregatorService_Collections_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AggregatorServiceServer).Collections(ctx, req.(*CollectionsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _AggregatorService_History_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(HistoryRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AggregatorServiceServer).History(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: AggregatorService_History_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AggregatorServiceServer).History(ctx, req.(*HistoryRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// AggregatorService_ServiceDesc is the grpc.ServiceDesc for AggregatorService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var AggregatorService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "aggregator.AggregatorService",
	HandlerType: (*AggregatorServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Search",
			Handler:    _AggregatorService_Search_Handler,
		},
		{
			MethodName: "Login",
			Handler:    _AggregatorService_Login_Handler,
		},
		{
			MethodName: "WatchCreation",
			Handler:    _AggregatorService_WatchCreation_Handler,
		},
		{
			MethodName: "SimilarCreations",
			Handler:    _AggregatorService_SimilarCreations_Handler,
		},
		{
			MethodName: "InitialComments",
			Handler:    _AggregatorService_InitialComments_Handler,
		},
		{
			MethodName: "GetTopComments",
			Handler:    _AggregatorService_GetTopComments_Handler,
		},
		{
			MethodName: "GetSecondComments",
			Handler:    _AggregatorService_GetSecondComments_Handler,
		},
		{
			MethodName: "GetUserReviews",
			Handler:    _AggregatorService_GetUserReviews_Handler,
		},
		{
			MethodName: "GetCreationReviews",
			Handler:    _AggregatorService_GetCreationReviews_Handler,
		},
		{
			MethodName: "GetCommentReviews",
			Handler:    _AggregatorService_GetCommentReviews_Handler,
		},
		{
			MethodName: "GetNewUserReviews",
			Handler:    _AggregatorService_GetNewUserReviews_Handler,
		},
		{
			MethodName: "GetNewCreationReviews",
			Handler:    _AggregatorService_GetNewCreationReviews_Handler,
		},
		{
			MethodName: "GetNewCommentReviews",
			Handler:    _AggregatorService_GetNewCommentReviews_Handler,
		},
		{
			MethodName: "HomePage",
			Handler:    _AggregatorService_HomePage_Handler,
		},
		{
			MethodName: "Collections",
			Handler:    _AggregatorService_Collections_Handler,
		},
		{
			MethodName: "History",
			Handler:    _AggregatorService_History_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "aggregator/agg_service.proto",
}
