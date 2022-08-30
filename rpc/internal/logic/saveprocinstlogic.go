package logic

import (
	"context"

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
	// todo: add your logic here and delete this line

	return &act.ProcInstReply{}, nil
}
