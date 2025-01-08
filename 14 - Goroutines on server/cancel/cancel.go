package main

import (
	"context"
	"fmt"
	"net/http"
	"time"
)

func longTask(ctx context.Context) {
	select {
	case <-time.After(5 * time.Second):
		fmt.Println("Task finished")
	case <-ctx.Done():
		fmt.Println("Task cancelled due connection closing")
	}
}

func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Println("---------------------------")
	ctx, cancel := context.WithCancel(r.Context())
	defer cancel()

	go longTask(ctx)

	select {
	case <-time.After(10 * time.Second):
		fmt.Fprintf(w, "Task done, returning after 10 seconds")
	case <-r.Context().Done():
		fmt.Fprintf(w, "Connection closed by client")
		cancel()
	}
}

func main() {
	http.HandleFunc("/", handler)
	fmt.Println("Server running on localhost:8080")
	http.ListenAndServe(":8080", nil)
}
