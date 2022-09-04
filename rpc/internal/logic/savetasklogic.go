package logic

import (
	"act/common/act/task"
	"act/rpc/internal/svc"
	"act/rpc/types/act"
	"context"
	"github.com/zeromicro/go-zero/core/logx"
)

type SaveTaskLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewSaveTaskLogic(ctx context.Context, svcCtx *svc.ServiceContext) *SaveTaskLogic {
	return &SaveTaskLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *SaveTaskLogic) SaveTask(in *act.TaskReq) (*act.TaskReply, error) {
	tx, err := l.svcCtx.CommonStore.Tx(l.ctx)
	if err != nil {
		return nil, err
	}
	//a, err = tx.Task.Create().SetDataID(in.DataId).SetNodeID(in.NodeId).SetProcInstID(in.ProcInstId).SetIsFinished(int8(in.IsFinished)).SetLevel(in.Level).
	//	SetMemberApprover(in.MemberApprover).SetStep(in.Step).SetMode(mode).Save(l.ctx)
	//
	//t, err := tx.Task.Create().SetDataID(in.DataId).SetNodeID(in.NodeId).SetProcInstID(in.ProcInstId).SetIsFinished(int8(in.IsFinished)).SetLevel(in.Level).
	//	SetMemberApprover(in.MemberApprover).SetStep(in.Step).SetMode(mode).Save(l.ctx)
	taskCreate := tx.Task.Create().SetDataID(in.DataId).SetNodeID(in.NodeId).SetProcInstID(in.ProcInstId).SetIsFinished(int8(in.IsFinished)).SetLevel(in.Level).
		SetMemberApprover(in.MemberApprover).SetStep(in.Step)
	if in.Mode != "" {
		var mode task.Mode
		if "or" == in.Mode {
			mode = task.ModeOr
		} else if "and" == in.Mode {
			mode = task.ModeAnd
		}
		taskCreate.SetMode(mode)
	}
	t, err := taskCreate.Save(l.ctx)
	if err != nil {
		tx.Rollback()
		return nil, err
	}
	tx.Commit()
	return &act.TaskReply{
		Id:     t.ID,
		NodeId: t.NodeID,
		Level:  t.Level,
		Step:   t.Step,
	}, nil
}
