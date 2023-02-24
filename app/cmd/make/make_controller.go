package make

import (
	"github.com/spf13/cobra"
)

var MakeController = &cobra.Command{
	Use:   "controller",
	Short: "Create api controller，exmaple: make controller user",
	Run:   runMakeAPIController,
	Args:  cobra.ExactArgs(1), // 只允许且必须传 1 个参数
}

func runMakeAPIController(cmd *cobra.Command, args []string) {

	filePath, model := getArgPathAndModel(args[0], "app/http/controllers/api", "_controller.go", 1)

	// 基于模板创建文件（做好变量替换）
	createFileFromStub(filePath, "controller", model)
}
