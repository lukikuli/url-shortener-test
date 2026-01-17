package url_shorten

import (
	"context"
	"doit/urlshortener/internal/domain/entity"
	"errors"
)

func (u *urlShortengRepo) FindByShortCode(ctx context.Context, shortCode string) (*entity.UrlShorten, error) {
	u.mu.Lock()
	defer u.mu.Unlock()

	result, ok := u.data[shortCode]
	if !ok {
		return nil, errors.New("not found")
	}

	return result, nil
}
