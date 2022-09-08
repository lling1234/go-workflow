package act

import (
	"act/api/flow"
	"act/api/internal/svc"
	"act/api/internal/types"
	"act/rpc/actclient"
	"context"
	"errors"
	"github.com/mumushuiding/util"
	"github.com/zeromicro/go-zero/core/logx"
	"log"
	"strconv"
	"strings"
	"time"
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
	inst, err := RPC.FindProcInstByDataId(l.ctx, &actclient.DataIdReq{
		DataId: req.DataId,
	})
	if err != nil {
		return types.GetErrorCommonResponse(err.Error())
	}
	err = l.Complete(req, RPC, inst.Id)

	return types.GetCommonResponse(err, "ok")
}

func (l *CompleteTaskLogic) Complete(req *types.CompleteTask, RPC actclient.Act, instId int64) error {
	// 1.找出该数据的流程最新审批节点
	task, err := RPC.FindLatestTask(l.ctx, &actclient.ProcInstIdArg{
		Id: instId,
	})
	taskId := task.Id
	if err != nil {
		return err
	}
	log.Println(1111111)
	var isTaskFinished int32 = 1
	if err != nil {
		return err
	}
	log.Println(2222222)
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
	log.Println(3333333)
	//3、更新审批人信息
	_, err = RPC.UpdateIdentityLink(l.ctx, &actclient.IdentityLinkReq{
		UserId:  identityLink.UserId,
		TaskId:  taskId,
		Comment: req.Comment,
		Result:  int32(req.Result),
		IsDeal:  1,
	})
	if err != nil {
		return err
	}
	log.Println(44444444)
	//4、根据审批状态和是否会签判断 a.继续当前节点审批 b.下一节点审批 c.驳回到上一节点审批 d.驳回到指定节点审批 e.直接结束流程
	//已通过 MoveNextStage
	//5.审批通过 生成新的流程节点和新的审批人表
	var isInstFinish int32 = 0
	var approvalState int32 = 0
	switch req.Result {
	case flow.HAVEPASS:
		nodeInfos, err := l.findNodeInfosByInstId(RPC, task.ProcInstId)
		log.Println(5555555)
		if err != nil {
			return err
		}
		if strings.ToLower(task.Mode) == "or" {
			if task.Level == int32(len(nodeInfos)-1) {
				isInstFinish = 1
				approvalState = flow.HAVEPASS
				_, err = l.finishTask(RPC, task.ProcInstId, task.Level+1, req.DataId)
				if err != nil {
					return err
				}
			} else {
				approvalState = flow.DEALING
				err = l.moveNextStage(RPC, task.ProcInstId, task.Level+1, req.DataId, nodeInfos)
				if err != nil {
					return err
				}
			}
		} else if strings.ToLower(task.Mode) == "and" {
			//1.判断是否该节点的所有人都已经审批
			log.Println("andand")
			//2.是否为末级节点
			if l.isTaskFinished(task.MemberApprover, task.AgreeApprover, identityLink.UserId) {
				if task.Level == int32(len(nodeInfos)-1) {
					isInstFinish = 1
					approvalState = flow.HAVEPASS
					_, err = l.finishTask(RPC, task.ProcInstId, task.Level+1, req.DataId)
					if err != nil {
						return err
					}
				} else {
					log.Println("andmoveNextStage")
					approvalState = flow.DEALING
					err = l.moveNextStage(RPC, task.ProcInstId, task.Level+1, req.DataId, nodeInfos)
					if err != nil {
						return err
					}
				}
			} else {
				log.Println("andDEALING")
				approvalState = flow.DEALING
				isTaskFinished = 0
			}

		}
		//6、审批不通过 直接结束流程 更新流程实例 is_finished=1 state = 7
	case flow.NOTPASS:
		log.Println("req.Result", req.Result)
		isInstFinish = 1
		approvalState = flow.NOTPASS
		//7.驳回  驳回到上一节点审批 驳回到指定节点审批 直接结束流程
	case flow.REJECT:
		approvalState = flow.REJECT
	}

	_, err = RPC.UpdateTask(l.ctx, &actclient.TaskReq{
		Id:         taskId,
		IsFinished: isTaskFinished,
	})
	//log.Println("approvalState", approvalState)
	var flowCode string
	if approvalState == flow.HAVEPASS {
		flowCode = l.generateCode()
	}
	//8.反向更新流程实例表
	_, err = RPC.UpdateProcInst(l.ctx, &actclient.UpdateProcInstReq{
		DataId:   req.DataId,
		State:    approvalState,
		IsFinish: isInstFinish,
		Code:     flowCode,
		TaskId:   taskId,
	})
	return nil
}

func (l *CompleteTaskLogic) generateCode() string {
	return "ORGINONE" + time.Now().Format("20060102")
}

func (l *CompleteTaskLogic) isTaskFinished(memberStr string, agreerStr string, userId int64) bool {
	userIdStr := strconv.FormatInt(userId, 10)
	members := strings.Split(memberStr, ",")
	if agreerStr == "" {
		if len(members) == 1 {
			return members[0] == userIdStr
		}
	}
	agreers := strings.Split(agreerStr+","+userIdStr, ",")
	flag := true
	if len(members) == len(agreers) {
		for _, v := range agreers {
			flag = strings.Contains(memberStr, v)
			if !flag {
				break
			}
		}
	} else {
		flag = false
	}
	return flag
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

func (l *CompleteTaskLogic) moveNextStage(RPC actclient.Act, instId int64, level int32, dataId int64, nodeInfos []*flow.NodeInfo) error {
	nodeInfo := nodeInfos[level-1] //计数是从0开始，因此需要-1
	//4.保存流程任务
	task := actclient.TaskReq{
		NodeId:         nodeInfo.NodeID,
		ProcInstId:     instId,
		Level:          level,
		IsFinished:     0,
		Step:           level,
		MemberApprover: nodeInfo.ApproverIds,
		Mode:           nodeInfo.Mode,
	}
	newTask, err := RPC.SaveTask(l.ctx, &task)
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
			TaskId:     newTask.Id,
			UserId:     userId,
			UserName:   userNames[k],
			Step:       level,
		}
		_, err = RPC.SaveIdentityLink(l.ctx, &newIdentityLink)
	}

	return nil
}
func (l *CompleteTaskLogic) finishTask(RPC actclient.Act, instId int64, level int32, dataId int64) (int64, error) {
	task := actclient.TaskReq{
		NodeId:     "结束",
		ProcInstId: instId,
		Level:      level,
		IsFinished: 1,
		Step:       level,
	}
	log.Println(666666)
	newTask, err := RPC.SaveTask(l.ctx, &task)
	log.Println(7777777)
	if newTask == nil {
		return 0, err
	}
	return newTask.Id, err
}
