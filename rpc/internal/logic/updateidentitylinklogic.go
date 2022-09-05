package logic

import (
	"act/common/act/identitylink"
	"act/rpc/general"
	"context"
	"time"

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
	tx, err := l.svcCtx.CommonStore.Tx(l.ctx)
	if err != nil {
		return nil, err
	}

	err = tx.IdentityLink.Update().Where(identitylink.TaskIDEQ(in.TaskId), identitylink.UserIDEQ(general.MyUserId), identitylink.IsDelEQ(0), identitylink.IsDealEQ(0)).
		SetUpdateTime(time.Now()).SetComment(in.Comment).SetResult(in.Result).SetIsDeal(1).Exec(l.ctx)
	if err != nil {
		tx.Rollback()
		return nil, err
	}
	tx.Commit()
	return &act.IdentityLinkReply{}, nil
}
