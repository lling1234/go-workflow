package logic

import (
	act2 "act/common/act"
	"act/common/act/task"
	"context"

	"act/rpc/internal/svc"
	"act/rpc/types/act"

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

func (l *FindLatestTaskLogic) FindLatestTask(in *act.ProcInstIdArg) (*act.TaskReply, error) {
	tx, err := l.svcCtx.CommonStore.Tx(l.ctx)
	if err != nil {
		return nil, err
	}

	t, err := tx.Task.Query().Where(task.ProcInstID(in.GetId()), task.IsDelEQ(0)).Order(act2.Desc(task.FieldStep)).First(l.ctx)

	if err != nil {
		return nil, err
	}

	return &act.TaskReply{
		Id:         t.ID,
		Step:       t.Step,
		Level:      t.Step,
		Mode:       string(t.Mode),
		ProcInstId: t.ProcInstID,
	}, nil
}
