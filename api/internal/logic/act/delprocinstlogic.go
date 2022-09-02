package act

import (
	"context"

	"act/api/internal/svc"
	"act/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type DelProcinstLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewDelProcinstLogic(ctx context.Context, svcCtx *svc.ServiceContext) *DelProcinstLogic {
	return &DelProcinstLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *DelProcinstLogic) DelProcinst(req *types.DataIdReq) (resp *types.CommonResponse, err error) {
	// todo: add your logic here and delete this line

	return
}
