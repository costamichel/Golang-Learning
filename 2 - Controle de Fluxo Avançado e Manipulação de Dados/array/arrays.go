package main

import "fmt"

func main() {
	var arr [4]int
	arr[0] = 10
	fmt.Println(arr)

	arr2 := [3]string{"Sheldon", "Leonard", "Penny"}
	fmt.Println(arr2)
	fmt.Println(arr2[1])

	arr3 := [...]int{1, 2, 3, 5, 7, 11}
	fmt.Println(arr3)

	for i := 0; i < len(arr); i++ {
		fmt.Printf("Valor de arr na posição %d: %d\n", i, arr[i])
	}

	arr[1] = 7
	for idx, val := range arr {
		fmt.Printf("Iterando com range. Idx: %d. Val: %d.\n", idx, val)
	}
}
