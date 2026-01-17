package shortener

import (
	"context"
	"errors"
)

func (s *shortenerUC) RedirectURL(ctx context.Context, shortCode string) (string, error) {
	urlEntity, err := s.repo.FindByShortCode(ctx, shortCode)
	if err != nil {
		return "", err
	}

	now := s.clock.Now()
	if urlEntity.IsExpired(now) {
		return "", errors.New("url expired")
	}

	err = s.repo.IncrementClick(ctx, urlEntity.ShortCode(), now)
	if err != nil {
		return "", err
	}

	return urlEntity.LongUrl(), nil
}
