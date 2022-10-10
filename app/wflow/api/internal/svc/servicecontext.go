package svc

import (
	"go-wflow/app/wflow/api/internal/config"
	wflowclient "go-wflow/app/wflow/rpc/wflow"
	"github.com/qkbyte/go-zero/zrpc"
)

type ServiceContext struct {
	Config config.Config
	Rpc    wflowclient.Wflow
}

func NewServiceContext(c config.Config) *ServiceContext {
	return &ServiceContext{
		Config: c,
		Rpc: wflowclient.NewWflow(zrpc.MustNewClient(c.Rpc)),
	}
}