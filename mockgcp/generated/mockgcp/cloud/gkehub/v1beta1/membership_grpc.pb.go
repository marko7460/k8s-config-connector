// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.12.4
// source: mockgcp/cloud/gkehub/v1beta1/membership.proto

package gkehubpb

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

// GkeHubMembershipServiceClient is the client API for GkeHubMembershipService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type GkeHubMembershipServiceClient interface {
	// Lists Memberships in a given project and location.
	ListMemberships(ctx context.Context, in *ListMembershipsRequest, opts ...grpc.CallOption) (*ListMembershipsResponse, error)
	// Gets the details of a Membership.
	GetMembership(ctx context.Context, in *GetMembershipRequest, opts ...grpc.CallOption) (*Membership, error)
	// Creates a new Membership.
	//
	// **This is currently only supported for GKE clusters on Google Cloud**.
	// To register other clusters, follow the instructions at
	// https://cloud.google.com/anthos/multicluster-management/connect/registering-a-cluster.
	CreateMembership(ctx context.Context, in *CreateMembershipRequest, opts ...grpc.CallOption) (*longrunningpb.Operation, error)
	// Removes a Membership.
	//
	// **This is currently only supported for GKE clusters on Google Cloud**.
	// To unregister other clusters, follow the instructions at
	// https://cloud.google.com/anthos/multicluster-management/connect/unregistering-a-cluster.
	DeleteMembership(ctx context.Context, in *DeleteMembershipRequest, opts ...grpc.CallOption) (*longrunningpb.Operation, error)
	// Updates an existing Membership.
	UpdateMembership(ctx context.Context, in *UpdateMembershipRequest, opts ...grpc.CallOption) (*longrunningpb.Operation, error)
	// Generates the manifest for deployment of the GKE connect agent.
	//
	// **This method is used internally by Google-provided libraries.**
	// Most clients should not need to call this method directly.
	GenerateConnectManifest(ctx context.Context, in *GenerateConnectManifestRequest, opts ...grpc.CallOption) (*GenerateConnectManifestResponse, error)
	// ValidateExclusivity validates the state of exclusivity in the cluster.
	// The validation does not depend on an existing Hub membership resource.
	ValidateExclusivity(ctx context.Context, in *ValidateExclusivityRequest, opts ...grpc.CallOption) (*ValidateExclusivityResponse, error)
	// GenerateExclusivityManifest generates the manifests to update the
	// exclusivity artifacts in the cluster if needed.
	//
	// Exclusivity artifacts include the Membership custom resource definition
	// (CRD) and the singleton Membership custom resource (CR). Combined with
	// ValidateExclusivity, exclusivity artifacts guarantee that a Kubernetes
	// cluster is only registered to a single GKE Hub.
	//
	// The Membership CRD is versioned, and may require conversion when the GKE
	// Hub API server begins serving a newer version of the CRD and
	// corresponding CR. The response will be the converted CRD and CR if there
	// are any differences between the versions.
	GenerateExclusivityManifest(ctx context.Context, in *GenerateExclusivityManifestRequest, opts ...grpc.CallOption) (*GenerateExclusivityManifestResponse, error)
}

type gkeHubMembershipServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewGkeHubMembershipServiceClient(cc grpc.ClientConnInterface) GkeHubMembershipServiceClient {
	return &gkeHubMembershipServiceClient{cc}
}

func (c *gkeHubMembershipServiceClient) ListMemberships(ctx context.Context, in *ListMembershipsRequest, opts ...grpc.CallOption) (*ListMembershipsResponse, error) {
	out := new(ListMembershipsResponse)
	err := c.cc.Invoke(ctx, "/mockgcp.cloud.gkehub.v1beta1.GkeHubMembershipService/ListMemberships", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *gkeHubMembershipServiceClient) GetMembership(ctx context.Context, in *GetMembershipRequest, opts ...grpc.CallOption) (*Membership, error) {
	out := new(Membership)
	err := c.cc.Invoke(ctx, "/mockgcp.cloud.gkehub.v1beta1.GkeHubMembershipService/GetMembership", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *gkeHubMembershipServiceClient) CreateMembership(ctx context.Context, in *CreateMembershipRequest, opts ...grpc.CallOption) (*longrunningpb.Operation, error) {
	out := new(longrunningpb.Operation)
	err := c.cc.Invoke(ctx, "/mockgcp.cloud.gkehub.v1beta1.GkeHubMembershipService/CreateMembership", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *gkeHubMembershipServiceClient) DeleteMembership(ctx context.Context, in *DeleteMembershipRequest, opts ...grpc.CallOption) (*longrunningpb.Operation, error) {
	out := new(longrunningpb.Operation)
	err := c.cc.Invoke(ctx, "/mockgcp.cloud.gkehub.v1beta1.GkeHubMembershipService/DeleteMembership", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *gkeHubMembershipServiceClient) UpdateMembership(ctx context.Context, in *UpdateMembershipRequest, opts ...grpc.CallOption) (*longrunningpb.Operation, error) {
	out := new(longrunningpb.Operation)
	err := c.cc.Invoke(ctx, "/mockgcp.cloud.gkehub.v1beta1.GkeHubMembershipService/UpdateMembership", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *gkeHubMembershipServiceClient) GenerateConnectManifest(ctx context.Context, in *GenerateConnectManifestRequest, opts ...grpc.CallOption) (*GenerateConnectManifestResponse, error) {
	out := new(GenerateConnectManifestResponse)
	err := c.cc.Invoke(ctx, "/mockgcp.cloud.gkehub.v1beta1.GkeHubMembershipService/GenerateConnectManifest", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *gkeHubMembershipServiceClient) ValidateExclusivity(ctx context.Context, in *ValidateExclusivityRequest, opts ...grpc.CallOption) (*ValidateExclusivityResponse, error) {
	out := new(ValidateExclusivityResponse)
	err := c.cc.Invoke(ctx, "/mockgcp.cloud.gkehub.v1beta1.GkeHubMembershipService/ValidateExclusivity", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *gkeHubMembershipServiceClient) GenerateExclusivityManifest(ctx context.Context, in *GenerateExclusivityManifestRequest, opts ...grpc.CallOption) (*GenerateExclusivityManifestResponse, error) {
	out := new(GenerateExclusivityManifestResponse)
	err := c.cc.Invoke(ctx, "/mockgcp.cloud.gkehub.v1beta1.GkeHubMembershipService/GenerateExclusivityManifest", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// GkeHubMembershipServiceServer is the server API for GkeHubMembershipService service.
// All implementations must embed UnimplementedGkeHubMembershipServiceServer
// for forward compatibility
type GkeHubMembershipServiceServer interface {
	// Lists Memberships in a given project and location.
	ListMemberships(context.Context, *ListMembershipsRequest) (*ListMembershipsResponse, error)
	// Gets the details of a Membership.
	GetMembership(context.Context, *GetMembershipRequest) (*Membership, error)
	// Creates a new Membership.
	//
	// **This is currently only supported for GKE clusters on Google Cloud**.
	// To register other clusters, follow the instructions at
	// https://cloud.google.com/anthos/multicluster-management/connect/registering-a-cluster.
	CreateMembership(context.Context, *CreateMembershipRequest) (*longrunningpb.Operation, error)
	// Removes a Membership.
	//
	// **This is currently only supported for GKE clusters on Google Cloud**.
	// To unregister other clusters, follow the instructions at
	// https://cloud.google.com/anthos/multicluster-management/connect/unregistering-a-cluster.
	DeleteMembership(context.Context, *DeleteMembershipRequest) (*longrunningpb.Operation, error)
	// Updates an existing Membership.
	UpdateMembership(context.Context, *UpdateMembershipRequest) (*longrunningpb.Operation, error)
	// Generates the manifest for deployment of the GKE connect agent.
	//
	// **This method is used internally by Google-provided libraries.**
	// Most clients should not need to call this method directly.
	GenerateConnectManifest(context.Context, *GenerateConnectManifestRequest) (*GenerateConnectManifestResponse, error)
	// ValidateExclusivity validates the state of exclusivity in the cluster.
	// The validation does not depend on an existing Hub membership resource.
	ValidateExclusivity(context.Context, *ValidateExclusivityRequest) (*ValidateExclusivityResponse, error)
	// GenerateExclusivityManifest generates the manifests to update the
	// exclusivity artifacts in the cluster if needed.
	//
	// Exclusivity artifacts include the Membership custom resource definition
	// (CRD) and the singleton Membership custom resource (CR). Combined with
	// ValidateExclusivity, exclusivity artifacts guarantee that a Kubernetes
	// cluster is only registered to a single GKE Hub.
	//
	// The Membership CRD is versioned, and may require conversion when the GKE
	// Hub API server begins serving a newer version of the CRD and
	// corresponding CR. The response will be the converted CRD and CR if there
	// are any differences between the versions.
	GenerateExclusivityManifest(context.Context, *GenerateExclusivityManifestRequest) (*GenerateExclusivityManifestResponse, error)
	mustEmbedUnimplementedGkeHubMembershipServiceServer()
}

// UnimplementedGkeHubMembershipServiceServer must be embedded to have forward compatible implementations.
type UnimplementedGkeHubMembershipServiceServer struct {
}

func (UnimplementedGkeHubMembershipServiceServer) ListMemberships(context.Context, *ListMembershipsRequest) (*ListMembershipsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListMemberships not implemented")
}
func (UnimplementedGkeHubMembershipServiceServer) GetMembership(context.Context, *GetMembershipRequest) (*Membership, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetMembership not implemented")
}
func (UnimplementedGkeHubMembershipServiceServer) CreateMembership(context.Context, *CreateMembershipRequest) (*longrunningpb.Operation, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateMembership not implemented")
}
func (UnimplementedGkeHubMembershipServiceServer) DeleteMembership(context.Context, *DeleteMembershipRequest) (*longrunningpb.Operation, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteMembership not implemented")
}
func (UnimplementedGkeHubMembershipServiceServer) UpdateMembership(context.Context, *UpdateMembershipRequest) (*longrunningpb.Operation, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateMembership not implemented")
}
func (UnimplementedGkeHubMembershipServiceServer) GenerateConnectManifest(context.Context, *GenerateConnectManifestRequest) (*GenerateConnectManifestResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GenerateConnectManifest not implemented")
}
func (UnimplementedGkeHubMembershipServiceServer) ValidateExclusivity(context.Context, *ValidateExclusivityRequest) (*ValidateExclusivityResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ValidateExclusivity not implemented")
}
func (UnimplementedGkeHubMembershipServiceServer) GenerateExclusivityManifest(context.Context, *GenerateExclusivityManifestRequest) (*GenerateExclusivityManifestResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GenerateExclusivityManifest not implemented")
}
func (UnimplementedGkeHubMembershipServiceServer) mustEmbedUnimplementedGkeHubMembershipServiceServer() {
}

// UnsafeGkeHubMembershipServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to GkeHubMembershipServiceServer will
// result in compilation errors.
type UnsafeGkeHubMembershipServiceServer interface {
	mustEmbedUnimplementedGkeHubMembershipServiceServer()
}

func RegisterGkeHubMembershipServiceServer(s grpc.ServiceRegistrar, srv GkeHubMembershipServiceServer) {
	s.RegisterService(&GkeHubMembershipService_ServiceDesc, srv)
}

func _GkeHubMembershipService_ListMemberships_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListMembershipsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GkeHubMembershipServiceServer).ListMemberships(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/mockgcp.cloud.gkehub.v1beta1.GkeHubMembershipService/ListMemberships",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GkeHubMembershipServiceServer).ListMemberships(ctx, req.(*ListMembershipsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _GkeHubMembershipService_GetMembership_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetMembershipRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GkeHubMembershipServiceServer).GetMembership(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/mockgcp.cloud.gkehub.v1beta1.GkeHubMembershipService/GetMembership",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GkeHubMembershipServiceServer).GetMembership(ctx, req.(*GetMembershipRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _GkeHubMembershipService_CreateMembership_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateMembershipRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GkeHubMembershipServiceServer).CreateMembership(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/mockgcp.cloud.gkehub.v1beta1.GkeHubMembershipService/CreateMembership",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GkeHubMembershipServiceServer).CreateMembership(ctx, req.(*CreateMembershipRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _GkeHubMembershipService_DeleteMembership_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteMembershipRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GkeHubMembershipServiceServer).DeleteMembership(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/mockgcp.cloud.gkehub.v1beta1.GkeHubMembershipService/DeleteMembership",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GkeHubMembershipServiceServer).DeleteMembership(ctx, req.(*DeleteMembershipRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _GkeHubMembershipService_UpdateMembership_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateMembershipRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GkeHubMembershipServiceServer).UpdateMembership(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/mockgcp.cloud.gkehub.v1beta1.GkeHubMembershipService/UpdateMembership",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GkeHubMembershipServiceServer).UpdateMembership(ctx, req.(*UpdateMembershipRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _GkeHubMembershipService_GenerateConnectManifest_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GenerateConnectManifestRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GkeHubMembershipServiceServer).GenerateConnectManifest(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/mockgcp.cloud.gkehub.v1beta1.GkeHubMembershipService/GenerateConnectManifest",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GkeHubMembershipServiceServer).GenerateConnectManifest(ctx, req.(*GenerateConnectManifestRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _GkeHubMembershipService_ValidateExclusivity_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ValidateExclusivityRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GkeHubMembershipServiceServer).ValidateExclusivity(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/mockgcp.cloud.gkehub.v1beta1.GkeHubMembershipService/ValidateExclusivity",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GkeHubMembershipServiceServer).ValidateExclusivity(ctx, req.(*ValidateExclusivityRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _GkeHubMembershipService_GenerateExclusivityManifest_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GenerateExclusivityManifestRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GkeHubMembershipServiceServer).GenerateExclusivityManifest(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/mockgcp.cloud.gkehub.v1beta1.GkeHubMembershipService/GenerateExclusivityManifest",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GkeHubMembershipServiceServer).GenerateExclusivityManifest(ctx, req.(*GenerateExclusivityManifestRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// GkeHubMembershipService_ServiceDesc is the grpc.ServiceDesc for GkeHubMembershipService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var GkeHubMembershipService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "mockgcp.cloud.gkehub.v1beta1.GkeHubMembershipService",
	HandlerType: (*GkeHubMembershipServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "ListMemberships",
			Handler:    _GkeHubMembershipService_ListMemberships_Handler,
		},
		{
			MethodName: "GetMembership",
			Handler:    _GkeHubMembershipService_GetMembership_Handler,
		},
		{
			MethodName: "CreateMembership",
			Handler:    _GkeHubMembershipService_CreateMembership_Handler,
		},
		{
			MethodName: "DeleteMembership",
			Handler:    _GkeHubMembershipService_DeleteMembership_Handler,
		},
		{
			MethodName: "UpdateMembership",
			Handler:    _GkeHubMembershipService_UpdateMembership_Handler,
		},
		{
			MethodName: "GenerateConnectManifest",
			Handler:    _GkeHubMembershipService_GenerateConnectManifest_Handler,
		},
		{
			MethodName: "ValidateExclusivity",
			Handler:    _GkeHubMembershipService_ValidateExclusivity_Handler,
		},
		{
			MethodName: "GenerateExclusivityManifest",
			Handler:    _GkeHubMembershipService_GenerateExclusivityManifest_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "mockgcp/cloud/gkehub/v1beta1/membership.proto",
}
