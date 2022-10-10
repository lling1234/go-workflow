package logic

import (
	"context"
	"log"

	"go-wflow/app/wflow/rpc/internal/svc"
	"go-wflow/kernel"
	"go-wflow/kernel/core/wflow"

	"github.com/qkbyte/go-zero/core/logx"
)

type AddProcDefLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
	WflowInterface wflow.WflowInterface
}

func NewAddProcDefLogic(ctx context.Context, svcCtx *svc.ServiceContext) *AddProcDefLogic {
	return &AddProcDefLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
		// TODO personId 从token中获取
		WflowInterface: wflow.New(111, svcCtx.CommonStore.Client, ctx),
	}
}

func (l *AddProcDefLogic) AddProcDef(in *kernel.SaveProcDefReq) (*kernel.Procdef, error) {
	log.Println("AddProcDef rpc 222", in)
	
	
	return l.WflowInterface.CreateProcDef(in)
}