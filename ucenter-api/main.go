package main

import (
	"flag"
	"fmt"
	"github.com/zeromicro/go-zero/core/logx"

	"ucenter-api/internal/config"
	"ucenter-api/internal/handler"
	"ucenter-api/internal/svc"

	"github.com/zeromicro/go-zero/core/conf"
	"github.com/zeromicro/go-zero/rest"
)

var configFile = flag.String("f", "etc/conf.yaml", "the config file")

func main() {
	flag.Parse()
	// 日志打印格式
	logx.MustSetup(logx.LogConf{
		Encoding: "plain",
		Stat:     false,
	})
	var c config.Config
	conf.MustLoad(*configFile, &c)

	server := rest.MustNewServer(c.RestConf)
	defer server.Stop()

	ctx := svc.NewServiceContext(c)
	router := handler.NewRouters(server)
	handler.RegisterHandlers(router, ctx)

	fmt.Printf("Starting server at %s:%d...\n", c.Host, c.Port)
	server.Start()
}
