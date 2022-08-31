package act

import (
	"context"

	"act/api/internal/svc"
	"act/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type SetActiveLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewSetActiveLogic(ctx context.Context, svcCtx *svc.ServiceContext) *SetActiveLogic {
	return &SetActiveLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *SetActiveLogic) SetActive(req *types.ProcDefIdReq) (resp *types.CommonResponse, err error) {
	// todo: add your logic here and delete this line

	return
}
