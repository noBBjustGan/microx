package config

import (
	"microx/common/config"
	conf "microx/pkg/config"
)

var (
	App     config.App
	Consul  config.Consul
	Logger  config.Logger
	Hystrix config.Hystrix
)

func Init(addr string) {
	conf.Init(addr, "gate-micro")
	conf.Get(&App, "app")
	conf.Get(&Consul, "consul")
	conf.Get(&Logger, "logger")
	conf.Get(&Hystrix, "hystrix")
}
