package routes

import (
	controllers "forum/app/http/controllers/api"
	"forum/app/http/controllers/api/auth"
	"github.com/gin-gonic/gin"
)

// RegisterRoutes 注册路由
func RegisterRoutes(route *gin.Engine) {
	api := route.Group("/api")
	{
		index := new(controllers.IndexController)
		api.GET("/index", index.Index)
		authGroup := api.Group("/auth")
		{
			signup := new(auth.SignUpController)
			authGroup.POST("/phone/exist", signup.IsPhoneExist)
		}
	}
}
