package act

import (
	"act/rpc/types/act"
	"context"
	"encoding/json"

	"act/api/internal/svc"
	"act/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type AddProcdefLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewAddProcdefLogic(ctx context.Context, svcCtx *svc.ServiceContext) *AddProcdefLogic {
	return &AddProcdefLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *AddProcdefLogic) AddProcdef(req *types.SaveProcdef) (resp *types.CommonResponse, err error) {
	if len(req.FormId) == 0 {
		return types.GetErrorCommonResponse("业务表单数据未获取到！")
	}
	if req.Resource == nil || len(req.Resource.Name) == 0 {
		return types.GetErrorCommonResponse("节点数据未获取到！")
	}
	resource, err := json.Marshal(req.Resource)
	if err != nil {
		return types.GetErrorCommonResponse(err.Error())
	}
	reply, err := l.svcCtx.Rpc.SaveProcDef(l.ctx, &act.ProcDefReq{
		UserId:      101,
		UserName:    "赵本山",
		Name:        req.Name,
		Code:        req.Code,
		FormId:      req.FormId,
		FormName:    req.FormName,
		Resource:    string(resource),
		RemainHours: req.RemainHours,
	})
	if err != nil {
		return types.GetErrorCommonResponse(err.Error())
	}
	return types.GetCommonResponse(err, reply)
}
