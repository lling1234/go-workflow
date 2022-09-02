package act

import (
	"context"

	"act/api/internal/svc"
	"act/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
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

func (l *FindMyPendingLogic) FindMyPending(req *types.PageReq) (resp *types.CommonResponse, err error) {
	// todo: add your logic here and delete this line

	return
}
