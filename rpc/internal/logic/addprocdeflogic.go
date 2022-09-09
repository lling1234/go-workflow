package logic

import (
	act2 "act/common/act"
	"act/common/act/procdef"
	"act/common/tools/date"
	"act/common/tools/linq"
	"act/rpc/general"
	"context"
	"time"

	"act/rpc/internal/svc"
	"act/rpc/types/act"

	"github.com/zeromicro/go-zero/core/logx"
)

type AddProcDefLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewAddProcDefLogic(ctx context.Context, svcCtx *svc.ServiceContext) *AddProcDefLogic {
	return &AddProcDefLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *AddProcDefLogic) AddProcDef(in *act.AddProcDefReq) (*act.ProcDefReply, error) {
	version, err := l.GetNewVersion(in.FormId)
	if err != nil {
		return nil, err
	}
	_, err = l.svcCtx.CommonStore.ProcDef.Create().
		SetName(in.Name).SetCode(in.Code).SetFormID(in.FormId).SetFormName(in.FormName).
		SetAppName(in.AppName).SetAppID(in.AppId).
		SetRemainHours(in.RemainHours).SetResource(in.Resource).
		SetCreateUserID(general.MyUserId).SetCreateUserName(general.UserName0).
		SetCreateTime(time.Now()).SetVersion(version).
		SetTargetID(general.TargetId).Save(l.ctx)

	if err != nil {
		return nil, err
	}

	reply := l.convert(in, version)
	return reply, nil
}

func (l *AddProcDefLogic) convert(in *act.AddProcDefReq, version int32) *act.ProcDefReply {
	return &act.ProcDefReply{
		Name:        in.Name,
		Code:        in.Code,
		FormId:      in.FormId,
		FormName:    in.FormName,
		AppId:       in.AppId,
		AppName:     in.AppName,
		RemainHours: in.RemainHours,
		Resource:    in.Resource,
		CreateTime:  date.NowStr(),
		Version:     version,
		IsDel:       0,
		IsActive:    0,
	}
}

func (l *AddProcDefLogic) GetNewVersion(formId int64) (int32, error) {
	defs, err := l.svcCtx.CommonStore.ProcDef.Query().Where(procdef.FormIDEQ(formId)).Select(procdef.FieldVersion).All(l.ctx)
	if err != nil {
		return 1, err
	}
	maxVersion := linq.From(defs).SelectT(func(a *act2.ProcDef) int32 {
		return a.Version
	}).Max()
	return maxVersion.(int32) + 1, nil
}
