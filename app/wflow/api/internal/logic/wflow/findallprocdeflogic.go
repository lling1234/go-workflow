package wflow

import (
	"context"

	"go-wflow/app/wflow/api/internal/svc"
	"go-wflow/app/wflow/api/internal/types"

	"github.com/qkbyte/go-zero/core/logx"
)

type FindAllProcDefLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewFindAllProcDefLogic(ctx context.Context, svcCtx *svc.ServiceContext) *FindAllProcDefLogic {
	return &FindAllProcDefLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *FindAllProcDefLogic) FindAllProcDef(req *types.IdPageReq) (resp *types.CommonResponse, err error) {
	// todo: add your logic here and delete this line

	return
}
