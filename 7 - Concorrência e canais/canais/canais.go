package main

import "fmt"

func sendNumbers(c chan int) {
	for i := 0; i < 5; i++ {
		fmt.Println("Goroutine sending number ", i)
		c <- i
	}
	close(c)
}

func main() {
	c := make(chan int)

	go sendNumbers(c)

	for number := range c {
		fmt.Println("Main received number ", number)
	}

	fmt.Println("Channel -> ", c) // Mem address

	fmt.Println("End!")
}
