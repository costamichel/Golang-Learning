package main

import (
	"fmt"
	"net/http"
)

func handler(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case "GET":
		fmt.Fprint(w, "You'd made a GET requisition")
	case "POST":
		fmt.Fprintf(w, "You'd made a POST requisition")
	case "OPTIONS":
		fmt.Fprintf(w, "You'd made a OPTIONS requisition")
	default:
		http.Error(w, "This method is not allowed", http.StatusMethodNotAllowed)
	}
}

func main() {
	http.HandleFunc("/", handler)

	fmt.Println("Server running on localhost:8080")
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		fmt.Println("Error on server: ", err)
	}
}
