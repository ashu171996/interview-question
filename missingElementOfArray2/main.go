package main

import (
	"fmt"
)

func main() {
	input := []int{2, 3, -7, 6, 8, 1, -10, 15}
	fmt.Println("missing element", missingElement(input))
	if missingElement(input) != 0 {
		fmt.Println("missing element", missingElement(input))
		return
	}
	fmt.Println("missing element not found")
}

func missingElement(input []int) int {
	ans := 1
	for _, v := range input {
		if v == ans {
			ans++
		}
	}
	return ans
}
