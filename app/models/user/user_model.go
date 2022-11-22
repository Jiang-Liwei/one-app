package user

import (
	"forum/app/models"
	"forum/pkg/database"
)

type User struct {
	models.BaseModel

	Name     string `json:"name,omitempty"`
	Email    string `json:"-"`
	Phone    string `json:"-"`
	Password string `json:"-"`
}

// Create 创建用户
func (userModel *User) Create() {
	database.DB.Create(&userModel)
}
