package act

import (
	"context"

	"act/api/internal/svc"
	"act/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type FindOverTimeLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewFindOverTimeLogic(ctx context.Context, svcCtx *svc.ServiceContext) *FindOverTimeLogic {
	return &FindOverTimeLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *FindOverTimeLogic) FindOverTime(req *types.UserReq) (resp *types.CommonResponse, err error) {
	// todo: add your logic here and delete this line

	return
}
