package logic

import (
	"context"

	"act/rpc/internal/svc"
	"act/rpc/types/act"

	"github.com/zeromicro/go-zero/core/logx"
)

type SaveIdentityLinkLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewSaveIdentityLinkLogic(ctx context.Context, svcCtx *svc.ServiceContext) *SaveIdentityLinkLogic {
	return &SaveIdentityLinkLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *SaveIdentityLinkLogic) SaveIdentityLink(in *act.IdentityLinkReq) (*act.IdentityLinkReply, error) {
	// todo: add your logic here and delete this line

	return &act.IdentityLinkReply{}, nil
}
