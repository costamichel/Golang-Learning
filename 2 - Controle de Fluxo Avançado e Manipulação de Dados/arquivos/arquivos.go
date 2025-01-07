package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	arquivo, err := os.Open("exemplo.txt")
	if err != nil {
		fmt.Println("Erro ao abrir o arquivo")
		return
	}
	defer arquivo.Close()

	scanner := bufio.NewScanner(arquivo)
	for scanner.Scan() {
		linha := scanner.Text()
		fmt.Println(linha)
	}

	if err := scanner.Err(); err != nil {
		fmt.Println("Erro ao ler o arquivo ", err)
	}
}
