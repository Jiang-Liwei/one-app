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
				// 手机是否注册
				signupGroup.POST("/phone/exist", signup.IsPhoneExist)
				// 判断 Email 是否已注册
				signupGroup.POST("/email/exist", signup.IsEmailExist)
				// 手机号注册
				signupGroup.POST("/using-phone", signup.SignupUsingPhone)
			}
			verifyCode := new(auth.VerifyCodeController)
			// 图片验证码
			authGroup.GET("/verify-codes/captcha", verifyCode.ShowCaptcha)
			// 发送短信
			authGroup.POST("/verify-codes/sms", verifyCode.SendUsingPhone)

			// 登录模块
			login := new(auth.LoginController)
			loginGroup := authGroup.Group("login")
			{
				// 验证码登录
				loginGroup.POST("using-phone", login.LoginByPhone)
				// 账号登录
				loginGroup.POST("using-password", login.LoginByPassword)
				// 刷新token
				loginGroup.POST("refresh-token", login.RefreshToken)
			}

			//密码操作模块
			password := new(auth.PasswordController)
			passwordGroup := authGroup.Group("password")
			{
				// 手机号重置密码
				passwordGroup.POST("reset/using-phone", password.ResetByPhone)
			}
		}
	}
}
