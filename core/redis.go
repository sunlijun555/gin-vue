package core

import (
	"go-vue/config"
	"go-vue/global"

	"github.com/go-redis/redis/v8"
	"github.com/sirupsen/logrus"
)

var (
	RedisDb *redis.Client
)

func RedisInit() error {
	RedisDb = redis.NewClient(&redis.Options{
		Addr:     config.Config.Redis.Address,
		Password: config.Config.Redis.Password,
		DB:       config.Config.Redis.Db,
	})
	_, err := RedisDb.Ping(global.Ctx).Result()
	if err != nil {
		logrus.Errorf("%v", err)
		return err
	}
	logrus.Infof("[redis]连接成功")
	return nil
}
