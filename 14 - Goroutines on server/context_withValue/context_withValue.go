package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"time"
)

type key string

const reqIDKey key = "requestID"

func generateRequestID() string {
	return fmt.Sprintf("%d", time.Now().UnixNano())
}

// Middleware
func withRequestID(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		reqID := generateRequestID()
		ctx := context.WithValue(r.Context(), reqIDKey, reqID)
		next.ServeHTTP(w, r.WithContext(ctx))
	})
}

func handler(w http.ResponseWriter, r *http.Request) {
	reqID := r.Context().Value(reqIDKey).(string)

	time.Sleep(2 * time.Second)

	fmt.Fprintf(w, "ID da requisição: %s", reqID)

	log.Printf("Requisição atendida: %s\n", reqID)
}

func main() {
	http.Handle("/", withRequestID(http.HandlerFunc(handler)))

	fmt.Println("Server running on localhost:8080")
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		fmt.Println("Error when starting server")
	}
}
