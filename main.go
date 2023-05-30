package main

import (
	"flag"
	"fmt"
	"front-gin/config/env"
	"front-gin/config/etc"
	"front-gin/middleware/global/log"
	"front-gin/router"
	"front-gin/rpc/server/svc"
	"github.com/zeromicro/go-zero/core/conf"
	"strconv"
)

var configFile = flag.String("f", "config/etc/front-api.yaml", "the config file")

func main() {
	flag.Parse()
	var c etc.Config
	conf.MustLoad(*configFile, &c)
	ctx := svc.NewServiceContext(c)

	config := env.InitConfig()
	r := router.InitRouter(ctx)

	// 使用日志中间件
	r.Use(log.LoggerMiddleware(config))

	// int 转 string
	httpPort := strconv.Itoa(config.BaseConfig.HTTP_PORT)
	err := r.Run(":" + httpPort)
	if err != nil {
		fmt.Println("服务器启动失败！")
	}
}