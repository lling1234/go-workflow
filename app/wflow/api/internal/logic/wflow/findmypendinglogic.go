package wflow

import (
	"context"

	"go-wflow/app/wflow/api/internal/svc"
	"go-wflow/app/wflow/api/internal/types"

	"github.com/qkbyte/go-zero/core/logx"
)

type FindMyPendingLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewFindMyPendingLogic(ctx context.Context, svcCtx *svc.ServiceContext) *FindMyPendingLogic {
	return &FindMyPendingLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *FindMyPendingLogic) FindMyPending(req *types.IdPageReq) (resp *types.CommonResponse, err error) {
	// todo: add your logic here and delete this line

	return
}
