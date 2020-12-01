package redisC

import (
	"github.com/gomodule/redigo/redis"
)

type stringRds struct {
}

// 设置值
func (s *stringRds) Set(db int64, key string, value interface{}, expire ...int64) *Reply {
	c := pool.Get()
	defer c.Close()

	if r := selectDB(c, db); r.error != nil {
		return r
	}

	if len(expire) == 0 {
		return getReply(c.Do("set", key, value))
	}
	return getReply(c.Do("set", key, value, "ex", expire[0]))
}

// 获取值
func (s *stringRds) Get(db int64, key string) *Reply {
	c := pool.Get()
	defer c.Close()

	if r := selectDB(c, db); r.error != nil {
		return r
	}

	return getReply(c.Do("get", key))
}

// key不存在是在设置值
func (s *stringRds) SetNX(db int64, key string, value interface{}, expire int64) *Reply {
	c := pool.Get()
	defer c.Close()

	if r := selectDB(c, db); r.error != nil {
		return r
	}

	if expire == 0 {
		return getReply(c.Do("setnx", key, value))
	}

	return getReply(c.Do("set", key, value, "EX", expire, "NX"))
}

// 	设置并返回旧值
func (s *stringRds) GetSet(db int64, key string, value interface{}) *Reply {
	c := pool.Get()
	defer c.Close()

	if r := selectDB(c, db); r.error != nil {
		return r
	}

	return getReply(c.Do("getset", key, value))
}

// 	设置key并指定生存时间
func (s *stringRds) SetEX(db int64, key string, value interface{}, seconds int64) *Reply {
	c := pool.Get()
	defer c.Close()

	if r := selectDB(c, db); r.error != nil {
		return r
	}

	return getReply(c.Do("setex", key, seconds, value))
}

// 	设置key值并指定生存时间(毫秒)
func (s *stringRds) PSetEX(db int64, key string, value interface{}, milliseconds int64) *Reply {
	c := pool.Get()
	defer c.Close()

	if r := selectDB(c, db); r.error != nil {
		return r
	}

	return getReply(c.Do("psetex", key, milliseconds, value))
}

// 设置子字符串
func (s *stringRds) SetRange(db int64, key string, value interface{}, offset int64) *Reply {
	c := pool.Get()
	defer c.Close()

	if r := selectDB(c, db); r.error != nil {
		return r
	}

	return getReply(c.Do("setrange", key, offset, value))
}

// 	获取子字符串
func (s *stringRds) GetRange(db int64, key string, start, end int64) *Reply {
	c := pool.Get()
	defer c.Close()

	if r := selectDB(c, db); r.error != nil {
		return r
	}

	return getReply(c.Do("getrange", key, start, end))
}

// 设置多个值
func (s *stringRds) MSet(db int64, kv map[string]interface{}) *Reply {
	c := pool.Get()
	defer c.Close()

	if r := selectDB(c, db); r.error != nil {
		return r
	}

	return getReply(c.Do("mset", redis.Args{}.AddFlat(kv)))
}

// key不存在时设置多个值
func (s *stringRds) MSetNx(db int64, kv map[string]interface{}) *Reply {
	c := pool.Get()
	defer c.Close()

	if r := selectDB(c, db); r.error != nil {
		return r
	}

	return getReply(c.Do("msetnx", redis.Args{}.AddFlat(kv)))
}

// 返回多个key的值
func (s *stringRds) MGet(db int64, keys []string) *Reply {
	c := pool.Get()
	defer c.Close()

	if r := selectDB(c, db); r.error != nil {
		return r
	}

	return getReply(c.Do("mget", redis.Args{}.AddFlat(keys)...))
}

// 自增
func (s *stringRds) Incr(db int64, key string) *Reply {
	c := pool.Get()
	defer c.Close()

	if r := selectDB(c, db); r.error != nil {
		return r
	}

	return getReply(c.Do("incr", key))
}

// 增加指定值
func (s *stringRds) IncrBy(db int64, key string, increment int64) *Reply {
	c := pool.Get()
	defer c.Close()

	if r := selectDB(c, db); r.error != nil {
		return r
	}

	return getReply(c.Do("incrby", key, increment))
}

// 增加一个浮点值
func (s *stringRds) IncrByFloat(db int64, key string, increment float64) *Reply {
	c := pool.Get()
	defer c.Close()

	if r := selectDB(c, db); r.error != nil {
		return r
	}

	return getReply(c.Do("incrbyfloat", key, increment))
}

// 自减
func (s *stringRds) Decr(db int64, key string) *Reply {
	c := pool.Get()
	defer c.Close()

	if r := selectDB(c, db); r.error != nil {
		return r
	}

	return getReply(c.Do("decr", key))
}

// 自减指定值
func (s *stringRds) DecrBy(db int64, key string, increment int64) *Reply {
	c := pool.Get()
	defer c.Close()

	if r := selectDB(c, db); r.error != nil {
		return r
	}

	return getReply(c.Do("decrby", key, increment))
}
