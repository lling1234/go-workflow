package logic

import (
	"context"

	"act/rpc/internal/svc"
	"act/rpc/types/act"

	"github.com/zeromicro/go-zero/core/logx"
)

type FindElapsedTimeLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewFindElapsedTimeLogic(ctx context.Context, svcCtx *svc.ServiceContext) *FindElapsedTimeLogic {
	return &FindElapsedTimeLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *FindElapsedTimeLogic) FindElapsedTime(in *act.UserReq) (*act.ElapsedTimeReply, error) {
	// todo: add your logic here and delete this line

	return &act.ElapsedTimeReply{}, nil
}
