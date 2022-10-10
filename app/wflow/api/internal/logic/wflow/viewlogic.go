package wflow

import (
	"context"

	"go-wflow/app/wflow/api/internal/svc"
	"go-wflow/app/wflow/api/internal/types"

	"github.com/qkbyte/go-zero/core/logx"
)

type ViewLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewViewLogic(ctx context.Context, svcCtx *svc.ServiceContext) *ViewLogic {
	return &ViewLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *ViewLogic) View(req *types.CompleteProcInstReq) (resp *types.CommonResponse, err error) {
	// todo: add your logic here and delete this line

	return
}
