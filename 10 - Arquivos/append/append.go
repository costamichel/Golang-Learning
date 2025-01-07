package main

import (
	"fmt"
	"log"
	"os"
)

func main() {
	filePath := ".\\arquivo.txt"
	file, err := os.OpenFile(filePath, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		fmt.Println("Erro ao abrir o arquivo")
		log.Fatal(err)
	}
	defer file.Close()

	texto := "\nEsta é uma linha adicionada"
	_, err = file.WriteString(texto)
	if err != nil {
		fmt.Println("Erro ao escrever conteúdo no final do arquivo")
		log.Fatal(err)
	}
	fmt.Println("Conteúdo foi adicionado ao arquivo!")
}
