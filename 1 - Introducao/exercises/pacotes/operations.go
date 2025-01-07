package pacotes

import "fmt"

func Sum(x int, y int) {
	fmt.Printf("A soma é %d\n", x+y)
}

func Sub(x int, y int) {
	fmt.Printf("A diferença é %d\n", x-y)
}

func Div(x int, y int) {
	fmt.Printf("O resultado da divisão de %d por %d é %d\n", x, y, x/y)
}

func Mul(x int, y int) {
	fmt.Printf("O resultado da multiplicação é %d\n", x*y)
}
