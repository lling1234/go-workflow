package wflow

import (
	"context"

	"go-wflow/app/wflow/api/internal/svc"
	"go-wflow/app/wflow/api/internal/types"

	"github.com/qkbyte/go-zero/core/logx"
)

type FindMyApprovedLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewFindMyApprovedLogic(ctx context.Context, svcCtx *svc.ServiceContext) *FindMyApprovedLogic {
	return &FindMyApprovedLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *FindMyApprovedLogic) FindMyApproved(req *types.IdPageReq) (resp *types.CommonResponse, err error) {
	// todo: add your logic here and delete this line

	return
}
