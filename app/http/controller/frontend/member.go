package frontend

import (
	"github.com/gin-gonic/gin"
	"go-template/app/services/member"
)

func (ctrl *controller) Register(c *gin.Context)  {
	var registerFormData member.RegisterForm
	if err := c.ShouldBindJSON(&registerFormData); err != nil {
		ctrl.Fail(c, err)
		return
	}

	register, err := member.Register(c, &registerFormData)
	if err != nil {
		ctrl.Fail(c, err)
		return
	}

	ctrl.Success(c, register)
}

func (ctrl *controller) Login(c *gin.Context) {
	var loginForm member.LoginForm
	if err := c.ShouldBindJSON(&loginForm); err != nil {
		ctrl.Fail(c, err)
		return
	}

	memberData, err := member.Login(c, &loginForm)
	if err != nil {
		ctrl.Fail(c, err)
		return
	}

	ctrl.Success(c, memberData)
}