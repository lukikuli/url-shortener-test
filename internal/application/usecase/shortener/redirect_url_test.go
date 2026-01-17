package shortener_test

import (
	"context"
	"doit/urlshortener/internal/application/usecase/shortener"
	"doit/urlshortener/internal/domain/entity"
	"testing"
	"time"

	"github.com/go-playground/assert/v2"
	"github.com/stretchr/testify/require"
)

func TestRedirectURL_Success(t *testing.T) {
	ctx := context.Background()

	repo := newFakeRepo()
	svc := &fakeShortenerService{
		codes: []string{"dup1234", "uniq567"},
	}

	now := time.Date(2026, 1, 1, 12, 0, 0, 0, time.UTC)
	clock := &fakeClock{now: now}

	urlEntity, _ := entity.NewUrlShorten(
		"https://example.com",
		"redir123",
		24*time.Hour,
		now.Add(-time.Hour),
	)
	repo.SaveShortUrl(ctx, urlEntity)

	uc := shortener.NewShortenerUsecase(repo, svc, clock)
	longURL, err := uc.RedirectURL(ctx, "redir123")
	require.NoError(t, err)

	assert.Equal(t, "https://example.com", longURL)
	assert.Equal(t, 1, urlEntity.ClickCount())
	assert.Equal(t, now, urlEntity.LastAccessedAt())
}

func TestRedirectURL_Expired(t *testing.T) {
	ctx := context.Background()

	repo := newFakeRepo()
	svc := &fakeShortenerService{
		codes: []string{"dup1234", "uniq567"},
	}
	now := time.Date(2026, 1, 2, 0, 0, 0, 0, time.UTC)
	clock := &fakeClock{now: now}

	urlEntity, _ := entity.NewUrlShorten(
		"https://example.com",
		"expired123",
		1*time.Hour,
		now.Add(-2*time.Hour),
	)
	repo.SaveShortUrl(ctx, urlEntity)

	uc := shortener.NewShortenerUsecase(repo, svc, clock)
	_, err := uc.RedirectURL(ctx, "expired123")
	require.Error(t, err)
}
