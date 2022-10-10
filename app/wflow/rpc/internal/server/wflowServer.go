// Code generated by goctl. DO NOT EDIT!
// Source: wflow.proto

package server

import (
	"context"

	"go-wflow/app/wflow/rpc/internal/logic"
	"go-wflow/app/wflow/rpc/internal/svc"
	"go-wflow/kernel"
)

type WflowServer struct {
	svcCtx *svc.ServiceContext
	kernel.UnimplementedWflowServer
}

func NewWflowServer(svcCtx *svc.ServiceContext) *WflowServer {
	return &WflowServer{
		svcCtx: svcCtx,
	}
}

func (s *WflowServer) AddProcDef(ctx context.Context, in *kernel.SaveProcDefReq) (*kernel.Procdef, error) {
	l := logic.NewAddProcDefLogic(ctx, s.svcCtx)
	return l.AddProcDef(in)
}

func (s *WflowServer) DelProcDef(ctx context.Context, in *kernel.IdReq) (*kernel.Nil, error) {
	l := logic.NewDelProcDefLogic(ctx, s.svcCtx)
	return l.DelProcDef(in)
}

func (s *WflowServer) UpdateProcDef(ctx context.Context, in *kernel.SaveProcDefReq) (*kernel.Procdef, error) {
	l := logic.NewUpdateProcDefLogic(ctx, s.svcCtx)
	return l.UpdateProcDef(in)
}

func (s *WflowServer) FindAllProcDef(ctx context.Context, in *kernel.PageReq) (*kernel.ProcdefArray, error) {
	l := logic.NewFindAllProcDefLogic(ctx, s.svcCtx)
	return l.FindAllProcDef(in)
}

func (s *WflowServer) FindOneProcDef(ctx context.Context, in *kernel.IdReq) (*kernel.Procdef, error) {
	l := logic.NewFindOneProcDefLogic(ctx, s.svcCtx)
	return l.FindOneProcDef(in)
}

func (s *WflowServer) SetActive(ctx context.Context, in *kernel.QueryProcDefReq) (*kernel.Nil, error) {
	l := logic.NewSetActiveLogic(ctx, s.svcCtx)
	return l.SetActive(in)
}

func (s *WflowServer) StartView(ctx context.Context, in *kernel.StartViewReq) (*kernel.StartViewResp, error) {
	l := logic.NewStartViewLogic(ctx, s.svcCtx)
	return l.StartView(in)
}

func (s *WflowServer) StartProcInst(ctx context.Context, in *kernel.StartProcInstReq) (*kernel.Nil, error) {
	l := logic.NewStartProcInstLogic(ctx, s.svcCtx)
	return l.StartProcInst(in)
}

func (s *WflowServer) Complete(ctx context.Context, in *kernel.CompleteProcInstReq) (*kernel.Nil, error) {
	l := logic.NewCompleteLogic(ctx, s.svcCtx)
	return l.Complete(in)
}

func (s *WflowServer) View(ctx context.Context, in *kernel.CompleteProcInstReq) (*kernel.Procinst, error) {
	l := logic.NewViewLogic(ctx, s.svcCtx)
	return l.View(in)
}

func (s *WflowServer) FindByDataId(ctx context.Context, in *kernel.DataIdReq) (*kernel.Procinst, error) {
	l := logic.NewFindByDataIdLogic(ctx, s.svcCtx)
	return l.FindByDataId(in)
}

func (s *WflowServer) Withdraw(ctx context.Context, in *kernel.DataIdReq) (*kernel.Nil, error) {
	l := logic.NewWithdrawLogic(ctx, s.svcCtx)
	return l.Withdraw(in)
}

func (s *WflowServer) DelProcInst(ctx context.Context, in *kernel.DataIdReq) (*kernel.Nil, error) {
	l := logic.NewDelProcInstLogic(ctx, s.svcCtx)
	return l.DelProcInst(in)
}
