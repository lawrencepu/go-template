package controllers

import (
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"go-template/app/models/mysql"
	"go-template/app/models/redis"
	"go-template/common"
	"go-template/config"
	"net/http"
)

var Ctrl = &BaseController{}

type BaseController struct {
	common.Response
}



// version
func (ctrl *BaseController)Version(c *gin.Context) {
	baseConfig, err := config.GetBaseConf()
	if err != nil {
		panic(err)
	}
	str := baseConfig.AppName + ":" + baseConfig.AppVersion
	ctrl.Success(c, str)
	return
}

// 存活
func (ctrl *BaseController) Ready(c *gin.Context) {
	db, err := mysql.DB.DB()
	if err != nil {
		logrus.Info("mysql gg")
		c.AbortWithError(http.StatusInternalServerError, err)
	}
	if err = db.Ping(); err != nil {
		logrus.Info("mysql gg")
		c.AbortWithError(http.StatusInternalServerError, err)
	}

	err = redis.Rdb.Ping(c).Err()

	if err != nil {
		logrus.Info("redis gg", err)
		c.AbortWithError(http.StatusInternalServerError, err)
	}

	ctrl.Success(c, "ok")

}
