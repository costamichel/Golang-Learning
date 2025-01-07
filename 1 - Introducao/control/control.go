package main

import "fmt"

func main() {
	x := 2
	var y int = 3

	if x > y {
		fmt.Println("x é maior que y")
	} else {
		fmt.Println(x, "não é maior que", y)
	}

	for i := 0; i < 3; i++ {
		fmt.Println("Imprimindo do 'for' ...", i)
		if i == 1 {
			break
		}
	}

	op := "c"
	switch op {
	case "A":
		fmt.Println("Escolheste a opção A")
	case "B":
		fmt.Println("Escolheste o B")
	case "C":
		fmt.Println("Desta vez foi o C")
	default:
		fmt.Println("Nenhuma opção válida. Tem gente precisando trocar o óculos.")
	}

	j := 5
	for j < 10 {
		fmt.Println("Simulando o while que não existe no Go ...", j)
		j += 3
	}

}
