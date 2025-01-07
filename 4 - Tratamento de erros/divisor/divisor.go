package main

import (
	"errors"
	"fmt"
)

func dividir(numerador, divisor float64) (float64, error) {
	if divisor == 0 {
		return 0, errors.New("Não é possível divisão por 0")
	}
	return numerador / divisor, nil
}

func main() {
	resultado, err := dividir(10, 0)
	if err != nil {
		fmt.Println("Erro -> ", err)
	} else {
		fmt.Println("Resultado -> ", resultado)
	}
}
