package main

import (
	"fmt"
	"sync"
	"time"
)

func downloadArquivo(nome string, tempo time.Duration, wg *sync.WaitGroup) {
	defer wg.Done()
	fmt.Printf("Baixando arquivo %s\n", nome)
	time.Sleep(tempo)
	fmt.Printf("Arquivo %s baixado em %v segundos\n", nome, tempo)
}

func main() {
	var wg sync.WaitGroup

	wg.Add(3)
	go downloadArquivo("arquivo 1", 3*time.Second, &wg)
	go downloadArquivo("arquivo 2", 5*time.Second, &wg)
	go downloadArquivo("arquivo 3", 2*time.Second, &wg)

	wg.Wait()
	fmt.Println("Todos os arquivos foram baixados!")
}
