package logic

import (
	"act/common/act/procdef"
	"act/rpc/general"
	"context"

	"act/rpc/internal/svc"
	"act/rpc/types/act"

	"github.com/zeromicro/go-zero/core/logx"
)

type DelProcDefLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewDelProcDefLogic(ctx context.Context, svcCtx *svc.ServiceContext) *DelProcDefLogic {
	return &DelProcDefLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *DelProcDefLogic) DelProcDef(in *act.FindProcDefReq) (*act.Nil, error) {
	tx, err := l.svcCtx.CommonStore.Tx(l.ctx)
	if err != nil {
		return nil, err
	}
	err = tx.ProcDef.Update().Where(procdef.FormIDEQ(in.FormId), procdef.TargetIDEQ(general.TargetId), procdef.VersionEQ(in.Version)).SetIsDel(1).Exec(l.ctx)
	if err != nil {
		tx.Rollback()
		return nil, err
	}
	return &act.Nil{}, nil
}
