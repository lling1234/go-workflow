package act

import (
	"act/api/flow"
	"act/api/internal/svc"
	"act/api/internal/types"
	"act/rpc/actclient"
	"context"
	"errors"
	"github.com/mumushuiding/util"
	"strconv"
	"strings"

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
	err = l.Complete(req, RPC)

	return types.GetCommonResponse(err, "ok")
}

func (l *CompleteTaskLogic) Complete(req *types.CompleteTask, RPC actclient.Act) error {
	newTask, err := RPC.FindLatestTask(l.ctx, &actclient.DataIdReq{
		DataId: req.DataId,
	})
	taskId := newTask.Id
	if err != nil {
		return err
	}
	identityLink, err := RPC.FindIdentityLinkByTaskId(l.ctx, &actclient.TaskIdArg{
		Id: taskId,
	})
	if err != nil {
		return err
	}
	if identityLink == nil {
		return errors.New("该用户没有审批权限！")
	}
	_, err = RPC.UpdateIdentityLink(l.ctx, &actclient.IdentityLinkReq{
		UserId:  identityLink.UserId,
		TaskId:  taskId,
		Comment: req.Comment,
		Result:  int32(req.Result),
	})
	if err != nil {
		return err
	}
	//已通过 MoveNextStage
	if req.Result == flow.HAVEPASS {
		//exec, err := RPC.FindExecutionByInstId(l.ctx, &actclient.ProcInstIdArg{
		//	Id: identityLink.ProcInstId,
		//})
		//if err != nil {
		//	return err
		//}
		//nodeInfo := exec.NodeInfos

	}
	return nil
}

func (l *CompleteTaskLogic) MoveNextStage(RPC actclient.Act, taskId int64, instId int64, level int32, dataId int64, mode string) error {
	exec, err := RPC.FindExecutionByInstId(l.ctx, &actclient.ProcInstIdArg{
		Id: instId,
	})
	if err != nil {
		return err
	}
	var nodeInfos []*flow.NodeInfo
	util.Str2Struct(exec.NodeInfos, nodeInfos)
	nodeInfo := nodeInfos[level]
	newTask := actclient.TaskReq{
		NodeId:         nodeInfo.NodeID,
		ProcInstId:     instId,
		DataId:         dataId,
		Level:          level,
		IsFinished:     1,
		Step:           level,
		MemberApprover: nodeInfo.ApproverIds,
		Mode:           mode,
	}
	_, err = RPC.SaveTask(l.ctx, &newTask)
	if err != nil {
		return err
	}
	userIds := strings.Split(nodeInfo.ApproverIds, ",")
	userNames := strings.Split(nodeInfo.ApproverNames, ",")
	for k, v := range userIds {
		userId, _ := strconv.ParseInt(v, 10, 64)
		newIdentityLink := actclient.IdentityLinkReq{
			ProcInstId: instId,
			TaskId:     taskId,
			UserId:     userId,
			UserName:   userNames[k],
			Step:       level,
		}
		_, err = RPC.SaveIdentityLink(l.ctx, &newIdentityLink)
	}

	return nil
}
