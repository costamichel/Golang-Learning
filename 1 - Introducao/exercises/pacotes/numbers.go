package pacotes

import "fmt"

func NumbersDiv3(max_val int, divisor int) {
	fmt.Printf("Calculando números inteiros divisíveis por %d de 1 até %d\n", divisor, max_val)
	contador := 0
	for i := 1; i <= max_val; i++ {
		if i%divisor == 0 {
			fmt.Printf("%d\t", i)
			contador += 1
		}
	}
	fmt.Printf("\nHá %d números divisíveis por %d de 1 até %d", contador, divisor, max_val)
}
