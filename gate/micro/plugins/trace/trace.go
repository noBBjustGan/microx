package trace

import (
	"net/http"

	"github.com/micro/cli"
	"github.com/micro/micro/plugin"
	"github.com/opentracing/opentracing-go"

	"microx/gate/micro/config"
	"microx/pkg/log"
	"microx/pkg/tracer"
)

type trace struct {
}

func (*trace) Flags() []cli.Flag {
	return nil
}

func (*trace) Commands() []cli.Command {
	return nil
}

func (*trace) Handler() plugin.Handler {
	return func(h http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			log.Infof("trace plugins received: %s %s", r.Method, r.RequestURI)
			spanCtx, _ := opentracing.GlobalTracer().Extract(opentracing.HTTPHeaders, opentracing.HTTPHeadersCarrier(r.Header))
			sp := opentracing.GlobalTracer().StartSpan(r.URL.Path, opentracing.ChildOf(spanCtx))
			defer sp.Finish()

			if err := opentracing.GlobalTracer().Inject(
				sp.Context(),
				opentracing.HTTPHeaders,
				opentracing.HTTPHeadersCarrier(r.Header)); err != nil {
			}

			h.ServeHTTP(w, r)
		})
	}
}

func (*trace) Init(*cli.Context) error {
	t, err := tracer.Init("gate", config.App.TracerAddr)
	if err != nil {
		log.Error(err)
		return nil
	}

	opentracing.SetGlobalTracer(t)
	return nil
}

func (*trace) String() string {
	return "trace"
}

func NewPlugin() plugin.Plugin {
	return new(trace)
}
