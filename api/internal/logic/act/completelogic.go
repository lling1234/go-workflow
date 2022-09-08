package act

import (
	"act/rpc/actclient"
	"act/rpc/types/act"
	"context"

	"act/api/internal/svc"
	"act/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type CompleteLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewCompleteLogic(ctx context.Context, svcCtx *svc.ServiceContext) *CompleteLogic {
	return &CompleteLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *CompleteLogic) Complete(req *types.CompleteProcinst) (resp *types.CommonResponse, err error) {
	RPC := l.svcCtx.Rpc
	inst, err := RPC.FindProcInstByDataId(l.ctx, &actclient.DataIdReq{
		DataId: req.DataId,
	})
	if err != nil {
		return types.GetErrorCommonResponse(err.Error())
	}
	// 1、普通流程  2、并行流程
	if inst.FlowType == 1 {
		_, err = l.svcCtx.Rpc.CompleteNormal(l.ctx, &act.CompleteNormalProcInstReq{
			InstId:  inst.Id,
			DataId:  req.DataId,
			Result:  req.Result,
			Comment: req.Comment,
		})
		if err != nil {
			return types.GetErrorCommonResponse(err.Error())
		}
	} else if inst.FlowType == 2 {

	}
	return types.GetCommonResponse(err, "ok")
}
