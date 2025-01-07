package main

import (
	"fmt"
	"os"
	"os/exec"
	"strings"
)

func main() {
	agenda := make(map[string]string)

	for {
		fmt.Printf("\n-------------------\n")
		fmt.Println("Escolha uma opção")
		fmt.Printf("A - Adicionar novo contato\t\t")
		fmt.Printf("U - Atualizar contato existente\n")
		fmt.Printf("E - Exibir contato específico\t\t")
		fmt.Printf("M - Mostrar todos os contatos\n")
		fmt.Printf("D - Deletar um contato\n")
		fmt.Printf("S - Sair\n")

		var input string
		fmt.Printf(" >> ")
		_, err := fmt.Scanln(&input)
		if err != nil {
			fmt.Println("Erro ao ler opção")
		} else {
			opcao := strings.ToUpper(input)
			switch opcao {
			case "A":
				fmt.Printf("Informe o nome do novo contato\n >> ")
				var nome, fone string
				fmt.Scanln(&nome)
				fmt.Println("Informe o telefone do contato\n >> ")
				fmt.Scanln(&fone)
				if !adicionarContato(agenda, nome, fone) {
					fmt.Println("Erro ao adicionar contato!")
				} else {
					fmt.Println("Novo contato adicionado na agenda")
				}
			case "U":
				fmt.Printf("Informe o nome do contato a ser atualizado\n >> ")
				var nome, fone string
				fmt.Scanln(&nome)
				_, exists := agenda[nome]
				if exists {
					fmt.Printf("Informe o novo telefone de %s\n >> ", nome)
					fmt.Scanln(&fone)
					if !atualizarContato(agenda, nome, fone) {
						fmt.Println("Erro ao atualizar contato")
					} else {
						fmt.Printf("Contato de %s foi atualizado\n", nome)
					}
				} else {
					fmt.Printf("%s não existe na agenda!\n", nome)
				}
			case "E":
				fmt.Printf("Informe o nome do contato a ser exibido\n >> ")
				var nome string
				fmt.Scanln(&nome)
				exibirContato(agenda, nome)
			case "M":
				mostrarTodos(agenda)
			case "D":
				fmt.Printf("Informe o nome do contato a ser removido\n >> ")
				var nome string
				fmt.Scanln(&nome)
				_, exists := agenda[nome]
				if exists {
					if !removerContato(agenda, nome) {
						fmt.Println("Erro ao remover contato")
					} else {
						fmt.Printf("%s foi removido da agenda\n", nome)
					}
				} else {
					fmt.Printf("%s não existe na agenda!\n", nome)
				}
			case "S":
				fmt.Println("Valeu! Até a próxima!")
				return
			default:
				fmt.Printf("%s não é uma opção válida!\n", opcao)
			}
		}
		fmt.Println("APERTE ENTER PARA CONTINUAR")
		fmt.Scanln()
		clearScreen()
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

func adicionarContato(agenda map[string]string, nome string, fone string) bool {
	_, exists := agenda[nome]
	if exists {
		fmt.Printf("%s já existe na agenda e não será adicionado!\n", nome)
		return false
	}
	agenda[nome] = fone
	return true
}

func atualizarContato(agenda map[string]string, nome string, fone string) bool {
	_, exists := agenda[nome]
	if !exists {
		fmt.Printf("%s não existe na agenda!\n", nome)
		return false
	}
	agenda[nome] = fone
	return true
}

func exibirContato(agenda map[string]string, nome string) {
	contato, exists := agenda[nome]
	if exists {
		fmt.Printf("Contato de %s -> %s\n", nome, contato)
	} else {
		fmt.Printf("Tá tentando colher onde não plantou né? Não tem %s na agenda\n", nome)
	}
}

func mostrarTodos(agenda map[string]string) {
	for nome, contato := range agenda {
		fmt.Printf("Contato de %s -> %s\n", nome, contato)
	}
}

func removerContato(agenda map[string]string, nome string) bool {
	_, exists := agenda[nome]
	if !exists {
		fmt.Printf("%s não existe na agenda!\n", nome)
		return false
	}
	delete(agenda, nome)
	return true
}
