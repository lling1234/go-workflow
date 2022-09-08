package logic

import (
	"act/api/flow"
	act2 "act/common/act"
	"act/common/act/execution"
	"act/common/act/identitylink"
	"act/common/act/procinst"
	"act/common/act/task"
	"act/rpc/constant"
	"act/rpc/general"
	"context"
	"errors"
	"github.com/mumushuiding/util"
	"log"
	"strconv"
	"strings"
	"time"

	"act/rpc/internal/svc"
	"act/rpc/types/act"

	"github.com/zeromicro/go-zero/core/logx"
)

type CompleteNormalLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewCompleteNormalLogic(ctx context.Context, svcCtx *svc.ServiceContext) *CompleteNormalLogic {
	return &CompleteNormalLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *CompleteNormalLogic) CompleteNormal(in *act.CompleteNormalProcInstReq) (*act.Nil, error) {
	tx, err := l.svcCtx.CommonStore.Tx(l.ctx)
	if err != nil {
		return nil, err
	}
	// 1.找出该数据的流程最新审批节点
	t, err := tx.Task.Query().Where(task.ProcInstID(in.InstId), task.IsDelEQ(0)).Order(act2.Desc(task.FieldStep)).First(l.ctx)
	taskId := t.ID
	if err != nil {
		return nil, err
	}
	log.Println(1111111)
	var isTaskFinished int32 = 1
	identityLink, err := l.findIdentityLinkByTaskId(taskId)
	//2、根据最新节点ID和用户ID找到对应的审批人
	//identityLink, err := RPC.FindIdentityLinkByTaskId(l.ctx, &actclient.TaskIdArg{
	//	Id: taskId,
	//})
	if err != nil {
		return nil, err
	}
	if identityLink == nil {
		return nil, errors.New("该用户没有审批权限！")
	}
	log.Println(3333333)
	err = l.updateIdentityLink(taskId, in, tx)
	//3、更新审批人信息
	//_, err = RPC.UpdateIdentityLink(l.ctx, &actclient.IdentityLinkReq{
	//	UserId:  identityLink.UserId,
	//	TaskId:  taskId,
	//	Comment: req.Comment,
	//	Result:  int32(req.Result),
	//	IsDeal:  1,
	//})
	if err != nil {
		return nil, err
	}
	log.Println(44444444)
	//4、根据审批状态和是否会签判断 a.继续当前节点审批 b.下一节点审批 c.驳回到上一节点审批 d.驳回到指定节点审批 e.直接结束流程
	//已通过 MoveNextStage
	//5.审批通过 生成新的流程节点和新的审批人表
	var isInstFinish int32 = 0
	var approvalState int32 = 0
	switch in.Result {
	case flow.HAVEPASS:
		nodeInfos, err := l.findNodeInfosByInstId(in.GetInstId())
		log.Println(5555555)
		if err != nil {
			return nil, err
		}
		if strings.ToLower(t.Mode) == "or" {
			if t.Level == int32(len(nodeInfos)-1) {
				isInstFinish = 1
				approvalState = flow.HAVEPASS
				_, err = l.saveTask(t.NodeID, in.GetInstId(), isInstFinish, t.Level, t.Level, t.MemberApprover, t.Mode, tx)
				if err != nil {
					return nil, err
				}
			} else {
				approvalState = flow.DEALING
				err = l.moveNextStage(in.GetInstId(), t.Level, nodeInfos, tx)
				if err != nil {
					return nil, err
				}
			}
		} else if strings.ToLower(t.Mode) == "and" {
			//1.判断是否该节点的所有人都已经审批
			log.Println("andand")
			//2.是否为末级节点
			if l.isTaskFinished(t.MemberApprover, t.AgreeApprover, identityLink.UserID) {
				if t.Level == int32(len(nodeInfos)-1) {
					isInstFinish = 1
					approvalState = flow.HAVEPASS
					_, err = l.saveTask(t.NodeID, in.GetInstId(), isInstFinish, t.Level, t.Level, t.MemberApprover, t.Mode, tx)
					if err != nil {
						return nil, err
					}
				} else {
					log.Println("andmoveNextStage")
					approvalState = flow.DEALING
					err = l.moveNextStage(in.GetInstId(), t.Level, nodeInfos, tx)
					if err != nil {
						return nil, err
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
		log.Println("req.Result", in.Result)
		isInstFinish = 1
		approvalState = flow.NOTPASS
		//7.驳回  驳回到上一节点审批 驳回到指定节点审批 直接结束流程
	case flow.REJECT:
		approvalState = flow.REJECT
	}

	err = l.updateTask(taskId, isTaskFinished, tx)
	//_, err = RPC.UpdateTask(l.ctx, &actclient.TaskReq{
	//	Id:         taskId,
	//	IsFinished: isTaskFinished,
	//})
	//log.Println("approvalState", approvalState)
	var flowCode string
	if approvalState == flow.HAVEPASS {
		flowCode = l.generateCode()
	}
	//8.反向更新流程实例表
	//_, err = RPC.UpdateProcInst(l.ctx, &actclient.UpdateProcInstReq{
	//	DataId:   req.DataId,
	//	State:    approvalState,
	//	IsFinish: isInstFinish,
	//	Code:     flowCode,
	//	TaskId:   taskId,
	//})
	err = l.UpdateProcInst("", in.GetDataId(), taskId, approvalState, isInstFinish, flowCode, tx)
	return &act.Nil{}, err
}

func (l *CompleteNormalLogic) findLatestTask(instId int64) (*act2.Task, error) {
	t, err := l.svcCtx.CommonStore.Task.Query().Where(task.ProcInstID(instId), task.IsDelEQ(0)).Order(act2.Desc(task.FieldStep)).First(l.ctx)
	return t, err
}

func (l *CompleteNormalLogic) findIdentityLinkByTaskId(taskId int64) (*act2.IdentityLink, error) {

	ils, err := l.svcCtx.CommonStore.IdentityLink.Query().
		Where(identitylink.TaskIDEQ(taskId), identitylink.UserIDEQ(general.MyUserId), identitylink.IsDelEQ(0), identitylink.IsDealEQ(0)).
		Select(identitylink.FieldID, identitylink.FieldUserID, identitylink.FieldUserName, identitylink.FieldStep, identitylink.FieldProcInstID, identitylink.FieldTaskID, identitylink.FieldTargetID).
		All(l.ctx)
	if err != nil {
		return nil, err
	}
	if ils == nil || len(ils) == 0 {
		return nil, errors.New("用户没有审批权限")
	}
	return ils[0], nil
}

func (l *CompleteNormalLogic) updateIdentityLink(taskId int64, in *act.CompleteNormalProcInstReq, tx *act2.Tx) error {
	err := tx.IdentityLink.Update().Where(identitylink.TaskIDEQ(taskId), identitylink.UserIDEQ(general.MyUserId), identitylink.IsDelEQ(0), identitylink.IsDealEQ(0)).
		SetUpdateTime(time.Now()).SetComment(in.Comment).SetResult(in.Result).SetIsDeal(1).Exec(l.ctx)

	return err
}

func (l *CompleteNormalLogic) findNodeInfosByInstId(instId int64) ([]*flow.NodeInfo, error) {
	exec, err := l.svcCtx.CommonStore.Execution.Query().Where(execution.ProcInstID(instId), execution.IsDelEQ(0)).First(l.ctx)
	if err != nil {
		return nil, err
	}
	//3.根据执行流程找到对应的流程定义节点
	var nodeInfos []*flow.NodeInfo
	util.Str2Struct(exec.NodeInfos, &nodeInfos)
	return nodeInfos, err
}

func (l *CompleteNormalLogic) saveTask(nodeId string, instId int64, isFinish int32, level int32, step int32, memberApprover string, mode string, tx *act2.Tx) (int64, error) {
	taskCreate := tx.Task.Create().SetNodeID(nodeId).SetProcInstID(instId).
		SetIsFinished(isFinish).SetLevel(level).SetStep(step).SetMemberApprover(memberApprover)
	if mode != "" {
		taskCreate.SetMode(mode)
	}
	t, err := taskCreate.Save(l.ctx)
	return t.ID, err
}

func (l *CompleteNormalLogic) saveIdentityLink(instId int64, taskId int64, step int32, userId int64, userName string, tx *act2.Tx) error {
	_, err := tx.IdentityLink.Create().SetProcInstID(instId).SetTaskID(taskId).SetTargetID(general.TargetId).
		SetIsDeal(0).SetStep(step).SetUserID(userId).SetUserName(userName).Save(l.ctx)

	return err
}

func (l *CompleteNormalLogic) generateCode() string {
	return "ORGINONE" + time.Now().Format("20060102")
}

func (l *CompleteNormalLogic) isTaskFinished(memberStr string, agreerStr string, userId int64) bool {
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

func (l *CompleteNormalLogic) moveNextStage(instId int64, level int32, nodeInfos []*flow.NodeInfo, tx *act2.Tx) error {
	nodeInfo := nodeInfos[level-1] //计数是从0开始，因此需要-1
	//4.保存流程任务
	taskId, err := l.saveTask(nodeInfo.NodeID, instId, 0, level, level, nodeInfo.ApproverIds, nodeInfo.Mode, tx)
	if err != nil {
		return err
	}

	//保存审批人
	userIds := strings.Split(nodeInfo.ApproverIds, ",")
	userNames := strings.Split(nodeInfo.ApproverNames, ",")
	for k, v := range userIds {
		userId, _ := strconv.ParseInt(v, 10, 64)
		err = l.saveIdentityLink(instId, taskId, level, userId, userNames[k], tx)
		if err != nil {
			return err
		}
	}

	return nil
}

func (l *CompleteNormalLogic) updateTask(taskId int64, isFinished int32, tx *act2.Tx) error {
	oldTask, err := tx.Task.Query().Where(task.IDEQ(taskId)).First(l.ctx)
	if err != nil {
		return err
	}
	agreers := ""
	if oldTask.AgreeApprover == "" {
		agreers = strconv.FormatInt(general.MyUserId, 10)
	} else {
		agreers = oldTask.AgreeApprover + "," + strconv.FormatInt(general.MyUserId, 10)
	}
	err = tx.Task.Update().Where(task.IDEQ(taskId)).SetUpdateTime(time.Now()).SetIsFinished(isFinished).SetAgreeApprover(agreers).SetClaimTime(time.Now()).Exec(l.ctx)

	return err
}

func (l *CompleteNormalLogic) UpdateProcInst(nodeId string, dataId int64, taskId int64, state int32, isFinished int32, code string, tx *act2.Tx) error {
	procInstUpdate := tx.ProcInst.Update().Where(procinst.DataIDEQ(dataId), procinst.StateNotIn(constant.WITHDRAW, constant.DISCARD), procinst.IsDelEQ(0))

	if taskId != 0 {
		procInstUpdate.SetNodeID(nodeId).SetTaskID(taskId)
	}
	if state != 0 {
		procInstUpdate.SetState(state)
	}
	if state == constant.HAVEPASS {
		procInstUpdate.SetCode(code)
	}
	if isFinished == 1 {
		procInstUpdate.SetIsFinished(1).SetEndTime(time.Now())
	}
	err := procInstUpdate.SetUpdateTime(time.Now()).SetUpdateUserID(general.MyUserId).Exec(l.ctx)

	return err
}
