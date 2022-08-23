package act

import (
	"context"

	"act/api/internal/svc"
	"act/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type FindMyApproveProcessLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewFindMyApproveProcessLogic(ctx context.Context, svcCtx *svc.ServiceContext) *FindMyApproveProcessLogic {
	return &FindMyApproveProcessLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *FindMyApproveProcessLogic) FindMyApproveProcess(req *types.FindMyApproveProcess) (resp *types.CommonResponse, err error) {
	// todo: add your logic here and delete this line

	return
}
