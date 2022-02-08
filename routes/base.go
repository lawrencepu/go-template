package routes

import (
	"github.com/gin-gonic/gin"
	"go-template/app/http/controllers"
	"go-template/app/http/controllers/backend"
)

func Register(r *gin.Engine) {
	r.GET("/ready", controllers.Ctrl.Ready)
	r.GET("/version", controllers.Ctrl.Version)

	// 用户相关路由注册
	MemberRouter(r)
	// 商品相关
	r.GET("/backend/goods/create", backend.Ctrl.Create)
}
