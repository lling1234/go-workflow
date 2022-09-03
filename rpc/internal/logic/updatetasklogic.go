package logic

import (
	"act/common/act/task"
	"context"
	"time"

	"act/rpc/internal/svc"
	"act/rpc/types/act"

	"github.com/zeromicro/go-zero/core/logx"
)

type UpdateTaskLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewUpdateTaskLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UpdateTaskLogic {
	return &UpdateTaskLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *UpdateTaskLogic) UpdateTask(in *act.TaskReq) (*act.TaskReply, error) {
	tx, err := l.svcCtx.CommonStore.Tx(l.ctx)
	if err != nil {
		return nil, err
	}
	err = tx.Task.Update().Where(task.IDEQ(in.Id)).SetUpdateTime(time.Now()).SetIsFinished(int8(in.IsFinished)).SetClaimTime(time.Now()).Exec(l.ctx)
	if err != nil {
		tx.Rollback()
		return nil, err
	}
	return &act.TaskReply{}, nil
}
