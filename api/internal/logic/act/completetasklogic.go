package act

import (
	"act/rpc/actclient"
	"context"

	"act/api/internal/svc"
	"act/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type CompleteTaskLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewCompleteTaskLogic(ctx context.Context, svcCtx *svc.ServiceContext) *CompleteTaskLogic {
	return &CompleteTaskLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *CompleteTaskLogic) CompleteTask(req *types.CompleteTask) (resp *types.CommonResponse, err error) {
	taskId, err := l.svcCtx.Rpc.FindLatestTaskId(l.ctx, &actclient.DataIdReq{
		DataId: req.DataId,
	})

	return types.GetCommonResponse(err, taskId)
}
