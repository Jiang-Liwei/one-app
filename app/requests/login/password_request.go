package login

import (
	"forum/app/requests"
	"forum/app/requests/validators"
	"github.com/gin-gonic/gin"
	"github.com/thedevsaddam/govalidator"
)

type ByPasswordRequest struct {
	CaptchaID     string `json:"captcha_id,omitempty" valid:"captcha_id"`
	CaptchaAnswer string `json:"captcha_answer,omitempty" valid:"captcha_answer"`

	AccountNumber string `json:"account_number" valid:"account_number"`
	Password      string `json:"password,omitempty" valid:"password"`
}

func ByPassword(data interface{}, c *gin.Context) map[string][]string {

	rules := govalidator.MapData{
		"account_number": []string{"required", "min:3"},
		"password":       []string{"required", "min:6"},
		"captcha_id":     []string{"required"},
		"captcha_answer": []string{"required", "digits:6"},
	}

	msg := govalidator.MapData{
		"account_number": []string{
			"required:账号不能为空",
			"min:账号长度必须大于3",
		},
		"password": []string{
			"required:密码不能为空",
			"min:密码长度必须大于6",
		},
		"captcha_id": []string{
			"required:验证码错误",
		},
		"captcha_answer": []string{
			"required:验证码不能为空",
			"digits:验证码错误",
		},
	}
	errs := requests.NewValidate(data, rules, msg)
	// 图片验证码
	_data := data.(*ByPasswordRequest)

	errs = validators.ValidateCaptcha(_data.CaptchaID, _data.CaptchaAnswer, errs)

	return errs
}
