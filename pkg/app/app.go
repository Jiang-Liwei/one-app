package app

import (
	"forum/pkg/config"
	"time"
)

func IsLocal() bool {
	return config.Get[string]("app.env") == "local"
}

func IsProduction() bool {
	return config.Get[string]("app.env") == "production"
}

func IsTesting() bool {
	return config.Get[string]("app.env") == "testing"
}

// TimeNowInTimezone 获取当前时间，支持时区
func TimeNowInTimezone() time.Time {
	chinaTimezone, _ := time.LoadLocation(config.Get[string]("app.timezone"))
	return time.Now().In(chinaTimezone)
}
