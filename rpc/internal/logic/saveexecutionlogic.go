package logic

import (
	"context"
	"time"

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
	tx, err := l.svcCtx.CommonStore.Tx(l.ctx)
	if err != nil {
		return nil, err
	}
	_, err = tx.Execution.Create().
		SetProcDefID(in.ProcDefId).SetProcInstID(in.ProcInstId).SetNodeInfos(in.NodeInfos).SetStartTime(time.Now()).Save(l.ctx)
	if err != nil {
		tx.Rollback()
		return nil, err
	}
	tx.Commit()
	return &act.ExecutionReply{}, nil
}
