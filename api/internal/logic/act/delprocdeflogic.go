package act

import (
	"context"

	"act/api/internal/svc"
	"act/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type DelProcdefLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewDelProcdefLogic(ctx context.Context, svcCtx *svc.ServiceContext) *DelProcdefLogic {
	return &DelProcdefLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *DelProcdefLogic) DelProcdef(req *types.FormIdReq) (resp *types.CommonResponse, err error) {
	// todo: add your logic here and delete this line

	return
}
