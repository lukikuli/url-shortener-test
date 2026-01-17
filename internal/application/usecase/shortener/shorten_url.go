package shortener

import (
	"context"
	"doit/urlshortener/internal/domain/entity"
	"doit/urlshortener/internal/domain/repository"
	"errors"
	"strconv"
	"time"
)

const DEFAULT_TTL_SECONDS = 86400

func (s *shortenerUC) ShortenURL(ctx context.Context, url, ttl string) (string, error) {
	ttlTime, err := strconv.Atoi(ttl)
	if err != nil || ttlTime <= 0 {
		ttlTime = DEFAULT_TTL_SECONDS
	}

	ttlDuration := time.Duration(ttlTime) * time.Second

	maxAttempt := 3
	for i := 0; i < maxAttempt; i++ {
		shortCode, err := s.svc.GenerateShortenUrlCode(ctx, url)
		if err != nil {
			return "", err
		}

		now := s.clock.Now()
		urlEntity, err := entity.NewUrlShorten(url, shortCode, ttlDuration, now)
		if err != nil {
			return "", err
		}

		err = s.repo.SaveShortUrl(ctx, urlEntity)
		if err == nil {
			return shortCode, nil
		}

		if !errors.Is(err, repository.ErrDuplicateShortCode) {
			return "", err
		}
	}

	return "", errors.New("failed to generate unique short code")
}
