package redis

import "doit/urlshortener/pkg/dotenv"

type RedisConfig struct {
	RedisDB           int
	RedisURL          string
	RedisPass         string
	RedisTLS          bool
	RedisPoolSize     int
	RedisMinIdleConns int
}

func RedisConfiguration() *RedisConfig {
	return &RedisConfig{
		RedisDB:           dotenv.GetInt("REDIS_DB", 0),
		RedisURL:          dotenv.REDISURL(),
		RedisPass:         dotenv.REDISPASS(),
		RedisTLS:          dotenv.ISREDISTLS(),
		RedisPoolSize:     dotenv.GetInt("REDIS_POOL_SIZE", 100),
		RedisMinIdleConns: dotenv.GetInt("REDIS_MIN_IDLE_CONS", 30),
	}
}
