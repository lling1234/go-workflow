package logic

import (
	"act/common/act/procinst"
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
	err = tx.ProcInst.Update().Where(procinst.DataIDEQ(in.DataId)).SetIsDel(1).Exec(l.ctx)
	if err != nil {
		tx.Rollback()
		return nil, err
	}
	return &act.Nil{}, nil
}
