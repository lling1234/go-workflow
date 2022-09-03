package logic

import (
	act2 "act/common/act"
	"act/common/act/task"
	"context"
	"log"

	"act/rpc/internal/svc"
	"act/rpc/types/act"

	"github.com/zeromicro/go-zero/core/logx"
)

type FindLatestTaskIdLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewFindLatestTaskIdLogic(ctx context.Context, svcCtx *svc.ServiceContext) *FindLatestTaskIdLogic {
	return &FindLatestTaskIdLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *FindLatestTaskIdLogic) FindLatestTaskId(in *act.DataIdReq) (*act.TaskIdArg, error) {
	tx, err := l.svcCtx.CommonStore.Tx(l.ctx)
	if err != nil {
		return nil, err
	}
	var t *act2.Task
	if in.Step != 0 {
		log.Printf("step1:%d", in.Step)
		t, err = tx.Task.Query().Where(task.DataID(in.DataId), task.IsDelEQ(0), task.StepEQ(in.Step)).First(l.ctx)
	} else {
		log.Printf("step2:%d", in.Step)
		t, err = tx.Task.Query().Where(task.DataID(in.DataId), task.IsDelEQ(0)).Order(act2.Desc(task.FieldStep)).First(l.ctx)
	}
	if err != nil {
		return nil, err
	}
	return &act.TaskIdArg{
		Id: t.ID,
	}, nil
}
