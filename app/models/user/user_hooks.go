package user

import (
	"forum/pkg/hash"
	"gorm.io/gorm"
)

// BeforeSave 将密码加密 Gorm钩子
func (userModel *User) BeforeSave(db *gorm.DB) (err error) {

	if !hash.BcryptIsHashed(userModel.Password) {
		userModel.Password = hash.BcryptHash(userModel.Password)
	}
	return
}
