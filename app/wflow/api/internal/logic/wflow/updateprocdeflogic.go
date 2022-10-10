package wflow

import (
	"context"

	"go-wflow/app/wflow/api/internal/svc"
	"go-wflow/app/wflow/api/internal/types"

	"github.com/qkbyte/go-zero/core/logx"
)

type UpdateProcDefLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewUpdateProcDefLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UpdateProcDefLogic {
	return &UpdateProcDefLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *UpdateProcDefLogic) UpdateProcDef(req *types.SaveProcDefReq) (resp *types.CommonResponse, err error) {
	// todo: add your logic here and delete this line

	return
}
