package main

import (
	"fmt"
	"time"
)

func sendMessage(c chan string, message string, interval time.Duration) {
	for i := 0; i < 5; i++ {
		c <- fmt.Sprintf("%s %d", message, i)
		time.Sleep(interval)
	}

	close(c)
}

func main() {
	chan1 := make(chan string)
	chan2 := make(chan string)

	go sendMessage(chan1, "Globo", 700*time.Millisecond)
	go sendMessage(chan2, "SBT", 500*time.Millisecond)

	for i := 0; i < 200; i++ {
		select {
		case msg1, ok := <-chan1:
			if ok {
				fmt.Println(msg1)
			}
		case msg2, ok := <-chan2:
			if ok {
				fmt.Println(msg2)
			}
		default:
			fmt.Println("Nothing to receive. Waiting...")
			time.Sleep(900 * time.Millisecond)
		}
	}

	fmt.Println("End!")
}
