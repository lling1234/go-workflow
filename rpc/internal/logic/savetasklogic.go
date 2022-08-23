package logic

import (
	"context"

	"act/rpc/internal/svc"
	"act/rpc/ms"

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

func (l *SaveTaskLogic) SaveTask(in *ms.TaskReq) (*ms.TaskReply, error) {
	// todo: add your logic here and delete this line

	return &ms.TaskReply{}, nil
}
