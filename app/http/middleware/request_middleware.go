package middleware

import (
	"bytes"
	"github.com/gin-gonic/gin"
	"github.com/go-basic/uuid"
	"go-template/tools"
	"io/ioutil"
)

// 设置请求标识
func RequestTrace(c *gin.Context) {
	c.Header("Content-Type", "application/json; charset=utf-8")
	requestId := c.GetHeader("request_id")
	if requestId== "" {
		c.Request.Header.Set("request_id", uuid.New())
	}
	c.Next()
}

// 初始化日志 打印请求参数,设置日志request id
func SetLog(c *gin.Context)  {
	requestId := c.GetHeader("request_id")
	tools.Logger.AddHook(tools.NewLogHook(requestId))
	requestData, _ := c.GetRawData()
	c.Request.Body = ioutil.NopCloser(bytes.NewBuffer(requestData))
	tools.Logger.Info(string(requestData))
	c.Next()
}
