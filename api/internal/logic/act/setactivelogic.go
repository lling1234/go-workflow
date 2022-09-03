package act

import (
	"act/rpc/types/act"
	"context"

	"act/api/internal/svc"
	"act/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type SetActiveLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewSetActiveLogic(ctx context.Context, svcCtx *svc.ServiceContext) *SetActiveLogic {
	return &SetActiveLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *SetActiveLogic) SetActive(req *types.FormIdReq) (resp *types.CommonResponse, err error) {
	_, err = l.svcCtx.Rpc.SetProcDefActive(l.ctx, &act.FindProcDefReq{
		FormId:  req.FormId,
		Version: req.Version,
	})
	return types.GetCommonResponse(err, "ok")
}
