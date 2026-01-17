package shortener

import "context"

func (s *shortenerUC) GetStats(ctx context.Context, shortCode string) (UrlStatsDTO, error) {
	urlEntity, err := s.repo.FindByShortCode(ctx, shortCode)
	if err != nil {
		return UrlStatsDTO{}, err
	}

	return UrlStatsDTO{
		LongURL:        urlEntity.LongUrl(),
		CreatedAt:      urlEntity.CreatedAt(),
		ExpiresAt:      urlEntity.ExpiredAt(),
		ClickCount:     urlEntity.ClickCount(),
		LastAccessedAt: urlEntity.LastAccessedAt(),
	}, nil
}
