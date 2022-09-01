package act

import (
	"act/rpc/types/act"
	"context"

	"act/api/internal/svc"
	"act/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type SetProcessActiveLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewSetProcessActiveLogic(ctx context.Context, svcCtx *svc.ServiceContext) *SetProcessActiveLogic {
	return &SetProcessActiveLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *SetProcessActiveLogic) SetProcessActive(req *types.SetProcessActiveReq) (resp *types.CommonResponse, err error) {
	reply, err := l.svcCtx.Rpc.SetProcDefActive(l.ctx, &act.SetProcessActiveReq{
		FormId:  req.FormId,
		Version: req.Version,
	})
	if err != nil {
		return types.GetErrorCommonResponse(err.Error())
	}
	return types.GetCommonResponse(err, reply)
}
