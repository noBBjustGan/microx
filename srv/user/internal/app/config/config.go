package config

import (
	"microx/common/config"
	conf "microx/pkg/config"
)

var (
	App    config.App
	Broker config.Broker
	Consul config.Consul
	Mysql  config.Mysql
	Logger config.Logger
)

func Init(addr string) {
	conf.Init(addr, "srv-user")
	conf.Get(&App, "app")
	conf.Get(&Broker, "broker")
	conf.Get(&Consul, "consul")
	conf.Get(&Mysql, "mysql")
	conf.Get(&Logger, "logger")
}
