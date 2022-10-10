package logic

import (
	"context"

	"go-wflow/app/wflow/rpc/internal/svc"
	"go-wflow/kernel"

	"github.com/qkbyte/go-zero/core/logx"
)

type FindAllProcDefLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewFindAllProcDefLogic(ctx context.Context, svcCtx *svc.ServiceContext) *FindAllProcDefLogic {
	return &FindAllProcDefLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *FindAllProcDefLogic) FindAllProcDef(in *kernel.PageReq) (*kernel.ProcdefArray, error) {
	// todo: add your logic here and delete this line

	return &kernel.ProcdefArray{}, nil
}
