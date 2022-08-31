package logic

import (
	"act/common/act/procdef"
	"context"

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

func (l *SetProcDefActiveLogic) SetProcDefActive(in *act.ProcDefIdReq) (*act.ProcDefReply, error) {
	id := in.GetId()
	tx, err := l.svcCtx.CommonStore.Tx(l.ctx)
	if err != nil {
		return nil, err
	}
	def, err := tx.ProcDef.Query().Where(procdef.IDEQ(id)).First(l.ctx)
	if err != nil {
		return nil, err
	}
	tx.ProcDef.UpdateOne(def).SetIsActive(1)
	tx.ProcDef.Update().Where(procdef.FormIDEQ(def.FormID), procdef.IDNEQ(id)).SetIsActive(0)
	return &act.ProcDefReply{}, nil
}
