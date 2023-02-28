package category

import (
	"forum/app/http/controllers/api"
	"forum/app/models/category"
	"forum/app/requests"
	requestsCategory "forum/app/requests/category"
	"forum/pkg/response"

	"github.com/gin-gonic/gin"
)

type CategoriesController struct {
	api.Controller
}

func (ctrl *CategoriesController) Store(c *gin.Context) {

	request := requestsCategory.CategoryRequest{}
	if ok := requests.Validate(c, &request, requestsCategory.CategorySave); !ok {
		return
	}

	categoryModel := category.Category{
		Name:        request.Name,
		Description: request.Description,
	}
	categoryModel.Create()
	if categoryModel.ID > 0 {
		response.Created(c, categoryModel)
	} else {
		response.Abort500(c, "创建失败，请稍后尝试~")
	}
}

func (ctrl *CategoriesController) Update(c *gin.Context) {

	// 验证 url 参数 id 是否正确
	categoryModel := category.Get(c.Param("id"))
	if categoryModel.ID == 0 {
		response.Abort404(c)
		return
	}

	// 表单验证
	request := requestsCategory.CategoryRequest{}
	if ok := requests.Validate(c, &request, requestsCategory.CategorySave); !ok {
		return
	}

	// 保存数据
	categoryModel.Name = request.Name
	categoryModel.Description = request.Description
	rowsAffected := categoryModel.Save()

	if rowsAffected > 0 {
		response.Data(c, categoryModel)
	} else {
		response.Abort500(c)
	}
}
