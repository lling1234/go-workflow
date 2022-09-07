package act

import (
	"act/api/flow"
	"act/rpc/actclient"
	"container/list"
	"context"
	"github.com/mumushuiding/util"
	"strconv"
	"strings"

	"act/api/internal/svc"
	"act/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type StartProcInstLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewStartProcInstLogic(ctx context.Context, svcCtx *svc.ServiceContext) *StartProcInstLogic {
	return &StartProcInstLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *StartProcInstLogic) StartProcInst(req *types.StartProcInst) (resp *types.CommonResponse, err error) {
	RPC := l.svcCtx.Rpc
	def, err := l.svcCtx.Rpc.FindDefByFormId(l.ctx, &actclient.FindProcDefReq{
		FormId: req.FormId,
	})
	if err != nil {
		return types.GetErrorCommonResponse(err.Error())
	}
	resource := def.Resource
	instReq := &actclient.ProcInstReq{
		ProcDefId:   def.Id,
		Title:       req.Title,
		FormId:      req.FormId,
		DataId:      req.DataId,
		RemainHours: def.RemainHours,
		State:       flow.PENDING,
	}
	inst, err := RPC.SaveProcInst(l.ctx, instReq)
	if err != nil {
		return types.GetErrorCommonResponse(err.Error())
	}
	var STEP int32 = 1
	task := actclient.TaskReq{
		NodeId:     "开始",
		ProcInstId: inst.Id,
		Level:      STEP,
		IsFinished: 1,
		Step:       STEP,
		//MemberCount:   1,
		Mode: "or",
	}
	_, err = RPC.SaveTask(l.ctx, &task)
	if err != nil {
		return types.GetErrorCommonResponse(err.Error())
	}
	var node *flow.Node
	err = util.Str2Struct(resource, &node)
	if err != nil {
		return types.GetErrorCommonResponse(err.Error())
	}
	list, err := flow.ParseProcessConfig(node)
	listStr, err := GenerateExec(list)
	exec := &actclient.ExecutionReq{
		ProcDefId:  def.Id,
		ProcInstId: inst.Id,
		NodeInfos:  listStr,
	}
	_, err = RPC.SaveExecution(l.ctx, exec)
	if err != nil {
		return types.GetErrorCommonResponse(err.Error())
	}
	// 获取执行流信息
	var nodeInfos []*flow.NodeInfo
	util.Str2Struct(exec.NodeInfos, &nodeInfos)
	// -----------------生成新任务-------------
	firstNode := nodeInfos[1]
	//if firstNode.Mode == "AND" {
	//task.MemberCount = int32(firstNode.MemberCount)
	//}
	task.NodeId = firstNode.NodeID
	STEP = 2
	//task.AgreeNum = 0
	task.IsFinished = 0
	task.Level = STEP
	task.Step = STEP
	task.Mode = firstNode.Mode
	task.MemberApprover = firstNode.ApproverIds
	newTask, err := RPC.SaveTask(l.ctx, &task)
	if err != nil {
		return types.GetErrorCommonResponse(err.Error())
	}
	//newTaskReply, err := RPC.FindLatestTask(l.ctx, &actclient.DataIdReq{
	//	DataId: req.DataId,
	//	Step:   STEP,
	//})
	//if err != nil {
	//	return types.GetErrorCommonResponse(err.Error())
	//}
	updateInst := &actclient.UpdateProcInstReq{
		ProcDefId:   def.Id,
		FormId:      req.FormId,
		DataId:      req.DataId,
		RemainHours: def.RemainHours,
		State:       flow.PENDING,
		TaskId:      newTask.Id,
		NodeId:      firstNode.NodeID,
	}
	_, err = RPC.UpdateProcInst(l.ctx, updateInst)
	if err != nil {
		return types.GetErrorCommonResponse(err.Error())
	}
	userIds := strings.Split(firstNode.ApproverIds, ",")
	userNames := strings.Split(firstNode.ApproverNames, ",")
	for k, v := range userIds {
		userId, _ := strconv.ParseInt(v, 10, 64)
		identityLink := actclient.IdentityLinkReq{
			ProcInstId: inst.Id,
			TaskId:     newTask.Id,
			UserId:     userId,
			UserName:   userNames[k],
			Step:       STEP,
		}
		_, err = RPC.SaveIdentityLink(l.ctx, &identityLink)
	}

	return types.GetCommonResponse(err, inst)
}

func GenerateExec(list *list.List) (string, error) {
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
