package redis

import (
	"fmt"
	"github.com/go-redis/redis"
	"github.com/z1px/framework/conf"
	"github.com/z1px/framework/logs"
)

// RedisClient Redis缓存客户端单例
var Redis *redis.Client

// 连接REDIS数据库
func Connect() {
	// 获取数据库配置
	redisConf := conf.Base.Redis

	// 连接
	client := redis.NewClient(&redis.Options{
		Addr:       fmt.Sprintf("%s:%d", redisConf.Host, redisConf.Port),
		Password:   redisConf.Password,
		DB:         redisConf.Db,
		MaxRetries: 1,
	})

	if _, err := client.Ping().Result(); err != nil {
		// handle error
		logs.ErrPrintf("redis connect error：%v\n", err)
	}

	Redis = client
}

// 检测Mysql是否有连接
func IsOpen() bool {
	if _, err := Redis.Ping().Result(); err != nil {
		logs.ErrPrintf("redis ping error：%v\n", err)
		return false
	} else {
		return true
	}
}

// 关闭连接
func Close() {
	err := Redis.Close()
	if err != nil {
		logs.ErrPrintf("redis close error：%v\n", err)
	}
}
