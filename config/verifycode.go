package config

import "forum/pkg/config"

func init() {
	config.Add("verifycode", func() map[string]interface{} {
		return map[string]interface{}{

			// 验证码的长度
			"code_length": config.Env("VERIFY_CODE_LENGTH", 6),

			// 过期时间，单位是分钟
			"expire_time": config.Env("VERIFY_CODE_EXPIRE", 15),
		}
	})
}
