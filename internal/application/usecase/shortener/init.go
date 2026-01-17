package shortener

import (
	"context"
	"doit/urlshortener/internal/domain/repository"
	"doit/urlshortener/internal/domain/service"
)

type ShortenerUsecase interface {
	ShortenURL(ctx context.Context, url, ttl string) (string, error)
	RedirectURL(ctx context.Context, shortCode string) (string, error)
}

type shortenerUC struct {
	repo  repository.UrlShortenRepo
	svc   service.ShortenerService
	clock service.Clock
}

func NewShortenerUsecase(
	r repository.UrlShortenRepo,
	svc service.ShortenerService,
	clock service.Clock,
) ShortenerUsecase {
	return &shortenerUC{
		repo:  r,
		svc:   svc,
		clock: clock,
	}
}
