package api

import (
	"forum/app/models/user"
	"forum/app/requests"
	userRequest "forum/app/requests/user"
	"forum/pkg/auth"
	"forum/pkg/config"
	"forum/pkg/file"
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

	request := userRequest.UpdateProfileRequest{}
	if ok := requests.Validate(c, &request, userRequest.UpdateProfile); !ok {
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

	request := userRequest.UpdateEmailRequest{}
	if ok := requests.Validate(c, &request, userRequest.UpdateEmail); !ok {
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

func (ctrl *UsersController) UpdatePhone(c *gin.Context) {

	request := userRequest.UpdatePhoneRequest{}
	if ok := requests.Validate(c, &request, userRequest.UpdatePhone); !ok {
		return
	}

	currentUser := auth.User(c)
	currentUser.Phone = request.Phone
	rowsAffected := currentUser.Save()

	if rowsAffected > 0 {
		response.Success(c)
	} else {
		response.Abort500(c, "更新失败，请稍后尝试~")
	}
}

func (ctrl *UsersController) UpdatePassword(c *gin.Context) {

	request := userRequest.UpdatePasswordRequest{}
	if ok := requests.Validate(c, &request, userRequest.UpdatePassword); !ok {
		return
	}

	currentUser := auth.User(c)
	// 验证原始密码是否正确
	isTrue := currentUser.ComparePassword(request.Password)
	if !isTrue {
		// 失败，显示错误提示
		response.Unauthorized(c, "原密码不正确")
		return
	}

	isTrue = currentUser.ComparePassword(request.NewPassword)
	if isTrue {
		response.Unauthorized(c, "新密码与旧密码一致")
		return
	}

	// 更新密码为新密码
	currentUser.Password = request.NewPassword
	currentUser.Save()

	response.Success(c)

}

func (ctrl *UsersController) UpdateAvatar(c *gin.Context) {

	request := userRequest.UpdateAvatarRequest{}
	if ok := requests.Validate(c, &request, userRequest.UpdateAvatar); !ok {
		return
	}

	avatar, err := file.SaveUploadAvatar(c, request.Avatar)
	if err != nil {
		response.Abort500(c, "上传头像失败，请稍后尝试~")
		return
	}

	currentUser := auth.User(c)
	currentUser.Avatar = config.Get[string]("app.url") + avatar
	currentUser.Save()

	response.Data(c, currentUser)
}
