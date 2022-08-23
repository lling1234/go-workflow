package svc

import (
	"act/api/internal/config"
	"act/rpc/act"

	"github.com/zeromicro/go-zero/zrpc"
)

type ServiceContext struct {
	Config config.Config
	ActRpc act.Act
}

func NewServiceContext(c config.Config) *ServiceContext {
	return &ServiceContext{
		Config: c,
		ActRpc: act.NewAct(zrpc.MustNewClient(c.ActRpc)),
	}
}
