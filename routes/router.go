package routes

import "github.com/gin-gonic/gin"

// SetupRoute 路由初始化
func SetupRoute(router *gin.Engine) {
	RegisterRoutes(router)
}
