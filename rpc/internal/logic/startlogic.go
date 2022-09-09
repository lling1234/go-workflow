package logic

import (
	act2 "act/common/act"
	"act/common/flow"
	"act/rpc/internal/svc"
	"act/rpc/types/act"
	"container/list"
	"context"
	"github.com/mumushuiding/util"
	"github.com/zeromicro/go-zero/core/logx"
	"strconv"
	"strings"
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
	if err != nil {
		tx.Rollback()
		return nil, err
	}
	var STEP int32 = 1
	_, err = SaveTask(SaveTaskReq{
		NodeID:     "开始",
		InstId:     inst.ID,
		IsFinished: 1,
		Level:      STEP,
		Step:       STEP,
	}, tx, l.ctx)
	if err != nil {
		tx.Rollback()
		return nil, err
	}
	var node *flow.Node
	err = util.Str2Struct(in.Resource, &node)
	if err != nil {
		return nil, err
	}
	//根据流程定义json判断是否为并行流程
	if flow.CheckConCurrentFlow(node) {

	} else {
		err = l.GenNormalProcess(node, inst.ID, in.DataId, STEP+1, tx)
		if err != nil {
			tx.Rollback()
			return nil, err
		}
	}

	tx.Commit()
	return &act.Nil{}, nil
}

func (l *StartLogic) GenExec(list *list.List) (string, error) {
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

func (l *StartLogic) GenNormalProcess(node *flow.Node, instId int64, dataId int64, step int32, tx *act2.Tx) error {
	list, err := flow.ParseProcessConfig(node)
	listStr, err := l.GenExec(list)
	if err != nil {
		return err
	}
	exec, err := SaveExecution(SaveExecutionReq{
		InstId:    instId,
		NodeInfos: listStr,
	}, tx, l.ctx)
	if err != nil {
		tx.Rollback()
		return err
	}
	// 获取执行流信息
	var nodeInfos []*flow.NodeInfo
	util.Str2Struct(exec.NodeInfos, &nodeInfos)
	// -----------------生成新任务-------------
	firstNode := nodeInfos[1]
	//STEP = 2
	secondTaskId, err := SaveTask(SaveTaskReq{
		InstId:         instId,
		NodeID:         firstNode.NodeID,
		IsFinished:     0,
		Level:          step,
		Step:           step,
		MemberApprover: firstNode.ApproverIds,
		Mode:           firstNode.Mode,
	}, tx, l.ctx)
	err = UpdateProcInst(UpdateProcInstReq{
		NodeId: firstNode.NodeID,
		DataId: dataId,
		TaskId: secondTaskId,
		State:  flow.PENDING,
	}, tx, l.ctx)
	if err != nil {
		tx.Rollback()
		return err
	}
	userIds := strings.Split(firstNode.ApproverIds, ",")
	userNames := strings.Split(firstNode.ApproverNames, ",")
	for k, v := range userIds {
		userId, _ := strconv.ParseInt(v, 10, 64)
		err = SaveIdentityLink(SaveIdentityLinkReq{
			InstId:   instId,
			TaskId:   secondTaskId,
			Step:     step,
			UserId:   userId,
			UserName: userNames[k],
		}, tx, l.ctx)
		if err != nil {
			tx.Rollback()
			return err
		}
	}
	return nil
}
