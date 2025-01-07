package main

import (
	"fmt"
	"net/http"
)

func helloHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello, Go HTTP!")
}

func main() {
	http.HandleFunc("/", helloHandler)

	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		fmt.Println("Error when starting server: ", err)
		return
	}
	fmt.Println("Server listening at port 8080")
}
