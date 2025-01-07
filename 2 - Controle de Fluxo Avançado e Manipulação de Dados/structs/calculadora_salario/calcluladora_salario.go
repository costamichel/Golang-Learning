package main

import "fmt"

type Funcionario struct {
	Nome             string
	SalarioBase      float32
	HorasTrabalhadas float32
}

func (f Funcionario) CalcularSalario() float32 {
	horasMensais := 220.0
	baseCalculo := f.HorasTrabalhadas / float32(horasMensais)
	return f.SalarioBase * baseCalculo
}

func main() {
	nome := "Michel"
	salarioBase := 3800.90
	horasTrabalhadas := 224.3

	funcionario1 := Funcionario{Nome: nome, SalarioBase: float32(salarioBase), HorasTrabalhadas: float32(horasTrabalhadas)}
	fmt.Println("O salário final de ", funcionario1.Nome, " é ", funcionario1.CalcularSalario())
}
