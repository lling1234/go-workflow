package logic

import (
	"context"

	"go-wflow/app/wflow/rpc/internal/svc"
	"go-wflow/kernel"

	"github.com/qkbyte/go-zero/core/logx"
)

type UpdateProcDefLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewUpdateProcDefLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UpdateProcDefLogic {
	return &UpdateProcDefLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *UpdateProcDefLogic) UpdateProcDef(in *kernel.SaveProcDefReq) (*kernel.Procdef, error) {
	// todo: add your logic here and delete this line

	return &kernel.Procdef{}, nil
}
