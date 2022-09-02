// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.19.4
// source: act.proto

package act

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

// ActClient is the client API for Act service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type ActClient interface {
	SaveProcDef(ctx context.Context, in *ProcDefReq, opts ...grpc.CallOption) (*ProcDefReply, error)
	FindDefByFormId(ctx context.Context, in *FindProcdefReq, opts ...grpc.CallOption) (*ProcDefReply, error)
	SetProcDefActive(ctx context.Context, in *FindProcdefReq, opts ...grpc.CallOption) (*ProcDefReply, error)
	DelProcDef(ctx context.Context, in *FindProcdefReq, opts ...grpc.CallOption) (*Nil, error)
	SaveProcInst(ctx context.Context, in *ProcInstReq, opts ...grpc.CallOption) (*ProcInstReply, error)
	SaveExecution(ctx context.Context, in *ExecutionReq, opts ...grpc.CallOption) (*ExecutionReply, error)
	SaveTask(ctx context.Context, in *TaskReq, opts ...grpc.CallOption) (*TaskReply, error)
	SaveIdentityLink(ctx context.Context, in *IdentityLinkReq, opts ...grpc.CallOption) (*IdentityLinkReply, error)
	FindLeastTaskId(ctx context.Context, in *DataIdReq, opts ...grpc.CallOption) (*TaskIdReply, error)
	DelProcInst(ctx context.Context, in *DataIdReq, opts ...grpc.CallOption) (*ProcInstReply, error)
	FindAllProcInst(ctx context.Context, in *ProcInstReq, opts ...grpc.CallOption) (*ProcInstReply, error)
	FindMyProcInst(ctx context.Context, in *MyProcInstReq, opts ...grpc.CallOption) (*ProcInstReply, error)
	FindMyApproval(ctx context.Context, in *MyProcInstReq, opts ...grpc.CallOption) (*ProcInstReply, error)
	FindOverTime(ctx context.Context, in *UserReq, opts ...grpc.CallOption) (*ProcInstReply, error)
	FindElapsedTime(ctx context.Context, in *UserReq, opts ...grpc.CallOption) (*ElapsedTimeReply, error)
}

type actClient struct {
	cc grpc.ClientConnInterface
}

func NewActClient(cc grpc.ClientConnInterface) ActClient {
	return &actClient{cc}
}

func (c *actClient) SaveProcDef(ctx context.Context, in *ProcDefReq, opts ...grpc.CallOption) (*ProcDefReply, error) {
	out := new(ProcDefReply)
	err := c.cc.Invoke(ctx, "/act.act/saveProcDef", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *actClient) FindDefByFormId(ctx context.Context, in *FindProcdefReq, opts ...grpc.CallOption) (*ProcDefReply, error) {
	out := new(ProcDefReply)
	err := c.cc.Invoke(ctx, "/act.act/findDefByFormId", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *actClient) SetProcDefActive(ctx context.Context, in *FindProcdefReq, opts ...grpc.CallOption) (*ProcDefReply, error) {
	out := new(ProcDefReply)
	err := c.cc.Invoke(ctx, "/act.act/setProcDefActive", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *actClient) DelProcDef(ctx context.Context, in *FindProcdefReq, opts ...grpc.CallOption) (*Nil, error) {
	out := new(Nil)
	err := c.cc.Invoke(ctx, "/act.act/delProcDef", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *actClient) SaveProcInst(ctx context.Context, in *ProcInstReq, opts ...grpc.CallOption) (*ProcInstReply, error) {
	out := new(ProcInstReply)
	err := c.cc.Invoke(ctx, "/act.act/saveProcInst", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *actClient) SaveExecution(ctx context.Context, in *ExecutionReq, opts ...grpc.CallOption) (*ExecutionReply, error) {
	out := new(ExecutionReply)
	err := c.cc.Invoke(ctx, "/act.act/saveExecution", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *actClient) SaveTask(ctx context.Context, in *TaskReq, opts ...grpc.CallOption) (*TaskReply, error) {
	out := new(TaskReply)
	err := c.cc.Invoke(ctx, "/act.act/saveTask", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *actClient) SaveIdentityLink(ctx context.Context, in *IdentityLinkReq, opts ...grpc.CallOption) (*IdentityLinkReply, error) {
	out := new(IdentityLinkReply)
	err := c.cc.Invoke(ctx, "/act.act/saveIdentityLink", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *actClient) FindLeastTaskId(ctx context.Context, in *DataIdReq, opts ...grpc.CallOption) (*TaskIdReply, error) {
	out := new(TaskIdReply)
	err := c.cc.Invoke(ctx, "/act.act/findLeastTaskId", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *actClient) DelProcInst(ctx context.Context, in *DataIdReq, opts ...grpc.CallOption) (*ProcInstReply, error) {
	out := new(ProcInstReply)
	err := c.cc.Invoke(ctx, "/act.act/delProcInst", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *actClient) FindAllProcInst(ctx context.Context, in *ProcInstReq, opts ...grpc.CallOption) (*ProcInstReply, error) {
	out := new(ProcInstReply)
	err := c.cc.Invoke(ctx, "/act.act/findAllProcInst", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *actClient) FindMyProcInst(ctx context.Context, in *MyProcInstReq, opts ...grpc.CallOption) (*ProcInstReply, error) {
	out := new(ProcInstReply)
	err := c.cc.Invoke(ctx, "/act.act/findMyProcInst", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *actClient) FindMyApproval(ctx context.Context, in *MyProcInstReq, opts ...grpc.CallOption) (*ProcInstReply, error) {
	out := new(ProcInstReply)
	err := c.cc.Invoke(ctx, "/act.act/findMyApproval", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *actClient) FindOverTime(ctx context.Context, in *UserReq, opts ...grpc.CallOption) (*ProcInstReply, error) {
	out := new(ProcInstReply)
	err := c.cc.Invoke(ctx, "/act.act/findOverTime", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *actClient) FindElapsedTime(ctx context.Context, in *UserReq, opts ...grpc.CallOption) (*ElapsedTimeReply, error) {
	out := new(ElapsedTimeReply)
	err := c.cc.Invoke(ctx, "/act.act/findElapsedTime", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// ActServer is the server API for Act service.
// All implementations must embed UnimplementedActServer
// for forward compatibility
type ActServer interface {
	SaveProcDef(context.Context, *ProcDefReq) (*ProcDefReply, error)
	FindDefByFormId(context.Context, *FindProcdefReq) (*ProcDefReply, error)
	SetProcDefActive(context.Context, *FindProcdefReq) (*ProcDefReply, error)
	DelProcDef(context.Context, *FindProcdefReq) (*Nil, error)
	SaveProcInst(context.Context, *ProcInstReq) (*ProcInstReply, error)
	SaveExecution(context.Context, *ExecutionReq) (*ExecutionReply, error)
	SaveTask(context.Context, *TaskReq) (*TaskReply, error)
	SaveIdentityLink(context.Context, *IdentityLinkReq) (*IdentityLinkReply, error)
	FindLeastTaskId(context.Context, *DataIdReq) (*TaskIdReply, error)
	DelProcInst(context.Context, *DataIdReq) (*ProcInstReply, error)
	FindAllProcInst(context.Context, *ProcInstReq) (*ProcInstReply, error)
	FindMyProcInst(context.Context, *MyProcInstReq) (*ProcInstReply, error)
	FindMyApproval(context.Context, *MyProcInstReq) (*ProcInstReply, error)
	FindOverTime(context.Context, *UserReq) (*ProcInstReply, error)
	FindElapsedTime(context.Context, *UserReq) (*ElapsedTimeReply, error)
	mustEmbedUnimplementedActServer()
}

// UnimplementedActServer must be embedded to have forward compatible implementations.
type UnimplementedActServer struct {
}

func (UnimplementedActServer) SaveProcDef(context.Context, *ProcDefReq) (*ProcDefReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SaveProcDef not implemented")
}
func (UnimplementedActServer) FindDefByFormId(context.Context, *FindProcdefReq) (*ProcDefReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method FindDefByFormId not implemented")
}
func (UnimplementedActServer) SetProcDefActive(context.Context, *FindProcdefReq) (*ProcDefReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SetProcDefActive not implemented")
}
func (UnimplementedActServer) DelProcDef(context.Context, *FindProcdefReq) (*Nil, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DelProcDef not implemented")
}
func (UnimplementedActServer) SaveProcInst(context.Context, *ProcInstReq) (*ProcInstReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SaveProcInst not implemented")
}
func (UnimplementedActServer) SaveExecution(context.Context, *ExecutionReq) (*ExecutionReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SaveExecution not implemented")
}
func (UnimplementedActServer) SaveTask(context.Context, *TaskReq) (*TaskReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SaveTask not implemented")
}
func (UnimplementedActServer) SaveIdentityLink(context.Context, *IdentityLinkReq) (*IdentityLinkReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SaveIdentityLink not implemented")
}
func (UnimplementedActServer) FindLeastTaskId(context.Context, *DataIdReq) (*TaskIdReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method FindLeastTaskId not implemented")
}
func (UnimplementedActServer) DelProcInst(context.Context, *DataIdReq) (*ProcInstReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DelProcInst not implemented")
}
func (UnimplementedActServer) FindAllProcInst(context.Context, *ProcInstReq) (*ProcInstReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method FindAllProcInst not implemented")
}
func (UnimplementedActServer) FindMyProcInst(context.Context, *MyProcInstReq) (*ProcInstReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method FindMyProcInst not implemented")
}
func (UnimplementedActServer) FindMyApproval(context.Context, *MyProcInstReq) (*ProcInstReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method FindMyApproval not implemented")
}
func (UnimplementedActServer) FindOverTime(context.Context, *UserReq) (*ProcInstReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method FindOverTime not implemented")
}
func (UnimplementedActServer) FindElapsedTime(context.Context, *UserReq) (*ElapsedTimeReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method FindElapsedTime not implemented")
}
func (UnimplementedActServer) mustEmbedUnimplementedActServer() {}

// UnsafeActServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to ActServer will
// result in compilation errors.
type UnsafeActServer interface {
	mustEmbedUnimplementedActServer()
}

func RegisterActServer(s grpc.ServiceRegistrar, srv ActServer) {
	s.RegisterService(&Act_ServiceDesc, srv)
}

func _Act_SaveProcDef_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ProcDefReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ActServer).SaveProcDef(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/act.act/saveProcDef",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ActServer).SaveProcDef(ctx, req.(*ProcDefReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _Act_FindDefByFormId_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(FindProcdefReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ActServer).FindDefByFormId(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/act.act/findDefByFormId",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ActServer).FindDefByFormId(ctx, req.(*FindProcdefReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _Act_SetProcDefActive_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(FindProcdefReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ActServer).SetProcDefActive(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/act.act/setProcDefActive",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ActServer).SetProcDefActive(ctx, req.(*FindProcdefReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _Act_DelProcDef_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(FindProcdefReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ActServer).DelProcDef(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/act.act/delProcDef",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ActServer).DelProcDef(ctx, req.(*FindProcdefReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _Act_SaveProcInst_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ProcInstReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ActServer).SaveProcInst(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/act.act/saveProcInst",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ActServer).SaveProcInst(ctx, req.(*ProcInstReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _Act_SaveExecution_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ExecutionReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ActServer).SaveExecution(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/act.act/saveExecution",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ActServer).SaveExecution(ctx, req.(*ExecutionReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _Act_SaveTask_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(TaskReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ActServer).SaveTask(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/act.act/saveTask",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ActServer).SaveTask(ctx, req.(*TaskReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _Act_SaveIdentityLink_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(IdentityLinkReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ActServer).SaveIdentityLink(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/act.act/saveIdentityLink",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ActServer).SaveIdentityLink(ctx, req.(*IdentityLinkReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _Act_FindLeastTaskId_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DataIdReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ActServer).FindLeastTaskId(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/act.act/findLeastTaskId",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ActServer).FindLeastTaskId(ctx, req.(*DataIdReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _Act_DelProcInst_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DataIdReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ActServer).DelProcInst(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/act.act/delProcInst",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ActServer).DelProcInst(ctx, req.(*DataIdReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _Act_FindAllProcInst_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ProcInstReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ActServer).FindAllProcInst(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/act.act/findAllProcInst",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ActServer).FindAllProcInst(ctx, req.(*ProcInstReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _Act_FindMyProcInst_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(MyProcInstReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ActServer).FindMyProcInst(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/act.act/findMyProcInst",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ActServer).FindMyProcInst(ctx, req.(*MyProcInstReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _Act_FindMyApproval_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(MyProcInstReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ActServer).FindMyApproval(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/act.act/findMyApproval",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ActServer).FindMyApproval(ctx, req.(*MyProcInstReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _Act_FindOverTime_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UserReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ActServer).FindOverTime(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/act.act/findOverTime",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ActServer).FindOverTime(ctx, req.(*UserReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _Act_FindElapsedTime_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UserReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ActServer).FindElapsedTime(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/act.act/findElapsedTime",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ActServer).FindElapsedTime(ctx, req.(*UserReq))
	}
	return interceptor(ctx, in, info, handler)
}

// Act_ServiceDesc is the grpc.ServiceDesc for Act service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Act_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "act.act",
	HandlerType: (*ActServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "saveProcDef",
			Handler:    _Act_SaveProcDef_Handler,
		},
		{
			MethodName: "findDefByFormId",
			Handler:    _Act_FindDefByFormId_Handler,
		},
		{
			MethodName: "setProcDefActive",
			Handler:    _Act_SetProcDefActive_Handler,
		},
		{
			MethodName: "delProcDef",
			Handler:    _Act_DelProcDef_Handler,
		},
		{
			MethodName: "saveProcInst",
			Handler:    _Act_SaveProcInst_Handler,
		},
		{
			MethodName: "saveExecution",
			Handler:    _Act_SaveExecution_Handler,
		},
		{
			MethodName: "saveTask",
			Handler:    _Act_SaveTask_Handler,
		},
		{
			MethodName: "saveIdentityLink",
			Handler:    _Act_SaveIdentityLink_Handler,
		},
		{
			MethodName: "findLeastTaskId",
			Handler:    _Act_FindLeastTaskId_Handler,
		},
		{
			MethodName: "delProcInst",
			Handler:    _Act_DelProcInst_Handler,
		},
		{
			MethodName: "findAllProcInst",
			Handler:    _Act_FindAllProcInst_Handler,
		},
		{
			MethodName: "findMyProcInst",
			Handler:    _Act_FindMyProcInst_Handler,
		},
		{
			MethodName: "findMyApproval",
			Handler:    _Act_FindMyApproval_Handler,
		},
		{
			MethodName: "findOverTime",
			Handler:    _Act_FindOverTime_Handler,
		},
		{
			MethodName: "findElapsedTime",
			Handler:    _Act_FindElapsedTime_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "act.proto",
}
