package logic

import (
	"context"

	"act/rpc/internal/svc"
	"act/rpc/types/act"

	"github.com/zeromicro/go-zero/core/logx"
)

type SaveTaskLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewSaveTaskLogic(ctx context.Context, svcCtx *svc.ServiceContext) *SaveTaskLogic {
	return &SaveTaskLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *SaveTaskLogic) SaveTask(in *act.TaskReq) (*act.TaskReply, error) {
	// todo: add your logic here and delete this line

	return &act.TaskReply{}, nil
}
