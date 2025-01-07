package main

import (
	"bufio"
	"exercises/pacotes"
	"fmt"
	"os"
	"strings"
)

func main() {
	x := 5
	y := 3
	pacotes.Sum(x, y)
	pacotes.Sub(x, y)
	pacotes.Mul(x, y)
	pacotes.Div(x, y)

	reader := bufio.NewReader(os.Stdin)
	fmt.Printf("Digite uma palavra ou frase -> ")
	input, _ := reader.ReadString('\n')
	input = strings.TrimSuffix(input, "\r\n")
	pacotes.QtdVogais(input)

	pacotes.NumbersDiv3(100, 3)
}
