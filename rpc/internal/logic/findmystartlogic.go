package logic

import (
	"context"

	"act/rpc/internal/svc"
	"act/rpc/types/act"

	"github.com/zeromicro/go-zero/core/logx"
)

type FindMyStartLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewFindMyStartLogic(ctx context.Context, svcCtx *svc.ServiceContext) *FindMyStartLogic {
	return &FindMyStartLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *FindMyStartLogic) FindMyStart(in *act.MyProcInstReq) (*act.ProcInstReply, error) {
	// todo: add your logic here and delete this line

	return &act.ProcInstReply{}, nil
}
