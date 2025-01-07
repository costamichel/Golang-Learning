package main

import (
	"fmt"
	"time"
)

func task1() {
	for i := 1; i <= 5; i++ {
		fmt.Printf("Task 1 iter %d\n", i)
		time.Sleep(700 * time.Millisecond)
	}
}

func task2() {
	for i := 1; i <= 5; i++ {
		fmt.Printf("Task 2 iter %d\n", i)
		time.Sleep(500 * time.Millisecond)
	}
}

func main() {
	go task1()
	go task2()

	time.Sleep(4 * time.Second)
	fmt.Println("End!")
}
