package main

import (
	"encoding/json"
	"fmt"
)

type Produto struct {
	Nome      string  `json:"nome"`
	Preco     float32 `json:"preco"`
	EmEstoque bool    `json:"em_estoque"`
}

func main() {
	p1 := Produto{Nome: "Picanha", Preco: 80.3, EmEstoque: true}

	jsonData, err := json.Marshal(p1)
	if err != nil {
		fmt.Println("Error when serializing json... -> ", err)
		return
	}
	fmt.Printf("JSON serialized!\n%s\n", string(jsonData))
	fmt.Println("-------------------------------")

	jsonStr := `{"nome": "Café", "preco": 10.5, "em_estoque": true}`
	var p2 Produto
	err = json.Unmarshal([]byte(jsonStr), &p2)
	if err != nil {
		fmt.Println("Error when deserializing: ", err)
		return
	}
	fmt.Printf("JSON deserialized!\n%+v\n", p2) //%+v imprime a struct com o nome dos campos
	fmt.Println("------------------------")

	jsonStr = `[
		{"nome": "Papel Higiênico", "preco": 30, "em_estoque": false},
		{"nome": "Chá verde", "preco": 5.1, "em_estoque": true}
	]`
	var produtos []Produto
	err = json.Unmarshal([]byte(jsonStr), &produtos)
	if err != nil {
		fmt.Println("Error when deserialized array: ", err)
		return
	}
	fmt.Printf("Array deserialized!\n%+v\n", produtos)
	fmt.Println("-------------------")

	produtos = append(produtos, p1, p2)
	fmt.Println("Final -> ", produtos)
}
