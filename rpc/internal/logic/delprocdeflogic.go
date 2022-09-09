package logic

import (
	"act/common/act/procdef"
	"act/common/act/procinst"
	"act/rpc/general"
	"context"
	"errors"
	"time"

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
	def, err := l.svcCtx.CommonStore.ProcDef.Query().Where(procdef.FormIDEQ(in.FormId), procdef.TargetIDEQ(general.TargetId), procdef.VersionEQ(in.Version)).First(l.ctx)
	if err != nil {
		return nil, err
	}
	insts, err := l.svcCtx.CommonStore.ProcInst.Query().Where(procinst.ProcDefIDEQ(def.ID), procinst.IsDelEQ(0)).All(l.ctx)
	if err != nil {
		return nil, err
	}
	if insts != nil && len(insts) > 0 {
		return nil, errors.New("该流程定义已被引用，无法删除。")
	}

	err = l.svcCtx.CommonStore.ProcDef.Update().Where(procdef.IDEQ(def.ID)).SetIsDel(1).SetDelUserID(general.MyUserId).SetDelTime(time.Now()).Exec(l.ctx)
	return &act.Nil{}, err
}
