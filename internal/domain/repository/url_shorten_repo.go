package repository

import (
	"context"
	"doit/urlshortener/internal/domain/entity"
	"errors"
	"time"
)

var ErrDuplicateShortCode = errors.New("duplicate short code")

type UrlShortenRepo interface {
	SaveShortUrl(ctx context.Context, urlm *entity.UrlShorten) error
	FindByShortCode(ctx context.Context, shortCode string) (*entity.UrlShorten, error)
	CheckUrlExists(ctx context.Context, shortCode string) (bool, error)
	IncrementClick(ctx context.Context, shortCode string, accessedAt time.Time) error
}
