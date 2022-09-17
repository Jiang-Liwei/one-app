package requests

import (
	"fmt"
	"github.com/thedevsaddam/govalidator"
	"regexp"
)

func init() {
	mobile()
}

func mobile() {
	// 添加手机正则验证
	govalidator.AddCustomRule("mobile", func(field string, rule string, message string, value interface{}) error {
		val := value.(string)
		result, _ := regexp.MatchString(`^1((34[0-8]\d{7})|((3[0-3|5-9])|(4[5-7|9])|(5[0-3|5-9])|(66)|(7[2-3|5-8])|(8[0-9])|(9[1|8-9]))\d{8})$`, val)
		if !result {
			return fmt.Errorf(message)
		}
		return nil
	})
}
