package factory

import (
	"fmt"
	"github.com/go-xorm/xorm"
	"microx/pkg/orm"
	"microx/srv/user/internal/app/config"
	"microx/srv/user/internal/app/handler"
	"microx/srv/user/internal/app/service"
	"microx/srv/user/internal/domain/repository"
	"microx/srv/user/internal/infra/persistence/mysql"
)

var (
	// engine
	engine *xorm.Engine
	// handler
	userHandler *handler.UserHandler
	// subscriber
	subscriber *handler.Subscriber
	// service
	userService *service.UserService
	// repo
	userRepo repository.UserRepo
)

func Init() {
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
	userService = service.NewUserService(userRepo)
	userHandler = handler.NewUserHandler(userService)
	subscriber = handler.NewSubscriber(engine, userService)
}

func GetUserHandler() *handler.UserHandler {
	return userHandler
}

func GetSubscriber() *handler.Subscriber {
	return subscriber
}

func GetEngine() *xorm.Engine {
	return engine
}
