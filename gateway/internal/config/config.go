// Code scaffolded by goctl. Safe to edit.
// goctl 1.10.1

package config

import (
	"github.com/zeromicro/go-zero/rest"
	"github.com/zeromicro/go-zero/zrpc"
)

type Config struct {
	rest.RestConf

	ZapLog struct {
		Level string `json:",default=info"`
	}

	UserRpc   zrpc.RpcClientConf
	DriverRpc zrpc.RpcClientConf
	MapRpc    zrpc.RpcClientConf

	// JWT 鉴权配置（rest.RestConf 无内置 Auth，需手动添加）
	Auth struct {
		AccessSecret string `json:",optional"`
		AccessExpire int64  `json:",optional"`
	}
}
