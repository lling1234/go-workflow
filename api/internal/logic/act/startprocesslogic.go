package act

import (
	"context"

	"act/api/internal/svc"
	"act/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type StartProcessLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewStartProcessLogic(ctx context.Context, svcCtx *svc.ServiceContext) *StartProcessLogic {
	return &StartProcessLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *StartProcessLogic) StartProcess(req *types.StartProcess) (resp *types.CommonResponse, err error) {
	// todo: add your logic here and delete this line

	return
}
