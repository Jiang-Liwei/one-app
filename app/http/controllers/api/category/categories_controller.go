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

func (ctrl *CategoriesController) Create(c *gin.Context) {

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
