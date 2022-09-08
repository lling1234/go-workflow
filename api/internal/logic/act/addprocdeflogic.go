package act

import (
	"act/api/flow"
	"act/api/internal/svc"
	"act/api/internal/types"
	"act/rpc/types/act"
	"context"
	"encoding/json"
	"log"

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
	if len(req.FormId) == 0 {
		return types.GetErrorCommonResponse("业务表单数据未获取到！")
	}
	if req.Resource == nil || len(req.Resource.Name) == 0 {
		return types.GetErrorCommonResponse("节点数据未获取到！")
	}
	if flow.IfProcessConifgIsValid(req.Resource) != nil {
		return types.GetErrorCommonResponse(err.Error())
	}
	resource, err := json.Marshal(req.Resource)
	if err != nil {
		return types.GetErrorCommonResponse(err.Error())
	}
	log.Println(111111)
	reply, err := l.svcCtx.Rpc.SaveProcDef(l.ctx, &act.SaveProcDefReq{
		//UserId:      101,
		//UserName:    "赵本山",
		Name:        req.Name,
		Code:        req.Code,
		FormId:      req.FormId,
		FormName:    req.FormName,
		Resource:    string(resource),
		RemainHours: req.RemainHours,
	})
	log.Println(222222)
	return types.GetCommonResponse(err, reply)
}
