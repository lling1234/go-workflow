package logic

import (
	"act/common/act/procdef"
	"act/rpc/general"
	"context"
	"time"

	"act/rpc/internal/svc"
	"act/rpc/types/act"

	"github.com/zeromicro/go-zero/core/logx"
)

type SetProcDefActiveLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewSetProcDefActiveLogic(ctx context.Context, svcCtx *svc.ServiceContext) *SetProcDefActiveLogic {
	return &SetProcDefActiveLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *SetProcDefActiveLogic) SetProcDefActive(in *act.FindProcDefReq) (*act.Nil, error) {
	formId := in.FormId
	version := in.Version
	tx, err := l.svcCtx.CommonStore.Tx(l.ctx)
	if err != nil {
		return nil, err
	}
	err = tx.ProcDef.Update().Where(procdef.FormIDEQ(formId), procdef.AppIDEQ(in.AppId), procdef.TargetIDEQ(general.TargetId), procdef.VersionEQ(version)).SetIsActive(1).SetUpdateTime(time.Now()).Exec(l.ctx)
	if err != nil {
		tx.Rollback()
		return nil, err
	}
	err = tx.ProcDef.Update().Where(procdef.FormIDEQ(formId), procdef.AppIDEQ(in.AppId), procdef.TargetIDEQ(general.TargetId), procdef.VersionNEQ(version)).SetIsActive(0).Exec(l.ctx)
	if err != nil {
		tx.Rollback()
		return nil, err
	}

	tx.Commit()
	return &act.Nil{}, nil
}
