package Redis

import (
	"github.com/go-redis/redis"
	log "github.com/sirupsen/logrus"
)

var (
	redis_cache_addr       = "localhost:6379"//os.Getenv("TRENDING_CACHE_ADDR")
	RedisClient            *redis.Client
	RedisMaxConnection     = 10

)

func init() {

	redisOption := redis.Options{}
	redisOption.PoolSize = RedisMaxConnection

	log.Info("redis address is ", redis_cache_addr)
	redisOption.Addr = redis_cache_addr

	redisOption.Password = ""
	redisOption.DB = 0

	RedisClient = redis.NewClient(&redisOption)
}

func GetInstance() *redis.Client {
	return RedisClient
}