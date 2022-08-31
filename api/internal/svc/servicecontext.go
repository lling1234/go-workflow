package svc

import (
	"act/api/internal/config"
	"act/rpc/act"
	"github.com/zeromicro/go-zero/zrpc"
)

type ServiceContext struct {
	Config config.Config
	Rpc    act.Act
}

func NewServiceContext(c config.Config) *ServiceContext {
	return &ServiceContext{
		Config: c,
		Rpc:    act.NewAct(zrpc.MustNewClient(c.Rpc)),
	}
}
