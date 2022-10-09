package redis

import (
	"context"
	"forum/pkg/logger"
	redis "github.com/go-redis/redis/v8"
	"sync"
	"time"
)

// Client Redis 服务
type Client struct {
	Client  *redis.Client
	Context context.Context
}

// once 保证只只实例化一次
var once sync.Once

// Redis 全局 Redis，使用 db 1
var Redis *Client

// ConnectRedis 连接redis数据库，设置全局的 Redis 对象
func ConnectRedis(address string, username string, password string, db int) {
	once.Do(func() {
		Redis = NewClient(address, username, password, db)
	})
}

// NewClient 创建redis连接
func NewClient(address string, username string, password string, db int) *Client {

	// 初始化 Client 实例
	c := &Client{}
	// 使用默认的 context
	c.Context = context.Background()

	// 使用 redis 库里的 NewClient 初始化连接
	c.Client = redis.NewClient(&redis.Options{
		Addr:     address,
		Username: username,
		Password: password,
		DB:       db,
	})

	err := c.Ping()
	logger.LogRecord(err)

	return c
}

// Ping 用以测试 redis 连接是否正常
func (c Client) Ping() error {
	_, err := c.Client.Ping(c.Context).Result()
	return err
}

// Set 存储 key 对应的 value，且设置 expiration 过期时间
func (c Client) Set(key string, value interface{}, expiration time.Duration) bool {
	if err := c.Client.Set(c.Context, key, value, expiration).Err(); err != nil {
		logger.ErrorString("Redis", "Set", err.Error())
		return false
	}
	return true
}

// Get 或者对应 key 的值
func (c Client) Get(key string) string {
	result, err := c.Client.Get(c.Context, key).Result()
	if err != nil {
		if err != redis.Nil {
			logger.ErrorString("Redis", "Get", err.Error())
		}

		return ""
	}
	return result
}

// Has 判断 key 是否存在
func (c Client) Has(key string) bool {
	_, err := c.Client.Get(c.Context, key).Result()
	if err != nil {
		if err != redis.Nil {
			logger.ErrorString("Redis", "Has", err.Error())
		}
		return false
	}
	return true
}

// Del 删除存储在 redis 里的数据，支持多个 key 传参
func (c Client) Del(keys ...string) bool {
	if err := c.Client.Del(c.Context, keys...).Err(); err != nil {
		logger.ErrorString("Redis", "Del", err.Error())
		return false
	}

	return true
}

// FlushDB 清空当前 redis db 里的所有数据
func (c Client) FlushDB() bool {
	if err := c.Client.FlushDB(c.Context).Err(); err != nil {
		logger.ErrorString("Redis", "FlushDB", err.Error())
		return false
	}
	return true
}

// Increment 当参数只有 1 个时，为 key，其值增加 1。
// 当参数有 2 个时，第一个参数为 key ，第二个参数为要增加的值 int64 类型。
func (c Client) Increment(parameters ...interface{}) bool {
	switch len(parameters) {
	case 1:
		key := parameters[0].(string)
		if err := c.Client.Incr(c.Context, key).Err(); err != nil {
			logger.ErrorString("Redis", "Increment", err.Error())
			return false
		}
	case 2:
		key := parameters[0].(string)
		value := parameters[1].(int64)
		if err := c.Client.IncrBy(c.Context, key, value).Err(); err != nil {
			logger.ErrorString("Redis", "Increment", err.Error())
			return false
		}
	default:
		logger.ErrorString("Redis", "Increment", "参数过多")
		return false
	}
	return true
}

// Decrement 当参数只有 1 个时，为 key，其值减去 1。
// 当参数有 2 个时，第一个参数为 key ，第二个参数为要减去的值 int64 类型。
func (c Client) Decrement(parameters ...interface{}) bool {
	switch len(parameters) {
	case 1:
		key := parameters[0].(string)
		if err := c.Client.Decr(c.Context, key).Err(); err != nil {
			logger.ErrorString("Redis", "Decrement", err.Error())
			return false
		}
	case 2:
		key := parameters[0].(string)
		value := parameters[1].(int64)
		if err := c.Client.DecrBy(c.Context, key, value).Err(); err != nil {
			logger.ErrorString("Redis", "Decrement", err.Error())
			return false
		}
	default:
		logger.ErrorString("Redis", "Decrement", "参数过多")
		return false
	}
	return true
}
