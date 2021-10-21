package mysql

type MemberLoginLog struct {
	Model
	MemberId       uint `gorm:"column:member_id;comment:'用户id'" json:"member_id"`
	Ip       string `gorm:"column:ip;comment:'登录ip'" json:"ip"`
	UserAgent string `gorm:"column:user_agent;comment:'用户代理'" json:"user_agent"`
	Platform uint8 `gorm:"column:platform;default:1;comment:'登录端；1-pc;2-h5;3-app'" json:"platform"`
}
