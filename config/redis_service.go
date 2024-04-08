package application

import "github.com/redis/go-redis/v9"

type RedisConfig struct {
	Addr     string `json:"addr"`
	Password string `json:"password"`
	DB       int    `json:"db"`
}

var RedisClient *redis.Client

func NewRedisCache(c *RedisConfig) *redis.Client {
	RedisClient = redis.NewClient(&redis.Options{
		Addr:     c.Addr,
		Password: c.Password, // no password set
		DB:       c.DB,       // use default DB
	})
	return RedisClient
}
