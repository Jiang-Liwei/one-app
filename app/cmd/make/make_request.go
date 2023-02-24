package make

import (
	"github.com/spf13/cobra"
)

var MakeRequest = &cobra.Command{
	Use:   "request",
	Short: "Create request file, example make request user",
	Run:   runMakeRequest,
	Args:  cobra.ExactArgs(1), // 只允许且必须传 1 个参数
}

func runMakeRequest(cmd *cobra.Command, args []string) {

	filePath, model := getArgPathAndModel(args[0], "app/requests", "_request.go", 0)

	println("\n" + filePath + "\n")
	// 基于模板创建文件（做好变量替换）
	createFileFromStub(filePath, "request", model)
}
