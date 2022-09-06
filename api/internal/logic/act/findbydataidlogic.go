package act

import (
	"act/rpc/types/act"
	"context"

	"act/api/internal/svc"
	"act/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type FindByDataIdLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewFindByDataIdLogic(ctx context.Context, svcCtx *svc.ServiceContext) *FindByDataIdLogic {
	return &FindByDataIdLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *FindByDataIdLogic) FindByDataId(req *types.DataIdReq) (resp *types.CommonResponse, err error) {
	reply, err := l.svcCtx.Rpc.FindProcInstByDataId(l.ctx, &act.DataIdReq{
		DataId: req.DataId,
	})

	return types.GetCommonResponse(err, reply)
}
