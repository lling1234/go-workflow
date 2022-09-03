package logic

import (
	"act/common/act/procinst"
	"context"

	"act/rpc/internal/svc"
	"act/rpc/types/act"

	"github.com/zeromicro/go-zero/core/logx"
)

type UpdateProcInstLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewUpdateProcInstLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UpdateProcInstLogic {
	return &UpdateProcInstLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *UpdateProcInstLogic) UpdateProcInst(in *act.ProcInstReq) (*act.ProcInstReply, error) {
	tx, err := l.svcCtx.CommonStore.Tx(l.ctx)
	if err != nil {
		return nil, err
	}
	tx.ProcInst.Update().Where(procinst.ProcDefIDEQ(in.ProcDefId), procinst.StateNotIn(4, 7), procinst.IsDelEQ(0)).
		SetNodeID(in.NodeId).SetIsFinished(1).SetTaskID(in.TaskId)

	return &act.ProcInstReply{}, nil
}
