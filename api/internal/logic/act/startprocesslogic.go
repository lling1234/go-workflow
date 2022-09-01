package act

import (
	"act/api/flow"
	"act/rpc/actclient"
	"container/list"
	"context"
	"github.com/mumushuiding/util"
	"strconv"

	"act/api/internal/svc"
	"act/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type StartProcessLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewStartProcessLogic(ctx context.Context, svcCtx *svc.ServiceContext) *StartProcessLogic {
	return &StartProcessLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *StartProcessLogic) StartProcess(req *types.StartProcess) (resp *types.CommonResponse, err error) {
	def, err := l.svcCtx.Rpc.FindDefByFormId(l.ctx, &actclient.FormIdReq{
		FormId: req.FormId,
	})
	if err != nil {
		return types.GetErrorCommonResponse(err.Error())
	}
	resource := def.Resource

	inst, err := l.svcCtx.Rpc.SaveProcInst(l.ctx, &actclient.ProcInstReq{
		Title:       req.Title,
		FormId:      req.FormId,
		DataId:      req.DataId,
		RemainHours: def.RemainHours,
		State:       flow.PENDING,
	})
	if err != nil {
		return types.GetErrorCommonResponse(err.Error())
	}
	task := actclient.TaskReq{
		NodeId:        "开始",
		ProcInstId:    inst.Id,
		DataId:        req.DataId,
		Level:         1,
		IsFinished:    1,
		Step:          1,
		MemberCount:   1,
		UnCompleteNum: 0,
		ActMode:       "or",
		AgreeNum:      1,
	}
	_, err = l.svcCtx.Rpc.SaveTask(l.ctx, &task)
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
	l.svcCtx.Rpc.SaveExecution(l.ctx, exec)
	// 获取执行流信息
	var nodeInfos []*flow.NodeInfo
	util.Str2Struct(exec.NodeInfos, &nodeInfos)

	// -----------------生成新任务-------------
	firstNode := nodeInfos[1]
	if firstNode.ActMode == "and" {
		task.UnCompleteNum = int32(firstNode.MemberCount)
		task.MemberCount = int32(firstNode.MemberCount)
	}
	task.NodeId = firstNode.NodeID
	task.AgreeNum = 0
	task.IsFinished = 2
	task.Level = 1
	task.Step = 1
	task.ActMode = firstNode.ActMode
	newTask, err := l.svcCtx.Rpc.SaveTask(l.ctx, &task)
	userId, _ := strconv.ParseInt(firstNode.ApproverIds, 10, 64)
	identityLink := actclient.IdentityLinkReq{
		ProcInstId: inst.Id,
		TaskId:     newTask.Id,
		UserId:     userId,
		UserName:   firstNode.ApproverNames,
	}
	l.svcCtx.Rpc.SaveIdentityLink(l.ctx, &identityLink)
	return types.GetCommonResponse(nil, inst)
}

func GenerateExec(list *list.List) (string, error) {
	list.PushBack(flow.NodeInfo{
		Level:  list.Len() + 1,
		NodeID: "结束",
	})
	list.PushFront(flow.NodeInfo{
		Level:  1,
		NodeID: "开始",
		Type:   "starter",
	})
	arr := util.List2Array(list)
	listStr, err := util.ToJSONStr(arr)
	return listStr, err
}
