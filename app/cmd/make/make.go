package make

import (
	"embed"
	"fmt"
	"forum/pkg/console"
	"forum/pkg/file"
	"forum/pkg/helpers"
	"forum/pkg/str"
	"github.com/iancoleman/strcase"
	"github.com/spf13/cobra"
	"os"
	"strings"
)

type Model struct {
	// 单个单词时：小写且为复数；两个词或者以上：蛇形拼接
	TableName string
	// 单个单词时：首字母大写；两个词或者以上：大驼峰
	StructName string
	// 单个单词时：首字母大写且为复数；两个词或者以上：大驼峰且为复数
	StructNamePlural string
	// 单个单词时：小写；两个词或者以上：小驼峰
	VariableName string
	// 单个单词时：小写且为复数；两个词或者以上：小驼峰且为复数
	VariableNamePlural string
	// 单个单词时：小写；两个词或者以上：蛇形拼接
	PackageName string
}

// stubsFS 方便我们后面打包这些 .stub 为后缀名的文件
//
//go:embed stubs
var stubsFS embed.FS

// Make 说明 cobra 命令
var Make = &cobra.Command{
	Use:   "make",
	Short: "生成文件、代码",
}

func init() {
	// 注册 make 的子命令
	Make.AddCommand(
		MakeCMD,
		MakeModel,
		MakeController,
		MakeRequest,
		MakeMigration,
	)
}

// makeModelFromString 格式化用户输入的内容
func makeModelFromString(name string) Model {
	model := Model{}
	model.StructName = str.Singular(strcase.ToCamel(name))
	model.StructNamePlural = str.Plural(model.StructName)
	model.TableName = str.Snake(model.StructNamePlural)
	model.VariableName = str.LowerCamel(model.StructName)
	model.PackageName = str.Snake(model.StructName)
	model.VariableNamePlural = str.LowerCamel(model.StructNamePlural)
	return model
}

// createFileFromStub 读取 stub 文件并进行变量替换
// 最后一个选项可选，如若传参，应传 map[string]string 类型，作为附加的变量搜索替换
func createFileFromStub(filePath string, stubName string, model Model, variables ...interface{}) {

	// 实现最后一个参数可选
	replaces := make(map[string]string)
	if len(variables) > 0 {
		replaces = variables[0].(map[string]string)
	}

	// 目标文件已存在
	if file.Exists(filePath) {
		console.Exit(filePath + " already exists!")
	}

	// 读取 stub 模板文件
	modelData, err := stubsFS.ReadFile("stubs/" + stubName + ".stub")
	if err != nil {
		console.Exit(err.Error())
	}
	modelStub := string(modelData)

	// 添加默认的替换变量
	replaces["{{VariableName}}"] = model.VariableName
	replaces["{{VariableNamePlural}}"] = model.VariableNamePlural
	replaces["{{StructName}}"] = model.StructName
	replaces["{{StructNamePlural}}"] = model.StructNamePlural
	replaces["{{PackageName}}"] = model.PackageName
	replaces["{{TableName}}"] = model.TableName

	// 对模板内容做变量替换
	for search, replace := range replaces {
		modelStub = strings.ReplaceAll(modelStub, search, replace)
	}

	isHave, folder := helpers.FolderIsExist(filePath)
	if !isHave {
		err = os.MkdirAll(folder, 0775)
		if err != nil {
			console.Exit(err.Error())
		}
	}

	// 存储到目标文件中
	err = file.Put([]byte(modelStub), filePath)
	if err != nil {
		console.Exit(err.Error())
	}

	// 提示成功
	console.Success(fmt.Sprintf("[%s] created.", filePath))
}

// getArgPathAndModel 获取命令参数内的路径以及Model返回
func getArgPathAndModel(path string, filePath string, file string, modelType int) (string, Model) {

	array := strings.Split(path, "/")

	arrayLen := len(array)
	if !(arrayLen > 0) {
		console.Exit("controller名称不能为空")
	}
	// 最后一个下标
	maxKey := arrayLen - 1
	var modelValue string
	var model Model
	for k, v := range array {

		if k == maxKey {
			// name 用来生成 cmd.Model 实例
			model = makeModelFromString(v)
			switch modelType {
			case 1:
				modelValue = model.TableName
			case 2:
				modelValue = model.StructName
			case 3:
				modelValue = model.StructNamePlural
			case 4:
				modelValue = model.VariableName
			case 5:
				modelValue = model.VariableNamePlural
			default:
				modelValue = model.PackageName
			}

			filePath = filePath + "/" + modelValue + file
			continue
		}
		filePath = filePath + "/" + v
	}
	return filePath, model
}
