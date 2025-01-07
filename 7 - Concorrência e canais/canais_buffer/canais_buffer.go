package main

import "fmt"

func sendMessages(c chan string) {
	var text string
	for i := 0; i < 5; i++ {
		if i%2 == 0 {
			text = fmt.Sprintf("%d é par", i)
		} else {
			text = fmt.Sprintf("%d é ímpar", i)
		}
		fmt.Println("Sendig text -> ", text)
		c <- text
		fmt.Println("Text sended -> ", text)
	}
	close(c)
}

func main() {
	c := make(chan string, 3)

	go sendMessages(c)

	for text := range c {
		fmt.Println("Text received -> ", text)
	}
}
