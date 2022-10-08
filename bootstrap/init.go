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

	// 设置 gin 的运行模式，支持 debug, release, test
	// release 会屏蔽调试信息，官方建议生产环境中使用
	// 非 release 模式 gin 终端打印太多信息，干扰到我们程序中的 Log
	// 故此设置为 release，有特殊情况手动改为 debug 即可
	gin.SetMode(gin.ReleaseMode)

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
