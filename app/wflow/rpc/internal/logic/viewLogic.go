package logic

import (
	"context"

	"go-wflow/app/wflow/rpc/internal/svc"
	"go-wflow/kernel"

	"github.com/qkbyte/go-zero/core/logx"
)

type ViewLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewViewLogic(ctx context.Context, svcCtx *svc.ServiceContext) *ViewLogic {
	return &ViewLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *ViewLogic) View(in *kernel.CompleteProcInstReq) (*kernel.Procinst, error) {
	// todo: add your logic here and delete this line

	return &kernel.Procinst{}, nil
}
