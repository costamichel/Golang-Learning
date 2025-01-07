package main

import "fmt"

// Função que retorna uma closure
func contador() func() int {
	i := 0
	// Retorna uma função anônima que incrementa e retorna 'i'
	return func() int {
		i++
		return i
	}
}

func main() {
	// Cria uma nova instância da closure
	incrementa := contador()

	fmt.Println(incrementa()) // Saída: 1
	fmt.Println(incrementa()) // Saída: 2
	fmt.Println(incrementa()) // Saída: 3

	// Outra instância independente da closure
	novoIncrementa := contador()
	fmt.Println(novoIncrementa()) // Saída: 1
}
