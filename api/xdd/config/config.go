package config

import (
	"microx/common/config"
	conf "microx/pkg/config"
)

var (
	App    config.App
	Consul config.Consul
	Logger config.Logger
)

func Init(addr string) {
	conf.Init(addr, "api-xdd")
	conf.Get(&App, "app")
	conf.Get(&Consul, "consul")
	conf.Get(&Logger)
}
