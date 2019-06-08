package main

import (
	"time"

	"github.com/micro/cli"
	"github.com/micro/go-micro/client"
	"github.com/micro/go-micro/registry"
	"github.com/micro/go-micro/registry/consul"
	"github.com/micro/go-micro/web"
	"github.com/opentracing/opentracing-go"

	"microx/api/xdd/config"
	"microx/api/xdd/factory"
	"microx/api/xdd/router"
	"microx/pkg/log"
	"microx/pkg/tracer"
)

func main() {
	var configAddr string

	service := web.NewService(
		web.Flags(cli.StringFlag{
			Name:  "config_address",
			Usage: "Config center address",
		}),
	)

	// 获取配置中心IP地址
	service.Init(
		web.Action(func(c *cli.Context) {
			if configAddr = c.String("config_address"); configAddr == "" {
				configAddr = "127.0.0.1:9600"
			}
		}),
	)

	// Config与Logger要先初始化
	// Config
	config.Init(configAddr)
	// Logger
	log.Init(config.Logger)

	// Consul
	reg := consul.NewRegistry(func(op *registry.Options) {
		op.Addrs = config.Consul.Addrs
	})

	// Tracer
	t, err := tracer.Init(config.App.Name, config.App.TracerAddr)
	if err != nil {
		log.Error(err)
	}
	opentracing.SetGlobalTracer(t)

	service.Init(
		web.Name(config.App.Name),
		web.Version(config.App.Version),
		web.RegisterTTL(time.Second*time.Duration(config.App.RegisterTTL)),
		web.RegisterInterval(time.Second*time.Duration(config.App.RegisterInterval)),
		web.Registry(reg),
		web.Handler(router.Router()),
	)

	factory.Init(client.Registry(reg))
	// run service
	if err := service.Run(); err != nil {
		log.Fatal(err)
	}
}
