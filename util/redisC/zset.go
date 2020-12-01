package redisC

type zSetRds struct {
}

// []{score1, member1, score2, member2,.....}  添加元素
func (z *zSetRds) ZAdd(db int64, key string, value []interface{}) *Reply {
	var (
		valueHandled = make([]interface{}, 0, len(value)+1)
	)

	c := pool.Get()
	defer c.Close()

	if r := selectDB(c, db); r.error != nil {
		return r
	}

	valueHandled = append(valueHandled, key)
	for _, v := range value {
		valueHandled = append(valueHandled, v)
	}
	return getReply(c.Do("zadd", valueHandled...))
}

// 	增加元素权重
func (z *zSetRds) ZUncrBy(db int64, key string, increment, member interface{}) *Reply {
	c := pool.Get()
	defer c.Close()

	if r := selectDB(c, db); r.error != nil {
		return r
	}

	return getReply(c.Do("zuncrby", key, increment, member))
}

// 	增加元素权重
func (z *zSetRds) ZCard(db int64, key string) *Reply {
	c := pool.Get()
	defer c.Close()

	if r := selectDB(c, db); r.error != nil {
		return r
	}

	return getReply(c.Do("zcard", key))
}

// 	返回指定元素的排名
func (z *zSetRds) ZEank(db int64, key string, member interface{}) *Reply {
	c := pool.Get()
	defer c.Close()

	if r := selectDB(c, db); r.error != nil {
		return r
	}

	return getReply(c.Do("zrank", key, member))
}

// 	返回指定元素的权重
func (z *zSetRds) ZScore(db int64, key string, member interface{}) *Reply {
	c := pool.Get()
	defer c.Close()

	if r := selectDB(c, db); r.error != nil {
		return r
	}

	return getReply(c.Do("zscore", key, member))
}

// 返回集合两个权重间的元素数
func (z *zSetRds) ZCount(db int64, key string, min, max interface{}) *Reply {
	c := pool.Get()
	defer c.Close()

	if r := selectDB(c, db); r.error != nil {
		return r
	}

	return getReply(c.Do("zcount", key, min, max))
}

// 返回指定区间内的元素
func (z *zSetRds) ZRange(db int64, key string, start, stop interface{}, withScore ...bool) *Reply {
	c := pool.Get()
	defer c.Close()

	if r := selectDB(c, db); r.error != nil {
		return r
	}

	if len(withScore) > 0 && withScore[0] {
		return getReply(c.Do("zrange", key, start, stop, withScore))
	}
	return getReply(c.Do("zrange", key, start, stop))
}

// 通过score返回指定区间内的元素
func (z *zSetRds) ZRangeByScore(db int64, key string, start, stop interface{}, withScore ...bool) *Reply {
	c := pool.Get()
	defer c.Close()

	if r := selectDB(c, db); r.error != nil {
		return r
	}

	if len(withScore) > 0 && withScore[0] {
		return getReply(c.Do("zrangebyscore", key, start, stop, withScore))
	}
	return getReply(c.Do("zrangebyscore", key, start, stop))
}

// 倒序返回指定区间内的元素
func (z *zSetRds) ZRevrange(db int64, key string, start, stop interface{}, withScore ...bool) *Reply {
	c := pool.Get()
	defer c.Close()

	if r := selectDB(c, db); r.error != nil {
		return r
	}

	if len(withScore) > 0 && withScore[0] {
		return getReply(c.Do("zrevrange", key, start, stop, withScore))
	}
	return getReply(c.Do("zrevrange", key, start, stop))
}
