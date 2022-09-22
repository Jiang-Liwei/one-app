package logger

import "forum/pkg/config"

// SetupLogger 初始化 Logger
func SetupLogger() {
	logger := LoggerStruct{
		Filename:  config.Get[string]("log.filename"),
		MaxSize:   config.Get[int]("log.max_size"),
		MaxBackup: config.Get[int]("log.max_backup"),
		MaxAge:    config.Get[int]("log.max_age"),
		Compress:  config.Get[bool]("log.compress"),
		LogType:   config.Get[string]("log.type"),
		Level:     config.Get[string]("log.level"),
	}

	logger.InitLogger()
}
