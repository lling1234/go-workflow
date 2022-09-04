package logic

import (
	"act/common/act/execution"
	"act/common/act/identitylink"
	"act/common/act/procinst"
	"act/common/act/task"
	"context"

	"act/rpc/internal/svc"
	"act/rpc/types/act"

	"github.com/zeromicro/go-zero/core/logx"
)

type DelProcInstLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewDelProcInstLogic(ctx context.Context, svcCtx *svc.ServiceContext) *DelProcInstLogic {
	return &DelProcInstLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *DelProcInstLogic) DelProcInst(in *act.DataIdReq) (*act.Nil, error) {
	tx, err := l.svcCtx.CommonStore.Tx(l.ctx)
	if err != nil {
		return nil, err
	}
	inst, err := tx.ProcInst.Query().Where(procinst.DataIDEQ(in.DataId), procinst.IsDelEQ(0)).First(l.ctx)
	if err != nil {
		return nil, err
	}
	instId := inst.ID
	err = tx.ProcInst.Update().Where(procinst.IDEQ(instId)).SetIsDel(1).Exec(l.ctx)
	if err != nil {
		tx.Rollback()
		return nil, err
	}
	err = tx.Execution.Update().Where(execution.IDEQ(instId)).SetIsDel(1).Exec(l.ctx)
	if err != nil {
		tx.Rollback()
		return nil, err
	}
	err = tx.Task.Update().Where(task.IDEQ(instId)).SetIsDel(1).Exec(l.ctx)
	if err != nil {
		tx.Rollback()
		return nil, err
	}
	err = tx.IdentityLink.Update().Where(identitylink.IDEQ(instId)).SetIsDel(1).Exec(l.ctx)
	if err != nil {
		tx.Rollback()
		return nil, err
	}
	tx.Commit()
	return &act.Nil{}, nil
}
