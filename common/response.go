package common

import (
	"github.com/gin-gonic/gin"
	"net/http"
)


type Response struct{}

type response struct {
	StatusCode uint32      `json:"status_code"`
	Message    string      `json:"message"`
	Data       interface{} `json:"data"`
	RequestID  string      `json:"request_id"`
}

func (res *Response) Success(c *gin.Context, data interface{}, msg ...string) {
	rr := new(response)
	rr.StatusCode = http.StatusOK
	if len(msg) == 0 {
		rr.Message = "success"
	} else {
		for _, v := range msg {
			rr.Message += "," + v
		}
	}
	rr.Data = data
	rr.RequestID = c.GetHeader("request_id")
	c.JSON(http.StatusOK, &rr)
}

func (res *Response) Fail(c *gin.Context, err error) {
	rr := new(response)
	rr.StatusCode = 400
	rr.Message = err.Error()
	rr.RequestID = c.GetHeader("request_id")
	c.JSON(http.StatusOK, &rr)
}

// RawJSONString json 数据返回
func (res *Response) RawJSONString(c *gin.Context, data string) {
	w := c.Writer
	w.WriteHeader(200)
	w.Header().Add("Content-Type", "application/json; charset=utf-8")
	_, _ = w.Write([]byte(data))
}

// RawString raw 数据返回
func (res *Response) RawString(c *gin.Context, data string) {
	w := c.Writer
	w.WriteHeader(200)
	_, _ = w.Write([]byte(data))
}
