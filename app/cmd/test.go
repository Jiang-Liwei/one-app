package cmd

import (
	"github.com/spf13/cobra"
)

var Test = &cobra.Command{
	Use:   "test",
	Short: "测试命令修改 runPlay 函数自定义自己需要测试的内容",
	Run:   runPlay,
}

// runPlay 自定义自己需要用来测试的内容
func runPlay(cmd *cobra.Command, args []string) {

	// 不执行命令不会影响
}
