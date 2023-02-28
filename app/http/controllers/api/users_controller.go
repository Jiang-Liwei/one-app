package api

import (
	"forum/pkg/auth"
	"forum/pkg/response"

	"github.com/gin-gonic/gin"
)

type UsersController struct {
	Controller
}

// CurrentUser 当前登录用户信息
func (ctrl *UsersController) CurrentUser(c *gin.Context) {
	userModel := auth.User(c)
	response.Data(c, userModel)
}
