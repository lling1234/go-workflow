package logic

import (
	"act/common/flow"
	"act/rpc/internal/svc"
	"act/rpc/types/act"
	"container/list"
	"context"
	"github.com/mumushuiding/util"
	"strconv"
	"strings"

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
	inst, err := SaveProcInst(SaveProcInstReq{
		DefId:       in.DefId,
		DataId:      in.DataId,
		Title:       in.Title,
		RemainHours: in.RemainHours,
	}, tx, l.ctx)
	var STEP int32 = 1
	_, err = SaveTask(SaveTaskReq{
		NodeID:     "开始",
		InstId:     inst.ID,
		IsFinished: 1,
		Level:      STEP,
		Step:       STEP,
	}, tx, l.ctx)
	var node *flow.Node
	err = util.Str2Struct(in.Resource, &node)
	if err != nil {
		return nil, err
	}
	list, err := flow.ParseProcessConfig(node)
	listStr, err := l.GenerateExec(list)
	exec, err := SaveExecution(SaveExecutionReq{
		InstId:    inst.ID,
		NodeInfos: listStr,
	}, tx, l.ctx)
	// 获取执行流信息
	var nodeInfos []*flow.NodeInfo
	util.Str2Struct(exec.NodeInfos, &nodeInfos)
	// -----------------生成新任务-------------
	firstNode := nodeInfos[1]
	STEP = 2
	secondTaskId, err := SaveTask(SaveTaskReq{
		InstId:         inst.ID,
		NodeID:         firstNode.NodeID,
		IsFinished:     0,
		Level:          STEP,
		Step:           STEP,
		MemberApprover: firstNode.ApproverIds,
		Mode:           firstNode.Mode,
	}, tx, l.ctx)
	UpdateProcInst(UpdateProcInstReq{
		NodeId: firstNode.NodeID,
		DataId: in.DataId,
		TaskId: secondTaskId,
		State:  flow.PENDING,
	}, tx, l.ctx)
	userIds := strings.Split(firstNode.ApproverIds, ",")
	userNames := strings.Split(firstNode.ApproverNames, ",")
	for k, v := range userIds {
		userId, _ := strconv.ParseInt(v, 10, 64)
		err = SaveIdentityLink(SaveIdentityLinkReq{
			InstId:   inst.ID,
			TaskId:   secondTaskId,
			Step:     STEP,
			UserId:   userId,
			UserName: userNames[k],
		}, tx, l.ctx)
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
