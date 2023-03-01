package sundry

import (
	"forum/app/http/controllers/api"
	"forum/app/models/link"
	"forum/pkg/response"
	"github.com/gin-gonic/gin"
)

type LinksController struct {
	api.Controller
}

func (ctrl *LinksController) Index(c *gin.Context) {
	links := link.All()
	response.Data(c, links)
}
