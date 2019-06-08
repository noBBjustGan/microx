package app

import (
	"io/ioutil"
	"os"
	"time"

	"github.com/google/uuid"
	"github.com/micro/cli"
	"github.com/micro/go-micro"
	"github.com/micro/go-micro/broker"
	"github.com/micro/go-micro/registry"
	"github.com/micro/go-micro/registry/consul"
	"github.com/micro/go-plugins/broker/stan"
	"github.com/micro/go-plugins/wrapper/trace/opentracing"

	"microx/pkg/capx"
	"microx/pkg/log"
	"microx/pkg/tracer"
	passport "microx/srv/passport/api"
	"microx/srv/passport/internal/app/config"
	"microx/srv/passport/internal/app/factory"
)

type Application interface {
	Run()
}

type application struct {
	service micro.Service
}

func getClientID() (cid string) {
	var (
		f   *os.File
		err error
		b   []byte
	)

	if f, err = os.OpenFile("cid", os.O_CREATE|os.O_RDWR, 0666); err != nil {
		panic(err)
	}
	if b, err = ioutil.ReadAll(f); err != nil {
		panic(err)
	}
	if len(b) == 0 {
		cid = uuid.New().String()
		f.Write([]byte(cid))
	} else {
		cid = string(b)
	}
	f.Close()
	return
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

	// Broker
	b := stan.NewBroker(
		broker.Addrs(config.Broker.Addrs...),
		stan.ClientID(getClientID()),
		stan.ClusterID(config.Broker.ClusterID),
		stan.ConnectRetry(true),
	)

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
		micro.Broker(b),
		micro.Registry(reg),
		micro.WrapHandler(opentracing.NewHandlerWrapper(t)),
	)

	// Factory
	factory.Init(app.service)

	// Register Handler
	passport.RegisterPassportHandler(app.service.Server(), factory.GetPassportHandler())

	capx.Init(factory.GetEngine())
	return app
}

func (a *application) Run() {
	// Run service
	if err := a.service.Run(); err != nil {
		log.Fatal(err)
	}
}
