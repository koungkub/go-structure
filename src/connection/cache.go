package connection

import (
	"github.com/go-redis/redis"
	"github.com/spf13/viper"
)

func GetCacheConnection() *redis.Client {

	cache := redis.NewClient(&redis.Options{
		Addr:     viper.GetString("APP.CACHE.HOST"),
		Password: viper.GetString("APP.CACHE.PASSWORD"),
		DB:       viper.GetInt("APP.CACHE.DATABASE"),
	})

	return cache
}
