package logic

import (
	"act/common/act/procdef"
	"act/common/act/procinst"
	"act/rpc/general"
	"context"
	"errors"

	"act/rpc/internal/svc"
	"act/rpc/types/act"

	"github.com/zeromicro/go-zero/core/logx"
)

type DelProcDefLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewDelProcDefLogic(ctx context.Context, svcCtx *svc.ServiceContext) *DelProcDefLogic {
	return &DelProcDefLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *DelProcDefLogic) DelProcDef(in *act.FindProcDefReq) (*act.Nil, error) {
	tx, err := l.svcCtx.CommonStore.Tx(l.ctx)
	if err != nil {
		return &act.Nil{}, err
	}
	def, err := tx.ProcDef.Query().Where(procdef.FormIDEQ(in.FormId), procdef.TargetIDEQ(general.TargetId), procdef.VersionEQ(in.Version)).First(l.ctx)
	if err != nil {
		return &act.Nil{}, err
	}
<<<<<<< HEAD
	inst, err := tx.ProcInst.Query().Where(procinst.ProcDefIDEQ(def.ID), procinst.IsDelEQ(0)).First(l.ctx)
	if err != nil {
		return &act.Nil{}, err
	}
	if inst != nil {
		return &act.Nil{}, errors.New("该流程定义已被引用，无法删除。")
	}
=======
	insts, err := tx.ProcInst.Query().Where(procinst.ProcDefIDEQ(def.ID), procinst.IsDelEQ(0)).All(l.ctx)
	if err != nil {
		return &act.Nil{}, err
	}
	if insts != nil && len(insts) > 0 {
		return &act.Nil{}, errors.New("该流程定义已被引用，无法删除。")
	}

>>>>>>> 2b4417e13dae4513883e0b8957c2674704c971fb
	err = tx.ProcDef.Update().Where(procdef.IDEQ(def.ID)).SetIsDel(1).Exec(l.ctx)
	if err != nil {
		tx.Rollback()
		return &act.Nil{}, err
	}
<<<<<<< HEAD
=======
	tx.Commit()
>>>>>>> 2b4417e13dae4513883e0b8957c2674704c971fb
	return &act.Nil{}, nil
}
