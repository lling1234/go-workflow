package svc

import (
	"act/rpc/internal/config"
	"act/rpc/internal/store"
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
