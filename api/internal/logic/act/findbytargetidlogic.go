package act

import (
	"context"

	"act/api/internal/svc"
	"act/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type FindByTargetIdLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewFindByTargetIdLogic(ctx context.Context, svcCtx *svc.ServiceContext) *FindByTargetIdLogic {
	return &FindByTargetIdLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *FindByTargetIdLogic) FindByTargetId(req *types.FindByTargetId) (resp *types.CommonResponse, err error) {
	// todo: add your logic here and delete this line

	return
}
