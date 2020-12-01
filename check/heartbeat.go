package check

import (
	"time"
	"log"

	"shaim/util/redisC"
	"shaim/conf"
)


func Heartbeat() {
	var err error
	go checkOnline()
	for {
		if err = redisC.Conn.String.Set(0, conf.RedisCheckIncKey+conf.GetHostStr(), 0, 5).Error(); err != nil {
			log.Println(err)
			return
		}
		if err = redisC.Conn.Set.SAddString(0, conf.RedisCheckListKey, []string{conf.GetHostStr()}).Error(); err != nil {
			log.Println(err)
			return
		}
		time.Sleep(3 * time.Second)
	}
}

func SetHash(room, username string) {
	redisC.Conn.Hash.HSet(0, conf.RedisCheckHashKey+room, username, conf.GetHostStr())
}

func DelHash(room, username string) {
	redisC.Conn.Hash.HDel(0, conf.RedisCheckHashKey+room, username)
}