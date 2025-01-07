package main

import "fmt"

type Transporte interface {
	Mover()
}

type Carro struct{}

func (c Carro) Mover() {
	fmt.Println("O carro se dirige")
}

type Bicicleta struct{}

func (b Bicicleta) Mover() {
	fmt.Println("A bicicleta se pedala")
}

func main() {
	c1 := Carro{}
	b1 := Bicicleta{}

	verLocomocao(c1)
	verLocomocao(b1)
}

func verLocomocao(t Transporte) {
	t.Mover()
}
