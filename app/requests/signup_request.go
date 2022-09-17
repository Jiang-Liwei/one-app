package requests

import (
	"github.com/gin-gonic/gin"
	"github.com/thedevsaddam/govalidator"
	"net/url"
)

type PhoneExistRequest struct {
	Phone string `json:"phone,omitempty" valid:"phone"`
}

func ValidatePhoneExist(data interface{}, c *gin.Context) url.Values {

	// 自定义规则
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

	// 配置初始化
	opts := govalidator.Options{
		Data:          data,
		Rules:         rules,
		TagIdentifier: "valid", // 模型中的 Struct 标签标识符
		Messages:      messages,
	}

	// 开始验证
	return govalidator.New(opts).ValidateStruct()
}