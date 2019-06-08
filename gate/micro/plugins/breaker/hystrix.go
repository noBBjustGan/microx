package breaker

import (
	"errors"
	"fmt"
	"net/http"

	"github.com/afex/hystrix-go/hystrix"
	"github.com/micro/cli"
	"github.com/micro/micro/plugin"

	"microx/gate/micro/config"
	"microx/pkg/log"
)

type breaker struct {
}

func (*breaker) Flags() []cli.Flag {
	return nil
}

func (*breaker) Commands() []cli.Command {
	return nil
}

func (*breaker) Handler() plugin.Handler {
	return func(h http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			log.Infof("breaker plugins received: %s %s", r.Method, r.RequestURI)
			name := r.Method + " " + r.RequestURI
			err := hystrix.Do(name, func() error {
				sct := &statusCodeTracker{ResponseWriter: w, status: http.StatusOK}
				h.ServeHTTP(sct.wrappedResponseWriter(), r)

				if sct.status >= http.StatusBadRequest {
					errmsg := fmt.Sprintf("%d %s", sct.status, http.StatusText(sct.status))
					return errors.New(errmsg)
				}
				return nil
			}, nil)
			if err != nil {
				log.Error(err)
				return
			}
		})
	}
}

func (*breaker) Init(*cli.Context) error {
	if config.Hystrix.Timeout != 0 {
		hystrix.DefaultTimeout = config.Hystrix.Timeout
	}

	if config.Hystrix.MaxConcurrent != 0 {
		hystrix.DefaultMaxConcurrent = config.Hystrix.MaxConcurrent
	}

	if config.Hystrix.RequestVolumeThreshold != 0 {
		hystrix.DefaultVolumeThreshold = config.Hystrix.RequestVolumeThreshold
	}

	if config.Hystrix.SleepWindow != 0 {
		hystrix.DefaultSleepWindow = config.Hystrix.SleepWindow
	}

	if config.Hystrix.ErrorPercentThreshold != 0 {
		hystrix.DefaultErrorPercentThreshold = config.Hystrix.ErrorPercentThreshold
	}

	return nil
}

func (*breaker) String() string {
	return "breaker"
}

func NewPlugin() plugin.Plugin {
	return new(breaker)
}
