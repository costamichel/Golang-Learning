package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	arquivo, err := os.Open("exemplo.txt")
	if err != nil {
		fmt.Println("File cannot be opened")
		return
	}
	defer arquivo.Close()

	scanner := bufio.NewScanner(arquivo)
	i := 0
	for scanner.Scan() {
		fmt.Printf("%d: %s\n", i, scanner.Text())
		i++
	}

	if err := scanner.Err(); err != nil {
		fmt.Println("Error when reading file -> ", err)
	}
}
