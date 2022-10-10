package logic

import (
	"context"

	"go-wflow/app/wflow/rpc/internal/svc"
	"go-wflow/kernel"
	"go-wflow/kernel/core/wflow"

	"github.com/qkbyte/go-zero/core/logx"
)

type CompleteLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
	WflowInterface wflow.WflowInterface
}

func NewCompleteLogic(ctx context.Context, svcCtx *svc.ServiceContext) *CompleteLogic {
	return &CompleteLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
		// TODO personId 从token中获取
		WflowInterface: wflow.New(111, svcCtx.CommonStore.Client, ctx),
	}
}

func (l *CompleteLogic) Complete(in *kernel.CompleteProcInstReq) (*kernel.Nil, error) {

	return l.WflowInterface.CompleteProcess(in)
}
