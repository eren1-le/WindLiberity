/*
 * @Author: eren dengdengd1222@mail.com
 * @Date: 2026-09-23 11:15:37
 * @LastEditors: eren dengdengd1222@mail.com
 * @LastEditTime: 2026-09-23 13:49:12
 * @FilePath: /WindLiberity/pkg/app/redis.go
 * @Description: e
 * 
 */
package app

import (
	"log"
	"time"

	"WindLiberity/pkg/config"
	redis2 "github.com/binbinly/pkg/storage/redis"
	"github.com/redis/go-redis/v9"
	"github.com/spf13/viper"
)
// redis client getter with default parameters
var Redis *redis.Client

type RedisConfig struct {
	Default redis2.Config
}

func InitRedis() *redis.Client {
	var cfg = &RedisConfig{}
	if err := loadRedisConf(cfg); err != nil {
		log.Fatalf("load redis conf err: %v", err)
	}

	c, err := redis2.NewClient(&cfg.Default)
	if err != nil {
		log.Fatalf("new redis client err: %v", err)
	}

	Redis = c
	return c
}


func loadRedisConf(cfg *RedisConfig) error {
	if err := config.Load("redis", cfg, func(v *viper.Viper) {
		v.SetDefault("default", map[string]any{
			"Addr":         "127.0.0.1:6379",
			"Password":     "",
			"DB":           0,
			"MinIdleConn":  200,
			"DialTimeout":  60 * time.Second,
			"ReadTimeout":  500 * time.Millisecond,
			"WriteTimeout": 500 * time.Millisecond,
			"PoolSize":     100,
			"PoolTimeout":  240 * time.Second,
		})
		v.BindEnv("default.url", "ANIME_REDIS_URL")
		v.BindEnv("default.addr", "ANIME_REDIS_ADDR")
		v.BindEnv("default.password", "ANIME_REDIS_PASSWORD")
		v.BindEnv("default.DB", "ANIME_REDIS_DB")
	}); err != nil {
		return err
	}
	return nil
}