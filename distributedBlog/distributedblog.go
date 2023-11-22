package main

import (
	"distributedBlog/internal/config"
	"distributedBlog/internal/handler"
	"distributedBlog/internal/handler/login"
	"distributedBlog/internal/handler/regist"
	"distributedBlog/internal/svc"
	"flag"
	"fmt"

	"github.com/zeromicro/go-zero/core/conf"
	"github.com/zeromicro/go-zero/rest"
)

var configFile = flag.String("f", "etc/distributedblog-api.yaml", "the config file")

func main() {
	//参数传递
	flag.Parse()
	//配置加载
	var c config.Config
	conf.MustLoad(*configFile, &c)
	//服务
	server := rest.MustNewServer(c.RestConf)
	defer server.Stop()
	//加载配置生成context
	ctx := svc.NewServiceContext(c)

	//请求处理
	handler.RegisterHandlers(server, ctx)
	login.RegisterHandlers(server, ctx)
	regist.RegisterHandlers(server, ctx)
	fmt.Printf("Starting server at %s:%d...\n", c.Host, c.Port)
	server.Start()
}
