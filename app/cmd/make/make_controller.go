package make

import (
	"forum/pkg/console"
	"github.com/spf13/cobra"
	"strings"
)

var MakeController = &cobra.Command{
	Use:   "controller",
	Short: "Create api controller，exmaple: make controller user",
	Run:   runMakeAPIController,
	Args:  cobra.ExactArgs(1), // 只允许且必须传 1 个参数
}

func runMakeAPIController(cmd *cobra.Command, args []string) {

	// 处理参数，要求附带 API 版本（v1 或者 v2）
	array := strings.Split(args[0], "/")

	arrayLen := len(array)
	if !(arrayLen > 0) {
		console.Exit("controller名称不能为空")
	}
	// 最后一个下标
	maxKey := arrayLen - 1

	var filePath = "app/http/controllers/api"
	var model Model
	for k, v := range array {
		if k == maxKey {
			// name 用来生成 cmd.Model 实例
			model = makeModelFromString(v)
			filePath = filePath + "/" + model.TableName + "_controller.go"
			continue
		}
		filePath = filePath + "/" + v
	}

	// 基于模板创建文件（做好变量替换）
	createFileFromStub(filePath, "controller", model)
}
