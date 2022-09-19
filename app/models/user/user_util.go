package user

import (
	"fmt"
	"forum/pkg/database"
)

// IsEmailExist 判断 Email 已被注册
func IsEmailExist(email string) bool {
	result := map[string]interface{}{}
	database.DB.Model(User{}).Where("email = ?", email).First(&result)
	fmt.Println(result)
	if result["email"] != email {
		return false
	}
	return true
}

// IsPhoneExist 判断手机号已被注册
func IsPhoneExist(phone string) bool {
	var user User
	database.DB.Model(User{}).Where("phone = ?", phone).First(&user)
	if user.Phone != phone {
		return false
	}
	return true
}
