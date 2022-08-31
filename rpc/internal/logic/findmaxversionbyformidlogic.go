package logic

import (
	"act/common/act/procdef"
	"context"

	"act/rpc/internal/svc"
	"act/rpc/types/act"

	"github.com/zeromicro/go-zero/core/logx"
)

type FindMaxVersionByFormIdLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewFindMaxVersionByFormIdLogic(ctx context.Context, svcCtx *svc.ServiceContext) *FindMaxVersionByFormIdLogic {
	return &FindMaxVersionByFormIdLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *FindMaxVersionByFormIdLogic) FindMaxVersionByFormId(in *act.FormIdReq) (*act.MaxVersionReply, error) {
	tx, err := l.svcCtx.CommonStore.Tx(l.ctx)
	if err != nil {
		return nil, err
	}
	prodef, err := tx.ProcDef.Query().Where(procdef.MaxVersion()).First(l.ctx)
	if err != nil {
		return nil, err
	}
	return &act.MaxVersionReply{
		Version: prodef.Version,
	}, nil
}
