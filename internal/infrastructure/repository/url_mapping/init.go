package urlmapping

import (
	"database/sql"

	"github.com/redis/go-redis/v9"
)

type urlMappingRepo struct {
	redis *redis.Client
	sqlDB *sql.DB
}

// func NewUrlMappingRepository(redis *redis.Client, sqlDB *sql.DB) repository.UrlMappingRepo {
// 	return &urlMappingRepo{
// 		redis: redis,
// 		sqlDB: sqlDB,
// 	}
// }
