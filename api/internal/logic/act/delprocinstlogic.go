package act

import (
	"act/rpc/types/act"
	"context"

	"act/api/internal/svc"
	"act/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type DelProcInstLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewDelProcInstLogic(ctx context.Context, svcCtx *svc.ServiceContext) *DelProcInstLogic {
	return &DelProcInstLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

//TODO 流程删除应该删除 流程实例、执行流程、节点任务、审批人
func (l *DelProcInstLogic) DelProcInst(req *types.DataIdReq) (resp *types.CommonResponse, err error) {
	_, err = l.svcCtx.Rpc.DelProcInst(l.ctx, &act.DataIdReq{
		DataId: req.DataId,
	})

	return types.GetCommonResponse(err, "ok")
}
