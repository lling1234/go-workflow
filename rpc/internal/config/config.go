package config

import (
	"act/common/models"
	"github.com/zeromicro/go-zero/zrpc"
)

type Config struct {
	zrpc.RpcServerConf
	Store models.DbStoreConfig
}
