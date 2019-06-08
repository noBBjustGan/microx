package factory

import (
	"github.com/micro/go-micro/client"
	passport "microx/srv/passport/api"
)

var (
	PassportClient passport.PassportService
)

func Init(opt ...client.Option) {
	cli := client.NewClient(opt...)
	PassportClient = passport.NewPassportService("", cli)
}
