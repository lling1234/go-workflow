// Code generated by goctl. DO NOT EDIT!
// Source: act.proto

package actclient

import (
	"context"

	"act/rpc/types/act"

	"github.com/zeromicro/go-zero/zrpc"
	"google.golang.org/grpc"
)

type (
	CommonRpcRes      = act.CommonRpcRes
	DataIdReq         = act.DataIdReq
	ElapsedTimeReply  = act.ElapsedTimeReply
	ExecutionReply    = act.ExecutionReply
	ExecutionReq      = act.ExecutionReq
	FindProcDefReq    = act.FindProcDefReq
	IdRequest         = act.IdRequest
	IdentityLinkReply = act.IdentityLinkReply
	IdentityLinkReq   = act.IdentityLinkReq
	MyProcInstReq     = act.MyProcInstReq
	Nil               = act.Nil
	PageReq           = act.PageReq
	PageRequest       = act.PageRequest
	ProcDefReply      = act.ProcDefReply
	ProcInstIdArg     = act.ProcInstIdArg
	ProcInstReply     = act.ProcInstReply
	ProcInstReq       = act.ProcInstReq
	SaveProcDefReq    = act.SaveProcDefReq
	SearchProcInstReq = act.SearchProcInstReq
	TaskIdArg         = act.TaskIdArg
	TaskReply         = act.TaskReply
	TaskReq           = act.TaskReq
	UpdateProcInstReq = act.UpdateProcInstReq
	UserReq           = act.UserReq

	Act interface {
		SaveProcDef(ctx context.Context, in *SaveProcDefReq, opts ...grpc.CallOption) (*ProcDefReply, error)
		FindDefByFormId(ctx context.Context, in *FindProcDefReq, opts ...grpc.CallOption) (*ProcDefReply, error)
		UpdateProcDef(ctx context.Context, in *FindProcDefReq, opts ...grpc.CallOption) (*ProcDefReply, error)
		SetProcDefActive(ctx context.Context, in *FindProcDefReq, opts ...grpc.CallOption) (*Nil, error)
		DelProcDef(ctx context.Context, in *FindProcDefReq, opts ...grpc.CallOption) (*Nil, error)
		SaveProcInst(ctx context.Context, in *ProcInstReq, opts ...grpc.CallOption) (*ProcInstReply, error)
		SaveExecution(ctx context.Context, in *ExecutionReq, opts ...grpc.CallOption) (*ExecutionReply, error)
		SaveTask(ctx context.Context, in *TaskReq, opts ...grpc.CallOption) (*TaskReply, error)
		SaveIdentityLink(ctx context.Context, in *IdentityLinkReq, opts ...grpc.CallOption) (*IdentityLinkReply, error)
		FindLatestTask(ctx context.Context, in *ProcInstIdArg, opts ...grpc.CallOption) (*TaskReply, error)
		UpdateProcInst(ctx context.Context, in *UpdateProcInstReq, opts ...grpc.CallOption) (*ProcInstReply, error)
		UpdateTask(ctx context.Context, in *TaskReq, opts ...grpc.CallOption) (*TaskReply, error)
		UpdateIdentityLink(ctx context.Context, in *IdentityLinkReq, opts ...grpc.CallOption) (*IdentityLinkReply, error)
		DelProcInst(ctx context.Context, in *DataIdReq, opts ...grpc.CallOption) (*Nil, error)
		FindIdentityLinkByTaskId(ctx context.Context, in *TaskIdArg, opts ...grpc.CallOption) (*IdentityLinkReply, error)
		FindExecutionByInstId(ctx context.Context, in *ProcInstIdArg, opts ...grpc.CallOption) (*ExecutionReq, error)
		DelIdentityLink(ctx context.Context, in *ProcInstIdArg, opts ...grpc.CallOption) (*Nil, error)
		Withdraw(ctx context.Context, in *DataIdReq, opts ...grpc.CallOption) (*Nil, error)
<<<<<<< HEAD
=======
		FindProcInstByDataId(ctx context.Context, in *DataIdReq, opts ...grpc.CallOption) (*ProcInstReply, error)
>>>>>>> 4a54178f732840048a6d221f9338a6e8f5d12ba3
		FindAllProcInst(ctx context.Context, in *IdRequest, opts ...grpc.CallOption) (*CommonRpcRes, error)
		FindMyProcInst(ctx context.Context, in *MyProcInstReq, opts ...grpc.CallOption) (*ProcInstReply, error)
		FindMyApproval(ctx context.Context, in *MyProcInstReq, opts ...grpc.CallOption) (*ProcInstReply, error)
		FindOverTime(ctx context.Context, in *UserReq, opts ...grpc.CallOption) (*ProcInstReply, error)
		FindElapsedTime(ctx context.Context, in *UserReq, opts ...grpc.CallOption) (*ElapsedTimeReply, error)
	}

	defaultAct struct {
		cli zrpc.Client
	}
)

func NewAct(cli zrpc.Client) Act {
	return &defaultAct{
		cli: cli,
	}
}

func (m *defaultAct) SaveProcDef(ctx context.Context, in *SaveProcDefReq, opts ...grpc.CallOption) (*ProcDefReply, error) {
	client := act.NewActClient(m.cli.Conn())
	return client.SaveProcDef(ctx, in, opts...)
}

func (m *defaultAct) FindDefByFormId(ctx context.Context, in *FindProcDefReq, opts ...grpc.CallOption) (*ProcDefReply, error) {
	client := act.NewActClient(m.cli.Conn())
	return client.FindDefByFormId(ctx, in, opts...)
}

func (m *defaultAct) UpdateProcDef(ctx context.Context, in *FindProcDefReq, opts ...grpc.CallOption) (*ProcDefReply, error) {
	client := act.NewActClient(m.cli.Conn())
	return client.UpdateProcDef(ctx, in, opts...)
}

func (m *defaultAct) SetProcDefActive(ctx context.Context, in *FindProcDefReq, opts ...grpc.CallOption) (*Nil, error) {
	client := act.NewActClient(m.cli.Conn())
	return client.SetProcDefActive(ctx, in, opts...)
}

func (m *defaultAct) DelProcDef(ctx context.Context, in *FindProcDefReq, opts ...grpc.CallOption) (*Nil, error) {
	client := act.NewActClient(m.cli.Conn())
	return client.DelProcDef(ctx, in, opts...)
}

func (m *defaultAct) SaveProcInst(ctx context.Context, in *ProcInstReq, opts ...grpc.CallOption) (*ProcInstReply, error) {
	client := act.NewActClient(m.cli.Conn())
	return client.SaveProcInst(ctx, in, opts...)
}

func (m *defaultAct) SaveExecution(ctx context.Context, in *ExecutionReq, opts ...grpc.CallOption) (*ExecutionReply, error) {
	client := act.NewActClient(m.cli.Conn())
	return client.SaveExecution(ctx, in, opts...)
}

func (m *defaultAct) SaveTask(ctx context.Context, in *TaskReq, opts ...grpc.CallOption) (*TaskReply, error) {
	client := act.NewActClient(m.cli.Conn())
	return client.SaveTask(ctx, in, opts...)
}

func (m *defaultAct) SaveIdentityLink(ctx context.Context, in *IdentityLinkReq, opts ...grpc.CallOption) (*IdentityLinkReply, error) {
	client := act.NewActClient(m.cli.Conn())
	return client.SaveIdentityLink(ctx, in, opts...)
}

func (m *defaultAct) FindLatestTask(ctx context.Context, in *ProcInstIdArg, opts ...grpc.CallOption) (*TaskReply, error) {
	client := act.NewActClient(m.cli.Conn())
	return client.FindLatestTask(ctx, in, opts...)
}

func (m *defaultAct) UpdateProcInst(ctx context.Context, in *UpdateProcInstReq, opts ...grpc.CallOption) (*ProcInstReply, error) {
	client := act.NewActClient(m.cli.Conn())
	return client.UpdateProcInst(ctx, in, opts...)
}

func (m *defaultAct) UpdateTask(ctx context.Context, in *TaskReq, opts ...grpc.CallOption) (*TaskReply, error) {
	client := act.NewActClient(m.cli.Conn())
	return client.UpdateTask(ctx, in, opts...)
}

func (m *defaultAct) UpdateIdentityLink(ctx context.Context, in *IdentityLinkReq, opts ...grpc.CallOption) (*IdentityLinkReply, error) {
	client := act.NewActClient(m.cli.Conn())
	return client.UpdateIdentityLink(ctx, in, opts...)
}

func (m *defaultAct) DelProcInst(ctx context.Context, in *DataIdReq, opts ...grpc.CallOption) (*Nil, error) {
	client := act.NewActClient(m.cli.Conn())
	return client.DelProcInst(ctx, in, opts...)
}

func (m *defaultAct) FindIdentityLinkByTaskId(ctx context.Context, in *TaskIdArg, opts ...grpc.CallOption) (*IdentityLinkReply, error) {
	client := act.NewActClient(m.cli.Conn())
	return client.FindIdentityLinkByTaskId(ctx, in, opts...)
}

func (m *defaultAct) FindExecutionByInstId(ctx context.Context, in *ProcInstIdArg, opts ...grpc.CallOption) (*ExecutionReq, error) {
	client := act.NewActClient(m.cli.Conn())
	return client.FindExecutionByInstId(ctx, in, opts...)
}

func (m *defaultAct) DelIdentityLink(ctx context.Context, in *ProcInstIdArg, opts ...grpc.CallOption) (*Nil, error) {
	client := act.NewActClient(m.cli.Conn())
	return client.DelIdentityLink(ctx, in, opts...)
}

func (m *defaultAct) Withdraw(ctx context.Context, in *DataIdReq, opts ...grpc.CallOption) (*Nil, error) {
	client := act.NewActClient(m.cli.Conn())
	return client.Withdraw(ctx, in, opts...)
}

<<<<<<< HEAD
=======
func (m *defaultAct) FindProcInstByDataId(ctx context.Context, in *DataIdReq, opts ...grpc.CallOption) (*ProcInstReply, error) {
	client := act.NewActClient(m.cli.Conn())
	return client.FindProcInstByDataId(ctx, in, opts...)
}

>>>>>>> 4a54178f732840048a6d221f9338a6e8f5d12ba3
func (m *defaultAct) FindAllProcInst(ctx context.Context, in *IdRequest, opts ...grpc.CallOption) (*CommonRpcRes, error) {
	client := act.NewActClient(m.cli.Conn())
	return client.FindAllProcInst(ctx, in, opts...)
}

func (m *defaultAct) FindMyProcInst(ctx context.Context, in *MyProcInstReq, opts ...grpc.CallOption) (*ProcInstReply, error) {
	client := act.NewActClient(m.cli.Conn())
	return client.FindMyProcInst(ctx, in, opts...)
}

func (m *defaultAct) FindMyApproval(ctx context.Context, in *MyProcInstReq, opts ...grpc.CallOption) (*ProcInstReply, error) {
	client := act.NewActClient(m.cli.Conn())
	return client.FindMyApproval(ctx, in, opts...)
}

func (m *defaultAct) FindOverTime(ctx context.Context, in *UserReq, opts ...grpc.CallOption) (*ProcInstReply, error) {
	client := act.NewActClient(m.cli.Conn())
	return client.FindOverTime(ctx, in, opts...)
}

func (m *defaultAct) FindElapsedTime(ctx context.Context, in *UserReq, opts ...grpc.CallOption) (*ElapsedTimeReply, error) {
	client := act.NewActClient(m.cli.Conn())
	return client.FindElapsedTime(ctx, in, opts...)
}
