package controller

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"
	"urlshortner/utils"
)

func TestShortenURLHandler(t *testing.T) {
	payload := []byte(`{"url":"https://example.com"}`)
	req, err := http.NewRequest("POST", "/shorten", bytes.NewBuffer(payload))
	if err != nil {
		t.Fatal(err)
	}
	req.Header.Set("Content-Type", "application/json")

	rr := httptest.NewRecorder()

	handler := http.HandlerFunc(ShortenURL)
	handler.ServeHTTP(rr, req)

	if status := rr.Code; status != http.StatusCreated {
		t.Errorf("wrong status code: got %v want %v", status, http.StatusCreated)
	}

	var urlData utils.URLData
	err = json.Unmarshal(rr.Body.Bytes(), &urlData)
	if err != nil {
		t.Fatal(err)
	}
	if urlData.OriginalURL != "https://example.com" {
		t.Errorf("unexpected original URL: got %v want %v", urlData.OriginalURL, "https://example.com")
	}
	if urlData.ShortURL == "" {
		t.Errorf(" empty short URL")
	}
}

func TestRedirectURLHandler(t *testing.T) {
	req, err := http.NewRequest("GET", "/redirect?short_url=example_short_url", nil)
	if err != nil {
		t.Fatal(err)
	}

	rr := httptest.NewRecorder()

	handler := http.HandlerFunc(RedirectURL)
	handler.ServeHTTP(rr, req)

	if status := rr.Code; status != http.StatusSeeOther {
		t.Errorf("wrong status code: got %v want %v", status, http.StatusNotFound)
	}

}

func TestGetMetricsHandler(t *testing.T) {
	req, err := http.NewRequest("GET", "/metrics", nil)
	if err != nil {
		t.Fatal(err)
	}

	rr := httptest.NewRecorder()

	handler := http.HandlerFunc(GetMetrics)
	handler.ServeHTTP(rr, req)

	if status := rr.Code; status != http.StatusOK {
		t.Errorf("wrong status code: got %v want %v", status, http.StatusOK)
	}

	var domainMetrics []utils.Metrics
	err = json.Unmarshal(rr.Body.Bytes(), &domainMetrics)
	if err != nil {
		t.Fatal(err)
	}
	if len(domainMetrics) != 0 {
		t.Errorf(" non-empty metrics")
	}
}

func TestGenerateShortURL(t *testing.T) {
	originalURL := "https://example.com"
	shortURL := generateShortURL(originalURL)
	if shortURL == "" {
		t.Error("no shorturl")
	}
}
