package wflow

import (
	"context"
	"encoding/json"
	"log"
	"strings"

	"go-wflow/app/wflow/api/internal/svc"
	"go-wflow/app/wflow/api/internal/types"
	"go-wflow/common/flow"
	"go-wflow/kernel"

	// "go-wflow/common/flow"

	"github.com/qkbyte/go-zero/core/logx"
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

// TODO 结构体互转
func (l *AddProcDefLogic) AddProcDef(req *types.SaveProcDefReq) (resp *types.CommonResponse, err error) {
	if req.FormId == 0 {
		return types.GetErrorCommonResponse("业务表单数据未获取到！")
	}
	if len(req.Resource) == 0 {
		return types.GetErrorCommonResponse("节点数据未获取到！")
	}
	//  string(req.Resource)转为node *flow.Node
	req.Resource = strings.Replace(req.Resource, "\\", "", -1)
	var resource *flow.Node
	json.Unmarshal([]byte(req.Resource), &resource)
	if flow.IfProcessConfigIsValid(resource) != nil {
		return types.GetErrorCommonResponse(err.Error())
	}
	log.Println("req 11111", req)
	p, err := l.svcCtx.Rpc.AddProcDef(l.ctx, &kernel.SaveProcDefReq{
		Id:          req.Id,
		Name:        req.Name,
		Code:        req.Code,
		FormId:      req.FormId,
		FormName:    req.FormName,
		AppId:       req.AppId,
		AppName:     req.AppName,
		RemainHours: req.RemainHours,
		Resource:    req.Resource,
		Version:     req.Version,
	})
	if err != nil {
		return types.GetErrorCommonResponse(err.Error())
	}
	return types.GetSuccessCommonResponse(p)
}
