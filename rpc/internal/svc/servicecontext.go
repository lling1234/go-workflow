package svc

import (
	"act/rpc/internal/config"
	"act/rpc/internal/store"
)

type ServiceContext struct {
	Config   config.Config
	ActStore store.ActStore
}

func NewServiceContext(c config.Config) *ServiceContext {
	return &ServiceContext{
		Config:   c,
		ActStore: *store.NewActrStore(c.Store),
	}
}
