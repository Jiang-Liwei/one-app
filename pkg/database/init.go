package database

import (
	"errors"
	"fmt"
	"forum/pkg/config"
	"forum/pkg/logger"
	"time"

	"gorm.io/driver/mysql"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

// SetupDB 初始化数据库和 ORM
func SetupDB() {

	var dbConfig gorm.Dialector
	switch config.Get[string]("database.connection") {
	case "mysql":
		// 构建 DSN 信息
		dsn := fmt.Sprintf("%v:%v@tcp(%v:%v)/%v?charset=%v&parseTime=True&multiStatements=true&loc=Local",
			config.Get[string]("database.mysql.username"),
			config.Get[string]("database.mysql.password"),
			config.Get[string]("database.mysql.host"),
			config.Get[string]("database.mysql.port"),
			config.Get[string]("database.mysql.database"),
			config.Get[string]("database.mysql.charset"),
		)
		dbConfig = mysql.New(mysql.Config{
			DSN: dsn,
		})
	case "sqlite":
		// 初始化 sqlite
		database := config.Get[string]("database.sqlite.database")
		dbConfig = sqlite.Open(database)
	default:
		panic(errors.New("database connection not supported"))
	}

	// 连接数据库，并设置 GORM 的日志模式
	Connect(dbConfig, logger.NewGormLogger())

	// 设置最大连接数
	SQLDB.SetMaxOpenConns(config.Get[int]("database.mysql.max_open_connections"))
	// 设置最大空闲连接数
	SQLDB.SetMaxIdleConns(config.Get[int]("database.mysql.max_idle_connections"))
	// 设置每个链接的过期时间
	SQLDB.SetConnMaxLifetime(time.Duration(config.Get[int]("database.mysql.max_life_seconds")) * time.Second)
}
