package main

import (
	"fmt"
	"go-limiter/internal/middleware"
	"net/http"
)

func handleRequest(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Status: OK")
}

func main() {
	ipRateLimiter := middleware.NewIPRateLimiter()

	mux := http.NewServeMux()
	mux.HandleFunc("/", middleware.RateLimitMiddleware(ipRateLimiter, handleRequest))

	fmt.Println("Server starting on :8080")
	http.ListenAndServe(":8080", mux)
}
