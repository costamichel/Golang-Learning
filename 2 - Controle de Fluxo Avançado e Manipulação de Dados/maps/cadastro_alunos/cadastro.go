package main

import (
	"fmt"
	"os"
	"os/exec"
	"strings"
)

func main() {
	alunos := make(map[string]float64)

	for {
		fmt.Printf("\n-----------------------------\n")
		fmt.Println("Escolha uma opção:")
		fmt.Printf("C - cadastrar aluno\t")
		fmt.Printf("E - exibir nota de aluno\n")
		fmt.Printf("T - mostrar todos os alunos e suas notas\t")
		fmt.Printf("D - deletar um aluno\n")
		fmt.Printf("S - sair\n")

		var input string
		fmt.Printf(" >> ")
		_, err := fmt.Scanln(&input)
		if err != nil {
			fmt.Println("Erro ao ler opção")
			continue
		} else {
			opcao := strings.ToUpper(input)
			switch opcao {
			case "C":
				var nome string
				var nota float64
				fmt.Printf("Informe o nome do aluno >> ")
				fmt.Scanln(&nome)
				fmt.Printf("Informe a nota do aluno >> ")
				fmt.Scanln(&nota)
				cadastrarAluno(alunos, nome, nota)
				fmt.Printf("Aluno '%s' adicionado\n", nome)
			case "E":
				var nome string
				fmt.Printf("Digite o nome do aluno que deseja ver a nota >> ")
				fmt.Scanln(&nome)
				exibeAluno(alunos, nome)
			case "T":
				fmt.Println("Mostrando todos os alunos")
				mostrarTodos(alunos)
			case "D":
				var nome string
				fmt.Printf("Informe o nome do aluno a ser removido >> ")
				fmt.Scanln(&nome)
				removeAluno(alunos, nome)
			case "S":
				fmt.Println("Saindo do programa...")
				return
			default:
				fmt.Printf("Opção %s é inválida!\n", opcao)
			}
		}
		fmt.Println("Aperte ENTER para continuar...")
		fmt.Scanln()
		clearScreen()
	}
}

func cadastrarAluno(mapa map[string]float64, nome string, nota float64) {
	mapa[nome] = nota
}

func exibeAluno(mapa map[string]float64, nome string) {
	nota, existe := mapa[nome]
	if existe {
		fmt.Printf("A nota de %s é %.2f\n", nome, nota)
	} else {
		fmt.Printf("Aluno %s não está cadastrado!\n", nome)
	}
}

func clearScreen() {
	var cmd *exec.Cmd

	switch os.Getenv("OS") {
	case "Windows_NT":
		cmd = exec.Command("cmd", "/c", "cls")
	default:
		cmd = exec.Command("clear")
	}

	cmd.Stdout = os.Stdout
	err := cmd.Run()
	if err != nil {
		fmt.Println("Erro ao limpar a tela:", err)
	}
}

func mostrarTodos(mapa map[string]float64) {
	for nome, nota := range mapa {
		fmt.Printf("%s tem nota %.2f\n", nome, nota)
	}
}

func removeAluno(mapa map[string]float64, nome string) {
	_, existe := mapa[nome]
	if existe {
		delete(mapa, nome)
		fmt.Printf("%s removido!\n", nome)
	} else {
		fmt.Printf("Aluno %s não encontrado\n", nome)
	}
}
