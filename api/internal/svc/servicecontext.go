package svc

import (
	"act/api/internal/config"
	"act/rpc/actclient"
	"github.com/zeromicro/go-zero/zrpc"
)

type ServiceContext struct {
	Config config.Config
	Rpc    actclient.Act
}

func NewServiceContext(c config.Config) *ServiceContext {
	return &ServiceContext{
		Config: c,
		Rpc:    actclient.NewAct(zrpc.MustNewClient(c.Rpc)),
	}
}
