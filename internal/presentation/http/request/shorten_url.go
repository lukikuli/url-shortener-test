package request

type RequestShortenUrl struct {
	LongUrl    string `json:"long_url"`
	TTLSeconds string `json:"ttl_seconds,omitempty"`
}
