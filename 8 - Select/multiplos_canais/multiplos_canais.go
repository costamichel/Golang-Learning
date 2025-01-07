package main

import (
	"fmt"
	"time"
)

func sendMessage(c chan string, message string, interval time.Duration) {
	for i := 0; i < 3; i++ {
		c <- fmt.Sprintf("%s %d", message, i)
		time.Sleep(interval)
	}

	close(c)
}

func main() {
	channel1 := make(chan string)
	channel2 := make(chan string)

	go sendMessage(channel1, "CANAL_1", 500*time.Millisecond)
	go sendMessage(channel2, "CANAL_2", 700*time.Millisecond)

	for i := 0; i < 25; i++ {
		select {
		case msg1, ok := <-channel1:
			if ok {
				fmt.Printf("Mensagem recebida de canal 1: %s\n", msg1)
			} else {
				fmt.Println("Canal 1 vazio")
			}
		case msg2, ok := <-channel2:
			if ok {
				fmt.Printf("Mensagem recebida de canal 2: %s\n", msg2)
			} else {
				fmt.Println("Canal 2 vazio")
			}
		default:
			fmt.Println("Nothing received. Waiting...")
			time.Sleep(300 * time.Millisecond)
		}
	}

	fmt.Println("End!")
}
