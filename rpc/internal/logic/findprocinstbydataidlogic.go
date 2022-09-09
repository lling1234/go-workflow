package logic

import (
	"act/common/act/procinst"
	"context"
	"errors"

	"act/rpc/internal/svc"
	"act/rpc/types/act"

	"github.com/zeromicro/go-zero/core/logx"
)

type FindProcInstByDataIdLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewFindProcInstByDataIdLogic(ctx context.Context, svcCtx *svc.ServiceContext) *FindProcInstByDataIdLogic {
	return &FindProcInstByDataIdLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *FindProcInstByDataIdLogic) FindProcInstByDataId(in *act.DataIdReq) (*act.ProcInstReply, error) {
	insts, err := l.svcCtx.CommonStore.ProcInst.Query().Where(procinst.DataIDEQ(in.DataId), procinst.IsDelEQ(0)).All(l.ctx)
	if err != nil {
		return nil, err
	}
	if insts == nil || len(insts) == 0 {
		return nil, errors.New("该数据没有对应的流程实例")
	}
	inst := insts[0]
	return &act.ProcInstReply{
		Id:       inst.ID,
		FlowType: inst.FlowType,
	}, nil
}
