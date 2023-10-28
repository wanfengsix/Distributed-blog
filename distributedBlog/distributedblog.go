package main

import (
	"flag"
	"fmt"

	"distributedBlog/internal/config"
	"distributedBlog/internal/handler"
	auth "distributedBlog/internal/logic/Auth"
	"distributedBlog/internal/svc"

	"github.com/zeromicro/go-zero/core/conf"
	"github.com/zeromicro/go-zero/rest"
)

var configFile = flag.String("f", "etc/distributedblog-api.yaml", "the config file")

func main() {
	flag.Parse()

	var c config.Config
	conf.MustLoad(*configFile, &c)

	server := rest.MustNewServer(c.RestConf)
	defer server.Stop()
	var a auth.Auth
	p := a.GetProtectionlist()
	fmt.Println(p)
	ctx := svc.NewServiceContext(c)
	handler.RegisterHandlers(server, ctx)
	fmt.Printf("Starting server at %s:%d...\n", c.Host, c.Port)
	server.Start()
}
