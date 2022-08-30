package logic

import (
	"context"

	"act/rpc/internal/svc"
	"act/rpc/types/act"

	"github.com/zeromicro/go-zero/core/logx"
)

type FindLeastTaskIdLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewFindLeastTaskIdLogic(ctx context.Context, svcCtx *svc.ServiceContext) *FindLeastTaskIdLogic {
	return &FindLeastTaskIdLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *FindLeastTaskIdLogic) FindLeastTaskId(in *act.DataIdReq) (*act.TaskIdReply, error) {
	// todo: add your logic here and delete this line

	return &act.TaskIdReply{}, nil
}
