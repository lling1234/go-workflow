package logic

import (
	"context"

	"go-wflow/app/wflow/rpc/internal/svc"
	"go-wflow/kernel"

	"github.com/qkbyte/go-zero/core/logx"
)

type DelProcDefLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewDelProcDefLogic(ctx context.Context, svcCtx *svc.ServiceContext) *DelProcDefLogic {
	return &DelProcDefLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *DelProcDefLogic) DelProcDef(in *kernel.IdReq) (*kernel.Nil, error) {
	// todo: add your logic here and delete this line

	return &kernel.Nil{}, nil
}
