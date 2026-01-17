package entity_test

import (
	"doit/urlshortener/internal/domain/entity"
	"testing"
	"time"
)

func TestNewUrlShorten(t *testing.T) {
	fixedTime := time.Date(2024, 1, 1, 12, 0, 0, 0, time.UTC)
	ttl := 24 * time.Hour

	tests := []struct {
		name        string
		rawUrl      string
		shortCode   string
		expectError bool
	}{
		{
			name:        "valid http url",
			rawUrl:      "http://example.com",
			shortCode:   "abc123",
			expectError: false,
		},
		{
			name:        "valid https url",
			rawUrl:      "https://example.com/path",
			shortCode:   "xyz789",
			expectError: false,
		},
		{
			name:        "invalid url",
			rawUrl:      "invalid-url",
			shortCode:   "abc123",
			expectError: true,
		},
		{
			name:        "empty url",
			rawUrl:      "",
			shortCode:   "abc123",
			expectError: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			urlShorten, err := entity.NewUrlShorten(tt.rawUrl, tt.shortCode, ttl, fixedTime)

			if tt.expectError {
				if err == nil {
					t.Errorf("expected error but got none")
				}
				return
			}

			if err != nil {
				t.Errorf("unexpected error: %v", err)
				return
			}

			if urlShorten.LongUrl() != tt.rawUrl {
				t.Errorf("expected longUrl %s, got %s", tt.rawUrl, urlShorten.LongUrl())
			}

			if urlShorten.ShortCode() != tt.shortCode {
				t.Errorf("expected shortCode %s, got %s", tt.shortCode, urlShorten.ShortCode())
			}

			if urlShorten.ClickCount() != 0 {
				t.Errorf("expected clickCount 0, got %d", urlShorten.ClickCount())
			}

			if !urlShorten.CreatedAt().Equal(fixedTime) {
				t.Errorf("expected createdAt %v, got %v", fixedTime, urlShorten.CreatedAt())
			}
		})
	}
}

func TestUrlShorten_IncreaseClick(t *testing.T) {
	now := time.Now()
	urlShorten, _ := entity.NewUrlShorten("https://example.com", "abc123", time.Hour, now)

	urlShorten.IncreaseClick()

	if urlShorten.ClickCount() != 1 {
		t.Errorf("expected clickCount 1, got %d", urlShorten.ClickCount())
	}

	urlShorten.IncreaseClick()

	if urlShorten.ClickCount() != 2 {
		t.Errorf("expected clickCount 2, got %d", urlShorten.ClickCount())
	}
}

func TestUrlShorten_SetLastAccessedAt(t *testing.T) {
	now := time.Now()
	urlShorten, _ := entity.NewUrlShorten("https://example.com", "abc123", time.Hour, now)
	accessTime := time.Date(2024, 1, 2, 10, 0, 0, 0, time.UTC)

	// Initially should be zero
	if !urlShorten.LastAccessedAt().IsZero() {
		t.Error("lastAccessedAt should be zero initially")
	}

	urlShorten.SetLastAccessedAt(accessTime)

	if !urlShorten.LastAccessedAt().Equal(accessTime) {
		t.Errorf("expected lastAccessedAt %v, got %v", accessTime, urlShorten.LastAccessedAt())
	}
}

func TestUrlShorten_ExpiredAt(t *testing.T) {
	fixedTime := time.Date(2024, 1, 1, 12, 0, 0, 0, time.UTC)
	ttl := 2 * time.Hour
	expectedExpiry := fixedTime.Add(ttl)

	urlShorten, _ := entity.NewUrlShorten("https://example.com", "abc123", ttl, fixedTime)

	if !urlShorten.ExpiredAt().Equal(expectedExpiry) {
		t.Errorf("expected expiredAt %v, got %v", expectedExpiry, urlShorten.ExpiredAt())
	}
}

func TestUrlShorten_IsExpired(t *testing.T) {
	fixedTime := time.Date(2024, 1, 1, 12, 0, 0, 0, time.UTC)
	ttl := time.Hour

	urlShorten, _ := entity.NewUrlShorten("https://example.com", "abc123", ttl, fixedTime)

	// Not expired yet
	if urlShorten.IsExpired(fixedTime) {
		t.Error("URL should not be expired yet")
	}

	// Check expiration after TTL
	futureTime := fixedTime.Add(2 * time.Hour)
	if !urlShorten.IsExpired(futureTime) {
		t.Error("URL should be expired")
	}
}