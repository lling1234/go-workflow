package logic

import (
	act2 "act/common/act"
	"act/common/act/task"
	"context"

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

func (l *FindLatestTaskIdLogic) FindLatestTaskId(in *act.DataIdReq) (*act.TaskIdReply, error) {
	tx, err := l.svcCtx.CommonStore.Tx(l.ctx)
	if err != nil {
		return nil, err
	}
	t, err := tx.Task.Query().Where(task.DataID(in.DataId)).Order(act2.Desc(task.FieldCreateTime)).First(l.ctx)
	if err != nil {
		return nil, err
	}
	return &act.TaskIdReply{
		Id: t.ID,
	}, nil
}
