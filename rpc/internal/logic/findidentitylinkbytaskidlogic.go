package logic

import (
	"act/common/act/identitylink"
	"act/rpc/general"
	"context"
	"errors"

	"act/rpc/internal/svc"
	"act/rpc/types/act"

	"github.com/zeromicro/go-zero/core/logx"
)

type FindIdentityLinkByTaskIdLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewFindIdentityLinkByTaskIdLogic(ctx context.Context, svcCtx *svc.ServiceContext) *FindIdentityLinkByTaskIdLogic {
	return &FindIdentityLinkByTaskIdLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *FindIdentityLinkByTaskIdLogic) FindIdentityLinkByTaskId(in *act.TaskIdArg) (*act.IdentityLinkReply, error) {
	tx, err := l.svcCtx.CommonStore.Tx(l.ctx)
	if err != nil {
		return nil, err
	}
	il, err := tx.IdentityLink.Query().Where(identitylink.TaskIDEQ(in.Id), identitylink.UserIDEQ(general.MyUserId), identitylink.IsDelEQ(0), identitylink.IsDealEQ(0)).First(l.ctx)
	if err != nil {
		return nil, err
	}
	if il == nil {
		return nil, errors.New("用户没有审批权限")
	}
	return &act.IdentityLinkReply{
		Id:         il.ID,
		UserId:     il.UserID,
		UserName:   il.UserName,
		Step:       il.Step,
		ProcInstId: il.ProcInstID,
		TargetId:   il.TargetID,
		TaskId:     il.TaskID,
	}, nil
}
