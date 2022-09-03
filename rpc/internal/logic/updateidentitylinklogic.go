package logic

import (
	"context"

	"act/rpc/internal/svc"
	"act/rpc/types/act"

	"github.com/zeromicro/go-zero/core/logx"
)

type UpdateIdentityLinkLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewUpdateIdentityLinkLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UpdateIdentityLinkLogic {
	return &UpdateIdentityLinkLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *UpdateIdentityLinkLogic) UpdateIdentityLink(in *act.IdentityLinkReq) (*act.IdentityLinkReply, error) {
	// todo: add your logic here and delete this line

	return &act.IdentityLinkReply{}, nil
}
