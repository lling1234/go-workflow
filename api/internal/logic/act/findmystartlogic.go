package act

import (
	"context"

	"act/api/internal/svc"
	"act/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type FindMyStartLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewFindMyStartLogic(ctx context.Context, svcCtx *svc.ServiceContext) *FindMyStartLogic {
	return &FindMyStartLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *FindMyStartLogic) FindMyStart(req *types.PageReq) (resp *types.CommonResponse, err error) {
	// todo: add your logic here and delete this line

	return
}
