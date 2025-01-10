package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"time"
)

func timeoutHandler(w http.ResponseWriter, r *http.Request) {
	ctx, cancel := context.WithTimeout(r.Context(), 2*time.Second)
	defer cancel()

	log.Println("Task started")

	select {
	case <-time.After(3 * time.Second):
		tmp_message := fmt.Sprintln("Task finished")
		fmt.Fprint(w, tmp_message)
		log.Println(tmp_message)
	case <-ctx.Done():
		tmp_message := fmt.Sprintf("Timeout error: %s", ctx.Err().Error())
		log.Println(tmp_message)
		http.Error(w, tmp_message, http.StatusServiceUnavailable)
	}

	fmt.Println("End of function")
}
