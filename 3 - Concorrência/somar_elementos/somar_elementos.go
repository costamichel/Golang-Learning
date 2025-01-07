package main

import (
	"fmt"
	"sync"
)

func main() {
	numbers := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 10}

	mid := len(numbers) / 2
	part1 := numbers[:mid]
	part2 := numbers[mid:]

	var wg sync.WaitGroup
	var sum1, sum2 int

	wg.Add(1)
	go func(nums []int) {
		defer wg.Done()
		for _, num := range nums {
			sum1 += num
		}
	}(part1)

	wg.Add(1)
	go func(nums []int) {
		defer wg.Done()
		for _, num := range nums {
			sum2 += num
		}
	}(part2)

	wg.Wait()

	total := sum1 + sum2

	fmt.Println("A soma dos elementos é ", total)
}
