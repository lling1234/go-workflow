package act

import (
	"act/api/internal/svc"
	"act/api/internal/types"
	"act/rpc/actclient"
	"context"

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
	RPC := l.svcCtx.Rpc
	newTask, err := RPC.FindLatestTaskId(l.ctx, &actclient.DataIdReq{
		DataId: req.DataId,
	})
	taskId := newTask.Id
	if err != nil {
		return types.GetErrorCommonResponse(err.Error())
	}
	identityLink, err := RPC.FindIdentityLinkByTaskId(l.ctx, &actclient.TaskIdArg{
		Id: taskId,
	})
	if err != nil {
		return types.GetErrorCommonResponse(err.Error())
	}
	if identityLink == nil {
		return types.GetErrorCommonResponse("该用户没有审批权限！")
	}
	_, err = RPC.UpdateIdentityLink(l.ctx, &actclient.IdentityLinkReq{
		UserId:  identityLink.UserId,
		TaskId:  taskId,
		Comment: req.Comment,
		Result:  int32(req.Result),
	})
	//TODO MoveStage
	return types.GetCommonResponse(err, identityLink)
}
