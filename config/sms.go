package config

import "forum/pkg/config"

func init() {
	config.Add("sms", func() map[string]interface{} {
		return map[string]interface{}{

			"platform": config.Env("SMS_Platform"),

			// 腾讯云配置
			"tencent": map[string]interface{}{
				"secret_id":   config.Env("SMS_TENCENT_Secret_Id"),
				"secret_key":  config.Env("SMS_TENCENT_Secret_Key"),
				"app_id":      config.Env("SMS_TENCENT_App_Id"),
				"sign_name":   config.Env("SMS_TENCENT_Sign_Name"),
				"template_id": config.Env("SMS_TENCENT_Template_Id"),
			},

			// 阿里云配置
			"aliyun": map[string]interface{}{
				"access_key_id":     config.Env("SMS_ALIYUN_ACCESS_ID"),
				"access_key_secret": config.Env("SMS_ALIYUN_ACCESS_SECRET"),
				"sign_name":         config.Env("SMS_ALIYUN_SIGN_NAME"),
				"template_code":     config.Env("SMS_ALIYUN_TEMPLATE_CODE"),
			},
		}
	})
}
