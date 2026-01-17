package url_shorten_test

import (
	"context"
	"doit/urlshortener/internal/domain/entity"
	"doit/urlshortener/internal/domain/repository"
	"doit/urlshortener/internal/infrastructure/repository/url_shorten"
	"sync"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestSaveAndFindByShortCode(t *testing.T) {
	repo := url_shorten.NewUrlMappingRepository()
	ctx := context.Background()

	now := time.Now()
	urlEntity, err := entity.NewUrlShorten(
		"https://example.com",
		"abc123",
		24*time.Hour,
		now,
	)
	require.NoError(t, err)

	err = repo.SaveShortUrl(ctx, urlEntity)
	require.NoError(t, err)

	found, err := repo.FindByShortCode(ctx, "abc123")
	require.NoError(t, err)

	assert.Equal(t, "https://example.com", found.LongUrl())
	assert.Equal(t, "abc123", found.ShortCode())
	assert.Equal(t, 0, found.ClickCount())
}

func TestSaveShortUrl_DuplicateShortCode(t *testing.T) {
	repo := url_shorten.NewUrlMappingRepository()
	ctx := context.Background()
	now := time.Now()

	first, _ := entity.NewUrlShorten(
		"https://example.com/1",
		"dup123",
		24*time.Hour,
		now,
	)

	second, _ := entity.NewUrlShorten(
		"https://example.com/2",
		"dup123",
		24*time.Hour,
		now,
	)

	err := repo.SaveShortUrl(ctx, first)
	require.NoError(t, err)

	err = repo.SaveShortUrl(ctx, second)
	require.Error(t, err)
	assert.Equal(t, err, repository.ErrDuplicateShortCode)
}

func TestIncrementClick_ConcurrentSafety(t *testing.T) {
	repo := url_shorten.NewUrlMappingRepository()
	ctx := context.Background()

	now := time.Now()
	urlEntity, err := entity.NewUrlShorten(
		"https://example.com",
		"concurrent123",
		24*time.Hour,
		now,
	)
	require.NoError(t, err)

	err = repo.SaveShortUrl(ctx, urlEntity)
	require.NoError(t, err)

	const workers = 200
	var wg sync.WaitGroup
	wg.Add(workers)

	for i := 0; i < workers; i++ {
		go func() {
			defer wg.Done()
			err := repo.IncrementClick(ctx, "concurrent123", now)
			require.NoError(t, err)
		}()
	}

	wg.Wait()

	found, err := repo.FindByShortCode(ctx, "concurrent123")
	require.NoError(t, err)

	assert.Equal(t, workers, found.ClickCount())
	assert.Equal(t, now, found.LastAccessedAt())
}
