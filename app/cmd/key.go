package cmd

import (
	"forum/pkg/console"
	"forum/pkg/helpers"
	"github.com/spf13/cobra"
)

var Key = &cobra.Command{
	Use:   "key",
	Short: "生成应用程序密钥，打印生成的密钥",
	Run:   runKeyGenerate,
	Args:  cobra.NoArgs, // 不允许传参
}

func runKeyGenerate(cmd *cobra.Command, args []string) {
	console.Success("---")
	console.Success("App Key:")
	console.Success(helpers.RandomString(32))
	console.Success("---")
	console.Warning("please go to .env file to change the APP_KEY option")
}
