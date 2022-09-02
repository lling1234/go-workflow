package act

import (
	"act/api/flow"
	"act/rpc/actclient"
	"container/list"
	"context"
	"github.com/mumushuiding/util"
	"log"
	"strconv"

	"act/api/internal/svc"
	"act/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type StartProcinstLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewStartProcinstLogic(ctx context.Context, svcCtx *svc.ServiceContext) *StartProcinstLogic {
	return &StartProcinstLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *StartProcinstLogic) StartProcinst(req *types.StartProcinst) (resp *types.CommonResponse, err error) {
	def, err := l.svcCtx.Rpc.FindDefByFormId(l.ctx, &actclient.FindProcdefReq{
		FormId: req.FormId,
	})
	if err != nil {
		return types.GetErrorCommonResponse(err.Error())
	}
	resource := def.Resource
	inst, err := l.svcCtx.Rpc.SaveProcInst(l.ctx, &actclient.ProcInstReq{
		ProcDefId:   def.Id,
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
		NodeId:     "开始",
		ProcInstId: inst.Id,
		DataId:     req.DataId,
		Level:      1,
		IsFinished: 1,
		Step:       1,
		//MemberCount:   1,
		Mode: "or",
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
	log.Printf("listStr:%s", listStr)
	exec := &actclient.ExecutionReq{
		ProcDefId:  def.Id,
		ProcInstId: inst.Id,
		NodeInfos:  listStr,
	}
	_, err = l.svcCtx.Rpc.SaveExecution(l.ctx, exec)
	if err != nil {
		return types.GetErrorCommonResponse(err.Error())
	}
	// 获取执行流信息
	var nodeInfos []*flow.NodeInfo
	util.Str2Struct(exec.NodeInfos, &nodeInfos)
	log.Printf("11111,nodeInfos.id:%s,nodeInfos.name:%s,nodeInfos.mode:%s", nodeInfos[1].ApproverIds, nodeInfos[1].ApproverNames, nodeInfos[1].Mode)
	// -----------------生成新任务-------------
	firstNode := nodeInfos[1]
	if firstNode.Mode == "AND" {
		//task.MemberCount = int32(firstNode.MemberCount)
	}
	task.NodeId = firstNode.NodeID
	//task.AgreeNum = 0
	task.IsFinished = 2
	task.Level = 1
	task.Step = 1
	task.Mode = firstNode.Mode
	log.Println("task", task.Mode)
	newTask, err := l.svcCtx.Rpc.SaveTask(l.ctx, &task)
	if err != nil {
		return types.GetErrorCommonResponse(err.Error())
	}
	log.Println(22222, newTask.Id)
	userId, _ := strconv.ParseInt(firstNode.ApproverIds, 10, 64)
	log.Println(33333, userId)
	identityLink := actclient.IdentityLinkReq{
		ProcInstId: inst.Id,
		TaskId:     newTask.Id,
		UserId:     userId,
		UserName:   firstNode.ApproverNames,
		Comment:    "",
		Result:     1,
	}
	log.Println(333333)
	_, err = l.svcCtx.Rpc.SaveIdentityLink(l.ctx, &identityLink)
	log.Println(44444)
	if err != nil {
		return types.GetErrorCommonResponse(err.Error())
	}
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
