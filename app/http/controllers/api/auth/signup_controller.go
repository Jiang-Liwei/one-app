package auth

import (
	"forum/app/http/controllers/api"
	"forum/app/models/user"
	"forum/app/requests"
	"forum/pkg/jwt"
	"forum/pkg/response"
	"github.com/gin-gonic/gin"
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
	response.JSON(c, gin.H{
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
	response.JSON(c, gin.H{
		"exist": user.IsEmailExist(request.Email),
	})
}

// SignupUsingPhone 短信注册账号
func (signup *SignUpController) SignupUsingPhone(c *gin.Context) {

	// 表单验证
	request := requests.SignupUsingPhoneRequest{}
	if ok := requests.Validate(c, &request, requests.SignupUsingPhone); !ok {
		return
	}

	// 注册
	_user := user.User{
		Name:     request.Name,
		Phone:    request.Phone,
		Password: request.Password,
	}

	_user.Create()

	if _user.ID > 0 {
		token := jwt.NewJWT().IssueToken(_user.GetStringID(), _user.Name)
		response.CreatedJSON(c, gin.H{
			"token": token,
			"data":  _user,
		})
		return
	}

	response.Abort500(c, "创建用户失败，请稍后尝试~")
}
