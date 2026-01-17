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
		RedisURL:          dotenv.GetString("REDIS_URL", "localhost"),
		RedisPass:         dotenv.GetString("REDIS_PASS", ""),
		RedisTLS:          dotenv.GetBool("IS_REDIS_TLS", false),
		RedisPoolSize:     dotenv.GetInt("REDIS_POOL_SIZE", 100),
		RedisMinIdleConns: dotenv.GetInt("REDIS_MIN_IDLE_CONS", 30),
	}
}
