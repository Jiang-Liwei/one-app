package api

import (
	"forum/app/models/user"
	"forum/app/requests"
	userRequest "forum/app/requests/user"
	"forum/pkg/auth"
	"forum/pkg/response"
	"time"

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

// Index 所有用户
func (ctrl *UsersController) Index(c *gin.Context) {

	request := requests.PaginationRequest{}
	if ok := requests.Validate(c, &request, requests.Pagination); !ok {
		return
	}
	data, pager := user.Paginate(c, 10)
	response.JSON(c, gin.H{
		"data":  data,
		"pager": pager,
	})
}

func (ctrl *UsersController) UpdateProfile(c *gin.Context) {

	request := userRequest.UserUpdateProfileRequest{}
	if ok := requests.Validate(c, &request, userRequest.UserUpdateProfile); !ok {
		return
	}

	currentUser := auth.User(c)
	currentUser.Name = request.Name
	currentUser.City = request.City
	currentUser.Introduction = request.Introduction

	// 如果注册时间为空后面补一个
	if currentUser.CreatedAt.IsZero() {
		currentUser.CreatedAt = time.Now()
	}

	rowsAffected := currentUser.Save()
	if rowsAffected > 0 {
		response.Data(c, currentUser)
	} else {
		response.Abort500(c, "更新失败，请稍后尝试~")
	}
}

func (ctrl *UsersController) UpdateEmail(c *gin.Context) {

	request := userRequest.UserUpdateEmailRequest{}
	if ok := requests.Validate(c, &request, userRequest.UserUpdateEmail); !ok {
		return
	}

	currentUser := auth.User(c)
	currentUser.Email = request.Email
	rowsAffected := currentUser.Save()

	if rowsAffected > 0 {
		response.Success(c)
	} else {
		// 失败，显示错误提示
		response.Abort500(c, "更新失败，请稍后尝试~")
	}
}
