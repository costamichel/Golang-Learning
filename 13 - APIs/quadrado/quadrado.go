package main

import (
	"fmt"
	"math"
	"net/http"
	"strconv"
	"strings"
)

func squareHandler(w http.ResponseWriter, r *http.Request) {
	parts := strings.Split(r.URL.Path, "/")
	if len(parts) != 3 || parts[2] == "" {
		http.Error(w, "No argument passed", http.StatusBadRequest)
		return
	}

	number, err := strconv.ParseFloat(parts[2], 64)
	if err != nil {
		http.Error(w, "Argument is not a valid number", http.StatusBadRequest)
		return
	}

	square := math.Pow(number, 2)
	fmt.Fprintf(w, "Result: %.2f", square)
}

func main() {
	http.HandleFunc("/square/", squareHandler)

	fmt.Println("Server running on localhost:8080")
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		fmt.Println("Error when starting server")
	}
}
