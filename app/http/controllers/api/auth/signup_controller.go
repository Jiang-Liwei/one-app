package auth

import (
	"forum/app/http/controllers/api"
	"forum/app/models/user"
	"forum/app/requests"
	"github.com/gin-gonic/gin"
	"net/http"
)

type SignUpController struct {
	api.Controller
}

// IsPhoneExist 检测手机是否被注册
func (signup *SignUpController) IsPhoneExist(c *gin.Context) {

	// 初始化请求对象
	request := requests.PhoneExistRequest{}
	// 解析 JSON 请求
	if ok := requests.Validate(c, &request, requests.PhoneExist); !ok {
		return
	}
	//  检查数据库并返回响应
	c.JSON(http.StatusOK, gin.H{
		"exist": user.IsPhoneExist(request.Phone),
	})
}

// IsEmailExist 检测邮箱是否已注册
func (signup *SignUpController) IsEmailExist(c *gin.Context) {

	// 初始化请求对象
	request := requests.EmailExistRequest{}

	if ok := requests.Validate(c, &request, requests.EmailExist); !ok {
		return
	}
	//  检查数据库并返回响应
	c.JSON(http.StatusOK, gin.H{
		"exist": user.IsEmailExist(request.Email),
	})
}
