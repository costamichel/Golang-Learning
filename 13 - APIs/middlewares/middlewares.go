package main

import (
	"fmt"
	"net/http"
	"time"
)

func logger(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		start := time.Now()
		fmt.Printf("Método: %s | URL: %s | Hora: %s\n", r.Method, r.URL.Path, start.Format(time.RFC3339))
		next.ServeHTTP(w, r)
	})
}

func helloHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Hello, middleware!")
}

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/", helloHandler)

	fmt.Println("Server running on port 8080")
	err := http.ListenAndServe(":8080", logger(mux))
	if err != nil {
		fmt.Println("Error when starting server")
	}
}
