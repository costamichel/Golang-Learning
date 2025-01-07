package main

import (
	"context"
	"fmt"
	"sync"
	"time"
)

func printNumbers(ctx context.Context, prefix string, wg *sync.WaitGroup) {
	defer wg.Done()

	for i := 0; ; i++ {
		select {
		case <-ctx.Done():
			fmt.Println("Cancelling task with prefix -> ", prefix)
			return
		default:
			fmt.Printf("%s %d\n", prefix, i)
			time.Sleep(500 * time.Millisecond)
		}
	}
}

func main() {
	ctx, cancel := context.WithCancel(context.Background())
	var wg sync.WaitGroup

	for i := 1; i <= 5; i++ {
		prefix := fmt.Sprintf("TASK_%d", i)
		wg.Add(1)
		go printNumbers(ctx, prefix, &wg)
	}

	time.Sleep(4 * time.Second)
	fmt.Println("Tasks will be cancelled!")
	cancel()
	time.Sleep(1 * time.Second)

	wg.Wait()
	fmt.Println("Tasks ended!")

	fmt.Println("End!")
}
