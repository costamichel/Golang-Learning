package main

import "fmt"

type Operacao interface {
	Calcular() float32
}

type Soma struct {
	A, B float32
}

func (s Soma) Calcular() float32 {
	return s.A + s.B
}

type Subtracao struct {
	A, B float32
}

func (s Subtracao) Calcular() float32 {
	return s.A - s.B
}

type Multiplicacao struct {
	A, B float32
}

func (m Multiplicacao) Calcular() float32 {
	return m.A * m.B
}

type Divisao struct {
	A, B float32
}

func (d Divisao) Calcular() float32 {
	return d.A / d.B
}

func main() {
	x := float32(3)
	y := float32(2)

	soma := Soma{A: x, B: y}
	realizarCalculo(soma)

	subtracao := Subtracao{A: x, B: y}
	realizarCalculo(subtracao)

	multiplicacao := Multiplicacao{A: x, B: y}
	realizarCalculo(multiplicacao)

	divisao := Divisao{A: x, B: y}
	realizarCalculo(divisao)
}

func realizarCalculo(op Operacao) {
	res := op.Calcular()
	fmt.Println("Resultado: ", res)
}
