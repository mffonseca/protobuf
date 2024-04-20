// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             v3.21.12
// source: webhook.proto

package generated

import (
	"context"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)


// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.32.0 or later.
const _ = grpc.SupportPackageIsVersion7

const (
	WorkflowService_TriggerWorkflow_FullMethodName = "/grpc.WorkflowService/TriggerWorkflow"
)

// WorkflowServiceClient is the client API for WorkflowService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type WorkflowServiceClient interface {
	// Define uma chamada RPC com os parâmetros esperados e a resposta
	TriggerWorkflow(ctx context.Context, in *WorkflowRequest, opts ...grpc.CallOption) (*WorkflowResponse, error)
}

type workflowServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewWorkflowServiceClient(cc grpc.ClientConnInterface) WorkflowServiceClient {
	return &workflowServiceClient{cc}
}

func (c *workflowServiceClient) TriggerWorkflow(ctx context.Context, in *WorkflowRequest, opts ...grpc.CallOption) (*WorkflowResponse, error) {
	out := new(WorkflowResponse)
	err := c.cc.Invoke(ctx, WorkflowService_TriggerWorkflow_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// WorkflowServiceServer is the server API for WorkflowService service.
// All implementations must embed UnimplementedWorkflowServiceServer
// for forward compatibility
type WorkflowServiceServer interface {
	// Define uma chamada RPC com os parâmetros esperados e a resposta
	TriggerWorkflow(context.Context, *WorkflowRequest) (*WorkflowResponse, error)
	mustEmbedUnimplementedWorkflowServiceServer()
}

// UnimplementedWorkflowServiceServer must be embedded to have forward compatible implementations.
type UnimplementedWorkflowServiceServer struct {
}

func (UnimplementedWorkflowServiceServer) TriggerWorkflow(context.Context, *WorkflowRequest) (*WorkflowResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method TriggerWorkflow not implemented")
}
func (UnimplementedWorkflowServiceServer) mustEmbedUnimplementedWorkflowServiceServer() {}

// UnsafeWorkflowServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to WorkflowServiceServer will
// result in compilation errors.
type UnsafeWorkflowServiceServer interface {
	mustEmbedUnimplementedWorkflowServiceServer()
}

func RegisterWorkflowServiceServer(s grpc.ServiceRegistrar, srv WorkflowServiceServer) {
	s.RegisterService(&WorkflowService_ServiceDesc, srv)
}

func _WorkflowService_TriggerWorkflow_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(WorkflowRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(WorkflowServiceServer).TriggerWorkflow(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: WorkflowService_TriggerWorkflow_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(WorkflowServiceServer).TriggerWorkflow(ctx, req.(*WorkflowRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// WorkflowService_ServiceDesc is the grpc.ServiceDesc for WorkflowService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var WorkflowService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "grpc.WorkflowService",
	HandlerType: (*WorkflowServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "TriggerWorkflow",
			Handler:    _WorkflowService_TriggerWorkflow_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "webhook.proto",
}
