package auth

import (
	"errors"
	"forum/app/models/user"
	"forum/pkg/helpers"
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
