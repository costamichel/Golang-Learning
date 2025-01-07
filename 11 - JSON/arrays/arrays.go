package main

import (
	"encoding/json"
	"fmt"
)

type Pessoa struct {
	Nome  string `json:"nome"`
	Idade int    `json:"idade"`
}

func main() {
	pessoas := []Pessoa{
		{Nome: "Michel", Idade: 32},
		{Nome: "Naira", Idade: 24},
	}
	fmt.Println("Serializando (Marshaling)")
	jsonData, err := json.Marshal(pessoas)
	if err != nil {
		fmt.Println("Erro ao serializar: ", err)
		return
	}
	fmt.Println(string(jsonData))
	fmt.Println("----------------------")

	fmt.Println("Desserializando (unmarshaling)")
	jsonStr := `[{"nome": "Zé Trovão", "idade": 33}, {"nome": "Ana Raio", "idade": 25}]`
	err = json.Unmarshal([]byte(jsonStr), &pessoas)
	if err != nil {
		fmt.Println("Erro ao desserializar: ", err)
		return
	}

	for i, p := range pessoas {
		fmt.Printf("Pessoa %d: %s tem %d anos\n", i, p.Nome, p.Idade)
	}
}
