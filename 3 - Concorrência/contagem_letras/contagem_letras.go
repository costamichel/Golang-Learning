package main

import (
	"fmt"
	"sync"
	"unicode"
)

func contarLetras(texto string, wg *sync.WaitGroup, resultado chan<- map[rune]int) {
	defer wg.Done()

	contagem := make(map[rune]int)

	for _, char := range texto {
		if unicode.IsLetter(char) {
			contagem[char]++
		}
	}

	resultado <- contagem
}

func combinarResultados(resultado1, resultado2 map[rune]int) map[rune]int {
	resultadoFinal := make(map[rune]int)

	for letra, qtd := range resultado1 {
		resultadoFinal[letra] += qtd
	}
	for letra, qtd := range resultado2 {
		resultadoFinal[letra] += qtd
	}

	return resultadoFinal
}

func main() {
	texto := "Este é um exemplo de texto para contar as letras de forma concorrente!"

	metade := len(texto) / 2
	texto1 := texto[:metade]
	texto2 := texto[metade:]

	var wg sync.WaitGroup
	wg.Add(2)
	resultado := make(chan map[rune]int, 2)

	go contarLetras(texto1, &wg, resultado)
	go contarLetras(texto2, &wg, resultado)

	wg.Wait()
	close(resultado)

	contagem1 := <-resultado
	contagem2 := <-resultado

	resultadoFinal := combinarResultados(contagem1, contagem2)
	for letra, qtd := range resultadoFinal {
		fmt.Printf("%c: %d\n", letra, qtd)
	}
}
