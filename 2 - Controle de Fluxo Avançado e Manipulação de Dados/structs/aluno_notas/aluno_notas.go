package main

import "fmt"

type Aluno struct {
	Nome  string
	Nota1 float32
	Nota2 float32
	Nota3 float32
}

func (a Aluno) Media() float32 {
	return (a.Nota1 + a.Nota2 + a.Nota3) / 3
}

func (a Aluno) Situacao() {
	media := a.Media()
	fmt.Printf("A média do aluno %s é %.2f\t", a.Nome, media)
	if media < 7 {
		fmt.Println("REPROVADO!")
	} else {
		fmt.Println("APROVADO")
	}
}

func main() {
	aluno1 := Aluno{
		Nome:  "Leonard",
		Nota1: 10.0,
		Nota2: 9.9,
		Nota3: 1.05,
	}
	aluno1.Situacao()
}
