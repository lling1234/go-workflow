// Code generated by goctl. DO NOT EDIT!
// Source: act.proto

package server

import (
	"context"

	"act/rpc/internal/logic"
	"act/rpc/internal/svc"
	"act/rpc/types/act"
)

type ActServer struct {
	svcCtx *svc.ServiceContext
	act.UnimplementedActServer
}

func NewActServer(svcCtx *svc.ServiceContext) *ActServer {
	return &ActServer{
		svcCtx: svcCtx,
	}
}

func (s *ActServer) SaveProcDef(ctx context.Context, in *act.ProcDefReq) (*act.ProcDefReply, error) {
	l := logic.NewSaveProcDefLogic(ctx, s.svcCtx)
	return l.SaveProcDef(in)
}

func (s *ActServer) FindDefByFormId(ctx context.Context, in *act.FindProcdefReq) (*act.ProcDefReply, error) {
	l := logic.NewFindDefByFormIdLogic(ctx, s.svcCtx)
	return l.FindDefByFormId(in)
}

func (s *ActServer) SetProcDefActive(ctx context.Context, in *act.FindProcdefReq) (*act.ProcDefReply, error) {
	l := logic.NewSetProcDefActiveLogic(ctx, s.svcCtx)
	return l.SetProcDefActive(in)
}

func (s *ActServer) DelProcDef(ctx context.Context, in *act.FindProcdefReq) (*act.Nil, error) {
	l := logic.NewDelProcDefLogic(ctx, s.svcCtx)
	return l.DelProcDef(in)
}

func (s *ActServer) SaveProcInst(ctx context.Context, in *act.ProcInstReq) (*act.ProcInstReply, error) {
	l := logic.NewSaveProcInstLogic(ctx, s.svcCtx)
	return l.SaveProcInst(in)
}

func (s *ActServer) SaveExecution(ctx context.Context, in *act.ExecutionReq) (*act.ExecutionReply, error) {
	l := logic.NewSaveExecutionLogic(ctx, s.svcCtx)
	return l.SaveExecution(in)
}

func (s *ActServer) SaveTask(ctx context.Context, in *act.TaskReq) (*act.TaskReply, error) {
	l := logic.NewSaveTaskLogic(ctx, s.svcCtx)
	return l.SaveTask(in)
}

func (s *ActServer) SaveIdentityLink(ctx context.Context, in *act.IdentityLinkReq) (*act.IdentityLinkReply, error) {
	l := logic.NewSaveIdentityLinkLogic(ctx, s.svcCtx)
	return l.SaveIdentityLink(in)
}

func (s *ActServer) FindLeastTaskId(ctx context.Context, in *act.DataIdReq) (*act.TaskIdReply, error) {
	l := logic.NewFindLeastTaskIdLogic(ctx, s.svcCtx)
	return l.FindLeastTaskId(in)
}

func (s *ActServer) DelProcInst(ctx context.Context, in *act.DataIdReq) (*act.ProcInstReply, error) {
	l := logic.NewDelProcInstLogic(ctx, s.svcCtx)
	return l.DelProcInst(in)
}

func (s *ActServer) FindAllProcInst(ctx context.Context, in *act.ProcInstReq) (*act.ProcInstReply, error) {
	l := logic.NewFindAllProcInstLogic(ctx, s.svcCtx)
	return l.FindAllProcInst(in)
}

func (s *ActServer) FindMyProcInst(ctx context.Context, in *act.MyProcInstReq) (*act.ProcInstReply, error) {
	l := logic.NewFindMyProcInstLogic(ctx, s.svcCtx)
	return l.FindMyProcInst(in)
}

func (s *ActServer) FindMyApproval(ctx context.Context, in *act.MyProcInstReq) (*act.ProcInstReply, error) {
	l := logic.NewFindMyApprovalLogic(ctx, s.svcCtx)
	return l.FindMyApproval(in)
}

func (s *ActServer) FindOverTime(ctx context.Context, in *act.UserReq) (*act.ProcInstReply, error) {
	l := logic.NewFindOverTimeLogic(ctx, s.svcCtx)
	return l.FindOverTime(in)
}

func (s *ActServer) FindElapsedTime(ctx context.Context, in *act.UserReq) (*act.ElapsedTimeReply, error) {
	l := logic.NewFindElapsedTimeLogic(ctx, s.svcCtx)
	return l.FindElapsedTime(in)
}
