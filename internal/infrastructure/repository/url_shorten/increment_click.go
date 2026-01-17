package url_shorten

import (
	"context"
	"errors"
	"time"
)

func (u *urlShortengRepo) IncrementClick(
	ctx context.Context,
	shortCode string,
	accessedAt time.Time,
) error {
	u.mu.Lock()
	defer u.mu.Unlock()

	if _, ok := u.data[shortCode]; !ok {
		return errors.New("data not found")
	}

	u.data[shortCode].IncreaseClick()
	u.data[shortCode].SetLastAccessedAt(accessedAt)

	return nil
}
