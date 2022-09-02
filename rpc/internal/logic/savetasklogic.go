package logic

import (
	"act/common/act/task"
	"act/rpc/internal/svc"
	"act/rpc/types/act"
	"context"
	"time"

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
	var mode task.Mode
	if "or" == in.Mode {
		mode = task.ModeOr
	} else if "and" == in.Mode {
		mode = task.ModeAnd
	}
	_, err = tx.Task.Create().SetDataID(in.DataId).SetProcInstID(in.ProcInstId).SetCreateTime(time.Now()).SetClaimTime(time.Now()).SetLevel(in.Level).
		SetStep(in.Step).SetMode(mode).Save(l.ctx)
	if err != nil {
		tx.Rollback()
		return nil, err
	}
	tx.Commit()
	return &act.TaskReply{}, nil
}
