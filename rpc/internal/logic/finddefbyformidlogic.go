package logic

import (
	act2 "act/common/act"
	"act/common/act/procdef"
	"act/common/tools/date"
	"context"

	"act/rpc/internal/svc"
	"act/rpc/types/act"
	"github.com/zeromicro/go-zero/core/logx"
)

type FindDefByFormIdLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewFindDefByFormIdLogic(ctx context.Context, svcCtx *svc.ServiceContext) *FindDefByFormIdLogic {
	return &FindDefByFormIdLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *FindDefByFormIdLogic) FindDefByFormId(in *act.FormIdReq) (*act.ProcDefReply, error) {
	formId := in.GetFormId()
	tx, err := l.svcCtx.CommonStore.Tx(l.ctx)
	if err != nil {
		return nil, err
	}
	procdef, err := tx.ProcDef.Query().Where(procdef.FormIDEQ(formId), procdef.TargetIDEQ(123), procdef.IsActiveEQ(1)).First(l.ctx)
	if err != nil {
		return nil, err
	}
	return convert(procdef), nil
}

func convert(actpd *act2.ProcDef) *act.ProcDefReply {
	return &act.ProcDefReply{
		Id:             actpd.ID,
		Name:           actpd.Name,
		Code:           actpd.Code,
		Version:        int32(actpd.Version),
		CreateUserId:   actpd.CreateUserID,
		CreateUserName: actpd.CreateUserName,
		CreateTime:     date.NowStr(),
		TargetId:       actpd.TargetID,
		FormId:         actpd.FormID,
		FormName:       actpd.FormName,
		Resource:       actpd.Resource,
		RemainHours:    int32(actpd.RemainHours),
		IsActive:       int32(actpd.IsActive),
	}
}
