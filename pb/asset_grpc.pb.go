// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.6.1
// source: asset.proto

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

// AssetServiceClient is the client API for AssetService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type AssetServiceClient interface {
	Insert(ctx context.Context, in *CreateAsset, opts ...grpc.CallOption) (*Asset, error)
	Read(ctx context.Context, in *ID, opts ...grpc.CallOption) (*Asset, error)
	Delete(ctx context.Context, in *ID, opts ...grpc.CallOption) (*ID, error)
	Update(ctx context.Context, in *Asset, opts ...grpc.CallOption) (*Asset, error)
	StreamList(ctx context.Context, opts ...grpc.CallOption) (AssetService_StreamListClient, error)
	GetAll(ctx context.Context, in *Void, opts ...grpc.CallOption) (*Assets, error)
}

type assetServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewAssetServiceClient(cc grpc.ClientConnInterface) AssetServiceClient {
	return &assetServiceClient{cc}
}

func (c *assetServiceClient) Insert(ctx context.Context, in *CreateAsset, opts ...grpc.CallOption) (*Asset, error) {
	out := new(Asset)
	err := c.cc.Invoke(ctx, "/pb.AssetService/Insert", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *assetServiceClient) Read(ctx context.Context, in *ID, opts ...grpc.CallOption) (*Asset, error) {
	out := new(Asset)
	err := c.cc.Invoke(ctx, "/pb.AssetService/Read", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *assetServiceClient) Delete(ctx context.Context, in *ID, opts ...grpc.CallOption) (*ID, error) {
	out := new(ID)
	err := c.cc.Invoke(ctx, "/pb.AssetService/Delete", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *assetServiceClient) Update(ctx context.Context, in *Asset, opts ...grpc.CallOption) (*Asset, error) {
	out := new(Asset)
	err := c.cc.Invoke(ctx, "/pb.AssetService/Update", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *assetServiceClient) StreamList(ctx context.Context, opts ...grpc.CallOption) (AssetService_StreamListClient, error) {
	stream, err := c.cc.NewStream(ctx, &AssetService_ServiceDesc.Streams[0], "/pb.AssetService/StreamList", opts...)
	if err != nil {
		return nil, err
	}
	x := &assetServiceStreamListClient{stream}
	return x, nil
}

type AssetService_StreamListClient interface {
	Send(*CreateAsset) error
	CloseAndRecv() (*Assets, error)
	grpc.ClientStream
}

type assetServiceStreamListClient struct {
	grpc.ClientStream
}

func (x *assetServiceStreamListClient) Send(m *CreateAsset) error {
	return x.ClientStream.SendMsg(m)
}

func (x *assetServiceStreamListClient) CloseAndRecv() (*Assets, error) {
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	m := new(Assets)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *assetServiceClient) GetAll(ctx context.Context, in *Void, opts ...grpc.CallOption) (*Assets, error) {
	out := new(Assets)
	err := c.cc.Invoke(ctx, "/pb.AssetService/GetAll", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// AssetServiceServer is the server API for AssetService service.
// All implementations must embed UnimplementedAssetServiceServer
// for forward compatibility
type AssetServiceServer interface {
	Insert(context.Context, *CreateAsset) (*Asset, error)
	Read(context.Context, *ID) (*Asset, error)
	Delete(context.Context, *ID) (*ID, error)
	Update(context.Context, *Asset) (*Asset, error)
	StreamList(AssetService_StreamListServer) error
	GetAll(context.Context, *Void) (*Assets, error)
	mustEmbedUnimplementedAssetServiceServer()
}

// UnimplementedAssetServiceServer must be embedded to have forward compatible implementations.
type UnimplementedAssetServiceServer struct {
}

func (UnimplementedAssetServiceServer) Insert(context.Context, *CreateAsset) (*Asset, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Insert not implemented")
}
func (UnimplementedAssetServiceServer) Read(context.Context, *ID) (*Asset, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Read not implemented")
}
func (UnimplementedAssetServiceServer) Delete(context.Context, *ID) (*ID, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Delete not implemented")
}
func (UnimplementedAssetServiceServer) Update(context.Context, *Asset) (*Asset, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Update not implemented")
}
func (UnimplementedAssetServiceServer) StreamList(AssetService_StreamListServer) error {
	return status.Errorf(codes.Unimplemented, "method StreamList not implemented")
}
func (UnimplementedAssetServiceServer) GetAll(context.Context, *Void) (*Assets, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetAll not implemented")
}
func (UnimplementedAssetServiceServer) mustEmbedUnimplementedAssetServiceServer() {}

// UnsafeAssetServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to AssetServiceServer will
// result in compilation errors.
type UnsafeAssetServiceServer interface {
	mustEmbedUnimplementedAssetServiceServer()
}

func RegisterAssetServiceServer(s grpc.ServiceRegistrar, srv AssetServiceServer) {
	s.RegisterService(&AssetService_ServiceDesc, srv)
}

func _AssetService_Insert_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateAsset)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AssetServiceServer).Insert(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pb.AssetService/Insert",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AssetServiceServer).Insert(ctx, req.(*CreateAsset))
	}
	return interceptor(ctx, in, info, handler)
}

func _AssetService_Read_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ID)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AssetServiceServer).Read(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pb.AssetService/Read",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AssetServiceServer).Read(ctx, req.(*ID))
	}
	return interceptor(ctx, in, info, handler)
}

func _AssetService_Delete_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ID)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AssetServiceServer).Delete(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pb.AssetService/Delete",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AssetServiceServer).Delete(ctx, req.(*ID))
	}
	return interceptor(ctx, in, info, handler)
}

func _AssetService_Update_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Asset)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AssetServiceServer).Update(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pb.AssetService/Update",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AssetServiceServer).Update(ctx, req.(*Asset))
	}
	return interceptor(ctx, in, info, handler)
}

func _AssetService_StreamList_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(AssetServiceServer).StreamList(&assetServiceStreamListServer{stream})
}

type AssetService_StreamListServer interface {
	SendAndClose(*Assets) error
	Recv() (*CreateAsset, error)
	grpc.ServerStream
}

type assetServiceStreamListServer struct {
	grpc.ServerStream
}

func (x *assetServiceStreamListServer) SendAndClose(m *Assets) error {
	return x.ServerStream.SendMsg(m)
}

func (x *assetServiceStreamListServer) Recv() (*CreateAsset, error) {
	m := new(CreateAsset)
	if err := x.ServerStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func _AssetService_GetAll_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Void)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AssetServiceServer).GetAll(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pb.AssetService/GetAll",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AssetServiceServer).GetAll(ctx, req.(*Void))
	}
	return interceptor(ctx, in, info, handler)
}

// AssetService_ServiceDesc is the grpc.ServiceDesc for AssetService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var AssetService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "pb.AssetService",
	HandlerType: (*AssetServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Insert",
			Handler:    _AssetService_Insert_Handler,
		},
		{
			MethodName: "Read",
			Handler:    _AssetService_Read_Handler,
		},
		{
			MethodName: "Delete",
			Handler:    _AssetService_Delete_Handler,
		},
		{
			MethodName: "Update",
			Handler:    _AssetService_Update_Handler,
		},
		{
			MethodName: "GetAll",
			Handler:    _AssetService_GetAll_Handler,
		},
	},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "StreamList",
			Handler:       _AssetService_StreamList_Handler,
			ClientStreams: true,
		},
	},
	Metadata: "asset.proto",
}
