package service

import (
	"context"
	"crypto/sha256"
	"fmt"
	"math/rand"
	"time"
)

type ShortenerService interface {
	GenerateShortenUrlCode(ctx context.Context, longUrl string) (string, error)
}

type shortenerServiceImpl struct {
}

func NewShortenerService() ShortenerService {
	return &shortenerServiceImpl{}
}

func (service *shortenerServiceImpl) GenerateShortenUrlCode(ctx context.Context, longUrl string) (string, error) {
	raw := fmt.Sprintf("%s:%d:%d", longUrl, time.Now().UnixNano(), rand.Int()) // #nosec G404
	hash := sha256Of(raw)
	encodedStr := generateCodeFromHash(hash)

	return encodedStr[:7], nil
}

func sha256Of(input string) []byte {
	algorithm := sha256.New()
	algorithm.Write([]byte(input))
	return algorithm.Sum(nil)
}

func generateCodeFromHash(hash []byte) string {
	const alphabet = "23456789ABCDEFGHJKLMNPQRSTUVWXYZabcdefghijkmnopqrstuvwxyz"
	const length = 7

	code := make([]byte, length)
	for i := 0; i < length; i++ {
		code[i] = alphabet[int(hash[i])%len(alphabet)]
	}

	return string(code)
}
