package act

import (
	"context"

	"act/api/internal/svc"
	"act/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type FindOneProcDefLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewFindOneProcDefLogic(ctx context.Context, svcCtx *svc.ServiceContext) *FindOneProcDefLogic {
	return &FindOneProcDefLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *FindOneProcDefLogic) FindOneProcDef(req *types.FormIdReq) (resp *types.CommonResponse, err error) {
	// todo: add your logic here and delete this line

	return
}
