package act

import (
	"context"

	"act/api/internal/svc"
	"act/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type FindElapsedTimeLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewFindElapsedTimeLogic(ctx context.Context, svcCtx *svc.ServiceContext) *FindElapsedTimeLogic {
	return &FindElapsedTimeLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *FindElapsedTimeLogic) FindElapsedTime(req *types.UserReq) (resp *types.CommonResponse, err error) {
	// todo: add your logic here and delete this line

	return
}
