package make

import (
	"github.com/spf13/cobra"
)

var MakeModel = &cobra.Command{
	Use:   "model",
	Short: "Crate model file, example: make model user",
	Run:   runMakeModel,
	Args:  cobra.ExactArgs(1), // 只允许且必须传 1 个参数
}

func runMakeModel(cmd *cobra.Command, args []string) {

	filePath, model := getArgPathAndModel(args[0], "app/models", "/", 0)
	println(filePath)
	// 替换变量
	createFileFromStub(filePath+model.PackageName+"_model.go", "model/model", model)
	createFileFromStub(filePath+model.PackageName+"_util.go", "model/model_util", model)
	createFileFromStub(filePath+model.PackageName+"_hooks.go", "model/model_hooks", model)
}
