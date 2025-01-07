package main

import (
	"fmt"
	"net/http"
	"strings"
)

func greetHandler(w http.ResponseWriter, r *http.Request) {
	parts := strings.Split(r.URL.Path, "/")
	if len(parts) < 3 || parts[2] == "" {
		http.Error(w, "Name not passed", http.StatusBadRequest)
		return
	}

	name := parts[2]
	fmt.Fprintf(w, "Hello %s", name)
}

func main() {
	http.HandleFunc("/greet/", greetHandler)

	fmt.Println("Server running on localhost:8080")
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		fmt.Println("Error when starting server")
	}
}
