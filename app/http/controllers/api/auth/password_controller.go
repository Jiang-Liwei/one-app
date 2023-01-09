package auth

import (
	"forum/app/http/controllers/api"
	"forum/app/models/user"
	"forum/app/requests"
	authRequest "forum/app/requests/auth"
	"forum/pkg/response"
	"github.com/gin-gonic/gin"
)

// PasswordController 用户控制器
type PasswordController struct {
	api.Controller
}

// ResetByPhone 手机号找回密码
func (controller *PasswordController) ResetByPhone(c *gin.Context) {

	// 表单验证
	request := authRequest.ResetByPhoneRequest{}
	if ok := requests.Validate(c, &request, authRequest.ResetByPhone); !ok {
		return
	}

	// 更新密码
	userModel := user.GetByPhone(request.Phone)
	if userModel.ID == 0 {
		response.Abort404(c, "未知用户")
		return
	}
	userModel.Password = request.Password
	userModel.Save()

	response.Success(c)
}
