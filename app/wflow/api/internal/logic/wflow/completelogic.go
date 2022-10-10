package wflow

import (
	"context"

	"go-wflow/app/wflow/api/internal/svc"
	"go-wflow/app/wflow/api/internal/types"
	"go-wflow/kernel"

	"github.com/qkbyte/go-zero/core/logx"
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

func (l *CompleteLogic) Complete(req *types.CompleteProcInstReq) (resp *types.CommonResponse, err error) {
	c,err:=l.svcCtx.Rpc.Complete(l.ctx, &kernel.CompleteProcInstReq{
		DataId:  req.DataId,
		Result:  req.Result,
		Comment: req.Comment,
	})
	if err != nil {
		types.GetErrorCommonResponse(err.Error())
	}
	return types.GetSuccessCommonResponse(c)
}
