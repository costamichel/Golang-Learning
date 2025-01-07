package main

import (
	"fmt"
	"sort"
	"strings"
)

type Pair struct {
	Palavra    string
	Frequencia int
}

func main() {
	frase := "O pelo do peito do pé pai do padre Pedro é preto. E quem disser que o pelo do peito do pé do pai do padre Pedro é preto, tem o pelo do peito do pé mais preto que o pai do padre Pedro."
	contagem := contarPalavras(frase)

	topPalavras := topNPalavras(contagem, 3)
	fmt.Println("Top 3 palavras mais frequentes")
	for _, pair := range topPalavras {
		fmt.Printf("%s aparece %d vezes\n", pair.Palavra, pair.Frequencia)
	}
}

func contarPalavras(documento string) map[string]int {
	tokenizado := strings.Fields(documento)
	contagem := make(map[string]int)

	for _, palavra := range tokenizado {
		contagem[palavra]++
	}

	return contagem
}

func topNPalavras(contagem map[string]int, n int) []Pair {
	var pares []Pair
	for palavra, frequencia := range contagem {
		pares = append(pares, Pair{Palavra: palavra, Frequencia: frequencia})
	}

	sort.Slice(pares, func(i, j int) bool {
		return pares[i].Frequencia > pares[j].Frequencia
	})

	if n > len(pares) {
		n = len(pares)
	}
	return pares[:n]
}
