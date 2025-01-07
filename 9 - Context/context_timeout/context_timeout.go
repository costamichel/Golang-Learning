package main

import (
	"context"
	"fmt"
	"time"
)

func task(ctx context.Context) {
	select {
	case <-time.After(2 * time.Second):
		fmt.Println("Task finished!")
	case <-ctx.Done():
		fmt.Println("Task error: ", ctx.Err())
	}
}

func main() {
	ctx, cancel := context.WithTimeout(context.Background(), 1*time.Second)

	defer cancel()
	go task(ctx)

	time.Sleep(3 * time.Second)
	fmt.Println("End!")
}
