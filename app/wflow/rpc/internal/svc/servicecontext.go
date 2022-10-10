package svc

import (
	"go-wflow/app/wflow/rpc/internal/config"
	"go-wflow/common/store"
)

type ServiceContext struct {
	Config      config.Config
	CommonStore store.CommonStore
}

func NewServiceContext(c config.Config) *ServiceContext {
	return &ServiceContext{
		Config:      c,
		CommonStore: *store.NewCommonStore(c.Store),
	}
}
