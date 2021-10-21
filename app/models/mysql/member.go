package mysql

import (
	"time"
)

type Member struct {
	Model
	Phone    string `gorm:"column:phone;comment:'手机号'" json:"phone"`
	Avatar   string `gorm:"column:avatar;comment:'头像'" json:"avatar"`
	Password string `gorm:"column:password;comment:'密码'" json:"password,omitempty"`
	Nickname string `gorm:"column:nickname;comment:'用户昵称'" json:"nickname"`
	Integral uint `gorm:"column:integral;comment:'积分'" json:"integral"`
	Garden uint8 `gorm:"column:garden;default:1;comment:'性别；1-保密；2-男；3-女'" json:"garden"`
	Email string `gorm:"column:email;comment:'email'" json:"email"`
	Status   uint8  `gorm:"column:status;default:1;comment:'状态:1-启用;2禁用'" json:"status"`
	Ip       string `gorm:"column:ip;comment:'注册ip'" json:"ip"`
	Token    string `gorm:"-" json:"token"` //入库不会插入该字段
}

const DefaultAvatar = "https://d2.shopxo.vip/static/index/default/images/default-user-avatar.jpg"

// 禁用账户
const StatusDisabled = 2
// 启用账户
const StatusEnabled = 1

// 密码加盐
const PasswordSalt = "password"

// 登录信息缓存前缀
const MemberCachePrefix = "member_"

const MemberTokenCacheKey = "member_tokens"

// 登录过期时间
const MemberCacheTtl = 86400 * 30 * time.Second
