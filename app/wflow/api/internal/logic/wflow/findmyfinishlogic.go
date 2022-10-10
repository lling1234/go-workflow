package wflow

import (
	"context"

	"go-wflow/app/wflow/api/internal/svc"
	"go-wflow/app/wflow/api/internal/types"

	"github.com/qkbyte/go-zero/core/logx"
)

type FindMyFinishLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewFindMyFinishLogic(ctx context.Context, svcCtx *svc.ServiceContext) *FindMyFinishLogic {
	return &FindMyFinishLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *FindMyFinishLogic) FindMyFinish(req *types.IdPageReq) (resp *types.CommonResponse, err error) {
	// todo: add your logic here and delete this line

	return
}
