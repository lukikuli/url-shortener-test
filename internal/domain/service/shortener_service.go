package service

import "context"

type ShortenerService interface {
	ShortenUrl(ctx context.Context, url string) (string, error)
	ExpandUrl(ctx context.Context, shortCode string) (string, error)
}
