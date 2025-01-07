package main

import "fmt"

type Pessoa struct {
	Nome   string
	Idade  int
	Cidade string
}

func (p Pessoa) Exibir() {
	fmt.Printf("%s de %s, %d anos.\n", p.Nome, p.Cidade, p.Idade)
}

func main() {
	var base_dados []Pessoa
	for {
		fmt.Println("Seja bem-vindo ao cadastro de pessoas. Escolha uma opção abaixo:")
		fmt.Printf("1 - Cadastrar nova pessoa\t")
		fmt.Printf("2 - Exibir todas as pessoas\t")
		fmt.Printf("0 - Sair\n>> ")
		var input int
		_, err := fmt.Scanln(&input)
		if err != nil {
			fmt.Println("Erro ao ler valor inserido\n----------\n")
		} else {
			switch input {
			case 1:
				var nome, cidade string
				var idade int
				fmt.Printf("Digite o nome >> ")
				fmt.Scanln(&nome)
				fmt.Printf("Quantos anos tem? >> ")
				fmt.Scanln(&idade)
				fmt.Printf("Informe a cidade em que reside o(a) %s >> ", nome)
				fmt.Scanln(&cidade)
				base_dados = append(base_dados, Pessoa{Nome: nome, Idade: idade, Cidade: cidade})
				fmt.Printf("%s foi adicionado ao slice\n", base_dados[len(base_dados)-1].Nome)
			case 2:
				fmt.Println("Exibindo todos os contatos registrados!")
				for _, pessoa := range base_dados {
					pessoa.Exibir()
				}
			case 0:
				fmt.Println("Obrigado por usar o programa. Até mais!")
				return
			default:
				fmt.Printf("Opção inválida!\n---------\n")
			}
		}
	}
}
