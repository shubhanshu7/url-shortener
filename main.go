package main

import (
	"log"
	"net/http"
	"urlshortner/controller"
)

func main() {
	http.HandleFunc("/shorten", controller.ShortenURL)
	http.HandleFunc("/redirect", controller.RedirectURL)
	http.HandleFunc("/metrics", controller.GetMetrics)
	log.Fatal(http.ListenAndServe(":8080", nil))
}
