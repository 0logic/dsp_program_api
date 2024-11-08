package Redis

import (
	"context"
	"dsp_program_api/config"
	"github.com/redis/go-redis/v9"
)

var ctx = context.Background()
var RedisCon *redis.Client

func init() {

	rdb := redis.NewClient(&redis.Options{
		Addr:     config.REDIS_ADDRESS,
		Password: config.REDIS_PASSWORD,  // 密码
		DB:       config.REDIS_DB,        // 数据库
		PoolSize: config.REDIS_POOL_SIZE, // 连接池大小
	})

	_, err := rdb.Ping(ctx).Result()
	if err != nil {
		defer rdb.Close()
	}
	RedisCon = rdb
}
