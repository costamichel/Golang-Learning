package calculadora

import "errors"

func Soma(a, b int) int {
	return a + b
}

func Divisao(a, b float64) (float64, error) {
	if b == 0 {
		return 0, errors.New("Não é permitido divisão por 0")
	}
	return a / b, nil
}
