package shortener

import (
	"context"
	"errors"
)

func (s *shortenerUC) GetOriginalUrl(ctx context.Context, shortCode string) (string, error) {
	if shortCode == "" {
		return "", errors.New("url not found")
	}

	return "", nil
}
