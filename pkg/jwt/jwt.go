package jwt

import (
	"errors"
	"forum/pkg/app"
	"forum/pkg/config"
	"forum/pkg/logger"
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v4"
	"strings"
	"time"
)

var (
	ErrTokenExpired           error = errors.New("令牌已过期")
	ErrTokenExpiredMaxRefresh error = errors.New("令牌已过最大刷新时间")
	ErrTokenMalformed         error = errors.New("请求令牌格式有误")
	ErrTokenInvalid           error = errors.New("请求令牌无效")
	ErrHeaderEmpty            error = errors.New("需要认证才能访问！")
	ErrHeaderMalformed        error = errors.New("请求头中 Authorization 格式有误")
)

type JWT struct {

	// 秘钥
	SignKey []byte

	// Token 刷新的最大过期时间
	MaxRefresh time.Duration
}

// CustomClaims 自定义载荷
type CustomClaims struct {
	UserID       string `json:"user_id"`
	UserName     string `json:"user_name"`
	ExpireAtTime int64  `json:"expire_time"`

	// RegisteredClaims 结构体实现了 Claims 接口继承了  Valid() 方法
	// JWT 规定了7个官方字段，提供使用:
	// - iss (issuer)：发布者
	// - sub (subject)：主题
	// - iat (Issued At)：生成签名的时间
	// - exp (expiration time)：签名过期时间
	// - aud (audience)：观众，相当于接受者
	// - nbf (Not Before)：生效时间
	// - jti (JWT ID)：编号
	jwt.RegisteredClaims
}

func NewJWT() *JWT {
	return &JWT{
		SignKey:    []byte(config.Get[string]("app.key")),
		MaxRefresh: time.Duration(config.Get[int64]("jwt.max_refresh_time")) * time.Minute,
	}
}

// ParserToken 解析 Token，中间件中调用
func (j *JWT) ParserToken(c *gin.Context) (*CustomClaims, error) {

	tokenString, parseErr := j.getTokenFromHeader(c)
	if parseErr != nil {
		return nil, parseErr
	}

	// 解析token
	token, err := j.parseTokenString(tokenString)

	// 解析出错
	if err != nil {
		validationErr, ok := err.(*jwt.ValidationError)
		// 满足 refresh 的条件：只是单一的报错 ValidationErrorExpired
		if !ok {
			if validationErr.Errors == jwt.ValidationErrorMalformed {
				return nil, ErrTokenMalformed
			}
			if validationErr.Errors == jwt.ValidationErrorExpired {
				return nil, ErrTokenExpired
			}
		}
		return nil, ErrTokenInvalid
	}
	// 将 token 中的 claims 信息解析出来和 JWTCustomClaims 数据结构进行校验
	if claims, ok := token.Claims.(*CustomClaims); ok && token.Valid {
		return claims, nil
	}

	return nil, ErrTokenInvalid
}

// RefreshToken 更新 Token
func (j *JWT) RefreshToken(c *gin.Context) (string, error) {

	// 获取token
	stringToken, parseErr := j.getTokenFromHeader(c)
	if parseErr != nil {
		return "", parseErr
	}

	// 解析 Token
	token, err := j.parseTokenString(stringToken)
	if err != nil {
		validationErr, ok := err.(*jwt.ValidationError)
		// 满足 refresh 的条件：只是单一的报错 ValidationErrorExpired
		if !ok || validationErr.Errors != jwt.ValidationErrorExpired {
			return "", err
		}
	}

	// 解析 CustomClaims 的数据
	claims := token.Claims.(*CustomClaims)

	// 检查是否过了『最大允许刷新的时间』
	unix := app.TimeNowInTimezone().Add(-j.MaxRefresh).Unix()
	if claims.IssuedAt.Time.Unix() > unix {
		// 修改过期时间
		claims.RegisteredClaims.ExpiresAt.Time = j.expireAtTime()
		return j.createToken(*claims)
	}

	return "", ErrTokenExpiredMaxRefresh
}

// IssueToken 生成Token
func (j *JWT) IssueToken(userID string, userName string) string {

	// 构建用户信息
	expireAtTime := j.expireAtTime()
	numericDate := jwt.NumericDate{
		Time: app.TimeNowInTimezone(),
	}

	claims := CustomClaims{
		userID,
		userName,
		expireAtTime.Unix(),
		jwt.RegisteredClaims{
			// 签名生效时间
			NotBefore: &numericDate,
			// 首次签名时间
			IssuedAt: &numericDate,
			// 签名过期时间
			ExpiresAt: &jwt.NumericDate{Time: expireAtTime},
			// 签名颁发者
			Issuer: config.Get[string]("app.name"),
		},
	}

	// 生成token对象
	token, err := j.createToken(claims)
	if err != nil {
		logger.LogRecord(err)
		return ""
	}
	return token
}

// getTokenFromHeader 使用 jwt.ParseWithClaims 解析 Token
// Authorization:Bearer xxx
func (j *JWT) getTokenFromHeader(c *gin.Context) (string, error) {
	authHeader := c.Request.Header.Get("Authorization")
	if authHeader == "" {
		return "", ErrHeaderEmpty
	}
	// 按空格分割
	parts := strings.SplitN(authHeader, " ", 2)
	if !(len(parts) == 2 && parts[0] == "Bearer") {
		return "", ErrHeaderMalformed
	}
	return parts[1], nil
}

// parseTokenString 使用 jwt.ParseWithClaims 解析 Token
func (j *JWT) parseTokenString(tokenString string) (*jwt.Token, error) {
	return jwt.ParseWithClaims(tokenString, &CustomClaims{}, func(token *jwt.Token) (interface{}, error) {
		return j.SignKey, nil
	})
}

// expireAtTime 过期时间
func (j *JWT) expireAtTime() time.Time {
	timeNow := app.TimeNowInTimezone()

	var expireTime int64
	if config.Get[bool]("app.debug") {
		expireTime = config.Get[int64]("jwt.debug_expire_time")
	} else {
		expireTime = config.Get[int64]("jwt.expire_time")
	}
	expire := time.Duration(expireTime) * time.Minute

	return timeNow.Add(expire)
}

// createToken 创建token
func (j *JWT) createToken(claims CustomClaims) (string, error) {
	// 使用HS256算法进行token生成
	token := jwt.NewWithClaims(jwt.SigningMethodES256, claims)

	return token.SigningString()
}
