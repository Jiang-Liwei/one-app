package auth

import (
	"forum/app/http/controllers/api"
	"forum/app/requests"
	"forum/pkg/captcha"
	"forum/pkg/logger"
	"forum/pkg/response"
	"forum/pkg/verifycode"
	"github.com/gin-gonic/gin"
)

// VerifyCodeController 控制器
type VerifyCodeController struct {
	api.Controller
}

// ShowCaptcha 显示图片验证码
func (vc *VerifyCodeController) ShowCaptcha(c *gin.Context) {

	// 生成验证码
	id, base64, err := captcha.NewCaptcha().GenerateCaptcha()
	// 记录错误日志 (error等级)
	logger.LogRecord(err, "err")

	// 返回给用户
	response.JSON(c, gin.H{
		"id":    id,
		"image": base64,
	})
}

// SendUsingPhone 发送手机验证码
func (vc *VerifyCodeController) SendUsingPhone(c *gin.Context) {

	// 表单验证
	request := requests.VerifyCodePhoneRequest{}
	if ok := requests.Validate(c, &request, requests.VerifyCodePhone); !ok {
		return
	}

	// 发送短信
	if ok := verifycode.NewVerifyCode().SendSMS([]string{request.Phone}); !ok {
		response.Abort500(c, "短信发送失败")
	}
	response.Success(c)
}
