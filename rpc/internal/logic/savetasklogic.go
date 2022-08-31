package logic

import (
	"act/common/act/task"
	"context"
	"time"

	"act/rpc/internal/svc"
	"act/rpc/types/act"

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
	var actMode task.ActMode
	if "or" == in.ActMode {
		actMode = task.ActModeOr
	} else if "and" == in.ActMode {
		actMode = task.ActModeAnd
	}
	tx.Task.Create().SetDataID(in.DataId).SetCreateTime(time.Now()).SetClaimTime(time.Now()).SetLevel(in.Level).
		SetStep(in.Step).SetIsDel(0).SetActMode(actMode).SetMemberCount(in.MemberCount).SetUnCompleteNum(in.UnCompleteNum).
		SetAgreeNum(in.AgreeNum).
		Save(l.ctx)
	return &act.TaskReply{}, nil
}
