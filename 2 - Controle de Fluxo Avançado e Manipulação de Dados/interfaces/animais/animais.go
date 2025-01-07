package main

import "fmt"

type Animal interface {
	Som()
}

type Cachorro struct{}

type Gato struct{}

func (g Gato) Som() {
	fmt.Println("Miau")
}

func (c Cachorro) Som() {
	fmt.Println("Au au")
}

type Pato struct{}

func (p Pato) Som() {
	fmt.Println("Quá quá")
}

func main() {
	meuCachorro := Cachorro{}
	meuGato := Gato{}
	meuPato := Pato{}

	emitirSom(meuGato)
	emitirSom(meuCachorro)
	emitirSom(meuPato)
}

func emitirSom(a Animal) {
	a.Som()
}
