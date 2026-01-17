package shortener

import (
	"context"
	"doit/urlshortener/internal/domain/repository"
	"doit/urlshortener/internal/domain/service"
)

type ShortenerUsecase interface {
	CreateShortUrl(ctx context.Context, url string) (string, error)
	GetOriginalUrl(ctx context.Context, shortCode string) (string, error)
}

type shortenerUC struct {
	repo repository.UrlMappingRepo
	svc  service.ShortenerService
}

func NewShortenerUsecase(
	r repository.UrlMappingRepo,
	svc service.ShortenerService) ShortenerUsecase {
	return &shortenerUC{
		repo: r,
		svc:  svc,
	}
}
