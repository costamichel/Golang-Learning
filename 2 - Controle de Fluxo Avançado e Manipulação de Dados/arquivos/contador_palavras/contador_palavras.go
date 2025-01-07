package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	arquivo, err := os.Open("texto.txt")
	if err != nil {
		fmt.Println("Erro ao abrir o arquivo -> ", err)
		return
	}
	defer arquivo.Close()

	scanner := bufio.NewScanner(arquivo)
	wordCount := 0
	for scanner.Scan() {
		line := scanner.Text()
		words := strings.Fields(line)
		wordCount += len(words)
	}

	if err := scanner.Err(); err != nil {
		fmt.Println("Erro ao ler arquivo -> ", err)
		return
	}

	fmt.Printf("Foram encontradas %d palavras!\n", wordCount)
}
