package main

import (
	"fmt"
	"time"
)

func main() {
	c := make(chan string)

	go func() {
		time.Sleep(2000 * time.Millisecond)
		c <- "Message received"
	}()

	select {
	case msg := <-c:
		fmt.Println(msg)
	case <-time.After(1 * time.Second):
		fmt.Println("Timeout!")
	}

	fmt.Println("End!")
}
