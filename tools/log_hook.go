package tools

import "github.com/sirupsen/logrus"

type LogHook struct {
	RequestId string
}

// 设置request_id
func (hook *LogHook) Fire(entry *logrus.Entry) error {
	entry.Data["request_id"] = hook.RequestId
	return nil
}

func (hook *LogHook) Levels() []logrus.Level {
	return logrus.AllLevels
}

func NewLogHook(requestId string) *LogHook {
	return &LogHook{RequestId: requestId}
}
