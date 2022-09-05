package logic

import (
	"act/common/act/execution"
	"context"

	"act/rpc/internal/svc"
	"act/rpc/types/act"

	"github.com/zeromicro/go-zero/core/logx"
)

type FindExecutionByInstIdLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewFindExecutionByInstIdLogic(ctx context.Context, svcCtx *svc.ServiceContext) *FindExecutionByInstIdLogic {
	return &FindExecutionByInstIdLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *FindExecutionByInstIdLogic) FindExecutionByInstId(in *act.ProcInstIdArg) (*act.ExecutionReq, error) {
	tx, err := l.svcCtx.CommonStore.Tx(l.ctx)
	if err != nil {
		return nil, err
	}
	exec, err := tx.Execution.Query().Where(execution.ProcInstID(in.Id), execution.IsDelEQ(0)).First(l.ctx)
	if err != nil {
		return nil, err
	}
	return &act.ExecutionReq{
		NodeInfos:  exec.NodeInfos,
		ProcInstId: in.Id,
	}, nil
}
