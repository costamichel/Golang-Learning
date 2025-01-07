package main

import (
	"fmt"
	"os"
	"time"
)

func main() {
	logFile := "log.txt"

	mensagem := fmt.Sprintf("Mensagem de log registrada em %s\n", time.Now().Format(time.RFC3339))

	file, err := os.OpenFile(logFile, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		fmt.Println("Erro ao abrir arquivo -> ", err)
		return
	}
	defer file.Close()

	if _, err := file.WriteString(mensagem); err != nil {
		fmt.Println("Erro ao gravar no arquivo de log => ", err)
		return
	}

	fmt.Println("Arquivo de log atualizado!")
}
