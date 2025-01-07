package main

import "fmt"

func add(x int, y int) int {
	return x + y
}

func main() {
	sum := add(3, 4)
	fmt.Println(sum)

	mul := func(x int, y int) int {
		return x * y
	}
	fmt.Println(mul(5, 2))

	quociente, resto := divide(5, 2)
	fmt.Printf("Quociente é %d e resto é %d\n", quociente, resto)
}

func divide(x int, y int) (int, int) {
	return x / y, x % y
}
