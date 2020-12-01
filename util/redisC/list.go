package redisC

type listRds struct {
}

//向列表头插入元素
func (l *listRds) LPush(db int64, key string, value interface{}) *Reply {
	c := pool.Get()
	defer c.Close()

	if r := selectDB(c, db); r.error != nil {
		return r
	}

	return getReply(c.Do("lpush", key, value))
}

//当列表存在则将元素插入表头
func (l *listRds) LPushx(db int64, key string, value interface{}) *Reply {
	c := pool.Get()
	defer c.Close()

	if r := selectDB(c, db); r.error != nil {
		return r
	}

	return getReply(c.Do("lpushx", key, value))
}

//将指定元素插入列表末尾
func (l *listRds) RPush(db int64, key string, value interface{}) *Reply {
	c := pool.Get()
	defer c.Close()

	if r := selectDB(c, db); r.error != nil {
		return r
	}

	return getReply(c.Do("rpush", key, value))
}

//将指定元素列表插入列表末尾
func (l *listRds) RPushL(db int64, key string, value []interface{}) *Reply {
	var (
		valueHandled = make([]interface{}, 0 ,len(value) + 1)
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
	return getReply(c.Do("rpush", valueHandled...))
}

//当列表存在则将元素插入表尾
func (l *listRds) RPushx(db int64, key string, value interface{}) *Reply {
	c := pool.Get()
	defer c.Close()

	if r := selectDB(c, db); r.error != nil {
		return r
	}

	return getReply(c.Do("rpushx", key, value))
}

//将元素插入指定位置position:BEFORE|AFTER,当 pivot 不存在于列表 key 时，不执行任何操作。当 key 不存在时， key 被视为空列表，不执行任何操作。
func (l *listRds) LInsert(db int64, key, position, pivot, value string) *Reply {
	c := pool.Get()
	defer c.Close()

	if r := selectDB(c, db); r.error != nil {
		return r
	}

	return getReply(c.Do("linsert", key, position, pivot, value))
}

//返回列表头元素
func (l *listRds) LPop(db int64, key string) *Reply {
	c := pool.Get()
	defer c.Close()

	if r := selectDB(c, db); r.error != nil {
		return r
	}

	return getReply(c.Do("lpop", key))
}

//阻塞并弹出头元素
func (l *listRds) BLpop(db int64, key, timeout interface{}) *Reply {
	c := pool.Get()
	defer c.Close()

	if r := selectDB(c, db); r.error != nil {
		return r
	}

	return getReply(c.Do("blpop", key, timeout))
}

//返回列表尾元素
func (l *listRds) RPop(db int64, key string) *Reply {
	c := pool.Get()
	defer c.Close()

	if r := selectDB(c, db); r.error != nil {
		return r
	}

	return getReply(c.Do("cpop", key))
}

//阻塞并弹出末尾元素
func (l *listRds) BRpop(db int64, key, timeout interface{}) *Reply {
	c := pool.Get()
	defer c.Close()

	if r := selectDB(c, db); r.error != nil {
		return r
	}

	return getReply(c.Do("brpop", key, timeout))
}

//返回指定位置的元素
func (l *listRds) LIndex(db int64, key string, index interface{}) *Reply {
	c := pool.Get()
	defer c.Close()

	if r := selectDB(c, db); r.error != nil {
		return r
	}

	return getReply(c.Do("lindex", key, index))
}

//获取指定区间的元素
func (l *listRds) LRange(db int64, key string, start, stop interface{}) *Reply {
	c := pool.Get()
	defer c.Close()

	if r := selectDB(c, db); r.error != nil {
		return r
	}

	return getReply(c.Do("lrange", key, start, stop))
}

//设置指定位元素
func (l *listRds) LSet(db int64, key string, index, value interface{}) *Reply {
	c := pool.Get()
	defer c.Close()

	if r := selectDB(c, db); r.error != nil {
		return r
	}

	return getReply(c.Do("lset", key, index, value))
}

//弹出source尾元素并返回，将弹出元素插入destination列表的开头
func (l *listRds) RPoplpush(db int64, key, source, destination string) *Reply {
	c := pool.Get()
	defer c.Close()

	if r := selectDB(c, db); r.error != nil {
		return r
	}

	return getReply(c.Do("rpoplpush ", key, source, destination))
}

//阻塞并弹出尾元素，将弹出元素插入另一列表的开头
func (l *listRds) BRpoplpush(db int64, key, source, destination string, timeout interface{}) *Reply {
	c := pool.Get()
	defer c.Close()

	if r := selectDB(c, db); r.error != nil {
		return r
	}

	return getReply(c.Do("brpoplpush ", key, source, destination, timeout))
}

//移除元素,count = 0 : 移除表中所有与 value 相等的值,count!=0,移除与 value 相等的元素，数量为 count的绝对值
func (l *listRds) LRem(db int64, key string, count, value interface{}) *Reply {
	c := pool.Get()
	defer c.Close()

	if r := selectDB(c, db); r.error != nil {
		return r
	}

	return getReply(c.Do("lrem", key, count, value))
}

//列表裁剪，让列表只保留指定区间内的元素，不在指定区间之内的元素都将被删除。-1 表示尾部
func (l *listRds) LTrim(db int64, key string, start, stop interface{}) *Reply {
	c := pool.Get()
	defer c.Close()

	if r := selectDB(c, db); r.error != nil {
		return r
	}

	return getReply(c.Do("ltrim", key, start, stop))
}

