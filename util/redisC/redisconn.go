package redisC

import (
	"fmt"
	"github.com/gomodule/redigo/redis"
	"time"
	"shaim/conf"
)

var (
	pool        *redis.Pool
	redisPrefix int64
)

//newRedisPool:创建redis连接池
func newRedisPool(redisHost, redisPass string, prefix int64) *redis.Pool {
	redisPrefix = prefix
	return &redis.Pool{
		MaxIdle:     100,
		MaxActive:   200,
		IdleTimeout: 300 * time.Second,
		Dial: func() (redis.Conn, error) {
			c, err := redis.Dial("tcp", redisHost)
			if err != nil {
				panic(fmt.Sprintf("redis init err: %v", err))
			}
			if _, err := c.Do("AUTH", redisPass); err != nil {
				c.Close()
				panic(fmt.Sprintf("redis init err: %v", err))
			}
			return c, nil
		},
		//定时检查redis是否出状况
		TestOnBorrow: func(conn redis.Conn, t time.Time) error {
			if time.Since(t) < time.Minute {
				return nil
			}
			_, err := conn.Do("PING")
			return err
		},
	}
}

//初始化redis连接池
func init() {
	pool = newRedisPool(conf.C.RedisHost, conf.C.RedisPwd, conf.C.RedisPrefix)
}

//对外暴露连接池
func RedisPool() *redis.Pool {
	return pool
}
