package act

import (
	"act/rpc/actclient"
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
	taskId, err := l.svcCtx.Rpc.FindLatestTaskId(l.ctx, &actclient.DataIdReq{
		DataId: req.DataId,
	})

	return types.GetCommonResponse(err, taskId)
}
