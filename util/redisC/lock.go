package redisC

import (
	"fmt"

	"time"
)

const (
	DefaultTimeOut = 10
)

type Lock struct {
	timeout int64
	key     string
	conn    *RedigoPack
}

func (lock *Lock) Lock(token string) (ok bool, err error) {
	var (
		success bool
		ttlTime int64
		end     time.Time
	)
	end = time.Now().Add(time.Second * time.Duration(lock.timeout))
	for time.Now().Before(end) {
		if success, err = lock.conn.String.SetNX(0, lock.getKey(), token, lock.timeout).Bool(); success {
			return true, nil
		} else if ttlTime, _ = lock.conn.Key.TTL(0, lock.getKey()).Int64(); ttlTime == -1 {
			lock.conn.Key.Expire(0, lock.getKey(), lock.timeout)
		}
		time.Sleep(time.Microsecond)
	}

	return
}

func (lock *Lock) TryLock(token string) (ok bool, err error) {
	var (
		success string
	)
	if success, err = lock.conn.String.SetNX(0, lock.getKey(), token, lock.timeout).String(); success == "OK" {
		return true, nil
	}
	if err != nil {
		return false, err
	}
	return false, nil
}

func (lock *Lock) Unlock() (err error) {
	_, err = lock.conn.Key.Del(0, lock.getKey()).Result()
	return
}

func (lock *Lock) getKey() string {
	return fmt.Sprintf("redislock:%s", lock.key)
}

func NewLock(key string) (lock *Lock) {
	return NewLockWithTimeout(key, DefaultTimeOut)
}

func NewLockWithTimeout(key string, timeout int64) (lock *Lock) {
	if timeout == 0 {
		return nil
	}
	lock = &Lock{timeout, key, Conn}

	return
}
