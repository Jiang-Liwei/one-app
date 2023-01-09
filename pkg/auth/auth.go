package auth

import (
	"errors"
	"forum/app/models/user"
	"forum/pkg/helpers"
	"forum/pkg/logger"
	"github.com/gin-gonic/gin"
)

// LoginByAccount 账号登录
func LoginByAccount(account string, password string) (user.User, error) {
	if helpers.Empty(account) {
		return user.User{}, errors.New("账号不存在")
	}

	userInfo := user.GetByMulti(account)
	if userInfo.ID == 0 {
		return user.User{}, errors.New("密码错误")
	}

	if !userInfo.ComparePassword(password) {
		return user.User{}, errors.New("密码错误")
	}

	return userInfo, nil
}

// LoginByPhone 手机登录
func LoginByPhone(phone string) (user.User, error) {
	userInfo := user.GetByPhone(phone)

	return userInfo, nil
}

// User 获取登录用户信息
func User(c *gin.Context) user.User {

	userModel, ok := c.MustGet("user").(user.User)
	if !ok {
		logger.LogRecord(errors.New("无法获取用户"))
		return user.User{}
	}
	// db is now a *DB value
	return userModel
}

// UserID 获取当前登录用户 ID
func UserID(c *gin.Context) string {
	return c.GetString("user_id")
}
