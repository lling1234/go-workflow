package act

import (
	"act/rpc/actclient"
	"context"

	"act/api/internal/svc"
	"act/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type StartProcInstLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewStartProcInstLogic(ctx context.Context, svcCtx *svc.ServiceContext) *StartProcInstLogic {
	return &StartProcInstLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *StartProcInstLogic) StartProcInst(req *types.StartProcInst) (resp *types.CommonResponse, err error) {
	def, err := l.svcCtx.Rpc.FindDefByFormId(l.ctx, &actclient.FindProcDefReq{
		FormId: req.FormId,
		AppId:  req.AppId,
	})
	if err != nil {
		return types.GetErrorCommonResponse(err.Error())
	}
	resource := def.Resource
	_, err = l.svcCtx.Rpc.Start(l.ctx, &actclient.StartProcInstReq{
		FormId:      req.FormId,
		AppId:       req.AppId,
		DataId:      req.DataId,
		RemainHours: def.RemainHours,
		Title:       req.Title,
		Resource:    resource,
	})
	return types.GetCommonResponse(err, "ok")
}
