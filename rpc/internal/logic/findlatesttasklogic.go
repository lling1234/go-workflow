package logic

import (
	act2 "act/common/act"
	"act/common/act/task"
<<<<<<< HEAD
	"act/rpc/internal/svc"
	"act/rpc/types/act"
	"context"
=======
	"context"

	"act/rpc/internal/svc"
	"act/rpc/types/act"
>>>>>>> 2b4417e13dae4513883e0b8957c2674704c971fb

	"github.com/zeromicro/go-zero/core/logx"
)

type FindLatestTaskLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewFindLatestTaskLogic(ctx context.Context, svcCtx *svc.ServiceContext) *FindLatestTaskLogic {
	return &FindLatestTaskLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

<<<<<<< HEAD
func (l *FindLatestTaskLogic) FindLatestTask(in *act.DataIdReq) (*act.TaskReply, error) {
=======
func (l *FindLatestTaskLogic) FindLatestTask(in *act.ProcInstIdArg) (*act.TaskReply, error) {
>>>>>>> 2b4417e13dae4513883e0b8957c2674704c971fb
	tx, err := l.svcCtx.CommonStore.Tx(l.ctx)
	if err != nil {
		return nil, err
	}
<<<<<<< HEAD
	var t *act2.Task
	if in.Step != 0 {
		t, err = tx.Task.Query().Where(task.DataID(in.DataId), task.IsDelEQ(0), task.StepEQ(in.Step)).First(l.ctx)
	} else {
		t, err = tx.Task.Query().Where(task.DataID(in.DataId), task.IsDelEQ(0)).Order(act2.Desc(task.FieldStep)).First(l.ctx)
	}
	if err != nil {
		return nil, err
	}
	//var mode string
	//if t.Mode == task.ModeOr{
	//	mode = "or"
	//}else if t.Mode == task.ModeAnd {
	//	mode = "and"
	//}
=======

	t, err := tx.Task.Query().Where(task.ProcInstID(in.GetId()), task.IsDelEQ(0)).Order(act2.Desc(task.FieldStep)).First(l.ctx)

	if err != nil {
		return nil, err
	}

>>>>>>> 2b4417e13dae4513883e0b8957c2674704c971fb
	return &act.TaskReply{
		Id:         t.ID,
		Step:       t.Step,
		Level:      t.Step,
		Mode:       string(t.Mode),
		ProcInstId: t.ProcInstID,
	}, nil
}
