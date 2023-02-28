package routes

import (
	controllers "forum/app/http/controllers/api"
	"forum/app/http/controllers/api/auth"
	"forum/app/http/controllers/api/category"
	"forum/app/http/middlewares"
	"github.com/gin-gonic/gin"
)

// RegisterRoutes 注册路由
func RegisterRoutes(route *gin.Engine) {
	api := route.Group("/api")
	{
		index := new(controllers.IndexController)
		api.GET("/index", index.Index)

		// 注册模块
		authGroup := api.Group("/auth")
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
		userGroup := api.Group("/user")
		userGroup.Use()
		{
			uc := new(controllers.UsersController)
			userGroup.GET("info", middlewares.AuthJWT(), uc.CurrentUser)
			userGroup.GET("users", uc.Index)
		}

		// 分类模块
		categoryGroup := api.Group("category")
		categoryGroup.Use()
		{
			cgc := new(category.CategoriesController)
			// 创建分类
			categoryGroup.POST("create", middlewares.AuthJWT(), cgc.Create)
		}
	}
}
