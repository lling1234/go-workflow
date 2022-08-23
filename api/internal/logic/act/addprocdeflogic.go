package act

import (
	"context"
	"log"

	"act/api/internal/svc"
	"act/api/internal/types"
	"act/rpc/act"

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

func (l *AddProcDefLogic) AddProcDef(req *types.SaveProcProdef) (resp *types.CommonResponse, err error) {
	// todo: add your logic here and delete this line
	log.Println("AddProcDef-----api 11111")
	resps, err := l.svcCtx.ActRpc.SaveProcDef(l.ctx, &act.ProcDefReq{
		Name:        req.Name,
		Code:        req.Code,
		YewuFormId:  req.YewuFormId,
		YewuName:    req.YewuName,
		RemainHours: req.RemainHours,
		Resource:    req.Resource,
	})
	if err != nil {
		return nil, err
	}
	log.Println("resps",resps)
	return &types.CommonResponse{
		Code: 200,
		Data: resps,
		Msg: "ok",
		Success: true,
	},nil
}
