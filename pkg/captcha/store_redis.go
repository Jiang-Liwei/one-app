package captcha

import (
	"errors"
	"forum/pkg/app"
	"forum/pkg/config"
	"forum/pkg/redis"
	"time"
)

// RedisStore 实现 base64Captcha.Store interface
type RedisStore struct {
	Client    *redis.Client
	KeyPrefix string
}

// Set 实现 base64Captcha.Store interface 的 Set 方法
func (s RedisStore) Set(key string, value string) error {

	ExpireTime := time.Minute * time.Duration(config.Get[int64]("captcha.expire_time"))
	// 方便本地开发调试
	if app.IsLocal() {
		ExpireTime = time.Minute * time.Duration(config.Get[int64]("captcha.debug_expire_time"))
	}

	if ok := s.Client.Set(s.KeyPrefix+key, value, ExpireTime); !ok {
		return errors.New("无法存储图片验证码答案")
	}
	return nil
}

// Get 实现 base64Captcha.Store interface 的 Get 方法
func (s RedisStore) Get(key string, clear bool) string {
	key = s.KeyPrefix + key
	val := s.Client.Get(key)
	if clear {
		s.Client.Del(key)
	}

	return val
}

// Verify 实现 base64Captcha.Store interface 的 Verify 方法
func (s RedisStore) Verify(key, answer string, clear bool) bool {
	val := s.Get(key, clear)
	return val == answer
}
