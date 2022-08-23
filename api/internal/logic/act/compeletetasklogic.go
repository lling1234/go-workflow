package act

import (
	"context"

	"act/api/internal/svc"
	"act/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type CompeleteTaskLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewCompeleteTaskLogic(ctx context.Context, svcCtx *svc.ServiceContext) *CompeleteTaskLogic {
	return &CompeleteTaskLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *CompeleteTaskLogic) CompeleteTask(req *types.CompeleteTask) (resp *types.CommonResponse, err error) {
	// todo: add your logic here and delete this line

	return
}
