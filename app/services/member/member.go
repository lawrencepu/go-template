package member

import (
	"encoding/json"
	"errors"
	"github.com/gin-gonic/gin"
	"go-template/app/models/mysql"
	"go-template/app/models/redis"
	"go-template/common"
	"go-template/tools"
	"strconv"
	"time"
)

// 注册验证器
type RegisterForm struct {
	Nickname string `form:"nickname" json:"nickname" binding:"required"`
	Phone    string `form:"phone" json:"phone" binding:"required"`
	Password string `form:"password" json:"password" binding:"required"`
}

// 登录验证器
type LoginForm struct {
	Phone    string `form:"phone" json:"phone" binding:"required"`
	Password string `form:"password" json:"password" binding:"required"`
}

// 注册用户
func Register(ctx *gin.Context, form *RegisterForm) (*mysql.Member, error) {
	var member mysql.Member
	result := mysql.DB.
		Where("phone", form.Phone).
		First(&member)
	if result.RowsAffected > 0 {
		return nil, errors.New("该账号已注册！")
	}
	member.Password = common.Md5(form.Password + mysql.PasswordSalt)
	member.Phone = form.Phone
	member.Nickname = form.Nickname
	member.Ip = ctx.ClientIP()
	member.Avatar = mysql.DefaultAvatar
	register := mysql.DB.Create(&member)
	if register.Error != nil || register.RowsAffected <= 0 {
		return nil, register.Error
	}
	member.Password = ""
	return &member, nil
}

// 用户登录
func Login(ctx *gin.Context, form *LoginForm) (*mysql.Member, error) {
	var member mysql.Member
	result := mysql.DB.Where("phone", form.Phone).First(&member)

	if result.RowsAffected <= 0 {
		return nil, errors.New("账号不存在，请先注册！")
	}

	if member.Status == mysql.StatusDisabled {
		return nil, errors.New("账号已被禁用！")
	}

	pwd := common.Md5(form.Password + mysql.PasswordSalt)
	if member.Password != pwd {
		return &member, errors.New("账号或密码错误！")
	}

	strCmd := redis.Rdb.HGet(ctx, mysql.MemberTokenCacheKey, strconv.Itoa(int(member.ID)))
	if strCmd.Val() != "" {
		redis.Rdb.HDel(ctx, mysql.MemberTokenCacheKey, strconv.Itoa(int(member.ID)))
		redis.Rdb.Del(ctx, strCmd.Val())
	}

	member.Token = common.Md5(time.Now().String() + strconv.Itoa(int(member.ID)))
	memberJson, _ := json.Marshal(&member)
	cmd := redis.Rdb.Set(ctx, mysql.MemberCachePrefix + member.Token, memberJson, mysql.MemberCacheTtl)
	if cmd.Err() != nil {
		return &mysql.Member{}, cmd.Err()
	}
	redis.Rdb.HSet(ctx, mysql.MemberTokenCacheKey, strconv.Itoa(int(member.ID)), mysql.MemberCachePrefix + member.Token)
	member.Password = ""
	go LoginLog(ctx, member)
	return &member, nil
}

func LoginLog(ctx *gin.Context, member mysql.Member)  {
	var loginLog mysql.MemberLoginLog
	loginLog.MemberId = member.ID
	loginLog.Ip = ctx.ClientIP()
	loginLog.UserAgent = ctx.GetHeader("user-agent")
	mysql.DB.Create(&loginLog)
	tools.Logger.Info("login log", loginLog)
}