package logic

import (
	"act/common/act/procdef"
	"act/rpc/internal/svc"
	"act/rpc/types/act"
	"context"
	"time"

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
	formId := in.FormId
	version := in.Version
	tx, err := l.svcCtx.CommonStore.Tx(l.ctx)
	if err != nil {
		return nil, err
	}
	// 如果传入的参数中有is_active=1。我们就认为是要启用该流程定义。
	//否则就是要修改该流程定义。
	if in.IsActive == 1 {
		err = tx.ProcDef.Update().Where(procdef.FormIDEQ(formId), procdef.TargetIDEQ(1727882), procdef.VersionEQ(version)).SetIsActive(1).SetUpdateTime(time.Now()).Exec(l.ctx)
		if err != nil {
			tx.Rollback()
			return nil, err
		}
		err = tx.ProcDef.Update().Where(procdef.FormIDEQ(formId), procdef.TargetIDEQ(1727882), procdef.VersionNEQ(version)).SetIsActive(0).Exec(l.ctx)
		if err != nil {
			tx.Rollback()
			return nil, err
		}
	} else {
		err = tx.ProcDef.Update().Where(procdef.FormIDEQ(formId), procdef.TargetIDEQ(1727882), procdef.VersionEQ(version)).SetResource(in.Resource).
			SetRemainHours(in.RemainHours).SetName(in.Name).SetCode(in.Code).Exec(l.ctx)
		if err != nil {
			tx.Rollback()
			return nil, err
		}
	}

	tx.Commit()
	return &act.ProcDefReply{}, nil
}
