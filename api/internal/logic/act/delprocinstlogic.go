package act

import (
	"context"

	"act/api/internal/svc"
	"act/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type DelProcInstLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewDelProcInstLogic(ctx context.Context, svcCtx *svc.ServiceContext) *DelProcInstLogic {
	return &DelProcInstLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *DelProcInstLogic) DelProcInst(req *types.DataIdReq) (resp *types.CommonResponse, err error) {
	// todo: add your logic here and delete this line

	return
}
