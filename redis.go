package main

import (
	"time"

	"github.com/go-redis/redis"
	"github.com/sirupsen/logrus"
)

var redisClient *redis.Client

func initRedis() {
	for {
		redisClient = redis.NewClient(&redis.Options{
			Addr:     conf.RedisHost + ":" + conf.RedisPort,
			Password: conf.RedisPassword, // no password set
			DB:       conf.RedisDB,       // use default DB
		})
		_, err := redisClient.Ping().Result()
		if err != nil {
			logrus.Error("redis connect error:", err, " Retry in 2 seconds!")
			time.Sleep(time.Second * 2)
			continue
		}
		logrus.Info("Redis connect successful!")
		break
	}
}
