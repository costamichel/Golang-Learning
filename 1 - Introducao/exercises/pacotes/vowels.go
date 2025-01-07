package pacotes

import (
	"fmt"
	"strings"
)

func QtdVogais(str string) {
	vogais := [5]rune{'a', 'e', 'i', 'o', 'u'}

	str = strings.ToLower(str)

	contador := 0
	for _, char := range str {
		for _, v := range vogais {
			if char == v {
				contador++
			}
		}
	}

	fmt.Printf("A quantidade de vogais na string `%s` é %d\n", str, contador)
}
