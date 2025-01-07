package main

import "fmt"

func main() {
	var idades map[string]int32
	idades = make(map[string]int32)
	idades["Michel"] = 32
	idades["Sheldon"] = 40
	idades["Yamine Lamal"] = 18
	fmt.Println(idades)

	salarios := map[string]float32{
		"Neymar": 2000000.50,
		"Marta":  10000.30,
	}
	fmt.Println(salarios)

	idades["Rebecca"] = 26
	idadeSheldon := idades["Sheldon"]
	fmt.Println(idadeSheldon)
	idadeSheldon = 45
	fmt.Println(idades)

	delete(idades, "Yamine Lamal")
	fmt.Println(idades)

	idadeSilvio, existe := idades["Silvio"]
	if existe {
		fmt.Printf("A idade do Silvio existe e é %d\n", idadeSilvio)
	} else {
		fmt.Printf("Silvio não existe no mapa!\n")
	}

	for pessoa, idade := range idades {
		fmt.Printf("%s tem %d anos\n", pessoa, idade)
	}
}
