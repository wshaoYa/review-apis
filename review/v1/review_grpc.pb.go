// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.5.1
// - protoc             v5.28.0
// source: api/review/v1/review.proto

package v1

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
	Review_CreateReview_FullMethodName = "/api.review.v1.Review/CreateReview"
	Review_UpdateReview_FullMethodName = "/api.review.v1.Review/UpdateReview"
	Review_DeleteReview_FullMethodName = "/api.review.v1.Review/DeleteReview"
	Review_GetReview_FullMethodName    = "/api.review.v1.Review/GetReview"
	Review_ListReview_FullMethodName   = "/api.review.v1.Review/ListReview"
	Review_ReplyReview_FullMethodName  = "/api.review.v1.Review/ReplyReview"
	Review_AppealReview_FullMethodName = "/api.review.v1.Review/AppealReview"
	Review_AuditReview_FullMethodName  = "/api.review.v1.Review/AuditReview"
	Review_AuditAppeal_FullMethodName  = "/api.review.v1.Review/AuditAppeal"
)

// ReviewClient is the client API for Review service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type ReviewClient interface {
	// C端 创建评价
	CreateReview(ctx context.Context, in *CreateReviewRequest, opts ...grpc.CallOption) (*CreateReviewReply, error)
	UpdateReview(ctx context.Context, in *UpdateReviewRequest, opts ...grpc.CallOption) (*UpdateReviewReply, error)
	DeleteReview(ctx context.Context, in *DeleteReviewRequest, opts ...grpc.CallOption) (*DeleteReviewReply, error)
	GetReview(ctx context.Context, in *GetReviewRequest, opts ...grpc.CallOption) (*GetReviewReply, error)
	ListReview(ctx context.Context, in *ListReviewRequest, opts ...grpc.CallOption) (*ListReviewReply, error)
	// B端 回复评价
	ReplyReview(ctx context.Context, in *ReplyReviewRequest, opts ...grpc.CallOption) (*ReplyReviewReply, error)
	// B端 申诉评价
	AppealReview(ctx context.Context, in *AppealReviewRequest, opts ...grpc.CallOption) (*AppealReviewReply, error)
	// O端 审核评价
	AuditReview(ctx context.Context, in *AuditReviewRequest, opts ...grpc.CallOption) (*AuditReviewReply, error)
	// O端 审核申诉
	AuditAppeal(ctx context.Context, in *AuditAppealRequest, opts ...grpc.CallOption) (*AuditAppealReply, error)
}

type reviewClient struct {
	cc grpc.ClientConnInterface
}

func NewReviewClient(cc grpc.ClientConnInterface) ReviewClient {
	return &reviewClient{cc}
}

func (c *reviewClient) CreateReview(ctx context.Context, in *CreateReviewRequest, opts ...grpc.CallOption) (*CreateReviewReply, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(CreateReviewReply)
	err := c.cc.Invoke(ctx, Review_CreateReview_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *reviewClient) UpdateReview(ctx context.Context, in *UpdateReviewRequest, opts ...grpc.CallOption) (*UpdateReviewReply, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(UpdateReviewReply)
	err := c.cc.Invoke(ctx, Review_UpdateReview_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *reviewClient) DeleteReview(ctx context.Context, in *DeleteReviewRequest, opts ...grpc.CallOption) (*DeleteReviewReply, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(DeleteReviewReply)
	err := c.cc.Invoke(ctx, Review_DeleteReview_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *reviewClient) GetReview(ctx context.Context, in *GetReviewRequest, opts ...grpc.CallOption) (*GetReviewReply, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(GetReviewReply)
	err := c.cc.Invoke(ctx, Review_GetReview_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *reviewClient) ListReview(ctx context.Context, in *ListReviewRequest, opts ...grpc.CallOption) (*ListReviewReply, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(ListReviewReply)
	err := c.cc.Invoke(ctx, Review_ListReview_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *reviewClient) ReplyReview(ctx context.Context, in *ReplyReviewRequest, opts ...grpc.CallOption) (*ReplyReviewReply, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(ReplyReviewReply)
	err := c.cc.Invoke(ctx, Review_ReplyReview_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *reviewClient) AppealReview(ctx context.Context, in *AppealReviewRequest, opts ...grpc.CallOption) (*AppealReviewReply, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(AppealReviewReply)
	err := c.cc.Invoke(ctx, Review_AppealReview_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *reviewClient) AuditReview(ctx context.Context, in *AuditReviewRequest, opts ...grpc.CallOption) (*AuditReviewReply, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(AuditReviewReply)
	err := c.cc.Invoke(ctx, Review_AuditReview_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *reviewClient) AuditAppeal(ctx context.Context, in *AuditAppealRequest, opts ...grpc.CallOption) (*AuditAppealReply, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(AuditAppealReply)
	err := c.cc.Invoke(ctx, Review_AuditAppeal_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// ReviewServer is the server API for Review service.
// All implementations must embed UnimplementedReviewServer
// for forward compatibility.
type ReviewServer interface {
	// C端 创建评价
	CreateReview(context.Context, *CreateReviewRequest) (*CreateReviewReply, error)
	UpdateReview(context.Context, *UpdateReviewRequest) (*UpdateReviewReply, error)
	DeleteReview(context.Context, *DeleteReviewRequest) (*DeleteReviewReply, error)
	GetReview(context.Context, *GetReviewRequest) (*GetReviewReply, error)
	ListReview(context.Context, *ListReviewRequest) (*ListReviewReply, error)
	// B端 回复评价
	ReplyReview(context.Context, *ReplyReviewRequest) (*ReplyReviewReply, error)
	// B端 申诉评价
	AppealReview(context.Context, *AppealReviewRequest) (*AppealReviewReply, error)
	// O端 审核评价
	AuditReview(context.Context, *AuditReviewRequest) (*AuditReviewReply, error)
	// O端 审核申诉
	AuditAppeal(context.Context, *AuditAppealRequest) (*AuditAppealReply, error)
	mustEmbedUnimplementedReviewServer()
}

// UnimplementedReviewServer must be embedded to have
// forward compatible implementations.
//
// NOTE: this should be embedded by value instead of pointer to avoid a nil
// pointer dereference when methods are called.
type UnimplementedReviewServer struct{}

func (UnimplementedReviewServer) CreateReview(context.Context, *CreateReviewRequest) (*CreateReviewReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateReview not implemented")
}
func (UnimplementedReviewServer) UpdateReview(context.Context, *UpdateReviewRequest) (*UpdateReviewReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateReview not implemented")
}
func (UnimplementedReviewServer) DeleteReview(context.Context, *DeleteReviewRequest) (*DeleteReviewReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteReview not implemented")
}
func (UnimplementedReviewServer) GetReview(context.Context, *GetReviewRequest) (*GetReviewReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetReview not implemented")
}
func (UnimplementedReviewServer) ListReview(context.Context, *ListReviewRequest) (*ListReviewReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListReview not implemented")
}
func (UnimplementedReviewServer) ReplyReview(context.Context, *ReplyReviewRequest) (*ReplyReviewReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ReplyReview not implemented")
}
func (UnimplementedReviewServer) AppealReview(context.Context, *AppealReviewRequest) (*AppealReviewReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method AppealReview not implemented")
}
func (UnimplementedReviewServer) AuditReview(context.Context, *AuditReviewRequest) (*AuditReviewReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method AuditReview not implemented")
}
func (UnimplementedReviewServer) AuditAppeal(context.Context, *AuditAppealRequest) (*AuditAppealReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method AuditAppeal not implemented")
}
func (UnimplementedReviewServer) mustEmbedUnimplementedReviewServer() {}
func (UnimplementedReviewServer) testEmbeddedByValue()                {}

// UnsafeReviewServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to ReviewServer will
// result in compilation errors.
type UnsafeReviewServer interface {
	mustEmbedUnimplementedReviewServer()
}

func RegisterReviewServer(s grpc.ServiceRegistrar, srv ReviewServer) {
	// If the following call pancis, it indicates UnimplementedReviewServer was
	// embedded by pointer and is nil.  This will cause panics if an
	// unimplemented method is ever invoked, so we test this at initialization
	// time to prevent it from happening at runtime later due to I/O.
	if t, ok := srv.(interface{ testEmbeddedByValue() }); ok {
		t.testEmbeddedByValue()
	}
	s.RegisterService(&Review_ServiceDesc, srv)
}

func _Review_CreateReview_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateReviewRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ReviewServer).CreateReview(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Review_CreateReview_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ReviewServer).CreateReview(ctx, req.(*CreateReviewRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Review_UpdateReview_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateReviewRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ReviewServer).UpdateReview(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Review_UpdateReview_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ReviewServer).UpdateReview(ctx, req.(*UpdateReviewRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Review_DeleteReview_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteReviewRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ReviewServer).DeleteReview(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Review_DeleteReview_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ReviewServer).DeleteReview(ctx, req.(*DeleteReviewRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Review_GetReview_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetReviewRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ReviewServer).GetReview(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Review_GetReview_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ReviewServer).GetReview(ctx, req.(*GetReviewRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Review_ListReview_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListReviewRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ReviewServer).ListReview(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Review_ListReview_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ReviewServer).ListReview(ctx, req.(*ListReviewRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Review_ReplyReview_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ReplyReviewRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ReviewServer).ReplyReview(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Review_ReplyReview_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ReviewServer).ReplyReview(ctx, req.(*ReplyReviewRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Review_AppealReview_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AppealReviewRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ReviewServer).AppealReview(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Review_AppealReview_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ReviewServer).AppealReview(ctx, req.(*AppealReviewRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Review_AuditReview_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AuditReviewRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ReviewServer).AuditReview(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Review_AuditReview_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ReviewServer).AuditReview(ctx, req.(*AuditReviewRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Review_AuditAppeal_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AuditAppealRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ReviewServer).AuditAppeal(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Review_AuditAppeal_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ReviewServer).AuditAppeal(ctx, req.(*AuditAppealRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// Review_ServiceDesc is the grpc.ServiceDesc for Review service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Review_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "api.review.v1.Review",
	HandlerType: (*ReviewServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "CreateReview",
			Handler:    _Review_CreateReview_Handler,
		},
		{
			MethodName: "UpdateReview",
			Handler:    _Review_UpdateReview_Handler,
		},
		{
			MethodName: "DeleteReview",
			Handler:    _Review_DeleteReview_Handler,
		},
		{
			MethodName: "GetReview",
			Handler:    _Review_GetReview_Handler,
		},
		{
			MethodName: "ListReview",
			Handler:    _Review_ListReview_Handler,
		},
		{
			MethodName: "ReplyReview",
			Handler:    _Review_ReplyReview_Handler,
		},
		{
			MethodName: "AppealReview",
			Handler:    _Review_AppealReview_Handler,
		},
		{
			MethodName: "AuditReview",
			Handler:    _Review_AuditReview_Handler,
		},
		{
			MethodName: "AuditAppeal",
			Handler:    _Review_AuditAppeal_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "api/review/v1/review.proto",
}
