package main

import (
	"fmt"
	"log"
	"os"
)

func main() {
	data, err := os.ReadFile("exemplo.txta")
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("File content:")
	fmt.Println(string(data))
}
