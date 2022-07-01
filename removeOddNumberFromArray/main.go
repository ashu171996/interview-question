// remove odd number from array without allocating extra variable
package main

import (
	"fmt"
)

func main() {
	input := []int{1, 2, 3, 4, 5, 6}
	fmt.Println("result", op(input))
}
func op(input []int) []int {
	low := 0
	high := len(input)
	for low < high {
		if input[low]%2 != 0 {
			input = append(input[:low], input[low+1:]...)
			high--
		} else {
			low++
		}
	}
	return input[:low]
}
