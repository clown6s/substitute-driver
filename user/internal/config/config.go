package config

import (
	"github.com/zeromicro/go-zero/core/stores/cache"
	"github.com/zeromicro/go-zero/zrpc"
)

type Config struct {
	zrpc.RpcServerConf

	Mysql struct {
		DataSource string
	}

	CacheRedis cache.CacheConf

	// JWT 签名配置（字段名不能包含 "auth"，否则与内嵌的 Auth bool 冲突）
	JWTConfig struct {
		AccessSecret string `json:",optional"`
		AccessExpire int64  `json:",default=86400"`
	}
}
