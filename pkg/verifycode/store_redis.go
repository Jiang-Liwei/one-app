package verifycode

import (
	"forum/pkg/config"
	"forum/pkg/redis"
	"time"
)

// RedisStore 实现 Store interface
type RedisStore struct {
	RedisClient *redis.Client
	KeyPrefix   string
}

// SetStore 实现 verifycode.Store interface 的 SetStore 方法
func (s *RedisStore) SetStore(key, value string) bool {

	ExpireTime := time.Minute * time.Duration(config.Get[int64]("verify_code.expire_time"))

	return s.RedisClient.Set(s.KeyPrefix+key, value, ExpireTime)
}

// GetStore 实现 verifycode.Store interface 的 GetStore 方法
func (s *RedisStore) GetStore(key string, clear bool) string {
	key = s.KeyPrefix + key
	val := s.RedisClient.Get(key)
	if clear {
		s.RedisClient.Del(key)
	}

	return val
}

// VerifyStore 实现 verifycode.Store interface 的 VerifyStore 方法
func (s *RedisStore) VerifyStore(key, answer string, clear bool) bool {
	val := s.GetStore(key, clear)

	return val == answer
}
