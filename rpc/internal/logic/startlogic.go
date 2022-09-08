package logic

import (
	"act/api/flow"
	act2 "act/common/act"
	"act/common/act/procinst"
	"act/rpc/constant"
	"act/rpc/general"
	"container/list"
	"context"
	"github.com/mumushuiding/util"
	"strconv"
	"strings"
	"time"

	"act/rpc/internal/svc"
	"act/rpc/types/act"

	"github.com/zeromicro/go-zero/core/logx"
)

type StartLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewStartLogic(ctx context.Context, svcCtx *svc.ServiceContext) *StartLogic {
	return &StartLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *StartLogic) Start(in *act.StartProcInstReq) (*act.Nil, error) {
	tx, err := l.svcCtx.CommonStore.Tx(l.ctx)
	if err != nil {
		return nil, err
	}
	inst, err := l.saveProcInst(in.DefId, in.DataId, in.Title, in.RemainHours, tx)
	var STEP int32 = 1
	//task := actclient.TaskReq{
	//	NodeId:     "开始",
	//	ProcInstId: inst.Id,
	//	Level:      STEP,
	//	IsFinished: 1,
	//	Step:       STEP,
	//	Mode:       "or",
	//}
	_, err = l.saveTask("开始", inst.ID, 1, STEP, STEP, "", "or", tx)
	//_, err = RPC.SaveTask(l.ctx, &task)
	//log.Println(33333)
	//if err != nil {
	//	return types.GetErrorCommonResponse(err.Error())
	//}
	var node *flow.Node
	err = util.Str2Struct(in.Resource, &node)
	if err != nil {
		return nil, err
	}
	list, err := flow.ParseProcessConfig(node)
	listStr, err := l.GenerateExec(list)
	exec, err := l.SaveExecution(inst.ID, listStr, tx)
	//exec := &actclient.ExecutionReq{
	//	ProcDefId:  def.Id,
	//	ProcInstId: inst.Id,
	//	NodeInfos:  listStr,
	//}
	//_, err = RPC.SaveExecution(l.ctx, exec)
	//if err != nil {
	//	return types.GetErrorCommonResponse(err.Error())
	//}
	// 获取执行流信息
	var nodeInfos []*flow.NodeInfo
	util.Str2Struct(exec.NodeInfos, &nodeInfos)
	// -----------------生成新任务-------------
	firstNode := nodeInfos[1]
	//if firstNode.Mode == "AND" {
	//task.MemberCount = int32(firstNode.MemberCount)
	//}
	//task.NodeId = firstNode.NodeID
	STEP = 2
	//task.AgreeNum = 0
	//task.IsFinished = 0
	//task.Level = STEP
	//task.Step = STEP
	//task.Mode = firstNode.Mode
	//task.MemberApprover = firstNode.ApproverIds
	//newTask, err := RPC.SaveTask(l.ctx, &task)
	//if err != nil {
	//	return types.GetErrorCommonResponse(err.Error())
	//}
	secondTaskId, err := l.saveTask(firstNode.NodeID, inst.ID, 0, STEP, STEP, firstNode.ApproverIds, firstNode.Mode, tx)
	//newTaskReply, err := RPC.FindLatestTask(l.ctx, &actclient.DataIdReq{
	//	DataId: req.DataId,
	//	Step:   STEP,
	//})
	//if err != nil {
	//	return types.GetErrorCommonResponse(err.Error())
	//}
	err = l.UpdateProcInst(firstNode.NodeID, in.DataId, secondTaskId, flow.PENDING, 0, "", tx)
	//updateInst := &actclient.UpdateProcInstReq{
	//	ProcDefId:   def.Id,
	//	FormId:      req.FormId,
	//	DataId:      req.DataId,
	//	RemainHours: def.RemainHours,
	//	State:       flow.PENDING,
	//	TaskId:      newTask.Id,
	//	NodeId:      firstNode.NodeID,
	//}
	//_, err = RPC.UpdateProcInst(l.ctx, updateInst)
	//if err != nil {
	//	return types.GetErrorCommonResponse(err.Error())
	//}
	userIds := strings.Split(firstNode.ApproverIds, ",")
	userNames := strings.Split(firstNode.ApproverNames, ",")
	for k, v := range userIds {
		userId, _ := strconv.ParseInt(v, 10, 64)
		err = l.saveIdentityLink(inst.ID, secondTaskId, STEP, userId, userNames[k], tx)
		//identityLink := actclient.IdentityLinkReq{
		//	ProcInstId: inst.Id,
		//	TaskId:     newTask.Id,
		//	UserId:     userId,
		//	UserName:   userNames[k],
		//	Step:       STEP,
		//}
		//_, err = RPC.SaveIdentityLink(l.ctx, &identityLink)
	}

	return &act.Nil{}, nil
}

func (l *StartLogic) GenerateExec(list *list.List) (string, error) {
	list.PushBack(flow.NodeInfo{
		Level:  list.Len() + 2,
		NodeID: "结束",
	})
	list.PushFront(flow.NodeInfo{
		Level:  1,
		NodeID: "开始",
		Type:   "starter",
		Mode:   "or",
	})
	arr := util.List2Array(list)
	listStr, err := util.ToJSONStr(arr)
	return listStr, err
}

func (l *StartLogic) saveProcInst(defId int64, dataId int64, title string, remainHours int32, tx *act2.Tx) (*act2.ProcInst, error) {
	inst, err := tx.ProcInst.Create().SetProcDefID(defId).SetDataID(dataId).SetTargetID(general.TargetId).SetStartTime(time.Now()).
		SetTitle(title).SetIsFinished(0).SetRemainHours(remainHours).SetStartUserID(general.UserId3).SetStartUserName(general.UserName3).
		SetState(constant.PENDING).Save(l.ctx)
	return inst, err
}

func (l *StartLogic) saveTask(nodeId string, instId int64, isFinish int32, level int32, step int32, memberApprover string, mode string, tx *act2.Tx) (int64, error) {
	taskCreate := tx.Task.Create().SetNodeID(nodeId).SetProcInstID(instId).
		SetIsFinished(isFinish).SetLevel(level).SetStep(step).SetMemberApprover(memberApprover)
	if mode != "" {
		taskCreate.SetMode(mode)
	}
	t, err := taskCreate.Save(l.ctx)
	return t.ID, err
}

func (l *StartLogic) SaveExecution(instId int64, nodeInfos string, tx *act2.Tx) (*act2.Execution, error) {
	exec, err := tx.Execution.Create().SetProcInstID(instId).SetNodeInfos(nodeInfos).SetStartTime(time.Now()).
		Save(l.ctx)
	return exec, err
}

func (l *StartLogic) UpdateProcInst(nodeId string, dataId int64, taskId int64, state int32, isFinished int32, code string, tx *act2.Tx) error {
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

func (l *StartLogic) saveIdentityLink(instId int64, taskId int64, step int32, userId int64, userName string, tx *act2.Tx) error {
	_, err := tx.IdentityLink.Create().SetProcInstID(instId).SetTaskID(taskId).SetTargetID(general.TargetId).
		SetIsDeal(0).SetStep(step).SetUserID(userId).SetUserName(userName).Save(l.ctx)

	return err
}
