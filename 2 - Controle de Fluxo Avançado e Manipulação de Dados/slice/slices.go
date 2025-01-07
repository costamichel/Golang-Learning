package main

import (
	"fmt"
	"reflect"
)

func main() {
	var s []int
	s = []int{2, 4, 6, 8}
	fmt.Println(s)

	arr := [4]string{"Michelangelo", "Raphael", "Donatello", "Leonardo"}
	s2 := arr[1:3]
	fmt.Println(s2)

	s3 := make([]int, 5)
	fmt.Println(s3)

	s4 := make([]int, 3, 5)
	s4[1] = 11
	fmt.Println(s4)

	fmt.Println(arr)
	s2[0] = "RAPHAEL"
	fmt.Println(arr)

	fmt.Println(s4)
	s4 = append(s4, 20, 30, 40)
	fmt.Println(s4)

	for i := 0; i < len(s4); i++ {
		fmt.Printf("Iterando com `for` -> s4[%d] = %d\n", i, s4[i])
	}

	for index, value := range s4 {
		value = value + 1
		fmt.Printf("Iterando com range -> index = %d e value = %d\n", index, value)
	}
	fmt.Println(s4)

	for _, value := range s4 {
		fmt.Printf("Iterando com range sem index -> %d\n", value)
	}

	s5 := []int{1, 2, 3}
	s6 := []int{1, 2, 3}
	areEqual := true
	if len(s5) != len(s6) {
		areEqual = false
	} else {
		for i := range s5 {
			if s5[i] != s6[i] {
				areEqual = false
				break
			}
		}
	}
	fmt.Println("Comparação manual de igualdade -> ", areEqual)
	fmt.Println("Comparação de igualdade usando reflect.DeepEqual -> ", reflect.DeepEqual(s5, s6))
}
