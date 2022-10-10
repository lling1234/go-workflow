package wflow

import (
	"context"

	"go-wflow/app/wflow/api/internal/svc"
	"go-wflow/app/wflow/api/internal/types"
	"go-wflow/kernel"

	"github.com/qkbyte/go-zero/core/logx"
)

type SetActiveLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewSetActiveLogic(ctx context.Context, svcCtx *svc.ServiceContext) *SetActiveLogic {
	return &SetActiveLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *SetActiveLogic) SetActive(req *types.QueryProcDefReq) (resp *types.CommonResponse, err error) {
	s,err:=l.svcCtx.Rpc.SetActive(l.ctx, &kernel.QueryProcDefReq{
		FormId:  req.FormId,
		AppId:   req.AppId,
		Version: req.Version,
	})
	if err != nil {
		types.GetErrorCommonResponse(err.Error())
	}
	return types.GetSuccessCommonResponse(s)
}
