package logic

import (
	act2 "act/common/act"
	"act/common/act/task"
	"context"

	"act/rpc/internal/svc"
	"act/rpc/types/act"

	"github.com/zeromicro/go-zero/core/logx"
)

type FindLeastTaskIdLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewFindLeastTaskIdLogic(ctx context.Context, svcCtx *svc.ServiceContext) *FindLeastTaskIdLogic {
	return &FindLeastTaskIdLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *FindLeastTaskIdLogic) FindLeastTaskId(in *act.DataIdReq) (*act.TaskIdReply, error) {
	tx, err := l.svcCtx.CommonStore.Tx(l.ctx)
	if err != nil {
		return nil, err
	}
	t := tx.Task.Query().Where(task.DataID(in.DataId)).Order(act2.Desc(task.FieldCreateTime))
	ds, err := t.IDs(l.ctx)
	if err != nil {
		return nil, err
	}
	return &act.TaskIdReply{
		Id: ds[0],
	}, nil
}
