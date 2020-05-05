package repository

import "github.com/gomodule/redigo/redis"

func newPool() *redis.Pool {
	return &redis.Pool{
		MaxIdle:   80,
		MaxActive: 12000, // max number of connections
		Dial: func() (redis.Conn, error) {
			c, err := redis.Dial("tcp", ":6379")
			_, err = c.Do("AUTH", "thisissohardpassword")
			if err != nil {
				panic(err.Error())
			}
			return c, err
		},
	}

}

var pool = newPool()

// RedisInit : Initialize redis connection
func RedisInit() redis.Conn {
	c := pool.Get()
	return c
}
