package main

import (
	"fmt"
	"os"
)

func main() {
	origem := "origem.txt"
	destino := "destino.txt"

	conteudo, err := os.ReadFile(origem)
	if err != nil {
		fmt.Println("Erro ao ler arquivo de origem -> ", err)
		return
	}

	err = os.WriteFile(destino, conteudo, 0644)
	if err != nil {
		fmt.Println("Erro ao gravar arquivo destino -> ", err)
		return
	}

	fmt.Printf("Cópia de %s para %s efetuada com sucesso!\n", origem, destino)
}
