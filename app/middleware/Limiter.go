package middleware

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/go-redis/redis"
	"net/http"
	"slim/app/libs"
	"time"
)

func Limiter(ctx *gin.Context) {
	now := time.Now().UnixNano()
	key := "REDIS_LIMITER"
	userCntKey := fmt.Sprint(ctx.ClientIP(), ":", key)

	//五秒限流
	var limit int64 = 10
	dura := time.Second * 5

	redisInstance := libs.GetRedisInstance(0)
	//删除有序集合中的五秒之前的数据
	redisInstance.ZRemRangeByScore(userCntKey,
		"0",
		fmt.Sprint(now-(dura.Nanoseconds()))).Result()

	reqs, _ := redisInstance.ZCard(userCntKey).Result()

	if reqs >= limit {
		ctx.AbortWithStatusJSON(http.StatusTooManyRequests, gin.H{
			"status":  http.StatusTooManyRequests,
			"message": "too many request",
		})
		return
	}

	ctx.Next()
	redisInstance.ZAddNX(userCntKey, redis.Z{Score: float64(now), Member: float64(now)})
	redisInstance.Expire(userCntKey, dura)
}
