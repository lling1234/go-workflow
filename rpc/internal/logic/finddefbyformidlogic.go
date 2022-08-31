package logic

import (
	"act/common/act/procdef"
	"context"

	act2 "act/common/act"
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
	procdef, err := tx.ProcDef.Query().Where(procdef.FormIDEQ(formId), procdef.TargetIDEQ(123)).First(l.ctx)
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
		CreateTime:     actpd.CreateTime.Format("2006-01-02 15:04:05"),
		TargetId:       actpd.TargetID,
		FormId:         actpd.FormID,
		FormName:       actpd.FormName,
		Resource:       actpd.Resource,
		RemainHours:    actpd.RemainHours,
		IsActive:       int32(actpd.IsActive),
	}
}
