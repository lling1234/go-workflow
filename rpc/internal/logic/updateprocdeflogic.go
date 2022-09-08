package logic

import (
	"act/common/act/procdef"
	"act/rpc/general"
	"context"
	"time"

	"act/rpc/internal/svc"
	"act/rpc/types/act"

	"github.com/zeromicro/go-zero/core/logx"
)

type UpdateProcDefLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewUpdateProcDefLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UpdateProcDefLogic {
	return &UpdateProcDefLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *UpdateProcDefLogic) UpdateProcDef(in *act.FindProcDefReq) (*act.ProcDefReply, error) {
	procDefUpdate := l.svcCtx.CommonStore.ProcDef.Update().Where(procdef.FormIDEQ(in.FormId), procdef.TargetIDEQ(general.TargetId), procdef.VersionEQ(in.Version))
	if in.Resource != "" {
		procDefUpdate.SetResource(in.Resource)
	}
	if in.Code != "" {
		procDefUpdate.SetCode(in.Code)
	}
	if in.Name != "" {
		procDefUpdate.SetName(in.Name)
	}
	if in.RemainHours != 0 {
		procDefUpdate.SetRemainHours(in.RemainHours)
	}
	err := procDefUpdate.SetUpdateTime(time.Now()).SetUpdateUserID(general.MyUserId).Exec(l.ctx)
	return &act.ProcDefReply{}, err
}
