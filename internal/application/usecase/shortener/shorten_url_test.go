package shortener_test

import (
	"context"
	"doit/urlshortener/internal/application/usecase/shortener"
	"testing"
	"time"

	"github.com/go-playground/assert/v2"
	"github.com/stretchr/testify/require"
)

func TestShortenURL_Success(t *testing.T) {
	ctx := context.Background()

	repo := newFakeRepo()
	svc := &fakeShortenerService{
		codes: []string{"abc1234"},
	}
	clock := &fakeClock{
		now: time.Date(2026, 1, 1, 0, 0, 0, 0, time.UTC),
	}

	uc := shortener.NewShortenerUsecase(repo, svc, clock)
	code, err := uc.ShortenURL(ctx, "https://example.com", "3600")
	require.NoError(t, err)
	assert.Equal(t, "abc1234", code)
}

func TestShortenURL_RetryOnCollision(t *testing.T) {
	ctx := context.Background()

	repo := newFakeRepo()
	svc := &fakeShortenerService{
		codes: []string{"dup1234", "uniq567"},
	}
	clock := &fakeClock{
		now: time.Now(),
	}

	uc := shortener.NewShortenerUsecase(repo, svc, clock)
	code1, err := uc.ShortenURL(ctx, "https://example.com/1", "3600")
	require.NoError(t, err)
	assert.Equal(t, "dup1234", code1)

	code2, err := uc.ShortenURL(ctx, "https://example.com/2", "3600")
	require.NoError(t, err)
	assert.Equal(t, "uniq567", code2)
}
