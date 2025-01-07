package main

import (
	"fmt"
	"sync"
)

func processarDado(dado int, wg *sync.WaitGroup, resultados chan<- int) {
	defer wg.Done()
	resultado := dado * 2
	resultados <- resultado
}

func processarEmParalelo(dados []int) []int {
	var wg sync.WaitGroup

	resultados := make(chan int, len(dados))

	for _, dado := range dados {
		wg.Add(1)
		processarDado(dado, &wg, resultados)
	}

	wg.Wait()
	close(resultados)

	var resultadosFinal []int
	for resultado := range resultados {
		resultadosFinal = append(resultadosFinal, resultado)
	}

	return resultadosFinal
}

func main() {

	dados := []int{1, 2, 3, 4, 5}
	resultados := processarEmParalelo(dados)

	fmt.Println(resultados)
}
