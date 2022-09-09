package act

import (
	"act/api/internal/svc"
	"act/api/internal/types"
	"act/common/flow"
	"act/rpc/types/act"
	"context"
	"encoding/json"
	"github.com/zeromicro/go-zero/core/logx"
)

type AddProcDefLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewAddProcDefLogic(ctx context.Context, svcCtx *svc.ServiceContext) *AddProcDefLogic {
	return &AddProcDefLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *AddProcDefLogic) AddProcDef(req *types.SaveProcDef) (resp *types.CommonResponse, err error) {
	if req.FormId == 0 {
		return types.GetErrorCommonResponse("业务表单数据未获取到！")
	}
	if req.Resource == nil || len(req.Resource.Name) == 0 {
		return types.GetErrorCommonResponse("节点数据未获取到！")
	}
	if flow.IfProcessConfigIsValid(req.Resource) != nil {
		return types.GetErrorCommonResponse(err.Error())
	}
	resource, err := json.Marshal(req.Resource)
	if err != nil {
		return types.GetErrorCommonResponse(err.Error())
	}
	reply, err := l.svcCtx.Rpc.AddProcDef(l.ctx, &act.AddProcDefReq{
		Name:        req.Name,
		Code:        req.Code,
		FormId:      req.FormId,
		FormName:    req.FormName,
		AppId:       req.AppId,
		AppName:     req.AppName,
		Resource:    string(resource),
		RemainHours: req.RemainHours,
	})
	return types.GetCommonResponse(err, reply)
}
