package entity

import (
	valueobject "doit/urlshortener/internal/domain/value_object"
	"time"
)

type UrlShorten struct {
	longUrl        string
	shortCode      string
	createdAt      time.Time
	expiredAt      time.Time
	lastAccessedAt time.Time
	clickCount     int
}

func NewUrlShorten(raw, code string, ttl time.Duration, now time.Time) (*UrlShorten, error) {
	validUrl, err := valueobject.NewLongUrl(raw)
	if err != nil {
		return nil, err
	}

	return &UrlShorten{
		longUrl:    validUrl,
		shortCode:  code,
		createdAt:  now,
		expiredAt:  now.Add(ttl),
		clickCount: 0,
	}, nil
}

func (u *UrlShorten) IncreaseClick() {
	u.clickCount++
}

func (u *UrlShorten) SetLastAccessedAt(t time.Time) {
	u.lastAccessedAt = t
}

func (u *UrlShorten) IsExpired(now time.Time) bool {
	return now.After(u.expiredAt)
}

func (u *UrlShorten) LongUrl() string {
	return u.longUrl
}

func (u *UrlShorten) ShortCode() string {
	return u.shortCode
}

func (u *UrlShorten) CreatedAt() time.Time {
	return u.createdAt
}

func (u *UrlShorten) ExpiredAt() time.Time {
	return u.expiredAt
}

func (u *UrlShorten) LastAccessedAt() time.Time {
	return u.lastAccessedAt
}

func (u *UrlShorten) ClickCount() int {
	return u.clickCount
}
