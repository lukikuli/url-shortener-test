package entity

import (
	valueobject "doit/urlshortener/internal/domain/value_object"
	"errors"
	"time"
)

type Clock interface {
	Now() time.Time
}

type UrlShorten struct {
	longUrl        string
	shortCode      string
	createdAt      time.Time
	expiredAt      time.Time
	lastAccessedAt time.Time
	clickCount     int
}

func NewUrlShorten(raw string, ttl time.Duration, clock Clock) (*UrlShorten, error) {
	validUrl, err := valueobject.NewLongUrl(raw)
	if err != nil {
		return nil, err
	}

	now := clock.Now()

	return &UrlShorten{
		longUrl:   validUrl,
		createdAt: now,
	}, nil
}

func (u *UrlShorten) SetShortCode(coded string) {
	u.shortCode = coded
}

func (u *UrlShorten) SetClick(count int) {
	u.clickCount = count
}

func (u *UrlShorten) IncreaseClick(clock Clock) error {
	if u.IsExpired(clock) {
		return errors.New("expired")
	}

	u.clickCount++
	u.lastAccessedAt = clock.Now()

	return nil
}

func (u *UrlShorten) LongUrl() string {
	return u.longUrl
}

func (u *UrlShorten) ShortCode() string {
	return u.shortCode
}

func (u *UrlShorten) ClickCount() int {
	return u.clickCount
}

func (u *UrlShorten) CreatedAt() time.Time {
	return u.createdAt
}

func (u *UrlShorten) IsExpired(clock Clock) bool {
	return clock.Now().After(u.expiredAt)
}
