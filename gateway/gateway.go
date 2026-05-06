// Code scaffolded by goctl. Safe to edit.
// goctl 1.10.1

package main

import (
	"flag"
	"log"

	"substitute-driver/common/zapx"
	"substitute-driver/gateway/internal/config"
	"substitute-driver/gateway/internal/handler"
	"substitute-driver/gateway/internal/middleware"
	"substitute-driver/gateway/internal/svc"

	"github.com/zeromicro/go-zero/core/conf"
	"github.com/zeromicro/go-zero/core/logx"
	"github.com/zeromicro/go-zero/rest"
)

var configFile = flag.String("f", "etc/gateway-api.yaml", "the config file")

func main() {
	flag.Parse()

	var c config.Config
	conf.MustLoad(*configFile, &c)

	// 初始化 zap 日志
	if err := zapx.Init(c.ZapLog.Level); err != nil {
		log.Fatalf("init zap logger failed: %v", err)
	}
	defer zapx.Sync()

	// 将 go-zero 日志桥接到 zap
	logx.SetWriter(zapx.NewZapLogWriter())

	// 设置全局错误处理（拦截 gRPC status error 转为结构化 JSON）
	middleware.SetupErrorHandler()

	server := rest.MustNewServer(c.RestConf)
	defer server.Stop()

	ctx := svc.NewServiceContext(c)
	handler.RegisterHandlers(server, ctx)

	zapx.Logger.Infof("Starting server at %s:%d...", c.Host, c.Port)
	server.Start()
}
