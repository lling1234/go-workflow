package logic

import (
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
	//targetIdStr, err := l.svcCtx.CommonStore.Cache.Get(l.ctx, "targetId")
	//if err != nil {
	//	return nil, err
	//}

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
	reply := l.convert(in)
	return reply, nil
}

func (l *SaveProcDefLogic) convert(in *act.ProcDefReq) *act.ProcDefReply {
	return &act.ProcDefReply{
		Name:        in.Name,
		Code:        in.Code,
		YewuFormId:  in.FormId,
		YewuName:    in.FormName,
		RemainHours: in.RemainHours,
		Resource:    in.Resource,
		CreateTime:  time.Now().Format("2006-01-02 15:04:05"),
		Version:     1,
		TargetId:    in.TargetId,
		IsDel:       0,
		IsActive:    1,
	}
}
