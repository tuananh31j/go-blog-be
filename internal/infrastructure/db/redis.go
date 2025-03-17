package db

import (
	"context"

	"nta-blog/internal/lib/logger"

	"github.com/redis/go-redis/v9"
)

func ConnectRedis(url string) *redis.Client {
	ctx := context.Background()
	opt, _ := redis.ParseURL(url)
	rdb := redis.NewClient(opt)
	err := rdb.Set(ctx, "key", "value", 0).Err()
	if err != nil {
		panic(err)
	}

	val, err := rdb.Get(ctx, "key").Result()
	if err != nil {
		panic(err)
	}
	logger.ZeroLog.Info().Msgf("key: %v", val)

	val2, err := rdb.Get(ctx, "key2").Result()
	if err == redis.Nil {
		logger.ZeroLog.Info().Msgf("key2 does not exist")
	} else if err != nil {
		panic(err)
	} else {
		logger.ZeroLog.Info().Msgf("key2: %v", val2)
	}
	return rdb
}

func DisconnectRedis(r *redis.Client) {
	if err := r.Close(); err != nil {
		logger.ZeroLog.Info().Msgf("Somethings wrong when close redis: >>>>>>>>%v", err.Error())
	}
}
