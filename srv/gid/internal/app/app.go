package app

import (
	"time"

	"github.com/micro/cli"
	"github.com/micro/go-micro"
	"github.com/micro/go-micro/registry"
	"github.com/micro/go-micro/registry/consul"
	"github.com/micro/go-plugins/wrapper/trace/opentracing"

	"microx/pkg/log"
	"microx/pkg/tracer"
	gid "microx/srv/gid/api"
	"microx/srv/gid/internal/app/config"
	"microx/srv/gid/internal/app/factory"
)

type Application interface {
	Run()
}

type application struct {
	service micro.Service
}

func NewApplication() Application {
	var configAddr string

	app := new(application)
	app.service = micro.NewService(
		micro.Flags(cli.StringFlag{
			Name:  "config_address",
			Usage: "Config center address",
		}),
	)
	// 获取配置中心IP地址
	app.service.Init(
		micro.Action(func(c *cli.Context) {
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

	// Initialise service
	app.service.Init(
		micro.Name(config.App.Name),
		micro.Version(config.App.Version),
		micro.RegisterTTL(time.Second*time.Duration(config.App.RegisterTTL)),
		micro.RegisterInterval(time.Second*time.Duration(config.App.RegisterInterval)),
		micro.Registry(reg),
		micro.WrapHandler(opentracing.NewHandlerWrapper(t)),
	)

	// Factory
	factory.Init()
	// Register Handler
	gid.RegisterGidHandler(app.service.Server(), factory.GetGidHandler())
	return app
}

func (a *application) Run() {
	// Run service
	if err := a.service.Run(); err != nil {
		log.Fatal(err)
	}
}
