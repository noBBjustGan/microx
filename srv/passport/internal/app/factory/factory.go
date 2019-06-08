package factory

import (
	"fmt"
	"github.com/go-xorm/xorm"

	"github.com/micro/go-micro"

	"microx/pkg/orm"
	gid "microx/srv/gid/api"
	"microx/srv/passport/internal/app/config"
	"microx/srv/passport/internal/app/handler"
	"microx/srv/passport/internal/app/service"
	"microx/srv/passport/internal/domain/repository"
	"microx/srv/passport/internal/infra/persistence/mysql"
	user "microx/srv/user/api"
)

var (
	// engine
	engine *xorm.Engine
	// client
	gidClient  gid.GidService
	userClient user.UserService
	// handler
	passportHandler *handler.PassportHandler
	// service
	passportService *service.PassportService
	// repo
	userRepo      repository.UserRepo
	userTokenRepo repository.UserTokenRepo
)

func Init(s micro.Service) {
	gidClient = gid.NewGidService("", s.Client())
	userClient = user.NewUserService("", s.Client())

	c := orm.Config{
		DriverName:     "mysql",
		DataSourceName: config.Mysql.DataSource,
		MaxIdleConn:    config.Mysql.MaxIdle,
		MaxOpenConn:    config.Mysql.MaxOpen,
	}
	fmt.Println(c)
	engine = orm.GetEngine(c)
	if engine == nil {
		panic("database init error")
	}
	userRepo = mysql.NewUserRepo(engine)
	userTokenRepo = mysql.NewUserTokenRepo(engine)
	passportService = service.NewPassportService(engine, gidClient, userClient, userRepo, userTokenRepo)
	passportHandler = handler.NewPassportHandler(passportService)
}

func GetEngine() *xorm.Engine {
	return engine
}

func GetPassportHandler() *handler.PassportHandler {
	return passportHandler
}

func GetGidClient() gid.GidService {
	return gidClient
}
