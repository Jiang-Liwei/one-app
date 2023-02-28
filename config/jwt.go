package config

import "forum/pkg/config"

func init() {
	config.Add("jwt", func() map[string]interface{} {
		return map[string]interface{}{

			// 过期时间 单位分钟
			"expire_time": config.Env("JWT_EXPIRE_TIME", 120000),
			// 允许刷新时间 单位分钟
			"max_refresh_time": config.Env("MAX_REFRESH_TIME", 86400),
			//
		}
	})
}
