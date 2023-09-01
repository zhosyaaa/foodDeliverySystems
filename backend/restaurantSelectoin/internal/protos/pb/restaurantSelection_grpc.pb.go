// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v4.24.1
// source: internal/protos/restaurantSelection.proto

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

// RestaurantSelectionServiceClient is the client API for RestaurantSelectionService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type RestaurantSelectionServiceClient interface {
	GetRestaurants(ctx context.Context, in *GetRestaurantsRequest, opts ...grpc.CallOption) (*GetRestaurantsResponse, error)
	GetRestaurantDetails(ctx context.Context, in *GetRestaurantDetailsRequest, opts ...grpc.CallOption) (*RestaurantDetails, error)
	AddNewRestaurant(ctx context.Context, in *AddNewRestaurantRequest, opts ...grpc.CallOption) (*Response, error)
	UpdateRestaurantInfo(ctx context.Context, in *UpdateRestaurantInfoRequest, opts ...grpc.CallOption) (*Response, error)
	SearchRestaurants(ctx context.Context, in *SearchRestaurantsRequest, opts ...grpc.CallOption) (*SearchRestaurantsResponse, error)
	DeleteRestaurant(ctx context.Context, in *DeleteRestaurantRequest, opts ...grpc.CallOption) (*Response, error)
}

type restaurantSelectionServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewRestaurantSelectionServiceClient(cc grpc.ClientConnInterface) RestaurantSelectionServiceClient {
	return &restaurantSelectionServiceClient{cc}
}

func (c *restaurantSelectionServiceClient) GetRestaurants(ctx context.Context, in *GetRestaurantsRequest, opts ...grpc.CallOption) (*GetRestaurantsResponse, error) {
	out := new(GetRestaurantsResponse)
	err := c.cc.Invoke(ctx, "/restaurant.RestaurantSelectionService/GetRestaurants", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *restaurantSelectionServiceClient) GetRestaurantDetails(ctx context.Context, in *GetRestaurantDetailsRequest, opts ...grpc.CallOption) (*RestaurantDetails, error) {
	out := new(RestaurantDetails)
	err := c.cc.Invoke(ctx, "/restaurant.RestaurantSelectionService/GetRestaurantDetails", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *restaurantSelectionServiceClient) AddNewRestaurant(ctx context.Context, in *AddNewRestaurantRequest, opts ...grpc.CallOption) (*Response, error) {
	out := new(Response)
	err := c.cc.Invoke(ctx, "/restaurant.RestaurantSelectionService/AddNewRestaurant", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *restaurantSelectionServiceClient) UpdateRestaurantInfo(ctx context.Context, in *UpdateRestaurantInfoRequest, opts ...grpc.CallOption) (*Response, error) {
	out := new(Response)
	err := c.cc.Invoke(ctx, "/restaurant.RestaurantSelectionService/UpdateRestaurantInfo", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *restaurantSelectionServiceClient) SearchRestaurants(ctx context.Context, in *SearchRestaurantsRequest, opts ...grpc.CallOption) (*SearchRestaurantsResponse, error) {
	out := new(SearchRestaurantsResponse)
	err := c.cc.Invoke(ctx, "/restaurant.RestaurantSelectionService/SearchRestaurants", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *restaurantSelectionServiceClient) DeleteRestaurant(ctx context.Context, in *DeleteRestaurantRequest, opts ...grpc.CallOption) (*Response, error) {
	out := new(Response)
	err := c.cc.Invoke(ctx, "/restaurant.RestaurantSelectionService/DeleteRestaurant", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// RestaurantSelectionServiceServer is the server API for RestaurantSelectionService service.
// All implementations must embed UnimplementedRestaurantSelectionServiceServer
// for forward compatibility
type RestaurantSelectionServiceServer interface {
	GetRestaurants(context.Context, *GetRestaurantsRequest) (*GetRestaurantsResponse, error)
	GetRestaurantDetails(context.Context, *GetRestaurantDetailsRequest) (*RestaurantDetails, error)
	AddNewRestaurant(context.Context, *AddNewRestaurantRequest) (*Response, error)
	UpdateRestaurantInfo(context.Context, *UpdateRestaurantInfoRequest) (*Response, error)
	SearchRestaurants(context.Context, *SearchRestaurantsRequest) (*SearchRestaurantsResponse, error)
	DeleteRestaurant(context.Context, *DeleteRestaurantRequest) (*Response, error)
	mustEmbedUnimplementedRestaurantSelectionServiceServer()
}

// UnimplementedRestaurantSelectionServiceServer must be embedded to have forward compatible implementations.
type UnimplementedRestaurantSelectionServiceServer struct {
}

func (UnimplementedRestaurantSelectionServiceServer) GetRestaurants(context.Context, *GetRestaurantsRequest) (*GetRestaurantsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetRestaurants not implemented")
}
func (UnimplementedRestaurantSelectionServiceServer) GetRestaurantDetails(context.Context, *GetRestaurantDetailsRequest) (*RestaurantDetails, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetRestaurantDetails not implemented")
}
func (UnimplementedRestaurantSelectionServiceServer) AddNewRestaurant(context.Context, *AddNewRestaurantRequest) (*Response, error) {
	return nil, status.Errorf(codes.Unimplemented, "method AddNewRestaurant not implemented")
}
func (UnimplementedRestaurantSelectionServiceServer) UpdateRestaurantInfo(context.Context, *UpdateRestaurantInfoRequest) (*Response, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateRestaurantInfo not implemented")
}
func (UnimplementedRestaurantSelectionServiceServer) SearchRestaurants(context.Context, *SearchRestaurantsRequest) (*SearchRestaurantsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SearchRestaurants not implemented")
}
func (UnimplementedRestaurantSelectionServiceServer) DeleteRestaurant(context.Context, *DeleteRestaurantRequest) (*Response, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteRestaurant not implemented")
}
func (UnimplementedRestaurantSelectionServiceServer) mustEmbedUnimplementedRestaurantSelectionServiceServer() {
}

// UnsafeRestaurantSelectionServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to RestaurantSelectionServiceServer will
// result in compilation errors.
type UnsafeRestaurantSelectionServiceServer interface {
	mustEmbedUnimplementedRestaurantSelectionServiceServer()
}

func RegisterRestaurantSelectionServiceServer(s grpc.ServiceRegistrar, srv RestaurantSelectionServiceServer) {
	s.RegisterService(&RestaurantSelectionService_ServiceDesc, srv)
}

func _RestaurantSelectionService_GetRestaurants_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetRestaurantsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RestaurantSelectionServiceServer).GetRestaurants(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/restaurant.RestaurantSelectionService/GetRestaurants",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RestaurantSelectionServiceServer).GetRestaurants(ctx, req.(*GetRestaurantsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _RestaurantSelectionService_GetRestaurantDetails_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetRestaurantDetailsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RestaurantSelectionServiceServer).GetRestaurantDetails(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/restaurant.RestaurantSelectionService/GetRestaurantDetails",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RestaurantSelectionServiceServer).GetRestaurantDetails(ctx, req.(*GetRestaurantDetailsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _RestaurantSelectionService_AddNewRestaurant_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AddNewRestaurantRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RestaurantSelectionServiceServer).AddNewRestaurant(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/restaurant.RestaurantSelectionService/AddNewRestaurant",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RestaurantSelectionServiceServer).AddNewRestaurant(ctx, req.(*AddNewRestaurantRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _RestaurantSelectionService_UpdateRestaurantInfo_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateRestaurantInfoRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RestaurantSelectionServiceServer).UpdateRestaurantInfo(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/restaurant.RestaurantSelectionService/UpdateRestaurantInfo",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RestaurantSelectionServiceServer).UpdateRestaurantInfo(ctx, req.(*UpdateRestaurantInfoRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _RestaurantSelectionService_SearchRestaurants_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SearchRestaurantsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RestaurantSelectionServiceServer).SearchRestaurants(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/restaurant.RestaurantSelectionService/SearchRestaurants",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RestaurantSelectionServiceServer).SearchRestaurants(ctx, req.(*SearchRestaurantsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _RestaurantSelectionService_DeleteRestaurant_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteRestaurantRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RestaurantSelectionServiceServer).DeleteRestaurant(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/restaurant.RestaurantSelectionService/DeleteRestaurant",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RestaurantSelectionServiceServer).DeleteRestaurant(ctx, req.(*DeleteRestaurantRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// RestaurantSelectionService_ServiceDesc is the grpc.ServiceDesc for RestaurantSelectionService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var RestaurantSelectionService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "restaurant.RestaurantSelectionService",
	HandlerType: (*RestaurantSelectionServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetRestaurants",
			Handler:    _RestaurantSelectionService_GetRestaurants_Handler,
		},
		{
			MethodName: "GetRestaurantDetails",
			Handler:    _RestaurantSelectionService_GetRestaurantDetails_Handler,
		},
		{
			MethodName: "AddNewRestaurant",
			Handler:    _RestaurantSelectionService_AddNewRestaurant_Handler,
		},
		{
			MethodName: "UpdateRestaurantInfo",
			Handler:    _RestaurantSelectionService_UpdateRestaurantInfo_Handler,
		},
		{
			MethodName: "SearchRestaurants",
			Handler:    _RestaurantSelectionService_SearchRestaurants_Handler,
		},
		{
			MethodName: "DeleteRestaurant",
			Handler:    _RestaurantSelectionService_DeleteRestaurant_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "internal/protos/restaurantSelection.proto",
}