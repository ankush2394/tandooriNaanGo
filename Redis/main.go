package Redis

import (
	"context"
	"fmt"
	"github.com/go-redis/redis"
	log "github.com/sirupsen/logrus"
)

var (
	redis_cache_addr       = "127.0.0.1:6379"//os.Getenv("TRENDING_CACHE_ADDR")
	RedisClient            *redis.Client
	MaxConnection          = 10
)

func init() {

	redisOption := redis.Options{}
	redisOption.PoolSize = MaxConnection

	log.Info("redis address is ", redis_cache_addr)
	redisOption.Addr = redis_cache_addr

	redisOption.Password = ""
	redisOption.DB = 0                   		//0 is used for default settings
	redisOption.ReadTimeout = 0

	RedisClient = redis.NewClient(&redisOption)
	var ctx context.Context
	pong , err := RedisClient.Ping(ctx).Result()
	fmt.Println(err, pong)
}

func GetInstance() *redis.Client {
	return RedisClient
}