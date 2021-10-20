package tools

import (
	"fmt"
	"github.com/sirupsen/logrus"
)

var Logger *logrus.Logger

type Writer struct {
	*logrus.Logger
}

func InitLogger() {
	Logger = logrus.New()
	// 日期格式设置
	Logger.SetFormatter(&logrus.JSONFormatter{
		TimestampFormat: "2006-01-02 15:04:05",
	})
}

func (writer *Writer)Printf(format string, v...interface{})  {
	logstr  := fmt.Sprintf(format, v...)
	//利用loggus记录日志
	writer.Info(logstr)
}