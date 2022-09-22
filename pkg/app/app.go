package app

import "forum/pkg/config"

func IsLocal() bool {
	return config.Get[string]("app.env") == "local"
}

func IsProduction() bool {
	return config.Get[string]("app.env") == "production"
}

func IsTesting() bool {
	return config.Get[string]("app.env") == "testing"
}
