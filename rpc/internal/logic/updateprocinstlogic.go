package logic

import (
	"act/common/act/procinst"
	"act/rpc/constant"
	"context"
	"time"

	"act/rpc/internal/svc"
	"act/rpc/types/act"

	"github.com/zeromicro/go-zero/core/logx"
)

type UpdateProcInstLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewUpdateProcInstLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UpdateProcInstLogic {
	return &UpdateProcInstLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *UpdateProcInstLogic) UpdateProcInst(in *act.UpdateProcInstReq) (*act.ProcInstReply, error) {
	tx, err := l.svcCtx.CommonStore.Tx(l.ctx)
	if err != nil {
		return nil, err
	}

	procInstUpdate := tx.ProcInst.Update().Where(procinst.DataIDEQ(in.DataId), procinst.StateNotIn(constant.WITHDRAW, constant.DISCARD), procinst.IsDelEQ(0))

	if in.TaskId != 0 {
		procInstUpdate.SetNodeID(in.NodeId).SetTaskID(in.TaskId)
	}
	if in.State != 0 {
		procInstUpdate.SetState(in.State)
	}
	if in.State == constant.HAVEPASS {
		procInstUpdate.SetCode(in.Code)
	}
	if in.IsFinish == 1 {
		procInstUpdate.SetIsFinished(1).SetEndTime(time.Now())
	}
	err = procInstUpdate.Exec(l.ctx)
	if err != nil {
		tx.Rollback()
		return &act.ProcInstReply{}, err
	}
	tx.Commit()
	return &act.ProcInstReply{}, nil
}
