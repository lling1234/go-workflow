package logic

import (
	"context"

	"act/rpc/internal/svc"
	"act/rpc/types/act"

	"github.com/zeromicro/go-zero/core/logx"
)

type FindMyFinishApprovalLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewFindMyFinishApprovalLogic(ctx context.Context, svcCtx *svc.ServiceContext) *FindMyFinishApprovalLogic {
	return &FindMyFinishApprovalLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *FindMyFinishApprovalLogic) FindMyFinishApproval(in *act.MyProcInstReq) (*act.ProcInstReply, error) {
	// todo: add your logic here and delete this line

	return &act.ProcInstReply{}, nil
}
