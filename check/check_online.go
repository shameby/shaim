package check

import (
	"math/rand"
	"time"

	"shaim/util/redisC"
	"shaim/conf"
)

const (
	redisLockKey = "im"
	timeWaitMin = 2000
	timeWait = 2000
)

var mutex *redisC.Lock

func init() {
	rand.Seed(time.Now().Unix())
	mutex = redisC.NewLockWithTimeout(redisLockKey, (timeWaitMin+rand.Int63n(timeWait))/1000)
}

func checkOnline() {
	var (
		suc bool
		err error
	)
	timer := time.NewTimer(time.Duration(timeWaitMin+rand.Int63n(timeWait)) * time.Millisecond)
	for {
		token := conf.GetHostStr()
		timer.Reset(time.Duration(timeWaitMin + rand.Int63n(rand.Int63n(timeWait))) * time.Millisecond)
		if suc, err = mutex.TryLock(token); !suc || err != nil {
			<-timer.C
			continue
		}
		hostList, err := redisC.Conn.Set.SMembers(0, conf.RedisCheckListKey).Strings()
		if err != nil {
			<-timer.C
			continue
		}
		for _, host := range hostList {
			if exist, err := redisC.Conn.Key.Exists(0, conf.RedisCheckIncKey+host).Bool(); !exist || err != nil {
				redisC.Conn.Set.SRem(0, conf.RedisCheckListKey, []interface{}{host}).Bool()
			}
		}
		<-timer.C
	}
}
