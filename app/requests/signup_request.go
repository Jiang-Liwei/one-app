package requests

import (
	"forum/app/requests/validators"
	"github.com/gin-gonic/gin"
	"github.com/thedevsaddam/govalidator"
)

type PhoneExistRequest struct {
	Phone string `json:"phone,omitempty" valid:"phone"`
}

type EmailExistRequest struct {
	Email string `json:"email,omitempty" valid:"email"`
}

// PhoneExist 验证手机号是否已经注册
func PhoneExist(data interface{}, c *gin.Context) map[string][]string {

	// 验证规则
	rules := govalidator.MapData{
		"phone": []string{"required", "digits:11", "mobile"},
	}

	// 自定义验证出错时的提示
	messages := govalidator.MapData{
		"phone": []string{
			"required:手机号为必填项，参数名称 phone",
			"digits:手机号长度必须为 11 位的数字",
			"mobile:错误的手机格式",
		},
	}

	return NewValidate(data, rules, messages)
}

// EmailExist 验证邮箱是否注册
func EmailExist(data interface{}, c *gin.Context) map[string][]string {

	// 自定义验证规则
	rules := govalidator.MapData{
		"email": []string{"required", "min:4", "max:30", "email"},
	}

	// 自定义验证出错时的提示
	messages := govalidator.MapData{
		"email": []string{
			"required:Email 为必填项",
			"min:Email 长度需大于 4",
			"max:Email 长度需小于 30",
			"email:Email 格式不正确，请提供有效的邮箱地址",
		},
	}
	return NewValidate(data, rules, messages)
}

// SignupUsingPhoneRequest 通过手机注册的请求信息
type SignupUsingPhoneRequest struct {
	Phone           string `json:"phone,omitempty"  valid:"phone"`
	Name            string `json:"name," valid:"name"`
	Password        string `json:"password,omitempty" valid:"password"`
	PasswordConfirm string `json:"password_confirm,omitempty" valid:"password_confirm"`
	VerifyCode      string `json:"verify_code,omitempty" valid:"verify_code"`
}

func SignupUsingPhone(data interface{}, c *gin.Context) map[string][]string {

	rules := govalidator.MapData{
		"phone":            []string{"mobile"},
		"name":             []string{"required", "alpha_space", "between:3,20"},
		"password":         []string{"required", "min:6"},
		"password_confirm": []string{"required"},
		"verify_code":      []string{"required", "digits:6"},
	}

	msg := govalidator.MapData{
		"phone": []string{
			"mobile:手机号格式错误",
		},
		"name": []string{
			"required:昵称不能为空",
			"alpha_space:非法昵称",
			"between:仅支持长度为3-20的昵称",
		},
		"password": []string{
			"required:请输入密码",
			"min:密码长度需大于等于 6",
		},
		"password_confirm": []string{
			"required:请输入确认密码",
		},
		"verify_code": []string{
			"required:请输入验证码",
			"digits:验证码长度必须为6",
		},
	}
	errs := NewValidate(data, rules, msg)
	_data := data.(*SignupUsingPhoneRequest)
	errs = validators.ValidatePasswordConfirm(_data.Password, _data.PasswordConfirm, errs)
	errs = validators.ValidateVerifyCode(_data.Phone, _data.VerifyCode, errs)

	return errs
}
