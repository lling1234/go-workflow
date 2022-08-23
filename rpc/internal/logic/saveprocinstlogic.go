package logic

import (
	"context"

	"act/rpc/internal/svc"
	"act/rpc/ms"

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

func (l *SaveProcInstLogic) SaveProcInst(in *ms.ProcInstReq) (*ms.ProcInstReply, error) {
	// todo: add your logic here and delete this line

	return &ms.ProcInstReply{}, nil
}
