package main

import "fmt"

type Animal interface {
	Speak() string
}

type Dog struct{}
type Cat struct{}

func (d Dog) Speak() string {
	return "Woof"
}

func (c Cat) Speak() string {
	return "Meow"
}

func main() {
	var a Animal

	a = Dog{}
	fmt.Println(a.Speak())

	//Type assert
	//a = Cat{}
	v, ok := a.(Dog)
	if ok {
		fmt.Println("It is a Dog!")
		fmt.Println("Value: ", v)
	} else {
		fmt.Println("Sorry, it's not a Dog.")
	}

	a = Cat{}
	fmt.Println(a.Speak())
}
