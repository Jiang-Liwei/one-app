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
			signupGroup := authGroup.Group("signup")
			{
				signup := new(auth.SignUpController)
				signupGroup.POST("/phone/exist", signup.IsPhoneExist)
				// 判断 Email 是否已注册
				signupGroup.POST("/email/exist", signup.IsEmailExist)
			}
		}
	}
}
