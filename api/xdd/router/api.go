package router

import (
	"github.com/gin-gonic/gin"
	"microx/api/xdd/controller"
	"microx/api/xdd/middleware/tracer"
)

func RegisterAPI_v1(r *gin.Engine) {
	r.Use(tracer.Tracer())
	v1 := r.Group("xdd/")

	// 注册路由
	// 增加新的路由在此注册！！！！！
	RegisterPassportRouter(v1, &controller.PassportController{})
}
