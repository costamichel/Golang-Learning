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
	jsonStr := `{"name": "Michel", "age": 32}`

	var p Pessoa
	err := json.Unmarshal([]byte(jsonStr), &p)
	if err != nil {
		fmt.Println("Não foi possível desserializar o json")
		return
	}

	fmt.Printf("%s tem %d anos!", p.Nome, p.Idade)
}
