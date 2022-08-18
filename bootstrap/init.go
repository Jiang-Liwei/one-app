package bootstrap

import (
	"github.com/gin-gonic/gin"
	"one-app/routes"
)

func init() {

}

func Start() {
	r := gin.Default()
	routes.SetupRoute(r)
	err := r.Run()
	if err != nil {
		return
	} // 监听并在 0.0.0.0:8080 上启动服务
}
