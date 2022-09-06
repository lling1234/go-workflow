package logic

import (
	"act/common/act/identitylink"
	"context"

	"act/rpc/internal/svc"
	"act/rpc/types/act"

	"github.com/zeromicro/go-zero/core/logx"
)

type DelIdentityLinkLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewDelIdentityLinkLogic(ctx context.Context, svcCtx *svc.ServiceContext) *DelIdentityLinkLogic {
	return &DelIdentityLinkLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

//DelIdentityLink 删除或签中，没有审批的审批人数据
func (l *DelIdentityLinkLogic) DelIdentityLink(in *act.ProcInstIdArg) (*act.Nil, error) {
	tx, err := l.svcCtx.CommonStore.Tx(l.ctx)
	if err != nil {
		return nil, err
	}
	_, err = tx.IdentityLink.Delete().Where(identitylink.ProcInstIDEQ(in.Id), identitylink.IsDealEQ(0)).Exec(l.ctx)
	if err != nil {
		tx.Rollback()
		return nil, err
	}
	tx.Commit()
	return &act.Nil{}, nil
}
