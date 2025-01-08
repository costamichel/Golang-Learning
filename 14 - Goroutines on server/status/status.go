package main

import (
	"fmt"
	"log"
	"net/http"
	"sync"
	"time"
)

var (
	taskStatus     string = "Not started"
	taskStatusMu   sync.Mutex
	taskInProgress bool
)

func longTask() {
	taskStatusMu.Lock()
	taskStatus = "Started"
	taskInProgress = true
	taskStatusMu.Unlock()

	time.Sleep(10 * time.Second)

	taskStatusMu.Lock()
	taskStatus = "Finished"
	taskInProgress = false
	taskStatusMu.Unlock()
}

func taskHandler(w http.ResponseWriter, r *http.Request) {
	go longTask()
	fmt.Fprint(w, "Long task started. See status on /status")
}

func statusHandler(w http.ResponseWriter, r *http.Request) {
	taskStatusMu.Lock()
	defer taskStatusMu.Unlock()
	if taskInProgress {
		fmt.Fprintf(w, "Task in progress: %s", taskStatus)
	} else {
		fmt.Fprintf(w, "Task status: %s", taskStatus)
	}
}

func main() {
	http.HandleFunc("/task", taskHandler)
	http.HandleFunc("/status", statusHandler)

	fmt.Println("Server running on localhost:8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
