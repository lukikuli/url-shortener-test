package url_shorten

import (
	"context"
	"doit/urlshortener/internal/domain/entity"
	"doit/urlshortener/internal/domain/repository"
)

func (u *urlShortengRepo) SaveShortUrl(ctx context.Context, urlm *entity.UrlShorten) error {
	u.mu.Lock()
	defer u.mu.Unlock()

	if _, exists := u.data[urlm.ShortCode()]; exists {
		return repository.ErrDuplicateShortCode
	}

	u.data[urlm.ShortCode()] = urlm

	return nil
}
