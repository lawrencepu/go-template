package mysql

import (
	"database/sql/driver"
	"errors"
	"fmt"
	"strings"
	"time"
)

type Model struct {
	ID        uint `gorm:"primarykey" json:"id"`
	CreatedAt Time `json:"created_at"`
	UpdatedAt Time `json:"updated_at"`
	DeletedAt *Time `gorm:"index" json:"deleted_at"` // 空值不会显示`0001-01-01 00:00:00`
}

type Time time.Time

func (t *Time) UnmarshalJSON(data []byte) error {
	if string(data) == "null" {
		return nil
	}
	var err error
	//前端接收的时间字符串
	str := string(data)
	//去除接收的str收尾多余的"
	timeStr := strings.Trim(str, "\"")
	t1, err := time.Parse("2006-01-02 15:04:05", timeStr)
	*t = Time(t1)
	return err
}

func (t Time) MarshalJSON() ([]byte, error) {
	formatted := fmt.Sprintf("\"%v\"", time.Time(t).Format("2006-01-02 15:04:05"))
	return []byte(formatted), nil
}

func (t Time) Value() (driver.Value, error) {
	// MyTime 转换成 time.Time 类型
	tTime := time.Time(t)
	return tTime.Format("2006-01-02 15:04:05"), nil
}

func (t *Time) Scan(v interface{}) error {
	switch vt := v.(type) {
	case time.Time:
		// 字符串转成 time.Time 类型
		*t = Time(vt)
	default:
		return errors.New("类型处理错误")
	}
	return nil
}

func (t *Time) String() string {
	return fmt.Sprintf("%s", time.Time(*t).String())
}
