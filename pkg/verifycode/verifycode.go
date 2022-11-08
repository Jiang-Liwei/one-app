package verifycode

import (
	"forum/pkg/config"
	"forum/pkg/helpers"
	"forum/pkg/logger"
	"forum/pkg/redis"
	"forum/pkg/sms"
	"sync"
)

type VerifyCode struct {
	Store Store
}

var once sync.Once
var internalVerifyCode *VerifyCode

func NewVerifyCode() *VerifyCode {
	once.Do(func() {
		internalVerifyCode = &VerifyCode{
			Store: &RedisStore{
				RedisClient: redis.Redis,
				KeyPrefix:   config.Get[string]("app.name") + ":verify_code:",
			},
		}
	})

	return internalVerifyCode
}

// SendSMS 发送短信
func (vc *VerifyCode) SendSMS(phone []string) bool {
	code := vc.generateVerifyCode(phone)

	return sms.NewSMS().Send(phone, sms.Message{
		Template: config.Get[string]("sms.tencent.template_id"),
		Data:     map[string]string{"code": code, "expiration_time": config.Get[string]("verifycode.expire_time")},
	})
}

// CheckAnswer 验证提交的验证码是否正确
func (vc *VerifyCode) CheckAnswer(key, answer string) bool {
	logger.DebugJSON("验证码", "检查验证码", map[string]string{key: answer})

	return vc.Store.VerifyStore(key, answer, false)
}

// generateVerifyCode 生成验证码并存入 Redis
func (vc *VerifyCode) generateVerifyCode(keys []string) string {

	// 生成验证码
	code := helpers.RandomNumber(config.Get[int]("verifycode.code_length"))

	for _, phone := range keys {
		logger.DebugJSON("验证码", "生成验证码", map[string]string{phone: code})

		vc.Store.SetStore(phone, code)
	}

	return code
}
