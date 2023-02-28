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

// GetByPhone 手机号获取用户信息
func GetByPhone(phone string) (user User) {
	database.DB.Where("phone = ?", phone).First(&user)

	return
}

// GetByMulti 通过 手机号/Email/用户名 来获取用户
func GetByMulti(loginID string) (user User) {
	database.DB.
		Where("phone = ?", loginID).
		Or("email = ?", loginID).
		Or("name = ?", loginID).
		First(&user)
	return
}

// Get 通过 ID 获取用户
func Get(id string) (userModel User) {
	database.DB.Where("id", id).First(&userModel)
	return
}

// All 获取所有用户数据
func All() (users []User) {
	database.DB.Find(&users)
	return
}
