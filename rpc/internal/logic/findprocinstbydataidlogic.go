package logic

import (
	"context"

	"act/rpc/internal/svc"
	"act/rpc/types/act"

	"github.com/zeromicro/go-zero/core/logx"
)

type FindProcInstByDataIdLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewFindProcInstByDataIdLogic(ctx context.Context, svcCtx *svc.ServiceContext) *FindProcInstByDataIdLogic {
	return &FindProcInstByDataIdLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *FindProcInstByDataIdLogic) FindProcInstByDataId(in *act.DataIdReq) (*act.ProcInstReply, error) {
	// todo: add your logic here and delete this line

	return &act.ProcInstReply{}, nil
}
