package main

import (
	"errors"
	"fmt"
	"math/rand"
	"time"
)

func conectar() error {
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	numero_aleatorio := r.Float64()
	if numero_aleatorio < 0.15 { // 15% de probabilidade de falha
		return errors.New("Erro de conexão com o banco!")
	}
	return nil
}

func main() {
	for i := 0; i < 15; i++ {
		fmt.Print("\nConectando ao banco de dados ... ")
		time.Sleep(300 * time.Millisecond)
		err := conectar()
		if err != nil {
			fmt.Print("ERRO -> ", err)
		} else {
			fmt.Print("SUCESSO!")
		}
	}
}
