package logic

import (
	"context"
	"time"

	"act/rpc/internal/svc"
	"act/rpc/types/act"

	"github.com/zeromicro/go-zero/core/logx"
)

type SaveProcInstLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewSaveProcInstLogic(ctx context.Context, svcCtx *svc.ServiceContext) *SaveProcInstLogic {
	return &SaveProcInstLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *SaveProcInstLogic) SaveProcInst(in *act.ProcInstReq) (*act.ProcInstReply, error) {
	tx, err := l.svcCtx.CommonStore.Tx(l.ctx)
	if err != nil {
		return nil, err
	}
	inst, err := tx.ProcInst.Create().SetProcDefID(in.ProcDefId).SetDataID(in.DataId).SetTargetID(1727882).SetStartTime(time.Now()).
		SetTitle(in.Title).SetIsFinished(0).SetRemainHours(in.RemainHours).SetStartUserID(11025).SetStartUserName("xiaoming").
		SetState(1).Save(l.ctx)
	if err != nil {
		tx.Rollback()
		return nil, err
	}
	err = tx.Commit()
	if err != nil {
		return nil, err
	}
	return &act.ProcInstReply{
		Id:          inst.ID,
		DataId:      in.DataId,
		Title:       in.Title,
		RemainHours: in.RemainHours,
		State:       in.State,
		TargetId:    inst.TargetID,
	}, nil
}
