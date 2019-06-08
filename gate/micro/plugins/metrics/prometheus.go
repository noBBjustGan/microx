package metrics

import (
	"github.com/micro/cli"
	"github.com/micro/micro/plugin"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"net/http"
)

type metrics struct {
}

func (*metrics) Flags() []cli.Flag {
	return nil
}

func (*metrics) Commands() []cli.Command {
	return nil
}

func (*metrics) Handler() plugin.Handler {
	ph := promhttp.Handler()
	return func(h http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			if r.URL.Path == "/metrics" {
				ph.ServeHTTP(w, r)
				return
			}

			h.ServeHTTP(w, r)
		})
	}
}

func (*metrics) Init(*cli.Context) error {
	return nil
}

func (*metrics) String() string {
	return "metrics"
}

func NewPlugin() plugin.Plugin {
	return new(metrics)
}
