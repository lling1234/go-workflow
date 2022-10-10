package config

import (
	"github.com/qkbyte/go-zero/rest"
	"github.com/qkbyte/go-zero/zrpc"
)

type Config struct {
	rest.RestConf
	Rpc zrpc.RpcClientConf
}