package main

import (
	"fmt"
	"sync"
	"time"
)

func tarefa(nome string, wg *sync.WaitGroup) {
	defer wg.Done()
	fmt.Println("Iniciando ", nome)
	time.Sleep(2 * time.Second)
	fmt.Println("Terminando ", nome)
}

func main() {
	var wg sync.WaitGroup
	wg.Add(1)

	go tarefa("Task 1", &wg)
	wg.Add(1)
	go tarefa("Task 2", &wg)

	wg.Wait()
	fmt.Println("Fim do programa")
}
