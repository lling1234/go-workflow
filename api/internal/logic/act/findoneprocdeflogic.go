package act

import (
	"context"

	"act/api/internal/svc"
	"act/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type FindOneProcdefLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewFindOneProcdefLogic(ctx context.Context, svcCtx *svc.ServiceContext) *FindOneProcdefLogic {
	return &FindOneProcdefLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *FindOneProcdefLogic) FindOneProcdef(req *types.FormIdReq) (resp *types.CommonResponse, err error) {
	// todo: add your logic here and delete this line

	return
}
