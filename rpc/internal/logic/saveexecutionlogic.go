package logic

import (
	"context"

	"act/rpc/internal/svc"
	"act/rpc/types/act"

	"github.com/zeromicro/go-zero/core/logx"
)

type SaveExecutionLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewSaveExecutionLogic(ctx context.Context, svcCtx *svc.ServiceContext) *SaveExecutionLogic {
	return &SaveExecutionLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *SaveExecutionLogic) SaveExecution(in *act.ExecutionReq) (*act.ExecutionReply, error) {
	// todo: add your logic here and delete this line

	return &act.ExecutionReply{}, nil
}
