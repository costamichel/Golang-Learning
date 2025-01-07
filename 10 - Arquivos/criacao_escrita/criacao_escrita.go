package main

import (
	"fmt"
	"log"
	"os"
)

func main() {
	content := []byte("Hi, I'm line 1\nHet, call me line 2!")
	err := os.WriteFile("criado.txt", content, 0644)

	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("File created and writed successfull")
	fmt.Println("-----------------------------------")

	fmt.Println("Abrindo arquivo...")
	data, err := os.ReadFile("criado.txt")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Conteúdo do arquivo:")
	fmt.Println(string(data))
}
