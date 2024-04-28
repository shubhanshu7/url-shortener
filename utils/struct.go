package utils

// URLData will be used for URL Shortening service
type URLData struct {
	OriginalURL string `json:"original_url"`
	ShortURL    string `json:"short_url"`
}

// Metrics represents the metric API
type Metrics struct {
	Domain string `json:"domain"`
	Count  int    `json:"count"`
}
