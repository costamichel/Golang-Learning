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
		fmt.Println("Task ended!")
	case <-ctx.Done():
		fmt.Println("Task cancelled!")
	}
}

func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Println("-----------------")
	ctx, cancel := context.WithCancel(r.Context())
	defer cancel()

	go longTask(ctx)

	go func() {
		<-ctx.Done()
		cancel()
	}()

	time.Sleep(5 * time.Second) // apenas para poder forçar o fechamento da conexão no insomnia

	fmt.Fprint(w, "Task started, close connection to continue")
}

func main() {
	http.HandleFunc("/", handler)

	fmt.Println("Server running on localhost:8080")
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		fmt.Println("Error when starting server")
	}
}
