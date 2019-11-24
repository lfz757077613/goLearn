package myRedis

import (
	"errors"
	"github.com/go-redis/redis"
	"github.com/lfz757077613/goLearn/utils/myConf"
	"github.com/lfz757077613/goLearn/utils/shutDownhook"
	"github.com/sirupsen/logrus"
	"time"
)

var client *redis.Client

func init() {
	client = redis.NewClient(&redis.Options{
		Addr:        myConf.GetString("redis", "addr", "localhost:6379"),
		PoolSize:    myConf.GetInt("redis", "poolSize", 10),
		PoolTimeout: time.Duration(myConf.GetInt64("redis", "poolTimeout", 10)),
	})
	pong, err := Ping()
	if err != nil || pong != "PONG" {
		logrus.Panic("init redis client error")
	}
	shutDownhook.AddShutdownHook(func() {
		if err := client.Close(); err!=nil {
			logrus.Errorf("myRedis client close error: [%s]", err)
		}
	})
}

func Ping() (string, error) {
	return client.Ping().Result()
}

func Set(key, value string) error {
	resp, err := client.Set(key, value, 0).Result()
	if err != nil {
		return err
	}
	if resp != "ok" {
		return errors.New("redis set not ok")
	}
	return nil
}
