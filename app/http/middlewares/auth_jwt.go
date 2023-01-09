package middlewares

import (
	"forum/app/models/user"
	"forum/pkg/jwt"
	"forum/pkg/response"
	"github.com/gin-gonic/gin"
)

func AuthJWT() gin.HandlerFunc {

	return func(c *gin.Context) {

		// 获取token
		claims, err := jwt.NewJWT().ParserToken(c)

		// JWT 解析失败
		if err != nil {
			response.Unauthorized(c)
			return
		}

		// JWT 解析成功，设置用户信息
		userModel := user.Get(claims.UserID)
		if userModel.ID == 0 {
			response.Unauthorized(c, "无效用户")
			return
		}

		// 将用户信息存入 gin.Context 中
		c.Set("user_id", userModel.GetStringID())
		c.Set("user", userModel)

		c.Next()
	}
}
