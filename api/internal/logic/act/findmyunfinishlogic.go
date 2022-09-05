package act

import (
	"context"

	"act/api/internal/svc"
	"act/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type FindMyUnFinishLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewFindMyUnFinishLogic(ctx context.Context, svcCtx *svc.ServiceContext) *FindMyUnFinishLogic {
	return &FindMyUnFinishLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *FindMyUnFinishLogic) FindMyUnFinish(req *types.PageReq) (resp *types.CommonResponse, err error) {
	// todo: add your logic here and delete this line

	return
}
