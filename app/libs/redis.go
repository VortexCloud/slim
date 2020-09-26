package libs

import (
	"github.com/go-redis/redis"
	"sync"
)

var (
	RedisClientInstance *redis.Client
	redisOnce           sync.Once
)

func GetRedisInstance(db int) *redis.Client {
	if RedisClientInstance == nil {
		redisOnce.Do(func() {
			RedisClientInstance = redis.NewClient(&redis.Options{
				Addr:       "172.20.1.3:6379", //os.Getenv("REDIS_HOST")
				Password:   "",
				MaxRetries: 3,
			})
		})
	}
	RedisClientInstance.Do("select", db)
	return RedisClientInstance
}
