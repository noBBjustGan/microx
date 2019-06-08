package main

import (
	"github.com/micro/go-plugins/micro/gzip"
	"github.com/micro/micro/api"

	"microx/gate/micro/plugins/auth"
	"microx/gate/micro/plugins/breaker"
	"microx/gate/micro/plugins/metrics"
	"microx/gate/micro/plugins/trace"
)

func init() {
	// plugins.Register(gzip.NewPlugin())

	api.Register(gzip.NewPlugin())
	api.Register(trace.NewPlugin())
	api.Register(auth.NewPlugin())
	api.Register(breaker.NewPlugin())
	api.Register(metrics.NewPlugin())
}
