package bootstrap

import (
	"fmt"
	"forum/app/cmd"
	"forum/app/cmd/make"
	btsConfig "forum/config"
	"forum/pkg/cache"
	"forum/pkg/config"
	"forum/pkg/console"
	"forum/pkg/database"
	"forum/pkg/logger"
	"forum/pkg/redis"
	"forum/pkg/requests"
	"github.com/spf13/cobra"
	"os"
)

func init() {
	// 加载 config 目录下的配置信息
	btsConfig.Initialize()
}

func Start() {

	// 应用的主入口，默认调用 cmd.CmdServe 命令
	var rootCmd = &cobra.Command{
		Use:   "forum",
		Short: "A simple forum project",
		Long:  `Default will run "serve" command, you can use "-h" flag to see all subcommands`,

		// rootCmd 的所有子命令都会执行以下代码
		PersistentPreRun: func(command *cobra.Command, args []string) {

			// 配置初始化，依赖命令行 --env 参数
			config.InitConfig(cmd.Env)

			// 初始化 Logger
			logger.SetupLogger()

			// 初始化数据库
			database.SetupDB()

			// 初始化 Redis
			redis.SetupRedis()

			// 初始化缓存
			cache.SetupCache()

			// 这里进行自定义规则的添加
			requests.InitRequests()
		},
	}

	// 注册子命令
	rootCmd.AddCommand(
		cmd.Serve,
		cmd.Key,
		cmd.Test,
		make.Make,
		cmd.Migrate,
		cmd.Seed,
	)

	// 配置默认运行 Web 服务
	cmd.RegisterDefaultCmd(rootCmd, cmd.Serve)

	// 注册全局参数，--env
	cmd.RegisterGlobalFlags(rootCmd)

	// 执行主命令
	if err := rootCmd.Execute(); err != nil {
		console.Exit(fmt.Sprintf("Failed to run app with %v: %s", os.Args, err.Error()))
	}
}
