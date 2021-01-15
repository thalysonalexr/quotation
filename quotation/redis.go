package quotation

import (
	"log"
	"os"
	"time"

	"github.com/fatih/color"
	"github.com/gomodule/redigo/redis"
)

// ConfigRedis configurate redis pool
func ConfigRedis() *redis.Pool {
	var redisPool = &redis.Pool{
		MaxActive: 1,
		MaxIdle:   1,
		Wait:      true,
		Dial: func() (redis.Conn, error) {
			c, err := redis.Dial(
				"tcp",
				os.Getenv("REDIS_HOST")+":"+os.Getenv("REDIS_PORT"),
				redis.DialPassword(os.Getenv("REDIS_PASSWORD")),
				redis.DialClientName("quotation"),
				redis.DialConnectTimeout(time.Millisecond*8),
			)
			if err != nil {
				color.Red(ErrClientRedisFailedConnection.Error())
				log.Fatalln(err)
				return c, err
			}
			color.Green(RedisConnectionSuccess)
			return c, nil
		},
	}

	return redisPool
}
