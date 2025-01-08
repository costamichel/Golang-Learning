package main

import (
	"fmt"
	"net/http"
	"sync"
	"time"
)

func task(wg *sync.WaitGroup) {
	defer wg.Done()
	time.Sleep(3 * time.Second)
}

func handler(w http.ResponseWriter, r *http.Request) {
	var wg sync.WaitGroup

	wg.Add(3)

	fmt.Fprint(w, "Starting tasks.")
	go task(&wg)
	go task(&wg)
	time.Sleep(1 * time.Second)
	go task(&wg)
	fmt.Fprint(w, "Tasks started.")

	wg.Wait()
	fmt.Fprint(w, "Tasks finished.")
}

func main() {
	http.HandleFunc("/", handler)

	fmt.Println("Server running on localhost:8080")
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		fmt.Println("Error when starting server")
	}
}
