package act

import (
	"context"

	"act/api/internal/svc"
	"act/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
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

func (l *FindMyFinishLogic) FindMyFinish(req *types.PageReq) (resp *types.CommonResponse, err error) {
	// todo: add your logic here and delete this line

	return
}
