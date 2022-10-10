package logic

import (
	"context"

	"go-wflow/app/wflow/rpc/internal/svc"
	"go-wflow/kernel"
	"go-wflow/kernel/core/wflow"

	"github.com/qkbyte/go-zero/core/logx"
)

type StartViewLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
	WflowInterface wflow.WflowInterface
}

func NewStartViewLogic(ctx context.Context, svcCtx *svc.ServiceContext) *StartViewLogic {
	return &StartViewLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
		// TODO personId 从token中获取
		WflowInterface: wflow.New(111, svcCtx.CommonStore.Client, ctx),
	}
}

func (l *StartViewLogic) StartView(in *kernel.StartViewReq) (*kernel.StartViewResp, error) {
	return 	l.WflowInterface.StartView(in)

}
