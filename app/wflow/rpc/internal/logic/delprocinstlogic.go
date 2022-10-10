package logic

import (
	"context"

	"go-wflow/app/wflow/rpc/internal/svc"
	"go-wflow/kernel"

	"github.com/qkbyte/go-zero/core/logx"
)

type DelProcInstLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewDelProcInstLogic(ctx context.Context, svcCtx *svc.ServiceContext) *DelProcInstLogic {
	return &DelProcInstLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *DelProcInstLogic) DelProcInst(in *kernel.DataIdReq) (*kernel.Nil, error) {
	// todo: add your logic here and delete this line

	return &kernel.Nil{}, nil
}
