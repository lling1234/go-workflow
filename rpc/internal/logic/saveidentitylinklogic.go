package logic

import (
	"act/rpc/general"
	"act/rpc/internal/svc"
	"act/rpc/types/act"
	"context"

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
	tx, err := l.svcCtx.CommonStore.Tx(l.ctx)
	if err != nil {
		return nil, err
	}
	_, err = tx.IdentityLink.Create().SetProcInstID(in.ProcInstId).SetTaskID(in.TaskId).SetTargetID(general.TargetId).
		SetIsDeal(0).SetStep(in.Step).SetUserID(in.UserId).SetUserName(in.UserName).Save(l.ctx)
	if err != nil {
		tx.Rollback()
		return nil, err
	}
	tx.Commit()
	return &act.IdentityLinkReply{}, nil
}
