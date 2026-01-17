package valueobject

import (
	"fmt"
	"net/url"
)

func NewLongUrl(raw string) (string, error) {
	urlParse, err := url.ParseRequestURI(raw)
	if err != nil {
		return "", err
	}

	if urlParse.Scheme != "http" && urlParse.Scheme != "https" {
		return "", fmt.Errorf("invalid scheme: only http and https are allowed")
	}

	return urlParse.String(), nil
}
