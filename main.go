package main

import (
	"log"
	"net/http"
	"sync"
	"urlshortner/controller"
	"urlshortner/utils"
)

// Global variables for storing URL data and metrics
var (
	urlMap  = make(map[string]utils.URLData)
	metrics = make(map[string]int)
	mu      sync.Mutex
)

func main() {
	http.HandleFunc("/shorten", controller.ShortenURL)
	http.HandleFunc("/redirect", controller.RedirectURL)
	http.HandleFunc("/metrics", controller.GetMetrics)
	log.Fatal(http.ListenAndServe(":8080", nil))
}
