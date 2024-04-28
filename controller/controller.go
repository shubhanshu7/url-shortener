package controller

import "net/http"

// shortenURL generates a shortened URL for the given original URL
func ShortenURL(w http.ResponseWriter, r *http.Request) {
}

// redirectURL redirects the user to the original URL based on the shortened URL
func RedirectURL(w http.ResponseWriter, r *http.Request) {
}

// getMetrics returns the top 3 domains with the most shortened URLs
func GetMetrics(w http.ResponseWriter, r *http.Request) {
}
