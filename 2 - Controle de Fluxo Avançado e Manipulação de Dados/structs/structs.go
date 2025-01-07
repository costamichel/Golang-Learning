package main

import "fmt"

type Pessoa struct {
	Nome   string
	Idade  int
	Cidade string
}

func (p Pessoa) Saudacao() string {
	return fmt.Sprintf("Olá, me chamo %s e tenho %d anos!", p.Nome, p.Idade)
}

func main() {
	pessoa1 := Pessoa{
		Nome:   "Michel",
		Idade:  32,
		Cidade: "Nonoai",
	}

	fmt.Println("Nome: ", pessoa1.Nome)
	fmt.Println("Idade: ", pessoa1.Idade)
	fmt.Println("Cidade: ", pessoa1.Cidade)

	fmt.Println(pessoa1.Saudacao())
}
