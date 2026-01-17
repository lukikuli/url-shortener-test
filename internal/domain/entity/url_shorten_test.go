package entity_test

import (
	"doit/urlshortener/internal/domain/entity"
	"testing"
	"time"
)

func TestNewUrlShorten(t *testing.T) {
	tests := []struct {
		name        string
		rawUrl      string
		clickCount  int
		expectError bool
	}{
		{
			name:        "valid http url",
			rawUrl:      "http://example.com",
			clickCount:  0,
			expectError: false,
		},
		{
			name:        "valid https url",
			rawUrl:      "https://example.com/path",
			clickCount:  5,
			expectError: false,
		},
		{
			name:        "invalid url",
			rawUrl:      "invalid-url",
			clickCount:  0,
			expectError: true,
		},
		{
			name:        "empty url",
			rawUrl:      "",
			clickCount:  0,
			expectError: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			urlMapping, err := entity.NewUrlShorten(tt.rawUrl)

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

			if urlMapping.LongUrl() != tt.rawUrl {
				t.Errorf("expected longUrl %s, got %s", tt.rawUrl, urlMapping.LongUrl())
			}

			if urlMapping.ClickCount() != tt.clickCount {
				t.Errorf("expected clickCount %d, got %d", tt.clickCount, urlMapping.ClickCount())
			}

			if urlMapping.CreatedAt().IsZero() {
				t.Error("createdAt should not be zero")
			}

			if urlMapping.ShortCode() != "" {
				t.Error("shortCode should be empty initially")
			}
		})
	}
}

func TestUrlMapping_SeturlCode(t *testing.T) {
	urlMapping, _ := entity.NewUrlShorten("https://example.com")
	code := "abc123"

	urlMapping.SetShortCode(code)

	if urlMapping.ShortCode() != code {
		t.Errorf("expected urlCode %s, got %s", code, urlMapping.ShortCode())
	}
}

func TestUrlMapping_IncreaseClick(t *testing.T) {
	urlMapping, _ := entity.NewUrlShorten("https://example.com")

	urlMapping.IncreaseClick()

	if urlMapping.ClickCount() != 6 {
		t.Errorf("expected clickCount 6, got %d", urlMapping.ClickCount())
	}
}

func TestUrlMapping_CreatedAt(t *testing.T) {
	before := time.Now()
	urlMapping, _ := entity.NewUrlShorten("https://example.com")
	after := time.Now()

	createdAt := urlMapping.CreatedAt()
	if createdAt.Before(before) || createdAt.After(after) {
		t.Error("createdAt should be between before and after timestamps")
	}
}
