package main

import (
	"errors"
	"fmt"
	"os"
)

func abrirArquivo(nome string) error {
	arquivo, err := os.Open(nome)
	if err != nil {
		// %w mantém o erro original encapsulado no erro retornado
		return fmt.Errorf("Erro ao abrir o arquivo %s: %w", nome, err)
	}
	defer arquivo.Close()
	fmt.Println("Arquivo aberto com sucesso!")
	return nil
}

func main() {
	err := abrirArquivo("teste.txt")
	if err != nil {
		fmt.Println("Erro: ", err)
		if errors.Is(err, os.ErrNotExist) {
			fmt.Println("É erro de arquivo inexistente")
		}
	}
}
