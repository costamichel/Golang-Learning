package main

import (
	"fmt"
	"net/http"
	"time"
)

func longTask() {
	fmt.Println("Long task started...")
	time.Sleep(5 * time.Second)
	fmt.Println("Long task finished!")
}

func handler(w http.ResponseWriter, r *http.Request) {
	go longTask()
	fmt.Fprint(w, "Long task is running")
}

func main() {
	http.HandleFunc("/", handler)

	fmt.Println("Server running on localhost:8080")
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		fmt.Println("Error when starting server")
	}
}
