package main

import (
	"fmt"
	"time"
)

func saudacao() {
	for i := 0; i < 3; i++ {
		fmt.Println("Olá")
		time.Sleep(500 * time.Millisecond)
	}
}

func despedida() {
	for i := 0; i < 3; i++ {
		fmt.Println("Até logo!")
		time.Sleep(700 * time.Millisecond)
	}
}

func main() {
	go saudacao()
	go despedida()

	time.Sleep(3 * time.Second)
	fmt.Println("Fim do programa")
}
