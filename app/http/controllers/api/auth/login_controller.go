package auth

import (
	"forum/app/http/controllers/api"
	"forum/app/requests"
	"forum/app/requests/login"
	"forum/pkg/auth"
	"forum/pkg/jwt"
	"forum/pkg/response"
	"github.com/gin-gonic/gin"
)

type LoginController struct {
	api.Controller
}

func (l *LoginController) LoginByPhone(c *gin.Context) {

	// 表单验证
	request := login.ByPhoneRequest{}

	if ok := requests.Validate(c, &request, login.ByPhone); !ok {
		return
	}

	// 登录
	user, err := auth.LoginByPhone(request.Phone)
	if err != nil {
		// 失败，显示错误提示
		response.Error(c, err, "账号不存在")
		return
	}

	token := jwt.NewJWT().IssueToken(user.GetStringID(), user.Name)

	response.JSON(c, gin.H{
		"token": token,
	})
}
