package wflow

import (
	"context"

	"go-wflow/app/wflow/api/internal/svc"
	"go-wflow/app/wflow/api/internal/types"
	"go-wflow/kernel"

	"github.com/qkbyte/go-zero/core/logx"
)

type StartViewLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewStartViewLogic(ctx context.Context, svcCtx *svc.ServiceContext) *StartViewLogic {
	return &StartViewLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *StartViewLogic) StartView(req *types.StartViewReq) (resp *types.CommonResponse, err error) {
	s, err := l.svcCtx.Rpc.StartView(l.ctx, &kernel.StartViewReq{
		FormId: req.FormId,
		AppId:  req.AppId,
		DataId: req.DataId,
	})
	if err != nil {
		types.GetErrorCommonResponse(err.Error())
	}
	return types.GetSuccessCommonResponse(s)
}
