package main

import (
	"fmt"
	"sync"
)

var contador int
var mutex sync.Mutex

func incrementar(wg *sync.WaitGroup) {
	defer wg.Done()

	mutex.Lock()
	contador++
	mutex.Unlock()
}

func main() {
	var wg sync.WaitGroup

	for i := 0; i < 10000; i++ {
		wg.Add(1)
		go incrementar(&wg)
	}

	wg.Wait()

	fmt.Println("Final value of counter -> ", contador)
	fmt.Println("End!")
}
