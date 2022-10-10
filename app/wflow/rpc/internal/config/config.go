package config

import (
	"go-wflow/common/models"

	"github.com/qkbyte/go-zero/zrpc"
)

type Config struct {
	zrpc.RpcServerConf
	Store models.DbStoreConfig
}
