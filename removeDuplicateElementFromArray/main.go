package main

import "fmt"

func main() {
	input := []int{1, 1, 2, 3, 2, 1, 5, 6, 4, 11, 12, 4}
	testMap := make(map[int]bool)
	result := []int{}
	for _, v := range input {
		if _, ok := testMap[v]; !ok {
			testMap[v] = true
			result = append(result, v)
		}
	}
	fmt.Println("result", result)
}
