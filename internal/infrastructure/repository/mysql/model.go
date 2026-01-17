package mysql

import "time"

type ShortUrlMapping struct {
	LongUrl        string    `dbq:"long_url"`
	ShortCode      string    `dbq:"short_code"`
	ClickCount     int       `dbq:"click_count"`
	ExpiredAt      time.Time `dbq:"expired_at"`
	CreatedAt      time.Time `dbq:"created_at"`
	CreatedBy      string    `dbq:"created_by"`
	lastAccessedAt time.Time `dbq:"last_accessed_at"`
}

func (ShortUrlMapping) TableName() string {
	return "short_url_mapping"
}
