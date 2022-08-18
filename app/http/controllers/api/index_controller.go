package api

import "github.com/gin-gonic/gin"

type IndexController struct {
}

func (index *IndexController) Index(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "你好！世界",
	})
}
