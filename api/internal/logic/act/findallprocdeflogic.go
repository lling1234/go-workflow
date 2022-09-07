package act

import (
	"context"

	"act/api/internal/svc"
	"act/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
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

func (l *FindAllProcDefLogic) FindAllProcDef(req *types.FormIdReq) (resp *types.CommonResponse, err error) {
	// todo: add your logic here and delete this line

	return
}
