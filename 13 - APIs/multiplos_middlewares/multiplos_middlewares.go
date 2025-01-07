package main

import (
	"fmt"
	"log"
	"net/http"
	"strings"
)

func LoggingMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		log.Printf("Requisition received: %s %s\n", r.Method, r.URL.Path)

		next.ServeHTTP(w, r)
	})
}

func AuthMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		authHeader := r.Header.Get("Authorization")
		if authHeader == "" || !strings.HasPrefix(authHeader, "Bearer ") {
			http.Error(w, "Invalid token", http.StatusUnauthorized)
			return
		}

		token := strings.TrimPrefix(authHeader, "Bearer ")
		if token != "my-secret-token" {
			http.Error(w, "Unauthorized access", http.StatusUnauthorized)
			return
		}

		next.ServeHTTP(w, r)
	})
}

func ProtectedHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "You're in! This is a protected route")
}

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/", ProtectedHandler)

	fmt.Println("Server running on port :8080")
	err := http.ListenAndServe(":8080", AuthMiddleware(LoggingMiddleware(mux)))
	if err != nil {
		fmt.Println("Error when starting server")
	}
}
