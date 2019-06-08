package main

import (
	"os"
	"time"

	"github.com/micro/go-micro"
	"github.com/micro/go-micro/registry"
	"github.com/micro/go-micro/registry/consul"
	"github.com/micro/micro/cmd"
	"microx/gate/micro/config"
	"microx/pkg/log"
)

func main() {
	configAddr := os.Getenv("CONFIG_CENTER_ADDRESS")
	if configAddr == "" {
		configAddr = "127.0.0.1:9600"
	}
	config.Init(configAddr)
	log.Init(config.Logger)
	reg := consul.NewRegistry(func(op *registry.Options) {
		op.Addrs = config.Consul.Addrs
	})
	cmd.Init(
		micro.Name(config.App.Name),
		micro.Version(config.App.Version),
		micro.RegisterTTL(time.Second*time.Duration(config.App.RegisterTTL)),
		micro.RegisterInterval(time.Second*time.Duration(config.App.RegisterInterval)),
		micro.Registry(reg),
	)
}
