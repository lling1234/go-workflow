package logic

import (
	"context"

	"go-wflow/app/wflow/rpc/internal/svc"
	"go-wflow/kernel"

	"github.com/qkbyte/go-zero/core/logx"
)

type FindOneProcDefLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewFindOneProcDefLogic(ctx context.Context, svcCtx *svc.ServiceContext) *FindOneProcDefLogic {
	return &FindOneProcDefLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *FindOneProcDefLogic) FindOneProcDef(in *kernel.IdReq) (*kernel.Procdef, error) {
	// todo: add your logic here and delete this line

	return &kernel.Procdef{}, nil
}
