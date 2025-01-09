package main

import (
	"context"
	"fmt"
	"net/http"
	"time"
)

func longTask(ctx context.Context) {
	select {
	case <-time.After(10 * time.Second):
		fmt.Println("Task finished")
	case <-ctx.Done():
		fmt.Println("Task aborted: ", ctx.Err())
	}
}

func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Println("-------------------")
	ctx, cancel := context.WithTimeout(r.Context(), 5*time.Second)
	defer cancel()

	go longTask(ctx)

	time.Sleep(5 * time.Second)

	select {
	case <-ctx.Done():
		cancel()
		fmt.Fprintf(w, "Too long time. Cancelling task...")
	case <-time.After(2 * time.Second):
		fmt.Fprintln(w, "Task started, waiting finishing")
	}
}

func main() {
	http.HandleFunc("/start-task", handler)

	fmt.Println("Server running on localhost:8080")
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		fmt.Println("Error when starting server: ", err)
	}
}
