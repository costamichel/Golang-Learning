package main

import (
	"fmt"
	"os"
)

func main() {
	arquivo, err := os.Create("meu_arquivo.txt")
	if err != nil {
		fmt.Println("Erro ao criar arquivo!")
		return
	}
	defer arquivo.Close()

	_, err = arquivo.WriteString("Galopeeeeeeeeeeeeira\n")
	if err != nil {
		fmt.Println("Erro ao escrever no arquivo ", err)
		return
	}

	fmt.Println("Escrita efetuada com sucesso")
}
