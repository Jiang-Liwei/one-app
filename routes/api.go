package routes

import (
	controllers "forum/app/http/controllers/api"
	"forum/app/http/controllers/api/auth"
	"forum/app/http/controllers/api/category"
	"forum/app/http/controllers/api/sundry"
	"forum/app/http/controllers/api/topic"
	"forum/app/http/middlewares"
	"github.com/gin-gonic/gin"
)

// RegisterRoutes 注册路由
func RegisterRoutes(route *gin.Engine) {
	api := route.Group("api")
	{
		index := new(controllers.IndexController)
		api.GET("index", index.Index)

		// 注册模块
		authGroup := api.Group("auth")
		authGroup.Use(middlewares.LimitIP("1000-H"))
		{
			signupGroup := authGroup.Group("signup")
			{
				signup := new(auth.SignUpController)
				// 手机是否注册
				signupGroup.POST("/phone/exist", middlewares.GuestJWT(), signup.IsPhoneExist)
				// 判断 Email 是否已注册
				signupGroup.POST("/email/exist", middlewares.GuestJWT(), signup.IsEmailExist)
				// 手机号注册
				signupGroup.POST("/using-phone", middlewares.GuestJWT(), signup.SignupUsingPhone)
			}
			verifyCode := new(auth.VerifyCodeController)
			// 图片验证码
			authGroup.GET("/verify-codes/captcha", middlewares.LimitPerRoute("50-H"), verifyCode.ShowCaptcha)
			// 发送短信
			authGroup.POST("/verify-codes/sms", middlewares.LimitPerRoute("20-H"), verifyCode.SendUsingPhone)

			// 登录模块
			login := new(auth.LoginController)
			loginGroup := authGroup.Group("login")
			{
				// 验证码登录
				loginGroup.POST("using-phone", middlewares.GuestJWT(), login.LoginByPhone)
				// 账号登录
				loginGroup.POST("using-password", middlewares.GuestJWT(), login.LoginByPassword)
				// 刷新token
				loginGroup.POST("refresh-token", middlewares.AuthJWT(), login.RefreshToken)
			}

			// 密码操作模块
			password := new(auth.PasswordController)
			passwordGroup := authGroup.Group("password")
			{
				// 手机号重置密码
				passwordGroup.POST("reset/using-phone", middlewares.GuestJWT(), password.ResetByPhone)
			}
		}

		// 用户模块
		userGroup := api.Group("user")
		{
			uc := new(controllers.UsersController)
			// 个人详情
			userGroup.GET("info", middlewares.AuthJWT(), uc.CurrentUser)
			// 用户列表
			userGroup.GET("users", uc.Index)
			// 修改账户信息
			userGroup.PUT("self", middlewares.AuthJWT(), uc.UpdateProfile)
			// 更换邮箱账号
			userGroup.PUT("/email", middlewares.AuthJWT(), uc.UpdateEmail)
			// 更换手机账号
			userGroup.PUT("/phone", middlewares.AuthJWT(), uc.UpdatePhone)
			// 修改密码
			userGroup.PUT("/password", middlewares.AuthJWT(), uc.UpdatePassword)
		}

		// 分类模块
		categoryGroup := api.Group("categories")
		{
			cgc := new(category.CategoriesController)
			// 分类列表
			categoryGroup.GET("", cgc.Index)
			// 创建分类
			categoryGroup.POST("", middlewares.AuthJWT(), cgc.Store)
			// 更新分类
			categoryGroup.PUT("/:id", middlewares.AuthJWT(), cgc.Update)
			// 删除分类
			categoryGroup.DELETE("/:id", middlewares.AuthJWT(), cgc.Delete)
		}

		// 话题模块
		topicGroup := api.Group("topics")
		{
			tpc := new(topic.TopicsController)
			// 话题列表
			topicGroup.GET("", tpc.Index)
			// 创建话题
			topicGroup.POST("", middlewares.AuthJWT(), tpc.Store)
			// 修改话题
			topicGroup.PUT("/:id", middlewares.AuthJWT(), tpc.Update)
			// 删除话题
			topicGroup.DELETE("/:id", middlewares.AuthJWT(), tpc.Delete)
			// 话题详情
			topicGroup.GET("/:id", tpc.Show)
		}

		// 杂项模块
		sundryGroup := api.Group("sundry")
		{
			lsc := new(sundry.LinksController)
			sundryGroup.GET("link", lsc.Index)
		}
	}
}
