package shortener_test

import (
	"context"
	"doit/urlshortener/internal/domain/entity"
	"doit/urlshortener/internal/domain/repository"
	"errors"
	"time"
)

type fakeClock struct {
	now time.Time
}

func (f *fakeClock) Now() time.Time {
	return f.now
}

type fakeShortenerService struct {
	codes []string
	call  int
}

func (f *fakeShortenerService) GenerateShortenUrlCode(ctx context.Context, longUrl string) (string, error) {
	code := f.codes[f.call]
	f.call++
	return code, nil
}

type fakeRepo struct {
	data map[string]*entity.UrlShorten
}

func newFakeRepo() *fakeRepo {
	return &fakeRepo{
		data: make(map[string]*entity.UrlShorten),
	}
}

func (f *fakeRepo) SaveShortUrl(ctx context.Context, url *entity.UrlShorten) error {
	if _, exists := f.data[url.ShortCode()]; exists {
		return repository.ErrDuplicateShortCode
	}
	f.data[url.ShortCode()] = url
	return nil
}

func (f *fakeRepo) FindByShortCode(ctx context.Context, code string) (*entity.UrlShorten, error) {
	if v, ok := f.data[code]; ok {
		return v, nil
	}
	return nil, errors.New("not found")
}

func (f *fakeRepo) CheckUrlExists(ctx context.Context, code string) (bool, error) {
	_, ok := f.data[code]
	return ok, nil
}

func (f *fakeRepo) IncrementClick(ctx context.Context, code string, at time.Time) error {
	v, ok := f.data[code]
	if !ok {
		return errors.New("not found")
	}
	v.IncreaseClick()
	v.SetLastAccessedAt(at)
	return nil
}
