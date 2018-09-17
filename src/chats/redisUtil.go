package chats

import (
	"github.com/gomodule/redigo/redis"
	"strconv"
)

var pool *redis.Pool

func initRedis(addr string, port int) (err error) {
	pool = &redis.Pool{
		Dial: func() (redis.Conn, error) {
			return redis.Dial("tcp", addr+":"+strconv.Itoa(port))
		},
		MaxIdle:     10,
		MaxActive:   11,
		IdleTimeout: 11,
		Wait:        false,
	}
	return
}
