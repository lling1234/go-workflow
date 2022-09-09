package logic

import (
	act2 "act/common/act"
	"act/common/act/procdef"
	"act/common/tools/date"
	"act/rpc/general"
	"context"
	"errors"

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

func (l *FindDefByFormIdLogic) FindDefByFormId(in *act.FindProcDefReq) (*act.ProcDefReply, error) {
	defs, err := l.svcCtx.CommonStore.ProcDef.Query().Where(procdef.FormIDEQ(in.GetFormId()), procdef.AppIDEQ(in.AppId), procdef.TargetIDEQ(general.TargetId), procdef.IsActiveEQ(1)).All(l.ctx)
	if err != nil {
		return nil, err
	}
	if defs == nil || len(defs) == 0 {
		return nil, errors.New("找不到对应的流程定义")
	}
	return l.convert(defs[0]), nil
}

func (l *FindDefByFormIdLogic) convert(actpd *act2.ProcDef) *act.ProcDefReply {
	return &act.ProcDefReply{
		Id:             actpd.ID,
		Name:           actpd.Name,
		Code:           actpd.Code,
		Version:        actpd.Version,
		CreateUserId:   actpd.CreateUserID,
		CreateUserName: actpd.CreateUserName,
		CreateTime:     date.NowStr(),
		TargetId:       actpd.TargetID,
		FormId:         actpd.FormID,
		FormName:       actpd.FormName,
		AppId:          actpd.AppID,
		AppName:        actpd.AppName,
		Resource:       actpd.Resource,
		RemainHours:    actpd.RemainHours,
		IsActive:       actpd.IsActive,
	}
}
