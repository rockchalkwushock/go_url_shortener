package internal

import (
	"net/url"
	"time"
)

type URLMap struct {
	ClickStats  uint32     // json: "click_stats"
	CreatedAt   time.Time  // json: "created_at"
	ExpiredAt   *time.Time // json: "expired_at"
	OriginalURL string     // json: "original_url"
	ShortenedID string     // json: "shortened_id"
}

func IsValidURL(u string) bool {
	parsedURL, err := url.Parse(u)

	if err != nil {
		return false
	}

	// Verify that the URL has a scheme and a host
	return parsedURL.Scheme != "" && parsedURL.Host != ""
}
