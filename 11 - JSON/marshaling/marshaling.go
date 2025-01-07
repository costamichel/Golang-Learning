package main

import (
	"encoding/json"
	"fmt"
)

type Pessoa struct {
	Nome  string `json:"name"`
	Idade int    `json:"age"`
}

func main() {
	p := Pessoa{Nome: "Michel", Idade: 32}

	jsonData, err := json.Marshal(p)
	if err != nil {
		fmt.Println("Erro ao converter struct para JSON")
		return
	}

	fmt.Println(string(jsonData))
}
