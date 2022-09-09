package logic

import (
	act2 "act/common/act"
	"act/common/act/execution"
	"act/common/act/identitylink"
	"act/common/act/task"
	"act/common/flow"
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
	if err != nil {
		return nil, err
	}
	if identityLink == nil {
		return nil, errors.New("该用户没有审批权限！")
	}
	log.Println(3333333)
	err = UpdateIdentityLink(UpdateIdentityLinkReq{
		UserId:  identityLink.UserID,
		TaskId:  taskId,
		Comment: in.Comment,
		Result:  in.Result,
		IsDeal:  1,
	}, tx, l.ctx)
	//3、更新审批人信息
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
				_, err = SaveTask(SaveTaskReq{
					NodeID:         t.NodeID,
					InstId:         in.InstId,
					IsFinished:     isInstFinish,
					Level:          t.Level,
					Step:           t.Step,
					MemberApprover: t.MemberApprover,
					Mode:           t.Mode,
				}, tx, l.ctx)
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
					_, err = SaveTask(SaveTaskReq{
						NodeID:         t.NodeID,
						InstId:         in.InstId,
						IsFinished:     isInstFinish,
						Level:          t.Level,
						Step:           t.Step,
						MemberApprover: t.MemberApprover,
						Mode:           t.Mode,
					}, tx, l.ctx)
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
	err = UpdateTask(UpdateTaskReq{
		TaskId:     taskId,
		IsFinished: isTaskFinished,
	}, tx, l.ctx)
	var flowCode string
	if approvalState == flow.HAVEPASS {
		flowCode = l.generateCode()
	}
	//8.反向更新流程实例表
	err = UpdateProcInst(UpdateProcInstReq{
		TaskId:     taskId,
		DataId:     in.DataId,
		NodeId:     "",
		State:      approvalState,
		IsFinished: isInstFinish,
		Code:       flowCode,
	}, tx, l.ctx)
	//err = UpdateProcInst("", in.GetDataId(), taskId, approvalState, isInstFinish, flowCode, tx)
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
	taskId, err := SaveTask(SaveTaskReq{
		NodeID:         nodeInfo.NodeID,
		InstId:         instId,
		IsFinished:     0,
		Level:          level,
		Step:           level,
		MemberApprover: nodeInfo.ApproverIds,
		Mode:           nodeInfo.Mode,
	}, tx, l.ctx)
	if err != nil {
		return err
	}

	//保存审批人
	userIds := strings.Split(nodeInfo.ApproverIds, ",")
	userNames := strings.Split(nodeInfo.ApproverNames, ",")
	for k, v := range userIds {
		userId, _ := strconv.ParseInt(v, 10, 64)
		err = SaveIdentityLink(SaveIdentityLinkReq{
			InstId:   instId,
			TaskId:   taskId,
			Step:     level,
			UserId:   userId,
			UserName: userNames[k],
		}, tx, l.ctx)
		if err != nil {
			return err
		}
	}

	return nil
}
