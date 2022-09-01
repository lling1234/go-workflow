package logic

import (
	"act/common/act/procdef"
	"act/common/tools/date"
	"context"
	"time"

	"act/rpc/internal/svc"
	"act/rpc/types/act"

	"github.com/zeromicro/go-zero/core/logx"
)

type SaveProcDefLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewSaveProcDefLogic(ctx context.Context, svcCtx *svc.ServiceContext) *SaveProcDefLogic {
	return &SaveProcDefLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *SaveProcDefLogic) SaveProcDef(in *act.ProcDefReq) (*act.ProcDefReply, error) {
	tx, err := l.svcCtx.CommonStore.Tx(l.ctx)
	if err != nil {
		return nil, err
	}

	_, err = tx.ProcDef.Create().
		SetName(in.Name).SetCode(in.Code).SetFormID(in.FormId).SetFormName(in.FormName).
		SetRemainHours(int(in.RemainHours)).SetResource(in.Resource).
		SetCreateUserID(in.UserId).SetCreateUserName(in.UserName).SetCreateTime(time.Now()).SetVersion(1).SetTargetID(1727882).
		Save(l.ctx)

	if err != nil {
		tx.Rollback()
		return nil, err
	}
	err = tx.Commit()
	if err != nil {
		return nil, err
	}
	version, err := l.FindMaxVersionByFormId(in.FormId)
	if err != nil {
		return nil, err
	}
	reply := l.convert(in, version)
	return reply, nil
}

func (l *SaveProcDefLogic) convert(in *act.ProcDefReq, version int) *act.ProcDefReply {
	return &act.ProcDefReply{
		Name:        in.Name,
		Code:        in.Code,
		FormId:      in.FormId,
		FormName:    in.FormName,
		RemainHours: in.RemainHours,
		Resource:    in.Resource,
		CreateTime:  date.NowStr(),
		Version:     int32(version),
		TargetId:    in.TargetId,
		IsDel:       0,
		IsActive:    1,
	}
}

func (l *SaveProcDefLogic) FindMaxVersionByFormId(formId string) (int, error) {
	tx, err := l.svcCtx.CommonStore.Tx(l.ctx)
	if err != nil {
		return 0, err
	}
	prodef, err := tx.ProcDef.Query().Where(procdef.MaxVersion(formId), procdef.FormIDEQ(formId)).First(l.ctx)
	if err != nil {
		return 0, err
	}
	if prodef == nil {
		return 1, nil
	}
	return prodef.Version, nil
}
