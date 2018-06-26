package helpers

import (
	"github.com/cihub/seelog"
	"github.com/garyburd/redigo/redis"
	"go_service/system"
	"time"
)

var (
	RedisClient *redis.Pool
)

func InitRedis() {
	// 建立连接池
	RedisClient = &redis.Pool{
		// 从配置文件获取maxidle以及maxactive，取不到则用后面的默认值
		MaxIdle:     system.GetConfiguration().RedisMaxIdle,
		MaxActive:   system.GetConfiguration().RedisMaxActive,
		IdleTimeout: time.Duration(system.GetConfiguration().RedisIdleTimeout) * time.Second,
		Dial: func() (redis.Conn, error) {
			c, err := redis.Dial("tcp", system.GetConfiguration().RedisHost)
			if err != nil {
				seelog.Critical(err.Error())
				return nil, err
			}

			if RedisPWD := system.GetConfiguration().RedisPWD; RedisPWD != "" {
				if _, err := c.Do("AUTH", RedisPWD); err != nil {
					c.Close()
					seelog.Critical(err.Error())
					return nil, err
				}
			}

			// 选择db
			c.Do("SELECT", system.GetConfiguration().RedisDB)
			return c, nil
		},
	}
}
