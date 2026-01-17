package redis

import (
	"context"
	"crypto/tls"
	"doit/urlshortener/pkg/dotenv"
	"fmt"
	"time"

	"github.com/newrelic/go-agent/v3/integrations/nrredis-v9"
	"github.com/redis/go-redis/v9"
)

type redisClientFactory struct {
	cfg *RedisConfig
}

func NewRedisClientFactory(cfg *RedisConfig) *redisClientFactory {
	return &redisClientFactory{
		cfg: cfg,
	}
}

func (f *redisClientFactory) CreateClient(dbIndex int) *redis.Client {
	opts := &redis.Options{
		Addr:            f.cfg.RedisURL,
		Password:        f.cfg.RedisPass,
		DB:              dbIndex,
		PoolSize:        f.cfg.RedisPoolSize,
		MinIdleConns:    f.cfg.RedisMinIdleConns,
		MaxRetries:      3,
		ConnMaxIdleTime: 60 * time.Second,
	}

	if f.cfg.RedisTLS {
		opts.TLSConfig = &tls.Config{
			InsecureSkipVerify: true, // #nosec G402
		}
	}

	client := redis.NewClient(opts)

	if dotenv.IsAppEnvProduction() {
		client.AddHook(nrredis.NewHook(opts))
	}

	if err := client.Ping(context.Background()).Err(); err != nil {
		fmt.Printf("[Redis] Failed to connect (DB %d): %v\n", dbIndex, err)
	} else {
		fmt.Printf("[Redis] Connected (DB %d)\n", dbIndex)
	}

	return client
}

func (f *redisClientFactory) CreateClusterClient(addrs []string) *redis.ClusterClient {
	opts := &redis.ClusterOptions{
		Addrs:           addrs,
		Password:        f.cfg.RedisPass,
		PoolSize:        f.cfg.RedisPoolSize,
		MinIdleConns:    f.cfg.RedisMinIdleConns,
		MaxRetries:      3,
		ConnMaxIdleTime: 60 * time.Second,
	}

	if f.cfg.RedisTLS {
		opts.TLSConfig = &tls.Config{
			InsecureSkipVerify: true, // #nosec G402
		}
	}

	client := redis.NewClusterClient(opts)

	if dotenv.IsAppEnvProduction() {
		client.AddHook(nrredis.NewHook(nil))
	}

	if err := client.Ping(context.Background()).Err(); err != nil {
		fmt.Printf("[Redis] Failed to connect (cluster): %v\n", err)
	} else {
		fmt.Printf("[Redis] Connected to cluster: %v\n", addrs)
	}

	return client
}
