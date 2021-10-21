package routes

import (
	"github.com/gin-gonic/gin"
	"go-template/app/http/controller/frontend"
)

func MemberRouter(r *gin.Engine)  {
	r.POST("/member/register", frontend.Ctrl.Register)
	r.POST("/member/login", frontend.Ctrl.Login)
}
