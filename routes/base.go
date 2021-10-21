package routes

import (
	"github.com/gin-gonic/gin"
	"go-template/app/http/controller"
)

func Register(r *gin.Engine) {
	r.GET("/ready", controller.Ctrl.Ready)
	r.GET("/version", controller.Ctrl.Version)

	// 用户相关路由注册
	MemberRouter(r)
}
