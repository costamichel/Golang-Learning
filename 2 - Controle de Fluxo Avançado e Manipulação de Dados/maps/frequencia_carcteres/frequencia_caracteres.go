package main

import (
	"fmt"
)

func main() {
	texto := "Eu estou aprendendo maps no Go"
	frequencia := contarCaracteres(texto)

	for char, qtd := range frequencia {
		fmt.Printf("O caractere %c aparece %d vezes\n", char, qtd)
	}

}

func contarCaracteres(texto string) map[rune]int {
	frequencia := make(map[rune]int)

	for _, char := range texto {
		frequencia[char]++
	}

	return frequencia
}
