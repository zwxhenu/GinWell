package initialize

import (
	"GinWell-Server/global"
	"context"
	"github.com/go-redis/redis/v8"
	"go.uber.org/zap"
)

func Redis() {
	redisConfig := global.GW_CONFIG.Redis
	redisClient := redis.NewClient(&redis.Options{
		Addr:     redisConfig.Addr,
		Password: redisConfig.Password,
		DB:       redisConfig.DB,
	})
	ping, err := redisClient.Ping(context.Background()).Result()
	if err != nil {
		global.GW_LOG.Error("redis connect ping failed, err:", zap.Error(err))
	} else {
		global.GW_LOG.Info("redis connect ping response:", zap.String("pong", ping))
		global.GW_REDIS = redisClient
	}
}
