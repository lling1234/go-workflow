package wflow

import (
	"context"

	"go-wflow/app/wflow/api/internal/svc"
	"go-wflow/app/wflow/api/internal/types"
	"go-wflow/kernel"

	"github.com/qkbyte/go-zero/core/logx"
)

type StartProcInstLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewStartProcInstLogic(ctx context.Context, svcCtx *svc.ServiceContext) *StartProcInstLogic {
	return &StartProcInstLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *StartProcInstLogic) StartProcInst(req *types.StartProcInstReq) (resp *types.CommonResponse, err error) {
	s, err := l.svcCtx.Rpc.StartProcInst(l.ctx, &kernel.StartProcInstReq{
		Title:       req.Title,
		FormId:      req.FormId,
		AppId:       req.AppId,
		DataId:      req.DataId,
		DefId:       req.DefId,
		Resource:    req.Resource,
		Notifiers:   req.Notifiers,
		RemainHours: uint64(req.RemainHours),
		FlowType:    uint64(req.FlowType),
	})
	if err != nil {
		types.GetErrorCommonResponse(err.Error())
	}
	return types.GetSuccessCommonResponse(s)
}
