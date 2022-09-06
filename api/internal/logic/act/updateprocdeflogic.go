package act

import (
	"act/api/internal/svc"
	"act/api/internal/types"
	"act/rpc/types/act"
	"context"
	"encoding/json"
	"github.com/zeromicro/go-zero/core/logx"
)

type UpdateProcDefLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewUpdateProcDefLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UpdateProcDefLogic {
	return &UpdateProcDefLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *UpdateProcDefLogic) UpdateProcDef(req *types.SaveProcDef) (resp *types.CommonResponse, err error) {
	if req.Resource == nil || len(req.Resource.Name) == 0 {
		return types.GetErrorCommonResponse("节点数据未获取到！")
	}
	resource, err := json.Marshal(req.Resource)
	if err != nil {
		return types.GetErrorCommonResponse(err.Error())
	}
	reply, err := l.svcCtx.Rpc.UpdateProcDef(l.ctx, &act.FindProcDefReq{
		Name:        req.Name,
		Code:        req.Code,
		FormId:      req.FormId,
		Resource:    string(resource),
		RemainHours: req.RemainHours,
		Version:     req.Version,
	})

	if err != nil {
		return types.GetErrorCommonResponse(err.Error())
	}
	return types.GetCommonResponse(err, reply)
}
