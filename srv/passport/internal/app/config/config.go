package config

import (
	"microx/common/config"
	conf "microx/pkg/config"
)

var (
	App   config.App
	Mysql config.Mysql
	//Redis  config.Redis
	Consul config.Consul
	Broker config.Broker
	Logger config.Logger
)

func Init(addr string) {
	conf.Init(addr, "srv-passport")
	conf.Get(&App, "app")
	conf.Get(&Broker, "broker")
	conf.Get(&Consul, "consul")
	conf.Get(&Mysql, "mysql")
	//conf.Get(&Redis, "redis")
	conf.Get(&Logger, "logger")

}
