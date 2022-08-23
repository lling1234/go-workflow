package act

import (
	"context"

	"act/api/internal/svc"
	"act/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type FindMyStartProcessLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewFindMyStartProcessLogic(ctx context.Context, svcCtx *svc.ServiceContext) *FindMyStartProcessLogic {
	return &FindMyStartProcessLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *FindMyStartProcessLogic) FindMyStartProcess(req *types.FindMyStartProcess) (resp *types.CommonResponse, err error) {
	// todo: add your logic here and delete this line

	return
}
