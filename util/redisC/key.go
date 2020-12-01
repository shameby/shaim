package redisC

import (
	"github.com/gomodule/redigo/redis"
)

type keyRds struct {
}

// 	查找键 [*模糊查找]
func (k *keyRds) Keys(db int64, key string) *Reply {
	c := pool.Get()
	defer c.Close()

	if r := selectDB(c, db); r.error != nil {
		return r
	}

	return getReply(c.Do("keys", key))
}

// 监视 如果在事务执行之前这个(或这些) key 被其他命令所改动，那么事务将被打断
func (k *keyRds) Watch(db int64, key string) *Reply {
	c := pool.Get()
	defer c.Close()

	if r := selectDB(c, db); r.error != nil {
		return r
	}

	return getReply(c.Do("watch", key))
}

// 	判断key是否存在
func (k *keyRds) Exists(db int64, key string) *Reply {
	c := pool.Get()
	defer c.Close()

	if r := selectDB(c, db); r.error != nil {
		return r
	}

	return getReply(c.Do("exists", key))
}

// 随机返回一个key
func (k *keyRds) RandomKey(db int64) *Reply {
	c := pool.Get()
	defer c.Close()

	if r := selectDB(c, db); r.error != nil {
		return r
	}

	return getReply(c.Do("randomkey"))
}

// 返回值类型
func (k *keyRds) Type(db int64, key string) *Reply {
	c := pool.Get()
	defer c.Close()

	if r := selectDB(c, db); r.error != nil {
		return r
	}

	return getReply(c.Do("type", key))
}

// 删除key
func (k *keyRds) Del(db int64, keys ...string) *Reply {
	c := pool.Get()
	defer c.Close()

	if r := selectDB(c, db); r.error != nil {
		return r
	}

	return getReply(c.Do("del", redis.Args{}.AddFlat(keys)...))
}

//重命名
func (k *keyRds) Rename(db int64, key, newKey string) *Reply {
	c := pool.Get()
	defer c.Close()

	if r := selectDB(c, db); r.error != nil {
		return r
	}

	return getReply(c.Do("rename", key, newKey))
}

// 仅当newkey不存在时重命名
func (k *keyRds) RenameNX(db int64, key, newKey string) *Reply {
	c := pool.Get()
	defer c.Close()

	if r := selectDB(c, db); r.error != nil {
		return r
	}

	return getReply(c.Do("renamenx", key, newKey))
}

//	序列化key
func (k *keyRds) Dump(db int64, key string) *Reply {
	c := pool.Get()
	defer c.Close()

	if r := selectDB(c, db); r.error != nil {
		return r
	}

	return getReply(c.Do("dump", key))
}

// 反序列化
func (k *keyRds) Restore(db int64, key string, ttl, serializedValue interface{}) *Reply {
	c := pool.Get()
	defer c.Close()

	if r := selectDB(c, db); r.error != nil {
		return r
	}

	return getReply(c.Do("restore", key, ttl, serializedValue))
}

// 秒
func (k *keyRds) Expire(db int64, key string, seconds int64) *Reply {
	c := pool.Get()
	defer c.Close()

	if r := selectDB(c, db); r.error != nil {
		return r
	}

	return getReply(c.Do("expire", key, seconds))
}

// 秒
func (k *keyRds) ExpireAt(db int64, key string, timestamp int64) *Reply {
	c := pool.Get()
	defer c.Close()

	if r := selectDB(c, db); r.error != nil {
		return r
	}

	return getReply(c.Do("expireat", key, timestamp))
}

// 毫秒
func (k *keyRds) Persist(db int64, key string) *Reply {
	c := pool.Get()
	defer c.Close()

	if r := selectDB(c, db); r.error != nil {
		return r
	}

	return getReply(c.Do("persist", key))
}

// 毫秒
func (k *keyRds) PersistAt(db int64, key string, milliSeconds int64) *Reply {
	c := pool.Get()
	defer c.Close()

	if r := selectDB(c, db); r.error != nil {
		return r
	}

	return getReply(c.Do("persistat", key, milliSeconds))
}

// 秒
func (k *keyRds) TTL(db int64, key string) *Reply {
	c := pool.Get()
	defer c.Close()

	if r := selectDB(c, db); r.error != nil {
		return r
	}

	return getReply(c.Do("ttl", key))
}

// 毫秒
func (k *keyRds) PTTL(db int64, key string) *Reply {
	c := pool.Get()
	defer c.Close()

	if r := selectDB(c, db); r.error != nil {
		return r
	}

	return getReply(c.Do("pttl", key))
}

//	同实例不同库间的键移动
func (k *keyRds) Move(dbS int64, key string, dbT int64) *Reply {
	c := pool.Get()
	defer c.Close()

	if r := selectDB(c, dbS); r.error != nil {
		return r
	}

	return getReply(c.Do("move", key, dbT))
}
