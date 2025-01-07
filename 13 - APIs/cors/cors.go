package main

import (
	"fmt"
	"net/http"
)

func cors(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Header().Set("Access-Control-Allow-Methods", "GET, POST, DELETE, PUT, OPTIONS")
		w.Header().Set("Access-Control-Allow-Headers", "Content-Type, X-Auth-Token")

		if r.Method == "OPTIONS" {
			w.WriteHeader(http.StatusOK)
			return
		}

		next.ServeHTTP(w, r)
	})
}

func helloHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Hello, you've made a allowed cors requisition")
}

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/", helloHandler)

	fmt.Println("Server running on port 8080")
	err := http.ListenAndServe(":8080", cors(mux))
	if err != nil {
		fmt.Println("Error when starting server")
	}
}
