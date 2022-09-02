package act

import (
	"context"

	"act/api/internal/svc"
	"act/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type FindAllProcessLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewFindAllProcessLogic(ctx context.Context, svcCtx *svc.ServiceContext) *FindAllProcessLogic {
	return &FindAllProcessLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *FindAllProcessLogic) FindAllProcess(req *types.SearchProcess) (resp *types.CommonResponse, err error) {
	// todo: add your logic here and delete this line

	return
}
