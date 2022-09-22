package bootstrap

import (
	"flag"
	"forum/app/models/user"
	btsConfig "forum/config"
	"forum/pkg/config"
	"forum/pkg/database"
	"forum/pkg/logger"
	"forum/pkg/requests"
	"forum/routes"
	"github.com/gin-gonic/gin"
)

func init() {
	// 加载 config 目录下的配置信息
	btsConfig.Initialize()
}

func Start() {
	var env string
	flag.StringVar(&env, "env", "", "加载 .env 文件，如 --env=testing 加载的是 .env.testing 文件")
	flag.Parse()
	config.InitConfig(env)

	// 初始化 Logger
	logger.SetupLogger()

	r := gin.Default()

	database.SetupDB()
	requests.InitRequests()
	err := database.DB.AutoMigrate(&user.User{})
	if err != nil {
		return
	}
	routes.SetupRoute(r)

	err = r.Run(":" + config.Get[string]("app.port"))
	if err != nil {
		return
	} // 监听并在 0.0.0.0:8080 上启动服务
}
