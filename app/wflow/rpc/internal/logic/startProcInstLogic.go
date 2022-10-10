package logic

import (
	"context"

	"go-wflow/app/wflow/rpc/internal/svc"
	"go-wflow/kernel"
	"go-wflow/kernel/core/wflow"

	"github.com/qkbyte/go-zero/core/logx"
)

type StartProcInstLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
	WflowInterface wflow.WflowInterface
}

func NewStartProcInstLogic(ctx context.Context, svcCtx *svc.ServiceContext) *StartProcInstLogic {
	return &StartProcInstLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
		// TODO personId 从token中获取
		WflowInterface: wflow.New(111, svcCtx.CommonStore.Client, ctx),
	}
}

func (l *StartProcInstLogic) StartProcInst(in *kernel.StartProcInstReq) (*kernel.Nil, error) {
	
	return l.WflowInterface.StartProcess(in)
}
