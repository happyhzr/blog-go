package redis

import (
	"github.com/garyburd/redigo/redis"

	"github.com/insisthzr/blog-back/config"
)

var (
	pool *redis.Pool
)

func newPool(server string) *redis.Pool {
	return &redis.Pool{
		Dial: func() (redis.Conn, error) {
			c, err := redis.Dial("tcp", server)
			if err != nil {
				return nil, err
			}
			return c, err
		},
	}
}

func GetConn() redis.Conn {
	return pool.Get()
}

func Run() {
	pool = newPool(config.RedisAddr)
}
