package routes

import (
	controllers "forum/app/http/controllers/api"
	"github.com/gin-gonic/gin"
)

// RegisterRoutes 注册路由
func RegisterRoutes(route *gin.Engine) {
	api := route.Group("/api")
	{
		index := new(controllers.IndexController)
		api.GET("/index", index.Index)
	}
}
