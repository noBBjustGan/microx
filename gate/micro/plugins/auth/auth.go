package auth

import (
	"encoding/json"
	"net/http"
	"strings"
	"time"

	"github.com/micro/cli"
	"github.com/micro/go-micro/cmd"
	"github.com/micro/go-micro/util/ctx"
	"github.com/micro/micro/plugin"

	"microx/common/errors"
	"microx/pkg/log"
	passport "microx/srv/passport/api"
)

var (
	passportClient passport.PassportService
)

type auth struct {
}

func (*auth) Flags() []cli.Flag {
	return nil
}

func (*auth) Commands() []cli.Command {
	return nil
}

func (*auth) Handler() plugin.Handler {
	return func(h http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			log.Infof("auth plugins received: %s %s", r.Method, r.RequestURI)

			if strings.HasPrefix(r.URL.Path, "/xdd/passport/smslogin") ||
				strings.HasPrefix(r.URL.Path, "/xdd/passport/sms") ||
				strings.HasPrefix(r.URL.Path, "/xdd/passport/login") ||
				strings.HasPrefix(r.URL.Path, "/xdd/passport/oauthlogin") {
				h.ServeHTTP(w, r)
				return
			}

			cx := ctx.FromRequest(r)

			_, err := passportClient.ValidateToken(cx, &passport.TokenRequest{})
			if err != nil {
				log.Error(err)
				writeError(w, err)
				return
			}
			// 运行到此说明token认证通过
			h.ServeHTTP(w, r)
		})
	}
}

func (*auth) Init(*cli.Context) error {
	passportClient = passport.NewPassportService("", *cmd.DefaultCmd.Options().Client)
	return nil
}

func (*auth) String() string {
	return "auth"
}

func writeError(w http.ResponseWriter, err error) {
	if len(w.Header().Get("Content-Type")) == 0 {
		w.Header().Set("Content-Type", "application/json")
	}

	e := errors.Parse(err.Error())
	response := map[string]interface{}{
		"t": time.Now().UnixNano(),
	}
	response["errno"] = e.Errno
	response["errmsg"] = e.Errmsg

	if err := json.NewEncoder(w).Encode(response); err != nil {
		http.Error(w, err.Error(), 500)
		return
	}
}

func NewPlugin() plugin.Plugin {
	return new(auth)
}
