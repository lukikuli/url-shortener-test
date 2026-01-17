package url_shorten

import (
	"context"
	"errors"
)

func (u *urlShortengRepo) CheckUrlExists(ctx context.Context, shortCode string) (bool, error) {
	u.mu.Lock()
	defer u.mu.Unlock()

	if _, ok := u.data[shortCode]; !ok {
		return false, errors.New("not found")
	}

	return true, nil
}
