package factory

import (
	"microx/srv/gid/internal/app/handler"
)

var (
	// handler
	gidHandler *handler.Gid
)

func Init() {
	//gidHandler = handler.NewPassportHandler()
}

func GetGidHandler() *handler.Gid {
	return gidHandler
}
