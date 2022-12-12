package login

import (
	"forum/app/requests"
	"forum/app/requests/validators"
	"github.com/gin-gonic/gin"
	"github.com/thedevsaddam/govalidator"
)

type ByPhoneRequest struct {
	Phone      string `json:"phone,omitempty" valid:"phone"`
	VerifyCode string `json:"verify_code,omitempty" valid:"verify_code"`
}

// ByPhone 验证表单
func ByPhone(data interface{}, c *gin.Context) map[string][]string {

	rules := govalidator.MapData{
		"phone":       []string{"required", "mobile"},
		"verify_code": []string{"required", "digits:6"},
	}

	msg := govalidator.MapData{
		"phone": []string{
			"required:手机号为必填项，参数名称 phone",
			"mobile:手机格式错误",
		},
		"verify_code": []string{
			"required:验证码答案必填",
			"digits:验证码长度必须为 6 位的数字",
		},
	}

	errs := requests.NewValidate(data, rules, msg)
	// 手机验证码
	_data := data.(*ByPhoneRequest)
	errs = validators.ValidateVerifyCode(_data.Phone, _data.VerifyCode, errs)

	return errs
}
