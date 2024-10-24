package core

import (
	"context"
	"filterate/global"
	"github.com/go-redis/redis/v8"
	"go.uber.org/zap"
	"time"
)

func Redis() {
	redisCfg := global.Config.RedisConf
	var client redis.UniversalClient
	client = redis.NewClient(&redis.Options{
		Addr:         redisCfg.Addr,
		Password:     redisCfg.Password,
		DB:           redisCfg.DB,
		PoolSize:     10,                // 线程最大连接数量
		MinIdleConns: 5,                 // 最小活跃数量
		MaxConnAge:   300 * time.Second, // 连接最大存活时间
		IdleTimeout:  240 * time.Second, // 连接最大闲置时间
	})

	global.Logger.Info(" 连接 redis", zap.String("addr", redisCfg.Addr), zap.Int("DB", redisCfg.DB))

	ping, err := client.Ping(context.Background()).Result()
	if err != nil {
		global.Logger.Error("redis 连接失败", zap.Error(err))
		panic(err)
	} else {
		global.Logger.Info(ping)
		//log.Default().Println(ping)
		global.Redis = client

		global.Logger.Info("redis 连接成功")
	}
}
