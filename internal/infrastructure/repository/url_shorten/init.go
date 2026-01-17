package url_shorten

import (
	"doit/urlshortener/internal/domain/entity"
	"doit/urlshortener/internal/domain/repository"
	"sync"
)

type urlShortengRepo struct {
	// redis *redis.Client
	// sqlDB *sql.DB
	mu   sync.Mutex
	data map[string]*entity.UrlShorten
}

func NewUrlMappingRepository() repository.UrlShortenRepo {
	return &urlShortengRepo{
		data: make(map[string]*entity.UrlShorten),
	}
}
