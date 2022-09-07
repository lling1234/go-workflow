package logic

import (
	act2 "act/common/act"
	"act/common/act/procdef"
	"act/common/tools/date"
	"act/rpc/general"
	"context"
	"log"
	"time"

	"act/rpc/internal/svc"
	"act/rpc/types/act"

	linq "github.com/ahmetb/go-linq/v3"

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

func (l *SaveProcDefLogic) SaveProcDef(in *act.SaveProcDefReq) (*act.ProcDefReply, error) {
	tx, err := l.svcCtx.CommonStore.Tx(l.ctx)
	if err != nil {
		return nil, err
	}
	version, err := l.FindMaxVersionByFormId(in.FormId, tx)
	if err != nil {
		return nil, err
	}
	version = version + 1
	_, err = tx.ProcDef.Create().
		SetName(in.Name).SetCode(in.Code).SetFormID(in.FormId).SetFormName(in.FormName).
		SetRemainHours(in.RemainHours).SetResource(in.Resource).
		SetCreateUserID(in.UserId).SetCreateUserName(in.UserName).SetCreateTime(time.Now()).SetVersion(version).SetTargetID(general.TargetId).
		Save(l.ctx)

	if err != nil {
		tx.Rollback()
		return nil, err
	}
	err = tx.Commit()
	if err != nil {
		return nil, err
	}

	reply := l.convert(in, version)
	return reply, nil
}

func (l *SaveProcDefLogic) convert(in *act.SaveProcDefReq, version int32) *act.ProcDefReply {
	return &act.ProcDefReply{
		Name:        in.Name,
		Code:        in.Code,
		FormId:      in.FormId,
		FormName:    in.FormName,
		RemainHours: in.RemainHours,
		Resource:    in.Resource,
		CreateTime:  date.NowStr(),
		Version:     version,
		TargetId:    in.TargetId,
		IsDel:       0,
		IsActive:    1,
	}
}

func (l *SaveProcDefLogic) FindMaxVersionByFormId(formId string, tx *act2.Tx) (int32, error) {
	defs, err := tx.ProcDef.Query().Where(procdef.FormIDEQ(formId)).Select(procdef.FieldVersion).All(l.ctx)
	if err != nil {
		return 0, err
	}

	maxVersion := linq.From(defs).SelectT(func(a *act2.ProcDef) int32 {
		return a.Version
	}).Max()
	log.Println("maxVersion", maxVersion)
	// var maxVersion int32 = 0
	// for _, v := range defs {
	// 	if maxVersion < v.Version {
	// 		maxVersion = v.Version
	// 	}
	// }
	return maxVersion.(int32), nil
}
