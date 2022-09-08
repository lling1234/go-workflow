package logic

import (
	"context"

	"act/rpc/internal/svc"
	"act/rpc/types/act"

	"github.com/zeromicro/go-zero/core/logx"
)

type FindMyFinishStartLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewFindMyFinishStartLogic(ctx context.Context, svcCtx *svc.ServiceContext) *FindMyFinishStartLogic {
	return &FindMyFinishStartLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *FindMyFinishStartLogic) FindMyFinishStart(in *act.MyProcInstReq) (*act.ProcInstReply, error) {
	// todo: add your logic here and delete this line

	return &act.ProcInstReply{}, nil
}
