package logic

import (
	"act/common/act/procdef"
	"act/rpc/general"
	"act/rpc/internal/svc"
	"act/rpc/types/act"
	"context"
	"github.com/zeromicro/go-zero/core/logx"
<<<<<<< HEAD
=======
	"time"
>>>>>>> 2b4417e13dae4513883e0b8957c2674704c971fb
)

type UpdateProcDefLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewUpdateProcDefLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UpdateProcDefLogic {
	return &UpdateProcDefLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *UpdateProcDefLogic) UpdateProcDef(in *act.FindProcDefReq) (*act.ProcDefReply, error) {
	formId := in.FormId
	version := in.Version
	tx, err := l.svcCtx.CommonStore.Tx(l.ctx)
	if err != nil {
		return nil, err
	}

<<<<<<< HEAD
	err = tx.ProcDef.Update().Where(procdef.FormIDEQ(formId), procdef.TargetIDEQ(general.TargetId), procdef.VersionEQ(version)).SetResource(in.Resource).
		SetRemainHours(in.RemainHours).SetName(in.Name).SetCode(in.Code).Exec(l.ctx)
=======
	procDefUpdate := tx.ProcDef.Update().Where(procdef.FormIDEQ(formId), procdef.TargetIDEQ(general.TargetId), procdef.VersionEQ(version))
	if in.Resource != "" {
		procDefUpdate.SetResource(in.Resource)
	}
	if in.Code != "" {
		procDefUpdate.SetCode(in.Code)
	}
	if in.Name != "" {
		procDefUpdate.SetName(in.Name)
	}
	if in.RemainHours != 0 {
		procDefUpdate.SetRemainHours(in.RemainHours)
	}
	err = procDefUpdate.SetUpdateTime(time.Now()).Exec(l.ctx)
>>>>>>> 2b4417e13dae4513883e0b8957c2674704c971fb
	if err != nil {
		tx.Rollback()
		return nil, err
	}

	tx.Commit()
	return &act.ProcDefReply{}, nil
}
