// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.12.4
// source: mockgcp/cloud/gkemulticloud/v1/attached_service.proto

package gkemulticloudpb

import (
	longrunningpb "cloud.google.com/go/longrunning/autogen/longrunningpb"
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.32.0 or later.
const _ = grpc.SupportPackageIsVersion7

// AttachedClustersClient is the client API for AttachedClusters service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type AttachedClustersClient interface {
	// Creates a new
	// [AttachedCluster][mockgcp.cloud.gkemulticloud.v1.AttachedCluster] resource
	// on a given Google Cloud Platform project and region.
	//
	// If successful, the response contains a newly created
	// [Operation][google.longrunning.Operation] resource that can be
	// described to track the status of the operation.
	CreateAttachedCluster(ctx context.Context, in *CreateAttachedClusterRequest, opts ...grpc.CallOption) (*longrunningpb.Operation, error)
	// Updates an
	// [AttachedCluster][mockgcp.cloud.gkemulticloud.v1.AttachedCluster].
	UpdateAttachedCluster(ctx context.Context, in *UpdateAttachedClusterRequest, opts ...grpc.CallOption) (*longrunningpb.Operation, error)
	// Imports creates a new
	// [AttachedCluster][mockgcp.cloud.gkemulticloud.v1.AttachedCluster] resource
	// by importing an existing Fleet Membership resource.
	//
	// Attached Clusters created before the introduction of the Anthos Multi-Cloud
	// API can be imported through this method.
	//
	// If successful, the response contains a newly created
	// [Operation][google.longrunning.Operation] resource that can be
	// described to track the status of the operation.
	ImportAttachedCluster(ctx context.Context, in *ImportAttachedClusterRequest, opts ...grpc.CallOption) (*longrunningpb.Operation, error)
	// Describes a specific
	// [AttachedCluster][mockgcp.cloud.gkemulticloud.v1.AttachedCluster] resource.
	GetAttachedCluster(ctx context.Context, in *GetAttachedClusterRequest, opts ...grpc.CallOption) (*AttachedCluster, error)
	// Lists all [AttachedCluster][mockgcp.cloud.gkemulticloud.v1.AttachedCluster]
	// resources on a given Google Cloud project and region.
	ListAttachedClusters(ctx context.Context, in *ListAttachedClustersRequest, opts ...grpc.CallOption) (*ListAttachedClustersResponse, error)
	// Deletes a specific
	// [AttachedCluster][mockgcp.cloud.gkemulticloud.v1.AttachedCluster] resource.
	//
	// If successful, the response contains a newly created
	// [Operation][google.longrunning.Operation] resource that can be
	// described to track the status of the operation.
	DeleteAttachedCluster(ctx context.Context, in *DeleteAttachedClusterRequest, opts ...grpc.CallOption) (*longrunningpb.Operation, error)
	// Returns information, such as supported Kubernetes versions, on a given
	// Google Cloud location.
	GetAttachedServerConfig(ctx context.Context, in *GetAttachedServerConfigRequest, opts ...grpc.CallOption) (*AttachedServerConfig, error)
	// Generates the install manifest to be installed on the target cluster.
	GenerateAttachedClusterInstallManifest(ctx context.Context, in *GenerateAttachedClusterInstallManifestRequest, opts ...grpc.CallOption) (*GenerateAttachedClusterInstallManifestResponse, error)
}

type attachedClustersClient struct {
	cc grpc.ClientConnInterface
}

func NewAttachedClustersClient(cc grpc.ClientConnInterface) AttachedClustersClient {
	return &attachedClustersClient{cc}
}

func (c *attachedClustersClient) CreateAttachedCluster(ctx context.Context, in *CreateAttachedClusterRequest, opts ...grpc.CallOption) (*longrunningpb.Operation, error) {
	out := new(longrunningpb.Operation)
	err := c.cc.Invoke(ctx, "/mockgcp.cloud.gkemulticloud.v1.AttachedClusters/CreateAttachedCluster", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *attachedClustersClient) UpdateAttachedCluster(ctx context.Context, in *UpdateAttachedClusterRequest, opts ...grpc.CallOption) (*longrunningpb.Operation, error) {
	out := new(longrunningpb.Operation)
	err := c.cc.Invoke(ctx, "/mockgcp.cloud.gkemulticloud.v1.AttachedClusters/UpdateAttachedCluster", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *attachedClustersClient) ImportAttachedCluster(ctx context.Context, in *ImportAttachedClusterRequest, opts ...grpc.CallOption) (*longrunningpb.Operation, error) {
	out := new(longrunningpb.Operation)
	err := c.cc.Invoke(ctx, "/mockgcp.cloud.gkemulticloud.v1.AttachedClusters/ImportAttachedCluster", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *attachedClustersClient) GetAttachedCluster(ctx context.Context, in *GetAttachedClusterRequest, opts ...grpc.CallOption) (*AttachedCluster, error) {
	out := new(AttachedCluster)
	err := c.cc.Invoke(ctx, "/mockgcp.cloud.gkemulticloud.v1.AttachedClusters/GetAttachedCluster", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *attachedClustersClient) ListAttachedClusters(ctx context.Context, in *ListAttachedClustersRequest, opts ...grpc.CallOption) (*ListAttachedClustersResponse, error) {
	out := new(ListAttachedClustersResponse)
	err := c.cc.Invoke(ctx, "/mockgcp.cloud.gkemulticloud.v1.AttachedClusters/ListAttachedClusters", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *attachedClustersClient) DeleteAttachedCluster(ctx context.Context, in *DeleteAttachedClusterRequest, opts ...grpc.CallOption) (*longrunningpb.Operation, error) {
	out := new(longrunningpb.Operation)
	err := c.cc.Invoke(ctx, "/mockgcp.cloud.gkemulticloud.v1.AttachedClusters/DeleteAttachedCluster", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *attachedClustersClient) GetAttachedServerConfig(ctx context.Context, in *GetAttachedServerConfigRequest, opts ...grpc.CallOption) (*AttachedServerConfig, error) {
	out := new(AttachedServerConfig)
	err := c.cc.Invoke(ctx, "/mockgcp.cloud.gkemulticloud.v1.AttachedClusters/GetAttachedServerConfig", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *attachedClustersClient) GenerateAttachedClusterInstallManifest(ctx context.Context, in *GenerateAttachedClusterInstallManifestRequest, opts ...grpc.CallOption) (*GenerateAttachedClusterInstallManifestResponse, error) {
	out := new(GenerateAttachedClusterInstallManifestResponse)
	err := c.cc.Invoke(ctx, "/mockgcp.cloud.gkemulticloud.v1.AttachedClusters/GenerateAttachedClusterInstallManifest", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// AttachedClustersServer is the server API for AttachedClusters service.
// All implementations must embed UnimplementedAttachedClustersServer
// for forward compatibility
type AttachedClustersServer interface {
	// Creates a new
	// [AttachedCluster][mockgcp.cloud.gkemulticloud.v1.AttachedCluster] resource
	// on a given Google Cloud Platform project and region.
	//
	// If successful, the response contains a newly created
	// [Operation][google.longrunning.Operation] resource that can be
	// described to track the status of the operation.
	CreateAttachedCluster(context.Context, *CreateAttachedClusterRequest) (*longrunningpb.Operation, error)
	// Updates an
	// [AttachedCluster][mockgcp.cloud.gkemulticloud.v1.AttachedCluster].
	UpdateAttachedCluster(context.Context, *UpdateAttachedClusterRequest) (*longrunningpb.Operation, error)
	// Imports creates a new
	// [AttachedCluster][mockgcp.cloud.gkemulticloud.v1.AttachedCluster] resource
	// by importing an existing Fleet Membership resource.
	//
	// Attached Clusters created before the introduction of the Anthos Multi-Cloud
	// API can be imported through this method.
	//
	// If successful, the response contains a newly created
	// [Operation][google.longrunning.Operation] resource that can be
	// described to track the status of the operation.
	ImportAttachedCluster(context.Context, *ImportAttachedClusterRequest) (*longrunningpb.Operation, error)
	// Describes a specific
	// [AttachedCluster][mockgcp.cloud.gkemulticloud.v1.AttachedCluster] resource.
	GetAttachedCluster(context.Context, *GetAttachedClusterRequest) (*AttachedCluster, error)
	// Lists all [AttachedCluster][mockgcp.cloud.gkemulticloud.v1.AttachedCluster]
	// resources on a given Google Cloud project and region.
	ListAttachedClusters(context.Context, *ListAttachedClustersRequest) (*ListAttachedClustersResponse, error)
	// Deletes a specific
	// [AttachedCluster][mockgcp.cloud.gkemulticloud.v1.AttachedCluster] resource.
	//
	// If successful, the response contains a newly created
	// [Operation][google.longrunning.Operation] resource that can be
	// described to track the status of the operation.
	DeleteAttachedCluster(context.Context, *DeleteAttachedClusterRequest) (*longrunningpb.Operation, error)
	// Returns information, such as supported Kubernetes versions, on a given
	// Google Cloud location.
	GetAttachedServerConfig(context.Context, *GetAttachedServerConfigRequest) (*AttachedServerConfig, error)
	// Generates the install manifest to be installed on the target cluster.
	GenerateAttachedClusterInstallManifest(context.Context, *GenerateAttachedClusterInstallManifestRequest) (*GenerateAttachedClusterInstallManifestResponse, error)
	mustEmbedUnimplementedAttachedClustersServer()
}

// UnimplementedAttachedClustersServer must be embedded to have forward compatible implementations.
type UnimplementedAttachedClustersServer struct {
}

func (UnimplementedAttachedClustersServer) CreateAttachedCluster(context.Context, *CreateAttachedClusterRequest) (*longrunningpb.Operation, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateAttachedCluster not implemented")
}
func (UnimplementedAttachedClustersServer) UpdateAttachedCluster(context.Context, *UpdateAttachedClusterRequest) (*longrunningpb.Operation, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateAttachedCluster not implemented")
}
func (UnimplementedAttachedClustersServer) ImportAttachedCluster(context.Context, *ImportAttachedClusterRequest) (*longrunningpb.Operation, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ImportAttachedCluster not implemented")
}
func (UnimplementedAttachedClustersServer) GetAttachedCluster(context.Context, *GetAttachedClusterRequest) (*AttachedCluster, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetAttachedCluster not implemented")
}
func (UnimplementedAttachedClustersServer) ListAttachedClusters(context.Context, *ListAttachedClustersRequest) (*ListAttachedClustersResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListAttachedClusters not implemented")
}
func (UnimplementedAttachedClustersServer) DeleteAttachedCluster(context.Context, *DeleteAttachedClusterRequest) (*longrunningpb.Operation, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteAttachedCluster not implemented")
}
func (UnimplementedAttachedClustersServer) GetAttachedServerConfig(context.Context, *GetAttachedServerConfigRequest) (*AttachedServerConfig, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetAttachedServerConfig not implemented")
}
func (UnimplementedAttachedClustersServer) GenerateAttachedClusterInstallManifest(context.Context, *GenerateAttachedClusterInstallManifestRequest) (*GenerateAttachedClusterInstallManifestResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GenerateAttachedClusterInstallManifest not implemented")
}
func (UnimplementedAttachedClustersServer) mustEmbedUnimplementedAttachedClustersServer() {}

// UnsafeAttachedClustersServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to AttachedClustersServer will
// result in compilation errors.
type UnsafeAttachedClustersServer interface {
	mustEmbedUnimplementedAttachedClustersServer()
}

func RegisterAttachedClustersServer(s grpc.ServiceRegistrar, srv AttachedClustersServer) {
	s.RegisterService(&AttachedClusters_ServiceDesc, srv)
}

func _AttachedClusters_CreateAttachedCluster_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateAttachedClusterRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AttachedClustersServer).CreateAttachedCluster(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/mockgcp.cloud.gkemulticloud.v1.AttachedClusters/CreateAttachedCluster",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AttachedClustersServer).CreateAttachedCluster(ctx, req.(*CreateAttachedClusterRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _AttachedClusters_UpdateAttachedCluster_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateAttachedClusterRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AttachedClustersServer).UpdateAttachedCluster(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/mockgcp.cloud.gkemulticloud.v1.AttachedClusters/UpdateAttachedCluster",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AttachedClustersServer).UpdateAttachedCluster(ctx, req.(*UpdateAttachedClusterRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _AttachedClusters_ImportAttachedCluster_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ImportAttachedClusterRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AttachedClustersServer).ImportAttachedCluster(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/mockgcp.cloud.gkemulticloud.v1.AttachedClusters/ImportAttachedCluster",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AttachedClustersServer).ImportAttachedCluster(ctx, req.(*ImportAttachedClusterRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _AttachedClusters_GetAttachedCluster_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetAttachedClusterRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AttachedClustersServer).GetAttachedCluster(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/mockgcp.cloud.gkemulticloud.v1.AttachedClusters/GetAttachedCluster",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AttachedClustersServer).GetAttachedCluster(ctx, req.(*GetAttachedClusterRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _AttachedClusters_ListAttachedClusters_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListAttachedClustersRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AttachedClustersServer).ListAttachedClusters(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/mockgcp.cloud.gkemulticloud.v1.AttachedClusters/ListAttachedClusters",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AttachedClustersServer).ListAttachedClusters(ctx, req.(*ListAttachedClustersRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _AttachedClusters_DeleteAttachedCluster_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteAttachedClusterRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AttachedClustersServer).DeleteAttachedCluster(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/mockgcp.cloud.gkemulticloud.v1.AttachedClusters/DeleteAttachedCluster",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AttachedClustersServer).DeleteAttachedCluster(ctx, req.(*DeleteAttachedClusterRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _AttachedClusters_GetAttachedServerConfig_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetAttachedServerConfigRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AttachedClustersServer).GetAttachedServerConfig(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/mockgcp.cloud.gkemulticloud.v1.AttachedClusters/GetAttachedServerConfig",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AttachedClustersServer).GetAttachedServerConfig(ctx, req.(*GetAttachedServerConfigRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _AttachedClusters_GenerateAttachedClusterInstallManifest_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GenerateAttachedClusterInstallManifestRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AttachedClustersServer).GenerateAttachedClusterInstallManifest(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/mockgcp.cloud.gkemulticloud.v1.AttachedClusters/GenerateAttachedClusterInstallManifest",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AttachedClustersServer).GenerateAttachedClusterInstallManifest(ctx, req.(*GenerateAttachedClusterInstallManifestRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// AttachedClusters_ServiceDesc is the grpc.ServiceDesc for AttachedClusters service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var AttachedClusters_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "mockgcp.cloud.gkemulticloud.v1.AttachedClusters",
	HandlerType: (*AttachedClustersServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "CreateAttachedCluster",
			Handler:    _AttachedClusters_CreateAttachedCluster_Handler,
		},
		{
			MethodName: "UpdateAttachedCluster",
			Handler:    _AttachedClusters_UpdateAttachedCluster_Handler,
		},
		{
			MethodName: "ImportAttachedCluster",
			Handler:    _AttachedClusters_ImportAttachedCluster_Handler,
		},
		{
			MethodName: "GetAttachedCluster",
			Handler:    _AttachedClusters_GetAttachedCluster_Handler,
		},
		{
			MethodName: "ListAttachedClusters",
			Handler:    _AttachedClusters_ListAttachedClusters_Handler,
		},
		{
			MethodName: "DeleteAttachedCluster",
			Handler:    _AttachedClusters_DeleteAttachedCluster_Handler,
		},
		{
			MethodName: "GetAttachedServerConfig",
			Handler:    _AttachedClusters_GetAttachedServerConfig_Handler,
		},
		{
			MethodName: "GenerateAttachedClusterInstallManifest",
			Handler:    _AttachedClusters_GenerateAttachedClusterInstallManifest_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "mockgcp/cloud/gkemulticloud/v1/attached_service.proto",
}
