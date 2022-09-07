package act

import (
	"act/rpc/types/act"
	"context"

	"act/api/internal/svc"
	"act/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type DelProcDefLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewDelProcDefLogic(ctx context.Context, svcCtx *svc.ServiceContext) *DelProcDefLogic {
	return &DelProcDefLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *DelProcDefLogic) DelProcDef(req *types.FormIdReq) (resp *types.CommonResponse, err error) {
	_, err = l.svcCtx.Rpc.DelProcDef(l.ctx, &act.FindProcDefReq{
		FormId:  req.FormId,
		Version: req.Version,
	})

	return types.GetCommonResponse(err, "ok")
}
