package main

import (
	"fmt"
	"strings"
)

func main() {
	frase := "Tifo tifo tifo, quantos tifo deu?"
	contagem := contarPalavras(frase)

	for palavra, quantidade := range contagem {
		fmt.Printf("%s aparece %d vez(es)\n", palavra, quantidade)
	}
}

func contarPalavras(documento string) map[string]int32 {
	tokenizado := strings.Fields(documento)
	contagem := make(map[string]int32)

	for _, palavra := range tokenizado {
		contagem[palavra]++
	}

	return contagem
}
