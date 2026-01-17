package repository

import (
	"context"
	"doit/urlshortener/internal/domain/entity"
)

type UrlMappingRepo interface {
	SaveUrlMapping(ctx context.Context, urlm *entity.UrlShorten) error
	GetUrlMapping(ctx context.Context, shortUrl string) (*entity.UrlShorten, error)
	CheckUrlExists(ctx context.Context, shortUrl string) (bool, error)

	SaveShortUrlToCache(ctx context.Context, shortUrl string) error
	GetShortUrlFromCache(ctx context.Context, longUrl string) (string, error)
}
