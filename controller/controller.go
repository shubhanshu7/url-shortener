package controller

import (
	"crypto/sha256"
	"encoding/base64"
	"encoding/json"
	"net/http"
	"strings"
	"sync"
	"urlshortner/utils"
)

// Global variables for storing URL data and metrics
var (
	urlMap  = make(map[string]utils.URLData)
	metrics = make(map[string]int)
	mu      sync.Mutex
)

// shortenURL generates a shortened URL for the given original URL
func ShortenURL(w http.ResponseWriter, r *http.Request) {
	if r.Method != "POST" {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	decoder := json.NewDecoder(r.Body)
	var requestData struct {
		URL string `json:"url"`
	}
	err := decoder.Decode(&requestData)
	if err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	originalURL := requestData.URL
	if originalURL == "" {
		http.Error(w, "URL is required", http.StatusBadRequest)
		return
	}

	mu.Lock()
	defer mu.Unlock()

	shortURL, exists := urlMap[originalURL]
	if exists {
		// Increment count for the domain
		metrics[extractDomain(originalURL)]++
		json.NewEncoder(w).Encode(shortURL)
		return
	}

	shortURL = utils.URLData{
		OriginalURL: originalURL,
		ShortURL:    generateShortURL(originalURL),
	}
	urlMap[originalURL] = shortURL

	// Increment count for the domain
	metrics[extractDomain(originalURL)]++

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(shortURL)
}

// redirectURL redirects the user to the original URL based on the shortened URL
func RedirectURL(w http.ResponseWriter, r *http.Request) {
	if r.Method != "GET" {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	shortURL := r.URL.Query().Get("short_url")
	if shortURL == "" {
		http.Error(w, "Short URL is required", http.StatusBadRequest)
		return
	}

	mu.Lock()
	defer mu.Unlock()

	originalURL, exists := getOriginalURL(shortURL)
	if !exists {
		http.Error(w, "Short URL not found", http.StatusNotFound)
		return
	}

	http.Redirect(w, r, originalURL, http.StatusSeeOther)
}

// getMetrics returns the top 3 domains with the most shortened URLs
func GetMetrics(w http.ResponseWriter, r *http.Request) {
	if r.Method != "GET" {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	mu.Lock()
	defer mu.Unlock()

	var domainMetrics []utils.Metrics
	for domain, count := range metrics {
		domainMetrics = append(domainMetrics, utils.Metrics{Domain: domain, Count: count})
	}

	// Sorting the metrics by count in descending order
	sortMetricsByCount(domainMetrics)

	var top3Metrics []utils.Metrics
	if len(domainMetrics) > 3 {
		top3Metrics = domainMetrics[:3]
	} else {
		top3Metrics = domainMetrics
	}

	json.NewEncoder(w).Encode(top3Metrics)
}

// generateShortURL generates a short URL for the given original URL using hashing(didn't use bitly.com , just tried something new)
func generateShortURL(originalURL string) string {
	// Create a new SHA-256 hash
	hash := sha256.New()

	hash.Write([]byte(originalURL))

	hashSum := hash.Sum(nil)

	shortURL := base64.URLEncoding.EncodeToString(hashSum)

	return shortURL
}

// getOriginalURL retrieves the original URL based on the shortened URL
func getOriginalURL(shortURL string) (string, bool) {
	for _, urlData := range urlMap {
		if urlData.ShortURL == shortURL {
			return urlData.OriginalURL, true
		}
	}
	return "", false
}

// extractDomain extracts the domain from the given URL
func extractDomain(url string) string {
	parts := strings.Split(url, "//")
	if len(parts) > 1 {
		domain := strings.Split(parts[1], "/")
		if len(domain) > 0 {
			return domain[0]
		}
	}
	return ""
}

// sortMetricsByCount sorts the metrics by count in descending order
func sortMetricsByCount(metrics []utils.Metrics) {
	for i := range metrics {
		maxIndex := i
		for j := i + 1; j < len(metrics); j++ {
			if metrics[j].Count > metrics[maxIndex].Count {
				maxIndex = j
			}
		}
		metrics[i], metrics[maxIndex] = metrics[maxIndex], metrics[i]
	}
}
