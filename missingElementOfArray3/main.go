package main

import "fmt"

func main() {
	result := []int{}
	input := []int{1, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 17, 18, 19, 20, 21, 22, 23, 24, 25, 28, 29, 30, 31, 32, 33, 34, 35, 40, 41, 42, 43, 44, 45, 46, 48, 49, 50}
	for i := 0; i < len(input)-1; i++ {
		if input[i+1]-input[i] != 1 {
			result = append(result, input[i]+1)
		}
	}
	fmt.Println("ans", result)
}
