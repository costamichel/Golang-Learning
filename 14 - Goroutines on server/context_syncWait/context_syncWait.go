package main

import (
	"context"
	"fmt"
	"net/http"
	"sync"
	"time"
)

func longTask(ctx context.Context, wg *sync.WaitGroup, id int) {
	defer wg.Done()

	select {
	case <-time.After(5 * time.Second):
		fmt.Printf("Task %d finished\n", id)
	case <-ctx.Done():
		fmt.Printf("Task %d aborted: %s\n", id, ctx.Err())
	}
}

func handler(w http.ResponseWriter, r *http.Request) {
	ctx, cancel := context.WithTimeout(r.Context(), 10*time.Second)
	defer cancel()

	numTasks := 3
	var wg sync.WaitGroup
	for i := 1; i <= numTasks; i++ {
		wg.Add(1)
		go longTask(ctx, &wg, i)
	}

	wg.Wait()

	fmt.Fprintf(w, "Tasks has been executed")
}

func main() {
	http.HandleFunc("/", handler)

	fmt.Println("Server running on localhost:8080")
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		fmt.Println("Error when starting server")
	}
}
