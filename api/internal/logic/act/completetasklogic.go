package act

import (
	"act/api/flow"
	"act/api/internal/svc"
	"act/api/internal/types"
	"act/rpc/actclient"
	"context"
	"errors"
	"github.com/mumushuiding/util"
	"log"
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
	// 1.找出该数据的流程最新审批节点
	task, err := RPC.FindLatestTask(l.ctx, &actclient.DataIdReq{
		DataId: req.DataId,
	})
	taskId := task.Id
	if err != nil {
		return err
	}
	//2、根据最新节点ID和用户ID找到对应的审批人
	identityLink, err := RPC.FindIdentityLinkByTaskId(l.ctx, &actclient.TaskIdArg{
		Id: taskId,
	})
	if err != nil {
		return err
	}
	if identityLink == nil {
		return errors.New("该用户没有审批权限！")
	}
	//3、更新审批人信息
	_, err = RPC.UpdateIdentityLink(l.ctx, &actclient.IdentityLinkReq{
		UserId:  identityLink.UserId,
		TaskId:  taskId,
		Comment: req.Comment,
		Result:  int32(req.Result),
	})
	if err != nil {
		return err
	}
	//4、根据审批状态和是否会签判断 a.继续当前节点审批 b.下一节点审批 c.驳回到上一节点审批 d.驳回到指定节点审批 e.直接结束流程
	//已通过 MoveNextStage
	//5.审批通过 生成新的流程节点和新的审批人表
	isInstFinish := false
	var approvalState int32 = 0
	if req.Result == flow.HAVEPASS {
		nodeInfos, err := l.findNodeInfosByInstId(RPC, task.ProcInstId)
		if err != nil {
			return err
		}
		if task.Mode == "or" {
			if task.Level == int32(len(nodeInfos)-1) {
				isInstFinish = true
				approvalState = flow.HAVEPASS
			}
		} else if task.Mode == "and" {
			//1.判断是否该节点的所有人都已经审批 TODO 返回数组不会写，暂时写只有一个审批人的
			//2.是否为末级节点
			if task.Level == int32(len(nodeInfos)-1) {
				isInstFinish = true
			} else {
				approvalState = flow.DEALING
				log.Printf("task.Level:%d", task.Level)
				err = l.MoveNextStage(RPC, task.Id, task.ProcInstId, task.Level, req.DataId, task.Mode, nodeInfos)
				if err != nil {
					return err
				}
			}
		}
		//6、审批不通过 直接结束流程 更新流程实例 is_finished=1 state = 7
	} else if req.Result == flow.DISCARD {
		isInstFinish = true
		approvalState = flow.DISCARD
		//7.驳回  驳回到上一节点审批 驳回到指定节点审批 直接结束流程
	} else if req.Result == flow.REJECT {
		approvalState = flow.REJECT
	}
	//8.反向更新流程实例表
	if isInstFinish {
		_, err = RPC.UpdateProcInst(l.ctx, &actclient.UpdateProcInstReq{
			DataId: req.DataId,
			State:  approvalState,
		})
	}
	return nil
}

func (l *CompleteTaskLogic) findNodeInfosByInstId(RPC actclient.Act, instId int64) ([]*flow.NodeInfo, error) {
	exec, err := RPC.FindExecutionByInstId(l.ctx, &actclient.ProcInstIdArg{
		Id: instId,
	})
	if err != nil {
		return nil, err
	}
	//3.根据执行流程找到对应的流程定义节点
	var nodeInfos []*flow.NodeInfo
	util.Str2Struct(exec.NodeInfos, &nodeInfos)
	return nodeInfos, err
}
func (l *CompleteTaskLogic) MoveNextStage(RPC actclient.Act, taskId int64, instId int64, level int32, dataId int64, mode string, nodeInfos []*flow.NodeInfo) error {
	//1.判断是否已经是末级节点
	//2.根据流程实例ID找到对应的执行流
	//nodeInfos, err := l.findNodeInfosByInstId(RPC, instId)
	//if err != nil {
	//	return err
	//}
	////3.根据执行流程找到对应的流程定义节点
	//var nodeInfos []*flow.NodeInfo
	//util.Str2Struct(exec.NodeInfos, nodeInfos)
	nodeInfo := nodeInfos[level]
	//4.保存流程任务
	newTask := actclient.TaskReq{
		NodeId:         nodeInfo.NodeID,
		ProcInstId:     instId,
		DataId:         dataId,
		Level:          level,
		IsFinished:     0,
		Step:           level,
		MemberApprover: nodeInfo.ApproverIds,
		Mode:           mode,
	}
	_, err := RPC.SaveTask(l.ctx, &newTask)
	if err != nil {
		return err
	}
	//保存审批人
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
