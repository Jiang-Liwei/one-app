package captcha

import (
	"forum/pkg/config"
	"forum/pkg/redis"
	"github.com/mojocn/base64Captcha"
	"sync"
)

type Captcha struct {
	Base64Captcha *base64Captcha.Captcha
}

// 单例
var once sync.Once

// internalCaptcha 内部使用的 Captcha 对象
var internalCaptcha *Captcha

// NewCaptcha 单例获取
func NewCaptcha() *Captcha {
	once.Do(func() {

		// 初始化 Captcha对象
		internalCaptcha = &Captcha{}

		// 使用全局 Redis 对象，并储存 key 的前缀
		store := RedisStore{
			Client:    redis.Redis,
			KeyPrefix: config.Get[string]("app.name") + ":captcha",
		}

		// 配置 base64Captcha 驱动信息
		driver := base64Captcha.NewDriverDigit(
			config.Get[int]("captcha.height"),      // 宽
			config.Get[int]("captcha.width"),       // 高
			config.Get[int]("captcha.length"),      // 长度
			config.Get[float64]("captcha.maxskew"), // 数字的最大倾斜角度
			config.Get[int]("captcha.dotcount"),    // 图片背景里的混淆点数量
		)

		// 实例化 base64Captcha 并赋值给内部使用的 internalCaptcha 对象
		internalCaptcha.Base64Captcha = base64Captcha.NewCaptcha(driver, &store)
	})

	return internalCaptcha
}

// GenerateCaptcha 生成图片验证码
func (c Captcha) GenerateCaptcha() (id, base64 string, err error) {
	return c.Base64Captcha.Generate()
}

// VerifyCaptcha 验证验证码是否正确
func (c Captcha) VerifyCaptcha(id, answer string) (match bool) {

	// 第三个参数是验证后是否删除
	return c.Base64Captcha.Verify(id, answer, false)
}
