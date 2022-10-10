package logic

import (
	"context"

	"go-wflow/app/wflow/rpc/internal/svc"
	"go-wflow/kernel"

	"github.com/qkbyte/go-zero/core/logx"
)

type FindByDataIdLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewFindByDataIdLogic(ctx context.Context, svcCtx *svc.ServiceContext) *FindByDataIdLogic {
	return &FindByDataIdLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *FindByDataIdLogic) FindByDataId(in *kernel.DataIdReq) (*kernel.Procinst, error) {
	// todo: add your logic here and delete this line

	return &kernel.Procinst{}, nil
}
