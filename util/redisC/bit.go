package redisC

import (
	"github.com/gomodule/redigo/redis"
)

type bitRds struct {
}

func (b *bitRds) SetBit(db int64, key string, offset, value int64) *Reply {
	c := pool.Get()
	defer c.Close()

	if r := selectDB(c, db); r.error != nil {
		return r
	}

	return getReply(c.Do("setbit", key, offset, value))
}

func (b *bitRds) GetBit(db int64, key string, offset int64) *Reply {
	c := pool.Get()
	defer c.Close()

	if r := selectDB(c, db); r.error != nil {
		return r
	}

	return getReply(c.Do("getbit", key, offset))
}

func (b *bitRds) BitCount(db int64, key string, interval ...int64) *Reply {
	c := pool.Get()
	defer c.Close()

	if r := selectDB(c, db); r.error != nil {
		return r
	}

	if len(interval) == 2 {
		return getReply(c.Do("bitcount", key, interval[0], interval[1]))
	}
	return getReply(c.Do("bitcount", key))
}

// opt 包含 and、or、xor、not
func (b *bitRds) BitTop(db int64, opt, destKey string, keys ...string) *Reply {
	c := pool.Get()
	defer c.Close()

	if r := selectDB(c, db); r.error != nil {
		return r
	}

	return getReply(c.Do("bitop", opt, redis.Args{}.Add(keys).AddFlat(keys)))
}

