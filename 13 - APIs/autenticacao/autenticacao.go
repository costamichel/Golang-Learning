package main

import (
	"fmt"
	"net/http"
)

func authenticate(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		token := r.Header.Get("X-Auth-Token")

		if token != "mysecret" {
			http.Error(w, "Unauthorized access!", http.StatusUnauthorized)
			return
		}

		next.ServeHTTP(w, r)
	})
}

func helloHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello, you're authenticated! Welcome!")
}

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/", helloHandler)

	fmt.Println("Server running on port 8080")
	err := http.ListenAndServe(":8080", authenticate(mux))
	if err != nil {
		fmt.Println("Error when starting server")
	}
}
