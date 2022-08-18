package routes

import (
	"github.com/gin-gonic/gin"
	controllers "one-app/app/http/controllers/api"
)

// RegisterRoutes 注册路由
func RegisterRoutes(route *gin.Engine) {
	api := route.Group("/api")
	{
		index := new(controllers.IndexController)
		api.GET("/index", index.Index)
	}
}
