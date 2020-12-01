package redisC

import (
	"github.com/gomodule/redigo/redis"
)

func selectDB(c redis.Conn, db int64) *Reply {
	db += redisPrefix
	return getReply(c.Do("select", db))
}
