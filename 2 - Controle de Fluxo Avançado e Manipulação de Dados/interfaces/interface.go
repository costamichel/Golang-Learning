package main

import (
	"fmt"
	"math"
)

type Forma interface {
	Area() float32
}

type Retangulo struct {
	Largura float32
	Altura  float32
}

func (r Retangulo) Area() float32 {
	return r.Largura * r.Altura
}

type Circulo struct {
	Raio float32
}

func (c Circulo) Area() float32 {
	return math.Pi * c.Raio * c.Raio
}

func mostrarArea(f Forma) {
	fmt.Printf("O valor da área é %.2f\n", f.Area())
}

func main() {
	r1 := Retangulo{
		Largura: 5,
		Altura:  1.2,
	}
	c1 := Circulo{
		Raio: 1.1,
	}

	mostrarArea(r1)
	mostrarArea(c1)
}
