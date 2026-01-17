package valueobject_test

import (
	valueobject "doit/urlshortener/internal/domain/value_object"
	"testing"
)

func TestNewUrl(t *testing.T) {
	tests := []struct {
		name        string
		rawUrl      string
		expected    string
		expectError bool
	}{
		{
			name:        "valid http url",
			rawUrl:      "http://example.com",
			expected:    "http://example.com",
			expectError: false,
		},
		{
			name:        "valid https url",
			rawUrl:      "https://example.com/path?query=1",
			expected:    "https://example.com/path?query=1",
			expectError: false,
		},
		{
			name:        "valid https url with port",
			rawUrl:      "https://example.com:8080/path",
			expected:    "https://example.com:8080/path",
			expectError: false,
		},
		{
			name:        "invalid scheme - ftp",
			rawUrl:      "ftp://example.com",
			expected:    "",
			expectError: true,
		},
		{
			name:        "no scheme",
			rawUrl:      "example.com",
			expected:    "",
			expectError: true,
		},
		{
			name:        "relative url",
			rawUrl:      "/path/to/resource",
			expected:    "",
			expectError: true,
		},
		{
			name:        "empty url",
			rawUrl:      "",
			expected:    "",
			expectError: true,
		},
		{
			name:        "invalid url format",
			rawUrl:      "not-a-url",
			expected:    "",
			expectError: true,
		},
		{
			name:        "url with fragment",
			rawUrl:      "https://example.com/page#section",
			expected:    "https://example.com/page%23section",
			expectError: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result, err := valueobject.NewLongUrl(tt.rawUrl)

			if tt.expectError {
				if err == nil {
					t.Errorf("expected error but got none")
				}
				if result != "" {
					t.Errorf("expected empty result on error, got %s", result)
				}
				return
			}

			if err != nil {
				t.Errorf("unexpected error: %v", err)
				return
			}

			if result != tt.expected {
				t.Errorf("expected %s, got %s", tt.expected, result)
			}
		})
	}
}
