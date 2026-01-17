package service_test

import (
	"context"
	"doit/urlshortener/internal/domain/service"
	"log"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestGenerateShortenUrlCode_Success(t *testing.T) {
	svc := service.NewShortenerService()
	ctx := context.Background()

	code, err := svc.GenerateShortenUrlCode(ctx, "https://example.com")
	if err != nil {
		log.Print("error", err)
	}
	require.NoError(t, err)

	assert.Len(t, code, 7)
}

func TestGenerateShortenUrlCode_ForbiddenCharacters(t *testing.T) {
	var forbiddenChars = map[rune]bool{
		'0': true,
		'O': true,
		'I': true,
		'l': true,
		'1': true,
	}

	svc := service.NewShortenerService()
	ctx := context.Background()

	for i := 0; i < 100; i++ {
		code, err := svc.GenerateShortenUrlCode(ctx, "https://example.com")
		require.NoError(t, err)

		for _, c := range code {
			assert.False(
				t,
				forbiddenChars[c],
				"forbidden character %q found in code %s",
				c,
				code,
			)
		}
	}
}

func TestGenerateShortenUrlCode_StableAcrossCalls(t *testing.T) {
	svc := service.NewShortenerService()
	ctx := context.Background()

	code1, _ := svc.GenerateShortenUrlCode(ctx, "https://example.com")
	code2, _ := svc.GenerateShortenUrlCode(ctx, "https://example.com")

	assert.NotEmpty(t, code1)
	assert.NotEmpty(t, code2)
}
