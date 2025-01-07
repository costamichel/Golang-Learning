package main

import (
	"fmt"
	"sync"
	"time"
)

func task(id int, wg *sync.WaitGroup) {
	defer wg.Done()
	fmt.Printf("Task %d started...\n", id)
	time.Sleep(40 * time.Millisecond)
	fmt.Printf("... task %d done!\n", id)
}

func main() {
	var wg sync.WaitGroup

	for i := 0; i < 3; i++ {
		wg.Add(1)
		go task(i, &wg)
	}
	wg.Wait()

	fmt.Println("End!")
}
