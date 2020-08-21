package redis

import (
	"fmt"
	"github.com/go-redis/redis"
	"github.com/z1px/framework/config"
	"github.com/z1px/framework/logs"
)

// RedisClient Redis缓存客户端单例
var Redis *redis.Client

// Redis 在中间件中初始化redis连接
func Init() {
	// 获取数据库配置
	conf := config.DBConf.Redis

	// 连接
	client := redis.NewClient(&redis.Options{
		Addr:       fmt.Sprintf("%s:%d", conf.Host, conf.Port),
		Password:   "",
		DB:         conf.Db,
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
