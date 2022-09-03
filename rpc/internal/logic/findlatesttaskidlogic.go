package logic

import (
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

func (l *FindLatestTaskIdLogic) FindLatestTaskId(in *act.DataIdReq) (*act.TaskIdReply, error) {
	tx, err := l.svcCtx.CommonStore.Tx(l.ctx)
	if err != nil {
		return nil, err
	}
	log.Println(1111122222)
	t, err := tx.Task.Query().Where(task.DataID(in.DataId), task.IsDelEQ(0), task.StepEQ(in.Step)).First(l.ctx)
	log.Println(222221111)
	if err != nil {
		return nil, err
	}
	return &act.TaskIdReply{
		Id: t.ID,
	}, nil
}
